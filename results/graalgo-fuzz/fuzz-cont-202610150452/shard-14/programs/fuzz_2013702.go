package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int64; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint8; F1 int64; F2 uint64 }
func genFn0() int {
    v5 := []int{}
    var v6 float32 = 0.0
    _ = v5
    _ = v6
    return 128
}

func main() {
    v1 := []int{0, -9223372036854775808, 4294967295}
    v2 := make(map[string]string)
    var v3 float32 = 3.14159
    var v4 int16 = -286
    var v7 uint = 124
    var v8 int32 = 2147483647
    var v9 int = 9223372036854775807
    v10 := GenS1{F0: 430}
    v3 -= 3.14159
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 bool = false
    v11 = ((v11 || true) && (!v11))
    var v12 int16 = v4
    if len(v1) > 0 {
        v1[0] = v9
    }
    switch genFn0() {
    case 42:
        var v13 int8 = -1
        _ = v13
    case 32767:
        v3 = float32(0.1)
    }
    if v11 {
        v4 += (v12 << ((1) & 63))
        v14 := GenS2{F0: 1, F1: (v7 - v7)}
        if (0 >= -9223372036854775808) {
            v8 = v8
        }
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 uint32 = 0
        for fi17 := 0; fi17 < 2; fi17++ {
            if v11 {
                var v18 int16 = v12
                var v19 int64 = -24
                fmt.Println(v3, v19)
                _ = v18
                _ = v19
            } else {
                var v20 int64 = 1
                switch "hello" {
                case "hello":
                    v1 = append(v1, (fi17 - int(v16)))
                default:
                    v11 = true
                }
                _ = v20
            }
            switch 42 {
            case -907:
                v9 += v9
            case 127:
                v11 = true
            case 9223372036854775807:
                fmt.Println(v12, v16)
            default:
                if (float64(-2.25) <= (1e10 - 1.5)) {
                    fmt.Println(v7, v3, fi15)
                    v21 := GenS3{F0: v11, F1: ((v8 & (v8 - (v8 % ((731) | 1)))) == (v8 / ((v8) | 1)))}
                    _ = v21
                } else {
                    if (0.0 > v3) {
                        v2["a"] = ("\n\t" + "a")
                    } else {
                        if (v11 || ((!true) || ((fi17 << ((v7) & 63)) > -9223372036854775808))) {
                            v4 -= v4
                            var v22 uint8 = 102
                            _ = v22
                        } else {
                            var v23 int8 = -113
                            _ = v23
                        }
                        var v24 int16 = v12
                        _ = v24
                    }
                }
            }
            switch v9 {
            case 128:
                var v25 bool = false
                v4 = (v12 | -325)
                switch (^(v9)) {
                case 0:
                    v9 -= fi17
                case -2147483648:
                    v26 := GenS4{F0: 336, F1: 255, F2: v8}
                    _ = v26
                case -1:
                    v9 -= fi15
                default:
                    fmt.Println(v11)
                    switch fi15 {
                    case 9223372036854775807:
                        v16 = (v16 & 2)
                    case -953:
                        if ((float32(float64(float32(0.1))) >= (0.1 / float32(1e10))) && false) {
                            var v27 float32 = v3
                            fmt.Println(v4, v25, v9)
                            v28 := GenS2{F0: (^(v16)), F1: v7}
                            _ = v27
                            _ = v28
                        }
                    case -1:
                        v11 = (v25 && v11)
                    }
                }
                _ = v25
            case 4294967295:
                if v11 {
                    v8 = v8
                    if false {
                        var v29 int8 = -128
                        _ = v29
                    }
                    var v30 int64 = -9223372036854775808
                    _ = v30
                } else {
                    fmt.Println(fi17, v4)
                    switch "" {
                    case "":
                        v8 = v8
                    case "\n\t":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch (fi17 - 2147483647) {
                        case 32767:
                            switch "hello" {
                            case "\n\t":
                                var v31 int8 = 0
                                var v32 int8 = -128
                                var v33 int = (v9 / ((genFn0()) | 1))
                                v31 = v31
                                if (v11 && ((v16 ^ v16) == (v16 % (((v16 | (v16 & v16))) | 1)))) {
                                    var v34 bool = true
                                    var v35 string = "hello"
                                    v36 := GenS4{F0: 42, F1: 2, F2: v8}
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                                _ = v31
                                _ = v32
                                _ = v33
                            case "hello":
                                v8 += v8
                            case "foo bar":
                                v4 += 1000
                            }
                            var v37 float64 = (-2.25 - 0.1)
                            v38 := GenS1{F0: 718}
                            v39 := GenS1{F0: 2}
                            _ = v37
                            _ = v38
                            _ = v39
                        case -2147483648:
                            switch ((("hello" + "a") + "\n\t") + "hello") {
                            case "\n\t":
                                v7 = 869
                            case "hello":
                                if true {
                                    fmt.Println(v4, fi17, v16)
                                    if v11 {
                                        switch ("hello" + "") {
                                        case "hello":
                                            fmt.Println(v12, v12)
                                        }
                                        v8 += -1
                                        if (true && (!(v11 && v11))) {
                                            v11 = v11
                                            switch (fi15 % ((fi17) | 1)) {
                                            case 32768:
                                                delete(v2, "\n\t")
                                            case 0:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            default:
                                                v40 := GenS1{F0: 1}
                                                _ = v40
                                            }
                                            v4 += (v4 + -427)
                                        } else {
                                            var v41 int32 = (v8 + (v8 << (((v7 % ((286) | 1))) & 63)))
                                            var v42 int64 = 0
                                            _ = v41
                                            _ = v42
                                        }
                                        v8 -= (v8 + v8)
                                        fmt.Println(v12, v4, v8)
                                        v4 -= 42
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = genFn0()
                                        }
                                    }
                                }
                            }
                        case 127:
                            if (127 == 128) {
                                var v43 uint = v7
                                fmt.Println(v12, v7)
                                v44 := v2["a"]
                                if (42 >= 1) {
                                    v45 := GenS2{F0: (v16 ^ (v16 % ((v16) | 1))), F1: (v7 ^ (^(v43)))}
                                    v46 := GenS2{F0: v16, F1: (v43 / ((v7) | 1))}
                                    var v47 int16 = (v4 | v4)
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                }
                                _ = v43
                                _ = v44
                            }
                        default:
                            var v48 int16 = (v12 & 1000)
                            _ = v48
                        }
                    case "foo bar":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        v16 = 353
                    }
                }
            default:
                if v11 {
                    v9 -= int(v12)
                }
            }
            _ = fi17
        }
        var v49 uint64 = 787
        _ = fi15
        _ = v16
        _ = v49
    }
    switch int(v8) {
    case -9223372036854775808:
        delete(v2, "foo bar")
    case -1:
        switch "hello" {
        case "\n\t":
            if false {
                var v50 int16 = -32768
                _ = v50
            } else {
                var v51 uint64 = 18446744073709551615
                v52 := GenS5{F0: 0, F1: -1, F2: v51}
                _ = v51
                _ = v52
            }
        case "a":
            v53 := GenS3{F0: (v7 >= 531), F1: v11}
            _ = v53
        case "foo bar":
            v54 := GenS5{F0: 38, F1: -179, F2: 42}
            _ = v54
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
