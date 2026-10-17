package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string }
func genFn0() int {
    switch ("\n\t" + "hello") {
    case "":
        v5 := GenS1{F0: 1, F1: -1, F2: 42}
        _ = v5
    case "\n\t":
        var v6 uint8 = 2
        _ = v6
    case "hello":
        var v7 string = "\n\t"
        _ = v7
    default:
        var v8 int = 581
        _ = v8
    }
    switch -243 {
    case 9223372036854775807:
        var v9 bool = (3.14159 == (1e10 * (3.14159 / ((3.14159 - 3.14159) + 0.1))))
        _ = v9
    case 127:
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(fi10)
            _ = fi10
        }
    case 42:
        if false {
            var v11 uint16 = 255
            switch 9223372036854775807 {
            case 2147483647:
                var v12 int = 380
                _ = v12
            case 127:
                v13 := GenS2{F0: 127, F1: true}
                _ = v13
            default:
                v14 := []int{-2147483648, 127}
                _ = v14
            }
            _ = v11
        } else {
            var v15 uint8 = 2
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 uint64 = 943
                var v18 int16 = -1
                _ = fi16
                _ = v17
                _ = v18
            }
            if (0.0 < float64(0.0)) {
                v15 = 123
            }
            _ = v15
        }
    }
    var v19 float32 = ((1e10 + 0.0) + 3.14159)
    if (-9223372036854775808 == -9223372036854775808) {
        switch 32768 {
        case 128:
            for fi20 := 0; fi20 < 20; fi20++ {
                v21 := []int{}
                v19 += (v19 - (v19 * 1.5))
                _ = fi20
                _ = v21
            }
        case 32767:
            v22 := []int{}
            _ = v22
        case -787:
            v23 := GenS1{F0: 152, F1: 1, F2: 18446744073709551615}
            _ = v23
        }
        var v24 uint8 = 255
        _ = v24
    } else {
        if (42 < 2) {
            for fi25 := 0; fi25 < 16; fi25++ {
                var v26 float64 = 3.14159
                if false {
                    if true {
                        v27 := GenS3{F0: 933}
                        if (42 < 56) {
                            var v28 string = "hello"
                            fmt.Println(v28, fi25, v28)
                            _ = v28
                        }
                        if true {
                            v19 += float32(v26)
                        } else {
                            var v29 uint32 = 1
                            var v30 uint8 = 134
                            _ = v29
                            _ = v30
                        }
                        _ = v27
                    }
                }
                switch "foo bar" {
                case "a":
                    if (!((v26 / 1e10) <= 0.0)) {
                        v26 = (1.5 * v26)
                        if (true || true) {
                            v31 := GenS4{F0: 2147483647, F1: "a"}
                            if (66 == 22) {
                                v19 += v19
                            } else {
                                switch fi25 {
                                case -512:
                                    v19 += float32(v26)
                                }
                                var v32 uint = 429
                                v26 += v26
                                _ = v32
                            }
                            v26 = (v26 * 1.5)
                            _ = v31
                        }
                        fmt.Println(v19, fi25, fi25)
                    }
                }
                _ = fi25
                _ = v26
            }
        } else {
            v19 = v19
            var v33 uint64 = 528
            _ = v33
        }
    }
    v19 += v19
    _ = v19
    return 2147483647
}
func genFn1(p0 uint) uint8 {
    v34 := GenS4{F0: 42, F1: "a"}
    p0 = p0
    fmt.Println(p0)
    p0 = p0
    _ = p0
    _ = v34
    return 72
}

func main() {
    var v1 uint64 = 1
    var v2 int32 = 1
    v3 := make(map[int]uint32)
    v4 := []int{-9223372036854775808, 32768, 1}
    if (!true) {
        v35 := GenS5{F0: ((1 == int16(v1)) || true)}
        switch "" {
        case "a":
            v2 += (v2 % ((42) | 1))
        }
        if (((2 > 255) && (v2 != v2)) && (32767 >= 1000)) {
            fmt.Println(v2, v2)
        }
        _ = v35
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        if false {
            v1 += v1
        }
        _ = fi36
    }
    if false {
        var v37 int16 = 1000
        var v38 float32 = 0.0
        _ = v37
        _ = v38
    }
    switch -2147483648 {
    case 1:
        for fi39 := 0; fi39 < 15; fi39++ {
            switch "a" {
            case "hello":
                var v40 int = (fi39 >> ((2) & 63))
                _ = v40
            case "foo bar":
                fmt.Println(v2, v2, v2)
            default:
                var v41 float64 = ((((-2.25 + ((1e10 + 3.14159) - 1.5)) / 0.1) - ((0.1 - 1e10) / 1.5)) / 1.0)
                _ = v41
            }
            v3[genFn0()] = 65535
            v1 += v1
            _ = fi39
        }
    }
    v42 := GenS2{F0: 42, F1: false}
    if true {
        v2 -= (v2 / (((v2 - 42)) | 1))
        v43 := GenS2{F0: -507, F1: (!(!(true || (((3.14159 * 1e10) != ((1.5 * (1e10 + 3.14159)) + 1.5)) || false))))}
        _ = v43
    }
    if false {
        v44 := GenS6{F0: ""}
        var v45 float32 = 0.0
        var v46 uint8 = uint8(v1)
        _ = v44
        _ = v45
        _ = v46
    } else {
        fmt.Println(v2)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
