package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int32; F1 int; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 bool }

func main() {
    var v1 int16 = 0
    var v2 bool = false
    v3 := make(map[int]bool)
    v4 := v3[42]
    v4 = v4
    fmt.Println(v2)
    switch 127 {
    case 128:
        if true {
            v2 = v4
        }
    case -31:
        v2 = (false && v2)
    case 42:
        v5 := v3[-2147483648]
        _ = v5
    default:
        if v2 {
            var v6 int = 128
            fmt.Println(v1, v2, v1)
            var v7 uint8 = 1
            _ = v6
            _ = v7
        } else {
            var v8 float32 = ((-2.25 / (0.0 + 1.5)) / 1.5)
            var v9 int64 = 42
            _ = v8
            _ = v9
        }
    }
    v10 := []int{0, 0, 127}
    switch "foo bar" {
    case "a":
        var v11 int16 = 42
        _ = v11
    case "\n\t":
        var v12 uint64 = 460
        _ = v12
    case "hello":
        if false {
            var v13 uint64 = 42
            _ = v13
        }
    default:
        fmt.Println(v2, v1, v4)
    }
    var v14 uint64 = 42
    v15 := GenS1{F0: 130, F1: ""}
    v1 += v1
    v4 = v2
    if v4 {
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 uint32 = 1
            _ = fi16
            _ = v17
        }
        if v4 {
            v18 := v3[42]
            _ = v18
        }
    } else {
        var v19 uint = 18446744073709551615
        v10 = append(v10, 32767)
        if true {
            switch 32767 {
            case -71:
                fmt.Println(v19, v1)
            }
            var v20 uint64 = v14
            _ = v20
        } else {
            var v21 uint16 = 2
            v1 = (v1 - v1)
            v19 -= (^(v19))
            _ = v21
        }
        _ = v19
    }
    if false {
        fmt.Println(v1, v1)
    } else {
        var v22 string = "\n\t"
        if true {
            if false {
                v23 := GenS1{F0: 42, F1: v22}
                switch -343 {
                case -562:
                    var v24 int = 42
                    _ = v24
                }
                v25 := GenS2{F0: 2147483647, F1: 128, F2: v22}
                _ = v23
                _ = v25
            }
            _ = len(v3)
        }
        if len(v10) > 0 {
            v10[0] = -859
        }
        _ = v22
    }
    if false {
        fmt.Println(v14, v1, v1)
    } else {
        if (42 == 4294967295) {
            fmt.Println(v4)
        }
        for fi26 := 0; fi26 < 2; fi26++ {
            switch "\n\t" {
            case "a":
                if v4 {
                    v27 := GenS3{F0: 1}
                    var v28 float32 = ((1e10 - 1.5) * -2.25)
                    v14 = 2
                    _ = v27
                    _ = v28
                } else {
                    var v29 bool = v2
                    fmt.Println(v4)
                    fmt.Println(fi26)
                    _ = v29
                }
            case "foo bar":
                v14 = 0
            case "":
                var v30 int = fi26
                _ = v30
            default:
                for fi31 := 0; fi31 < 5; fi31++ {
                    if len(v10) > 0 {
                        v10[0] = -2147483648
                    }
                    v32 := GenS1{F0: 127, F1: "a"}
                    switch fi31 {
                    case 128:
                        fmt.Println(fi26, v1)
                    case 9223372036854775807:
                        var v33 uint = 2
                        _ = v33
                    }
                    _ = fi31
                    _ = v32
                }
            }
            fmt.Println(v4, v4)
            _ = fi26
        }
    }
    if (v1 >= (v1 << ((65535) & 63))) {
        v10 = append(v10, -9223372036854775808)
        for fi34 := 0; fi34 < 1; fi34++ {
            v35 := GenS2{F0: -14, F1: fi34, F2: ""}
            if true {
                fmt.Println(v4)
                var v36 int64 = -1
                var v37 int = (fi34 % (((fi34 & fi34)) | 1))
                _ = v36
                _ = v37
            }
            _ = fi34
            _ = v35
        }
        if len(v10) > 0 {
            v10[0] = -1
        }
    } else {
        v2 = v2
        v2 = ((v1 != -1) && true)
    }
    v38 := GenS4{F0: 0, F1: (0.1 <= 3.14159)}
    v10 = append(v10, 194)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
