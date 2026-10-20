package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }

func main() {
    var v1 int8 = 0
    v2 := []uint32{0, 816, 65535, 226}
    v3 := []string{"foo bar", "\n\t", ("\n\t" + "\n\t"), ("\n\t" + (("" + (("foo bar" + (("a" + ("" + "hello")) + "a")) + ("hello" + "a"))) + "\n\t"))}
    fmt.Println(v1)
    switch -1 {
    case 127:
        var v4 int = 128
        _ = v4
    default:
        v1 -= (-(v1))
    }
    if len(v2) > 0 {
        v2[0] = 808
    }
    v1 -= v1
    var v5 float64 = float64(0.1)
    v1 += (^(v1))
    v5 += 0.1
    switch 9223372036854775807 {
    case -568:
        var v6 uint8 = 233
        _ = v6
    case 32767:
        v5 = v5
    }
    v5 = v5
    var v7 uint16 = 0
    var v8 uint = 503
    fmt.Println(v1)
    for fi9 := 0; fi9 < 1; fi9++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v10 uint = (v8 | 0)
        _ = fi9
        _ = v10
    }
    switch 300 {
    case 9223372036854775807:
        if (-419 == -9223372036854775808) {
            if true {
                v11 := GenS1{F0: 9223372036854775807}
                _ = v11
            } else {
                var v12 bool = ((212 <= 0) || (false && (4294967295 != 0)))
                fmt.Println(v5)
                _ = v12
            }
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS1{F0: (fi13 & int(v8))}
                _ = len(v3)
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 19; fi15++ {
                v5 = ((v5 / (v5 / v5)) / 0.1)
                _ = fi15
            }
        } else {
            fmt.Println(v8, v5, v5)
            _ = len(v2)
        }
    case 4294967295:
        if true {
            var v16 float32 = 1.5
            v3 = append(v3, "\n\t")
            if false {
                for fi17 := 0; fi17 < 4; fi17++ {
                    if ((!false) && ((v8 - (v8 + v8)) < (v8 / (((^(v8))) | 1)))) {
                        switch "" {
                        case "foo bar":
                            _ = len(v3)
                        case "a":
                            if (128 != 2) {
                                var v18 float32 = 0.1
                                _ = len(v2)
                                _ = v18
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                for fi19 := 0; fi19 < 7; fi19++ {
                                    var v20 int64 = -1
                                    var v21 int32 = 2147483647
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                            }
                        default:
                            var v22 string = "foo bar"
                            _ = v22
                        }
                    } else {
                        v1 = (v1 ^ (v1 / (((v1 | v1)) | 1)))
                    }
                    for fi23 := 0; fi23 < 12; fi23++ {
                        v3 = append(v3, ("hello" + "hello"))
                        var v24 float32 = v16
                        switch fi23 {
                        case -956:
                            v24 = v24
                        default:
                            var v25 uint8 = 42
                            _ = v25
                        }
                        if (42 >= 2147483647) {
                            var v26 float64 = (-2.25 / v5)
                            _ = v26
                        }
                        var v27 float32 = (float32(float64(v24)) * ((float32(v5) * 1e10) - 0.0))
                        var v28 int64 = 1
                        _ = fi23
                        _ = v24
                        _ = v27
                        _ = v28
                    }
                    v1 = v1
                    _ = fi17
                }
            }
            _ = v16
        } else {
            var v29 uint8 = 42
            fmt.Println(v1)
            _ = v29
        }
    case 128:
        v30 := GenS2{F0: ("hello" + "a")}
        _ = v30
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        for fi32 := 0; fi32 < 14; fi32++ {
            v1 = int8(v8)
            if false {
                var v33 uint16 = v7
                _ = len(v2)
                var v34 int32 = -185
                v1 -= (^(v1))
                _ = v33
                _ = v34
            }
            _ = fi32
        }
        var v35 int16 = -1
        v36 := GenS1{F0: (fi31 / ((128) | 1))}
        _ = fi31
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v7
    _ = v8
}
