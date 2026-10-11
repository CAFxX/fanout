package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0(p0 int32, p1 int64) int {
    v6 := []int{2147483647}
    var v7 uint = 0
    if ((987 < 4294967295) && true) {
        var v8 float32 = 0.1
        _ = v8
    } else {
        p1 = p1
        v9 := GenS1{F0: "a", F1: 65535}
        switch (("hello" + "foo bar") + "foo bar") {
        case "\n\t":
            fmt.Println(v7, p0, v7)
        case "":
            v10 := GenS1{F0: ("a" + "hello"), F1: 345}
            _ = v10
        case "foo bar":
            for fi11 := 0; fi11 < 15; fi11++ {
                v12 := GenS1{F0: (((("hello" + "\n\t") + "foo bar") + "a") + ("" + "")), F1: 0}
                _ = fi11
                _ = v12
            }
        default:
            if (true || false) {
                switch "\n\t" {
                case "\n\t":
                    switch ((("" + ("\n\t" + (("foo bar" + (("" + "hello") + "\n\t")) + ("foo bar" + "hello")))) + ((("" + "hello") + ("hello" + "hello")) + "")) + "hello") {
                    case "foo bar":
                        fmt.Println(v7)
                    }
                case "":
                    for fi13 := 0; fi13 < 14; fi13++ {
                        v14 := GenS2{F0: false}
                        _ = fi13
                        _ = v14
                    }
                case "hello":
                    p0 = (p0 | p0)
                }
                if len(v6) > 0 {
                    v6[0] = -9223372036854775808
                }
                fmt.Println(v7, p1, p1)
            } else {
                for fi15 := 0; fi15 < 8; fi15++ {
                    if len(v6) > 0 {
                        v6[0] = (fi15 ^ (^(fi15)))
                    }
                    _ = fi15
                }
                switch "foo bar" {
                case "a":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    v7 = (^(v7))
                }
            }
        }
        _ = v9
    }
    if (!(0 == 40)) {
        var v16 int16 = -32768
        if true {
            var v17 uint32 = 2
            _ = v17
        } else {
            var v18 int16 = v16
            v16 = v18
            _ = v18
        }
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 750
}
func genFn1(p0 string) int16 {
    fmt.Println(p0)
    switch p0 {
    case "":
        if false {
            switch p0 {
            case "hello":
                var v19 int64 = -9223372036854775808
                _ = v19
            default:
                var v20 int64 = -1
                _ = v20
            }
            var v21 int8 = -1
            _ = v21
        }
    default:
        var v22 float64 = 3.14159
        _ = v22
    }
    if (((-1 <= 99) && (!false)) && false) {
        var v23 int8 = 0
        _ = v23
    }
    _ = p0
    return 42
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[string]string)
    var v3 uint8 = 0
    v4 := []int{2147483647}
    var v5 uint16 = 65535
    switch 128 {
    case 2147483647:
        var v24 int32 = -1
        _ = v24
    default:
        var v25 bool = (!(42 >= 1))
        _ = v25
    }
    var v26 uint16 = uint16(v1)
    var v27 uint8 = 127
    fmt.Println(v1)
    if (403 > 552) {
        v26 -= 42
        for fi28 := 0; fi28 < 8; fi28++ {
            v27 = (v3 - v27)
            _ = fi28
        }
    }
    if true {
        fmt.Println(v27, v1)
        switch ("a" + "hello") {
        case "hello":
            v1 -= (^(v1))
        case "":
            if len(v4) > 0 {
                v4[0] = 32768
            }
        }
    }
    var v29 int8 = -120
    if true {
        var v30 float64 = -2.25
        _ = v30
    }
    v26 = (^(v26))
    var v31 int = 2147483647
    var v32 uint64 = uint64(v26)
    var v33 uint8 = v27
    v26 += v5
    var v34 int8 = -8
    fmt.Println(v32, v34)
    v35 := GenS3{F0: (v26 + 255), F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
