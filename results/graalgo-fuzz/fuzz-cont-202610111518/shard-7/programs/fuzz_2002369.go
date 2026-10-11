package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint16; F1 int16; F2 int8 }
type GenS3 struct { F0 uint32; F1 int16; F2 string }
func genFn0(p0 uint32, p1 bool) int8 {
    var v5 uint32 = (p0 + p0)
    for fi6 := 0; fi6 < 6; fi6++ {
        if (float32(1e10) > -2.25) {
            v5 += 616
        }
        if (32768 != (fi6 >> ((158) & 63))) {
            v5 -= (^(v5))
            if true {
                p0 += v5
                if p1 {
                    var v7 string = ("a" + "\n\t")
                    _ = v7
                } else {
                    switch fi6 {
                    case -123:
                        fmt.Println(fi6, p1, v5)
                    case -9223372036854775808:
                        fmt.Println(p0, v5, v5)
                    }
                    var v8 int32 = 690
                    var v9 uint16 = 231
                    _ = v8
                    _ = v9
                }
                p0 -= v5
            }
        }
        v10 := GenS1{F0: -741, F1: false}
        _ = fi6
        _ = v10
    }
    if p1 {
        var v11 uint64 = 42
        v12 := []int{-21, 32767}
        switch -2147483648 {
        case -9223372036854775808:
            var v13 int = 80
            _ = v13
        }
        _ = v11
        _ = v12
    } else {
        var v14 int8 = -128
        if (60 != -9223372036854775808) {
            switch "" {
            case "":
                v15 := GenS1{F0: 1, F1: (0.0 <= 1.5)}
                _ = v15
            case "hello":
                var v16 bool = (p1 || (993 > 0))
                _ = v16
            case "a":
                for fi17 := 0; fi17 < 15; fi17++ {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        var v19 int64 = 0
                        fmt.Println(v14, fi18)
                        if (1 < 2) {
                            v20 := []int{(fi17 >> ((2) & 63))}
                            _ = v20
                        } else {
                            p0 = (v5 >> ((18446744073709551615) & 63))
                        }
                        v21 := GenS1{F0: -999, F1: (p1 || (1.5 >= 3.14159))}
                        p1 = (!(!p1))
                        v22 := []int{}
                        _ = fi18
                        _ = v19
                        _ = v21
                        _ = v22
                    }
                    p1 = ((((3.14159 / 0.1) + (3.14159 + 1e10)) != 1e10) || true)
                    fmt.Println(v5, v5)
                    _ = fi17
                }
            default:
                v5 += (p0 + (v5 * p0))
            }
        }
        switch 0 {
        case -317:
            var v23 string = "foo bar"
            _ = v23
        case -2147483648:
            v14 = (v14 % ((v14) | 1))
        }
        _ = v14
    }
    p0 -= v5
    _ = p0
    _ = p1
    _ = v5
    return 127
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 42
    var v3 float64 = (0.0 / ((1.5 / (3.14159 + -2.25)) - (1e10 * float64(1e10))))
    var v4 uint8 = 10
    if true {
        if true {
            v24 := GenS1{F0: 1000, F1: false}
            if (false || true) {
                v3 += v3
                var v25 string = ("foo bar" + "foo bar")
                _ = v25
            } else {
                v2 = (v2 >> ((255) & 63))
            }
            _ = v24
        }
        var v26 int32 = v2
        _ = v26
    } else {
        for fi27 := 0; fi27 < 19; fi27++ {
            v3 -= (3.14159 * (3.14159 * -2.25))
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 float32 = (0.0 * 0.1)
                var v30 int = (fi28 % ((fi28) | 1))
                var v31 string = "\n\t"
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            v3 = float64((1e10 * -2.25))
            _ = fi27
        }
    }
    var v32 uint16 = 255
    var v33 int8 = -1
    v34 := GenS2{F0: (^(v32)), F1: 1000, F2: (v33 >> ((284) & 63))}
    v35 := []int{190}
    fmt.Println(v33, v2, v32)
    var v36 uint64 = 4294967295
    if false {
        _ = len(v1)
        for fi37 := 0; fi37 < 2; fi37++ {
            v32 += v32
            v38 := v1["a"]
            _ = fi37
            _ = v38
        }
    } else {
        var v39 uint = uint(v36)
        _ = v39
    }
    v40 := GenS3{F0: 700, F1: -32768, F2: ("a" + "hello")}
    if false {
        v33 = (v33 / ((0) | 1))
        var v41 string = ("" + "")
        _ = v41
    } else {
        v33 -= (v33 >> ((uint(v36)) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
}
