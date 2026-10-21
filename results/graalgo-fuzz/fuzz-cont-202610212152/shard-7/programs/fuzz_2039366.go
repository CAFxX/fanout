package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int16 }
type GenS2 struct { F0 int32; F1 int32 }

func main() {
    var v1 float64 = 3.14159
    var v2 string = "hello"
    var v3 uint32 = 42
    var v4 uint64 = 2
    var v5 uint64 = v4
    var v6 int16 = 32767
    if true {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 float64 = v1
            _ = fi7
            _ = v8
        }
        if (-1 >= 109) {
            v9 := GenS1{F0: 18446744073709551615, F1: 1, F2: 935}
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := []int{-21, 78, (fi10 % (((-(fi10))) | 1)), (^(fi10))}
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            v1 = v1
            _ = fi12
        }
        switch -2147483648 {
        case 42:
            var v13 int32 = 0
            _ = v13
        case 9223372036854775807:
            var v14 int = -328
            _ = v14
        case -826:
            var v15 int64 = 1
            _ = v15
        default:
            v16 := GenS1{F0: 4294967295, F1: -128, F2: 42}
            _ = v16
        }
    }
    v4 = v4
    v2 = ""
    v6 = v6
    var v17 float64 = (float64(0.0) * (v1 / v1))
    if false {
        if ((!false) || false) {
            for fi18 := 0; fi18 < 9; fi18++ {
                if (76 <= 1) {
                    v19 := []int{9223372036854775807, 2147483647, fi18}
                    v20 := GenS2{F0: 42, F1: -1}
                    _ = v19
                    _ = v20
                }
                fmt.Println(v17, v2)
                var v21 int32 = -117
                _ = fi18
                _ = v21
            }
            var v22 string = v2
            _ = v22
        }
        var v23 string = v2
        _ = v23
    }
    v5 = v4
    v5 += v4
    fmt.Println(v4, v5)
    fmt.Println(v1, v4, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
