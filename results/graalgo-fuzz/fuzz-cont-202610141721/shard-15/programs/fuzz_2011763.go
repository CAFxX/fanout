package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 float64 = -2.25
    var v4 uint32 = 42
    var v5 string = "foo bar"
    v5 = ("hello" + (v5 + v5))
    v6 := GenS1{F0: 42}
    v7 := GenS1{F0: 0}
    delete(v1, 127)
    if (!(true && (false && true))) {
        fmt.Println(v3)
        v5 = v5
    }
    if (true && (((1e10 * float32(v3)) * 1.5) < 3.14159)) {
        v3 -= 0.0
        if true {
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 float32 = (float32(v3) * 0.1)
                _ = fi8
                _ = v9
            }
            v10 := []int{-2147483648, 2147483647, 128}
            v5 = v5
            _ = v10
        }
        if true {
            v3 = (v3 * float64(0.1))
            switch 0 {
            case 127:
                for fi11 := 0; fi11 < 2; fi11++ {
                    v12 := []int{(fi11 / ((fi11) | 1)), (-(fi11)), fi11}
                    _ = fi11
                    _ = v12
                }
            default:
                fmt.Println(v4)
            }
        }
    }
    if true {
        v3 += 1e10
    } else {
        fmt.Println(v4)
        v3 = 0.0
    }
    v3 -= (-2.25 - v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
