advent_of_code::solution!(4);

fn count_shared_elements<T: PartialEq>(vec1: &[T], vec2: &[T]) -> u32 {
    vec1.iter()
        .filter(|&x| vec2.contains(x))
        .count()
        .try_into()
        .unwrap()
}

fn get_matches<'a>(line: &'a str, winning: &mut Vec<&'a str>) -> usize {
    let mut numbers = line.split_whitespace().skip(2);
    winning.clear();
    while let Some(x) = numbers.next() {
        if x == "|" {
            break;
        }

        winning.push(x);
    }

    numbers.filter(|x| winning.contains(x)).count()

}

pub fn part_one(input: &str) -> Option<u32> {
    let total_shared_count: u32 = input
        .lines()
        .map(|line| {
            let collection = line.split(":").collect::<Vec<&str>>();
            let numbers = collection[1].split("|").collect::<Vec<&str>>();

            let winning_numbers: Vec<i32> = numbers[0]
                .split_whitespace()
                .filter_map(|s| s.parse().ok())
                .collect();

            let self_numbers: Vec<i32> = numbers[1]
                .split_whitespace()
                .filter_map(|s| s.parse().ok())
                .collect();

            let shared_count = count_shared_elements(&winning_numbers, &self_numbers);
            if shared_count == 0 {
                0
            } else {
                2u32.pow(shared_count - 1)
            }
        })
        .sum();

    return Some(total_shared_count);
}

pub fn part_two(input: &str) -> Option<u32> {
    let mut card_counts = vec![1u32];
    let mut count = 0;
    let mut winning = Vec::new();

    for (n, line) in input.lines().enumerate() {
        let end = n + get_matches(line, &mut winning) + 1;

        if end > card_counts.len() {
            card_counts.resize(end, 1);
        }

        for i in n + 1..end {
            card_counts[i] += card_counts[n];
        }

        count += card_counts[n];
    }

    Some(count)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(13));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(30));
    }
}
