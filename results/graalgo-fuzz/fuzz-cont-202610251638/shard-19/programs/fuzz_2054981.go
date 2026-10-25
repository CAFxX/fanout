package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }

func main() {
    v1 := []string{}
    var v2 uint = 409
    v3 := []int{42, 217, 2147483647, 2}
    var v4 int64 = -9223372036854775808
    for fi5 := 0; fi5 < 11; fi5++ {
        if ((false || (-51 < 1)) && (229 >= 2147483647)) {
            v4 += int64(v2)
        } else {
            fmt.Println(fi5)
        }
        _ = fi5
    }
    v4 = v4
    v4 = -55
    var v6 float32 = 0.0
    v7 := GenS1{F0: "a"}
    v2 = v2
    v1 = append(v1, "\n\t")
    fmt.Println(v4, v2)
    var v8 uint8 = 77
    if (32767 < 1) {
        v8 -= 2
        v9 := GenS2{F0: 32767}
        _ = v9
    }
    _ = len(v1)
    for fi10 := 0; fi10 < 17; fi10++ {
        v4 = (v4 * v4)
        fmt.Println(v4)
        fmt.Println(v6)
        _ = fi10
    }
    v6 += v6
    switch 9223372036854775807 {
    case 4294967295:
        if true {
            var v11 float64 = float64(1.5)
            _ = v11
        }
    }
    fmt.Println(v6, v6, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
