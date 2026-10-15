package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 bool; F1 int16; F2 int16 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 int; F1 uint; F2 uint32 }
func genFn0() int8 {
    var v4 uint32 = 1
    var v5 uint32 = (^(v4))
    v5 = (v4 - v4)
    if (!(767 > 2)) {
        fmt.Println(v5, v4)
        v6 := GenS1{F0: 102}
        for fi7 := 0; fi7 < 9; fi7++ {
            switch ("\n\t" + "") {
            case "\n\t":
                v8 := GenS1{F0: 1}
                _ = v8
            case "hello":
                v9 := GenS2{F0: -2147483648, F1: 42}
                _ = v9
            case "foo bar":
                v4 = (v4 - 969)
            default:
                for fi10 := 0; fi10 < 10; fi10++ {
                    v11 := []int{(fi10 << ((503) & 63)), int(v5), fi10, (fi7 - (fi10 % ((fi10) | 1)))}
                    fmt.Println(v4)
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = fi10
                    _ = v11
                }
            }
            var v12 int32 = 2147483647
            _ = fi7
            _ = v12
        }
        _ = v6
    } else {
        v5 -= v5
    }
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int16 = 816
    v2 := make(map[uint32]uint32)
    v3 := []int{-1}
    if false {
        var v13 int8 = -127
        if (false && true) {
            v13 = (v13 * (v13 % ((127) | 1)))
            v14 := v2[2147483647]
            _ = v14
        } else {
            v15 := GenS1{F0: 2}
            _ = v15
        }
        _ = v13
    }
    fmt.Println(v1)
    v1 = -1
    v2[42] = 869
    v1 = (v1 | -61)
    v1 -= v1
    switch -406 {
    case 63:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -273:
        v1 -= (v1 & (v1 * v1))
    case 128:
        _ = len(v2)
    default:
        switch "" {
        case "":
            v1 = v1
        case "foo bar":
            if ("" > ("a" + "")) {
                v1 = 32767
            }
        case "a":
            if false {
                fmt.Println(v1)
                v16 := GenS3{F0: 42, F1: 0}
                v17 := GenS4{F0: true, F1: (v1 << ((2) & 63)), F2: (v1 & (v1 * v1))}
                _ = v16
                _ = v17
            } else {
                switch ("\n\t" + ((((("foo bar" + "\n\t") + ("" + ("\n\t" + ("" + "foo bar")))) + "") + "\n\t") + "\n\t")) {
                case "\n\t":
                    fmt.Println(v1)
                case "hello":
                    var v18 string = ("a" + "foo bar")
                    _ = v18
                }
                v1 = v1
            }
        default:
            for fi19 := 0; fi19 < 5; fi19++ {
                v1 -= 32767
                var v20 uint16 = 65535
                _ = fi19
                _ = v20
            }
        }
    }
    if (!true) {
        fmt.Println(v1)
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        if ((202 < 128) && (0.1 <= ((3.14159 / 1.0) * 1.5))) {
            fmt.Println(v1, fi21)
        } else {
            fmt.Println(fi21, v1)
            v3 = append(v3, (fi21 << ((18446744073709551615) & 63)))
        }
        v1 = v1
        _ = fi21
    }
    var v22 int8 = 42
    switch 42 {
    case 1:
        v3 = append(v3, 0)
    case 128:
        var v23 uint64 = 18446744073709551615
        _ = v23
    case 42:
        for fi24 := 0; fi24 < 11; fi24++ {
            v25 := GenS5{F0: 383, F1: v22}
            v26 := GenS6{F0: fi24, F1: 65535, F2: 2147483648}
            _ = fi24
            _ = v25
            _ = v26
        }
    default:
        _ = len(v3)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
}
