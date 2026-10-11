package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }

func main() {
    var v1 uint8 = 74
    var v2 float32 = 1.5
    var v3 uint = 42
    var v4 float32 = (float32(3.14159) + float32(3.14159))
    v3 = (v3 - (v3 | (v3 ^ v3)))
    for fi5 := 0; fi5 < 18; fi5++ {
        v2 -= v4
        v2 += v4
        var v6 uint16 = 852
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint8 = (v1 + v1)
        _ = fi7
        _ = v8
    }
    if ("foo bar" <= "hello") {
        fmt.Println(v1)
        for fi9 := 0; fi9 < 13; fi9++ {
            v1 -= v1
            _ = fi9
        }
    } else {
        v10 := []int{9223372036854775807, 762}
        _ = v10
    }
    v11 := GenS1{F0: 2147483647, F1: -1, F2: "\n\t"}
    var v12 bool = ((!true) && false)
    v13 := GenS1{F0: 1, F1: 0, F2: "hello"}
    var v14 float32 = v4
    for fi15 := 0; fi15 < 18; fi15++ {
        v12 = (false || v12)
        _ = fi15
    }
    if (42 == 1) {
        v16 := []int{}
        v17 := GenS1{F0: 267, F1: 2147483647, F2: (("a" + "hello") + "hello")}
        if true {
            var v18 int64 = 990
            for fi19 := 0; fi19 < 7; fi19++ {
                if len(v16) > 0 {
                    v16[0] = fi19
                }
                v14 += 0.1
                if len(v16) > 0 {
                    v16[0] = (fi19 / (((-(fi19))) | 1))
                }
                _ = fi19
            }
            _ = v18
        } else {
            for fi20 := 0; fi20 < 12; fi20++ {
                for fi21 := 0; fi21 < 4; fi21++ {
                    v22 := GenS1{F0: 816, F1: 2147483647, F2: "hello"}
                    switch ("\n\t" + "") {
                    case "foo bar":
                        v2 = v4
                    }
                    switch "\n\t" {
                    case "":
                        v23 := GenS1{F0: 2147483648, F1: 942, F2: "a"}
                        _ = v23
                    case "\n\t":
                        var v24 int64 = -667
                        _ = v24
                    default:
                        v12 = (42 >= 423)
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
            var v25 uint64 = 18446744073709551615
            var v26 float64 = 1.5
            _ = v25
            _ = v26
        }
        _ = v16
        _ = v17
    } else {
        fmt.Println(v4)
        v1 -= v1
    }
    var v27 string = ""
    fmt.Println(v1, v4, v27)
    switch v27 {
    case "a":
        v14 = v4
    case "foo bar":
        v2 += 1.5
    case "":
        fmt.Println(v14, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
