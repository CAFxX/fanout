package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint16; F2 uint }
type GenS3 struct { F0 int8; F1 int; F2 int64 }
func genFn0() bool {
    if false {
        v5 := []int{9223372036854775807, 9223372036854775807}
        var v6 uint = 1
        _ = v5
        _ = v6
    }
    switch 256 {
    case -917:
        var v7 bool = false
        _ = v7
    case 128:
        v8 := []int{260, -1, -2147483648}
        _ = v8
    case 4294967295:
        if ((!true) && (-523 >= 42)) {
            switch 127 {
            case -656:
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 float64 = 0.1
                    _ = fi9
                    _ = v10
                }
            case 0:
                var v11 uint64 = 350
                _ = v11
            }
        } else {
            switch 32768 {
            case -9223372036854775808:
                v12 := []int{399, -2147483648, -9223372036854775808}
                _ = v12
            case 667:
                v13 := []int{-2147483648, 480, -9223372036854775808}
                _ = v13
            case -644:
                var v14 uint = 255
                _ = v14
            default:
                var v15 float32 = 3.14159
                _ = v15
            }
        }
    }
    return false
}
func genFn1(p0 uint16, p1 uint32, p2 int32) uint32 {
    switch "a" {
    case "a":
        var v16 int32 = p2
        _ = v16
    case "\n\t":
        p0 = p0
    case "hello":
        if (527 < 18446744073709551615) {
            var v17 uint32 = (p1 + p1)
            v18 := GenS1{F0: (false || (((1000 < 951) && (true || false)) && ("a" >= "")))}
            _ = v17
            _ = v18
        }
    }
    var v19 int32 = p2
    if (((p2 <= v19) || true) || (!("foo bar" != ("" + ("\n\t" + ("hello" + "\n\t")))))) {
        v20 := []int{721, 2147483647}
        for fi21 := 0; fi21 < 5; fi21++ {
            if (!false) {
                var v22 uint32 = (p1 << ((246) & 63))
                v23 := GenS1{F0: ((p1 >= 0) || ((^(p0)) != (p0 ^ p0)))}
                for fi24 := 0; fi24 < 18; fi24++ {
                    var v25 int = 2147483647
                    v20 = append(v20, -521)
                    _ = fi24
                    _ = v25
                }
                _ = v22
                _ = v23
            }
            for fi26 := 0; fi26 < 6; fi26++ {
                if (("" + "foo bar") > "") {
                    switch ("" + "") {
                    case "hello":
                        fmt.Println(fi21, fi21, fi21)
                    case "foo bar":
                        fmt.Println(p1, p2)
                    }
                } else {
                    var v27 bool = false
                    v19 += -1
                    _ = v27
                }
                fmt.Println(p0)
                _ = fi26
            }
            _ = fi21
        }
        _ = v20
    } else {
        v19 -= (p2 * v19)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 0
}

func main() {
    v1 := []uint32{848, 42, 2}
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    v4 := []int{-408}
    var v28 uint8 = 0
    v29 := GenS2{F0: -32768, F1: 2, F2: 65535}
    fmt.Println(v28)
    fmt.Println(v28)
    for fi30 := 0; fi30 < 4; fi30++ {
        v31 := GenS3{F0: -1, F1: (fi30 | 42), F2: 1}
        var v32 uint64 = 0
        _ = fi30
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        fmt.Println(fi33)
        _ = fi33
    }
    var v34 int32 = 1
    var v35 float64 = (3.14159 - float64(1e10))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
