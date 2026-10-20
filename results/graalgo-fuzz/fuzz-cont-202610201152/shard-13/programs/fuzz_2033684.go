package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int; F2 int16 }
type GenS3 struct { F0 uint32; F1 int64; F2 string }
type GenS4 struct { F0 int16 }

func main() {
    var v1 string = "foo bar"
    var v2 int = -856
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    var v5 float32 = (0.0 / float32(0.1))
    var v6 int64 = 138
    if false {
        var v7 uint64 = 793
        v6 -= (v6 - v6)
        _ = v7
    }
    if (!true) {
        if ((32767 == 271) || (1 < 948)) {
            if (9223372036854775807 < int64(v2)) {
                var v8 uint32 = 2147483647
                if (!((!(false && (false && false))) || true)) {
                    v9 := GenS1{F0: 0}
                    v2 = (^(v2))
                    v6 += v6
                    _ = v9
                } else {
                    v10 := v4[v1]
                    switch "foo bar" {
                    case "hello":
                        var v11 uint8 = 1
                        _ = v11
                    }
                    _ = v10
                }
                _ = v8
            } else {
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 int = -628
                    v14 := v3[65535]
                    var v15 int64 = -993
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            }
            var v16 float64 = (1e10 * 0.0)
            _ = v16
        }
    }
    switch 1 {
    case 4294967295:
        v1 = "\n\t"
    case 938:
        switch v2 {
        case 9223372036854775807:
            switch ("\n\t" + "") {
            case "a":
                v17 := GenS1{F0: (-(v2))}
                _ = v17
            case "":
                v18 := []int{-237}
                _ = v18
            }
        case 1:
            switch (("\n\t" + "") + "foo bar") {
            case "hello":
                v2 = (v2 / ((-869) | 1))
            }
        case 128:
            v1 = (v1 + v1)
        }
    }
    v19 := []int{v2, (v2 * (v2 >> ((483) & 63))), 0}
    switch v2 {
    case -333:
        if (false && (1000 != 1000)) {
            var v20 int64 = (v6 >> ((2) & 63))
            v5 -= ((v5 * -2.25) + 3.14159)
            fmt.Println(v20, v5)
            _ = v20
        } else {
            var v21 int = 219
            v6 += v6
            _ = v21
        }
    case 4294967295:
        fmt.Println(v6, v5, v1)
    default:
        switch (v2 * v2) {
        case 42:
            if false {
                if false {
                    var v22 float64 = 0.1
                    if true {
                        if (-20 <= -92) {
                            if (true && true) {
                                if (0.1 < (v5 + 3.14159)) {
                                    var v23 string = "hello"
                                    _ = v23
                                }
                            }
                        }
                    } else {
                        v24 := GenS2{F0: -1, F1: v2, F2: -32768}
                        v6 = -525
                        var v25 int = v2
                        _ = v24
                        _ = v25
                    }
                    _ = v22
                } else {
                    v2 += v2
                    switch (v1 + v1) {
                    case "foo bar":
                        if true {
                            for fi26 := 0; fi26 < 15; fi26++ {
                                for fi27 := 0; fi27 < 14; fi27++ {
                                    switch (v1 + v1) {
                                    case "a":
                                        var v28 uint = 18446744073709551615
                                        _ = v28
                                    case "hello":
                                        v5 += 0.1
                                    default:
                                        switch v1 {
                                        case "foo bar":
                                            var v29 uint64 = 2
                                            _ = v29
                                        case "a":
                                            var v30 int8 = -1
                                            _ = v30
                                        default:
                                            fmt.Println(fi26)
                                        }
                                    }
                                    v31 := GenS1{F0: 32767}
                                    _ = fi27
                                    _ = v31
                                }
                                v19 = append(v19, (^(v2)))
                                _ = fi26
                            }
                            _ = len(v4)
                            v32 := GenS1{F0: (v2 ^ v2)}
                            _ = v32
                        } else {
                            fmt.Println(v1)
                            var v33 float32 = 1e10
                            _ = v33
                        }
                    case "\n\t":
                        v34 := GenS1{F0: -168}
                        _ = v34
                    default:
                        switch v2 {
                        case 890:
                            v2 = (v2 << ((1) & 63))
                        default:
                            var v35 uint32 = 652
                            _ = v35
                        }
                    }
                }
            }
        case 4294967295:
            if (v6 >= v6) {
                v36 := GenS2{F0: v6, F1: (v2 * v2), F2: -32768}
                if false {
                    v6 = 1
                    var v37 uint8 = 2
                    if (true || ((v37 - v37) == 255)) {
                        var v38 float64 = (-2.25 + (float64((((v5 - v5) / float32(0.1)) + v5)) * 3.14159))
                        _ = v38
                    }
                    _ = v37
                }
                fmt.Println(v2, v6)
                _ = v36
            }
        }
    }
    var v39 float32 = 3.14159
    if (v6 <= (v6 / (((^(v6))) | 1))) {
        if false {
            v40 := GenS2{F0: v6, F1: (-(v2)), F2: 1000}
            fmt.Println(v5, v6, v2)
            _ = v40
        }
        v41 := GenS1{F0: v2}
        for fi42 := 0; fi42 < 16; fi42++ {
            var v43 int32 = -2147483648
            var v44 int8 = 90
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = v41
    } else {
        for fi45 := 0; fi45 < 15; fi45++ {
            v46 := GenS3{F0: 210, F1: v6, F2: ((((((v1 + ((v1 + v1) + ("\n\t" + v1))) + v1) + ("a" + v1)) + "a") + "a") + v1)}
            delete(v4, "")
            _ = fi45
            _ = v46
        }
        v47 := v3[42]
        var v48 int32 = -804
        _ = v47
        _ = v48
    }
    v49 := GenS2{F0: v6, F1: v2, F2: -219}
    if false {
        v6 = (v6 * (v6 - (v6 >> ((4294967295) & 63))))
        v50 := GenS4{F0: 0}
        _ = v50
    }
    fmt.Println(v6, v2, v1)
    var v51 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v49
    _ = v5
    _ = v51
    _ = v6
}
