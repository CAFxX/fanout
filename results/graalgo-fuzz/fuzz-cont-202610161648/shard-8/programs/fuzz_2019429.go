package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int; F1 bool; F2 bool }
func genFn0(p0 uint32, p1 uint, p2 uint8) int8 {
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{(fi7 % (((fi7 ^ -2147483648)) | 1))}
        fmt.Println(p2, p2, p1)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: "", F1: -199, F2: 459}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 127
}

func main() {
    var v1 string = "foo bar"
    v2 := []int{128, 2147483647}
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    var v5 int64 = 9223372036854775807
    var v6 string = "foo bar"
    if true {
        switch 4294967295 {
        case -9223372036854775808:
            delete(v4, "hello")
        case 0:
            var v10 int64 = (v5 % (((-(v5))) | 1))
            _ = v10
        }
    }
    var v11 int64 = (v5 | 0)
    switch 9223372036854775807 {
    case -2147483648:
        var v12 uint8 = 42
        _ = v12
    case -1:
        switch ("a" + v6) {
        case "a":
            v5 = -9223372036854775808
        case "hello":
            v11 += (v5 % ((v11) | 1))
        case "":
            if (v11 <= (^(v11))) {
                v1 = (("hello" + v1) + "hello")
            } else {
                if len(v2) > 0 {
                    v2[0] = 128
                }
                v13 := GenS2{F0: 107, F1: 65535, F2: 0}
                switch v6 {
                case "":
                    fmt.Println(v6, v6)
                default:
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v15 := GenS3{F0: -132, F1: (!false), F2: (!false)}
                        _ = fi14
                        _ = v15
                    }
                }
                _ = v13
            }
        default:
            fmt.Println(v11, v6, v11)
        }
    case 128:
        _ = len(v4)
    default:
        v11 = 0
    }
    switch -9223372036854775808 {
    case 1:
        var v16 uint16 = 0
        _ = v16
    }
    v17 := GenS2{F0: 4294967295, F1: 42, F2: 42}
    switch -2147483648 {
    case -769:
        v6 = v6
    }
    v18 := GenS2{F0: 1, F1: 42, F2: 311}
    var v19 int16 = 993
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    var v20 float32 = 0.1
    var v21 uint8 = 127
    switch 2147483647 {
    case 32767:
        switch v1 {
        case "a":
            v6 = ("\n\t" + "a")
        }
    case 128:
        v19 += (v19 * v19)
    case 1:
        for fi22 := 0; fi22 < 15; fi22++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v23 := GenS1{F0: "foo bar", F1: (fi22 ^ (fi22 << ((18446744073709551615) & 63))), F2: 1}
            _ = fi22
            _ = v23
        }
    default:
        delete(v3, -2147483648)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
