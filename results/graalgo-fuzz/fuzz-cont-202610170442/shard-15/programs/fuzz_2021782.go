package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
func genFn0(p0 uint32, p1 int32, p2 int64) int64 {
    var v4 float64 = (1e10 - -2.25)
    var v5 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}
func genFn1(p0 uint16, p1 uint16, p2 uint64) string {
    var v6 string = ""
    p2 -= (p2 >> ((0) & 63))
    switch 4294967295 {
    case 0:
        p1 += p1
    default:
        p1 += 32767
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint32 = 2147483647
        for fi9 := 0; fi9 < 10; fi9++ {
            p0 += p1
            if true {
                var v10 int32 = 1
                _ = v10
            } else {
                v11 := GenS1{F0: 469}
                var v12 int8 = 60
                v8 += (^(v8))
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        if (!false) {
            var v13 uint8 = 42
            var v14 string = (v6 + v6)
            p0 = p1
            _ = v13
            _ = v14
        }
        _ = fi7
        _ = v8
    }
    v15 := []int{-2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v6
    return ((v6 + v6) + ((v6 + v6) + "a"))
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{65535, 445, 42}
    var v3 uint64 = 2
    var v16 uint16 = 1
    v1 += (v3 ^ 723)
    v17 := GenS1{F0: 65535}
    fmt.Println(v1)
    var v18 uint8 = 244
    if false {
        var v19 int16 = 32767
        v20 := GenS1{F0: 65535}
        v18 = (v18 >> ((1) & 63))
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 10; fi21++ {
            fmt.Println(v18, v16)
            if false {
                v22 := GenS1{F0: 2}
                v23 := GenS1{F0: 255}
                if false {
                    v24 := GenS1{F0: uint(v3)}
                    _ = v24
                } else {
                    v25 := GenS1{F0: uint(v1)}
                    _ = v25
                }
                _ = v22
                _ = v23
            } else {
                fmt.Println(v18, fi21, v1)
            }
            _ = fi21
        }
        fmt.Println(v16, v18, v1)
        switch 2147483647 {
        case 0:
            v1 -= (v1 ^ (v3 * v1))
        case 9223372036854775807:
            var v26 uint32 = 463
            _ = v26
        default:
            switch "" {
            case "hello":
                v18 -= uint8(v3)
            case "":
                _ = len(v2)
            default:
                v27 := GenS2{F0: false}
                _ = v27
            }
        }
    }
    var v28 bool = ((!false) && (1e10 == 1.5))
    for fi29 := 0; fi29 < 10; fi29++ {
        v3 += (v3 << ((18446744073709551615) & 63))
        _ = fi29
    }
    v18 = v18
    var v30 bool = v28
    v31 := GenS2{F0: (0 <= 504)}
    for fi32 := 0; fi32 < 19; fi32++ {
        v28 = ((-2.25 + (((((3.14159 + (0.1 + (0.1 * (float64(1e10) / (((1.5 / 1.0) - 0.1) + -2.25))))) + 1.5) * (1e10 - (0.1 - ((((0.0 + ((float64(float32((-2.25 / 0.1))) * (float64(0.0) - ((float64((0.0 * ((-2.25 + (-2.25 / (0.1 * 1.5))) + float32(3.14159)))) + (1e10 / 1.5)) / 1e10))) / ((3.14159 * -2.25) - float64(1e10)))) + (-2.25 - (0.0 * -2.25))) * 1e10) - (1.5 / 1e10))))) * 3.14159) + 1.5)) != -2.25)
        v1 = v3
        for fi33 := 0; fi33 < 19; fi33++ {
            v1 -= v1
            if (255 >= 18446744073709551615) {
                if len(v2) > 0 {
                    v2[0] = uint32(v1)
                }
            }
            _ = fi33
        }
        _ = fi32
    }
    var v34 string = "hello"
    if false {
        v18 = (v18 ^ v18)
        fmt.Println(v34, v30)
        var v35 uint16 = v16
        _ = v35
    }
    v28 = (v28 && true)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v34
}
