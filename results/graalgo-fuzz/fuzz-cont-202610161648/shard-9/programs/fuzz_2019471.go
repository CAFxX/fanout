package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
func genFn0() uint {
    var v6 int = -708
    v6 = (^(v6))
    _ = v6
    return 1
}

func main() {
    var v1 int8 = 50
    var v2 string = ""
    var v3 bool = ((!false) || false)
    var v4 uint16 = 32767
    var v5 uint16 = (v4 ^ (v4 | (v4 % (((^(v4))) | 1))))
    v3 = (!((!v3) && (float32(1.5) > 0.0)))
    var v7 uint8 = 42
    v1 = v1
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(v5, fi8)
        _ = fi8
    }
    v7 = uint8(v5)
    fmt.Println(v5)
    for fi9 := 0; fi9 < 20; fi9++ {
        if true {
            v10 := []int{(fi9 ^ 2147483647), 0, (fi9 + fi9), (fi9 * 137)}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int8 = (v1 ^ (v1 & (-(v1))))
                _ = fi11
                _ = v12
            }
            v13 := []int{(fi9 ^ fi9)}
            if false {
                if len(v13) > 0 {
                    v13[0] = fi9
                }
                v14 := GenS1{F0: 154}
                switch fi9 {
                case -14:
                    if (v1 == (-(v1))) {
                        if v3 {
                            var v15 int = (fi9 % ((9223372036854775807) | 1))
                            v1 -= -128
                            _ = v15
                        } else {
                            switch ("foo bar" + "a") {
                            case "":
                                fmt.Println(v7, v1, v7)
                            default:
                                v5 = (v4 * v4)
                            }
                            if (v7 > v7) {
                                var v16 int32 = 0
                                var v17 string = v2
                                v13 = append(v13, 1)
                                _ = v16
                                _ = v17
                            }
                            v3 = v3
                        }
                        v4 += v5
                        switch "" {
                        case "":
                            var v18 float64 = ((0.1 * 0.0) / 1.0)
                            _ = v18
                        }
                    } else {
                        v13 = append(v13, (^(fi9)))
                    }
                case -9223372036854775808:
                    fmt.Println(v4)
                }
                _ = v14
            } else {
                v1 = int8(v4)
                v2 = ""
            }
            _ = v13
        }
        _ = fi9
    }
    var v19 int64 = -836
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS2{F0: 18446744073709551615}
        for fi22 := 0; fi22 < 15; fi22++ {
            switch v2 {
            case "\n\t":
                v2 = v2
            case "hello":
                var v23 int64 = (-(v19))
                _ = v23
            }
            if false {
                v24 := GenS2{F0: uint64(v4)}
                v25 := GenS3{F0: 2147483647}
                _ = v24
                _ = v25
            }
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    switch -1 {
    case 0:
        v26 := []int{127}
        _ = v26
    case -296:
        var v27 int64 = -92
        _ = v27
    case 32767:
        fmt.Println(v19, v19, v4)
    }
    var v28 string = v2
    v29 := []int{128}
    v30 := GenS3{F0: 726}
    v29 = append(v29, int(v19))
    var v31 int8 = -128
    fmt.Println(v4, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v7
}
