package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int64 }

func main() {
    var v1 int32 = 780
    var v2 int16 = 32767
    var v3 uint64 = 1
    v4 := make(map[string]int)
    var v5 int8 = 42
    switch "hello" {
    case "a":
        if ((-2147483648 < v1) || (float64(3.14159) <= (float64(1e10) / -2.25))) {
            for fi6 := 0; fi6 < 14; fi6++ {
                fmt.Println(v1, v5)
                v3 = v3
                v7 := GenS1{F0: (v1 * 0), F1: 0}
                _ = fi6
                _ = v7
            }
            v2 = (v2 + v2)
            v8 := GenS2{F0: 42, F1: 9223372036854775807}
            _ = v8
        } else {
            switch 9223372036854775807 {
            case -537:
                v9 := []int{-982, 4294967295, -617, 9223372036854775807}
                _ = v9
            default:
                if (true && false) {
                    var v10 uint32 = 2147483648
                    _ = v10
                } else {
                    fmt.Println(v2, v2, v2)
                }
            }
            var v11 uint32 = 2
            _ = v11
        }
    case "hello":
        for fi12 := 0; fi12 < 20; fi12++ {
            v3 -= (v3 / ((v3) | 1))
            v13 := GenS3{F0: 4294967295}
            var v14 float32 = 3.14159
            _ = fi12
            _ = v13
            _ = v14
        }
    case "foo bar":
        for fi15 := 0; fi15 < 19; fi15++ {
            _ = len(v4)
            v16 := GenS4{F0: v3, F1: 42}
            _ = fi15
            _ = v16
        }
    }
    switch ("hello" + ("" + "foo bar")) {
    case "\n\t":
        v2 -= (v2 ^ v2)
    case "a":
        v17 := GenS4{F0: v3, F1: 1}
        _ = v17
    default:
        switch 2147483647 {
        case 4294967295:
            var v18 int8 = (v5 ^ (v5 % ((-48) | 1)))
            _ = v18
        }
    }
    v1 += (v1 * 1)
    v3 += 115
    v19 := []int{0, 128}
    var v20 int32 = -150
    if (42 == 2) {
        if (!(!false)) {
            switch -456 {
            case 42:
                switch ("foo bar" + (("\n\t" + ("" + "")) + "\n\t")) {
                case "hello":
                    v3 += (^(v3))
                case "a":
                    switch "a" {
                    case "foo bar":
                        v1 += (v20 + v1)
                    default:
                        switch 42 {
                        case -481:
                            delete(v4, ("a" + "hello"))
                        default:
                            v21 := GenS2{F0: 4294967295, F1: 9223372036854775807}
                            _ = v21
                        }
                    }
                default:
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v5, v1, v20)
                    case "":
                        v1 += (v1 / ((2147483647) | 1))
                    case "a":
                        var v22 int32 = v20
                        _ = v22
                    default:
                        fmt.Println(v1, v20, v3)
                    }
                }
            default:
                switch "\n\t" {
                case "hello":
                    var v23 float64 = 0.1
                    _ = v23
                case "foo bar":
                    var v24 uint16 = 1
                    _ = v24
                }
            }
        }
        v2 = v2
        fmt.Println(v1, v3)
    } else {
        fmt.Println(v2, v2)
    }
    switch ((("hello" + ("a" + ("hello" + "hello"))) + (("hello" + ("hello" + "a")) + "a")) + ("\n\t" + "hello")) {
    case "a":
        var v25 int16 = 1000
        _ = v25
    case "":
        for fi26 := 0; fi26 < 9; fi26++ {
            v4[("a" + "\n\t")] = (fi26 + 0)
            _ = fi26
        }
    case "foo bar":
        var v27 int = 33
        _ = v27
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        _ = len(v4)
        fmt.Println(v20, v3)
        _ = fi28
    }
    fmt.Println(v3)
    v3 = (^(v3))
    var v29 int16 = v2
    for fi30 := 0; fi30 < 4; fi30++ {
        if true {
            switch 32767 {
            case 9223372036854775807:
                v2 += (v29 ^ (v2 ^ (^(v29))))
            case -1:
                v31 := GenS3{F0: 843}
                _ = v31
            }
            delete(v4, "foo bar")
            var v32 int8 = 89
            _ = v32
        }
        var v33 uint32 = 2
        _ = fi30
        _ = v33
    }
    fmt.Println(v5, v1, v2)
    for fi34 := 0; fi34 < 13; fi34++ {
        v35 := GenS2{F0: 4294967295, F1: -131}
        fmt.Println(v29)
        _ = fi34
        _ = v35
    }
    v36 := GenS5{F0: -9223372036854775808}
    v37 := GenS4{F0: (v3 << ((955) & 63)), F1: 32767}
    v38 := GenS2{F0: 706, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
