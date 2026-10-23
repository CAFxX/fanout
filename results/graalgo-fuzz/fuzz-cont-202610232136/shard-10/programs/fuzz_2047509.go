package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int8; F2 int }
type GenS4 struct { F0 int; F1 int64; F2 int8 }

func main() {
    v1 := []string{("" + "foo bar")}
    v2 := []int{1, -9223372036854775808, -157}
    var v3 uint32 = 81
    var v4 int = 2147483647
    var v5 float32 = 0.1
    _ = len(v1)
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 uint32 = 633
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 20; fi9++ {
            fmt.Println(v5)
            v3 = 880
            _ = fi9
        }
        fmt.Println(v5)
        _ = fi6
    }
    if ((false || false) || (v5 > ((1.5 + -2.25) / float32(float64(1e10))))) {
        if false {
            var v10 int8 = 101
            _ = v10
        } else {
            var v11 int64 = 956
            v12 := GenS1{F0: (^(v11)), F1: 32767, F2: v11}
            _ = v11
            _ = v12
        }
        v5 = float32(0.0)
    } else {
        v13 := GenS2{F0: 42}
        var v14 uint16 = 42
        v15 := GenS1{F0: 1, F1: v14, F2: 901}
        _ = v13
        _ = v14
        _ = v15
    }
    v5 = 3.14159
    switch "foo bar" {
    case "hello":
        var v16 float32 = (v5 - (v5 * 1e10))
        _ = v16
    case "":
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 uint8 = 118
            if ("\n\t" == (("hello" + "foo bar") + "hello")) {
                if ((((float64((v5 * 3.14159)) * (0.0 - (1e10 + 0.0))) * float64(0.0)) != 3.14159) || true) {
                    if (-1 < 100) {
                        v5 -= v5
                    }
                    var v19 uint32 = 0
                    _ = v19
                }
                v5 = v5
            }
            var v20 uint64 = 18446744073709551615
            _ = fi17
            _ = v18
            _ = v20
        }
    case "foo bar":
        var v21 uint32 = (^(v3))
        _ = v21
    }
    v22 := GenS2{F0: (^(v3))}
    var v23 string = "foo bar"
    if len(v1) > 0 {
        v1[0] = (("a" + "\n\t") + v23)
    }
    var v24 float64 = 1e10
    v23 = ""
    for fi25 := 0; fi25 < 15; fi25++ {
        switch v4 {
        case 32768:
            v26 := GenS1{F0: 9223372036854775807, F1: 372, F2: -1}
            _ = v26
        default:
            for fi27 := 0; fi27 < 5; fi27++ {
                if false {
                    v4 = (v4 / (((v4 - 127)) | 1))
                    _ = len(v2)
                    if len(v2) > 0 {
                        v2[0] = -614
                    }
                } else {
                    v28 := GenS1{F0: -333, F1: 65535, F2: -9223372036854775808}
                    if (-128 >= 100) {
                        switch (v23 + "hello") {
                        case "\n\t":
                            var v29 uint32 = (v3 / ((4294967295) | 1))
                            _ = v29
                        default:
                            fmt.Println(fi25)
                        }
                    }
                    var v30 bool = true
                    _ = v28
                    _ = v30
                }
                _ = len(v2)
                if true {
                    fmt.Println(v5, fi25)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if (42 <= -1) {
                        var v31 uint16 = 641
                        v31 = (v31 + 1)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v31
                    }
                }
                v32 := GenS1{F0: 1, F1: 255, F2: -10}
                _ = fi27
                _ = v32
            }
        }
        var v33 float32 = (0.0 / float32(v24))
        _ = fi25
        _ = v33
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 int64 = -9223372036854775808
        v36 := GenS3{F0: (fi34 >> ((0) & 63)), F1: -41, F2: (^(fi34))}
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 int64 = 42
    v3 += (^(v3))
    if true {
        v3 += v3
        var v38 uint = 2
        _ = v38
    } else {
        for fi39 := 0; fi39 < 12; fi39++ {
            fmt.Println(v5, v5)
            if (!false) {
                v24 += v24
            } else {
                v40 := GenS4{F0: (fi39 + fi39), F1: (v37 | v37), F2: -81}
                var v41 string = ("a" + "\n\t")
                _ = v40
                _ = v41
            }
            v4 = v4
            _ = fi39
        }
        fmt.Println(v5, v24, v37)
        if len(v1) > 0 {
            v1[0] = v23
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
