package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int32; F2 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 445
    v3 := make(map[uint32]bool)
    v4 := []string{}
    v1["foo bar"] = v2
    fmt.Println(v2)
    var v5 int32 = 553
    v6 := GenS1{F0: (("a" + "hello") + "foo bar")}
    if true {
        var v7 uint16 = 767
        var v8 bool = ((((3.14159 - 0.1) + -2.25) + 0.0) < 1.5)
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: ("" + "")}
        if true {
            fmt.Println(v2, v5)
            v10 := v3[v2]
            fmt.Println(v5, v10, v5)
            _ = v10
        } else {
            if false {
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 uint = 18446744073709551615
                    v12 -= v12
                    var v13 int32 = (^(v5))
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                for fi14 := 0; fi14 < 18; fi14++ {
                    delete(v3, 569)
                    _ = fi14
                }
            } else {
                for fi15 := 0; fi15 < 5; fi15++ {
                    fmt.Println(fi15)
                    v16 := GenS2{F0: int32(v2), F1: "\n\t"}
                    v5 += v5
                    _ = fi15
                    _ = v16
                }
                fmt.Println(v5, v5)
            }
            v17 := GenS3{F0: v2}
            for fi18 := 0; fi18 < 18; fi18++ {
                if true {
                    v2 = (v2 % (((v2 - 4294967295)) | 1))
                }
                for fi19 := 0; fi19 < 2; fi19++ {
                    var v20 uint64 = 443
                    if ((v5 + (v5 & (v5 % ((-1) | 1)))) > (v5 / ((v5) | 1))) {
                        delete(v3, 4294967295)
                        switch fi18 {
                        case -1:
                            delete(v3, uint32(v20))
                        case -9223372036854775808:
                            fmt.Println(fi19, v5, fi18)
                        }
                    } else {
                        v2 = (v2 - (v2 * v2))
                        v21 := v1["\n\t"]
                        if ((((3.14159 - 0.1) * (0.0 * 1.5)) / (-2.25 + (1e10 * (((((float32((3.14159 + 3.14159)) + 1e10) / 1.5) + (3.14159 / (-2.25 + -2.25))) + ((-2.25 * 0.0) + 1.5)) + 0.1)))) != 0.1) {
                            var v22 float64 = 0.0
                            v20 = (^(v20))
                            _ = v22
                        } else {
                            v21 -= 523
                        }
                        var v23 uint64 = 1
                        fmt.Println(v20, v5)
                        v23 = (v20 | (^(v20)))
                        _ = v21
                        _ = v23
                    }
                    v24 := GenS3{F0: v2}
                    _ = fi19
                    _ = v20
                    _ = v24
                }
                v5 -= v5
                _ = fi18
            }
            _ = v17
        }
        for fi25 := 0; fi25 < 7; fi25++ {
            if ((0.1 - 3.14159) >= 1e10) {
                v26 := GenS4{F0: uint16(v2), F1: -85, F2: 1000}
                _ = v26
            }
            _ = fi25
        }
        _ = v9
    }
    v27 := v3[(v2 / (((^(v2))) | 1))]
    var v28 float64 = 0.0
    if (v27 && v27) {
        v28 = (v28 * 1.5)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi29 := 0; fi29 < 7; fi29++ {
            v30 := GenS3{F0: v2}
            var v31 int64 = -9223372036854775808
            var v32 int64 = (v31 >> ((0) & 63))
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
    } else {
        var v33 uint8 = 128
        v27 = (4294967295 != 32768)
        for fi34 := 0; fi34 < 7; fi34++ {
            for fi35 := 0; fi35 < 8; fi35++ {
                switch (fi34 + (fi34 - (fi35 * (^(fi34))))) {
                case 2147483647:
                    v5 -= (v5 | v5)
                default:
                    v3[566] = (v28 >= v28)
                }
                fmt.Println(v27)
                fmt.Println(v2)
                _ = fi35
            }
            _ = fi34
        }
        _ = v33
    }
    v28 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
