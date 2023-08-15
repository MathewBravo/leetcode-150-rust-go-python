fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
    let mut ins_idx = 0;
    for i in 0..nums.len() {
        if nums[i] != val {
            nums[ins_idx] = nums[i];
            ins_idx += 1;
        }
    }
    ins_idx as i32
}

#[cfg(test)]
mod tests {
    use super::remove_element;
    #[test]
    fn remove_element_works() {
        let mut test_vec = vec![0, 1, 2, 2, 3, 0, 4, 2];
        let result = remove_element(&mut test_vec, 2);
        assert_eq!(test_vec[0..5], vec![0, 1, 3, 0, 4]);
        assert_eq!(result, 5);
    }
}
