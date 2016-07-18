import numpy, collections
from matplotlib import pyplot as plt 

data = numpy.genfromtxt('records.csv', delimiter=',', dtype=None)

urls = data[:,0]
times = data[:,1]

def time_collector():
	return collections.defaultdict(int)

def string_to_time(time):
	times = time.split(":")
	minutes = int(times[0])
	seconds = int(times[1])
	return 60*minutes + seconds

processed_data = collections.defaultdict(time_collector)

for i, url in enumerate(urls):
	processed_data[url][times[i]] +=1


for url, key_values in processed_data.iteritems():
	tuples = sorted(key_values.iteritems())
	video_times, counts = zip(*tuples)
	video_seconds = [string_to_time(time) for time in video_times]

	# import ipdb; ipdb.set_trace()

	plt.plot(video_seconds, counts)
	plt.xticks(video_seconds, video_times, rotation='vertical')
	axes = plt.gca()
	axes.set_ylim([0,5])
	plt.show()

