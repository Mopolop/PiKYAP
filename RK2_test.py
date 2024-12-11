import unittest
from operator import itemgetter
from RK2_refactoring import ( # Замените your_module на имя вашего файла
    Brw, Cmp, CmpBrw, get_one_to_many,
    calculate_total_memory, get_browsers_per_computer
)


class TestDataProcessing(unittest.TestCase):

    def setUp(self):
        self.cmps = [
            Cmp(1, 'Alexei','ноутбук','Windows 11'),
            Cmp(2, 'PC Artem','ультрабук','Windows 10'),
            Cmp(3, 'Dmitriy','настольный','Linux')
        ]
        self.brws = [
            Brw(1, 200,'Yandex',  'Яндекс',1),
            Brw(2, 350,'Chrome',  'Google',2)
        ]
        self.cmps_brws = [CmpBrw(1, 1), CmpBrw(2, 2)]

    def test_get_one_to_many(self):
        expected = [('Yandex', 200, 'Alexei'), ('Chrome', 350, 'PC Artem')]
        result = get_one_to_many(self.brws, self.cmps)
        self.assertEqual(result, expected)


    def test_calculate_total_memory(self):
        one_to_many_data = [('Yandex', 200, 'Alexei'), ('Chrome', 350, 'Alexei')]
        expected = [('Alexei', 550)]
        result = calculate_total_memory(one_to_many_data, self.cmps)
        self.assertEqual(result, expected)

    def test_get_browsers_per_computer(self):
        many_to_many_data = [('Yandex', 200, 'Alexei'), ('Chrome', 350, 'Alexei')]
        expected = {'Alexei': ['Yandex', 'Chrome'], 'Dmitriy': [], 'PC Artem': []}
        result = get_browsers_per_computer(many_to_many_data, self.cmps)
        self.assertEqual(result, expected)

if __name__ == '__main__':
    unittest.main()
