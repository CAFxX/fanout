package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 0.0
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]string)
    var v5 uint = 938
    fmt.Println(v2, v3, v2)
    var v6 int = 32767
    switch "" {
    case "a":
        var v7 uint = (v5 + 0)
        _ = v7
    case "hello":
        delete(v4, 2147483647)
    }
    var v8 uint32 = 486
    var v9 uint16 = 255
    v10 := []int{1, v6, (-(v6)), (v6 + v6)}
    v11 := GenS1{F0: true}
    var v12 bool = (-18 != 846)
    var v13 int64 = int64(v5)
    if (0.0 == 1.5) {
        if v12 {
            var v14 int8 = -10
            v15 := GenS2{F0: (v8 ^ v8), F1: (v9 ^ (^(v9)))}
            _ = v14
            _ = v15
        } else {
            var v16 uint32 = (v8 & 1)
            switch "a" {
            case "":
                var v17 bool = true
                _ = v17
            case "\n\t":
                if false {
                    switch (v6 - 1) {
                    case 0:
                        var v18 int = v6
                        _ = v18
                    case 923:
                        if true {
                            delete(v1, "a")
                            var v19 float32 = -2.25
                            _ = v19
                        } else {
                            if (v3 != v3) {
                                switch v6 {
                                case -20:
                                    var v20 uint32 = 4294967295
                                    _ = v20
                                default:
                                    var v21 uint64 = (^(v3))
                                    _ = v21
                                }
                                if (((1e10 * float64(0.1)) * 3.14159) != 3.14159) {
                                    var v22 float32 = 0.0
                                    var v23 uint16 = (v9 | (v9 / ((v9) | 1)))
                                    var v24 int16 = int16(v5)
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                            } else {
                                delete(v1, "a")
                                for fi25 := 0; fi25 < 1; fi25++ {
                                    v26 := GenS1{F0: (!false)}
                                    for fi27 := 0; fi27 < 17; fi27++ {
                                        v12 = v12
                                        switch ("\n\t" + "foo bar") {
                                        case "\n\t":
                                            if v12 {
                                                v16 -= v16
                                                _ = len(v10)
                                                switch fi25 {
                                                case 128:
                                                    if v12 {
                                                        fmt.Println(v3, v3, v12)
                                                        v28 := GenS2{F0: v8, F1: v9}
                                                        _ = v28
                                                    }
                                                case 764:
                                                    if true {
                                                        var v29 uint32 = (v16 % ((v8) | 1))
                                                        _ = v29
                                                    } else {
                                                        v9 = v9
                                                        var v30 string = (((("a" + "hello") + ("a" + ("a" + "a"))) + "a") + ("hello" + "hello"))
                                                        _ = v30
                                                    }
                                                case 127:
                                                    switch ("a" + "hello") {
                                                    case "hello":
                                                        var v31 float32 = 0.0
                                                        _ = v31
                                                    default:
                                                        switch (-(fi25)) {
                                                        case 0:
                                                            v12 = true
                                                        case -2147483648:
                                                            var v32 int = fi27
                                                            _ = v32
                                                        case 128:
                                                            fmt.Println(v8, fi27)
                                                        }
                                                    }
                                                default:
                                                    v3 -= 42
                                                }
                                            } else {
                                                var v33 int64 = (v13 >> (((v5 << ((42) & 63))) & 63))
                                                _ = v33
                                            }
                                        default:
                                            switch fi25 {
                                            case 0:
                                                switch -58 {
                                                case 560:
                                                    if v12 {
                                                        v9 += v9
                                                        v2 = -2.25
                                                    } else {
                                                        var v34 bool = ((v2 + (v2 * v2)) <= v2)
                                                        if len(v10) >= 2 {
                                                            v10 = v10[1:len(v10)]
                                                        }
                                                        var v35 bool = true
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                case 42:
                                                    switch -9223372036854775808 {
                                                    case 2147483647:
                                                        v36 := GenS3{F0: 127, F1: 42}
                                                        _ = v36
                                                    case 32767:
                                                        v16 = 750
                                                        var v37 int64 = (^(v13))
                                                        _ = v37
                                                    }
                                                case 665:
                                                    v16 += (v8 | v16)
                                                default:
                                                    if len(v10) > 0 {
                                                        v10[0] = fi25
                                                    }
                                                }
                                            default:
                                                fmt.Println(v13)
                                            }
                                        }
                                        _ = fi27
                                    }
                                    _ = fi25
                                    _ = v26
                                }
                                var v38 uint32 = 2147483647
                                _ = v38
                            }
                        }
                    default:
                        switch (v6 ^ v6) {
                        case 547:
                            v12 = v12
                        case -2147483648:
                            v6 = (v6 + (v6 | v6))
                        }
                    }
                } else {
                    v5 = v5
                    v10 = append(v10, v6)
                }
            case "hello":
                switch (^(v6)) {
                case 299:
                    switch v6 {
                    case 9223372036854775807:
                        v16 = (^(v16))
                    case 127:
                        var v39 uint64 = (v3 << ((217) & 63))
                        _ = v39
                    default:
                        switch (("a" + ("foo bar" + ("" + "hello"))) + "foo bar") {
                        case "a":
                            v16 = (v16 * (^(v16)))
                        }
                    }
                default:
                    v6 = (v6 >> (((v5 * v5)) & 63))
                }
            }
            v13 -= (-(v13))
            _ = v16
        }
        var v40 int16 = -926
        if false {
            fmt.Println(v3, v12, v8)
            fmt.Println(v8, v12)
            switch "a" {
            case "foo bar":
                v8 = v8
            }
        }
        _ = v40
    } else {
        var v41 int8 = -1
        _ = v41
    }
    var v42 uint16 = 65535
    if len(v10) > 0 {
        v10[0] = -9223372036854775808
    }
    var v43 uint16 = v42
    if (32767 != (v42 ^ uint16(v8))) {
        switch "hello" {
        case "hello":
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        }
    } else {
        var v44 uint = uint(v3)
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
