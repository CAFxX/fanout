package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint64; F2 int16 }

func main() {
    var v1 bool = false
    var v2 int32 = 1
    var v3 int32 = (v2 << ((775) & 63))
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := []int{(fi4 >> ((0) & 63)), (fi4 & (^(fi4))), (fi4 | (fi4 % ((fi4) | 1))), (fi4 / (((^(fi4))) | 1))}
        if len(v5) > 0 {
            v5[0] = (fi4 % (((fi4 << ((480) & 63))) | 1))
        }
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 float64 = 0.0
            switch "hello" {
            case "\n\t":
                v8 := GenS1{F0: -168, F1: 431}
                _ = v8
            default:
                v9 := GenS2{F0: 1000}
                _ = v9
            }
            fmt.Println(v3, v1)
            v10 := GenS3{F0: "", F1: 87, F2: 0}
            _ = fi6
            _ = v10
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    v1 = (!v1)
    var v11 bool = (0.1 == ((-2.25 / 1.0) * (((-2.25 / 1.0) + (0.0 / (3.14159 - (-2.25 * 1.5)))) - (float32(-2.25) + 1.5))))
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint16 = 0
        _ = fi12
        _ = v13
    }
    v2 += (v2 / ((42) | 1))
    if (42 <= 262) {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int32 = (v3 - v2)
            v16 := GenS3{F0: "a", F1: 18446744073709551615, F2: 1}
            _ = fi14
            _ = v15
            _ = v16
        }
        v1 = v1
        v17 := []int{}
        _ = v17
    } else {
        var v18 int8 = 4
        _ = v18
    }
    v2 = (-(v2))
    for fi19 := 0; fi19 < 6; fi19++ {
        if (v3 > (v3 << ((18446744073709551615) & 63))) {
            for fi20 := 0; fi20 < 3; fi20++ {
                if (!v1) {
                    var v21 uint32 = 4294967295
                    var v22 float64 = float64(((float32((0.1 * 1e10)) * 0.0) / 0.1))
                    var v23 uint16 = uint16(v21)
                    _ = v21
                    _ = v22
                    _ = v23
                }
                var v24 bool = ("hello" <= ("hello" + (("a" + (("foo bar" + ("a" + "foo bar")) + "")) + "foo bar")))
                v25 := []int{2147483647, fi19, (fi20 ^ (fi19 / ((128) | 1))), 32768}
                _ = fi20
                _ = v24
                _ = v25
            }
            v26 := []int{}
            _ = v26
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
