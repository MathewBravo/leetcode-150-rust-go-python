use std::collections::HashMap;

fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut contained: HashMap<i32, usize> = HashMap::new();
    for (i, num) in nums.iter().enumerate() {
        if contained.contains_key(&(target - num)) {
            return vec![
                *contained.get(&(target - num)).unwrap() as i32,
                i as i32,
            ];
        } else {
            contained.insert(*num, i);
        }
    }

    return vec![-1, -1];
}

#[cfg(test)]
mod tests {
    use super::two_sum;

    #[test]
    fn two_sum_works() {
        let test_v1 = vec![2, 7, 11, 15];
        let test_v1_t = 9;
        let test_v2 = vec![3, 2, 4];
        let test_v2_t = 6;
        assert_eq!(two_sum(test_v1, test_v1_t), vec![0, 1]);
        assert_eq!(two_sum(test_v2, test_v2_t), vec![1, 2]);
    }
}
