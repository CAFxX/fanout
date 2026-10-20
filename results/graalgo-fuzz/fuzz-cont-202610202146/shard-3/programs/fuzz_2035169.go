package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 int32; F2 uint8 }
func genFn0() int16 {
    if (((3.14159 / 1.5) + 0.1) == float32(1e10)) {
        switch -1 {
        case -9223372036854775808:
            var v7 float64 = 1e10
            _ = v7
        default:
            switch (("hello" + "foo bar") + ("foo bar" + "foo bar")) {
            case "":
                if (0.1 >= (float64(3.14159) + -2.25)) {
                    v8 := []int{}
                    _ = v8
                } else {
                    v9 := GenS1{F0: 1}
                    _ = v9
                }
            case "hello":
                if true {
                    var v10 uint = 18446744073709551615
                    var v11 uint64 = 987
                    _ = v10
                    _ = v11
                }
            case "foo bar":
                var v12 uint = 65535
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            v14 := GenS2{F0: -1}
            _ = fi13
            _ = v14
        }
    }
    var v15 bool = (!(65535 > 65535))
    if (!(42 > 0)) {
        v16 := GenS2{F0: -2147483648}
        v15 = (2 < 0)
        _ = v16
    }
    var v17 int8 = -128
    v18 := []int{32768, -2147483648, int(v17)}
    _ = v15
    _ = v17
    _ = v18
    return -717
}
func genFn1() uint {
    switch (((("hello" + "hello") + ("hello" + ("\n\t" + "hello"))) + "a") + "hello") {
    case "\n\t":
        var v19 uint32 = 2
        _ = v19
    case "hello":
        switch "foo bar" {
        case "\n\t":
            for fi20 := 0; fi20 < 19; fi20++ {
                if (1.5 == (1e10 * (0.1 * ((1.5 + 0.0) / (float64(3.14159) / 1.0))))) {
                    v21 := []int{}
                    var v22 int32 = -2147483648
                    _ = v21
                    _ = v22
                } else {
                    if false {
                        fmt.Println(fi20)
                    }
                    var v23 bool = false
                    _ = v23
                }
                _ = fi20
            }
        case "hello":
            v24 := []int{128, 914, -2147483648, 1}
            _ = v24
        case "a":
            var v25 int = -2147483648
            _ = v25
        }
    }
    var v26 int64 = -149
    var v27 int16 = genFn0()
    _ = v26
    _ = v27
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 2
    var v3 float64 = ((-2.25 + 0.1) / 1.5)
    var v4 int64 = 9223372036854775807
    v5 := make(map[string]bool)
    var v6 string = "\n\t"
    v28 := []int{-2147483648, 617}
    switch 128 {
    case 128:
        var v29 string = v6
        _ = v29
    case 2147483647:
        var v30 int = -1
        _ = v30
    case 4294967295:
        v2 += v2
    }
    fmt.Println(v6, v6)
    v31 := GenS3{F0: v6}
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    var v32 float32 = 0.1
    if ((v2 / (((v2 - v2)) | 1)) <= 42) {
        fmt.Println(v3)
        if true {
            for fi33 := 0; fi33 < 3; fi33++ {
                for fi34 := 0; fi34 < 5; fi34++ {
                    var v35 int = (fi33 >> ((genFn1()) & 63))
                    v35 = (fi34 | (v35 | v35))
                    _ = fi34
                    _ = v35
                }
                _ = fi33
            }
        }
    }
    if len(v28) > 0 {
        v28[0] = -564
    }
    if ((false || true) || ((!false) || (!(false && true)))) {
        v36 := v1[128]
        _ = v36
    } else {
        var v37 uint64 = (v2 % (((v2 + v2)) | 1))
        _ = v37
    }
    var v38 int = -1
    for fi39 := 0; fi39 < 11; fi39++ {
        for fi40 := 0; fi40 < 4; fi40++ {
            v41 := GenS4{F0: v4, F1: -2147483648, F2: 128}
            var v42 uint = 212
            v42 = genFn1()
            _ = fi40
            _ = v41
            _ = v42
        }
        if ((genFn0() < genFn0()) || true) {
            if false {
                if true {
                    v2 -= 18446744073709551615
                }
            }
        }
        _ = fi39
    }
    var v43 uint = genFn1()
    v43 = v43
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
