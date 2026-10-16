package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0() string {
    var v6 uint = 1
    var v7 uint32 = 12
    _ = v6
    _ = v7
    return "\n\t"
}
func genFn1() string {
    var v8 int = 128
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        switch "foo bar" {
        case "":
            var v9 int16 = 1
            _ = v9
        case "hello":
            fmt.Println(v8)
        case "foo bar":
            v8 -= v8
        default:
            var v10 int16 = 1
            _ = v10
        }
    case "foo bar":
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint16 = 228
            v8 += 186
            v13 := GenS1{F0: 65535}
            _ = fi11
            _ = v12
            _ = v13
        }
    default:
        fmt.Println(v8)
    }
    _ = v8
    return "hello"
}

func main() {
    v1 := []uint32{42, 285, 891}
    v2 := make(map[string]string)
    v3 := []uint32{0, 4294967295, 366, 0}
    v4 := []int{-180}
    var v5 uint = 18446744073709551615
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := GenS2{F0: 243, F1: -1, F2: 4294967295}
        fmt.Println(fi14, v5)
        _ = fi14
        _ = v15
    }
    v3 = append(v3, 343)
    fmt.Println(v5)
    v5 -= (^(v5))
    v16 := v2[("hello" + "hello")]
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 int64 = -1
        var v19 uint64 = 1
        fmt.Println(v19, v18, v16)
        _ = fi17
        _ = v18
        _ = v19
    }
    switch -9223372036854775808 {
    case -292:
        var v20 int16 = 1000
        _ = v20
    case 34:
        for fi21 := 0; fi21 < 14; fi21++ {
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS3{F0: 1}
                v5 = 18446744073709551615
                var v24 int16 = -1
                _ = fi22
                _ = v23
                _ = v24
            }
            if true {
                var v25 uint32 = 1
                var v26 string = ("a" + v16)
                for fi27 := 0; fi27 < 2; fi27++ {
                    v28 := GenS2{F0: 42, F1: -9223372036854775808, F2: v25}
                    _ = fi27
                    _ = v28
                }
                _ = v25
                _ = v26
            }
            v29 := GenS2{F0: 127, F1: -9223372036854775808, F2: 2}
            _ = fi21
            _ = v29
        }
    }
    fmt.Println(v16, v16)
    v30 := GenS3{F0: 48}
    v1 = append(v1, 2147483648)
    delete(v2, (v16 + v16))
    v16 = (v16 + ((("foo bar" + ("\n\t" + v16)) + (v16 + "hello")) + ((v16 + v16) + "\n\t")))
    v16 = (v16 + (("hello" + "\n\t") + v16))
    var v31 int = 184
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
