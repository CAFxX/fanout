package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int32 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int; F1 uint64; F2 int32 }
type GenS4 struct { F0 int }
func genFn0() int8 {
    var v4 int32 = 106
    for fi5 := 0; fi5 < 17; fi5++ {
        if false {
            v4 += 42
        } else {
            v6 := []int{127}
            v6 = append(v6, (fi5 >> ((433) & 63)))
            switch (("\n\t" + ("hello" + "hello")) + "a") {
            case "a":
                v4 += v4
            case "":
                v7 := GenS1{F0: 1, F1: 127, F2: v4}
                _ = v7
            case "hello":
                var v8 float64 = (3.14159 * 1e10)
                _ = v8
            default:
                v4 -= (^(v4))
            }
            _ = v6
        }
        _ = fi5
    }
    v9 := GenS1{F0: -169, F1: 127, F2: (-(v4))}
    var v10 uint32 = 2
    _ = v10
    _ = v4
    _ = v9
    return 65
}
func genFn1() uint8 {
    var v11 int32 = -19
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 float64 = 0.1
        var v14 int32 = 2147483647
        var v15 uint32 = 42
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: "a", F1: 2}
    fmt.Println(v11)
    v11 -= 2147483647
    _ = v11
    _ = v16
    return 61
}

func main() {
    var v1 int16 = 0
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    _ = len(v3)
    var v17 int8 = 100
    var v18 int32 = 0
    v1 = 855
    if ((v18 % ((42) | 1)) > v18) {
        v19 := GenS2{F0: "\n\t", F1: 963}
        var v20 string = "a"
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 16; fi21++ {
            fmt.Println(fi21)
            var v22 uint32 = 42
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 7; fi23++ {
            v18 -= 1
            _ = fi23
        }
        var v24 int16 = v1
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := GenS3{F0: (fi25 | (fi25 + (fi25 ^ (fi25 / (((-(fi25))) | 1))))), F1: 0, F2: int32(v1)}
        _ = fi25
        _ = v26
    }
    fmt.Println(v17, v1)
    var v27 bool = false
    var v28 bool = ((^(v18)) != 1)
    if v28 {
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS4{F0: 32767}
            _ = fi29
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
