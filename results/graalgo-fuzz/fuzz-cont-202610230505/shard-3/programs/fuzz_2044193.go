package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint64; F1 string; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int16; F2 bool }
type GenS5 struct { F0 uint64; F1 int64; F2 uint }
func genFn0() uint32 {
    var v4 int16 = 859
    v4 += v4
    v5 := GenS1{F0: 1, F1: 1}
    _ = v4
    _ = v5
    return 151
}
func genFn1(p0 uint64) uint32 {
    var v6 int64 = 1
    v7 := []int{4294967295, -604, 0, 1}
    _ = p0
    _ = v6
    _ = v7
    return 893
}

func main() {
    var v1 uint8 = 50
    var v2 float32 = 1e10
    var v3 string = "foo bar"
    var v8 int64 = 693
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            if true {
                v3 = v3
            } else {
                v3 = v3
            }
            _ = fi10
        }
        if ((1 != -384) || (477 < (-(v8)))) {
            v11 := GenS1{F0: 255, F1: 4294967295}
            _ = v11
        } else {
            if ((((-2.25 * ((3.14159 + 0.1) + 0.0)) <= (float64(((v2 + -2.25) - v2)) + 3.14159)) && (!true)) || (32767 >= 0)) {
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 float32 = -2.25
                    _ = fi12
                    _ = v13
                }
            }
            v14 := []int{(^(fi9)), fi9, fi9, fi9}
            _ = v14
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            v2 = float32((0.1 / ((1e10 / float64((float32(3.14159) + (v2 - (3.14159 + ((float32(-2.25) - ((1.5 * 3.14159) + float32(0.0))) / v2)))))) + float64(v2))))
            v16 := GenS1{F0: 18446744073709551615, F1: 0}
            _ = fi15
            _ = v16
        }
        _ = fi9
    }
    v17 := GenS2{F0: 58, F1: v3, F2: 42}
    var v18 uint64 = 18446744073709551615
    switch "\n\t" {
    case "":
        switch -11 {
        case -9223372036854775808:
            var v19 string = (v3 + v3)
            _ = v19
        default:
            if true {
                if true {
                    v20 := GenS3{F0: (-2.25 != 0.1)}
                    _ = v20
                }
                v18 += (^(v18))
                v2 = v2
            }
        }
    case "foo bar":
        v3 = v3
    default:
        var v21 int16 = 32767
        _ = v21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 float64 = (1.5 - 1.5)
        _ = fi22
        _ = v23
    }
    fmt.Println(v2)
    v2 -= v2
    fmt.Println(v8, v18)
    var v24 int8 = 16
    switch 9223372036854775807 {
    case -263:
        var v25 uint64 = v18
        _ = v25
    case 128:
        if ((v3 + ("a" + v3)) != v3) {
            switch -9223372036854775808 {
            case 32767:
                v26 := GenS3{F0: ((-805 == 0) && (!false))}
                _ = v26
            case 308:
                fmt.Println(v18)
            case 9223372036854775807:
                for fi27 := 0; fi27 < 12; fi27++ {
                    for fi28 := 0; fi28 < 4; fi28++ {
                        var v29 uint16 = 2
                        v30 := []int{fi28, fi27, 9223372036854775807, (fi28 + (fi28 - fi27))}
                        v31 := GenS4{F0: false, F1: 42, F2: true}
                        _ = fi28
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    _ = fi27
                }
            default:
                v32 := []int{1, 398}
                _ = v32
            }
        }
    case 127:
        switch "hello" {
        case "":
            for fi33 := 0; fi33 < 12; fi33++ {
                v34 := []int{(fi33 + (fi33 | (fi33 << ((65) & 63)))), (fi33 + (fi33 / ((fi33) | 1))), fi33, (fi33 % ((-9223372036854775808) | 1))}
                if (2 != 0) {
                    var v35 int8 = v24
                    v24 -= v24
                    v34 = append(v34, 128)
                    _ = v35
                }
                v24 += 42
                _ = fi33
                _ = v34
            }
        default:
            var v36 int32 = 2147483647
            _ = v36
        }
    default:
        var v37 uint8 = v1
        _ = v37
    }
    if (0 > 255) {
        for fi38 := 0; fi38 < 4; fi38++ {
            for fi39 := 0; fi39 < 16; fi39++ {
                v40 := GenS5{F0: (^(v18)), F1: 9223372036854775807, F2: 2}
                var v41 string = ((((("\n\t" + v3) + ((v3 + v3) + "\n\t")) + ("" + "hello")) + v3) + (v3 + v3))
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = fi38
        }
        fmt.Println(v3)
    }
    v42 := GenS5{F0: 42, F1: v8, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v42
    _ = v8
}
