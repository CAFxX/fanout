package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint }
func genFn0() bool {
    if true {
        var v6 uint = 18446744073709551615
        v7 := GenS1{F0: 0, F1: 100}
        v8 := []int{1, -2147483648}
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 int16 = -126
        _ = v9
    }
    var v10 uint16 = 42
    var v11 bool = false
    fmt.Println(v10)
    _ = v10
    _ = v11
    return (!((v11 || v11) || v11))
}

func main() {
    var v1 int32 = 707
    var v2 bool = (0 == 2147483648)
    var v3 int8 = -128
    var v4 uint = 1
    v5 := make(map[uint32]int)
    if ((!false) || (1 == 451)) {
        _ = len(v5)
        for fi12 := 0; fi12 < 2; fi12++ {
            fmt.Println(v4)
            _ = fi12
        }
        var v13 string = ""
        _ = v13
    } else {
        switch "foo bar" {
        case "\n\t":
            v1 = v1
        case "":
            v14 := GenS2{F0: ""}
            _ = v14
        }
        v15 := GenS2{F0: ((("\n\t" + "a") + "\n\t") + "")}
        v1 = v1
        _ = v15
    }
    v1 -= (v1 / (((v1 - (v1 ^ (-(v1))))) | 1))
    fmt.Println(v1, v3)
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 18; fi17++ {
            switch "a" {
            case "\n\t":
                v4 = (v4 + (v4 * (^(v4))))
                var v18 uint8 = 2
                var v19 int16 = 4
                _ = v18
                _ = v19
            }
            var v20 int64 = 758
            _ = fi17
            _ = v20
        }
        v2 = (0.1 > float64(1e10))
        v2 = v2
        _ = fi16
    }
    v21 := v5[419]
    v22 := GenS3{F0: "a", F1: v4}
    v3 += -1
    var v23 string = "\n\t"
    fmt.Println(v4, v2)
    var v24 uint64 = 914
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
