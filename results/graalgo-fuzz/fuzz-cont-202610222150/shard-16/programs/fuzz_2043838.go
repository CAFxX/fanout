package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 string; F1 int; F2 uint64 }
type GenS5 struct { F0 uint; F1 string; F2 uint }
func genFn0(p0 int32) int64 {
    switch 1 {
    case 108:
        for fi5 := 0; fi5 < 12; fi5++ {
            switch ("foo bar" + ("\n\t" + (((("foo bar" + "hello") + ("hello" + "hello")) + "hello") + ""))) {
            case "hello":
                v6 := []int{}
                _ = v6
            case "foo bar":
                switch (fi5 % ((fi5) | 1)) {
                case 32767:
                    p0 -= p0
                case 32768:
                    var v7 int32 = 2147483647
                    _ = v7
                case -1:
                    var v8 int8 = -88
                    _ = v8
                }
            default:
                p0 = p0
            }
            var v9 uint = 929
            switch 5 {
            case 127:
                v9 += 18446744073709551615
            case 9223372036854775807:
                v10 := []int{fi5, fi5}
                _ = v10
            }
            _ = fi5
            _ = v9
        }
    case -462:
        var v11 uint16 = 801
        _ = v11
    case -9223372036854775808:
        p0 = p0
    default:
        v12 := GenS1{F0: (p0 / ((387) | 1)), F1: "foo bar", F2: 96}
        _ = v12
    }
    p0 = (-(p0))
    _ = p0
    return 0
}

func main() {
    var v1 int = 127
    v2 := []uint32{414, 2147483647, 42, 65535}
    var v3 string = ""
    var v4 int8 = 127
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        var v13 int64 = 1
        v3 = v3
        switch v1 {
        case 42:
            for fi14 := 0; fi14 < 16; fi14++ {
                v15 := GenS1{F0: -1, F1: (v3 + ""), F2: 255}
                _ = fi14
                _ = v15
            }
        case 0:
            v16 := GenS2{F0: 1}
            _ = v16
        case 1:
            v17 := GenS2{F0: 1}
            _ = v17
        }
        _ = v13
    } else {
        v4 -= v4
        v2 = append(v2, 2147483648)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, 989)
    var v18 uint16 = 2
    if (float64(0.1) >= 1e10) {
        v19 := GenS1{F0: 417, F1: ((("" + "a") + v3) + "foo bar"), F2: 2}
        v20 := GenS3{F0: 42, F1: (v18 >> ((0) & 63))}
        v21 := GenS4{F0: "hello", F1: v1, F2: 0}
        _ = v19
        _ = v20
        _ = v21
    } else {
        if (42 == 65535) {
            fmt.Println(v4, v1)
            if true {
                if true {
                    v1 -= -2147483648
                }
                v18 = 0
                switch (v3 + (("a" + "a") + (((v3 + "\n\t") + v3) + ("a" + ("\n\t" + v3))))) {
                case "a":
                    for fi22 := 0; fi22 < 7; fi22++ {
                        var v23 uint = 2
                        v24 := GenS3{F0: 1, F1: v18}
                        fmt.Println(v23)
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
            }
        } else {
            fmt.Println(v1)
            var v25 int = (v1 + -2147483648)
            if (true && (127 <= uint8(v18))) {
                if len(v2) > 0 {
                    v2[0] = 2
                }
                var v26 int32 = 42
                _ = v26
            }
            _ = v25
        }
    }
    var v27 float64 = (1e10 * ((-2.25 * 3.14159) - 1.5))
    v4 -= (v4 / (((v4 / ((127) | 1))) | 1))
    v1 -= (v1 & 261)
    if true {
        v28 := GenS5{F0: 728, F1: "\n\t", F2: 255}
        v1 += (v1 / (((v1 * v1)) | 1))
        for fi29 := 0; fi29 < 19; fi29++ {
            v1 -= (fi29 + 758)
            _ = fi29
        }
        _ = v28
    }
    if false {
        v2 = append(v2, 1)
    } else {
        v30 := GenS3{F0: 65535, F1: 32767}
        v18 = (^(v18))
        _ = v30
    }
    var v31 int64 = genFn0(-2147483648)
    v32 := GenS4{F0: ((v3 + (v3 + (v3 + ""))) + v3), F1: 32768, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
