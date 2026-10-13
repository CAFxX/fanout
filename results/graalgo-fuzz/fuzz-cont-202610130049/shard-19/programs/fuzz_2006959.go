package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int64) int8 {
    if (1 > 1) {
        var v5 int64 = (^(p0))
        var v6 float32 = float32((((((0.0 / 1.0) * 0.0) - ((((0.0 * 3.14159) * (3.14159 - (((0.1 - float64(-2.25)) + -2.25) + 0.1))) * 1e10) / 3.14159)) + 3.14159) / 1e10))
        v7 := GenS1{F0: 42}
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := []int{}
    _ = p0
    _ = v8
    return 42
}

func main() {
    var v1 int32 = -126
    v2 := []int{4294967295, 9223372036854775807, 127}
    var v3 uint16 = 2
    v4 := make(map[uint32]bool)
    _ = len(v4)
    v1 -= (v1 % (((v1 + (v1 - -2147483648))) | 1))
    switch -469 {
    case 1:
        var v9 string = "\n\t"
        _ = v9
    default:
        if (((3.14159 - -2.25) < 3.14159) && false) {
            v3 += (v3 + (v3 << ((255) & 63)))
            var v10 int64 = 42
            _ = v10
        }
    }
    var v11 int16 = -1
    var v12 int64 = 1
    for fi13 := 0; fi13 < 9; fi13++ {
        if (692 >= 42) {
            switch ("" + "a") {
            case "foo bar":
                var v14 int64 = (v12 & v12)
                _ = v14
            case "":
                if true {
                    var v15 int64 = v12
                    _ = v15
                } else {
                    v3 = (^(v3))
                }
            case "a":
                v3 = (^(v3))
            }
        }
        if false {
            fmt.Println(v3)
            v3 = 65535
            for fi16 := 0; fi16 < 13; fi16++ {
                delete(v4, 0)
                _ = len(v4)
                var v17 int16 = int16(v12)
                _ = fi16
                _ = v17
            }
        } else {
            v18 := GenS1{F0: v11}
            if false {
                v11 = v11
            }
            if (42 > (v11 ^ (-(v11)))) {
                v19 := GenS1{F0: (v11 * (v11 | 0))}
                var v20 int64 = (v12 + -9223372036854775808)
                _ = v19
                _ = v20
            }
            _ = v18
        }
        delete(v4, 743)
        _ = fi13
    }
    v11 += v11
    fmt.Println(v11, v1, v1)
    v21 := GenS1{F0: (v11 - (v11 & (v11 - (v11 * (v11 | v11)))))}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, 32768)
    switch 9223372036854775807 {
    case 42:
        v3 = v3
    case -1:
        for fi22 := 0; fi22 < 3; fi22++ {
            if (25 > 42) {
                if false {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v12 = v12
                        var v24 uint32 = 935
                        _ = fi23
                        _ = v24
                    }
                    v12 = 1
                } else {
                    if false {
                        v25 := GenS1{F0: 1000}
                        v12 += v12
                        _ = v25
                    } else {
                        for fi26 := 0; fi26 < 14; fi26++ {
                            if false {
                                if false {
                                    if false {
                                        var v27 uint64 = 187
                                        v28 := GenS1{F0: (v11 + (^(v11)))}
                                        v29 := v4[uint32(v27)]
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    }
                                    v1 = v1
                                }
                                var v30 uint16 = v3
                                _ = v30
                            }
                            _ = fi26
                        }
                    }
                    if true {
                        var v31 uint32 = 458
                        v3 += v3
                        v32 := GenS2{F0: 140, F1: "a"}
                        _ = v31
                        _ = v32
                    } else {
                        v1 = v1
                        v11 += (v11 + int16(v1))
                    }
                    if (false || false) {
                        if true {
                            for fi33 := 0; fi33 < 6; fi33++ {
                                var v34 uint16 = 0
                                var v35 int8 = -1
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            for fi36 := 0; fi36 < 15; fi36++ {
                                var v37 uint = 18446744073709551615
                                _ = fi36
                                _ = v37
                            }
                        }
                    }
                }
                fmt.Println(v3, v11)
                var v38 float64 = (1e10 - 0.1)
                _ = v38
            }
            v39 := GenS3{F0: 128}
            _ = fi22
            _ = v39
        }
    default:
        v40 := GenS2{F0: 0, F1: (("a" + "") + ("a" + ("hello" + "foo bar")))}
        _ = v40
    }
    if (v11 <= v11) {
        v3 = (v3 + v3)
        switch -1 {
        case 32768:
            fmt.Println(v1, v12, v11)
        case 53:
            var v41 int8 = -41
            _ = v41
        default:
            v42 := GenS4{F0: 0}
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
