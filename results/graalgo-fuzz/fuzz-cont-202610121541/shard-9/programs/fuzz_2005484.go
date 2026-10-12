package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint16, p1 uint64, p2 string) bool {
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(p2, p1)
        _ = fi6
    }
    p2 = "\n\t"
    v7 := []int{-1}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int8 = 0
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return true
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 bool = true
    v5 := make(map[string]string)
    fmt.Println(v4, v1)
    v1 = ((1e10 + v1) / (3.14159 + v1))
    if v4 {
        switch "\n\t" {
        case "\n\t":
            v1 = v1
        }
        v1 = v1
        switch 9223372036854775807 {
        case 1:
            switch -9223372036854775808 {
            case 9223372036854775807:
                delete(v2, 2)
            case -1:
                v1 = float32((0.0 * 1.5))
            case 1:
                if v4 {
                    if v4 {
                        for fi11 := 0; fi11 < 18; fi11++ {
                            var v12 uint = 42
                            _ = fi11
                            _ = v12
                        }
                        if (2 == 380) {
                            fmt.Println(v4, v4)
                            v13 := []int{32767}
                            _ = v13
                        }
                        var v14 bool = v4
                        _ = v14
                    } else {
                        switch "hello" {
                        case "foo bar":
                            switch 128 {
                            case 32767:
                                var v15 int32 = -428
                                _ = v15
                            case 2147483647:
                                v1 = 0.0
                            default:
                                var v16 int32 = 2147483647
                                _ = v16
                            }
                        case "hello":
                            var v17 uint16 = 255
                            _ = v17
                        case "\n\t":
                            if (v4 && false) {
                                v4 = (-1 != 42)
                            }
                        }
                        v18 := []int{-9223372036854775808}
                        _ = v18
                    }
                }
            }
        case 0:
            delete(v2, 65535)
        }
    }
    var v19 int8 = 0
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 string = ("foo bar" + ("\n\t" + "a"))
        _ = fi20
        _ = v21
    }
    if v4 {
        var v22 float32 = (v1 / 1.0)
        for fi23 := 0; fi23 < 7; fi23++ {
            v22 -= v1
            _ = fi23
        }
        _ = v22
    }
    if v4 {
        v1 = ((float32(0.1) - (v1 + v1)) / (float32(float64(0.1)) * v1))
        var v24 string = "a"
        for fi25 := 0; fi25 < 11; fi25++ {
            if v4 {
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 uint16 = 0
                    var v28 int8 = v19
                    var v29 int64 = -1
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
            } else {
                v19 = v19
                v30 := GenS1{F0: 9223372036854775807}
                _ = v30
            }
            _ = fi25
        }
        _ = v24
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        switch fi31 {
        case 32768:
            v32 := GenS1{F0: -9223372036854775808}
            _ = v32
        default:
            v19 = (v19 ^ (v19 ^ v19))
        }
        _ = fi31
    }
    v33 := []int{32768, 32767}
    v1 += (0.1 * float32(1e10))
    for fi34 := 0; fi34 < 12; fi34++ {
        switch "a" {
        case "hello":
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
        default:
            v4 = true
        }
        _ = fi34
    }
    var v35 uint8 = 127
    fmt.Println(v35)
    v35 = 0
    v33 = append(v33, int(v35))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v4
    _ = v5
}
