package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int16, p1 uint, p2 string) uint8 {
    var v7 uint16 = 32767
    if false {
        v7 -= (^(v7))
    }
    if true {
        v8 := GenS1{F0: 222, F1: 127}
        p0 += 42
        switch 42 {
        case -310:
            v9 := []int{128, 4294967295, 592, 4294967295}
            _ = v9
        case -480:
            for fi10 := 0; fi10 < 5; fi10++ {
                if true {
                    var v11 bool = (0 < 146)
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v11 = (v11 || true)
                        var v13 uint32 = 910
                        _ = fi12
                        _ = v13
                    }
                    p2 = ("hello" + ((("" + ("" + "a")) + p2) + p2))
                    _ = v11
                }
                var v14 uint8 = 195
                _ = fi10
                _ = v14
            }
        case -1:
            if (p1 >= (^(p1))) {
                v15 := GenS1{F0: 4294967295, F1: 128}
                var v16 bool = false
                var v17 int = 32768
                _ = v15
                _ = v16
                _ = v17
            }
        default:
            v18 := GenS2{F0: 538, F1: 1, F2: uint64(p1)}
            _ = v18
        }
        _ = v8
    }
    if (v7 != (v7 / ((v7) | 1))) {
        switch 735 {
        case 2147483647:
            var v19 int8 = 42
            _ = v19
        case 69:
            p2 = p2
        }
    }
    v20 := []int{1, -653}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v7
    return 47
}

func main() {
    var v1 float32 = 1e10
    var v2 uint = 65535
    var v3 bool = ((1.5 * 0.0) < (((3.14159 / 1.0) / 1e10) - 0.0))
    v4 := []int{}
    v5 := make(map[string]int)
    v6 := make(map[string]bool)
    if (42 != 42) {
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 int64 = 1
            for fi23 := 0; fi23 < 11; fi23++ {
                switch (("foo bar" + "foo bar") + "a") {
                case "a":
                    if true {
                        _ = len(v6)
                        _ = len(v4)
                    } else {
                        v3 = true
                        v3 = ((0 <= -116) && (0 < -670))
                    }
                case "":
                    v22 = 9223372036854775807
                default:
                    var v24 uint32 = 2147483647
                    _ = v24
                }
                _ = fi23
            }
            var v25 int8 = 64
            _ = fi21
            _ = v22
            _ = v25
        }
        if (v3 || false) {
            var v26 uint = (v2 - v2)
            v3 = ((((("hello" + "\n\t") + "a") + "hello") != "a") || false)
            _ = v26
        }
        v2 = (^(v2))
    }
    if true {
        if v3 {
            if len(v4) > 0 {
                v4[0] = 297
            }
            _ = len(v4)
            v27 := GenS3{F0: v3, F1: 42, F2: 255}
            _ = v27
        } else {
            v2 = (v2 % ((v2) | 1))
            if (!(666 <= 32767)) {
                switch "hello" {
                case "foo bar":
                    if true {
                        fmt.Println(v1, v2)
                        _ = len(v5)
                        v28 := GenS3{F0: (42 > 65535), F1: 40, F2: 55}
                        _ = v28
                    }
                case "a":
                    fmt.Println(v2, v1, v1)
                }
                var v29 float64 = 1e10
                _ = v29
            } else {
                var v30 uint = 4294967295
                _ = v30
            }
        }
        switch -411 {
        case 127:
            v2 -= (^(v2))
        }
        v31 := GenS4{F0: 65535}
        _ = v31
    } else {
        v32 := GenS4{F0: 596}
        v33 := GenS4{F0: 2147483648}
        _ = v32
        _ = v33
    }
    if (-991 != 0) {
        if (!(((v3 && (!false)) && v3) && (false || v3))) {
            fmt.Println(v3, v3, v3)
            v34 := GenS4{F0: 42}
            switch "" {
            case "a":
                fmt.Println(v3)
            default:
                var v35 uint = v2
                _ = v35
            }
            _ = v34
        }
        v1 = (float32(0.0) + 1e10)
    } else {
        fmt.Println(v2, v2, v3)
        var v36 float32 = v1
        v4 = append(v4, 2147483647)
        _ = v36
    }
    var v37 int32 = -1
    v38 := GenS5{F0: 1}
    if (!(42 >= 56)) {
        v1 = (v1 * v1)
        v6[""] = v3
        v2 = v2
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v39 float64 = -2.25
    var v40 uint64 = 42
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v41 int64 = 353
    var v42 float64 = v39
    switch -369 {
    case 32767:
        v43 := GenS5{F0: (v41 / ((v41) | 1))}
        _ = v43
    case 1:
        var v44 string = ("a" + "\n\t")
        _ = v44
    case 127:
        fmt.Println(v3, v40)
    }
    v39 = 0.0
    fmt.Println(v37)
    v39 = v39
    for fi45 := 0; fi45 < 2; fi45++ {
        if len(v4) > 0 {
            v4[0] = (fi45 >> ((v2) & 63))
        }
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
