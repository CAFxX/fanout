package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 string }
type GenS2 struct { F0 int8 }
func genFn0(p0 int8, p1 int16) int8 {
    p0 -= 38
    switch "foo bar" {
    case "foo bar":
        v7 := []int{-932, -2147483648}
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 int16 = p1
            _ = fi8
            _ = v9
        }
    case "a":
        v10 := GenS1{F0: 1, F1: ("foo bar" + ("a" + ((("hello" + ("" + "foo bar")) + "foo bar") + "hello"))), F2: ("hello" + "\n\t")}
        _ = v10
    default:
        p0 += p0
    }
    if ((p0 % (((-(p0))) | 1)) >= p0) {
        var v11 float64 = 1e10
        v12 := []int{}
        _ = v11
        _ = v12
    } else {
        var v13 int8 = (p0 * (p0 ^ 0))
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 int32 = 2147483647
            v16 := []int{(fi14 + -850), 0, 4294967295}
            v15 -= 2147483647
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    }
    _ = p0
    _ = p1
    return (p0 & -1)
}
func genFn1(p0 string) int16 {
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := GenS1{F0: 0, F1: "a", F2: p0}
        _ = fi17
        _ = v18
    }
    var v19 uint32 = 4294967295
    _ = p0
    _ = v19
    return -32768
}

func main() {
    var v1 int16 = 42
    var v2 int = 769
    var v3 float32 = (0.0 / -2.25)
    var v4 uint32 = 943
    var v5 string = ("hello" + ("foo bar" + ("foo bar" + (("foo bar" + ("a" + (("a" + "a") + ("a" + ("\n\t" + "a"))))) + "hello"))))
    v6 := make(map[int]string)
    fmt.Println(v2, v4, v1)
    switch v5 {
    case "\n\t":
        v6[v2] = v5
    case "foo bar":
        if false {
            v2 += (v2 / (((^(v2))) | 1))
            var v20 int32 = -883
            _ = v20
        } else {
            switch (-(v2)) {
            case 32768:
                var v21 int64 = 0
                _ = v21
            case 9223372036854775807:
                fmt.Println(v3, v2, v3)
            default:
                v2 = (v2 | (v2 << ((2) & 63)))
            }
            v4 = (v4 ^ (v4 % (((v4 * (v4 | 150))) | 1)))
            var v22 uint = 18446744073709551615
            _ = v22
        }
    case "hello":
        v1 -= genFn1(v5)
    }
    var v23 bool = (((!((false && ((v3 / (-2.25 * ((float32((3.14159 - 3.14159)) - (v3 / 0.1)) - v3))) == 1e10)) && (v5 <= v5))) && true) || (v5 >= (v5 + ("\n\t" + ("" + "hello")))))
    var v24 float64 = 1.5
    var v25 uint32 = (v4 >> ((4294967295) & 63))
    v26 := GenS2{F0: -5}
    fmt.Println(v24, v4)
    var v27 int32 = -582
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
