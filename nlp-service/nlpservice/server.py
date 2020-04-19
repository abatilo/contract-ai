from concurrent import futures
import grpc
import nlp_pb2
import nlp_pb2_grpc
import logging
from allennlp.predictors.predictor import Predictor

logging.basicConfig(level=logging.NOTSET)
logger = logging.getLogger("nlpservice")

predictor = Predictor.from_path("https://storage.googleapis.com/allennlp-public-models/bidaf-elmo-model-2018.11.30-charpad.tar.gz")

class NLPService(nlp_pb2_grpc.NLPServiceServicer):

  def AskQuestion(self, request, context):
    logger.info("Received a question")
    output = predictor.predict(
      passage=request.corpus.body,
      question=request.question.body
    )
    logger.info("Completed prediction")
    return nlp_pb2.AskQuestionResponse(answer=nlp_pb2.Answer(body=output["best_span_str"]), question=request.question)

  def AskQuestions(self, request, context):
    answers = []
    for question in request.questions:
      output = predictor.predict(
        passage=request.corpus.body,
        question=question.body
      )
      answers += [nlp_pb2.Answer(body=output["best_span_str"])]
    return nlp_pb2.AskQuestionsResponse(answers=answers)

def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  nlp_pb2_grpc.add_NLPServiceServicer_to_server(
    NLPService(), server)
  server.add_insecure_port("[::]:9090")
  logger.info("Starting server...")
  server.start()
  server.wait_for_termination()

if __name__ == '__main__':
  serve()
