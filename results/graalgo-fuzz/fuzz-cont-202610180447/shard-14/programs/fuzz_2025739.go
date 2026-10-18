package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 uint64 = 42
    var v2 string = "\n\t"
    var v3 float64 = 3.14159
    var v4 uint = 765
    for fi5 := 0; fi5 < 9; fi5++ {
        if (100 >= 127) {
            v6 := GenS1{F0: ("a" + ((((((v2 + "") + v2) + (("foo bar" + ((v2 + v2) + v2)) + (((((v2 + (v2 + "\n\t")) + ("hello" + v2)) + (v2 + "hello")) + "") + v2))) + (v2 + v2)) + v2) + v2)), F1: 922}
            fmt.Println(v3)
            _ = v6
        } else {
            v7 := []int{1}
            _ = v7
        }
        var v8 uint = v4
        _ = fi5
        _ = v8
    }
    fmt.Println(v2, v4)
    var v9 uint64 = 913
    v4 += v4
    switch -248 {
    case 152:
        var v10 int32 = -2147483648
        _ = v10
    case -9223372036854775808:
        v11 := []int{}
        _ = v11
    case 127:
        var v12 float64 = v3
        _ = v12
    }
    v13 := []int{127}
    if false {
        if true {
            v13 = append(v13, 128)
        } else {
            fmt.Println(v4)
            if true {
                fmt.Println(v9, v2)
            } else {
                switch v2 {
                case "foo bar":
                    for fi14 := 0; fi14 < 12; fi14++ {
                        v15 := GenS2{F0: 1}
                        _ = fi14
                        _ = v15
                    }
                default:
                    v16 := GenS3{F0: v1, F1: 32767, F2: 1}
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                v1 = (v9 << ((v4) & 63))
                for fi18 := 0; fi18 < 11; fi18++ {
                    if true {
                        fmt.Println(fi17, v2)
                        v3 = 1e10
                        var v19 uint64 = v1
                        v9 = v19
                        fmt.Println(v1, v19, v2)
                        _ = v19
                    }
                    var v20 int16 = -1
                    fmt.Println(v4, v1, v20)
                    fmt.Println(v1, v1)
                    v9 = 18446744073709551615
                    _ = fi18
                    _ = v20
                }
                v13 = append(v13, -144)
                _ = fi17
            }
        }
        _ = len(v13)
    }
    v3 = -2.25
    _ = len(v13)
    switch ((v2 + "a") + v2) {
    case "a":
        if false {
            for fi21 := 0; fi21 < 9; fi21++ {
                var v22 int8 = 127
                var v23 int32 = -1
                _ = fi21
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 17; fi24++ {
                var v25 int64 = 502
                _ = fi24
                _ = v25
            }
        }
    case "hello":
        v26 := GenS3{F0: (v1 * 18446744073709551615), F1: -2147483648, F2: 51}
        _ = v26
    default:
        v27 := GenS1{F0: v2, F1: -32768}
        _ = v27
    }
    var v28 float64 = (v3 / 1.5)
    v29 := GenS4{F0: 42, F1: 1}
    var v30 bool = true
    var v31 int8 = -128
    v13 = append(v13, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v9
}
