package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 int32; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 int64 }

func main() {
    v1 := []int{2147483647}
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    var v4 uint = 42
    var v5 uint16 = 255
    v6 := make(map[uint32]bool)
    var v7 uint64 = 259
    var v8 bool = (42 >= -238)
    if (0 <= 1000) {
        v9 := GenS1{F0: v7, F1: 483}
        var v10 uint64 = 0
        v11 := GenS1{F0: (v10 / ((v10) | 1)), F1: v7}
        _ = v10
        _ = v11
        _ = v9
    } else {
        switch "hello" {
        case "a":
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 int64 = -1
                var v14 bool = ((0.1 - 0.1) < float32((1.5 / 0.1)))
                _ = fi12
                _ = v13
                _ = v14
            }
        case "hello":
            v8 = v8
        default:
            var v15 int64 = -1
            _ = v15
        }
    }
    if v8 {
        v4 = v4
    } else {
        if v8 {
            switch 28 {
            case 42:
                var v16 int = 230
                _ = v16
            }
            var v17 string = ((("a" + "a") + "") + "a")
            _ = v17
        } else {
            fmt.Println(v7, v8, v8)
            if v8 {
                delete(v6, 0)
                if false {
                    v4 = (^(v4))
                } else {
                    var v18 int32 = -967
                    switch -446 {
                    case 188:
                        var v19 bool = v8
                        _ = v19
                    default:
                        if v8 {
                            var v20 uint = v4
                            v18 += 1
                            var v21 uint8 = 2
                            _ = v20
                            _ = v21
                        } else {
                            if v8 {
                                if ((v5 % ((uint16(v4)) | 1)) != uint16(v7)) {
                                    v4 = 65535
                                    v18 += int32(v7)
                                    if (65535 == 1) {
                                        v1 = append(v1, -179)
                                    } else {
                                        if ((!(!false)) && (!v8)) {
                                            for fi22 := 0; fi22 < 3; fi22++ {
                                                var v23 float64 = 1e10
                                                var v24 int = fi22
                                                v25 := GenS1{F0: (^(v7)), F1: v7}
                                                _ = fi22
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            }
                                            for fi26 := 0; fi26 < 11; fi26++ {
                                                var v27 uint8 = 65
                                                var v28 string = ("\n\t" + "\n\t")
                                                var v29 uint64 = v7
                                                _ = fi26
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                            }
                                            fmt.Println(v4, v5)
                                        }
                                        var v30 int8 = int8(v5)
                                        var v31 int16 = -32768
                                        _ = v30
                                        _ = v31
                                    }
                                }
                                if ((100 != -1) && (v8 || (1.5 >= 3.14159))) {
                                    v32 := GenS2{F0: v5}
                                    _ = v32
                                } else {
                                    fmt.Println(v8, v4)
                                }
                                v6[0] = v8
                            }
                        }
                    }
                    v5 = v5
                    _ = v18
                }
                v33 := GenS1{F0: v7, F1: v7}
                _ = v33
            } else {
                var v34 uint64 = (v7 * (v7 + v7))
                _ = v34
            }
        }
        if ((!true) || true) {
            v35 := GenS3{F0: v5}
            v36 := GenS4{F0: -156}
            _ = v35
            _ = v36
        } else {
            fmt.Println(v8, v4, v7)
            v5 += v5
            v37 := GenS5{F0: ((!(!(false || v8))) || false), F1: -889, F2: -60}
            _ = v37
        }
        if v8 {
            var v38 string = "a"
            var v39 int8 = -127
            _ = v38
            _ = v39
        }
    }
    v40 := GenS6{F0: 916}
    v41 := GenS5{F0: (671 >= 41), F1: 0, F2: 42}
    v8 = (!true)
    var v42 int = 32767
    v8 = v8
    for fi43 := 0; fi43 < 11; fi43++ {
        var v44 uint32 = 2147483647
        fmt.Println(v5, v44, v4)
        if ("a" != ("foo bar" + "")) {
            switch ((("" + "") + "\n\t") + ("hello" + ("a" + (((("\n\t" + "\n\t") + "hello") + "foo bar") + "a")))) {
            case "a":
                for fi45 := 0; fi45 < 13; fi45++ {
                    var v46 uint8 = uint8(v5)
                    var v47 uint16 = 1
                    fmt.Println(v44)
                    switch ("a" + ("a" + "a")) {
                    case "\n\t":
                        if false {
                            v7 = (v7 | v7)
                        }
                    case "":
                        if false {
                            var v48 int = (v42 + (v42 * (v42 & v42)))
                            v49 := GenS5{F0: ("hello" > "hello"), F1: 908, F2: 1000}
                            _ = v48
                            _ = v49
                        } else {
                            var v50 int16 = -1
                            _ = v50
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = int(v7)
                    }
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            case "\n\t":
                var v51 bool = v8
                _ = v51
            }
            v8 = (651 != (v7 + (v7 ^ v7)))
        } else {
            for fi52 := 0; fi52 < 7; fi52++ {
                v53 := GenS1{F0: v7, F1: (v7 % (((v7 / (((v7 - v7)) | 1))) | 1))}
                switch "" {
                case "hello":
                    delete(v6, 4294967295)
                case "a":
                    if len(v1) > 0 {
                        v1[0] = (fi52 ^ 823)
                    }
                case "\n\t":
                    fmt.Println(v44, v8)
                }
                _ = fi52
                _ = v53
            }
            v54 := GenS7{F0: 42, F1: 9223372036854775807}
            _ = v54
        }
        _ = fi43
        _ = v44
    }
    var v55 uint32 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v55
    _ = v6
    _ = v7
    _ = v8
}
