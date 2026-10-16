package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int16; F1 uint; F2 int32 }
type GenS4 struct { F0 bool; F1 int; F2 uint8 }
func genFn0(p0 uint64, p1 uint32) int {
    if false {
        p1 = p1
    }
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 int8 = -1
        if (v5 != v5) {
            p0 += (p0 * (p0 | 2))
            v6 := GenS1{F0: true, F1: 0, F2: 42}
            _ = v6
        } else {
            v5 += (v5 & 100)
        }
        _ = fi4
        _ = v5
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return -974
}
func genFn1(p0 bool, p1 int32) string {
    v7 := []int{}
    switch "a" {
    case "hello":
        v8 := GenS2{F0: 42, F1: 2147483647}
        _ = v8
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return "a"
}

func main() {
    var v1 int64 = 213
    v2 := []string{"a", "\n\t"}
    var v3 uint8 = 29
    var v9 float32 = 1e10
    v1 = v1
    v9 -= v9
    var v10 bool = true
    if (("" <= "hello") || (42 >= 42)) {
        switch -9223372036854775808 {
        case 32768:
            var v11 uint8 = v3
            _ = v11
        }
        var v12 int32 = 610
        var v13 int32 = (v12 | v12)
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 17; fi14++ {
            if ((1 > 18446744073709551615) || ((v9 > float32((-2.25 + -2.25))) || v10)) {
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 int16 = 317
                    if (-60 <= -128) {
                        v17 := GenS2{F0: 0, F1: 128}
                        switch "" {
                        case "hello":
                            switch fi15 {
                            case 9223372036854775807:
                                v18 := GenS3{F0: 32767, F1: 1, F2: 1}
                                _ = v18
                            default:
                                fmt.Println(fi14)
                            }
                        }
                        v1 -= (v1 & v1)
                        _ = v17
                    } else {
                        var v19 int = fi14
                        _ = v19
                    }
                    _ = fi15
                    _ = v16
                }
                for fi20 := 0; fi20 < 7; fi20++ {
                    switch "foo bar" {
                    case "hello":
                        v9 += (3.14159 + v9)
                    case "a":
                        if (-88 > -128) {
                            if (v9 <= v9) {
                                v3 = v3
                            }
                            var v21 int16 = -1
                            _ = v21
                        }
                        var v22 string = "a"
                        _ = v22
                    default:
                        var v23 int8 = 43
                        _ = v23
                    }
                    _ = fi20
                }
            } else {
                fmt.Println(v9)
                var v24 uint16 = 65535
                _ = v24
            }
            for fi25 := 0; fi25 < 18; fi25++ {
                var v26 bool = false
                _ = fi25
                _ = v26
            }
            for fi27 := 0; fi27 < 8; fi27++ {
                v3 = (v3 & v3)
                if (157 != 599) {
                    switch fi14 {
                    case 127:
                        if true {
                            var v28 int64 = v1
                            v9 = (v9 / (v9 * (1.5 / v9)))
                            _ = v28
                        } else {
                            v10 = v10
                            var v29 int64 = v1
                            v3 = v3
                            _ = v29
                        }
                    }
                }
                var v30 bool = (((("hello" + "a") + "foo bar") + ("a" + "")) < "a")
                _ = fi27
                _ = v30
            }
            _ = fi14
        }
    }
    v31 := GenS3{F0: 1, F1: 360, F2: 42}
    var v32 int8 = 127
    for fi33 := 0; fi33 < 11; fi33++ {
        v34 := GenS4{F0: (-1 == 891), F1: fi33, F2: v3}
        _ = fi33
        _ = v34
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v10 {
        v9 += ((0.0 * (v9 - float32(0.1))) + v9)
    }
    var v35 uint64 = 139
    var v36 string = ""
    v9 -= (v9 * v9)
    for fi37 := 0; fi37 < 9; fi37++ {
        switch v36 {
        case "a":
            for fi38 := 0; fi38 < 16; fi38++ {
                v39 := GenS4{F0: false, F1: fi37, F2: (v3 | 42)}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi38
                _ = v39
            }
        default:
            v40 := GenS1{F0: v10, F1: v1, F2: 255}
            _ = v40
        }
        var v41 int32 = 235
        var v42 bool = (42 >= 1)
        _ = fi37
        _ = v41
        _ = v42
    }
    fmt.Println(v1)
    fmt.Println(v9, v32, v3)
    var v43 int16 = 1
    fmt.Println(v35)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v43:", v43)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v43
    _ = v9
}
