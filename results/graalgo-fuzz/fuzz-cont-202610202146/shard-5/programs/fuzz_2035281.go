package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int32; F1 int }

func main() {
    var v1 int8 = 0
    v2 := []uint32{2147483648, 0, 14}
    v3 := make(map[uint32]int)
    var v4 string = "a"
    var v5 string = ((v4 + "\n\t") + v4)
    var v6 uint32 = 374
    if (1 == 42) {
        v6 = (v6 & v6)
        v5 = "hello"
        var v7 float32 = (0.0 + (1.5 + 0.1))
        _ = v7
    }
    var v8 uint32 = (v6 << ((4294967295) & 63))
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(v5, v8)
        _ = fi9
    }
    if (-1 != 1) {
        v6 += (v6 / ((2147483647) | 1))
        if (-9223372036854775808 == -456) {
            if (0 <= 42) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v1 = (v1 - v1)
            }
            switch "a" {
            case "a":
                if (false || ("\n\t" > v4)) {
                    var v10 uint8 = 127
                    if (255 >= 42) {
                        var v11 uint32 = (v6 & (v8 - v8))
                        var v12 int64 = -512
                        v13 := GenS1{F0: v12, F1: v12}
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    switch -2147483648 {
                    case 722:
                        var v14 uint8 = 250
                        _ = v14
                    }
                    _ = v10
                } else {
                    delete(v3, 233)
                }
            case "hello":
                switch 32768 {
                case -694:
                    fmt.Println(v8)
                case 1:
                    delete(v3, (v8 | (^(v6))))
                case 32767:
                    switch 4294967295 {
                    case 42:
                        v2 = append(v2, (v8 - (v8 / ((v6) | 1))))
                    case 1:
                        v6 = v6
                    }
                default:
                    if false {
                        v4 = (v4 + v5)
                    } else {
                        if (v8 > 42) {
                            v15 := v3[2147483648]
                            var v16 string = v5
                            _ = v15
                            _ = v16
                        }
                        fmt.Println(v1, v5, v1)
                        fmt.Println(v4)
                    }
                }
            }
            if (((v5 + v4) + "foo bar") > v4) {
                v5 = v5
                v5 = v5
                v4 = ("a" + v5)
            } else {
                var v17 bool = (!(!(true && (612 <= 4))))
                var v18 uint64 = uint64(v8)
                var v19 string = ("a" + "\n\t")
                _ = v17
                _ = v18
                _ = v19
            }
        }
        v4 = (v4 + (v5 + v5))
    }
    if (((!true) && false) || (42 == -1)) {
        for fi20 := 0; fi20 < 6; fi20++ {
            for fi21 := 0; fi21 < 19; fi21++ {
                if ((((18446744073709551615 >= 42) || (!false)) && true) && true) {
                    var v22 uint32 = v8
                    _ = v22
                }
                v8 += (v8 | v8)
                _ = fi21
            }
            v6 = 2147483647
            var v23 uint32 = v8
            _ = fi20
            _ = v23
        }
        switch v5 {
        case "":
            v24 := GenS2{F0: 322, F1: "foo bar"}
            _ = v24
        default:
            var v25 int16 = -1
            _ = v25
        }
        for fi26 := 0; fi26 < 11; fi26++ {
            fmt.Println(v5, v8)
            _ = fi26
        }
    } else {
        if (2147483647 > -2147483648) {
            var v27 int8 = v1
            var v28 int = 0
            _ = v27
            _ = v28
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v29 float64 = 1e10
    v30 := GenS3{F0: 1, F1: 32767}
    var v31 int = 128
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        v1 += 127
        _ = fi32
    }
    v3[4294967295] = (v31 + 4294967295)
    v31 = 128
    switch (("a" + v5) + (((v5 + v4) + v5) + (v5 + ((v5 + "foo bar") + "hello")))) {
    case "":
        var v33 bool = false
        _ = v33
    case "a":
        fmt.Println(v31, v6)
    case "\n\t":
        fmt.Println(v8, v1, v5)
    }
    var v34 int32 = 0
    var v35 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
