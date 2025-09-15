# Gunakan image Python dari Docker Hub
FROM python:3

# Set direktori kerja
WORKDIR /usr/src/app

# Salin file requirements.txt ke dalam kontainer
COPY requirements.txt .

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Salin semua file dalam direktori proyek ke dalam kontainer
COPY . .

# Expose port yang akan digunakan oleh Jupyter Notebook
EXPOSE 8888

# Perintah untuk menjalankan Jupyter Notebook
CMD ["jupyter", "notebook", "--ip='*'", "--port=8888", "--no-browser", "--allow-root"]