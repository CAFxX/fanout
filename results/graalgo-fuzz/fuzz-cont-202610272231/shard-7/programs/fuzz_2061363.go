package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 int8 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint32 }
func genFn0(p0 int32, p1 int, p2 int64) int32 {
    v7 := GenS1{F0: 42, F1: (p0 % (((p0 << ((42) & 63))) | 1)), F2: 0}
    p2 = 508
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 407
    var v4 uint16 = v3
    v5 := make(map[string]uint32)
    var v6 float32 = (0.0 - (float32((((float64((0.0 + float32(0.1))) - 1e10) / -2.25) * 1.5)) * (float32((float64(-2.25) * 0.0)) - -2.25)))
    for fi8 := 0; fi8 < 15; fi8++ {
        if false {
            v3 = (v4 - 32767)
            delete(v1, 0)
        }
        _ = fi8
    }
    v3 -= 32767
    fmt.Println(v4)
    v9 := []int{4294967295, 32768, int(v2)}
    var v10 string = "hello"
    var v11 uint32 = 1
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 float32 = v6
            if len(v9) > 0 {
                v9[0] = (fi13 << ((916) & 63))
            }
            if true {
                v6 += 3.14159
                v2 = v2
                v2 = (^(v2))
            } else {
                v15 := GenS2{F0: 93, F1: 1, F2: -76}
                var v16 uint = 335
                _ = v15
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v17 uint16 = (v3 / (((v4 | (v3 >> ((236) & 63)))) | 1))
    v10 = ("hello" + v10)
    v1[(v11 | (v11 * 4294967295))] = (v11 / (((v11 >> ((1) & 63))) | 1))
    delete(v5, "a")
    var v18 bool = ((0 < v2) || (!true))
    if true {
        switch 669 {
        case -234:
            if (v18 && (116 >= 0)) {
                switch -497 {
                case -489:
                    v19 := GenS2{F0: 0, F1: 1, F2: -118}
                    _ = v19
                case 765:
                    var v20 uint = 944
                    _ = v20
                case 4294967295:
                    for fi21 := 0; fi21 < 2; fi21++ {
                        delete(v5, v10)
                        _ = fi21
                    }
                default:
                    _ = len(v9)
                }
            }
        default:
            switch 4294967295 {
            case -948:
                v18 = (-1 >= -998)
            case 128:
                if (194 != 42) {
                    if (v11 < v11) {
                        for fi22 := 0; fi22 < 20; fi22++ {
                            switch -928 {
                            case 337:
                                for fi23 := 0; fi23 < 19; fi23++ {
                                    v4 -= uint16(v11)
                                    var v24 int64 = (v2 / ((v2) | 1))
                                    var v25 uint = 1
                                    var v26 uint = v25
                                    v9 = append(v9, int(v2))
                                    v27 := v1[864]
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                            default:
                                var v28 uint8 = 0
                                _ = v28
                            }
                            if (42 > -1) {
                                v2 = 0
                                var v29 float32 = (0.1 * v6)
                                var v30 float32 = float32(1.5)
                                _ = v29
                                _ = v30
                            }
                            _ = fi22
                        }
                        v2 = 0
                        var v31 bool = v18
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 1; fi32++ {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        var v33 int64 = 823
                        _ = fi32
                        _ = v33
                    }
                } else {
                    fmt.Println(v11, v18, v17)
                    for fi34 := 0; fi34 < 16; fi34++ {
                        var v35 float64 = (0.1 / float64((v6 - float32(3.14159))))
                        for fi36 := 0; fi36 < 12; fi36++ {
                            switch ("foo bar" + "\n\t") {
                            case "hello":
                                v18 = v18
                                delete(v1, v11)
                                if (v18 && (v18 || (!(v2 > 359)))) {
                                    var v37 string = "hello"
                                    var v38 uint16 = v4
                                    _ = v37
                                    _ = v38
                                } else {
                                    if v18 {
                                        var v39 uint64 = 4294967295
                                        v40 := GenS3{F0: 2147483648, F1: 0, F2: 2}
                                        v1[2] = (v11 ^ (^(v11)))
                                        _ = v39
                                        _ = v40
                                    } else {
                                        if (!(v10 <= "a")) {
                                            if ((v11 >= 65535) && v18) {
                                                var v41 float64 = (1e10 / 0.1)
                                                var v42 float64 = (((v35 - float64(v6)) + -2.25) / v41)
                                                if v18 {
                                                    v3 = v4
                                                    v43 := GenS1{F0: (v17 >> ((42) & 63)), F1: -58, F2: 622}
                                                    _ = v43
                                                }
                                                _ = v41
                                                _ = v42
                                            }
                                            var v44 float32 = v6
                                            v9 = append(v9, fi34)
                                            _ = v44
                                        } else {
                                            var v45 int = -9223372036854775808
                                            _ = v45
                                        }
                                        v10 = ((v10 + (v10 + ((v10 + "hello") + (v10 + "a")))) + (v10 + v10))
                                    }
                                    fmt.Println(v35, v4)
                                }
                            case "\n\t":
                                fmt.Println(v3, fi34)
                            case "foo bar":
                                var v46 string = ""
                                _ = v46
                            }
                            fmt.Println(v10)
                            if (!(-32768 != 1000)) {
                                var v47 float32 = v6
                                var v48 uint64 = 18446744073709551615
                                switch (v10 + (v10 + v10)) {
                                case "a":
                                    var v49 uint64 = (v48 / (((v48 + v48)) | 1))
                                    _ = v49
                                }
                                var v50 uint16 = v4
                                fmt.Println(v2, v6, v4)
                                _ = v47
                                _ = v48
                                _ = v50
                            }
                            fmt.Println(v18)
                            _ = fi36
                        }
                        _ = fi34
                        _ = v35
                    }
                    _ = len(v1)
                }
            case 1:
                v51 := GenS2{F0: 0, F1: 42, F2: 127}
                _ = v51
            }
        }
    }
    var v52 int8 = -128
    v9 = append(v9, 2147483647)
    var v53 uint32 = v11
    if (!v18) {
        v53 = v11
    }
    v6 += v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v9
}
