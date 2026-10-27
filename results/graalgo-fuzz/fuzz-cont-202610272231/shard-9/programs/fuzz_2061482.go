package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint; F2 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 bool }

func main() {
    v1 := []string{"", (("\n\t" + "") + ("\n\t" + "hello")), "\n\t", "a"}
    v2 := make(map[uint32]int)
    var v3 int = 32768
    for fi4 := 0; fi4 < 16; fi4++ {
        v3 = v3
        switch "foo bar" {
        case "a":
            var v5 float64 = -2.25
            _ = v5
        case "\n\t":
            var v6 int64 = -1
            _ = v6
        }
        _ = len(v2)
        _ = fi4
    }
    v3 = v3
    _ = len(v2)
    if (2147483647 < 2147483648) {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 uint8 = 2
            for fi9 := 0; fi9 < 17; fi9++ {
                v8 -= 42
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
    } else {
        if ((2 <= 255) || (534 > 95)) {
            var v10 uint8 = 1
            _ = v10
        }
        if len(v1) > 0 {
            v1[0] = "hello"
        }
    }
    if false {
        if (float64((0.1 - ((-2.25 / 1.0) - 1e10))) > 1.5) {
            if (!true) {
                var v11 uint8 = 88
                _ = v11
            }
            for fi12 := 0; fi12 < 8; fi12++ {
                v3 = fi12
                switch ("" + ("hello" + "a")) {
                case "a":
                    if false {
                        switch (^(fi12)) {
                        case -2147483648:
                            v13 := GenS1{F0: "a", F1: 65535, F2: 229}
                            _ = v13
                        }
                    } else {
                        var v14 int32 = -169
                        v14 -= v14
                        for fi15 := 0; fi15 < 6; fi15++ {
                            v14 = -1
                            var v16 uint64 = 2
                            v17 := GenS1{F0: "foo bar", F1: 0, F2: 255}
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        _ = v14
                    }
                case "foo bar":
                    var v18 int32 = -90
                    _ = v18
                }
                var v19 float32 = 1.5
                _ = fi12
                _ = v19
            }
        }
    }
    var v20 string = ("" + ("foo bar" + "\n\t"))
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := v2[65535]
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        switch (v22 & (v22 & v3)) {
        case 42:
            delete(v2, 2147483648)
        default:
            for fi23 := 0; fi23 < 19; fi23++ {
                if true {
                    v24 := GenS2{F0: -412}
                    _ = v24
                } else {
                    fmt.Println(v22, fi21)
                    v3 += (-(fi23))
                    switch (^(fi21)) {
                    case 1:
                        if (255 < 0) {
                            switch "a" {
                            case "\n\t":
                                v25 := GenS2{F0: -431}
                                _ = len(v1)
                                _ = v25
                            }
                            if false {
                                v26 := GenS3{F0: -32768}
                                _ = v26
                            }
                            v1 = append(v1, "hello")
                        }
                    case 4294967295:
                        fmt.Println(fi21, v3)
                    case -71:
                        _ = len(v1)
                    }
                }
                if false {
                    v3 -= v3
                } else {
                    var v27 int32 = 307
                    _ = v27
                }
                _ = fi23
            }
        }
        _ = fi21
        _ = v22
    }
    var v28 int8 = 96
    v29 := GenS3{F0: 1000}
    var v30 bool = false
    if (true || v30) {
        var v31 uint = 1
        for fi32 := 0; fi32 < 3; fi32++ {
            if ((1e10 != 3.14159) && (!(65535 > v31))) {
                var v33 uint64 = 18446744073709551615
                for fi34 := 0; fi34 < 16; fi34++ {
                    v35 := GenS1{F0: v20, F1: 42, F2: 128}
                    fmt.Println(fi32)
                    v20 = (v20 + ("" + v20))
                    var v36 uint16 = 915
                    v28 -= 127
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                _ = v33
            } else {
                fmt.Println(v3)
                v37 := GenS4{F0: 42, F1: (v31 - v31), F2: 1}
                fmt.Println(v20)
                _ = v37
            }
            fmt.Println(v30)
            _ = fi32
        }
        v38 := GenS1{F0: v20, F1: v31, F2: 255}
        _ = v31
        _ = v38
    }
    fmt.Println(v30)
    if (0 < 255) {
        switch (v3 * v3) {
        case -1:
            var v39 int8 = v28
            _ = v39
        case 9223372036854775807:
            v20 = "\n\t"
        default:
            switch ("foo bar" + v20) {
            case "":
                if (!v30) {
                    switch (v3 | v3) {
                    case 830:
                        fmt.Println(v28, v3)
                    case 2147483647:
                        switch (v3 | 780) {
                        case 127:
                            _ = len(v2)
                        case -2147483648:
                            if (v30 || v30) {
                                v40 := GenS5{F0: 664}
                                if (!v30) {
                                    for fi41 := 0; fi41 < 16; fi41++ {
                                        var v42 int64 = 9223372036854775807
                                        v43 := GenS6{F0: 42, F1: v30}
                                        _ = fi41
                                        _ = v42
                                        _ = v43
                                    }
                                    v28 += (v28 ^ (v28 * v28))
                                    switch 127 {
                                    case 42:
                                        v30 = (2147483647 != 42)
                                    default:
                                        var v44 int32 = 522
                                        _ = v44
                                    }
                                }
                                var v45 int16 = 42
                                _ = v40
                                _ = v45
                            }
                        case -131:
                            fmt.Println(v3, v3, v20)
                        default:
                            v46 := v2[42]
                            _ = v46
                        }
                    case 0:
                        if len(v1) > 0 {
                            v1[0] = v20
                        }
                    }
                    v1 = append(v1, ("a" + ((((v20 + ("foo bar" + ((("" + (v20 + v20)) + "") + v20))) + "a") + v20) + v20)))
                }
            default:
                var v47 int64 = 0
                _ = v47
            }
        }
        _ = len(v2)
        v20 = v20
    } else {
        v48 := GenS5{F0: -1}
        v30 = (9223372036854775807 <= 1)
        _ = v48
    }
    v3 += v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (370 > 1) {
        fmt.Println(v3)
        for fi49 := 0; fi49 < 10; fi49++ {
            v28 = v28
            v50 := GenS1{F0: ("\n\t" + ("" + (v20 + (v20 + (v20 + ("\n\t" + ((v20 + "a") + "hello"))))))), F1: 0, F2: 127}
            var v51 uint32 = 4294967295
            _ = fi49
            _ = v50
            _ = v51
        }
    } else {
        v28 = (v28 + v28)
        _ = len(v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
