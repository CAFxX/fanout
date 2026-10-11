package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint64; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 bool }
func genFn0(p0 int8) uint16 {
    var v4 bool = (false && false)
    for fi5 := 0; fi5 < 15; fi5++ {
        v4 = v4
        v6 := []int{1}
        _ = fi5
        _ = v6
    }
    v4 = v4
    _ = p0
    _ = v4
    return 86
}
func genFn1() uint32 {
    switch "" {
    case "foo bar":
        var v7 uint = 2
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int16 = 0
            _ = fi8
            _ = v9
        }
    default:
        v10 := GenS1{F0: 65535}
        _ = v10
    }
    var v11 int = 981
    if false {
        var v12 int16 = 1
        _ = v12
    }
    switch (v11 ^ v11) {
    case 32767:
        var v13 uint = 4294967295
        _ = v13
    case 127:
        v11 -= v11
    default:
        v11 += (v11 % (((v11 % ((-852) | 1))) | 1))
    }
    switch ("hello" + ("\n\t" + (("a" + "a") + (("a" + "foo bar") + "")))) {
    case "foo bar":
        if false {
            var v14 float64 = -2.25
            _ = v14
        }
    case "":
        v15 := []int{(v11 ^ v11), v11, (v11 >> ((42) & 63))}
        _ = v15
    case "\n\t":
        if (false || true) {
            for fi16 := 0; fi16 < 15; fi16++ {
                v17 := GenS1{F0: 255}
                _ = fi16
                _ = v17
            }
        }
    default:
        switch v11 {
        case 1:
            var v18 float32 = (float32(((-2.25 * 0.0) + (float64((0.0 / -2.25)) * ((3.14159 + (0.1 * 0.0)) * 1e10)))) - (0.1 - 1.5))
            _ = v18
        case 0:
            v11 -= 245
        }
    }
    _ = v11
    return 770
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{32768}
    var v3 uint32 = 2
    if false {
        delete(v1, (v3 - (v3 >> ((65535) & 63))))
    }
    switch "foo bar" {
    case "a":
        v3 += 4294967295
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        var v19 uint = 611
        _ = v19
    default:
        v20 := GenS2{F0: 796, F1: 0, F2: -2147483648}
        _ = v20
    }
    if false {
        if true {
            if true {
                v21 := GenS3{F0: 0}
                var v22 int = 810
                var v23 float64 = (-2.25 - float64(1.5))
                _ = v21
                _ = v22
                _ = v23
            }
            v3 -= (^(v3))
            if len(v2) > 0 {
                v2[0] = -1
            }
        }
    }
    switch (("hello" + "a") + "") {
    case "\n\t":
        for fi24 := 0; fi24 < 17; fi24++ {
            v3 = (v3 & (v3 + 1))
            v3 = v3
            _ = fi24
        }
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        _ = len(v2)
        fmt.Println(fi25)
        _ = fi25
    }
    switch "foo bar" {
    case "a":
        v2 = append(v2, -586)
    default:
        if false {
            var v26 int64 = 42
            v26 = v26
            _ = v26
        } else {
            v27 := GenS1{F0: 4294967295}
            v3 = v3
            if false {
                if true {
                    v3 -= genFn1()
                }
                if true {
                    switch "" {
                    case "a":
                        for fi28 := 0; fi28 < 2; fi28++ {
                            v29 := GenS3{F0: -2147483648}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi28
                            _ = v29
                        }
                    }
                    fmt.Println(v3)
                }
                v30 := v1[65535]
                _ = v30
            } else {
                v31 := GenS1{F0: 943}
                v3 -= genFn1()
                _ = v31
            }
            _ = v27
        }
    }
    fmt.Println(v3)
    var v32 int8 = 100
    fmt.Println(v3, v3)
    var v33 bool = (!true)
    fmt.Println(v32)
    v34 := GenS4{F0: 215, F1: ((2 > 18446744073709551615) || true)}
    if v33 {
        v1[(v3 + v3)] = 157
        v35 := GenS4{F0: 255, F1: v33}
        _ = v35
    } else {
        if false {
            fmt.Println(v3, v33)
            var v36 uint = 692
            _ = v36
        } else {
            for fi37 := 0; fi37 < 1; fi37++ {
                for fi38 := 0; fi38 < 1; fi38++ {
                    var v39 bool = (1e10 >= 0.1)
                    var v40 uint64 = 1
                    v39 = (!v33)
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                _ = fi37
            }
            for fi41 := 0; fi41 < 20; fi41++ {
                fmt.Println(v33)
                _ = len(v2)
                _ = fi41
            }
        }
        v3 = v3
    }
    delete(v1, 93)
    if (4294967295 > 0) {
        var v42 string = ""
        _ = v42
    } else {
        var v43 string = "a"
        v33 = v33
        v33 = v33
        _ = v43
    }
    if v33 {
        var v44 uint = 670
        _ = v44
    }
    fmt.Println(v33)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
