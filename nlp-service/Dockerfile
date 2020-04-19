FROM python:3.7.3-slim AS builder

WORKDIR /src/

RUN apt-get update \
      && apt-get install --no-install-recommends -yq \
      make=4.1-9.1 \
      build-essential=12.3 \
      && apt-get clean \
      && rm -rf /var/lib/apt/lists/*

# Hack to install allennlp because it breaks poetry
# https://github.com/python-poetry/poetry/issues/2094
COPY ./requirements.txt /src/
RUN pip install -r requirements.txt

COPY ./pyproject.toml ./poetry.lock /src/
RUN pip install poetry==0.12.11 && \
	poetry install

COPY ./nlpservice /src/nlpservice
RUN poetry build
RUN pip install .

FROM python:3.7.3-slim

COPY --from=builder /usr/local/bin/entrypoint /usr/local/bin/
COPY --from=builder /usr/local/lib /usr/local/lib

# Download spacy model
RUN python -m spacy download en_core_web_sm
RUN python -c 'from allennlp.predictors.predictor import Predictor; predictor = Predictor.from_path("https://storage.googleapis.com/allennlp-public-models/bidaf-elmo-model-2018.11.30-charpad.tar.gz")';

ENV PYTHONPATH /usr/local/lib/python3.7/site-packages/nlpservice
CMD ["entrypoint"]
