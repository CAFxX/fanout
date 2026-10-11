package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint8; F2 int64 }

func main() {
    v1 := []string{"\n\t", "a", ("hello" + "\n\t"), ("a" + "foo bar")}
    v2 := []int{}
    v3 := []int{253, 613, -2147483648, 443}
    v4 := make(map[int]uint32)
    var v5 float32 = 1.5
    var v6 int32 = -1
    v5 = float32(3.14159)
    var v7 uint64 = 4294967295
    v7 += (^(v7))
    switch ("" + (("\n\t" + ("foo bar" + (("a" + "") + "\n\t"))) + ("hello" + ""))) {
    case "hello":
        var v8 uint8 = 72
        _ = v8
    case "\n\t":
        if false {
            var v9 bool = (-348 == -162)
            _ = v9
        } else {
            var v10 bool = true
            switch "hello" {
            case "\n\t":
                if false {
                    for fi11 := 0; fi11 < 12; fi11++ {
                        var v12 uint64 = (v7 | (v7 % ((2) | 1)))
                        if (100 > 100) {
                            switch "" {
                            case "\n\t":
                                v10 = (!false)
                            case "foo bar":
                                if ((float32(-2.25) / ((1e10 - -2.25) - 0.0)) == v5) {
                                    for fi13 := 0; fi13 < 1; fi13++ {
                                        switch (^(fi13)) {
                                        case 693:
                                            fmt.Println(v12, v12)
                                            v7 = (^(v12))
                                            v14 := GenS1{F0: (-(fi11)), F1: 127}
                                            _ = v14
                                        case 127:
                                            switch "\n\t" {
                                            case "":
                                                v5 += float32((0.0 + (1e10 / 1e10)))
                                            }
                                        case -1:
                                            v12 = (v7 * (^(v7)))
                                        default:
                                            var v15 float64 = (float64(v5) + float64(1e10))
                                            _ = v15
                                        }
                                        var v16 int64 = 1
                                        _ = fi13
                                        _ = v16
                                    }
                                    switch (((((("" + "\n\t") + (("" + "a") + "foo bar")) + "a") + ("\n\t" + "a")) + "hello") + "") {
                                    case "\n\t":
                                        var v17 uint = 1
                                        _ = v17
                                    default:
                                        v18 := GenS2{F0: 42}
                                        _ = v18
                                    }
                                } else {
                                    if v10 {
                                        v19 := GenS1{F0: fi11, F1: 2}
                                        fmt.Println(v7, fi11)
                                        _ = v19
                                    } else {
                                        var v20 uint = 1
                                        _ = v20
                                    }
                                    v21 := GenS2{F0: 32767}
                                    var v22 int32 = (^(v6))
                                    _ = v21
                                    _ = v22
                                }
                            }
                        } else {
                            var v23 int = fi11
                            _ = v23
                        }
                        _ = fi11
                        _ = v12
                    }
                    v10 = false
                } else {
                    switch ("" + "foo bar") {
                    case "hello":
                        v6 += v6
                    }
                }
            case "a":
                var v24 string = ""
                _ = v24
            case "hello":
                switch (("" + "hello") + "a") {
                case "":
                    var v25 float64 = 0.0
                    _ = v25
                case "foo bar":
                    v2 = append(v2, 469)
                }
            default:
                fmt.Println(v10, v5)
            }
            _ = v10
        }
    default:
        for fi26 := 0; fi26 < 15; fi26++ {
            if (65535 >= 0) {
                if (((v7 + v7) <= (v7 ^ v7)) || ((v7 & v7) == v7)) {
                    var v27 int64 = 828
                    _ = v27
                }
                if true {
                    v3 = append(v3, (fi26 | (fi26 >> ((481) & 63))))
                    var v28 uint16 = 255
                    var v29 uint32 = 0
                    _ = v28
                    _ = v29
                } else {
                    fmt.Println(fi26, fi26)
                    v5 = v5
                }
                fmt.Println(v6, v7)
            } else {
                v6 += (-(v6))
                v30 := GenS1{F0: fi26, F1: 127}
                var v31 int = (fi26 << ((965) & 63))
                _ = v30
                _ = v31
            }
            _ = fi26
        }
    }
    v6 += (v6 ^ (v6 + v6))
    for fi32 := 0; fi32 < 7; fi32++ {
        v33 := GenS3{F0: (fi32 >> ((2) & 63)), F1: (fi32 / ((fi32) | 1)), F2: (-9223372036854775808 > 9223372036854775807)}
        _ = fi32
        _ = v33
    }
    fmt.Println(v6)
    v34 := GenS4{F0: 4294967295}
    fmt.Println(v6, v6, v5)
    var v35 int16 = -32768
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 int32 = -2147483648
        _ = fi36
        _ = v37
    }
    switch int(v7) {
    case 1:
        v7 -= 2
    case 0:
        fmt.Println(v35, v35, v35)
    case 127:
        for fi38 := 0; fi38 < 6; fi38++ {
            v39 := GenS3{F0: fi38, F1: (fi38 | fi38), F2: false}
            _ = fi38
            _ = v39
        }
    }
    switch 42 {
    case -9223372036854775808:
        switch 864 {
        case -1:
            v5 = v5
        case 9223372036854775807:
            v7 = (v7 << ((255) & 63))
        case -180:
            v40 := GenS5{F0: 169, F1: 42, F2: 42}
            _ = v40
        }
    }
    fmt.Println(v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
