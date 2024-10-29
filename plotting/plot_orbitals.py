import pandas
import pathlib
import matplotlib.pyplot as plt

csv = pathlib.Path(__file__).parents[1].joinpath("main", "data.csv")

data = pandas.read_csv(f"{csv}", header=None)

plt.plot(data[data[0] == "B"][1], data[data[0] == "B"][2], "k-")
plt.plot(data[data[0] == "C"][1], data[data[0] == "C"][2], "r-")
plt.plot(data[data[0] == "D"][1], data[data[0] == "D"][2], "b-")
plt.savefig(f"{pathlib.Path.cwd().joinpath('orbitals.png')}")