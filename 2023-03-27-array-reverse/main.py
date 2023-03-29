import unittest


# def reverse(arr):
#     res = []
#     reverseRecurse(arr, res)
#     return res


def reverse(arr):
    res = []
    for i in reversed(arr):
        if isinstance(i, list):
            res.append(reverse(i))
        else:
            res.append(i)
    return res


class TestReverse(unittest.TestCase):
    def test_reverse(self):
        self.assertEqual(reverse([1, 2, 3, 4, 5]), [5, 4, 3, 2, 1])
        self.assertEqual(reverse([1, [2, 3, [4, 5]]]), [[[5, 4], 3, 2], 1])
        self.assertEqual(reverse([1, [2, [3, [4, 5], 6]]]), [[[6, [5, 4], 3], 2], 1])


# run test code for main.py
if __name__ == '__main__':
    unittest.main()
