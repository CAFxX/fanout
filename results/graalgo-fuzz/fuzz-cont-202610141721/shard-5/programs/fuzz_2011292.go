package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint32 }
type GenS2 struct { F0 string; F1 int64; F2 uint }
type GenS3 struct { F0 bool; F1 int16; F2 int8 }
func genFn0() uint16 {
    v6 := GenS1{F0: ("foo bar" + "hello"), F1: 1, F2: 284}
    var v7 float32 = 0.1
    v8 := GenS1{F0: "\n\t", F1: 25, F2: 42}
    v7 += v7
    _ = v6
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 int16 = 32767
    var v2 uint32 = 554
    var v3 uint32 = (^(v2))
    v4 := make(map[int]bool)
    var v5 uint8 = 0
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := []int{(-(fi9)), fi9, (^(fi9)), fi9}
        _ = fi9
        _ = v10
    }
    var v11 int16 = 42
    if true {
        var v12 uint8 = (v5 ^ (v5 >> ((18446744073709551615) & 63)))
        var v13 int32 = 843
        _ = v12
        _ = v13
    }
    v1 -= (v1 * (^(v1)))
    v3 = 2147483648
    if (v3 < v3) {
        v2 = v3
        v14 := GenS1{F0: "hello", F1: (^(v5)), F2: 127}
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 int32 = 0
            v17 := []int{}
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    } else {
        v5 += v5
        v11 -= 1000
        var v18 uint64 = 362
        _ = v18
    }
    v2 -= 2
    var v19 string = "a"
    fmt.Println(v1, v1)
    v5 += (v5 & v5)
    for fi20 := 0; fi20 < 4; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 uint64 = 0
            v23 := GenS2{F0: (v19 + v19), F1: 1, F2: 18446744073709551615}
            _ = fi21
            _ = v22
            _ = v23
        }
        var v24 uint64 = 1
        _ = fi20
        _ = v24
    }
    switch ("a" + v19) {
    case "a":
        var v25 uint64 = 418
        _ = v25
    default:
        switch 42 {
        case -787:
            if false {
                var v26 uint = 65535
                fmt.Println(v11, v1, v11)
                _ = v26
            }
        default:
            v27 := GenS3{F0: (true && ((v11 / ((v11) | 1)) < -673)), F1: -154, F2: 42}
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
