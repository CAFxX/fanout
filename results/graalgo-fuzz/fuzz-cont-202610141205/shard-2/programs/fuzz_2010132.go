package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = -2.25
    v3 := make(map[uint32]string)
    v2 -= 1.5
    if (((((1e10 * 1.5) - (((1.5 * 1.5) + -2.25) - 0.1)) * (-2.25 - (1.5 - (1.5 / 1.0)))) / 0.1) > -2.25) {
        var v4 uint = 1
        for fi5 := 0; fi5 < 19; fi5++ {
            fmt.Println(v4, v4, fi5)
            _ = fi5
        }
        if false {
            for fi6 := 0; fi6 < 5; fi6++ {
                v7 := []int{(fi6 << ((v4) & 63))}
                fmt.Println(fi6)
                if ((255 > 42) && (false || (9223372036854775807 >= -9223372036854775808))) {
                    v4 -= (v4 - (^(v4)))
                } else {
                    delete(v1, 4294967295)
                }
                _ = fi6
                _ = v7
            }
            if true {
                v4 = (v4 + 4294967295)
                fmt.Println(v4)
            } else {
                v4 += v4
            }
            fmt.Println(v2, v2)
        } else {
            v8 := GenS1{F0: (v4 + (v4 * v4))}
            _ = v8
        }
        _ = v4
    }
    v9 := []int{0, -9223372036854775808, 0, -2147483648}
    for fi10 := 0; fi10 < 7; fi10++ {
        switch (fi10 >> ((4294967295) & 63)) {
        case -9223372036854775808:
            switch ("" + ("\n\t" + "")) {
            case "foo bar":
                v11 := GenS2{F0: "\n\t", F1: -32768}
                _ = v11
            case "a":
                v2 -= ((v2 - 0.1) + 1.5)
            case "":
                fmt.Println(v2, fi10)
            }
        case 128:
            var v12 uint32 = 207
            _ = v12
        default:
            for fi13 := 0; fi13 < 9; fi13++ {
                fmt.Println(fi13, fi13, fi10)
                if false {
                    switch "a" {
                    case "hello":
                        v2 = v2
                    case "":
                        v14 := GenS3{F0: 0}
                        _ = v14
                    case "a":
                        v2 = v2
                    default:
                        v2 = 1e10
                    }
                    fmt.Println(fi10)
                }
                v15 := GenS1{F0: 4294967295}
                _ = fi13
                _ = v15
            }
        }
        v16 := GenS2{F0: "hello", F1: -32768}
        var v17 bool = false
        _ = fi10
        _ = v16
        _ = v17
    }
    var v18 int32 = 0
    v18 = -2147483648
    v19 := GenS1{F0: 0}
    if true {
        v18 = (v18 & (v18 >> ((18446744073709551615) & 63)))
        v20 := GenS2{F0: "\n\t", F1: -32768}
        if (((!true) && true) && ((((v2 + 1e10) / 1e10) + (((3.14159 / v2) / float64(float32(float64((3.14159 / -2.25))))) + (1.5 + 0.0))) == (v2 - float64((-2.25 + 0.1))))) {
            var v21 uint8 = 2
            _ = v21
        } else {
            var v22 int64 = 1
            _ = v22
        }
        _ = v20
    }
    var v23 uint64 = 1
    for fi24 := 0; fi24 < 20; fi24++ {
        fmt.Println(v2)
        var v25 uint32 = 713
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 string = "foo bar"
        var v28 string = v27
        v27 = v28
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 uint16 = 2
    v30 := GenS4{F0: 0}
    for fi31 := 0; fi31 < 3; fi31++ {
        v18 -= v18
        _ = fi31
    }
    if ("foo bar" > ("" + (("a" + (("hello" + ("\n\t" + "\n\t")) + "a")) + (("hello" + "") + "\n\t")))) {
        _ = len(v1)
        v18 += 859
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v9
}
