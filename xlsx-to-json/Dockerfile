FROM python

WORKDIR /app

COPY requirements.txt ./

RUN pip install -r requirements.txt

VOLUME /app/data

COPY . .

CMD [ "python3", "main.py" ]