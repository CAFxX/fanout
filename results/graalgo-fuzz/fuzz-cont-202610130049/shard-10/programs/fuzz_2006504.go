package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int64; F1 int8; F2 uint16 }
type GenS3 struct { F0 int32 }
func genFn0() string {
    v5 := []int{}
    var v6 int = 2147483647
    if ((4294967295 > 44) || false) {
        _ = len(v5)
        var v7 string = "\n\t"
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        v6 += (v6 & v6)
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 float64 = float64(-2.25)
            _ = fi9
            _ = v10
        }
        var v11 int = fi8
        _ = fi8
        _ = v11
    }
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    v1 := []uint32{2, 65535, 2147483647}
    v2 := []string{"a", "a", "a", ("foo bar" + "")}
    v3 := make(map[string]bool)
    var v4 uint16 = 255
    switch (("a" + ("a" + "hello")) + ("\n\t" + "\n\t")) {
    case "foo bar":
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                if ((false && (true && (492 != 983))) && true) {
                    var v14 uint = 4294967295
                    switch fi13 {
                    case 0:
                        _ = len(v2)
                    case 750:
                        var v15 uint8 = uint8(v14)
                        _ = v15
                    case 836:
                        var v16 uint16 = 1
                        _ = v16
                    }
                    v17 := GenS1{F0: 482, F1: 2}
                    _ = v14
                    _ = v17
                }
                switch (fi12 * fi12) {
                case 1:
                    v18 := GenS2{F0: -145, F1: 127, F2: (^(v4))}
                    _ = v18
                default:
                    _ = len(v1)
                }
                _ = fi13
            }
            if len(v2) > 0 {
                v2[0] = "a"
            }
            _ = fi12
        }
    case "hello":
        fmt.Println(v4)
    case "\n\t":
        v4 = 42
    default:
        switch -2147483648 {
        case -1:
            var v19 string = ("foo bar" + "a")
            _ = v19
        case 4294967295:
            if false {
                for fi20 := 0; fi20 < 4; fi20++ {
                    fmt.Println(fi20)
                    _ = fi20
                }
                fmt.Println(v4)
                v4 = (v4 - (v4 >> ((0) & 63)))
            } else {
                v4 -= v4
                var v21 bool = false
                v4 += 1
                _ = v21
            }
        default:
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v4)
            case "foo bar":
                if (38 >= -1) {
                    v22 := v3[("foo bar" + "hello")]
                    _ = v22
                }
            default:
                v4 += (v4 | v4)
            }
        }
    }
    var v23 int = -2147483648
    v24 := GenS1{F0: 108, F1: 2}
    _ = len(v3)
    v23 = v23
    switch "hello" {
    case "":
        v4 = (v4 >> ((470) & 63))
    case "\n\t":
        var v25 uint8 = 95
        _ = v25
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 uint64 = 732
        v28 := GenS2{F0: 9223372036854775807, F1: 42, F2: (v4 - (^(v4)))}
        if true {
            if false {
                if (!true) {
                    delete(v3, "\n\t")
                    var v29 int64 = 42
                    _ = v29
                }
            }
            switch (("foo bar" + "a") + ("" + "")) {
            case "\n\t":
                var v30 int = 784
                _ = v30
            case "hello":
                v31 := GenS3{F0: 651}
                _ = v31
            }
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v23, v23)
    v3["\n\t"] = (321 > 161)
    fmt.Println(v4)
    var v32 float64 = (0.0 - ((0.0 / 1.5) * 0.1))
    var v33 int = (v23 >> ((1) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
