package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint8; F2 int64 }
func genFn0(p0 int16) string {
    v7 := []int{2147483647, 42}
    for fi8 := 0; fi8 < 4; fi8++ {
        p0 = 0
        _ = fi8
    }
    switch -9223372036854775808 {
    case 429:
        if len(v7) > 0 {
            v7[0] = -9223372036854775808
        }
    case 1:
        for fi9 := 0; fi9 < 16; fi9++ {
            if (float32((((0.0 + (-2.25 / (-2.25 / 1.5))) - 0.1) - 1e10)) >= (0.0 / 1.5)) {
                var v10 int32 = 0
                for fi11 := 0; fi11 < 6; fi11++ {
                    v7 = append(v7, (fi11 + fi11))
                    p0 = (p0 & (p0 & (p0 - 0)))
                    switch ("hello" + "") {
                    case "foo bar":
                        fmt.Println(p0, p0)
                    case "\n\t":
                        if false {
                            p0 += (p0 | (p0 * (p0 << ((65535) & 63))))
                        } else {
                            if false {
                                v12 := GenS1{F0: 0, F1: 967}
                                v10 += (v10 / (((v10 - -782)) | 1))
                                v10 = v10
                                _ = v12
                            }
                            var v13 int8 = 42
                            var v14 int32 = (v10 % ((v10) | 1))
                            _ = v13
                            _ = v14
                        }
                    case "":
                        var v15 int16 = (p0 << ((0) & 63))
                        _ = v15
                    }
                    _ = fi11
                }
                _ = v10
            }
            p0 = (p0 * (p0 / ((p0) | 1)))
            for fi16 := 0; fi16 < 17; fi16++ {
                v17 := GenS2{F0: 127, F1: 65535}
                _ = fi16
                _ = v17
            }
            _ = fi9
        }
    case 299:
        if false {
            var v18 int = 2147483647
            _ = v18
        }
    default:
        _ = len(v7)
    }
    _ = p0
    _ = v7
    return ""
}

func main() {
    var v1 bool = (668 != 4294967295)
    var v2 bool = false
    var v3 uint = 2
    v4 := []int{32768}
    v5 := []int{}
    v6 := make(map[string]bool)
    if v2 {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        v19 := GenS3{F0: 42, F1: 255, F2: -104}
        _ = v19
    }
    fmt.Println(v2, v2, v2)
    v20 := GenS1{F0: 1000, F1: (v3 & v3)}
    var v21 string = ("hello" + "a")
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 uint32 = 0
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 uint8 = 241
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        var v27 int16 = 1000
        _ = fi26
        _ = v27
    }
    v1 = (!((true && v2) && true))
    fmt.Println(v21, v1)
    var v28 uint16 = 0
    v29 := GenS3{F0: -670, F1: 0, F2: 9223372036854775807}
    v30 := GenS2{F0: -128, F1: (v28 / ((v28) | 1))}
    v3 -= 944
    switch 213 {
    case 0:
        var v31 uint16 = 1
        _ = v31
    }
    switch -2147483648 {
    case 32768:
        if len(v5) > 0 {
            v5[0] = 835
        }
    case 9223372036854775807:
        switch -2147483648 {
        case 32767:
            v3 -= v3
        case 98:
            v6[v21] = (true || (2147483648 != 953))
        case 794:
            for fi32 := 0; fi32 < 17; fi32++ {
                var v33 int64 = 42
                _ = fi32
                _ = v33
            }
        default:
            for fi34 := 0; fi34 < 7; fi34++ {
                var v35 int = (fi34 - 32768)
                v36 := GenS3{F0: int64(v28), F1: 127, F2: 997}
                _ = fi34
                _ = v35
                _ = v36
            }
        }
    default:
        v2 = (v2 || true)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v37 := GenS3{F0: 0, F1: 255, F2: 955}
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
