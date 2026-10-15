package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool; F1 string; F2 int32 }

func main() {
    var v1 int32 = 1
    var v2 int8 = 127
    var v3 uint32 = 2147483647
    v4 := make(map[string]int)
    v3 += (v3 / (((^(v3))) | 1))
    switch "\n\t" {
    case "hello":
        for fi5 := 0; fi5 < 4; fi5++ {
            v3 += v3
            if ((v3 * v3) != 65535) {
                fmt.Println(v3)
            } else {
                if ((((1.5 + ((1.5 + 0.1) - 1e10)) * 0.0) * -2.25) < 3.14159) {
                    for fi6 := 0; fi6 < 15; fi6++ {
                        switch fi5 {
                        case 498:
                            var v7 float64 = 1e10
                            _ = v7
                        case -9223372036854775808:
                            var v8 bool = true
                            _ = v8
                        }
                        v3 += v3
                        _ = fi6
                    }
                }
            }
            var v9 uint = 18446744073709551615
            _ = fi5
            _ = v9
        }
    case "a":
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: 651, F1: (v3 / (((v3 - (v3 + v3))) | 1))}
            v12 := []int{}
            _ = fi10
            _ = v11
            _ = v12
        }
    default:
        v13 := []int{4294967295}
        _ = v13
    }
    var v14 int64 = 648
    v15 := v4[(("" + "foo bar") + (("hello" + "") + ("hello" + "hello")))]
    switch v15 {
    case 0:
        fmt.Println(v3, v1)
    case 2147483647:
        var v16 uint16 = 0
        _ = v16
    case 32767:
        if (127 > (v15 * v15)) {
            fmt.Println(v14, v2)
            var v17 int = v15
            _ = v17
        }
    default:
        if ((!(2 >= 1)) || (42 <= v3)) {
            v18 := GenS2{F0: false, F1: "hello", F2: v1}
            var v19 int8 = 0
            _ = v18
            _ = v19
        } else {
            v20 := GenS2{F0: false, F1: (("foo bar" + "a") + "foo bar"), F2: v1}
            _ = v20
        }
    }
    fmt.Println(v2)
    var v21 int64 = v14
    var v22 int = v15
    switch v15 {
    case 127:
        switch "foo bar" {
        case "a":
            _ = len(v4)
        case "hello":
            var v23 int64 = (-(v21))
            _ = v23
        }
    default:
        fmt.Println(v2, v3)
    }
    switch "hello" {
    case "\n\t":
        fmt.Println(v3, v3)
    case "":
        var v24 uint16 = 42
        _ = v24
    case "a":
        v25 := GenS2{F0: (v3 == (v3 ^ (v3 + 2147483648))), F1: "foo bar", F2: 1}
        _ = v25
    }
    var v26 uint8 = 2
    switch (^(v15)) {
    case 1:
        switch 9223372036854775807 {
        case -363:
            fmt.Println(v15, v3)
        }
    default:
        v22 = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
}
