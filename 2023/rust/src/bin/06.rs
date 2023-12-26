
advent_of_code::solution!(6);

fn get_races(input: &str) -> Vec<Vec<i32>> {
    input
        .lines()
        .map(|line| {
            let collection = line.split(":").collect::<Vec<&str>>();
            let numbers: Vec<i32> = collection[1]
                .split_whitespace()
                .filter_map(|s| s.parse().ok())
                .collect();
            numbers
        })
        .collect()
}

fn quadratic_solver(a: f64, b: f64, c: f64) -> (i32, i32) {
    let discriminant = b * b - 4.0 * a * c;

    if discriminant < 0.0 {
        return (0, 0);
    }

    let sqrt_discriminant = discriminant.sqrt();
    let root1 = (-b + sqrt_discriminant) / (2.0 * a);
    let root2 = (-b - sqrt_discriminant) / (2.0 * a);

    (root1.ceil() as i32, root2.floor() as i32)
}

fn number_of_ways(races: &Vec<Vec<i32>>) -> Option<u32>{
    let mut result: u32 = 1;
    for i in 0..races[0].len() {
        let t = races[0][i];
        let d = races[1][i];

        let number_of_ways = quadratic_solver(1., (-t).try_into().unwrap(), d.try_into().unwrap());
        let total_number_of_ways = number_of_ways.0 - number_of_ways.1 - 1;
        result *= (total_number_of_ways) as u32;
    }

    Some(result)
}

pub fn part_one(input: &str) -> Option<u32> {
    let races = get_races(input);
    number_of_ways(&races)
}

pub fn part_two(input: &str) -> Option<u32> {
    // I did this by hand and through it into wolframalpha
    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(288));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, None);
    }
}
