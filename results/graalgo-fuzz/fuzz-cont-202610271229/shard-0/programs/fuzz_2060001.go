package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int; F1 string; F2 int64 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 uint64; F1 bool }

func main() {
    v1 := []string{("a" + ("\n\t" + "foo bar")), ((("" + "foo bar") + ("a" + "a")) + "foo bar")}
    v2 := make(map[int]bool)
    var v3 uint = 255
    if (!true) {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    }
    for fi4 := 0; fi4 < 14; fi4++ {
        if ((fi4 % ((128) | 1)) > (fi4 + fi4)) {
            v3 += (v3 & v3)
            switch fi4 {
            case 42:
                _ = len(v2)
            case 32768:
                if (-1 >= -32768) {
                    fmt.Println(fi4, v3)
                    _ = len(v1)
                    for fi5 := 0; fi5 < 13; fi5++ {
                        switch fi5 {
                        case -2147483648:
                            v6 := GenS1{F0: 2, F1: "", F2: false}
                            _ = v6
                        case 89:
                            v3 = (v3 - v3)
                        default:
                            v7 := GenS1{F0: 1, F1: "", F2: true}
                            _ = v7
                        }
                        _ = fi5
                    }
                } else {
                    var v8 int = 32767
                    if false {
                        v3 = (v3 << ((v3) & 63))
                        if (0 < 0) {
                            switch v8 {
                            case 127:
                                switch fi4 {
                                case 32768:
                                    v9 := GenS2{F0: v3}
                                    _ = v9
                                default:
                                    v10 := v2[fi4]
                                    _ = v10
                                }
                            }
                        }
                    } else {
                        switch -9223372036854775808 {
                        case 42:
                            v8 = (fi4 >> ((v3) & 63))
                        }
                        for fi11 := 0; fi11 < 1; fi11++ {
                            v12 := GenS3{F0: 168, F1: uint8(v3)}
                            v1 = append(v1, "")
                            if (-20 >= -128) {
                                v13 := v2[fi11]
                                v14 := GenS4{F0: -128, F1: -2147483648}
                                v8 += (fi4 << ((v3) & 63))
                                _ = v13
                                _ = v14
                            } else {
                                fmt.Println(fi11, v3, v3)
                            }
                            _ = fi11
                            _ = v12
                        }
                        var v15 int32 = 2147483647
                        _ = v15
                    }
                    var v16 string = "foo bar"
                    _ = v16
                    _ = v8
                }
            }
            v3 = (v3 * 725)
        } else {
            fmt.Println(v3, v3)
            var v17 float32 = (float32(1.5) + float32((0.0 - -2.25)))
            var v18 int = 2147483647
            _ = v17
            _ = v18
        }
        var v19 int = (-(fi4))
        _ = fi4
        _ = v19
    }
    if ((v3 < v3) && (true && (!((-620 == 6) && (42 <= 0))))) {
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        v3 = 18446744073709551615
    }
    fmt.Println(v3)
    var v20 int16 = 42
    var v21 uint64 = 2
    for fi22 := 0; fi22 < 15; fi22++ {
        for fi23 := 0; fi23 < 2; fi23++ {
            v20 -= (v20 / ((42) | 1))
            if (0 == 127) {
                if (!true) {
                    switch "\n\t" {
                    case "a":
                        v21 += v21
                    case "foo bar":
                        v24 := GenS5{F0: (fi22 % ((fi22) | 1)), F1: "hello", F2: 55}
                        _ = v24
                    case "":
                        v20 = -1
                    default:
                        switch "foo bar" {
                        case "":
                            var v25 int16 = v20
                            _ = v25
                        }
                    }
                    v21 = (v21 % ((v21) | 1))
                    if false {
                        v26 := GenS1{F0: (^(v21)), F1: "foo bar", F2: (!((-1 <= 1) && false))}
                        delete(v2, 69)
                        _ = v26
                    }
                }
                var v27 uint32 = 1
                _ = v27
            } else {
                v21 = (v21 % ((211) | 1))
                v28 := GenS6{F0: 0, F1: v20}
                switch fi23 {
                case 344:
                    var v29 int8 = -128
                    _ = v29
                case 0:
                    v21 -= 2
                case 32767:
                    v21 = uint64(v3)
                }
                v30 := GenS5{F0: (fi23 | fi23), F1: ((("a" + "a") + "hello") + "\n\t"), F2: 9223372036854775807}
                _ = v28
                _ = v30
            }
            if true {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(fi22)
                case "foo bar":
                    v2[fi22] = (42 > 1)
                case "":
                    switch (fi23 - fi22) {
                    case -94:
                        v3 += v3
                    }
                default:
                    fmt.Println(v3, v21, fi23)
                }
                var v31 string = "hello"
                var v32 uint8 = 1
                _ = v31
                _ = v32
            } else {
                var v33 uint8 = 143
                if (true && false) {
                    v34 := GenS3{F0: v33, F1: v33}
                    v35 := GenS2{F0: (v3 / ((v3) | 1))}
                    _ = v34
                    _ = v35
                } else {
                    v3 = 1
                    v21 = 4294967295
                }
                v33 -= v33
                _ = v33
            }
            _ = fi23
        }
        _ = fi22
    }
    if false {
        v36 := v2[0]
        v36 = v36
        _ = v36
    }
    v1 = append(v1, "hello")
    if true {
        switch "" {
        case "hello":
            fmt.Println(v3, v20, v20)
        case "foo bar":
            fmt.Println(v21, v20, v20)
        case "\n\t":
            v2[667] = ((!(-1 != -128)) || true)
        default:
            v20 = v20
        }
    }
    switch 327 {
    case -2147483648:
        v37 := GenS7{F0: (v21 | (v21 + v21)), F1: true}
        _ = v37
    default:
        v20 += 32767
    }
    v1 = append(v1, (("hello" + ("hello" + "foo bar")) + ("a" + "foo bar")))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
