package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{0, 628, 2147483647}
    var v3 bool = true
    v4 := make(map[int]uint32)
    v5 := make(map[int]int)
    var v6 int64 = -130
    var v7 bool = v3
    if v7 {
        v8 := GenS1{F0: 961, F1: "hello"}
        _ = v8
    } else {
        v7 = (!v3)
        if v3 {
            v7 = false
            v3 = v7
            switch -1 {
            case 0:
                for fi9 := 0; fi9 < 15; fi9++ {
                    v10 := GenS1{F0: 482, F1: "foo bar"}
                    var v11 uint32 = 1
                    if (225 != 114) {
                        fmt.Println(v6)
                        var v12 int = -9223372036854775808
                        _ = v12
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case -403:
                var v13 float32 = ((-2.25 / 1.5) - (1.5 + 0.1))
                _ = v13
            case 127:
                var v14 bool = false
                _ = v14
            default:
                v15 := GenS1{F0: 106, F1: "foo bar"}
                _ = v15
            }
        }
        switch -1 {
        case 245:
            switch "foo bar" {
            case "":
                if (255 != 255) {
                    fmt.Println(v3, v6, v3)
                    v16 := GenS1{F0: 2147483648, F1: ("a" + ((((("" + "a") + "foo bar") + "a") + "\n\t") + "hello"))}
                    v17 := GenS1{F0: 65535, F1: "a"}
                    _ = v16
                    _ = v17
                }
            case "a":
                v18 := GenS2{F0: 0}
                _ = v18
            default:
                var v19 int = -2147483648
                _ = v19
            }
        default:
            v3 = v7
        }
    }
    v3 = (80 < 53)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v7 = (2 <= 2)
    for fi20 := 0; fi20 < 6; fi20++ {
        switch (("" + "a") + "a") {
        case "\n\t":
            if v3 {
                v21 := GenS2{F0: 4294967295}
                _ = v21
            }
        }
        fmt.Println(v7, fi20, v3)
        v6 = 544
        _ = fi20
    }
    v7 = (3.14159 != (-2.25 * 1.5))
    _ = len(v1)
    var v22 bool = (v3 && false)
    switch 32768 {
    case 32768:
        var v23 int16 = -1
        _ = v23
    case 0:
        for fi24 := 0; fi24 < 2; fi24++ {
            if v7 {
                v25 := GenS3{F0: 42}
                for fi26 := 0; fi26 < 7; fi26++ {
                    var v27 float64 = (1.5 - (3.14159 * -2.25))
                    v28 := v5[(fi26 | (fi24 % (((fi26 & fi26)) | 1)))]
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = v25
            }
            _ = fi24
        }
    case 9223372036854775807:
        for fi29 := 0; fi29 < 16; fi29++ {
            var v30 int32 = 42
            if (!false) {
                fmt.Println(v22, v3)
                var v31 int8 = -1
                _ = v31
            }
            for fi32 := 0; fi32 < 11; fi32++ {
                fmt.Println(v6, v6, v7)
                v22 = (-1 != -1)
                _ = fi32
            }
            _ = fi29
            _ = v30
        }
    }
    if v22 {
        for fi33 := 0; fi33 < 2; fi33++ {
            if false {
                var v34 float32 = 1e10
                _ = v34
            }
            var v35 uint64 = 496
            _ = fi33
            _ = v35
        }
        v22 = v3
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        fmt.Println(fi36)
        if (32767 != 478) {
            fmt.Println(v7, v7)
            var v37 string = "\n\t"
            var v38 float32 = 0.1
            _ = v37
            _ = v38
        }
        v39 := GenS2{F0: 18446744073709551615}
        _ = fi36
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
