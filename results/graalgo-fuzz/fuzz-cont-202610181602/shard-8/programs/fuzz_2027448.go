package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int8; F2 string }
type GenS4 struct { F0 int8; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint64 }
func genFn0() int64 {
    var v6 int = 2147483647
    v6 = (v6 + -9223372036854775808)
    var v7 uint8 = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        switch "\n\t" {
        case "":
            for fi9 := 0; fi9 < 4; fi9++ {
                if false {
                    v10 := []int{}
                    fmt.Println(v7, fi9)
                    var v11 uint32 = 2
                    if (174 > 2147483647) {
                        fmt.Println(fi8, fi9)
                    } else {
                        v12 := GenS1{F0: 32767}
                        switch 301 {
                        case 1:
                            var v13 int8 = 0
                            _ = v13
                        default:
                            v14 := GenS2{F0: 18446744073709551615}
                            _ = v14
                        }
                        _ = v12
                    }
                    switch fi8 {
                    case 9223372036854775807:
                        if true {
                            fmt.Println(v6)
                            v6 += fi8
                            var v15 bool = true
                            _ = v15
                        } else {
                            v16 := GenS1{F0: 32767}
                            var v17 uint64 = 42
                            v7 = v7
                            var v18 uint64 = (v17 / ((4294967295) | 1))
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    }
                    _ = v10
                    _ = v11
                } else {
                    v19 := []int{(-(fi9))}
                    v7 = v7
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    _ = v19
                }
                if (!((3.14159 * (0.0 * (((0.0 * 0.1) + ((3.14159 - 1e10) / 3.14159)) - 1e10))) > 3.14159)) {
                    fmt.Println(fi8)
                    var v20 float32 = (3.14159 / 1e10)
                    _ = v20
                }
                if (true || false) {
                    v21 := []int{}
                    _ = v21
                } else {
                    fmt.Println(v7, v7, v6)
                    v22 := GenS3{F0: 1, F1: 100, F2: (("" + "hello") + "hello")}
                    _ = v22
                }
                v6 -= 265
                v7 = (^(v7))
                _ = fi9
            }
        }
        _ = fi8
    }
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 int16 = 44
    v4 := []string{"hello"}
    var v5 string = "foo bar"
    var v23 uint64 = 708
    v3 += (v3 - v3)
    v23 = v23
    v24 := v2[32767]
    for fi25 := 0; fi25 < 19; fi25++ {
        v23 += v23
        _ = fi25
    }
    if (65535 <= 2147483648) {
        fmt.Println(v24, v5, v5)
    }
    v23 = (v23 + (v23 | (^(v23))))
    v23 += v23
    v26 := GenS2{F0: 255}
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 uint = 764
        delete(v2, fi27)
        _ = fi27
        _ = v28
    }
    v29 := GenS2{F0: 65535}
    v24 = v24
    if true {
        v23 = (v23 / ((42) | 1))
    }
    if true {
        var v30 int16 = (v3 & (v3 << ((284) & 63)))
        switch (v24 & v24) {
        case 1:
            v23 += (^(v23))
        default:
            if true {
                switch v5 {
                case "foo bar":
                    v2[32767] = int(v23)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        _ = len(v1)
        _ = v30
    } else {
        var v31 int = (v24 << ((255) & 63))
        v32 := GenS1{F0: 42}
        _ = v31
        _ = v32
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3, v24)
    v24 -= (v24 >> ((4294967295) & 63))
    for fi33 := 0; fi33 < 1; fi33++ {
        v34 := GenS4{F0: -1, F1: -800, F2: 348}
        v35 := GenS5{F0: 255, F1: 139}
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
