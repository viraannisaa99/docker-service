# Gunakan image dasar dengan Conda
FROM continuumio/miniconda3

# Set working directory di dalam container
WORKDIR /app

# Salin file environment.yml ke container
COPY environment.yml .

# Buat environment conda berdasarkan file environment.yml
RUN conda env create -f environment.yml

# Aktifkan environment yang telah dibuat
SHELL ["conda", "run", "-n", "myenv", "/bin/bash", "-c"]

# Install Streamlit
RUN pip install streamlit

# Salin kode aplikasi Streamlit Anda ke dalam container
COPY . .

# Tentukan perintah yang dijalankan saat container dijalankan
CMD ["streamlit", "run", "app.py"]
