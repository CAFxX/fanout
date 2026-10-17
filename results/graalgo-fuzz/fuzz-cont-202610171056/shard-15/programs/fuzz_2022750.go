package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int; F1 uint32; F2 string }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint16 }

func main() {
    var v1 bool = (!(!(1 >= -55)))
    var v2 int16 = -321
    var v3 uint8 = 137
    var v4 uint8 = (v3 >> ((65535) & 63))
    v5 := make(map[int]int)
    v6 := []uint32{2, 65535, 445, 65535}
    var v7 float32 = 1.5
    v2 = (^(v2))
    switch 2147483647 {
    case -9223372036854775808:
        var v8 uint = 18446744073709551615
        _ = v8
    case 1:
        switch "\n\t" {
        case "a":
            for fi9 := 0; fi9 < 8; fi9++ {
                for fi10 := 0; fi10 < 16; fi10++ {
                    if false {
                        v4 += (v4 * 255)
                        var v11 float32 = (float32(0.0) * (v7 / v7))
                        _ = v11
                    } else {
                        v3 -= v3
                    }
                    _ = fi10
                }
                _ = fi9
            }
        }
    case 0:
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 uint64 = 0
            _ = fi12
            _ = v13
        }
    default:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        switch (fi14 << ((894) & 63)) {
        case -507:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case -2147483648:
            fmt.Println(v2, v3, v7)
        case 644:
            var v15 int32 = 788
            _ = v15
        default:
            var v16 int16 = 42
            _ = v16
        }
        _ = fi14
    }
    switch ("a" + ((("" + ((("a" + ("a" + "a")) + "") + "foo bar")) + "hello") + ("a" + ("foo bar" + ("hello" + (("foo bar" + "hello") + ("a" + "a"))))))) {
    case "hello":
        for fi17 := 0; fi17 < 9; fi17++ {
            if (781 == -239) {
                v1 = v1
                var v18 string = "a"
                _ = v18
            } else {
                var v19 int = fi17
                fmt.Println(v1, v19, fi17)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v19
            }
            var v20 string = ("hello" + ((("" + "") + "") + "foo bar"))
            v3 = 191
            _ = fi17
            _ = v20
        }
    case "":
        delete(v5, 70)
    case "\n\t":
        for fi21 := 0; fi21 < 3; fi21++ {
            v7 = v7
            v3 -= (v3 >> ((0) & 63))
            v22 := GenS1{F0: 0, F1: "hello"}
            _ = fi21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 uint = 2
        switch fi23 {
        case -1:
            switch fi23 {
            case 32767:
                v25 := GenS2{F0: 0, F1: 1}
                _ = v25
            case 0:
                _ = len(v5)
            default:
                for fi26 := 0; fi26 < 2; fi26++ {
                    if (true || ((^(v24)) != (v24 ^ v24))) {
                        var v27 float32 = v7
                        _ = v27
                    } else {
                        if ((int64(v24) == -1) || ((v2 ^ v2) > (-(v2)))) {
                            var v28 uint64 = 319
                            v29 := GenS3{F0: (fi23 * (^(fi26))), F1: 730, F2: ("hello" + "hello")}
                            _ = v28
                            _ = v29
                        }
                        v30 := GenS4{F0: 42, F1: 1, F2: 32767}
                        fmt.Println(fi26)
                        _ = v30
                    }
                    switch fi23 {
                    case 311:
                        v2 -= 32767
                    }
                    v7 += ((v7 - (((v7 - (float32((1.5 * float64(v7))) / v7)) / ((0.0 * v7) / v7)) - 0.0)) + v7)
                    _ = fi26
                }
            }
        }
        _ = fi23
        _ = v24
    }
    var v31 uint64 = 0
    var v32 int16 = 42
    if v1 {
        var v33 float32 = v7
        v3 = (v4 / ((47) | 1))
        _ = v33
    }
    fmt.Println(v32, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
