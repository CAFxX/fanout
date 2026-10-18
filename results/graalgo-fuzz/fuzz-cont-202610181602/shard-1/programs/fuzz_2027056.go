package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }

func main() {
    var v1 uint64 = 18
    v2 := make(map[string]int)
    var v3 int16 = 0
    v4 := []int{865, -1}
    v5 := make(map[string]int)
    v3 += 1000
    var v6 uint64 = (^(v1))
    for fi7 := 0; fi7 < 17; fi7++ {
        v3 = (v3 / (((v3 << ((18446744073709551615) & 63))) | 1))
        _ = fi7
    }
    if false {
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 6; fi9++ {
                v6 = v1
                if false {
                    var v10 uint8 = 128
                    var v11 uint64 = (v1 ^ 130)
                    v1 -= v6
                    _ = v10
                    _ = v11
                }
                _ = len(v4)
                _ = fi9
            }
            switch (fi8 % ((-948) | 1)) {
            case 2147483647:
                for fi12 := 0; fi12 < 3; fi12++ {
                    v13 := GenS1{F0: "a", F1: "foo bar"}
                    _ = fi12
                    _ = v13
                }
            default:
                if ((v3 >> ((0) & 63)) >= v3) {
                    v1 -= v6
                } else {
                    var v14 uint16 = 32767
                    v6 = (v1 - v6)
                    _ = v14
                }
            }
            _ = fi8
        }
        var v15 int32 = 42
        fmt.Println(v3, v1)
        _ = v15
    } else {
        v16 := GenS1{F0: "hello", F1: "hello"}
        _ = v16
    }
    fmt.Println(v3)
    for fi17 := 0; fi17 < 4; fi17++ {
        v1 += v6
        var v18 int32 = 42
        v6 = (v6 & 977)
        _ = fi17
        _ = v18
    }
    v4 = append(v4, 128)
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := GenS1{F0: ("hello" + (("hello" + (("" + "hello") + "hello")) + ("a" + "a"))), F1: "a"}
        _ = len(v5)
        v6 = v6
        _ = fi19
        _ = v20
    }
    var v21 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
