package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{2, 191, 2147483647}
    var v3 uint32 = 2
    var v4 uint16 = 65535
    var v5 int64 = int64(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        v4 -= v4
    }
    var v6 bool = true
    for fi7 := 0; fi7 < 7; fi7++ {
        switch "foo bar" {
        case "hello":
            v5 -= (^(v5))
        default:
            v8 := GenS1{F0: v3, F1: 42}
            _ = v8
        }
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS2{F0: -1, F1: v1}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch ("" + ("hello" + "hello")) {
            case "foo bar":
                var v11 uint64 = v1
                _ = v11
            case "a":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "":
                v12 := GenS2{F0: 2147483647, F1: v1}
                _ = v12
            default:
                if false {
                    fmt.Println(v6, v1)
                    v5 = v5
                }
            }
            if len(v2) > 0 {
                v2[0] = uint32(v1)
            }
            fmt.Println(v6, v4, v5)
            _ = fi9
            _ = v10
        }
        var v13 uint8 = 154
        _ = fi7
        _ = v13
    }
    var v14 float64 = 1e10
    var v15 uint8 = 255
    if ((("foo bar" + ("\n\t" + ("" + ""))) == ("" + (("hello" + ("a" + "\n\t")) + ("hello" + ("" + "a"))))) && true) {
        _ = len(v2)
        v16 := GenS1{F0: uint32(v4), F1: v15}
        v6 = v6
        _ = v16
    }
    switch ("hello" + ("" + "hello")) {
    case "":
        v14 += float64((0.1 / float32((v14 + v14))))
    case "foo bar":
        var v17 float32 = float32(v14)
        _ = v17
    case "a":
        var v18 int8 = 42
        _ = v18
    }
    var v19 bool = true
    for fi20 := 0; fi20 < 19; fi20++ {
        if (uint64(v4) != (v1 * v1)) {
            for fi21 := 0; fi21 < 9; fi21++ {
                if false {
                    v15 = v15
                    var v22 bool = v19
                    var v23 int64 = (v5 / ((v5) | 1))
                    _ = v22
                    _ = v23
                }
                _ = fi21
            }
            v14 -= v14
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            v15 -= uint8(v4)
            fmt.Println(fi20)
            if (true || ((v3 | 1) < v3)) {
                switch "\n\t" {
                case "foo bar":
                    v15 = v15
                case "\n\t":
                    var v25 int32 = 42
                    _ = v25
                default:
                    v4 = v4
                }
                switch -720 {
                case 282:
                    var v26 string = "\n\t"
                    _ = v26
                case -1:
                    v27 := GenS3{F0: (v6 && (("a" + "\n\t") != ("\n\t" + "foo bar")))}
                    _ = v27
                case 127:
                    v3 = 2147483648
                }
                v3 = (v3 % (((v3 % ((v3) | 1))) | 1))
            } else {
                v5 = (v5 << ((385) & 63))
            }
            fmt.Println(v1, v14)
            if (!(2147483647 >= 315)) {
                fmt.Println(v1)
                var v28 float64 = float64(-2.25)
                _ = v28
            }
            _ = fi24
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
