package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int8; F2 int64 }
type GenS3 struct { F0 int64; F1 uint64; F2 int16 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint64 }
func genFn0() int32 {
    var v6 uint8 = 2
    var v7 uint64 = 2
    v7 -= (v7 + (v7 % (((v7 % (((^(v7))) | 1))) | 1)))
    v8 := GenS1{F0: -104}
    for fi9 := 0; fi9 < 9; fi9++ {
        if true {
            for fi10 := 0; fi10 < 14; fi10++ {
                if (!(435 < 65535)) {
                    var v11 int8 = 127
                    if true {
                        switch fi10 {
                        case -2147483648:
                            if true {
                                switch fi10 {
                                case 127:
                                    if true {
                                        v12 := []int{(^(fi9)), (fi10 % ((fi9) | 1))}
                                        _ = v12
                                    } else {
                                        v11 -= v11
                                        v11 = v11
                                    }
                                case 42:
                                    v7 = 838
                                case 186:
                                    var v13 int = fi10
                                    _ = v13
                                default:
                                    v14 := []int{}
                                    _ = v14
                                }
                                v15 := []int{fi9, 0, fi9, (fi10 * fi9)}
                                if (!false) {
                                    v11 = v11
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                } else {
                                    fmt.Println(v7)
                                    v6 -= v6
                                    if true {
                                        switch "hello" {
                                        case "":
                                            var v16 string = "\n\t"
                                            v17 := GenS2{F0: -32768, F1: v11, F2: 313}
                                            if true {
                                                v18 := GenS2{F0: int16(v6), F1: (v11 % (((v11 + -125)) | 1)), F2: -859}
                                                v19 := GenS3{F0: 42, F1: v7, F2: 32767}
                                                _ = v18
                                                _ = v19
                                            }
                                            _ = v16
                                            _ = v17
                                        default:
                                            v11 = 1
                                        }
                                    } else {
                                        _ = len(v15)
                                        switch "" {
                                        case "":
                                            v7 += (v7 % ((884) | 1))
                                        }
                                    }
                                }
                                var v20 string = "\n\t"
                                fmt.Println(v11, v20, v11)
                                _ = v15
                                _ = v20
                            }
                        case 2147483647:
                            v21 := GenS2{F0: -32768, F1: (^(v11)), F2: -9223372036854775808}
                            _ = v21
                        default:
                            v22 := []int{fi10, (fi10 ^ -9223372036854775808), 32768}
                            _ = v22
                        }
                    }
                    switch fi9 {
                    case 32767:
                        v23 := []int{}
                        _ = v23
                    }
                    _ = v11
                }
                v7 += v7
                _ = fi10
            }
            v7 += (^(v7))
            v24 := GenS1{F0: 1}
            _ = v24
        }
        _ = fi9
    }
    _ = v6
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 int32, p1 uint64) int {
    var v25 uint8 = 128
    v25 = (v25 / ((128) | 1))
    var v26 string = ("hello" + ("hello" + ("foo bar" + "foo bar")))
    p1 -= (p1 >> ((42) & 63))
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    return -461
}

func main() {
    var v1 bool = true
    v2 := make(map[string]uint32)
    var v3 uint = 2
    var v4 uint8 = 128
    var v5 uint16 = uint16(v4)
    var v27 string = "foo bar"
    var v28 float64 = 0.1
    switch -2147483648 {
    case -9223372036854775808:
        var v29 uint = (v3 / (((^(v3))) | 1))
        _ = v29
    case 2147483647:
        var v30 int64 = 1
        _ = v30
    default:
        v4 = v4
    }
    var v31 string = v27
    var v32 uint8 = 1
    v28 = v28
    fmt.Println(v28, v5, v1)
    for fi33 := 0; fi33 < 13; fi33++ {
        if (-1 == 32767) {
            v34 := []int{genFn1(1, uint64(v3)), fi33, (fi33 + (fi33 >> (((v3 & (v3 >> ((v3) & 63)))) & 63)))}
            v34 = append(v34, fi33)
            _ = v34
        } else {
            v27 = (v31 + "foo bar")
            switch genFn1(-1, 4294967295) {
            case 0:
                v31 = v27
            case 9223372036854775807:
                v35 := GenS1{F0: 127}
                _ = v35
            case 1:
                switch v31 {
                case "a":
                    var v36 string = v27
                    _ = v36
                default:
                    if false {
                        v37 := GenS2{F0: -1, F1: 0, F2: 9223372036854775807}
                        var v38 int16 = 42
                        if (v28 != ((3.14159 * v28) - 0.1)) {
                            var v39 float32 = (3.14159 / 0.1)
                            _ = v39
                        }
                        _ = v37
                        _ = v38
                    }
                }
            }
            if true {
                var v40 float32 = 3.14159
                _ = v40
            } else {
                v27 = v31
            }
        }
        v28 = float64(3.14159)
        for fi41 := 0; fi41 < 10; fi41++ {
            var v42 int64 = -512
            if false {
                delete(v2, ((v31 + "a") + v27))
                switch v31 {
                case "hello":
                    fmt.Println(v27, v5)
                case "foo bar":
                    if (v1 && (100 != -128)) {
                        if (v5 == (v5 >> (((v3 * (v3 | 24))) & 63))) {
                            v43 := []int{(fi33 % ((genFn1(-598, 4294967295)) | 1)), 127}
                            var v44 uint64 = 531
                            _ = v43
                            _ = v44
                        } else {
                            v5 += v5
                            v27 = "a"
                        }
                    }
                default:
                    v2[v31] = 2147483647
                }
                if v1 {
                    v45 := GenS4{F0: (-(v42)), F1: 2, F2: 140}
                    _ = v45
                }
            } else {
                fmt.Println(v4)
                v1 = ((fi33 * (fi41 + -126)) >= 9223372036854775807)
            }
            var v46 int16 = 32767
            if (v1 || false) {
                switch ("\n\t" + "") {
                case "hello":
                    var v47 int = fi41
                    _ = v47
                case "":
                    fmt.Println(v32)
                case "foo bar":
                    var v48 bool = (int64(v3) > -9223372036854775808)
                    _ = v48
                }
                v3 += v3
            } else {
                fmt.Println(v4)
            }
            var v49 uint = (v3 ^ 4294967295)
            var v50 string = "hello"
            _ = fi41
            _ = v42
            _ = v46
            _ = v49
            _ = v50
        }
        _ = fi33
    }
    var v51 int32 = 2147483647
    for fi52 := 0; fi52 < 9; fi52++ {
        fmt.Println(v31, v1)
        _ = fi52
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v51
}
