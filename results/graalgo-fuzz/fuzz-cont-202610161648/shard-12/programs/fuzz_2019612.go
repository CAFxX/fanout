package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int16 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint, p1 uint32) bool {
    var v6 uint = p0
    if (!false) {
        v6 += 1
        var v7 float32 = 1e10
        _ = v7
    }
    var v8 int64 = -297
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return false
}

func main() {
    v1 := []int{32768, 2147483647, -330, 4294967295}
    var v2 int = 32768
    var v3 uint16 = 0
    v4 := make(map[uint32]string)
    var v5 bool = true
    v5 = (v5 && (v2 > 32768))
    v5 = (v5 && v5)
    fmt.Println(v5)
    v9 := GenS1{F0: "hello", F1: 0, F2: 42}
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 string = "\n\t"
        delete(v4, 2147483647)
        v5 = (int8(v3) != -128)
        _ = fi10
        _ = v11
    }
    switch (v2 + (v2 * 32767)) {
    case 9223372036854775807:
        switch -226 {
        case 986:
            var v12 uint64 = 870
            _ = v12
        case 9223372036854775807:
            var v13 uint16 = v3
            _ = v13
        }
    case -1:
        if v5 {
            switch ("\n\t" + (("" + "hello") + "hello")) {
            case "":
                if ("a" < "") {
                    for fi14 := 0; fi14 < 17; fi14++ {
                        fmt.Println(v3)
                        switch "foo bar" {
                        case "\n\t":
                            delete(v4, 65535)
                        }
                        switch (fi14 + 32768) {
                        case 32768:
                            v15 := GenS2{F0: ("" + "a"), F1: 300}
                            _ = v15
                        case -2147483648:
                            v16 := GenS3{F0: -1, F1: -944}
                            _ = v16
                        case 9223372036854775807:
                            v17 := GenS4{F0: (!(v5 && v5)), F1: (v3 * (^(v3)))}
                            _ = v17
                        default:
                            var v18 int16 = 1
                            _ = v18
                        }
                        _ = fi14
                    }
                    var v19 uint8 = 127
                    _ = v19
                } else {
                    v2 = v2
                    v5 = v5
                }
            case "a":
                if (true || (!true)) {
                    fmt.Println(v5, v5, v5)
                } else {
                    if (!(0.1 <= 0.0)) {
                        var v20 int = 127
                        v21 := GenS4{F0: v5, F1: v3}
                        if true {
                            v2 = v20
                            for fi22 := 0; fi22 < 6; fi22++ {
                                v5 = false
                                for fi23 := 0; fi23 < 1; fi23++ {
                                    var v24 uint64 = 2
                                    if v5 {
                                        v5 = true
                                        v20 = (fi23 ^ 42)
                                        v25 := GenS2{F0: "", F1: fi22}
                                        _ = v25
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                                _ = fi22
                            }
                            v2 -= 128
                        } else {
                            fmt.Println(v5, v5)
                        }
                        _ = v20
                        _ = v21
                    }
                }
            case "hello":
                v3 = (^(v3))
            }
            var v26 float64 = (0.1 / 1.0)
            _ = v26
        } else {
            var v27 string = "a"
            var v28 string = (v27 + v27)
            _ = v27
            _ = v28
        }
    default:
        v1 = append(v1, v2)
    }
    var v29 int64 = 9223372036854775807
    switch v2 {
    case 219:
        var v30 int8 = 68
        _ = v30
    case -2147483648:
        if (!true) {
            v3 -= v3
        }
    case 4294967295:
        v2 = 32767
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
