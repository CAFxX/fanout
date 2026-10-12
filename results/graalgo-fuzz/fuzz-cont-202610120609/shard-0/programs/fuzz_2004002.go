package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    var v3 float64 = 3.14159
    var v4 int16 = -56
    v5 := make(map[int]int)
    var v6 int16 = 1
    v6 += 42
    v7 := GenS1{F0: 29}
    for fi8 := 0; fi8 < 18; fi8++ {
        if (true && false) {
            v6 += (v6 | 42)
        } else {
            fmt.Println(fi8, v6)
            v6 = v6
        }
        _ = fi8
    }
    switch "\n\t" {
    case "foo bar":
        if true {
            v3 -= 1.5
        } else {
            v9 := GenS2{F0: 395, F1: 130, F2: 9223372036854775807}
            v2["\n\t"] = 2
            if ((-(v4)) > v6) {
                var v10 float64 = 3.14159
                v10 -= (v10 * v3)
                var v11 int = 819
                _ = v10
                _ = v11
            }
            _ = v9
        }
    case "":
        fmt.Println(v3, v3)
    }
    if true {
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 int16 = (v4 % ((v6) | 1))
            _ = fi12
            _ = v13
        }
        v4 = (v6 + (^(v6)))
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS3{F0: ("a" + "\n\t")}
            _ = fi14
            _ = v15
        }
    }
    v16 := v2[("hello" + ("hello" + ("hello" + "hello")))]
    fmt.Println(v6, v16, v16)
    var v17 bool = (true && (1 > 18446744073709551615))
    fmt.Println(v17, v3)
    if false {
        v4 -= (v6 | -1)
    }
    v18 := []int{356, 862, 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
