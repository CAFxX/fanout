package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int8 = -29
    var v2 int8 = (v1 * (v1 & v1))
    var v3 int = -2147483648
    var v4 int16 = 42
    v5 := make(map[string]string)
    v6 := []string{"\n\t", "foo bar", ("a" + (("a" + "foo bar") + "foo bar")), "\n\t"}
    v2 -= v2
    if false {
        if (-1 == -1) {
            switch "foo bar" {
            case "a":
                v7 := GenS1{F0: v2, F1: v2, F2: 107}
                _ = v7
            case "\n\t":
                v8 := v5["foo bar"]
                _ = v8
            default:
                v9 := GenS1{F0: (v1 / ((-128) | 1)), F1: v2, F2: 255}
                _ = v9
            }
        } else {
            var v10 int8 = 100
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 uint64 = 18446744073709551615
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        if (!(!(false && ((2 >= 18446744073709551615) && true)))) {
            fmt.Println(v3)
            v4 = -1
        } else {
            switch "\n\t" {
            case "foo bar":
                v1 = v1
            case "":
                for fi13 := 0; fi13 < 10; fi13++ {
                    fmt.Println(v4, fi13)
                    if (2 < 253) {
                        var v14 int64 = 362
                        _ = v14
                    }
                    _ = fi13
                }
            default:
                fmt.Println(v2, v1, v4)
            }
        }
        var v15 string = "\n\t"
        _ = v15
    } else {
        v4 -= v4
        var v16 uint = 0
        var v17 int64 = -1
        _ = v16
        _ = v17
    }
    fmt.Println(v4, v2, v3)
    if (!true) {
        if (false && true) {
            if ((v1 ^ (v1 % ((-56) | 1))) >= v1) {
                var v18 int32 = 0
                switch "foo bar" {
                case "\n\t":
                    v19 := GenS2{F0: 1}
                    _ = v19
                case "foo bar":
                    switch (v3 & (v3 * 9223372036854775807)) {
                    case -429:
                        var v20 string = "a"
                        _ = v20
                    case 32768:
                        if (false && true) {
                            var v21 int64 = 42
                            var v22 uint32 = 2147483648
                            _ = v21
                            _ = v22
                        } else {
                            v1 = (v1 - v2)
                            v23 := GenS1{F0: v1, F1: (v1 / (((-(v2))) | 1)), F2: 4294967295}
                            _ = v23
                        }
                    default:
                        var v24 bool = (!true)
                        _ = v24
                    }
                default:
                    fmt.Println(v1, v4)
                }
                v25 := GenS3{F0: 4294967295, F1: 32}
                _ = v18
                _ = v25
            }
            if false {
                fmt.Println(v2, v1, v1)
            } else {
                var v26 uint64 = 18446744073709551615
                v2 = v1
                v4 = v4
                _ = v26
            }
            v27 := v5["foo bar"]
            _ = v27
        }
        if true {
            var v28 uint = 255
            if false {
                var v29 int32 = -902
                _ = v29
            }
            _ = v28
        }
        if true {
            v3 = 32767
            v30 := GenS3{F0: 1, F1: 1}
            _ = v30
        } else {
            var v31 uint16 = 2
            switch 9223372036854775807 {
            case 0:
                var v32 string = ""
                _ = v32
            }
            _ = v31
        }
    } else {
        v33 := GenS4{F0: 42, F1: (v1 | (^(v1)))}
        _ = v33
    }
    v34 := GenS1{F0: (v1 >> ((18446744073709551615) & 63)), F1: -128, F2: 427}
    var v35 float32 = 0.0
    if false {
        var v36 int64 = -1
        var v37 int16 = (v4 << ((96) & 63))
        _ = v36
        _ = v37
    } else {
        var v38 int32 = 42
        for fi39 := 0; fi39 < 6; fi39++ {
            v3 -= (v3 ^ (v3 + v3))
            for fi40 := 0; fi40 < 11; fi40++ {
                fmt.Println(v3, v4, v35)
                var v41 uint8 = 255
                switch (^(fi40)) {
                case -489:
                    fmt.Println(fi40, v1)
                }
                _ = fi40
                _ = v41
            }
            _ = fi39
        }
        _ = v38
    }
    for fi42 := 0; fi42 < 14; fi42++ {
        var v43 int32 = 42
        _ = fi42
        _ = v43
    }
    var v44 uint64 = 636
    v45 := v5["\n\t"]
    v1 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
