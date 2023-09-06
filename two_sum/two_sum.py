from typing import List

class Solution:
    def two_sum(self, nums: List[int], target: int) -> List[int]:
        result = []
        find_target = False

        for ind1 in range(len(nums)):
            if find_target:
                break
            for ind2 in range(len(nums)):
                if find_target:
                    break
                if (ind1 != ind2):
                    if (nums[ind1] + nums[ind2] == target):
                        result.append(ind1)
                        result.append(ind2)
                        find_target = True
        
        return result

    def two_sum_2(self, nums: List[int], target: int) -> List[int]:
        find_target = False
        ind1 = 0
        last_ind = len(nums) - 1

        while ind1 <= last_ind:
            ind2 = last_ind

            while ind2 > ind1:
                if (nums[ind1] + nums[ind2] == target):
                    find_target = True
                    break
                ind2 -= 1
            if find_target:
                break
            ind1 += 1
        
        return [ind1, ind2]

    def two_sum_3(self, nums: List[int], target: int) -> List[int]:
        ind = 0
        last_ind = len(nums) - 1
        hashmap = {}

        while ind <= last_ind:
            hashmap.update({(target - nums[ind]):ind})
            ind += 1

        ind1 = 0
        ind2 = 0
        while ind1 <= last_ind:
            val = hashmap.get(nums[ind1])
            if val != None and val != ind1:
                ind2 = val
                break
            ind1 += 1
        
        return [ind1, ind2]

    def two_sum_final(self, nums: List[int], target: int) -> List[int]:
        hashmap = {}
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in hashmap:
                return [i, hashmap[complement]]
            hashmap[nums[i]] = i

print("---first---")
print(Solution.two_sum(1, [2,7,11,15], 9))
print(Solution.two_sum(1, [3,2,4], 6))
print(Solution.two_sum(1, [3,3], 6))
print("---second---")
print(Solution.two_sum_2(1, [2,7,11,15], 9))
print(Solution.two_sum_2(1, [3,2,4], 6))
print(Solution.two_sum_2(1, [3,3], 6))
print("---third---")
print(Solution.two_sum_3(1, [2,7,11,15], 9))
print(Solution.two_sum_3(1, [3,2,4], 6))
print(Solution.two_sum_3(1, [3,3], 6))