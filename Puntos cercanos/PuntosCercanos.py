from tkinter import *
    
window = Tk()
points = []

class Point():

    x = 0
    y = 0
    
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def draw(self):
        canvas.create_oval(self.x,self.y,self.x,self.y, width = 10)
        

def createPoint(event):
    print( "clicked at", event.x, event.y)
    point = Point(event.x, event.y)
    point.draw()
    points.append(point)
    
    

def closestTwoPoints():

    pointsX = []
    for point in points:
        pointsX.append(point.x)
        print(point.x)
        

    pointsX.sort()
    print(pointsX)

    #Dividir el conjunto de puntos en dos partes iguales

    for 

canvas = Canvas(window, bg = "white", width =  400, height = 400)

button = Button(canvas, text = "Ordenar", command = closestTwoPoints)
button.place(x =0, y = 0)

canvas.bind("<Button-1>", createPoint)
canvas.pack()

window.mainloop()


        

