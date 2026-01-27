import numpy as np

arr_0d = np.array(54)

print(arr_0d)


arr_1d = np.array((54,32,12,46,56))

print(arr_1d)

arr_2d = np.array(([54,32,12,46,12], [12,34,21,56,21]))

print(arr_2d)

arr_3d = np.array([[12,45,21,52,2], [12,4,2,12,4],[1,4,12,42,1]])

print(arr_3d)

print(arr_0d.ndim)
print(arr_1d.ndim)
print(arr_2d.ndim)
print(arr_3d.ndim)

arr_5d = np.array([1,2,3,4,5,6], ndmin=5)

print(arr_5d)
print(arr_5d.ndim)


