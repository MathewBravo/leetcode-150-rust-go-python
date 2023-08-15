fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
    let (mut p1, mut p2, mut p3) = (m - 1, n - 1, m + n - 1);

    while p1 >= 0 && p2 >= 0 {
        if nums1[p1 as usize] >= nums2[p2 as usize] {
            nums1[p3 as usize] = nums1[p1 as usize];
            p1 -= 1;
        } else {
            nums1[p3 as usize] = nums2[p2 as usize];
            p2 -= 1;
        }
        p3 -= 1;
    }

    while p2 >= 0 {
        nums1[p3 as usize] = nums2[p2 as usize];
        p2 -= 1;
        p3 -= 1;
    }
}

#[cfg(test)]
mod tests {
    use super::merge;

    #[test]
    fn test_merge_nums1_empty() {
        let mut nums1 = vec![0];
        let m = 0;
        let mut nums2 = vec![1];
        let n = 1;
        merge(&mut nums1, m, &mut nums2, n);
        assert_eq!(nums1, vec![1]);
    }

    #[test]
    fn test_merge_nums2_empty() {
        let mut nums1 = vec![1];
        let m = 1;
        let mut nums2 = vec![];
        let n = 0;
        merge(&mut nums1, m, &mut nums2, n);
        assert_eq!(nums1, vec![1]);
    }

    #[test]
    fn test_merge() {
        let mut nums1 = vec![1, 2, 3, 0, 0, 0];
        let m = 3;
        let mut nums2 = vec![2, 5, 6];
        let n = 3;
        merge(&mut nums1, m, &mut nums2, n);
        assert_eq!(nums1, vec![1, 2, 2, 3, 5, 6]);
    }
}
