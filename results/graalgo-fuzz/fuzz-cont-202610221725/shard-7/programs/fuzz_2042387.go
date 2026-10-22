package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int8; F2 int64 }
type GenS4 struct { F0 int16; F1 int64 }
func genFn0(p0 int32) uint {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: "foo bar", F1: 54, F2: 2}
        if (1 == 383) {
            p0 += 2147483647
            switch "a" {
            case "":
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := []int{}
                    _ = fi9
                    _ = v10
                }
            case "hello":
                p0 += (p0 | (p0 >> ((0) & 63)))
            case "\n\t":
                v11 := []int{fi7, 9223372036854775807}
                _ = v11
            }
            v12 := GenS1{F0: "hello", F1: 255, F2: 1}
            _ = v12
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := GenS2{F0: -1, F1: 0}
            var v15 int8 = 42
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi7
        _ = v8
    }
    v16 := []int{-473}
    _ = p0
    _ = v16
    return 65535
}
func genFn1(p0 uint64, p1 int32, p2 uint) int8 {
    if ((false || (p2 != (p2 - 18446744073709551615))) || (false && false)) {
        var v17 int = 128
        var v18 string = ("a" + "")
        _ = v17
        _ = v18
    } else {
        var v19 float64 = -2.25
        switch 956 {
        case 0:
            v20 := []int{797, 4294967295}
            _ = v20
        case 9223372036854775807:
            v21 := GenS2{F0: -27, F1: (p0 * (p0 % ((p0) | 1)))}
            _ = v21
        case 1:
            var v22 int16 = -32768
            _ = v22
        default:
            if true {
                v23 := []int{}
                var v24 float32 = (0.1 + -2.25)
                if false {
                    var v25 int64 = -1
                    _ = v25
                }
                _ = v23
                _ = v24
            }
        }
        var v26 uint = p2
        _ = v19
        _ = v26
    }
    var v27 uint16 = 395
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 1
}

func main() {
    var v1 float64 = -2.25
    var v2 int8 = -9
    var v3 bool = (0 != 2147483647)
    var v4 int64 = -557
    var v5 uint64 = 4294967295
    v6 := make(map[int]bool)
    switch ("foo bar" + "") {
    case "hello":
        v1 = -2.25
    }
    var v28 uint32 = 1
    var v29 bool = ((!v3) && v3)
    delete(v6, 782)
    v6[32768] = v3
    fmt.Println(v29, v29, v4)
    var v30 int64 = int64(v28)
    switch -7 {
    case 2147483647:
        v31 := GenS3{F0: v28, F1: genFn1(42, -2147483648, 0), F2: 531}
        _ = v31
    default:
        v5 += v5
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        var v33 int = -2147483648
        switch (-(fi32)) {
        case 32768:
            fmt.Println(v4, v5, v3)
        case -9223372036854775808:
            var v34 string = "a"
            _ = v34
        default:
            v35 := GenS4{F0: 0, F1: (v4 << ((267) & 63))}
            _ = v35
        }
        v36 := GenS2{F0: genFn1(v5, 42, 65535), F1: 175}
        _ = fi32
        _ = v33
        _ = v36
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        v4 = (v30 % (((v4 ^ (v4 / ((v30) | 1)))) | 1))
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
