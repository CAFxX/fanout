package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 uint8; F2 uint64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64 }

func main() {
    var v1 uint = 1
    var v2 int32 = -1
    v3 := []string{"\n\t", "", "hello"}
    v4 := make(map[int]bool)
    var v5 uint = (v1 % ((v1) | 1))
    if len(v3) > 0 {
        v3[0] = "a"
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!true) {
        if (!((v1 * 273) < (v5 % (((v1 * v5)) | 1)))) {
            v5 += 738
            var v6 uint16 = 32767
            fmt.Println(v1, v2, v1)
            _ = v6
        }
        if (((true || (-128 < 42)) && (!true)) || true) {
            v7 := GenS1{F0: v2, F1: 42}
            if (0.1 != -2.25) {
                v5 = v1
            }
            _ = v7
        } else {
            switch "hello" {
            case "foo bar":
                v8 := GenS1{F0: -32, F1: 127}
                _ = v8
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            var v9 uint64 = 2
            _ = v9
        }
    }
    switch -298 {
    case 32767:
        var v10 int16 = 32767
        _ = v10
    case 128:
        if true {
            var v11 int16 = -1
            _ = v11
        }
    case 1:
        if (!(105 == 2)) {
            fmt.Println(v1, v2)
            for fi12 := 0; fi12 < 9; fi12++ {
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 uint32 = 1
                    _ = len(v4)
                    _ = fi13
                    _ = v14
                }
                var v15 uint16 = 1
                _ = fi12
                _ = v15
            }
        } else {
            var v16 bool = (3.14159 != float32(-2.25))
            v17 := GenS2{F0: -75, F1: (v2 << (((v5 | v5)) & 63))}
            _ = len(v4)
            _ = v16
            _ = v17
        }
    }
    if (!true) {
        v18 := GenS3{F0: 782, F1: v1}
        v19 := GenS2{F0: 89, F1: (^(v2))}
        switch -500 {
        case 0:
            switch "a" {
            case "\n\t":
                v1 += (v1 & (^(v1)))
            case "":
                v20 := GenS4{F0: 42}
                _ = v20
            }
        case 4294967295:
            var v21 int = 1
            _ = v21
        default:
            if (true || (32767 <= uint16(v1))) {
                if (!true) {
                    v22 := v4[-9223372036854775808]
                    if v22 {
                        switch -189 {
                        case 42:
                            _ = len(v3)
                        case 32768:
                            switch ("a" + ("foo bar" + "foo bar")) {
                            case "a":
                                if len(v3) > 0 {
                                    v3[0] = ("" + ("hello" + "\n\t"))
                                }
                            case "":
                                v3 = append(v3, "foo bar")
                            }
                        case 431:
                            for fi23 := 0; fi23 < 19; fi23++ {
                                var v24 float64 = ((0.0 * 0.0) * 0.1)
                                var v25 string = "foo bar"
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                        default:
                            if len(v3) > 0 {
                                v3[0] = ("\n\t" + "foo bar")
                            }
                        }
                        _ = len(v3)
                        v26 := GenS5{F0: ("foo bar" + ""), F1: 223, F2: 4294967295}
                        _ = v26
                    }
                    _ = v22
                }
            } else {
                var v27 string = "\n\t"
                _ = v27
            }
        }
        _ = v18
        _ = v19
    } else {
        v2 = 1
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        v29 := GenS6{F0: -32768}
        var v30 float64 = 1.5
        var v31 int8 = 0
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    var v32 int16 = 32767
    fmt.Println(v5, v5, v2)
    v33 := GenS7{F0: 216}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
