package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int64; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint16; F1 int64; F2 int64 }
func genFn0(p0 uint, p1 string, p2 int16) bool {
    v4 := GenS1{F0: 0, F1: 2147483647}
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS2{F0: 219, F1: (p1 != p1)}
        _ = fi5
        _ = v6
    }
    var v7 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return (-699 != 0)
}
func genFn1() uint {
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int8 = -40
        v9 += v9
        v9 -= 17
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 255
    v11 := GenS3{F0: 1, F1: 255, F2: 1}
    _ = v10
    _ = v11
    return 608
}

func main() {
    v1 := []int{}
    var v2 uint16 = 2
    var v3 bool = false
    v12 := GenS2{F0: 18446744073709551615, F1: false}
    v1 = append(v1, 0)
    switch ("foo bar" + "hello") {
    case "foo bar":
        v2 += v2
    default:
        v13 := GenS2{F0: 255, F1: v3}
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        if v3 {
            fmt.Println(v2, v2, v3)
            v3 = v3
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 float64 = (float64(1e10) + 3.14159)
                _ = fi15
                _ = v16
            }
        } else {
            if len(v1) > 0 {
                v1[0] = fi14
            }
            var v17 float32 = ((((((0.0 * 1e10) * ((0.1 - 1e10) / 1.5)) - 1.5) - float32(0.0)) - float32((float64(1e10) / -2.25))) / 0.1)
            if (true || (1 >= 127)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v3 = (v3 || ((v3 && (1 < 0)) || false))
                _ = len(v1)
            } else {
                var v18 int16 = 1000
                _ = v18
            }
            _ = v17
        }
        v2 = (v2 - v2)
        _ = fi14
    }
    switch 32768 {
    case 2147483647:
        v2 += v2
    }
    v19 := GenS2{F0: genFn1(), F1: ((-2.25 > float64(0.0)) || (v3 && false))}
    var v20 int64 = 0
    if v3 {
        var v21 float64 = -2.25
        var v22 string = "a"
        v2 = 0
        _ = v21
        _ = v22
    }
    if true {
        var v23 int8 = 66
        if (-2147483648 > 42) {
            var v24 string = ("a" + ("a" + "hello"))
            _ = v24
        }
        _ = v23
    }
    v25 := GenS3{F0: (v20 - v20), F1: v2, F2: (v20 / ((v20) | 1))}
    fmt.Println(v2)
    var v26 int64 = v20
    if true {
        switch ("" + "\n\t") {
        case "\n\t":
            var v27 bool = false
            _ = v27
        case "a":
            v28 := GenS4{F0: (v2 % (((^(v2))) | 1)), F1: (v20 + v26), F2: 564}
            _ = v28
        case "hello":
            v26 += v20
        default:
            for fi29 := 0; fi29 < 9; fi29++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi29
            }
        }
        if (646 >= 65535) {
            for fi30 := 0; fi30 < 10; fi30++ {
                var v31 float64 = float64(float32(float64(0.0)))
                v20 -= (v20 - -9223372036854775808)
                fmt.Println(v26, v2)
                _ = fi30
                _ = v31
            }
            var v32 uint = 826
            var v33 int32 = -1
            _ = v32
            _ = v33
        } else {
            v26 = (v20 * (v26 % ((v20) | 1)))
            for fi34 := 0; fi34 < 18; fi34++ {
                if false {
                    var v35 int32 = -620
                    _ = v35
                } else {
                    var v36 float32 = -2.25
                    v3 = v3
                    var v37 bool = false
                    _ = v36
                    _ = v37
                }
                _ = fi34
            }
        }
    }
    v26 -= v26
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
}
