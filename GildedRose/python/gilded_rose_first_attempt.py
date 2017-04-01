#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Wed Mar 22 16:08:34 2017

@author: johnyearsley
"""


def update_quality(items):
    for item in items:
        
        past_sell_date = is_past_sell_date(item)
        increasing_type = is_aged_brie(item)
        backstage_type = is_backstage_pass(item)
        conjured_type = is_conjured(item)
        legendary_type = is_sulfuras(item)
        
        item.sell_in -= 1
        
        if legendary_type:
            item.sell_in += 1
        
        elif increasing_type:
            if past_sell_date:
                item.quality += 2
            
            else:
                item.quality += 1
                
        elif conjured_type:
            if past_sell_date:
                item.quality -= 4
            
            else:
                item.quality -= 2
                
        elif backstage_type:
            if 11 <= item.sell_in <50:
                item.quality += 1
                
            elif 5 < item.sell_in < 11:
                item.quality += 2
                
            elif 0 < item.sell_in <= 5:
                item.quality += 3
                
            elif past_sell_date:
                item.quality = 0
                
        else:
            if past_sell_date:
                item.quality -= 2
                
            else:
                item.quality -= 1
                
class Item:
    def __init__(self, name, sell_in, quality):
        self.name = name
        self.sell_in = sell_in
        self.quality = quality

    def __repr__(self):
        return "%s, %s, %s" % (self.name, self.sell_in, self.quality)
    
def is_past_sell_date(item):
    if item.sell_in < 0:
        return True
    
    else:
        return False
    
def is_aged_brie(item):
    if item.name == "Aged Brie":
        return True
    
    else:
        return False
       
def is_backstage_pass(item):
    if item.name == "Backstage passes":
        return True
    
    else:
        return False
        
def is_conjured(item):
    if "Conjured" in item.name:
        return True
    
    else:
        return False
    
def is_sulfuras(item):
    if item.name == "Sulfuras":
        return True
    
    else:
        return False