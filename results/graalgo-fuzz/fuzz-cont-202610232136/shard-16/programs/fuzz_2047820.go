package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 string, p1 uint8, p2 int) int {
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 100, F1: 42}
        v9 := GenS2{F0: -2147483648, F1: p1, F2: 8}
        if false {
            fmt.Println(p2, p1, p0)
            v10 := GenS3{F0: -614}
            var v11 uint64 = 7
            _ = v10
            _ = v11
        } else {
            switch (fi7 * (p2 + p2)) {
            case -964:
                v12 := GenS1{F0: 1, F1: 255}
                _ = v12
            }
            if (-128 >= 127) {
                p2 = (p2 - -529)
                var v13 int = p2
                v14 := GenS3{F0: 0}
                _ = v13
                _ = v14
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 9223372036854775807 {
    case -1:
        p0 = p0
    case -9223372036854775808:
        v15 := GenS3{F0: 0}
        _ = v15
    }
    v16 := GenS1{F0: 1, F1: uint16(p1)}
    v17 := GenS3{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return p2
}
func genFn1() uint16 {
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 uint32 = 4294967295
        v20 := GenS3{F0: 733}
        var v21 uint64 = 148
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    if false {
        var v22 string = "a"
        switch "a" {
        case "a":
            v22 = v22
        case "foo bar":
            v23 := GenS2{F0: -466, F1: 127, F2: 2}
            _ = v23
        case "":
            v22 = v22
        }
        fmt.Println(v22)
        _ = v22
    } else {
        var v24 float64 = 1.5
        for fi25 := 0; fi25 < 1; fi25++ {
            v24 += v24
            v24 -= (((v24 - float64(1e10)) / 1.0) * float64(1.5))
            var v26 uint32 = 42
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    var v27 uint16 = 42
    _ = v27
    return v27
}

func main() {
    var v1 int16 = 0
    var v2 float32 = 0.1
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]bool)
    var v5 uint8 = 79
    v6 := []string{"a", "foo bar", "\n\t"}
    fmt.Println(v5, v1, v2)
    v28 := GenS4{F0: 4294967295}
    var v29 uint8 = 223
    v30 := GenS3{F0: (v3 ^ v3)}
    for fi31 := 0; fi31 < 13; fi31++ {
        v1 = v1
        if len(v6) > 0 {
            v6[0] = ""
        }
        _ = fi31
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch 910 {
    case 32767:
        var v32 uint = 255
        _ = v32
    case 32768:
        switch 42 {
        case -1:
            if (42 <= 215) {
                var v33 bool = (false || false)
                _ = v33
            }
        case 128:
            switch -2147483648 {
            case 32767:
                v34 := GenS3{F0: 0}
                _ = v34
            case 1:
                v5 = v5
            case 545:
                var v35 float32 = float32(0.0)
                _ = v35
            }
        case -2147483648:
            delete(v4, 42)
        }
    default:
        v36 := GenS4{F0: 42}
        _ = v36
    }
    fmt.Println(v29, v1, v3)
    v3 = v3
    _ = len(v6)
    var v37 int32 = 767
    for fi38 := 0; fi38 < 19; fi38++ {
        v1 = -1
        _ = fi38
    }
    v39 := v4[genFn0("\n\t", (^(v5)), 127)]
    var v40 uint16 = 2
    v40 -= (v40 | v40)
    v1 -= (v1 - (v1 - (v1 ^ (v1 + v1))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
