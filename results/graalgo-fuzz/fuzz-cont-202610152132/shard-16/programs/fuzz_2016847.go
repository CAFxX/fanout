package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int16, p1 int32) uint8 {
    fmt.Println(p1, p0)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 255
}
func genFn1(p0 uint16, p1 int16, p2 uint32) int16 {
    v6 := GenS1{F0: 42}
    p0 -= (p0 % ((p0) | 1))
    p2 -= p2
    var v7 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 bool = false
    v2 := make(map[int]uint32)
    var v3 float64 = 0.1
    var v4 float32 = 3.14159
    var v5 int8 = 44
    v5 += (v5 << ((255) & 63))
    v4 -= 1.5
    switch "" {
    case "":
        fmt.Println(v1)
    case "foo bar":
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 bool = false
                v10 = v1
                if (949 >= 65535) {
                    if ((false || v1) && false) {
                        if (!(42 >= fi8)) {
                            v5 = v5
                        }
                        v11 := []int{fi8}
                        _ = v11
                    }
                    v12 := v2[42]
                    _ = v12
                } else {
                    v3 -= v3
                    var v13 uint64 = 4294967295
                    v13 = (v13 - 1)
                    _ = v13
                }
                v3 = -2.25
                _ = fi9
                _ = v10
            }
            v14 := GenS2{F0: 179, F1: 42}
            _ = fi8
            _ = v14
        }
    }
    var v15 float64 = (3.14159 / (3.14159 * (v3 + -2.25)))
    var v16 int16 = 1
    if false {
        for fi17 := 0; fi17 < 1; fi17++ {
            switch fi17 {
            case -9223372036854775808:
                switch ((("\n\t" + "") + "hello") + "\n\t") {
                case "\n\t":
                    switch "\n\t" {
                    case "":
                        var v18 int64 = 9223372036854775807
                        _ = v18
                    default:
                        v19 := []int{(fi17 & 653), 408, (fi17 >> ((0) & 63)), fi17}
                        _ = v19
                    }
                }
            case 235:
                var v20 int16 = (^(v16))
                _ = v20
            }
            _ = fi17
        }
    }
    v1 = v1
    var v21 uint32 = 1
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 string = ("a" + "a")
        var v24 uint8 = 255
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int16 = (-(v16))
    var v26 uint64 = 0
    for fi27 := 0; fi27 < 3; fi27++ {
        switch ("a" + (("\n\t" + "hello") + "foo bar")) {
        case "hello":
            var v28 bool = v1
            _ = v28
        case "":
            var v29 string = ("hello" + (("foo bar" + ("\n\t" + "hello")) + ("hello" + "a")))
            _ = v29
        default:
            switch (fi27 / (((fi27 % (((fi27 + -860)) | 1))) | 1)) {
            case 1:
                var v30 int16 = (v16 & v16)
                _ = v30
            case 9223372036854775807:
                fmt.Println(v5, v26, v21)
            default:
                switch fi27 {
                case -1:
                    v31 := GenS2{F0: 874, F1: 562}
                    _ = v31
                case 0:
                    var v32 int8 = 1
                    _ = v32
                }
            }
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
