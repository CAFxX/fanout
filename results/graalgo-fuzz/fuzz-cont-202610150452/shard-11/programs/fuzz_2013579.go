package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int16 }

func main() {
    var v1 uint16 = 255
    var v2 string = "foo bar"
    v3 := make(map[string]string)
    var v4 string = ((("\n\t" + v2) + v2) + (v2 + ("" + "\n\t")))
    if true {
        v5 := []int{}
        var v6 int8 = 1
        _ = v5
        _ = v6
    } else {
        v2 = v4
        v1 -= (v1 % ((v1) | 1))
    }
    if (!(v1 < v1)) {
        for fi7 := 0; fi7 < 2; fi7++ {
            v4 = v4
            if (true && true) {
                var v8 int32 = 0
                v9 := GenS1{F0: v4, F1: (v1 ^ (v1 - v1)), F2: 411}
                v1 -= 862
                _ = v8
                _ = v9
            } else {
                if true {
                    v4 = (v4 + "hello")
                    v4 = v4
                    v10 := GenS1{F0: "hello", F1: v1, F2: 0}
                    _ = v10
                }
                v11 := GenS1{F0: ("foo bar" + v4), F1: (v1 << ((uint(v1)) & 63)), F2: 1000}
                fmt.Println(v4, v4, fi7)
                _ = v11
            }
            for fi12 := 0; fi12 < 10; fi12++ {
                v2 = v4
                _ = fi12
            }
            _ = fi7
        }
        var v13 uint32 = 677
        _ = v13
    }
    v2 = v4
    var v14 int16 = 102
    v15 := GenS1{F0: "a", F1: v1, F2: v14}
    v16 := GenS1{F0: ("" + v2), F1: v1, F2: (^(v14))}
    var v17 float32 = (-2.25 * (0.0 * ((1.5 * 1.5) + 0.0)))
    v4 = (v4 + "hello")
    var v18 float64 = 0.1
    switch v2 {
    case "foo bar":
        v19 := []int{}
        _ = v19
    case "a":
        if (!false) {
            if ((-2147483648 >= -70) && true) {
                var v20 int32 = 973
                v21 := []int{}
                var v22 int8 = -1
                _ = v20
                _ = v21
                _ = v22
            }
        }
    }
    var v23 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
