package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 bool; F1 int16; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint64; F2 bool }
type GenS6 struct { F0 bool; F1 int }
func genFn0() int32 {
    v5 := GenS1{F0: 18446744073709551615, F1: -1, F2: 2147483648}
    var v6 int = 50
    switch (("\n\t" + "foo bar") + "foo bar") {
    case "a":
        if (!(663 >= 65535)) {
            v6 = (v6 - 0)
            switch "\n\t" {
            case "\n\t":
                v7 := GenS1{F0: 2, F1: -1, F2: 2147483647}
                _ = v7
            case "hello":
                v8 := GenS1{F0: 763, F1: 42, F2: 138}
                _ = v8
            case "a":
                if (!true) {
                    var v9 float32 = float32(0.0)
                    _ = v9
                } else {
                    fmt.Println(v6)
                }
            }
            v10 := GenS2{F0: (false && (!false)), F1: 1, F2: v6}
            _ = v10
        } else {
            v11 := []int{2147483647, (v6 & 818), (v6 >> ((4294967295) & 63))}
            var v12 int = (-(v6))
            fmt.Println(v6, v6)
            _ = v11
            _ = v12
        }
    case "\n\t":
        v13 := GenS3{F0: 0}
        _ = v13
    }
    v6 = v6
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{42, 2}
    var v3 float64 = 0.1
    var v4 uint = 4294967295
    var v14 string = ((((("hello" + ("\n\t" + "")) + "\n\t") + ("a" + "")) + "\n\t") + ((("\n\t" + "\n\t") + "foo bar") + "\n\t"))
    var v15 int8 = 100
    v15 = v15
    for fi16 := 0; fi16 < 11; fi16++ {
        v4 = (v4 - 42)
        var v17 uint64 = 42
        var v18 bool = true
        _ = fi16
        _ = v17
        _ = v18
    }
    switch v14 {
    case "foo bar":
        var v19 uint64 = 121
        _ = v19
    case "a":
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS4{F0: (v14 + "foo bar")}
            var v22 int = fi20
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v1[-9223372036854775808] = 32767
    if (v4 != v4) {
        if false {
            v3 -= (-2.25 - v3)
        } else {
            _ = len(v2)
        }
        v23 := GenS5{F0: 51, F1: 18446744073709551615, F2: (42 > -37)}
        _ = v23
    }
    if false {
        var v24 uint16 = 65535
        _ = v24
    } else {
        var v25 uint32 = 42
        _ = len(v1)
        _ = v25
    }
    v2 = append(v2, 42)
    v26 := GenS6{F0: (!((!((-60 != -32768) && (-721 < 32768))) && true)), F1: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
