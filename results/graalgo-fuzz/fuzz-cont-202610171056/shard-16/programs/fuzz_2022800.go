package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint64, p1 int16) uint32 {
    var v5 uint64 = (^(p0))
    if true {
        p0 = 1
    } else {
        var v6 int8 = 42
        v7 := GenS1{F0: p1, F1: 2147483647}
        _ = v6
        _ = v7
    }
    var v8 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 42
}
func genFn1(p0 uint64, p1 int64) int8 {
    v9 := GenS2{F0: 42}
    v10 := GenS1{F0: 32767, F1: genFn0(2, 962)}
    if false {
        p1 -= p1
    }
    fmt.Println(p0, p1)
    v11 := GenS1{F0: -397, F1: 2}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return -66
}

func main() {
    v1 := []string{"a", ("hello" + "")}
    v2 := make(map[uint32]string)
    var v3 uint64 = 110
    var v4 bool = true
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 int8 = -60
        v4 = v4
        _ = fi12
        _ = v13
    }
    v14 := v2[genFn0((v3 + v3), -608)]
    switch v14 {
    case "":
        if true {
            if len(v1) > 0 {
                v1[0] = v14
            }
            var v15 uint32 = 1
            _ = v15
        } else {
            for fi16 := 0; fi16 < 11; fi16++ {
                if false {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        var v18 uint64 = (^(v3))
                        _ = fi17
                        _ = v18
                    }
                } else {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        v20 := GenS3{F0: 2147483648}
                        _ = fi19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v22 := GenS4{F0: (v14 >= v14), F1: 416}
                        _ = fi21
                        _ = v22
                    }
                }
                fmt.Println(v4, v4)
                for fi23 := 0; fi23 < 12; fi23++ {
                    fmt.Println(v4)
                    switch "" {
                    case "foo bar":
                        v24 := GenS4{F0: false, F1: -910}
                        _ = v24
                    default:
                        v4 = v4
                    }
                    if (1 <= 136) {
                        var v25 uint32 = genFn0((v3 >> ((0) & 63)), 1000)
                        _ = v25
                    }
                    _ = fi23
                }
                _ = fi16
            }
            var v26 int8 = 42
            _ = v26
        }
    case "foo bar":
        v27 := GenS1{F0: 32767, F1: 2147483647}
        _ = v27
    case "a":
        for fi28 := 0; fi28 < 12; fi28++ {
            fmt.Println(fi28, v14)
            _ = fi28
        }
    default:
        var v29 uint8 = 127
        _ = v29
    }
    fmt.Println(v4)
    if (-634 > -1) {
        var v30 int8 = 42
        var v31 int16 = -32768
        for fi32 := 0; fi32 < 6; fi32++ {
            _ = len(v2)
            for fi33 := 0; fi33 < 2; fi33++ {
                if v4 {
                    v34 := GenS5{F0: v30, F1: 255}
                    _ = v34
                } else {
                    var v35 uint8 = 127
                    _ = v35
                }
                _ = fi33
            }
            v36 := GenS5{F0: v30, F1: 255}
            _ = fi32
            _ = v36
        }
        _ = v30
        _ = v31
    } else {
        fmt.Println(v3, v4)
        for fi37 := 0; fi37 < 8; fi37++ {
            v3 = (v3 % ((v3) | 1))
            for fi38 := 0; fi38 < 2; fi38++ {
                v14 = "\n\t"
                v3 = 737
                _ = fi38
            }
            var v39 float32 = 3.14159
            _ = fi37
            _ = v39
        }
    }
    var v40 int16 = 42
    var v41 int64 = 0
    var v42 uint32 = 65535
    if (v14 <= v14) {
        fmt.Println(v14, v40, v42)
    }
    if v4 {
        switch -1 {
        case -360:
            var v43 float64 = (3.14159 / -2.25)
            _ = v43
        default:
            for fi44 := 0; fi44 < 18; fi44++ {
                var v45 int16 = 42
                _ = fi44
                _ = v45
            }
        }
        _ = len(v2)
    } else {
        delete(v2, v42)
    }
    var v46 int8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v46
}
