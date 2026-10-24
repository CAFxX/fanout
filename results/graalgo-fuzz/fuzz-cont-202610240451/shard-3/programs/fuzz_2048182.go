package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int16; F2 int8 }
type GenS4 struct { F0 uint; F1 int16; F2 string }
type GenS5 struct { F0 int; F1 uint; F2 int64 }
type GenS6 struct { F0 int16 }
func genFn0(p0 int16) uint16 {
    p0 -= (p0 - (p0 * p0))
    var v7 bool = (!false)
    var v8 uint16 = 710
    v9 := GenS1{F0: 2}
    if v7 {
        switch "hello" {
        case "hello":
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := GenS1{F0: 4294967295}
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 uint = 233
                    var v14 int32 = 2147483647
                    var v15 uint = 18446744073709551615
                    v13 = (^(v13))
                    var v16 int = (fi12 + fi12)
                    v17 := GenS1{F0: 0}
                    v16 += v16
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                var v18 float32 = (1e10 + (0.1 + 1e10))
                _ = fi10
                _ = v11
                _ = v18
            }
        case "foo bar":
            v19 := GenS1{F0: 65535}
            _ = v19
        default:
            v20 := GenS1{F0: 2}
            _ = v20
        }
    } else {
        var v21 int64 = -653
        for fi22 := 0; fi22 < 2; fi22++ {
            if true {
                v23 := GenS2{F0: (v8 ^ (v8 & v8))}
                v7 = ((128 < 128) && true)
                _ = v23
            }
            for fi24 := 0; fi24 < 20; fi24++ {
                v25 := GenS1{F0: 1}
                _ = fi24
                _ = v25
            }
            _ = fi22
        }
        _ = v21
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return v8
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-104, -585, 32767}
    var v3 uint = 2
    var v4 uint8 = 45
    var v5 uint64 = 18446744073709551615
    var v6 int16 = 1
    v4 = v4
    delete(v1, 65535)
    v26 := GenS1{F0: 65535}
    switch "a" {
    case "foo bar":
        delete(v1, 1)
    case "a":
        _ = len(v1)
    default:
        switch 94 {
        case 1:
            var v27 int8 = int8(v5)
            _ = v27
        case 128:
            _ = len(v1)
        default:
            var v28 int8 = -128
            _ = v28
        }
    }
    var v29 uint64 = (v5 - 42)
    fmt.Println(v3, v5)
    v2 = append(v2, 127)
    if true {
        v30 := GenS2{F0: genFn0(0)}
        v31 := GenS1{F0: 188}
        _ = v30
        _ = v31
    } else {
        v2 = append(v2, -725)
        var v32 float64 = 0.0
        var v33 int64 = 363
        _ = v32
        _ = v33
    }
    var v34 uint = (v3 ^ 0)
    var v35 uint = 0
    for fi36 := 0; fi36 < 17; fi36++ {
        var v37 int32 = -1
        _ = fi36
        _ = v37
    }
    switch -2147483648 {
    case 799:
        if (0.0 >= float64(1e10)) {
            var v38 int16 = (v6 % ((42) | 1))
            v34 = (^(v35))
            _ = v38
        } else {
            for fi39 := 0; fi39 < 13; fi39++ {
                for fi40 := 0; fi40 < 9; fi40++ {
                    switch "a" {
                    case "":
                        if (!(!(1000 > (v6 + (-(v6)))))) {
                            if false {
                                v5 = v5
                                var v41 uint16 = 475
                                var v42 float64 = 3.14159
                                v6 += (v6 * (v6 << ((v3) & 63)))
                                _ = v41
                                _ = v42
                            } else {
                                v43 := GenS2{F0: 32767}
                                v29 += v29
                                v2 = append(v2, (fi40 / (((fi39 / (((fi40 + 32767)) | 1))) | 1)))
                                _ = v43
                            }
                            v44 := GenS3{F0: "a", F1: (^(v6)), F2: int8(v4)}
                            var v45 int16 = (v6 + v6)
                            _ = v44
                            _ = v45
                        } else {
                            var v46 uint64 = v29
                            fmt.Println(v4)
                            v47 := GenS4{F0: 44, F1: v6, F2: "\n\t"}
                            if (!true) {
                                v6 -= (v6 * (^(v6)))
                                fmt.Println(v34)
                            } else {
                                v5 = (v29 / (((^(v29))) | 1))
                                v1[2147483647] = "foo bar"
                                if true {
                                    if (!true) {
                                        v29 -= v46
                                        v46 = 692
                                        var v48 uint32 = 2147483647
                                        _ = v48
                                    }
                                    var v49 int = fi39
                                    _ = v49
                                } else {
                                    if true {
                                        var v50 uint32 = 2147483647
                                        _ = v50
                                    }
                                }
                            }
                            var v51 int = fi40
                            _ = v46
                            _ = v47
                            _ = v51
                        }
                    case "hello":
                        fmt.Println(v29, v29)
                    case "a":
                        if len(v2) > 0 {
                            v2[0] = (fi39 >> ((65535) & 63))
                        }
                    }
                    var v52 int32 = 2147483647
                    _ = fi40
                    _ = v52
                }
                if (!(197 != uint32(v4))) {
                    v2 = append(v2, (^(fi39)))
                    var v53 uint32 = 0
                    v54 := GenS1{F0: 2147483647}
                    _ = v53
                    _ = v54
                }
                for fi55 := 0; fi55 < 9; fi55++ {
                    fmt.Println(v6, v29)
                    var v56 int32 = -513
                    if (v35 != (v34 % (((^(v34))) | 1))) {
                        fmt.Println(v5, v29, v4)
                        var v57 float32 = 0.1
                        _ = v57
                    }
                    _ = fi55
                    _ = v56
                }
                _ = fi39
            }
            switch "foo bar" {
            case "a":
                var v58 uint8 = 128
                _ = v58
            }
            v5 = (v29 % ((2) | 1))
        }
    case 42:
        fmt.Println(v34)
    case -9223372036854775808:
        var v59 float32 = ((1e10 * (3.14159 * (-2.25 - -2.25))) + float32(float64(1.5)))
        _ = v59
    default:
        v60 := GenS5{F0: 2147483647, F1: 42, F2: 0}
        _ = v60
    }
    v61 := GenS6{F0: (v6 - (^(v6)))}
    var v62 uint8 = (v4 >> ((uint(v5)) & 63))
    v34 += v35
    delete(v1, 65535)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
}
