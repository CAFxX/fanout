package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 bool }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int64 = 1
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]string)
    var v4 uint64 = 42
    var v5 bool = true
    v6 := make(map[string]uint32)
    fmt.Println(v4, v2, v4)
    var v7 uint16 = 0
    if v5 {
        switch (("hello" + ("" + "a")) + "a") {
        case "foo bar":
            var v8 int = 700
            _ = v8
        case "a":
            var v9 int32 = 574
            _ = v9
        default:
            var v10 int = int(v2)
            _ = v10
        }
    }
    var v11 bool = v5
    if true {
        v11 = false
    }
    if v5 {
        v2 -= -9223372036854775808
    }
    var v12 string = "hello"
    v4 -= (v4 + (v4 >> ((42) & 63)))
    if false {
        if (v7 != v7) {
            switch "hello" {
            case "\n\t":
                for fi13 := 0; fi13 < 5; fi13++ {
                    switch "a" {
                    case "":
                        var v14 int = fi13
                        _ = v14
                    default:
                        switch "a" {
                        case "hello":
                            v7 = v7
                        case "":
                            if (!v5) {
                                var v15 uint = 18446744073709551615
                                if false {
                                    var v16 float32 = 0.0
                                    var v17 int64 = v2
                                    _ = v16
                                    _ = v17
                                }
                                _ = len(v3)
                                _ = v15
                            }
                        case "foo bar":
                            switch fi13 {
                            case 1:
                                v18 := []int{-9223372036854775808, (fi13 | fi13)}
                                _ = v18
                            case 32767:
                                var v19 int8 = -72
                                _ = v19
                            case 609:
                                for fi20 := 0; fi20 < 12; fi20++ {
                                    var v21 int16 = -32768
                                    v2 = v2
                                    v22 := GenS1{F0: (!(0.1 <= 1e10)), F1: v7, F2: (2147483648 == 1)}
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            }
                        default:
                            switch (fi13 - fi13) {
                            case 128:
                                var v23 string = "hello"
                                _ = v23
                            case 1:
                                v1 = (^(v1))
                            }
                        }
                    }
                    v5 = (v11 && false)
                    fmt.Println(v11, v12, v7)
                    _ = fi13
                }
            case "":
                var v24 int64 = (v2 ^ (v2 + 9223372036854775807))
                _ = v24
            case "hello":
                v7 = v7
            default:
                switch ("hello" + "\n\t") {
                case "\n\t":
                    v25 := GenS2{F0: 0, F1: 100}
                    _ = v25
                case "hello":
                    var v26 uint = uint(v4)
                    _ = v26
                case "":
                    _ = len(v6)
                }
            }
            var v27 string = "\n\t"
            _ = v27
        }
        if (42 >= -128) {
            v7 += v7
        } else {
            v6[((v12 + "") + v12)] = 1
        }
        var v28 uint16 = (v7 * (v7 ^ v7))
        _ = v28
    }
    v5 = v5
    if (!false) {
        v29 := []int{32768, -9223372036854775808, -228}
        _ = v29
    } else {
        if ((51 != 1) && (0 <= 2147483647)) {
            v1 = v2
            var v30 uint64 = (v4 / ((v4) | 1))
            v31 := GenS3{F0: "hello", F1: v12}
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 7; fi32++ {
            switch v12 {
            case "foo bar":
                var v33 int = fi32
                _ = v33
            case "a":
                v4 = (v4 >> ((96) & 63))
            default:
                var v34 string = v12
                _ = v34
            }
            fmt.Println(v5, v7, v5)
            var v35 uint64 = v4
            _ = fi32
            _ = v35
        }
    }
    if (v7 <= (v7 + 184)) {
        v36 := GenS1{F0: (!v11), F1: (v7 % (((v7 ^ (^(v7)))) | 1)), F2: (v7 >= v7)}
        switch 32768 {
        case -9223372036854775808:
            switch 32767 {
            case -1:
                var v37 int16 = -184
                _ = v37
            case 0:
                v11 = v5
            default:
                v38 := GenS2{F0: 32768, F1: 50}
                _ = v38
            }
        case 127:
            v1 = -9223372036854775808
        case -2147483648:
            delete(v6, ((v12 + (((((v12 + v12) + v12) + (((v12 + v12) + (v12 + "foo bar")) + v12)) + (v12 + ("a" + "\n\t"))) + v12)) + (("foo bar" + (v12 + "\n\t")) + v12)))
        }
        switch -468 {
        case -1:
            v39 := GenS4{F0: (v7 << ((206) & 63))}
            _ = v39
        case 42:
            for fi40 := 0; fi40 < 7; fi40++ {
                for fi41 := 0; fi41 < 17; fi41++ {
                    var v42 int8 = 127
                    v7 = (v7 % ((0) | 1))
                    _ = fi41
                    _ = v42
                }
                _ = fi40
            }
        default:
            switch (v12 + "a") {
            case "a":
                for fi43 := 0; fi43 < 12; fi43++ {
                    fmt.Println(v12, fi43, v1)
                    _ = fi43
                }
            default:
                switch v12 {
                case "\n\t":
                    fmt.Println(v5)
                }
            }
        }
        _ = v36
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        var v45 float64 = float64(float32(0.0))
        delete(v3, "")
        fmt.Println(fi44, v2, v7)
        _ = fi44
        _ = v45
    }
    v46 := GenS1{F0: (!v5), F1: (v7 | 887), F2: v5}
    var v47 string = ("foo bar" + "")
    var v48 string = v47
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
