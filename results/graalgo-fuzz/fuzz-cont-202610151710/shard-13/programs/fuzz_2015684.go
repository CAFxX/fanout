package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint; F2 int }
func genFn0(p0 int8, p1 string, p2 uint) string {
    var v7 string = ((p1 + "a") + p1)
    if (255 > 809) {
        var v8 string = ((v7 + v7) + (("foo bar" + "foo bar") + ""))
        v8 = "a"
        if (-1 == -1) {
            var v9 uint8 = 177
            var v10 uint = (p2 << ((255) & 63))
            _ = v10
            _ = v9
        }
        _ = v8
    } else {
        v11 := GenS1{F0: int(p0), F1: -1, F2: false}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}
func genFn1() uint {
    var v12 int8 = 50
    v13 := []int{0, -730}
    if true {
        _ = len(v13)
        var v14 int64 = -9223372036854775808
        _ = v14
    }
    v12 -= 42
    _ = v12
    _ = v13
    return 773
}

func main() {
    var v1 uint64 = 1
    var v2 bool = true
    var v3 uint64 = v1
    v4 := make(map[uint32]uint32)
    var v5 bool = v2
    var v6 int8 = 1
    var v15 uint8 = 2
    _ = len(v4)
    for fi16 := 0; fi16 < 12; fi16++ {
        delete(v4, 189)
        var v17 uint32 = 2147483647
        _ = fi16
        _ = v17
    }
    var v18 string = "foo bar"
    if (true || (!((0 >= 1) && true))) {
        v18 = "foo bar"
    } else {
        var v19 int64 = -184
        var v20 float64 = (((1.5 + (((((-2.25 + ((-2.25 + 0.1) - (3.14159 / (float64(1.5) / 3.14159)))) - -2.25) + 1.5) + (0.0 * -2.25)) + 1e10)) + 0.0) / ((1e10 - float64((3.14159 * float32(float64(1e10))))) - 0.0))
        var v21 float64 = v20
        _ = v19
        _ = v20
        _ = v21
    }
    if v5 {
        if v2 {
            switch 944 {
            case 127:
                v22 := GenS2{F0: (v5 && v5)}
                _ = v22
            case -760:
                v18 = ("" + "\n\t")
            case 4294967295:
                for fi23 := 0; fi23 < 4; fi23++ {
                    var v24 uint64 = (v1 % (((v3 + 0)) | 1))
                    _ = fi23
                    _ = v24
                }
            default:
                for fi25 := 0; fi25 < 15; fi25++ {
                    v18 = v18
                    for fi26 := 0; fi26 < 13; fi26++ {
                        v15 -= (v15 << ((4294967295) & 63))
                        if (v2 && v2) {
                            v3 = 341
                            if (!v5) {
                                var v27 int32 = 700
                                _ = v27
                            }
                            v3 -= v3
                        } else {
                            var v28 float64 = ((0.0 * 3.14159) / 1.5)
                            _ = v28
                        }
                        var v29 float32 = 3.14159
                        var v30 uint64 = (^(v1))
                        _ = fi26
                        _ = v29
                        _ = v30
                    }
                    var v31 uint64 = 1
                    _ = fi25
                    _ = v31
                }
            }
            v32 := []int{-381, 32767}
            _ = v32
        }
        v3 -= (v1 | (v3 & (v1 % (((v1 + (v3 - v3))) | 1))))
        switch v18 {
        case "":
            v33 := v4[599]
            _ = v33
        }
    }
    v15 = uint8(v3)
    v34 := GenS3{F0: 2147483647, F1: genFn1(), F2: 128}
    for fi35 := 0; fi35 < 2; fi35++ {
        fmt.Println(v15, v6)
        _ = fi35
    }
    v18 = "a"
    v36 := GenS2{F0: v2}
    var v37 int8 = v6
    fmt.Println(v1)
    var v38 int64 = 0
    for fi39 := 0; fi39 < 15; fi39++ {
        _ = len(v4)
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
