package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint64) string {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := []int{(fi5 + fi5), (fi5 ^ 630)}
        var v7 uint8 = 22
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{}
    if true {
        switch 4294967295 {
        case 32767:
            fmt.Println(p0)
        default:
            var v9 int32 = -1
            _ = v9
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            p0 = (^(p0))
            _ = fi10
        }
    }
    v8 = append(v8, -46)
    _ = p0
    _ = v8
    return ""
}
func genFn1(p0 uint8, p1 int16, p2 int8) uint16 {
    var v11 int8 = -1
    var v12 uint8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    var v3 float32 = (3.14159 * float32((0.0 * 1.5)))
    var v4 int32 = -900
    v13 := GenS1{F0: 2, F1: 211, F2: 100}
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS2{F0: 0}
        if true {
            for fi16 := 0; fi16 < 16; fi16++ {
                fmt.Println(v4, v3)
                _ = fi16
            }
            switch "" {
            case "a":
                v3 = v3
            }
            var v17 uint = 0
            _ = v17
        } else {
            if false {
                var v18 int8 = -3
                if true {
                    switch ("hello" + "") {
                    case "a":
                        v19 := GenS1{F0: 0, F1: 43, F2: v18}
                        _ = v19
                    default:
                        for fi20 := 0; fi20 < 1; fi20++ {
                            var v21 uint8 = 137
                            var v22 int = 127
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                }
                _ = v18
            }
            _ = len(v2)
        }
        if false {
            fmt.Println(v4, v3, v3)
        }
        _ = fi14
        _ = v15
    }
    v4 = (v4 & 1)
    var v23 int64 = -1
    v3 += (v3 / float32((float64(float32(1.5)) * (3.14159 + -2.25))))
    if true {
        var v24 int = 9223372036854775807
        v25 := GenS1{F0: 2, F1: 35, F2: 42}
        delete(v2, 0)
        _ = v24
        _ = v25
    } else {
        if (!false) {
            if len(v1) > 0 {
                v1[0] = 147
            }
            if false {
                v26 := GenS1{F0: 127, F1: 48, F2: -128}
                v3 -= (1.5 / v3)
                v27 := v2[2147483647]
                _ = v26
                _ = v27
            } else {
                v28 := GenS2{F0: 4294967295}
                v23 += v23
                v29 := GenS3{F0: true}
                _ = v28
                _ = v29
            }
        }
    }
    var v30 int = 2147483647
    var v31 int = v30
    var v32 int8 = 0
    var v33 string = ("" + "")
    _ = len(v2)
    for fi34 := 0; fi34 < 10; fi34++ {
        switch ("hello" + "") {
        case "":
            if false {
                var v35 bool = false
                if v35 {
                    for fi36 := 0; fi36 < 8; fi36++ {
                        var v37 int64 = (v23 ^ v23)
                        _ = fi36
                        _ = v37
                    }
                } else {
                    var v38 uint64 = 1
                    var v39 uint = 255
                    _ = v38
                    _ = v39
                }
                _ = len(v2)
                _ = v35
            }
        case "foo bar":
            v40 := GenS4{F0: -128}
            _ = v40
        }
        _ = fi34
    }
    v4 = (^(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
