package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint64; F2 int64 }
func genFn0() string {
    var v4 int = 128
    for fi5 := 0; fi5 < 2; fi5++ {
        if false {
            for fi6 := 0; fi6 < 16; fi6++ {
                var v7 int32 = 166
                _ = fi6
                _ = v7
            }
            v8 := GenS1{F0: "hello"}
            _ = v8
        }
        _ = fi5
    }
    _ = v4
    return "\n\t"
}
func genFn1() uint16 {
    if true {
        var v9 uint16 = 0
        fmt.Println(v9)
        _ = v9
    } else {
        v10 := []int{649, -170, 515, 541}
        _ = v10
    }
    v11 := []int{}
    var v12 int64 = 1
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 int16 = 0
    var v2 int16 = v1
    v3 := []int{928}
    fmt.Println(v2, v2)
    if (175 > 0) {
        var v13 string = ("foo bar" + "a")
        _ = v13
    } else {
        for fi14 := 0; fi14 < 11; fi14++ {
            v1 += (v2 | (v2 & (^(v2))))
            _ = fi14
        }
        v15 := GenS2{F0: false}
        var v16 float32 = 1e10
        _ = v15
        _ = v16
    }
    if (!((1e10 * 3.14159) >= -2.25)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if len(v3) > 0 {
            v3[0] = -1
        }
        fmt.Println(v2, v1)
    }
    var v17 float64 = ((((1e10 / (0.0 + 0.1)) + (3.14159 + (0.0 * 3.14159))) - 0.1) / -2.25)
    fmt.Println(v17, v2, v17)
    var v18 int16 = v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch -835 {
    case 4294967295:
        for fi19 := 0; fi19 < 20; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                v1 = v1
                if (((true && true) || (!(0.1 <= (-2.25 + 1.5)))) && (2147483647 < 2147483647)) {
                    fmt.Println(v2, v2)
                    v21 := GenS3{F0: 42, F1: 0, F2: -9223372036854775808}
                    _ = v21
                }
                if true {
                    var v22 float64 = (0.1 + (((v17 / 0.1) / v17) + v17))
                    _ = v22
                }
                _ = fi20
            }
            v1 -= 32767
            var v23 uint32 = 444
            _ = fi19
            _ = v23
        }
    case 2147483647:
        v24 := GenS1{F0: ((("foo bar" + "") + "hello") + "hello")}
        _ = v24
    case -1:
        if (false && true) {
            fmt.Println(v1, v2)
            v17 = (0.0 - ((-2.25 - 1.5) / (1.5 + (3.14159 / (((1e10 / (v17 / 1.5)) / 1.5) - (((v17 - float64((-2.25 * float32(v17)))) - (v17 * 1e10)) / 1.0))))))
            v1 = (v2 << ((18446744073709551615) & 63))
        } else {
            var v25 uint8 = 42
            _ = v25
        }
    default:
        v1 = (v1 & (v2 >> ((1) & 63)))
    }
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
