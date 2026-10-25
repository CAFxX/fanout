package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int32 }

func main() {
    v1 := []string{("hello" + "a"), "\n\t"}
    var v2 float64 = (float64(0.1) - (((0.0 / -2.25) + 0.1) + 1e10))
    var v3 uint = 833
    v4 := []uint32{2, 384}
    v5 := []uint32{}
    v6 := make(map[uint32]bool)
    var v7 int8 = 127
    v8 := GenS1{F0: 42, F1: 127}
    switch ("foo bar" + ("a" + "\n\t")) {
    case "":
        switch -667 {
        case -1:
            var v9 int16 = -256
            _ = v9
        case -9223372036854775808:
            v10 := GenS2{F0: 42}
            _ = v10
        case -2147483648:
            var v11 uint8 = 141
            _ = v11
        default:
            if (!((((("" + "\n\t") + "hello") + (((("\n\t" + ("foo bar" + "foo bar")) + ("\n\t" + (("foo bar" + ("a" + (("hello" + ("a" + "foo bar")) + ("" + "hello")))) + ("hello" + ("hello" + "\n\t"))))) + "") + ("" + ""))) + "") != "a")) {
                var v12 uint = v3
                for fi13 := 0; fi13 < 13; fi13++ {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        v7 = v7
                        _ = fi14
                    }
                    _ = fi13
                }
                for fi15 := 0; fi15 < 3; fi15++ {
                    v12 += (v3 * (v12 >> ((v3) & 63)))
                    _ = fi15
                }
                _ = v12
            }
        }
    }
    v16 := GenS3{F0: 4294967295, F1: 42}
    v7 = 1
    _ = len(v6)
    for fi17 := 0; fi17 < 7; fi17++ {
        v6[42] = ((^(fi17)) != fi17)
        _ = fi17
    }
    var v18 uint64 = 0
    if ((v3 - v3) <= (^(v3))) {
        v18 = (v18 << (((v3 << ((v3) & 63))) & 63))
        v3 -= 692
        v7 += v7
    } else {
        v19 := GenS3{F0: v18, F1: 2147483647}
        var v20 int64 = 1
        if ((!((-1 == (-(v7))) || (240 < v3))) && false) {
            var v21 int16 = 409
            _ = v21
        }
        _ = v19
        _ = v20
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 float64 = v2
        if true {
            v24 := GenS1{F0: 227, F1: fi22}
            var v25 uint = (v3 * v3)
            var v26 float64 = (v23 / (v23 + (1e10 + float64((0.0 + ((1e10 * 1e10) * 1.5))))))
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 uint64 = 943
        var v29 uint64 = v28
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
