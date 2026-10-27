package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int8; F2 int8 }

func main() {
    var v1 uint = 0
    v2 := make(map[string]bool)
    var v3 uint8 = 128
    var v4 uint8 = v3
    var v5 float32 = 0.0
    v3 += 235
    v6 := GenS1{F0: -170}
    if false {
        v4 = (v3 | v3)
        if (true || (false || false)) {
            if true {
                for fi7 := 0; fi7 < 9; fi7++ {
                    v8 := GenS1{F0: -1}
                    _ = fi7
                    _ = v8
                }
            } else {
                var v9 int64 = 1
                _ = v9
            }
            v3 += (v4 ^ 255)
        }
    }
    fmt.Println(v3, v4)
    var v10 float64 = 1.5
    if false {
        var v11 uint64 = 0
        _ = v11
    } else {
        var v12 uint = (v1 | 0)
        _ = v12
    }
    if false {
        v13 := []int{1}
        switch 4294967295 {
        case 128:
            if false {
                if (2 < 42) {
                    v14 := GenS2{F0: 42, F1: 100, F2: -128}
                    fmt.Println(v10)
                    if (-322 != 224) {
                        v15 := GenS2{F0: -1, F1: 100, F2: 0}
                        _ = v15
                    }
                    _ = v14
                } else {
                    fmt.Println(v3, v5)
                    _ = len(v13)
                    v1 += (v1 & (^(v1)))
                }
                v16 := GenS2{F0: 9223372036854775807, F1: 1, F2: -128}
                _ = v16
            } else {
                v17 := GenS2{F0: -591, F1: -96, F2: 76}
                _ = v17
            }
        default:
            switch 148 {
            case 4294967295:
                switch 32768 {
                case 2147483647:
                    var v18 int16 = -1
                    _ = v18
                case -1:
                    v19 := GenS2{F0: 32767, F1: -44, F2: 33}
                    _ = v19
                case -847:
                    var v20 uint = v1
                    _ = v20
                }
            case -9223372036854775808:
                v21 := v2[("hello" + ("hello" + "a"))]
                _ = v21
            case 981:
                var v22 int32 = 2147483647
                _ = v22
            }
        }
        _ = v13
    }
    var v23 string = (("foo bar" + "foo bar") + "\n\t")
    v24 := GenS2{F0: -9223372036854775808, F1: 100, F2: 1}
    v25 := GenS2{F0: 238, F1: -95, F2: 127}
    v26 := GenS1{F0: -9223372036854775808}
    v3 = (v3 >> (((v1 % ((1) | 1))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
