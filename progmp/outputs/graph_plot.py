import pandas as pd
import matplotlib.pyplot as plt

df = pd.read_csv('data.txt')

print(df)
# df.plot(y=['KB/s in','KB/s in.1'])
# df.plot(y=['KB/s in'])
print(df['scheduler'])
index = df['scheduler'].tolist()
print(index)
df['time'] = df['time'].astype(float)
time = df['time'].tolist()
df = pd.DataFrame({'time': time}, index=index)
ax =df.plot.bar()
plt.xticks(fontsize=6, rotation=15)
# ax.tick_params(axis='both', which='minor', labelsize=1)
# plt.show()
plt.show(ax)