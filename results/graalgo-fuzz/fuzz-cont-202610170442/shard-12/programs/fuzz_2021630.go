package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int8; F1 bool }

func main() {
    var v1 int16 = 42
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 int32 = 287
    v4 -= (v4 % ((v4) | 1))
    v5 := GenS1{F0: 18446744073709551615, F1: 963}
    v1 += (v1 - v1)
    if false {
        switch "foo bar" {
        case "a":
            var v6 int8 = 42
            _ = v6
        default:
            v7 := GenS2{F0: 0, F1: v1}
            _ = v7
        }
        v1 -= v1
    }
    if (20 != 42) {
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 11; fi9++ {
                v10 := GenS1{F0: 4294967295, F1: 778}
                v11 := GenS1{F0: 286, F1: 65535}
                var v12 uint32 = 2147483648
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            switch (-(fi8)) {
            case 2147483647:
                if true {
                    v2[("\n\t" + (((("hello" + "") + (("a" + ("\n\t" + (("\n\t" + ((("\n\t" + (("a" + "foo bar") + "a")) + ((("hello" + "foo bar") + "") + ("a" + ("foo bar" + "foo bar")))) + (((("a" + "") + "\n\t") + "\n\t") + "\n\t"))) + ""))) + (("hello" + "a") + "foo bar"))) + "a") + ("a" + ("foo bar" + ("" + ("hello" + "hello"))))))] = true
                    var v13 uint32 = 4294967295
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 7; fi14++ {
                        v4 -= v4
                        fmt.Println(v1, fi8, v1)
                        var v15 bool = true
                        _ = fi14
                        _ = v15
                    }
                }
            case 0:
                var v16 int8 = 0
                _ = v16
            case -2147483648:
                switch "hello" {
                case "\n\t":
                    fmt.Println(v4, v1)
                case "":
                    v17 := []int{(-(fi8)), -2147483648}
                    _ = v17
                }
            default:
                v18 := GenS3{F0: 0, F1: ("foo bar" < "a")}
                _ = v18
            }
            _ = fi8
        }
    }
    fmt.Println(v4)
    fmt.Println(v1)
    var v19 int16 = v1
    fmt.Println(v19, v19, v1)
    var v20 uint8 = 42
    if true {
        var v21 float64 = 1.5
        var v22 uint = 18446744073709551615
        _ = v21
        _ = v22
    } else {
        fmt.Println(v19)
    }
    fmt.Println(v1, v1, v20)
    fmt.Println(v20, v4, v20)
    v4 = 0
    if false {
        fmt.Println(v19, v1)
        v23 := GenS1{F0: 4294967295, F1: 516}
        _ = v23
    } else {
        var v24 uint32 = 2147483648
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
