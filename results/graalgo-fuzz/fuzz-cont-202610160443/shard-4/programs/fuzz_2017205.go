package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 int32 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint16 }
func genFn0(p0 int16, p1 uint32) string {
    switch "hello" {
    case "\n\t":
        if (1 >= (p1 & (p1 << ((65535) & 63)))) {
            fmt.Println(p0, p0)
            var v4 int32 = 2147483647
            _ = v4
        }
    case "":
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 uint16 = 255
            if true {
                fmt.Println(p0, p0, v6)
            } else {
                if ((127 != 42) && (509 > 1)) {
                    fmt.Println(v6)
                } else {
                    v6 -= (v6 % ((v6) | 1))
                    var v7 int = -244
                    v8 := GenS1{F0: 18446744073709551615, F1: 100}
                    _ = v7
                    _ = v8
                }
                for fi9 := 0; fi9 < 3; fi9++ {
                    switch "\n\t" {
                    case "foo bar":
                        v10 := []int{2147483647, (fi5 - fi5), (fi9 & int(v6)), 200}
                        _ = v10
                    }
                    v11 := GenS2{F0: 858, F1: 660}
                    _ = fi9
                    _ = v11
                }
            }
            v12 := []int{(fi5 - 2147483647)}
            _ = fi5
            _ = v12
            _ = v6
        }
    }
    var v13 int32 = 42
    var v14 uint64 = 4294967295
    var v15 float64 = 0.0
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return "foo bar"
}

func main() {
    var v1 string = ""
    var v2 uint8 = 255
    var v3 int32 = -2147483648
    fmt.Println(v3, v1, v2)
    v16 := GenS1{F0: 42, F1: 42}
    v17 := GenS2{F0: 763, F1: 970}
    var v18 int32 = v3
    v18 += 42
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 uint8 = 182
        _ = fi19
        _ = v20
    }
    v21 := []int{-1}
    var v22 uint64 = 0
    switch v1 {
    case "\n\t":
        v23 := GenS3{F0: 255, F1: (v3 % ((int32(v22)) | 1)), F2: (v3 << ((42) & 63))}
        _ = v23
    default:
        var v24 uint8 = (v2 >> ((42) & 63))
        _ = v24
    }
    if false {
        switch v1 {
        case "a":
            var v25 int8 = 0
            _ = v25
        }
        if (!(!true)) {
            v18 -= -733
            if false {
                var v26 int = 537
                _ = v26
            }
        } else {
            var v27 uint16 = 549
            _ = v27
        }
        var v28 float64 = 0.0
        _ = v28
    } else {
        fmt.Println(v18, v1, v2)
    }
    var v29 bool = true
    var v30 float64 = ((1e10 + 0.0) / 1.5)
    v22 = (v22 >> ((255) & 63))
    if ((^(v22)) <= v22) {
        var v31 int8 = -1
        _ = v31
    } else {
        switch 4294967295 {
        case 9223372036854775807:
            if v29 {
                v29 = v29
                v32 := GenS3{F0: 696, F1: v18, F2: (v18 & 1)}
                _ = v32
            } else {
                v33 := GenS4{F0: 1, F1: 42, F2: 255}
                for fi34 := 0; fi34 < 17; fi34++ {
                    switch fi34 {
                    case 128:
                        switch -1 {
                        case 640:
                            var v35 uint32 = 507
                            _ = v35
                        case -814:
                            v2 = (v2 << ((18446744073709551615) & 63))
                        default:
                            v21 = append(v21, 4294967295)
                        }
                    default:
                        v36 := GenS4{F0: 65535, F1: (v2 >> ((1) & 63)), F2: 617}
                        _ = v36
                    }
                    var v37 uint8 = v2
                    v1 = (v1 + ((("" + v1) + "a") + v1))
                    _ = fi34
                    _ = v37
                }
                _ = v33
            }
        case -76:
            for fi38 := 0; fi38 < 13; fi38++ {
                if len(v21) > 0 {
                    v21[0] = fi38
                }
                switch 32768 {
                case 4294967295:
                    var v39 uint32 = 42
                    _ = v39
                }
                _ = fi38
            }
        case 4294967295:
            v18 = (v3 - v3)
        }
    }
    var v40 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v40
}
