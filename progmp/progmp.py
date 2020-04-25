'''
@author: Alexander Froemmgen
'''

import socket
from struct import *

class ProgMP:
	@staticmethod
	def MPTCP_RBS_SCHEDULER():
		return 44
	
	@staticmethod
	def MPTCP_RBS_REGISTER():
		return 45
	
	@staticmethod
	def MPTCP_RBS_SKB_PROPERTIES():
		return 46
		
	@staticmethod
	def R1():
		return 0
	
	@staticmethod
	def R2():
		return 1
		
	@staticmethod
	def R3():
		return 2

	@staticmethod
	def R4():
		return 3
		
	@staticmethod
	def R5():
		return 4
	
	@staticmethod
	def R6():
		return 5
		
	@staticmethod
	def getSchedulerName(schedulerFileName):
		with open(schedulerFileName, "r") as src:
			content = src.read()
			prefix = "SCHEDULER "
			startIndex = content.index(prefix) + len(prefix)
			endIndex = content[startIndex:].index(";")
			return content[startIndex:startIndex + endIndex]
		
	@staticmethod
	def loadScheduler(schedulerStr):
		with open("/proc/net/mptcp_net/rbs/schedulers", "w") as dst:
			dst.write(schedulerStr)
			
	@staticmethod
	def setScheduler(s, schedulerName):
		s.setsockopt(socket.IPPROTO_TCP, Progmp.MPTCP_RBS_SCHEDULER(), schedulerName)
		
	@staticmethod
	def setDefaultScheduler(schedulerName):
		with open("/proc/net/mptcp_net/rbs/default", "w") as dst:
			dst.write(schedulerName)

	@staticmethod
	def setRegister(s, register, value):
		# Pack it accordingly to the c struct and therefore import struct above."
		sValue = pack('II', register, value)
		s.setsockopt(socket.IPPROTO_TCP, Progmp.MPTCP_RBS_REGISTER(), sValue)
	
	@staticmethod
	def setUser(s, value):
		# Pack it accordingly to the c struct and therefore import struct above."
		sValue = pack('I', value)
		s.setsockopt(socket.IPPROTO_TCP, Progmp.MPTCP_RBS_SKB_PROPERTIES(), sValue)
	
	@staticmethod
	def removeScheduler(schedulerName):
		with open("/proc/net/mptcp_net/rbs/" + schedulerName + "/delete", "w") as dst:
			dst.write("1")
	
	@staticmethod
	def isMultipath():
		pass