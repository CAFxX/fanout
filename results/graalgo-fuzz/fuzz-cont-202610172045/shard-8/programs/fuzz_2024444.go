package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = false
    var v3 uint32 = 676
    if (0 <= 0) {
        v4 := []int{32767, 4294967295, 1, 127}
        if v2 {
            switch 127 {
            case 32767:
                var v5 int8 = 0
                _ = v5
            default:
                v6 := GenS1{F0: true}
                _ = v6
            }
        }
        var v7 uint32 = (^(v3))
        _ = v4
        _ = v7
    } else {
        var v8 uint8 = 2
        if ((!v2) || (-159 >= -1)) {
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := GenS1{F0: false}
                fmt.Println(v3, v8)
                if (!(-2.25 > 0.0)) {
                    var v11 int16 = 1000
                    _ = v11
                } else {
                    var v12 uint = 1
                    v3 = 860
                    v13 := []int{}
                    _ = v12
                    _ = v13
                }
                _ = fi9
                _ = v10
            }
            for fi14 := 0; fi14 < 9; fi14++ {
                v2 = v2
                v3 = (v3 | 608)
                var v15 int8 = 127
                _ = fi14
                _ = v15
            }
            var v16 int = 32768
            _ = v16
        } else {
            v8 = (v8 / (((v8 - (v8 ^ (^(v8))))) | 1))
            for fi17 := 0; fi17 < 16; fi17++ {
                if (false && v2) {
                    for fi18 := 0; fi18 < 7; fi18++ {
                        v2 = true
                        var v19 string = "a"
                        v20 := GenS1{F0: v2}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    if v2 {
                        delete(v1, 127)
                    } else {
                        fmt.Println(fi17)
                        var v21 int8 = 0
                        _ = v21
                    }
                } else {
                    for fi22 := 0; fi22 < 5; fi22++ {
                        v3 -= v3
                        _ = fi22
                    }
                    if (v2 && (4294967295 >= 18446744073709551615)) {
                        v8 += (^(v8))
                    }
                    var v23 int = fi17
                    _ = v23
                }
                v3 += (v3 * (v3 << ((366) & 63)))
                _ = fi17
            }
            if (1e10 >= float32(-2.25)) {
                var v24 uint8 = 147
                var v25 uint = 255
                if v2 {
                    _ = len(v1)
                } else {
                    var v26 float32 = -2.25
                    _ = v26
                }
                _ = v24
                _ = v25
            }
        }
        _ = v8
    }
    v27 := GenS1{F0: (!false)}
    var v28 float64 = ((0.0 * 1.5) / 3.14159)
    if v2 {
        _ = len(v1)
    }
    var v29 uint = 65535
    v30 := GenS1{F0: false}
    v28 = v28
    var v31 int16 = -32768
    var v32 uint = 4294967295
    v31 += v31
    v3 = v3
    var v33 uint32 = v3
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 bool = (!v2)
        _ = fi34
        _ = v35
    }
    v31 -= v31
    if (v2 || v2) {
        if (false || (!v2)) {
            for fi36 := 0; fi36 < 17; fi36++ {
                v37 := GenS2{F0: fi36, F1: -467}
                var v38 bool = v2
                _ = fi36
                _ = v37
                _ = v38
            }
        }
        for fi39 := 0; fi39 < 9; fi39++ {
            var v40 int8 = -1
            v41 := GenS2{F0: (fi39 & fi39), F1: -1}
            _ = len(v1)
            _ = fi39
            _ = v40
            _ = v41
        }
        v31 -= v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
