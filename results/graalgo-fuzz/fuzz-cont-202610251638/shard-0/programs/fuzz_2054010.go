package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int8; F2 uint16 }
type GenS3 struct { F0 int16; F1 int32; F2 bool }
func genFn0(p0 uint32, p1 int32) uint64 {
    fmt.Println(p0, p1)
    p0 += 4294967295
    p1 = p1
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 string = ("\n\t" + "a")
    var v2 bool = (true || (646 != 42))
    v3 := make(map[string]int)
    v4 := []int{1, 9223372036854775807, 1}
    var v5 string = v1
    switch v5 {
    case "foo bar":
        switch 32767 {
        case 42:
            switch 4294967295 {
            case 32767:
                switch 128 {
                case 1:
                    var v6 int16 = 1000
                    _ = v6
                default:
                    var v7 int32 = 555
                    _ = v7
                }
            case 4294967295:
                v5 = v5
            case 127:
                v2 = ((1 <= -1) || v2)
            default:
                var v8 uint = 4294967295
                _ = v8
            }
        case 2147483647:
            for fi9 := 0; fi9 < 18; fi9++ {
                v1 = "hello"
                switch v1 {
                case "foo bar":
                    if v2 {
                        v3[v1] = (-(fi9))
                        var v10 uint16 = 65535
                        var v11 int = fi9
                        _ = v10
                        _ = v11
                    } else {
                        var v12 uint64 = 1
                        if v2 {
                            v13 := GenS1{F0: 42}
                            v4 = append(v4, -2147483648)
                            var v14 int = (fi9 & (fi9 << ((609) & 63)))
                            _ = v13
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 13; fi15++ {
                            var v16 float32 = ((3.14159 / (-2.25 / (0.1 + (-2.25 / (((((3.14159 * 1.5) + (float32(1.5) + (-2.25 - float32(1.5)))) + float32(0.1)) - ((1.5 + (-2.25 / 1.5)) / 1e10)) - float32(float64(1e10))))))) - 3.14159)
                            _ = fi15
                            _ = v16
                        }
                        _ = v12
                    }
                case "":
                    if v2 {
                        if (!v2) {
                            fmt.Println(v1, v5)
                            v2 = true
                            for fi17 := 0; fi17 < 16; fi17++ {
                                v18 := v3[v1]
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                switch fi17 {
                                case 9223372036854775807:
                                    var v19 int64 = -1
                                    _ = v19
                                case 42:
                                    if (18446744073709551615 >= 601) {
                                        switch (-(v18)) {
                                        case 32768:
                                            v20 := GenS1{F0: -1}
                                            _ = v20
                                        }
                                        switch "" {
                                        case "hello":
                                            v5 = v1
                                        case "\n\t":
                                            v18 = fi9
                                        }
                                        v4 = append(v4, (fi9 % ((fi9) | 1)))
                                    }
                                default:
                                    switch -109 {
                                    case -423:
                                        v18 = (fi17 * -205)
                                    case -446:
                                        if v2 {
                                            v1 = (v5 + "foo bar")
                                            v3[v5] = fi9
                                            if false {
                                                if false {
                                                    fmt.Println(v1, fi9, v1)
                                                    _ = len(v3)
                                                    var v21 bool = v2
                                                    _ = v21
                                                }
                                                var v22 uint16 = 255
                                                _ = v22
                                            } else {
                                                if ((0.0 / 3.14159) != (3.14159 / 0.1)) {
                                                    v23 := GenS2{F0: v18, F1: -115, F2: 734}
                                                    _ = v23
                                                } else {
                                                    fmt.Println(fi17, fi17)
                                                    v24 := GenS2{F0: 0, F1: 100, F2: 623}
                                                    _ = v24
                                                }
                                                if false {
                                                    fmt.Println(v5)
                                                    var v25 uint8 = 1
                                                    v26 := v3[("" + v5)]
                                                    _ = v25
                                                    _ = v26
                                                } else {
                                                    v2 = v2
                                                    var v27 float64 = -2.25
                                                    _ = v27
                                                }
                                            }
                                        } else {
                                            v4 = append(v4, (v18 / (((v18 % ((fi17) | 1))) | 1)))
                                            var v28 uint = 42
                                            _ = v28
                                        }
                                    }
                                }
                                _ = fi17
                                _ = v18
                            }
                        }
                        if v2 {
                            v1 = v5
                            var v29 uint8 = 42
                            var v30 uint8 = (^(v29))
                            _ = v29
                            _ = v30
                        }
                    } else {
                        v1 = "foo bar"
                        var v31 uint16 = 955
                        v32 := GenS3{F0: 1000, F1: -2147483648, F2: (v2 && (v2 && (-2.25 <= 0.0)))}
                        _ = v31
                        _ = v32
                    }
                }
                _ = fi9
            }
        }
    default:
        for fi33 := 0; fi33 < 5; fi33++ {
            v2 = (v2 || (1.5 > -2.25))
            v1 = v1
            _ = fi33
        }
    }
    var v34 uint = 0
    v5 = v5
    v2 = (!v2)
    v35 := GenS3{F0: -1, F1: 2147483647, F2: false}
    var v36 uint = (v34 % ((65535) | 1))
    var v37 int16 = -32768
    var v38 uint64 = 645
    var v39 uint64 = 4294967295
    v4 = append(v4, 42)
    if v2 {
        for fi40 := 0; fi40 < 9; fi40++ {
            for fi41 := 0; fi41 < 17; fi41++ {
                var v42 float64 = 0.0
                var v43 uint16 = 432
                _ = fi41
                _ = v42
                _ = v43
            }
            fmt.Println(v34, v38, v34)
            _ = fi40
        }
        v36 = (^(v34))
    } else {
        delete(v3, v1)
        if len(v4) > 0 {
            v4[0] = 808
        }
        switch -82 {
        case -1:
            switch 351 {
            case 32768:
                fmt.Println(v39, v1)
            case 42:
                v37 = 1
            }
        case -9223372036854775808:
            if len(v4) > 0 {
                v4[0] = 127
            }
        }
    }
    for fi44 := 0; fi44 < 13; fi44++ {
        v5 = v1
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
