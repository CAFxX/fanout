package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 0
    var v3 int8 = -128
    var v4 int8 = -128
    var v5 float32 = 3.14159
    v6 := make(map[uint32]string)
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint32 = 4294967295
        var v10 bool = false
        if (715 <= (v2 / (((v2 & 1)) | 1))) {
            fmt.Println(v4, v7)
        } else {
            v4 = v3
            v7 += v7
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    switch 538 {
    case 127:
        var v11 int64 = 278
        _ = v11
    case 32768:
        switch "" {
        case "foo bar":
            var v12 uint = 42
            _ = v12
        case "a":
            v13 := GenS1{F0: 884, F1: false, F2: 1}
            _ = v13
        default:
            v14 := []int{728, 9223372036854775807, -1, 32767}
            _ = v14
        }
    }
    v3 -= v4
    fmt.Println(v2)
    var v15 int = -2147483648
    if false {
        v16 := []int{(v15 >> ((640) & 63)), 4294967295, 32767, (v15 / ((v15) | 1))}
        v7 = float64(1.5)
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := GenS1{F0: 1, F1: true, F2: 32767}
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    v5 = float32(((v7 - float64(0.0)) * 1.5))
    var v19 uint64 = 2
    fmt.Println(v15)
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 float32 = v5
        switch "foo bar" {
        case "foo bar":
            v6[42] = ("a" + "\n\t")
        case "a":
            var v22 uint8 = 255
            _ = v22
        default:
            v2 -= (v2 + v2)
        }
        var v23 float32 = (v5 / float32(v7))
        _ = fi20
        _ = v21
        _ = v23
    }
    v24 := []int{}
    fmt.Println(v5, v2)
    for fi25 := 0; fi25 < 16; fi25++ {
        switch "" {
        case "a":
            if (!("hello" == (("foo bar" + "\n\t") + "a"))) {
                fmt.Println(v7)
                fmt.Println(fi25, v7)
            } else {
                for fi26 := 0; fi26 < 5; fi26++ {
                    var v27 uint64 = v19
                    v5 = 0.0
                    _ = fi26
                    _ = v27
                }
            }
        case "hello":
            var v28 string = "\n\t"
            _ = v28
        case "foo bar":
            for fi29 := 0; fi29 < 2; fi29++ {
                v3 -= v4
                var v30 float32 = 3.14159
                v31 := GenS2{F0: 32767, F1: -9223372036854775808}
                if false {
                    var v32 uint16 = 1
                    var v33 int = (fi25 & (v15 % ((fi29) | 1)))
                    _ = v32
                    _ = v33
                } else {
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                    fmt.Println(v2)
                }
                switch -9223372036854775808 {
                case 127:
                    fmt.Println(v15, v19, v3)
                case 42:
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                }
                _ = fi29
                _ = v30
                _ = v31
            }
        default:
            v7 = v7
        }
        _ = fi25
    }
    fmt.Println(v2, v4)
    var v34 float32 = v5
    switch "hello" {
    case "hello":
        v34 = 0.0
    case "foo bar":
        v4 = v4
    case "":
        fmt.Println(v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
