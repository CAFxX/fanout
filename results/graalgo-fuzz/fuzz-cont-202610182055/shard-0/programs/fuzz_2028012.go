package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 bool; F1 uint32; F2 int8 }
type GenS6 struct { F0 int8 }

func main() {
    var v1 float32 = -2.25
    var v2 int64 = 0
    v3 := []uint32{65535, 0, 1, 2147483648}
    var v4 uint = 459
    var v5 bool = true
    var v6 uint = 0
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: (v6 * v6), F1: ("\n\t" + "hello"), F2: -32768}
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int8 = -1
            var v11 int8 = (v10 + (v10 | (v10 * int8(v2))))
            fmt.Println(v1)
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            if true {
                v5 = true
                v13 := GenS1{F0: v6, F1: "\n\t", F2: 32767}
                _ = v13
            }
            _ = fi12
        }
        _ = fi7
        _ = v8
    }
    var v14 int64 = v2
    v15 := GenS2{F0: (v4 / (((v4 * v4)) | 1))}
    if (127 <= -9) {
        v5 = false
        v2 = v2
        var v16 int64 = (^(v2))
        _ = v16
    }
    switch ("" + ("\n\t" + "\n\t")) {
    case "\n\t":
        var v17 uint64 = 608
        _ = v17
    }
    v2 -= v14
    switch "hello" {
    case "a":
        v1 -= 3.14159
    case "foo bar":
        _ = len(v3)
    case "":
        fmt.Println(v2, v1)
    }
    var v18 float32 = float32(1.5)
    fmt.Println(v5)
    if (!true) {
        var v19 bool = (98 > 105)
        _ = v19
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 int32 = -542
            _ = fi20
            _ = v21
        }
        v14 = (^(v2))
    }
    _ = len(v3)
    for fi22 := 0; fi22 < 1; fi22++ {
        switch "foo bar" {
        case "hello":
            if (!v5) {
                var v23 string = "foo bar"
                if ((^(v6)) < v4) {
                    v1 = ((v1 - v1) + 1e10)
                    switch "hello" {
                    case "\n\t":
                        var v24 int32 = -648
                        _ = v24
                    case "":
                        if len(v3) > 0 {
                            v3[0] = 2
                        }
                    case "a":
                        if (0.1 >= -2.25) {
                            v18 = 0.0
                        } else {
                            v23 = (v23 + "hello")
                            fmt.Println(v2, v1, v5)
                        }
                    }
                    for fi25 := 0; fi25 < 9; fi25++ {
                        var v26 int16 = 853
                        fmt.Println(fi22, v2, fi25)
                        _ = fi25
                        _ = v26
                    }
                }
                if (!(128 >= 255)) {
                    var v27 int16 = -1
                    switch ((v23 + "foo bar") + (v23 + "hello")) {
                    case "a":
                        switch (fi22 << ((255) & 63)) {
                        case 127:
                            var v28 bool = false
                            _ = v28
                        }
                    case "hello":
                        var v29 uint16 = 32767
                        _ = v29
                    case "":
                        var v30 int = -9223372036854775808
                        _ = v30
                    }
                    v18 = 1e10
                    _ = v27
                }
                _ = v23
            }
        }
        v31 := GenS3{F0: (fi22 ^ (fi22 ^ fi22))}
        var v32 int16 = -1
        _ = fi22
        _ = v31
        _ = v32
    }
    fmt.Println(v18, v4, v4)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v5, v5)
    default:
        for fi33 := 0; fi33 < 18; fi33++ {
            if (42 >= 127) {
                for fi34 := 0; fi34 < 1; fi34++ {
                    var v35 string = "hello"
                    if (("hello" + "foo bar") != ("\n\t" + "hello")) {
                        if v5 {
                            v5 = (false && ((1 == 32767) || true))
                            v36 := GenS1{F0: 4294967295, F1: v35, F2: 32767}
                            _ = v36
                        } else {
                            v37 := GenS4{F0: 65535, F1: 2147483647}
                            v14 += (v14 * 1)
                            _ = v37
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v2 -= v2
                    }
                    if (v35 <= ("" + v35)) {
                        v38 := GenS5{F0: (33 > 42), F1: 573, F2: 42}
                        var v39 float32 = v1
                        _ = v38
                        _ = v39
                    }
                    _ = fi34
                    _ = v35
                }
                var v40 uint32 = 2147483648
                _ = v40
            }
            var v41 float32 = v1
            v42 := GenS4{F0: 42, F1: -1}
            _ = fi33
            _ = v41
            _ = v42
        }
    }
    v43 := GenS6{F0: 0}
    switch 14 {
    case 42:
        var v44 int8 = 77
        _ = v44
    case 128:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        switch -1 {
        case 427:
            v14 = (v14 - 42)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
