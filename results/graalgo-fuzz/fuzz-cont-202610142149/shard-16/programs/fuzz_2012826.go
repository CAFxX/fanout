package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 string }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint }
type GenS4 struct { F0 string; F1 int32; F2 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 string = "\n\t"
    fmt.Println(v3)
    for fi4 := 0; fi4 < 10; fi4++ {
        v3 = ""
        _ = fi4
    }
    if (77 != 0) {
        v3 = ((v3 + "hello") + v3)
    }
    var v5 uint64 = 996
    if (1.5 >= 0.0) {
        v3 = v3
    } else {
        v6 := GenS1{F0: -2147483648, F1: v3, F2: "foo bar"}
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS2{F0: 0, F1: v5}
        fmt.Println(v5)
        switch "foo bar" {
        case "":
            switch 128 {
            case 0:
                v5 = (v5 >> ((42) & 63))
            default:
                v3 = (v3 + ((v3 + "a") + ("a" + "hello")))
            }
        case "foo bar":
            if (float64(1.5) < 0.1) {
                fmt.Println(fi7)
                v9 := GenS1{F0: fi7, F1: ((v3 + "") + ("" + v3)), F2: ""}
                _ = v9
            }
        }
        _ = fi7
        _ = v8
    }
    var v10 uint8 = 103
    v5 -= 18446744073709551615
    v11 := GenS3{F0: v5, F1: v5, F2: 18446744073709551615}
    v12 := GenS4{F0: v3, F1: 2147483647, F2: 120}
    var v13 int8 = -2
    var v14 float32 = 1.5
    var v15 uint = 671
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v5
}
