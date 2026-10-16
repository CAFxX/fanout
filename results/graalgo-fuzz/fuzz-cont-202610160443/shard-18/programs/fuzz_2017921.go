package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 string; F2 uint8 }
func genFn0() uint16 {
    var v5 uint64 = 2
    var v6 int32 = 42
    v5 += (v5 >> ((18446744073709551615) & 63))
    _ = v5
    _ = v6
    return 229
}
func genFn1() uint8 {
    if (42 >= 1) {
        v7 := GenS1{F0: 100}
        _ = v7
    } else {
        switch 128 {
        case 127:
            v8 := []int{}
            _ = v8
        case -901:
            v9 := GenS1{F0: -1}
            _ = v9
        case -701:
            var v10 uint16 = 2
            _ = v10
        default:
            var v11 int8 = 1
            _ = v11
        }
        var v12 int16 = 1
        v12 += (v12 / (((v12 ^ v12)) | 1))
        _ = v12
    }
    v13 := GenS1{F0: 42}
    var v14 int64 = 1
    _ = v13
    _ = v14
    return 0
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = (v1 >> ((0) & 63))
    var v3 string = ""
    v4 := []string{(v3 + ((((v3 + v3) + (v3 + "foo bar")) + "a") + (v3 + v3))), (("\n\t" + ("hello" + v3)) + (v3 + v3)), v3, (v3 + ("foo bar" + "a"))}
    _ = len(v4)
    for fi15 := 0; fi15 < 18; fi15++ {
        if false {
            v3 = v3
            v16 := GenS1{F0: 100}
            if (!(("" + v3) != v3)) {
                var v17 float32 = 1e10
                _ = v17
            } else {
                switch (^(fi15)) {
                case 32767:
                    var v18 float32 = (((1e10 * (3.14159 * 0.1)) - float32((1e10 + 0.1))) + 0.0)
                    _ = v18
                default:
                    var v19 int64 = (-(v2))
                    _ = v19
                }
            }
            _ = v16
        } else {
            if (963 > 18446744073709551615) {
                var v20 int8 = 0
                var v21 uint = 789
                _ = len(v4)
                _ = v20
                _ = v21
            } else {
                _ = len(v4)
                v1 += v2
                v2 -= (v2 % ((v2) | 1))
            }
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            fmt.Println(fi22)
            _ = fi22
        }
        fmt.Println(fi15, fi15)
        _ = fi15
    }
    fmt.Println(v2)
    if true {
        if (42 >= -201) {
            v3 = v3
        }
        if false {
            switch 0 {
            case -813:
                v1 += 0
            }
            if len(v4) > 0 {
                v4[0] = ("foo bar" + ("a" + v3))
            }
            var v23 float64 = (float64(1.5) + (-2.25 * 1e10))
            _ = v23
        }
        var v24 float32 = float32(0.1)
        _ = v24
    }
    v25 := GenS2{F0: 952, F1: ((v3 + "a") + (v3 + "foo bar")), F2: 1}
    switch 127 {
    case 32768:
        for fi26 := 0; fi26 < 7; fi26++ {
            var v27 float64 = (-2.25 - (0.0 * 1e10))
            _ = fi26
            _ = v27
        }
    default:
        var v28 int8 = 82
        _ = v28
    }
    if true {
        for fi29 := 0; fi29 < 17; fi29++ {
            if len(v4) > 0 {
                v4[0] = (v3 + "foo bar")
            }
            for fi30 := 0; fi30 < 12; fi30++ {
                var v31 uint64 = 4294967295
                _ = fi30
                _ = v31
            }
            fmt.Println(v2)
            _ = fi29
        }
        var v32 int = 32767
        v33 := GenS2{F0: 278, F1: v3, F2: genFn1()}
        _ = v32
        _ = v33
    } else {
        var v34 uint32 = 2
        v34 += (v34 + (v34 - 42))
        _ = v34
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 uint16 = 1
        _ = fi35
        _ = v36
    }
    var v37 uint64 = 2
    v38 := GenS2{F0: 42, F1: "foo bar", F2: genFn1()}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
