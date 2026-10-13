package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint; F1 string }
func genFn0(p0 int64, p1 uint8, p2 int16) int32 {
    var v5 float32 = 0.0
    if false {
        var v6 int64 = -259
        switch 9223372036854775807 {
        case 4294967295:
            var v7 int8 = 0
            _ = v7
        case 42:
            var v8 int = 2147483647
            _ = v8
        case 128:
            var v9 uint16 = 194
            _ = v9
        default:
            v10 := GenS1{F0: p1, F1: 2, F2: 761}
            _ = v10
        }
        _ = v6
    }
    v11 := []int{-706, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v5
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    var v4 bool = true
    var v12 int8 = 100
    v13 := []int{127, int(v12), 9223372036854775807, -897}
    v4 = (!("" > ("hello" + ("\n\t" + "a"))))
    var v14 uint16 = 255
    v4 = true
    if v4 {
        switch 930 {
        case 329:
            v15 := v1[128]
            _ = v15
        case 2147483647:
            var v16 uint64 = 0
            _ = v16
        case 1:
            fmt.Println(v4, v14)
        default:
            v14 += (v14 & v14)
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            if v4 {
                v14 = 244
            }
            v4 = (2147483647 != 65535)
            _ = fi17
        }
    } else {
        v13 = append(v13, 32767)
        var v18 bool = (2 >= 2)
        _ = v18
    }
    v4 = v4
    v4 = (42 > 0)
    if ((0.0 == 0.1) && (false && v4)) {
        v4 = (!(v4 || true))
    } else {
        v14 -= (^(v14))
    }
    v12 = v12
    switch 4294967295 {
    case 957:
        v19 := v1[127]
        _ = v19
    case 9223372036854775807:
        if v4 {
            v4 = (!(9223372036854775807 <= 1))
            switch 42 {
            case -2147483648:
                for fi20 := 0; fi20 < 20; fi20++ {
                    fmt.Println(v12, fi20, v12)
                    _ = fi20
                }
            case 995:
                v21 := GenS1{F0: 150, F1: 18446744073709551615, F2: 866}
                _ = v21
            default:
                if true {
                    var v22 float64 = (3.14159 - 3.14159)
                    switch 32767 {
                    case -166:
                        v23 := GenS1{F0: 0, F1: 431, F2: 2147483648}
                        _ = v23
                    case -2147483648:
                        var v24 bool = ((591 > v14) || (-2147483648 <= 0))
                        _ = v24
                    default:
                        v22 = v22
                    }
                    if ((!v4) || false) {
                        for fi25 := 0; fi25 < 15; fi25++ {
                            v22 = 1.5
                            v26 := v3["foo bar"]
                            var v27 int64 = 979
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        var v28 uint64 = uint64(v14)
                        var v29 float64 = 0.1
                        _ = v28
                        _ = v29
                    } else {
                        if v4 {
                            var v30 uint64 = 1
                            var v31 int = 42
                            _ = v30
                            _ = v31
                        }
                        v22 = 1.5
                        for fi32 := 0; fi32 < 1; fi32++ {
                            v22 += 1e10
                            if false {
                                v22 = 1.5
                            }
                            fmt.Println(v22, v22)
                            _ = fi32
                        }
                    }
                    _ = v22
                }
            }
            for fi33 := 0; fi33 < 10; fi33++ {
                var v34 int = (fi33 | fi33)
                _ = fi33
                _ = v34
            }
        }
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 uint8 = 255
        fmt.Println(v12, v12)
        v37 := GenS1{F0: v36, F1: 65535, F2: 2147483648}
        _ = fi35
        _ = v36
        _ = v37
    }
    v38 := GenS2{F0: 42, F1: "foo bar"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
