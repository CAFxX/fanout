package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int32 }
func genFn0(p0 int8, p1 uint64) int64 {
    v7 := GenS1{F0: 2147483647, F1: ("" + ("" + "hello")), F2: 0}
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS2{F0: (p1 % ((652) | 1))}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    v1 := []string{("foo bar" + "\n\t")}
    var v2 string = ((("" + "foo bar") + "\n\t") + "a")
    v3 := []uint32{318, 2147483647, 0}
    v4 := []uint32{4294967295}
    var v5 string = v2
    v6 := make(map[string]bool)
    v10 := GenS3{F0: 2, F1: 1, F2: -1}
    v5 = v5
    switch 42 {
    case 32768:
        var v11 int64 = genFn0(-117, 18446744073709551615)
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        _ = len(v6)
        v13 := GenS3{F0: 0, F1: 0, F2: -32768}
        _ = fi12
        _ = v13
    }
    v14 := GenS4{F0: 2147483647, F1: 1}
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint = 2
        v5 = v5
        _ = fi15
        _ = v16
    }
    if true {
        v17 := GenS5{F0: 2, F1: -1}
        if (42 >= -125) {
            var v18 int64 = genFn0(31, 393)
            fmt.Println(v5, v2, v2)
            _ = v18
        } else {
            var v19 uint16 = 1
            _ = v19
        }
        var v20 int64 = -9223372036854775808
        _ = v17
        _ = v20
    } else {
        if (true || (0 != 1)) {
            v5 = (("hello" + ("foo bar" + (v5 + v2))) + v2)
            for fi21 := 0; fi21 < 2; fi21++ {
                v2 = "hello"
                var v22 string = "hello"
                if false {
                    v23 := GenS6{F0: -2147483648}
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 17; fi24++ {
                        fmt.Println(fi24, v2, fi24)
                        v25 := GenS2{F0: 18446744073709551615}
                        var v26 float64 = -2.25
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                }
                _ = fi21
                _ = v22
            }
            v5 = (v5 + "")
        } else {
            var v27 uint32 = 2147483647
            _ = v27
        }
        var v28 int = 777
        _ = v28
    }
    v29 := GenS5{F0: 554, F1: -9223372036854775808}
    var v30 int = 1
    if (!false) {
        fmt.Println(v2)
        switch v30 {
        case -52:
            if ((true || (true || false)) || true) {
                var v31 float64 = (float64(((float32(1.5) / (1e10 * 1.5)) - ((float32(-2.25) + 0.0) / -2.25))) / 1.0)
                _ = v31
            } else {
                fmt.Println(v2, v2, v5)
                for fi32 := 0; fi32 < 9; fi32++ {
                    v6[v2] = (!((true || true) || (0 <= 42)))
                    _ = fi32
                }
                if true {
                    fmt.Println(v2)
                }
            }
        case -761:
            switch 32768 {
            case 0:
                _ = len(v6)
            case 32767:
                v33 := GenS3{F0: 120, F1: 2, F2: -32768}
                _ = v33
            case 127:
                v2 = ""
            default:
                for fi34 := 0; fi34 < 10; fi34++ {
                    var v35 uint = 255
                    _ = fi34
                    _ = v35
                }
            }
        default:
            switch ("" + v5) {
            case "foo bar":
                v6[(v5 + (v5 + ("hello" + "foo bar")))] = true
            case "hello":
                if (v30 != -9223372036854775808) {
                    var v36 uint64 = 787
                    _ = v36
                }
            case "a":
                v5 = v5
            }
        }
        switch "" {
        case "\n\t":
            _ = len(v3)
        case "foo bar":
            var v37 string = "\n\t"
            _ = v37
        default:
            v38 := GenS1{F0: 1, F1: (("a" + v5) + v5), F2: 42}
            _ = v38
        }
    }
    fmt.Println(v5)
    for fi39 := 0; fi39 < 18; fi39++ {
        fmt.Println(v2)
        fmt.Println(v2, v30)
        _ = fi39
    }
    var v40 uint32 = 2147483648
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
