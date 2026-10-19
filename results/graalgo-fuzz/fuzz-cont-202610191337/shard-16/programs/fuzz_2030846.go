package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
func genFn0(p0 int8, p1 uint8, p2 uint32) uint32 {
    v4 := []int{-949, 9223372036854775807, 32768}
    var v5 string = "a"
    v5 = (("hello" + "\n\t") + v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}
func genFn1() int16 {
    if false {
        if ((-2.25 + float32((3.14159 + 3.14159))) >= 1.5) {
            var v6 uint8 = 255
            v6 -= (v6 - v6)
            fmt.Println(v6)
            _ = v6
        }
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{637, fi7, (fi7 - 128)}
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 uint32 = 1
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: 42, F1: 100}
            _ = fi7
            _ = v11
            _ = v8
        }
    } else {
        var v12 float64 = float64(0.1)
        _ = v12
    }
    v13 := []int{-9223372036854775808, 551}
    _ = v13
    return 49
}

func main() {
    v1 := []uint32{2147483647}
    v2 := []string{"foo bar", ("foo bar" + "\n\t"), "\n\t"}
    var v3 int8 = -128
    v3 = v3
    v3 = v3
    fmt.Println(v3)
    if len(v2) > 0 {
        v2[0] = ""
    }
    var v14 float32 = ((3.14159 + -2.25) * -2.25)
    for fi15 := 0; fi15 < 9; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            v3 = v3
            var v17 float64 = 1e10
            _ = fi16
            _ = v17
        }
        v14 -= (v14 * (-2.25 / v14))
        _ = fi15
    }
    switch 32768 {
    case -2147483648:
        v18 := GenS1{F0: 42, F1: (v3 | (v3 | 63))}
        _ = v18
    case -44:
        for fi19 := 0; fi19 < 7; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                if ((42 >= 2) || false) {
                    var v21 uint = 0
                    if (v21 != (v21 - v21)) {
                        var v22 int32 = 39
                        _ = v22
                    }
                    var v23 int64 = 1
                    _ = v21
                    _ = v23
                }
                v1 = append(v1, 2147483648)
                _ = fi20
            }
            _ = fi19
        }
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 int8 = (v3 % (((v3 ^ 1)) | 1))
        v25 = v3
        var v26 float32 = (v14 - float32(1e10))
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        v2 = append(v2, ("hello" + ("a" + "")))
        v14 = v14
        _ = fi27
    }
    v14 -= v14
    var v28 uint16 = 577
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v28
    _ = v3
}
