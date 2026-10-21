package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int32; F1 uint8; F2 int8 }
type GenS3 struct { F0 string }
func genFn0(p0 bool, p1 int32) int {
    switch 859 {
    case 442:
        v4 := GenS1{F0: ("\n\t" < ("foo bar" + ((("hello" + ("foo bar" + ("a" + ("hello" + "")))) + ("foo bar" + "foo bar")) + ("\n\t" + ((("" + ("\n\t" + "\n\t")) + ("\n\t" + "foo bar")) + "\n\t"))))), F1: ("" != "hello")}
        _ = v4
    }
    v5 := GenS1{F0: p0, F1: ((42 < 42) && ((251 > 2) || p0))}
    p0 = false
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{fi6, -1}
        _ = fi6
        _ = v7
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v5
    return 684
}
func genFn1(p0 int64, p1 int8) int {
    var v8 uint64 = 369
    v9 := []int{2147483647, 909}
    for fi10 := 0; fi10 < 20; fi10++ {
        v8 = 18446744073709551615
        switch int(v8) {
        case -746:
            var v11 string = ("" + "")
            _ = v11
        case 802:
            var v12 int64 = p0
            _ = v12
        default:
            var v13 uint8 = 127
            _ = v13
        }
        v9 = append(v9, genFn0(false, 1))
        _ = fi10
    }
    var v14 uint32 = 946
    _ = p0
    _ = p1
    _ = v14
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 float32 = (0.1 / -2.25)
    v2 := []string{"\n\t", "foo bar", "", (("\n\t" + "\n\t") + "hello")}
    var v3 uint64 = 43
    for fi15 := 0; fi15 < 16; fi15++ {
        v1 = 0.1
        var v16 string = ((("" + (("\n\t" + "foo bar") + "a")) + "hello") + "foo bar")
        switch v16 {
        case "foo bar":
            v2 = append(v2, v16)
        case "":
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 uint64 = 902
                _ = fi17
                _ = v18
            }
        case "a":
            v1 -= v1
        default:
            fmt.Println(fi15, v3)
        }
        _ = fi15
        _ = v16
    }
    v19 := GenS1{F0: (!false), F1: false}
    v20 := GenS2{F0: 2147483647, F1: 127, F2: -111}
    fmt.Println(v1, v3)
    var v21 int32 = 2147483647
    switch "foo bar" {
    case "\n\t":
        if (0 == 255) {
            v21 = 0
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 uint = 34
                _ = len(v2)
                var v24 int16 = -915
                _ = fi22
                _ = v23
                _ = v24
            }
            var v25 int32 = 203
            _ = v25
        }
    default:
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v26 uint16 = 65535
            _ = v26
        }
    }
    var v27 uint32 = 4294967295
    var v28 int = 219
    if true {
        v21 = (v21 + (v21 ^ (v21 | -2147483648)))
    } else {
        v29 := GenS3{F0: "foo bar"}
        var v30 uint32 = 65535
        _ = v29
        _ = v30
    }
    var v31 uint32 = (v27 - (v27 % ((v27) | 1)))
    _ = len(v2)
    switch "foo bar" {
    case "hello":
        if false {
            v27 -= (v27 * v27)
        }
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    switch (v28 >> ((uint(v31)) & 63)) {
    case 9223372036854775807:
        for fi32 := 0; fi32 < 8; fi32++ {
            var v33 string = "hello"
            _ = fi32
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v31
}
