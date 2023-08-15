use std::collections::HashMap;

fn majority_element(nums: Vec<i32>) -> i32 {
    let mut count_map: HashMap<i32, i32> = HashMap::new();
    for num in nums {
        *count_map.entry(num).or_insert(0) += 1;
    }

    let (key_of_max_value, _max_value) = count_map
        .iter()
        .max_by(|a, b| a.1.cmp(&b.1))
        .unwrap_or((&0, &0));

    return *key_of_max_value;
}
