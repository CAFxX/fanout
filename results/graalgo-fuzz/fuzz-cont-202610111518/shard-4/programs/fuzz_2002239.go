package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int32; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 uint16; F2 int }

func main() {
    v1 := []string{(("foo bar" + "foo bar") + ("foo bar" + "foo bar"))}
    var v2 float32 = ((3.14159 + (3.14159 + (0.0 * 1e10))) / 0.1)
    var v3 string = "a"
    var v4 uint16 = 1
    var v5 int32 = 388
    fmt.Println(v5, v5, v4)
    switch "foo bar" {
    case "hello":
        var v6 int64 = -1
        _ = v6
    case "":
        fmt.Println(v3, v4)
    case "a":
        switch 127 {
        case 42:
            if false {
                v7 := GenS1{F0: 65535}
                v5 -= (v5 << ((4294967295) & 63))
                _ = v7
            } else {
                if true {
                    v4 = 42
                } else {
                    var v8 float32 = -2.25
                    _ = v8
                }
            }
        case 2147483647:
            var v9 uint = 42
            _ = v9
        default:
            if false {
                v5 += (-(v5))
            }
        }
    default:
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 int16 = -32768
            _ = fi10
            _ = v11
        }
    }
    var v12 int8 = 127
    fmt.Println(v2, v2, v5)
    var v13 int16 = -1
    v1 = append(v1, "")
    if true {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 bool = true
            _ = fi14
            _ = v15
        }
        switch -9223372036854775808 {
        case 128:
            v16 := GenS2{F0: 0, F1: 667, F2: 2}
            _ = v16
        case 42:
            switch (v3 + "a") {
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "hello":
                v17 := GenS3{F0: 4294967295, F1: (v5 ^ (v5 - (v5 - (v5 * v5)))), F2: 627}
                _ = v17
            }
        case 967:
            _ = len(v1)
        }
        if true {
            var v18 int64 = 9223372036854775807
            _ = v18
        } else {
            if false {
                if true {
                    var v19 uint32 = 0
                    if (float32((1e10 - (0.0 / 1e10))) != ((v2 - v2) + 0.0)) {
                        for fi20 := 0; fi20 < 18; fi20++ {
                            if (128 <= 624) {
                                if (1.5 == ((3.14159 + float64(v2)) - 3.14159)) {
                                    var v21 int64 = 915
                                    if len(v1) > 0 {
                                        v1[0] = "hello"
                                    }
                                    _ = v21
                                }
                                for fi22 := 0; fi22 < 18; fi22++ {
                                    v3 = v3
                                    if (((0.0 != float64(v2)) && true) && ((fi20 | (fi20 >> ((2) & 63))) >= fi20)) {
                                        switch (fi22 & (fi22 / (((fi22 * fi20)) | 1))) {
                                        case 9223372036854775807:
                                            v13 += v13
                                        case 42:
                                            v23 := GenS1{F0: 4294967295}
                                            _ = v23
                                        }
                                        fmt.Println(fi20, fi20, v3)
                                        if len(v1) > 0 {
                                            v1[0] = ""
                                        }
                                    }
                                    var v24 int64 = 569
                                    v1 = append(v1, "hello")
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi22
                                    _ = v24
                                }
                            }
                            switch v3 {
                            case "foo bar":
                                fmt.Println(v4, v13, v19)
                            }
                            _ = fi20
                        }
                        for fi25 := 0; fi25 < 6; fi25++ {
                            fmt.Println(v5, fi25, v4)
                            _ = fi25
                        }
                        var v26 int8 = (-(v12))
                        _ = v26
                    } else {
                        v13 -= (v13 ^ -504)
                    }
                    _ = v19
                } else {
                    switch int(v5) {
                    case 704:
                        v27 := GenS4{F0: 238}
                        _ = v27
                    }
                    if len(v1) > 0 {
                        v1[0] = v3
                    }
                    v12 = 102
                }
                var v28 uint = 42
                _ = len(v1)
                _ = v28
            } else {
                var v29 string = "\n\t"
                v30 := GenS2{F0: 42, F1: 0, F2: 1}
                _ = v29
                _ = v30
            }
        }
    }
    if ((v5 << ((447) & 63)) > v5) {
        v1 = append(v1, (v3 + v3))
    } else {
        v31 := GenS5{F0: 42, F1: v4, F2: 373}
        _ = v31
    }
    v4 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
