# -*- coding: utf-8 -*-
import unittest
from gilded_rose_refactored import update_quality, Item


class GildedRoseTest(unittest.TestCase):
    def test_aged_brie_quality_increases_by_two_after_sell_date(self):
        items = [Item("Aged Brie", -1, 47)]
        update_quality(items)
        self.assertEqual(49, items[0].quality)
        
    def test_aged_brie_before_sell_date_increase_by_one(self):
        item = [Item("Aged Brie", 10, 40)]
        update_quality(item)
        self.assertEqual(41, item[0].quality)
        
    def test_backstage_pass_after_concert(self):
        items = [Item('Backstage passes',-1,49)]
        update_quality(items)
        self.assertEqual(0, items[0].quality)
        
    def test_backstage_pass_10_days_before_concert(self):
        items = [Item('Backstage passes',10,45)]
        update_quality(items)
        self.assertEqual(47, items[0].quality)

    def test_backstage_pass_less_than_5_days_before_concert(self):
        items = [Item('Backstage passes',5,45)]
        update_quality(items)
        self.assertEqual(48, items[0].quality) 
            
    def test_sulfuras_never_changes(self):
        items = [Item('Sulfuras',10,40)]
        update_quality(items)
        self.assertEqual((10,40),(items[0].sell_in, items[0].quality))
        
    def test_after_sell_date_quality_decreases_by_2(self):
        items = [Item('something', -1, 35)]
        update_quality(items)
        self.assertEqual((-2,33), (items[0].sell_in, items[0].quality))
        
    def test_conjured_before_sell_date(self):
        items = [Item('Conjured', 10, 20)]
        update_quality(items)
        self.assertEqual(18,items[0].quality)
        


if __name__ == "__main__":
    unittest.main()
