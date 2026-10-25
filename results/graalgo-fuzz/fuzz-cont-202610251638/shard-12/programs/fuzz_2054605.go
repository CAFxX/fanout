package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    v3 := []uint32{1, 48, 42}
    v4 := make(map[string]uint32)
    v5 := []int{-2147483648, -308, 10, 1}
    var v6 bool = (18446744073709551615 == 18446744073709551615)
    var v7 bool = (65535 <= 65535)
    var v8 uint = 1
    var v9 bool = v6
    if (!v6) {
        v8 = 18446744073709551615
    } else {
        for fi10 := 0; fi10 < 4; fi10++ {
            if v7 {
                for fi11 := 0; fi11 < 10; fi11++ {
                    var v12 uint64 = 18446744073709551615
                    v2[525] = ""
                    v13 := GenS1{F0: 100, F1: 1}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                for fi14 := 0; fi14 < 18; fi14++ {
                    fmt.Println(v6, v8)
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v6, v9, fi14)
                    }
                    _ = fi14
                }
            } else {
                v8 = (v8 >> (((v8 ^ v8)) & 63))
                if len(v5) > 0 {
                    v5[0] = fi10
                }
                switch ("a" + "a") {
                case "":
                    var v15 uint16 = 255
                    _ = v15
                case "a":
                    v4["foo bar"] = uint32(v8)
                case "foo bar":
                    v6 = ((v7 && v6) || true)
                default:
                    var v16 uint32 = 264
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                v6 = (!(!(-14 >= 0)))
                _ = fi17
            }
            _ = fi10
        }
    }
    if ((!v7) || false) {
        var v18 uint = v8
        var v19 float64 = 3.14159
        _ = v18
        _ = v19
    }
    switch "\n\t" {
    case "":
        switch -921 {
        case 32767:
            v8 = (^(v8))
        case 0:
            var v20 float64 = (-2.25 + (((0.1 - 3.14159) / 1.0) / (-2.25 * 0.1)))
            _ = v20
        }
    case "hello":
        if true {
            fmt.Println(v7, v7)
            var v21 uint64 = 1
            var v22 string = (("\n\t" + "a") + ("foo bar" + ("" + (("foo bar" + "a") + (("foo bar" + "\n\t") + ("a" + "hello"))))))
            _ = v21
            _ = v22
        } else {
            v9 = true
            fmt.Println(v9)
        }
    case "foo bar":
        if (v9 && true) {
            switch -1 {
            case -1:
                v9 = v6
            }
            var v23 float64 = 1.5
            var v24 int16 = 0
            _ = v23
            _ = v24
        } else {
            var v25 float32 = -2.25
            v25 = v25
            var v26 uint = v8
            _ = v25
            _ = v26
        }
    default:
        v27 := GenS2{F0: 176}
        _ = v27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        v8 += v8
        _ = fi28
    }
    var v29 int16 = 1000
    var v30 int8 = 100
    var v31 uint16 = uint16(v8)
    v8 = (v8 | v8)
    v31 += (v31 + (v31 & v31))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
