package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int64 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]int)
    var v3 uint = 0
    var v4 float64 = float64(1e10)
    var v5 uint16 = 65535
    if false {
        v4 += float64(3.14159)
        v4 -= v4
        if true {
            v3 = (v3 * v3)
            v6 := GenS1{F0: -1, F1: 18446744073709551615, F2: true}
            _ = v6
        }
    }
    fmt.Println(v5, v5, v5)
    if (224 >= int64(v3)) {
        var v7 float32 = 1e10
        var v8 bool = true
        if (false || v8) {
            var v9 int8 = 100
            if v8 {
                var v10 bool = v8
                _ = v10
            } else {
                for fi11 := 0; fi11 < 15; fi11++ {
                    v8 = v8
                    var v12 int16 = -480
                    for fi13 := 0; fi13 < 13; fi13++ {
                        switch ("" + "a") {
                        case "":
                            v1 = append(v1, "foo bar")
                            var v14 uint = 319
                            v15 := GenS1{F0: fi11, F1: 4294967295, F2: (42 > -14)}
                            _ = v14
                            _ = v15
                        }
                        _ = fi13
                    }
                    _ = fi11
                    _ = v12
                }
                v8 = v8
            }
            _ = v9
        } else {
            var v16 bool = v8
            v7 -= float32(1.5)
            switch "hello" {
            case "hello":
                v3 += (v3 * v3)
            case "a":
                var v17 int16 = -82
                _ = v17
            default:
                v18 := GenS1{F0: -9223372036854775808, F1: 42, F2: ((v3 / ((v3) | 1)) == 255)}
                _ = v18
            }
            _ = v16
        }
        _ = v7
        _ = v8
    } else {
        v19 := GenS2{F0: -2147483648}
        _ = v19
    }
    if false {
        var v20 string = "\n\t"
        fmt.Println(v5)
        if (-2.25 < 1.5) {
            fmt.Println(v4, v20)
            var v21 uint16 = 1
            _ = v21
        }
        _ = v20
    } else {
        v4 -= v4
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        v3 = (v3 / (((^(v3))) | 1))
        v4 = 3.14159
        for fi23 := 0; fi23 < 19; fi23++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if false {
                v24 := GenS1{F0: (fi23 * (fi22 & fi23)), F1: 795, F2: (!((true && false) && (-21 != 436)))}
                v3 -= v3
                if true {
                    v2[0] = -916
                } else {
                    if (((("foo bar" + "\n\t") + "") + "\n\t") == "foo bar") {
                        v5 = uint16(v3)
                    }
                    var v25 int = 32768
                    v3 -= (^(v3))
                    v3 = 547
                    var v26 uint = (v3 / (((v3 & v3)) | 1))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v27 := GenS3{F0: 0, F1: 0}
                    if false {
                        var v28 uint64 = 272
                        var v29 int64 = -947
                        v3 = 57
                        _ = v28
                        _ = v29
                    }
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = v24
            }
            _ = fi23
        }
        _ = fi22
    }
    v30 := GenS1{F0: 2147483647, F1: 4294967295, F2: true}
    v31 := GenS3{F0: 1000, F1: int64(v3)}
    for fi32 := 0; fi32 < 14; fi32++ {
        for fi33 := 0; fi33 < 18; fi33++ {
            delete(v2, 374)
            _ = fi33
        }
        v3 -= (v3 - (v3 + 65535))
        _ = fi32
    }
    v3 = v3
    v3 = 4294967295
    v4 = float64(1.5)
    v5 = 2
    _ = len(v1)
    switch 32768 {
    case -2147483648:
        if (!false) {
            for fi34 := 0; fi34 < 12; fi34++ {
                v5 = (^(v5))
                var v35 int64 = 42
                _ = fi34
                _ = v35
            }
        }
    default:
        for fi36 := 0; fi36 < 17; fi36++ {
            var v37 uint32 = 148
            _ = fi36
            _ = v37
        }
    }
    if (!((-897 < 42) && true)) {
        for fi38 := 0; fi38 < 12; fi38++ {
            var v39 float32 = 1e10
            _ = fi38
            _ = v39
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
