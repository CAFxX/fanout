package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int16; F1 int64; F2 uint32 }
func genFn0(p0 int, p1 int8, p2 int16) int16 {
    var v4 float32 = 1.5
    if ((p0 ^ int(p1)) > 361) {
        v5 := []int{(-(p0)), (p0 / ((p0) | 1))}
        fmt.Println(p1, p2, p2)
        if false {
            var v6 uint16 = 2
            v7 := GenS1{F0: 255}
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 16; fi8++ {
                switch p0 {
                case -2147483648:
                    p2 = 92
                default:
                    var v9 uint8 = 127
                    _ = v9
                }
                v5 = append(v5, -916)
                _ = fi8
            }
        }
        _ = v5
    }
    if ((0.1 + (v4 / (float32(3.14159) / v4))) > (v4 / v4)) {
        var v10 float64 = (3.14159 + -2.25)
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := []int{(p0 * (p0 >> ((690) & 63))), (fi11 % ((127) | 1)), 42, (^(p0))}
        var v13 int64 = 804
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v4
    return 205
}

func main() {
    v1 := []uint32{}
    var v2 bool = true
    v3 := []int{397}
    v15 := GenS1{F0: 128}
    if (100 < 1) {
        var v16 int16 = 0
        fmt.Println(v16)
        _ = v16
    }
    switch 32768 {
    case -608:
        switch 0 {
        case -498:
            var v17 uint = 251
            _ = v17
        case 0:
            var v18 int16 = -32768
            _ = v18
        default:
            if (!(("a" < "") || (4294967295 > 2147483648))) {
                var v19 bool = (-1 <= 911)
                _ = v19
            } else {
                switch ("a" + "") {
                case "":
                    v2 = (255 <= 116)
                case "a":
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                case "foo bar":
                    fmt.Println(v2)
                default:
                    v20 := GenS2{F0: 4294967295, F1: ("\n\t" + ("a" + "a"))}
                    _ = v20
                }
            }
        }
    case 32768:
        if true {
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
        } else {
            v2 = v2
            if len(v3) > 0 {
                v3[0] = 0
            }
            var v21 uint16 = 65535
            _ = v21
        }
    default:
        switch (("a" + "hello") + "") {
        case "\n\t":
            fmt.Println(v2)
        case "hello":
            var v22 float32 = -2.25
            _ = v22
        default:
            if (!(4294967295 <= 4294967295)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    }
    if ((-1 != 127) || v2) {
        v23 := GenS3{F0: 267}
        var v24 uint16 = 1
        v3 = append(v3, 0)
        _ = v23
        _ = v24
    } else {
        fmt.Println(v2)
        if v2 {
            switch (("\n\t" + "hello") + ("\n\t" + "a")) {
            case "":
                if (!true) {
                    var v25 float32 = 0.1
                    _ = len(v3)
                    _ = v25
                }
            case "a":
                switch "a" {
                case "hello":
                    v26 := GenS2{F0: 4294967295, F1: "a"}
                    _ = v26
                case "foo bar":
                    for fi27 := 0; fi27 < 9; fi27++ {
                        fmt.Println(fi27)
                        v2 = true
                        _ = fi27
                    }
                case "a":
                    v28 := GenS3{F0: 42}
                    _ = v28
                default:
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                }
            default:
                var v29 int16 = 32767
                _ = v29
            }
            v2 = v2
        } else {
            v2 = (0 > -2147483648)
            if v2 {
                v2 = (v2 && (1 > -159))
                var v30 int8 = 42
                _ = v30
            }
        }
        var v31 float32 = 1e10
        _ = v31
    }
    v2 = v2
    var v32 int64 = 0
    if v2 {
        v2 = true
        v32 = 1
    }
    v2 = (!(371 < 32767))
    switch -222 {
    case 127:
        switch 2147483647 {
        case 32768:
            var v33 uint32 = 2147483647
            _ = v33
        }
    case -9223372036854775808:
        v2 = (1000 >= 1000)
    default:
        v34 := GenS4{F0: 990, F1: 2147483647, F2: 32767}
        _ = v34
    }
    v35 := GenS5{F0: -625, F1: 42, F2: 879}
    var v36 uint = 255
    var v37 int = -1
    for fi38 := 0; fi38 < 6; fi38++ {
        fmt.Println(v32, v36)
        _ = fi38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v37
}
