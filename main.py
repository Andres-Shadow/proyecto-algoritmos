import tkinter as tk
from tkinter import ttk, messagebox
from PIL import Image, ImageTk
import subprocess
class MainGUI:
    def __init__(self, root):
        self.root = root
        self.root.title("Proyecto Final de Análisis de Algoritmos")
        self.initialize_interface()

    def initialize_interface(self):
        # Configuración de estilo
        style = ttk.Style()
        style.configure("TButton", padding=6, relief="flat", background="#3EAAFA", foreground="blue")

        # Marco principal
        main_frame = ttk.Frame(self.root, padding="20")
        main_frame.grid(row=0, column=0, sticky=(tk.W, tk.E, tk.N, tk.S))

        # Etiqueta de propósito
        label = ttk.Label(main_frame, text="Proyecto final de análisis de algoritmos", font=("Helvetica", 14))
        label.grid(row=0, column=0, columnspan=3, pady=(0, 10))

        # Botones en la misma línea
        execute_button = ttk.Button(main_frame, text="Ejecutar todos los algoritmos", command=self.execute_all_algorithms)
        execute_button.grid(row=1, column=0, padx=5)

        info_button = ttk.Button(main_frame, text="Mostrar Información", command=self.show_info)
        info_button.grid(row=1, column=1, padx=5)

        # Nombre del realizador
        realizador_label = ttk.Label(main_frame, text="Andrés Mauricio Dussán Bastidas \n Luis Fernano Arenas Ramirez", font=("Helvetica", 10, "bold"), anchor='center')
        realizador_label.grid(row=3, column=0, pady=(30, 0), sticky=tk.W)

    def execute_all_algorithms(self):
        comando = "go run main.go ui 1"

        # Ejecutar el comando y esperar a que termine
        resultado = subprocess.run(comando, shell=True)

        # Verificar si ocurrió algún error
        if resultado.returncode != 0:
            messagebox.showerror("Error", f"Error en la ejecución: {resultado.stderr}")
        else:
            # Mostrar la imagen después de que el comando haya terminado
            self.show_image("./output_barras.png")

    def show_info(self):
        self.show_image("./output_barras.png")
    def show_image(self, image_path):
        try:
            # Cargar la imagen
            img = Image.open(image_path)

            # Mostrar la imagen en una ventana emergente
            image_window = tk.Toplevel(self.root)
            image_window.title("Resultado de la Ejecución")
            
            tk_image = ImageTk.PhotoImage(img)
            label = tk.Label(image_window, image=tk_image)
            label.image = tk_image
            label.pack()

        except Exception as e:
            messagebox.showerror("Error", f"Error al mostrar la imagen: {str(e)}")

def main():
    root = tk.Tk()
    app = MainGUI(root)
    root.mainloop()

if __name__ == "__main__":
    main()