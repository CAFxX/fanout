package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint; F2 bool }
func genFn0(p0 uint32, p1 int16) uint16 {
    p0 -= (p0 >> ((2) & 63))
    switch -9223372036854775808 {
    case 127:
        v4 := []int{-299, -124}
        _ = v4
    default:
        fmt.Println(p1, p1)
    }
    for fi5 := 0; fi5 < 5; fi5++ {
        if true {
            var v6 float64 = 0.1
            for fi7 := 0; fi7 < 6; fi7++ {
                fmt.Println(fi5, fi5, fi5)
                p0 = p0
                v8 := GenS1{F0: 32, F1: 42}
                p0 -= (p0 - 1)
                var v9 int8 = 127
                var v10 float32 = -2.25
                var v11 string = "foo bar"
                v12 := GenS1{F0: 2, F1: -2147483648}
                _ = fi7
                _ = v10
                _ = v11
                _ = v12
                _ = v8
                _ = v9
            }
            p1 = p1
            _ = v6
        }
        v13 := GenS2{F0: 2}
        _ = fi5
        _ = v13
    }
    p1 += p1
    if (!((true || (4294967295 > 2)) && true)) {
        var v14 int16 = p1
        var v15 uint64 = 0
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    return 255
}
func genFn1(p0 int64, p1 uint, p2 uint) bool {
    switch "a" {
    case "foo bar":
        v16 := GenS3{F0: 4294967295, F1: p1, F2: ("hello" == "foo bar")}
        _ = v16
    }
    v17 := []int{}
    var v18 uint64 = 4294967295
    _ = len(v17)
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return ((^(v18)) >= (v18 * (v18 * (^(v18)))))
}

func main() {
    v1 := []int{}
    v2 := []string{(("\n\t" + "hello") + "foo bar"), ""}
    var v3 string = (((("\n\t" + "hello") + "\n\t") + "foo bar") + ("a" + ("hello" + "foo bar")))
    v19 := GenS1{F0: 18446744073709551615, F1: -983}
    for fi20 := 0; fi20 < 5; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            v2 = append(v2, (((v3 + ((("a" + (v3 + (v3 + (v3 + v3)))) + "") + v3)) + ((v3 + (v3 + (v3 + v3))) + (v3 + ""))) + "foo bar"))
            v3 = "a"
            _ = fi21
        }
        _ = fi20
    }
    v3 = ""
    v1 = append(v1, 502)
    var v22 uint16 = 390
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    v23 := GenS3{F0: 2147483647, F1: 592, F2: ((v22 * v22) < v22)}
    switch v3 {
    case "foo bar":
        v2 = append(v2, v3)
    default:
        for fi24 := 0; fi24 < 8; fi24++ {
            v3 = v3
            fmt.Println(v3, v22)
            _ = fi24
        }
    }
    if ((-2147483648 != 356) || (4294967295 >= 2)) {
        if (2 != 2147483647) {
            v3 = ("foo bar" + (v3 + "hello"))
        }
    }
    fmt.Println(v3, v22)
    var v25 uint8 = 2
    v25 = (^(v25))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
}
