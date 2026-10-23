package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint32) int8 {
    v4 := []int{-1, -1}
    if ((p0 | p0) >= (p0 % ((p0) | 1))) {
        p0 += 695
    } else {
        if len(v4) > 0 {
            v4[0] = -580
        }
        if ((9223372036854775807 == 644) && (100 != 79)) {
            for fi5 := 0; fi5 < 2; fi5++ {
                var v6 uint8 = 255
                _ = fi5
                _ = v6
            }
            var v7 uint64 = 4294967295
            var v8 string = (("foo bar" + "hello") + ("\n\t" + "\n\t"))
            _ = v7
            _ = v8
        }
    }
    v9 := GenS1{F0: 0, F1: -2147483648, F2: 2}
    p0 = (p0 & 2147483648)
    p0 += 65535
    _ = p0
    _ = v4
    _ = v9
    return 42
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    var v10 float32 = (float32((((1e10 + 1.5) * (1.5 / float64(1.5))) * -2.25)) / (1.5 - -2.25))
    var v11 int8 = genFn0(1)
    if (32768 < 32768) {
        v12 := GenS2{F0: 690}
        _ = v12
    }
    var v13 float64 = (3.14159 * (0.1 - (((-2.25 + -2.25) / (float64(v10) - ((1e10 * 1e10) + 0.1))) + (float64(v10) - 3.14159))))
    v11 -= genFn0(1)
    if true {
        switch 4294967295 {
        case -9223372036854775808:
            switch 128 {
            case 128:
                v14 := GenS3{F0: 1000, F1: 2147483647}
                _ = v14
            case 1:
                delete(v3, 4294967295)
            case 32768:
                v13 += v13
            default:
                v15 := GenS4{F0: int8(v1)}
                _ = v15
            }
        case 9223372036854775807:
            v16 := GenS5{F0: 42}
            _ = v16
        default:
            v10 = ((float32((float64(v10) - (0.1 / v13))) * float32((float64(v10) + 0.0))) - v10)
        }
    } else {
        fmt.Println(v11, v13, v13)
        v10 += v10
        v17 := GenS3{F0: -32768, F1: 2147483647}
        _ = v17
    }
    var v18 string = ""
    if false {
        if ((^(v1)) >= (^(v1))) {
            var v19 string = ("a" + "\n\t")
            v20 := v2[v18]
            _ = v19
            _ = v20
        } else {
            var v21 string = (v18 + v18)
            _ = v21
        }
        v1 = 128
        v1 -= 1
    } else {
        if false {
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := GenS6{F0: 32767}
                v1 -= (^(v1))
                _ = fi22
                _ = v23
            }
            var v24 uint16 = 32767
            _ = v24
        } else {
            v25 := GenS6{F0: 125}
            var v26 string = v18
            _ = v25
            _ = v26
        }
        var v27 int = -1
        fmt.Println(v18, v18)
        _ = v27
    }
    v28 := GenS6{F0: 644}
    v1 -= v1
    v18 = ((("a" + "\n\t") + ((v18 + v18) + v18)) + "foo bar")
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v18
    _ = v2
    _ = v28
    _ = v3
}
