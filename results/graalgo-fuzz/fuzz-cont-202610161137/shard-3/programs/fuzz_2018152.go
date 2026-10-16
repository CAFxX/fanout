package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    var v1 int32 = 1
    var v2 uint32 = 2
    v3 := make(map[uint32]int)
    v4 := []string{("\n\t" + "a"), "hello", "foo bar"}
    var v5 int = 434
    v4 = append(v4, "foo bar")
    if (127 <= -1) {
        switch v5 {
        case 6:
            var v6 bool = true
            _ = v6
        case 42:
            if false {
                for fi7 := 0; fi7 < 5; fi7++ {
                    if true {
                        v8 := GenS1{F0: 4294967295, F1: (v1 >> ((0) & 63)), F2: 18446744073709551615}
                        _ = v8
                    } else {
                        if (4294967295 <= 4294967295) {
                            switch (fi7 + 935) {
                            case 32768:
                                var v9 string = "foo bar"
                                _ = v9
                            case 9223372036854775807:
                                var v10 int64 = -9223372036854775808
                                _ = v10
                            default:
                                for fi11 := 0; fi11 < 2; fi11++ {
                                    delete(v3, (v2 / (((v2 % (((v2 & (v2 >> ((918) & 63)))) | 1))) | 1)))
                                    v5 -= v5
                                    _ = fi11
                                }
                            }
                            v12 := GenS1{F0: 2, F1: -990, F2: 0}
                            v2 = (v2 | v2)
                            _ = v12
                        } else {
                            v2 = v2
                        }
                        for fi13 := 0; fi13 < 20; fi13++ {
                            v14 := GenS2{F0: 33, F1: -457}
                            _ = fi13
                            _ = v14
                        }
                    }
                    fmt.Println(v2, v1, v2)
                    _ = fi7
                }
                var v15 int32 = (^(v1))
                var v16 float64 = (float64(((3.14159 * 3.14159) + 1e10)) / 1.0)
                _ = v15
                _ = v16
            } else {
                if true {
                    var v17 int32 = (v1 % ((int32(v5)) | 1))
                    var v18 int32 = v1
                    _ = v17
                    _ = v18
                } else {
                    v2 = 65535
                    for fi19 := 0; fi19 < 20; fi19++ {
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                        _ = fi19
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            }
        }
    }
    var v20 float64 = -2.25
    var v21 int32 = v1
    if len(v4) > 0 {
        v4[0] = ""
    }
    if true {
        var v22 uint64 = 4294967295
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        v24 := GenS3{F0: (-(v1)), F1: 745}
        switch "hello" {
        case "\n\t":
            v25 := GenS1{F0: 0, F1: (v1 + v1), F2: 2}
            _ = v25
        case "foo bar":
            v26 := GenS1{F0: 2, F1: v21, F2: 479}
            _ = v26
        }
        delete(v3, 0)
        _ = fi23
        _ = v24
    }
    fmt.Println(v21)
    v27 := GenS3{F0: (v1 | v21), F1: 1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v28 int = (v5 >> ((0) & 63))
    v3[(v2 ^ (v2 * (v2 >> ((968) & 63))))] = v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v21 -= -757
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v29 string = "foo bar"
    if (!(-42 >= -67)) {
        v30 := GenS3{F0: (-(v1)), F1: -664}
        var v31 int = 42
        var v32 float64 = ((0.0 / v20) - v20)
        _ = v30
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
