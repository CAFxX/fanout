package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS4 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 1000
    v3 := []uint32{}
    var v4 uint16 = 2
    var v5 float64 = (1e10 - -2.25)
    fmt.Println(v5, v4)
    fmt.Println(v2, v5, v5)
    if (!((!((v2 | (v2 / (((^(v2))) | 1))) < (v2 ^ 1000))) || false)) {
        v4 -= 32767
    }
    fmt.Println(v5, v2)
    switch "" {
    case "a":
        v5 = float64(1.5)
    case "hello":
        for fi6 := 0; fi6 < 18; fi6++ {
            _ = len(v1)
            _ = fi6
        }
    case "":
        v1[("foo bar" + "\n\t")] = ("" + "\n\t")
    }
    switch "a" {
    case "":
        var v7 float32 = (1e10 - 1.5)
        _ = v7
    case "foo bar":
        if true {
            v8 := GenS1{F0: 1, F1: 811}
            _ = v8
        }
    default:
        switch "\n\t" {
        case "foo bar":
            v4 = v4
        default:
            v2 += v2
        }
    }
    var v9 bool = ((-2147483648 == -922) && true)
    v9 = true
    v10 := GenS2{F0: 0}
    var v11 bool = (true || v9)
    for fi12 := 0; fi12 < 10; fi12++ {
        switch ("" + "foo bar") {
        case "foo bar":
            v2 = v2
        default:
            for fi13 := 0; fi13 < 2; fi13++ {
                delete(v1, "a")
                _ = fi13
            }
        }
        _ = fi12
    }
    v2 = (v2 % ((957) | 1))
    if v11 {
        if v11 {
            v9 = ((v2 < (v2 / ((0) | 1))) || false)
        } else {
            var v14 bool = v9
            _ = len(v1)
            v5 = v5
            _ = v14
        }
        v4 -= 2
        if (0 > 717) {
            v15 := v1["a"]
            if false {
                if ("" != (v15 + "foo bar")) {
                    if v9 {
                        v16 := GenS1{F0: 4294967295, F1: 4294967295}
                        v17 := GenS3{F0: 4294967295, F1: 206, F2: 42}
                        fmt.Println(v5, v9)
                        _ = v16
                        _ = v17
                    }
                } else {
                    var v18 int32 = 1
                    _ = v18
                }
                var v19 string = (v15 + ("hello" + "hello"))
                _ = v19
            }
            _ = v15
        }
    } else {
        switch 9223372036854775807 {
        case 42:
            switch "a" {
            case "a":
                for fi20 := 0; fi20 < 2; fi20++ {
                    fmt.Println(v2)
                    var v21 int16 = (v2 | (v2 * -38))
                    _ = fi20
                    _ = v21
                }
            default:
                v22 := GenS3{F0: 2147483648, F1: 973, F2: 1}
                _ = v22
            }
        case 127:
            v23 := v1[("foo bar" + "\n\t")]
            _ = v23
        case 0:
            v24 := v1["hello"]
            _ = v24
        }
        if v9 {
            var v25 uint64 = 42
            v25 = v25
            _ = v25
        } else {
            for fi26 := 0; fi26 < 3; fi26++ {
                switch "" {
                case "\n\t":
                    v27 := GenS4{F0: -6, F1: 2}
                    _ = v27
                case "hello":
                    fmt.Println(v5, v11, fi26)
                default:
                    _ = len(v3)
                }
                _ = fi26
            }
        }
        v28 := GenS4{F0: 126, F1: 65535}
        _ = v28
    }
    fmt.Println(v11)
    v3 = append(v3, 1)
    v9 = (v11 && v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
