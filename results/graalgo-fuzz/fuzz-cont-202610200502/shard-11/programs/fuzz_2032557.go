package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint32 }

func main() {
    var v1 float64 = (-2.25 / 1.0)
    var v2 bool = false
    var v3 int8 = 127
    v4 := make(map[uint32]int)
    v5 := make(map[int]uint32)
    var v6 string = (("" + "\n\t") + ("a" + ("a" + "")))
    var v7 int32 = 0
    v5[-9223372036854775808] = 0
    switch ((("" + v6) + "foo bar") + (v6 + v6)) {
    case "hello":
        var v8 uint64 = 18446744073709551615
        _ = v8
    case "a":
        _ = len(v4)
    case "foo bar":
        for fi9 := 0; fi9 < 15; fi9++ {
            switch "hello" {
            case "a":
                v10 := v4[42]
                _ = v10
            }
            _ = fi9
        }
    default:
        if ((v1 - 1e10) == v1) {
            switch 127 {
            case 32768:
                fmt.Println(v7, v2, v3)
            case 0:
                for fi11 := 0; fi11 < 13; fi11++ {
                    v1 = v1
                    _ = fi11
                }
            }
            switch 4294967295 {
            case 4294967295:
                var v12 bool = (v3 == (v3 % (((v3 | v3)) | 1)))
                _ = v12
            case 949:
                v13 := []int{127, 4294967295, 42, 32768}
                _ = v13
            case 894:
                v7 = v7
            default:
                switch -128 {
                case -801:
                    v6 = "\n\t"
                default:
                    _ = len(v4)
                }
            }
        } else {
            fmt.Println(v3)
        }
    }
    var v14 uint8 = 42
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 uint8 = v14
        var v17 int = (fi15 * fi15)
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 0}
    if (520 == 53) {
        var v19 uint = 564
        if false {
            var v20 int32 = (v7 - v7)
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 float64 = v1
                v20 -= v20
                var v23 uint16 = 0
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v20
        } else {
            for fi24 := 0; fi24 < 20; fi24++ {
                v6 = ((v6 + v6) + v6)
                var v25 bool = ((v6 >= v6) || true)
                var v26 string = "hello"
                _ = fi24
                _ = v25
                _ = v26
            }
            v1 = v1
        }
        _ = v19
    }
    var v27 int16 = 0
    v3 = v3
    if (!v2) {
        for fi28 := 0; fi28 < 2; fi28++ {
            var v29 uint64 = 0
            _ = fi28
            _ = v29
        }
    } else {
        if v2 {
            for fi30 := 0; fi30 < 6; fi30++ {
                switch fi30 {
                case 0:
                    var v31 int32 = v7
                    _ = v31
                }
                _ = fi30
            }
            fmt.Println(v3)
            var v32 int32 = v7
            _ = v32
        }
        for fi33 := 0; fi33 < 13; fi33++ {
            for fi34 := 0; fi34 < 2; fi34++ {
                var v35 int16 = v27
                var v36 float64 = v1
                v37 := GenS2{F0: (v3 << ((2) & 63))}
                v5[fi34] = 476
                switch 670 {
                case 187:
                    if ((((^(v7)) < (v7 / (((v7 / (((^(v7))) | 1))) | 1))) || ((4294967295 < 2) && true)) && v2) {
                        if false {
                            var v38 int32 = v7
                            switch (fi33 ^ 2147483647) {
                            case 2147483647:
                                v39 := v5[fi34]
                                _ = v39
                            }
                            _ = v38
                        } else {
                            fmt.Println(v2, v2, fi34)
                        }
                        fmt.Println(v36, v27, fi33)
                        v36 -= v36
                    } else {
                        v2 = ((!(v2 && (true || false))) || v2)
                    }
                default:
                    v40 := GenS3{F0: ((v7 * 0) == v7), F1: 25}
                    _ = v40
                }
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
            var v41 uint8 = 128
            fmt.Println(fi33, v7)
            _ = fi33
            _ = v41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
