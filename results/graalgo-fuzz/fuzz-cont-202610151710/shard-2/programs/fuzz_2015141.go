package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 string }
type GenS2 struct { F0 uint8; F1 string; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }

func main() {
    var v1 int64 = 42
    var v2 uint16 = 65535
    var v3 bool = (1e10 > ((-2.25 * (3.14159 + 0.1)) + (1e10 + 0.1)))
    var v4 int32 = -299
    v5 := []uint32{2, 2147483647, 1, 2147483647}
    var v6 int8 = 1
    v7 := GenS1{F0: 703, F1: (v6 << ((42) & 63)), F2: "foo bar"}
    var v8 uint32 = 2147483648
    v9 := GenS1{F0: 359, F1: -34, F2: (("foo bar" + "") + (("" + "") + "hello"))}
    v2 += 433
    fmt.Println(v1, v6)
    for fi10 := 0; fi10 < 1; fi10++ {
        switch ((("a" + "hello") + ("foo bar" + ("" + "foo bar"))) + "hello") {
        case "foo bar":
            var v11 float64 = (3.14159 / 3.14159)
            _ = v11
        case "\n\t":
            if (-1 <= 32767) {
                for fi12 := 0; fi12 < 16; fi12++ {
                    fmt.Println(v8)
                    _ = fi12
                }
                v6 -= (v6 * (v6 & 127))
            }
        }
        if (-2147483648 != 2147483647) {
            if (-2.25 > ((-2.25 - (1.5 + 0.1)) + (0.0 + 0.1))) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if ((v8 ^ 42) < v8) {
                    switch "\n\t" {
                    case "\n\t":
                        var v13 int = fi10
                        _ = v13
                    case "foo bar":
                        v14 := GenS1{F0: 2, F1: (v6 / ((v6) | 1)), F2: ""}
                        _ = v14
                    }
                    if (fi10 != -81) {
                        fmt.Println(v3, v1, v8)
                        v15 := GenS2{F0: 127, F1: "", F2: v6}
                        _ = v15
                    }
                } else {
                    for fi16 := 0; fi16 < 19; fi16++ {
                        v17 := GenS2{F0: 171, F1: (((("\n\t" + ("a" + ("\n\t" + "foo bar"))) + "hello") + ("" + "")) + ("foo bar" + (((("\n\t" + "") + ("\n\t" + "\n\t")) + "") + "a"))), F2: v6}
                        var v18 int32 = -2147483648
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    v4 -= 166
                    v19 := GenS3{F0: (v8 << ((18446744073709551615) & 63))}
                    _ = v19
                }
            } else {
                if v3 {
                    v6 = v6
                } else {
                    v8 = v8
                }
                var v20 int16 = 1000
                _ = v20
            }
            for fi21 := 0; fi21 < 9; fi21++ {
                fmt.Println(v1, v8, v3)
                v5 = append(v5, 4294967295)
                v3 = (v3 || v3)
                _ = fi21
            }
            v3 = (2 == 18446744073709551615)
        } else {
            var v22 int8 = v6
            for fi23 := 0; fi23 < 3; fi23++ {
                if v3 {
                    var v24 int16 = -324
                    if len(v5) > 0 {
                        v5[0] = (v8 / ((22) | 1))
                    }
                    _ = v24
                }
                if (((^(v22)) != int8(v1)) || (!v3)) {
                    if true {
                        fmt.Println(v22, v6)
                    } else {
                        fmt.Println(v2, v2, v4)
                        v25 := GenS4{F0: 128}
                        _ = v25
                    }
                } else {
                    if len(v5) > 0 {
                        v5[0] = v8
                    }
                    v4 -= v4
                    var v26 int = (-(fi23))
                    _ = v26
                }
                v2 = v2
                _ = fi23
            }
            _ = v22
        }
        fmt.Println(fi10)
        _ = fi10
    }
    switch 4294967295 {
    case -1000:
        if v3 {
            var v27 int = -9223372036854775808
            for fi28 := 0; fi28 < 10; fi28++ {
                var v29 uint8 = 2
                v30 := GenS2{F0: 42, F1: ("\n\t" + "a"), F2: 42}
                var v31 string = (("a" + "foo bar") + "")
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            var v32 float32 = 1e10
            _ = v27
            _ = v32
        }
    case 127:
        var v33 float32 = 0.0
        _ = v33
    default:
        fmt.Println(v3)
    }
    v1 += (v1 & v1)
    for fi34 := 0; fi34 < 14; fi34++ {
        var v35 int16 = -32768
        v36 := GenS5{F0: v4}
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 uint8 = 127
    fmt.Println(v2, v2, v8)
    v37 -= (^(v37))
    for fi38 := 0; fi38 < 16; fi38++ {
        for fi39 := 0; fi39 < 11; fi39++ {
            v8 += (v8 - (^(v8)))
            fmt.Println(fi38)
            var v40 uint32 = (v8 * (v8 | (v8 * (v8 & 2147483647))))
            var v41 uint64 = 42
            _ = fi39
            _ = v40
            _ = v41
        }
        v37 += v37
        var v42 int8 = (v6 << ((361) & 63))
        _ = fi38
        _ = v42
    }
    switch "hello" {
    case "a":
        v2 += (^(v2))
    case "hello":
        if (v3 && true) {
            var v43 float32 = 1e10
            for fi44 := 0; fi44 < 10; fi44++ {
                v5 = append(v5, v8)
                for fi45 := 0; fi45 < 12; fi45++ {
                    var v46 int64 = v1
                    v4 = v4
                    switch fi44 {
                    case 0:
                        v1 += v1
                    }
                    _ = fi45
                    _ = v46
                }
                _ = fi44
            }
            fmt.Println(v3)
            _ = v43
        }
    default:
        var v47 uint16 = 0
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
