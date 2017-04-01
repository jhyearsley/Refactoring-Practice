def update_quality(items):

    for item in items:
        legendary = is_legendary(item)
        
        if not legendary:
            item.sell_in -= 1        
            compute_difference(item)


class Item:
    def __init__(self, name, sell_in, quality):
        self.name = name
        self.sell_in = sell_in
        self.quality = quality

    def __repr__(self):
        return "%s, %s, %s" % (self.name, self.sell_in, self.quality)
                
def compute_difference(item):
    if item.name not in ["Aged Brie", "Backstage passes", "Conjured"]:
        return default_strategy(item)
    
    if item.name == "Sulfuras":
        return legendary_strategy(item)
    
    if item.name == "Aged Brie":
        return aged_brie_strategy(item)
    
    if item.name == "Backstage passes":
        return backstage_strategy(item)
    
    if "Conjured" in item.name:
        return conjured_strategy(item)
    
def default_strategy(item):
    if item.sell_in < 0:
        item.quality -= 2
    
    else:
        item.quality -= 1
                
def legendary_strategy(item):
    pass

def aged_brie_strategy(item):
    if item.sell_in < 0:
        item.quality += 2
            
    else:
        item.quality += 1
        
def backstage_strategy(item):
    if 11 <= item.sell_in <50:
        item.quality += 1
                
    elif 5 < item.sell_in < 11:
        item.quality += 2
                
    elif 0 < item.sell_in <= 5:
        item.quality += 3
            
    elif item.sell_in <= 0:
        item.quality = 0 

def conjured_strategy(item):
    if item.sell_in >= 0:
        item.quality -= 2
        
    else:
        item.quality -= 4
    
def is_legendary(item):
    legendary_items = ["Sulfuras"]
    if item.name in legendary_items:
        return True
    
    else:
        return False