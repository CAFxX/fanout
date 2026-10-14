package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int; F1 uint32; F2 int }
type GenS5 struct { F0 int8 }
func genFn0() bool {
    for fi5 := 0; fi5 < 9; fi5++ {
        if true {
            var v6 int16 = -32768
            v7 := GenS1{F0: -32768, F1: 172}
            switch "hello" {
            case "hello":
                var v8 uint64 = 18446744073709551615
                _ = v8
            }
            _ = v6
            _ = v7
        }
        _ = fi5
    }
    if (!true) {
        if true {
            if true {
                for fi9 := 0; fi9 < 17; fi9++ {
                    var v10 int = (fi9 * (fi9 >> ((0) & 63)))
                    _ = fi9
                    _ = v10
                }
            }
            v11 := GenS1{F0: 365, F1: 32767}
            _ = v11
        }
        v12 := GenS2{F0: "", F1: 469}
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 string = ""
        _ = fi13
        _ = v14
    }
    return ((!(true || (!true))) || false)
}

func main() {
    var v1 int64 = -370
    v2 := []string{"hello", "foo bar"}
    var v3 string = "\n\t"
    var v4 bool = true
    for fi15 := 0; fi15 < 5; fi15++ {
        _ = len(v2)
        v3 = v3
        v3 = v3
        _ = fi15
    }
    if (((-(v1)) > (^(v1))) && (1 != 1)) {
        var v16 int64 = 1
        if true {
            if false {
                if v4 {
                    var v17 bool = (-1 <= 0)
                    var v18 uint = 18446744073709551615
                    _ = v17
                    _ = v18
                }
                for fi19 := 0; fi19 < 6; fi19++ {
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v16 = -9223372036854775808
                        fmt.Println(fi19, v3, v3)
                        var v21 int32 = -286
                        _ = fi20
                        _ = v21
                    }
                    fmt.Println(v3, v3)
                    fmt.Println(v1)
                    _ = fi19
                }
            }
            var v22 uint = 18446744073709551615
            _ = v22
        }
        _ = len(v2)
        _ = v16
    }
    var v23 int = 2147483647
    fmt.Println(v4, v1)
    for fi24 := 0; fi24 < 4; fi24++ {
        if v4 {
            var v25 uint = 2
            var v26 uint16 = 404
            _ = v25
            _ = v26
        }
        _ = fi24
    }
    if len(v2) > 0 {
        v2[0] = v3
    }
    v27 := GenS1{F0: 1000, F1: 905}
    switch v23 {
    case 9223372036854775807:
        v28 := GenS3{F0: 42, F1: -494}
        _ = v28
    default:
        for fi29 := 0; fi29 < 7; fi29++ {
            fmt.Println(v23, v3)
            _ = fi29
        }
    }
    switch (v23 - v23) {
    case 2147483647:
        if v4 {
            v4 = v4
        }
    case 9223372036854775807:
        for fi30 := 0; fi30 < 16; fi30++ {
            switch v3 {
            case "foo bar":
                switch (v23 << ((242) & 63)) {
                case -1:
                    v31 := GenS3{F0: 994, F1: 42}
                    _ = v31
                }
            case "":
                switch (v23 & (v23 >> ((35) & 63))) {
                case -2147483648:
                    if len(v2) > 0 {
                        v2[0] = v3
                    }
                case 42:
                    if (-2.25 >= 0.1) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        v32 := GenS1{F0: 1, F1: 742}
                        v23 -= (fi30 + 4294967295)
                        _ = v32
                    }
                default:
                    fmt.Println(v4, v1)
                }
            case "\n\t":
                fmt.Println(v4, fi30)
            default:
                var v33 float32 = 0.0
                _ = v33
            }
            for fi34 := 0; fi34 < 12; fi34++ {
                v23 -= (fi34 | 0)
                v3 = v3
                var v35 uint = 18446744073709551615
                v3 = "hello"
                var v36 float32 = (3.14159 / (3.14159 - 1.5))
                var v37 uint16 = 805
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
            _ = fi30
        }
    case 32768:
        if len(v2) > 0 {
            v2[0] = v3
        }
    }
    switch "" {
    case "hello":
        v38 := GenS4{F0: -499, F1: 42, F2: 4294967295}
        _ = v38
    case "foo bar":
        for fi39 := 0; fi39 < 6; fi39++ {
            v1 += v1
            _ = fi39
        }
    case "a":
        for fi40 := 0; fi40 < 15; fi40++ {
            var v41 uint16 = 42
            _ = fi40
            _ = v41
        }
    }
    v42 := GenS4{F0: (v23 / ((v23) | 1)), F1: 1, F2: (v23 | v23)}
    v2 = append(v2, v3)
    v43 := GenS5{F0: 26}
    v4 = (false || ((-1 != 862) || (v4 || false)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}
