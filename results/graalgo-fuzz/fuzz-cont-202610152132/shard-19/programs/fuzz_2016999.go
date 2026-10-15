package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 string }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]bool)
    var v3 uint8 = 42
    v4 := GenS1{F0: (!true), F1: -528, F2: ("hello" + "foo bar")}
    for fi5 := 0; fi5 < 17; fi5++ {
        if ((!(("foo bar" + "\n\t") != "a")) && true) {
            v6 := GenS1{F0: (-1 != 42), F1: 9223372036854775807, F2: ""}
            _ = v6
        }
        var v7 float32 = 1e10
        _ = fi5
        _ = v7
    }
    v3 += (v3 & v3)
    v8 := []int{32767, -1}
    _ = len(v2)
    v3 = 255
    if true {
        v9 := GenS2{F0: 128, F1: 127}
        _ = v9
    }
    var v10 int32 = 1
    v1 = ((v3 | (v3 | (v3 | v3))) < v3)
    var v11 uint16 = 42
    for fi12 := 0; fi12 < 3; fi12++ {
        v8 = append(v8, 32768)
        fmt.Println(v10, v10, v11)
        _ = fi12
    }
    if (!(-2.25 == 1.5)) {
        switch "hello" {
        case "\n\t":
            v13 := v2[42]
            _ = v13
        case "a":
            switch (("hello" + "\n\t") + "a") {
            case "foo bar":
                if v1 {
                    var v14 bool = (v1 && v1)
                    _ = v14
                }
            case "a":
                if v1 {
                    v15 := GenS3{F0: v10, F1: ("\n\t" + ("\n\t" + ""))}
                    _ = len(v8)
                    _ = v15
                } else {
                    if (false && (v1 || ((v10 * (v10 - v10)) == (^(v10))))) {
                        var v16 float64 = 1.5
                        v11 = (v11 >> ((18446744073709551615) & 63))
                        fmt.Println(v3)
                        _ = v16
                    }
                    v3 = 128
                    if false {
                        delete(v2, 4294967295)
                    } else {
                        v11 = v11
                        v17 := GenS4{F0: 262}
                        _ = v17
                    }
                }
            case "\n\t":
                v8 = append(v8, 238)
            }
        }
    }
    switch 4294967295 {
    case 127:
        v1 = false
    case 42:
        for fi18 := 0; fi18 < 6; fi18++ {
            v11 = 90
            _ = fi18
        }
    case 32768:
        v1 = (false || (42 == 900))
    default:
        v1 = (42 != 0)
    }
    if ((42 < -1) || true) {
        fmt.Println(v1, v3, v1)
        fmt.Println(v3)
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 int8 = 100
        _ = fi19
        _ = v20
    }
    switch 32767 {
    case 222:
        if v1 {
            fmt.Println(v3, v10, v3)
            var v21 int16 = 1000
            _ = len(v2)
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        switch "" {
        case "":
            for fi23 := 0; fi23 < 6; fi23++ {
                v2[0] = v1
                v1 = true
                _ = fi23
            }
        case "\n\t":
            switch fi22 {
            case -512:
                switch (fi22 / (((fi22 / (((fi22 & fi22)) | 1))) | 1)) {
                case -714:
                    for fi24 := 0; fi24 < 15; fi24++ {
                        v1 = false
                        if (v3 <= v3) {
                            if false {
                                if ((true || false) || v1) {
                                    fmt.Println(v1)
                                } else {
                                    v11 = 32767
                                    var v25 uint = 955
                                    var v26 int64 = 1
                                    _ = v25
                                    _ = v26
                                }
                                fmt.Println(v1, v3)
                            }
                            var v27 float32 = (1.5 + (3.14159 / ((3.14159 / 0.1) - ((3.14159 * -2.25) / float32(1e10)))))
                            fmt.Println(fi24, fi22, v10)
                            _ = v27
                        }
                        _ = fi24
                    }
                case 127:
                    if len(v8) > 0 {
                        v8[0] = (fi22 + fi22)
                    }
                }
            default:
                _ = len(v8)
            }
        case "foo bar":
            switch (fi22 & fi22) {
            case 4294967295:
                var v28 float32 = ((((0.1 / 1.5) * (0.0 - -2.25)) * 0.1) * float32(3.14159))
                _ = v28
            case 950:
                var v29 uint64 = 513
                _ = v29
            }
        default:
            v11 = v11
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
