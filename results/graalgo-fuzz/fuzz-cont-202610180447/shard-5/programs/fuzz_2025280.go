package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int32 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0() int16 {
    switch -1 {
    case -438:
        var v6 int32 = 42
        _ = v6
    default:
        var v7 uint8 = 1
        _ = v7
    }
    var v8 uint8 = 128
    if ((2147483647 != -1) && true) {
        fmt.Println(v8)
    } else {
        if true {
            var v9 bool = (255 == 1)
            v10 := GenS1{F0: 1, F1: -1, F2: 202}
            v9 = (!(32768 <= 128))
            _ = v10
            _ = v9
        }
    }
    if (554 >= 32767) {
        var v11 string = (("" + "a") + (("" + ("foo bar" + "\n\t")) + "hello"))
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(fi12, fi12, v8)
            if true {
                var v13 uint = 687
                _ = v13
            } else {
                var v14 uint = 352
                _ = v14
            }
            _ = fi12
        }
        _ = v11
    } else {
        fmt.Println(v8)
    }
    _ = v8
    return 1000
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{2147483648, 65535, 42}
    var v3 int64 = 9223372036854775807
    var v4 int = 32767
    v5 := []uint32{2147483648, 65535, 1}
    v4 += (v4 | v4)
    v2 = append(v2, 2)
    var v15 uint = 639
    fmt.Println(v4, v4, v4)
    v16 := GenS1{F0: 0, F1: 127, F2: -1}
    var v17 int8 = -108
    fmt.Println(v3, v17)
    if (false && false) {
        fmt.Println(v15)
    }
    v18 := GenS1{F0: (v3 + v3), F1: (v17 ^ (v17 & v17)), F2: 0}
    switch "hello" {
    case "hello":
        if (!false) {
            var v19 uint64 = 2
            for fi20 := 0; fi20 < 17; fi20++ {
                switch 42 {
                case 0:
                    var v21 int32 = -2147483648
                    _ = v21
                case -137:
                    for fi22 := 0; fi22 < 10; fi22++ {
                        delete(v1, 4294967295)
                        var v23 uint32 = 993
                        _ = fi22
                        _ = v23
                    }
                case 9223372036854775807:
                    var v24 float32 = float32(((0.0 * float64(float32(1.5))) / 3.14159))
                    _ = v24
                }
                v25 := GenS2{F0: 42, F1: "foo bar"}
                _ = fi20
                _ = v25
            }
            _ = v19
        } else {
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            delete(v1, 129)
            switch ("foo bar" + ("" + "foo bar")) {
            case "\n\t":
                if (false || true) {
                    var v26 uint32 = 42
                    fmt.Println(v15, v17)
                    _ = v26
                }
            case "a":
                if true {
                    var v27 uint32 = 2147483647
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 7; fi28++ {
                        var v29 bool = true
                        _ = fi28
                        _ = v29
                    }
                    if (true && (42 != 0)) {
                        v30 := GenS3{F0: (v4 & -2147483648), F1: 42, F2: 127}
                        switch "foo bar" {
                        case "\n\t":
                            v31 := GenS1{F0: (v3 >> (((v15 << (((v15 - v15)) & 63))) & 63)), F1: 120, F2: -1}
                            _ = v31
                        case "":
                            fmt.Println(v17, v17, v3)
                        }
                        if false {
                            fmt.Println(v3)
                        }
                        _ = v30
                    }
                    for fi32 := 0; fi32 < 5; fi32++ {
                        fmt.Println(v15)
                        var v33 uint = v15
                        _ = fi32
                        _ = v33
                    }
                }
            default:
                switch "foo bar" {
                case "hello":
                    v34 := GenS4{F0: 605}
                    _ = v34
                default:
                    var v35 int64 = (v3 % ((v3) | 1))
                    _ = v35
                }
            }
        }
    case "a":
        var v36 int8 = 0
        _ = v36
    default:
        var v37 bool = true
        _ = v37
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        fmt.Println(v17, fi38)
        _ = fi38
    }
    var v39 int = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
