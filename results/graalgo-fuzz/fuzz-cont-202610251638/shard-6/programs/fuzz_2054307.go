package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int; F1 uint32; F2 uint16 }
type GenS4 struct { F0 string; F1 int16; F2 int }
func genFn0() uint8 {
    var v5 int16 = -131
    v5 = -312
    var v6 float64 = ((-2.25 + float64(1.5)) / 1.0)
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int = (^(fi7))
        v9 := []int{4294967295, -9223372036854775808}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v5
    _ = v6
    return 128
}
func genFn1(p0 int8, p1 uint, p2 int64) int32 {
    fmt.Println(p1)
    v10 := []int{127, 9223372036854775807, 1, int(p2)}
    v11 := GenS1{F0: -1, F1: 128, F2: 0}
    switch -957 {
    case 0:
        if ((true && false) || (-1 == 4294967295)) {
            v12 := GenS1{F0: 1, F1: 177, F2: -1}
            var v13 int = -697
            p0 -= (p0 + p0)
            _ = v12
            _ = v13
        }
    case 4294967295:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    case 18:
        switch 128 {
        case -908:
            var v14 int64 = (p2 % (((p2 >> (((p1 | (p1 * p1))) & 63))) | 1))
            _ = v14
        case 300:
            v15 := GenS2{F0: p1, F1: ("\n\t" + ((("" + ("foo bar" + "\n\t")) + (("a" + "foo bar") + (("hello" + "hello") + (("foo bar" + ("hello" + "a")) + (("a" + "") + "foo bar"))))) + ("\n\t" + "a")))}
            _ = v15
        case 2147483647:
            var v16 string = ((("\n\t" + "a") + ("a" + "a")) + "foo bar")
            _ = v16
        }
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return -1
}

func main() {
    var v1 uint = 1
    var v2 int16 = 773
    v3 := []int{237, 944}
    var v4 int8 = 0
    v2 += v2
    var v17 int32 = -2147483648
    if (255 != 255) {
        var v18 int8 = (v4 - v4)
        _ = v18
    } else {
        if len(v3) > 0 {
            v3[0] = -858
        }
        switch 128 {
        case 9223372036854775807:
            if false {
                if true {
                    var v19 int16 = v2
                    _ = v19
                }
                v20 := GenS3{F0: 127, F1: 2147483648, F2: 22}
                _ = v20
            }
        case 42:
            var v21 bool = false
            _ = v21
        default:
            v17 = v17
        }
        if false {
            _ = len(v3)
        }
    }
    v22 := GenS4{F0: "", F1: 536, F2: 32767}
    fmt.Println(v2, v1, v4)
    var v23 uint64 = 658
    fmt.Println(v23, v2, v23)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
