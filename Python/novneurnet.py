import numpy, random
learning_rate = 1 #learning rate
bias = 1 

#2 neurons + bias = 3 weights
weights = [random.random(),random.random(),random.random()] 

def Perceptron(input1, input2, output) :
    percep_out = input1*weights[0]+input2*weights[1]+bias*weights[2]
    if percep_out > 0 :
        percep_out = 1
    else :
        percep_out = 0
    error = output - percep_out
    weights[0] += error * input1 * learning_rate
    weights[1] += error * input2 * learning_rate
    weights[2] += error * bias * learning_rate

for i in range(50) :
    Perceptron(1,1,1)
    Perceptron(1,0,1)
    Perceptron(0,1,1)
    Perceptron(0,0,0)

x = int(input())
y = int(input())
percep_out = x*weights[0] + y*weights[1] + bias*weights[2]
if percep_out > 0 :
    percep_out = 1
else :
    percep_out = 0
print(x, "or", y, "is : ", percep_out)

percep_out = 1/(1+numpy.exp(-percep_out))
print("Sigmoid output : ", percep_out)
print("Weights : ", weights)