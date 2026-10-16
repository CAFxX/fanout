package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 int32 }
func genFn0() uint8 {
    var v5 int8 = -128
    for fi6 := 0; fi6 < 6; fi6++ {
        v5 = 44
        v7 := GenS1{F0: 18446744073709551615, F1: 4294967295}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 765, F1: 4294967295}
    _ = v5
    _ = v8
    return 185
}
func genFn1(p0 int64, p1 uint32) int16 {
    var v9 int8 = 119
    var v10 uint32 = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 32767
    var v3 int8 = 0
    var v4 string = "a"
    fmt.Println(v4, v2, v4)
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS2{F0: 793, F1: -543}
        v4 = (v4 + (v4 + v4))
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        if ((^(v2)) < (^(v2))) {
            var v14 uint32 = 663
            fmt.Println(fi13)
            for fi15 := 0; fi15 < 4; fi15++ {
                v16 := GenS1{F0: 2, F1: 419}
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        _ = fi13
    }
    v4 = "a"
    var v17 int = -1
    if (2 == genFn0()) {
        var v18 float64 = ((1e10 / float64((-2.25 - 1e10))) / float64((0.1 / 1e10)))
        if true {
            for fi19 := 0; fi19 < 20; fi19++ {
                v18 = 1e10
                v3 -= v3
                v20 := GenS2{F0: 42, F1: -9223372036854775808}
                _ = fi19
                _ = v20
            }
            v1[v4] = "foo bar"
            var v21 uint = 765
            _ = v21
        } else {
            switch v4 {
            case "hello":
                fmt.Println(v3, v17, v17)
            }
            if false {
                v18 = (v18 * float64(1e10))
                for fi22 := 0; fi22 < 2; fi22++ {
                    v23 := GenS2{F0: 2147483648, F1: -9223372036854775808}
                    _ = fi22
                    _ = v23
                }
                var v24 float64 = -2.25
                _ = v24
            } else {
                for fi25 := 0; fi25 < 8; fi25++ {
                    v2 += 0
                    if (191 != 1) {
                        v26 := v1[""]
                        _ = v26
                    } else {
                        var v27 bool = true
                        v28 := GenS2{F0: 4294967295, F1: -9223372036854775808}
                        _ = v27
                        _ = v28
                    }
                    _ = fi25
                }
            }
        }
        v29 := GenS3{F0: (-(v17)), F1: v4}
        _ = v18
        _ = v29
    }
    v30 := GenS3{F0: (^(v17)), F1: v4}
    switch (("foo bar" + v4) + "") {
    case "":
        if false {
            var v31 float64 = (float64(float32((0.0 * 0.0))) / (((-2.25 + (1e10 - (1e10 - float64(0.0)))) / (1e10 / 0.1)) + ((float64((((1e10 - float32((1e10 / 3.14159))) + float32(0.1)) - -2.25)) * 3.14159) - (0.0 / 0.1))))
            var v32 float64 = v31
            _ = v31
            _ = v32
        }
    default:
        v33 := GenS4{F0: 1000, F1: false}
        _ = v33
    }
    delete(v1, v4)
    for fi34 := 0; fi34 < 14; fi34++ {
        v35 := GenS5{F0: 2147483647}
        var v36 float32 = 0.0
        _ = fi34
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        v38 := GenS1{F0: 1, F1: 619}
        _ = len(v1)
        _ = fi37
        _ = v38
    }
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
