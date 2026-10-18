package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 string; F1 uint32; F2 string }
type GenS4 struct { F0 uint8; F1 int64; F2 int64 }
type GenS5 struct { F0 string; F1 bool; F2 uint }
func genFn0() uint16 {
    var v4 int16 = 1000
    v5 := []int{-114}
    var v6 uint16 = 304
    var v7 int8 = -128
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (v6 + (v6 << ((18446744073709551615) & 63)))
}
func genFn1(p0 int8) uint16 {
    fmt.Println(p0)
    v8 := GenS1{F0: 168, F1: (0 != p0)}
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS2{F0: -1, F1: 9223372036854775807}
        _ = fi9
        _ = v10
    }
    v11 := []int{1, 501, 732}
    _ = p0
    _ = v11
    _ = v8
    return 65535
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]string)
    var v3 string = (("" + "a") + "hello")
    var v12 uint64 = 18446744073709551615
    v13 := []int{42, 42, 32767}
    if (false || (!true)) {
        if (42 != -1) {
            v3 = v3
            switch "a" {
            case "foo bar":
                fmt.Println(v12, v3)
            case "hello":
                fmt.Println(v12)
            default:
                v14 := GenS3{F0: ((v3 + (v3 + ((("hello" + ("hello" + "\n\t")) + v3) + "hello"))) + v3), F1: 2147483648, F2: v3}
                _ = v14
            }
            var v15 int64 = 300
            _ = v15
        }
    } else {
        var v16 string = v3
        v1 += 3.14159
        v3 = (v3 + "hello")
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS3{F0: (v3 + v3), F1: 770, F2: v3}
        v3 = ("foo bar" + "\n\t")
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            v3 = (v3 + (("a" + v3) + ""))
            _ = fi20
        }
        var v21 bool = (42 == -1)
        _ = fi19
        _ = v21
    }
    fmt.Println(v3, v3)
    v2[65535] = "a"
    fmt.Println(v1)
    fmt.Println(v1, v3)
    var v22 uint = 1
    for fi23 := 0; fi23 < 1; fi23++ {
        if len(v13) > 0 {
            v13[0] = 1
        }
        var v24 string = v3
        var v25 int = fi23
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 uint8 = 248
    if true {
        for fi27 := 0; fi27 < 4; fi27++ {
            if (991 != 956) {
                if ((v26 & (v26 | v26)) >= v26) {
                    fmt.Println(fi27, fi27)
                    v22 = 18446744073709551615
                }
            }
            if false {
                var v28 uint32 = 0
                _ = v28
            } else {
                v22 = (^(v22))
                for fi29 := 0; fi29 < 6; fi29++ {
                    if false {
                        if true {
                            v30 := GenS4{F0: (v26 ^ 255), F1: -224, F2: int64(v12)}
                            if true {
                                v26 = v26
                            } else {
                                if ((v26 + 128) != v26) {
                                    v31 := GenS5{F0: v3, F1: (v1 <= v1), F2: v22}
                                    _ = v31
                                }
                                var v32 int32 = 0
                                _ = v32
                            }
                            v26 -= (v26 ^ 128)
                            _ = v30
                        }
                        v3 = ""
                    }
                    var v33 int = fi27
                    _ = fi29
                    _ = v33
                }
            }
            _ = fi27
        }
        switch (v3 + (v3 + v3)) {
        case "\n\t":
            fmt.Println(v3, v26)
        default:
            v3 = v3
        }
        _ = len(v13)
    }
    var v34 bool = (668 != 678)
    var v35 float32 = 1e10
    for fi36 := 0; fi36 < 17; fi36++ {
        var v37 int32 = -1
        _ = fi36
        _ = v37
    }
    v3 = ("foo bar" + "a")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v34
    _ = v35
}
