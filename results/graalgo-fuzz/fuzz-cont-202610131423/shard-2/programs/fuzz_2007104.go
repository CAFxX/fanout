package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int32 }

func main() {
    var v1 string = ("hello" + "")
    v2 := make(map[string]uint32)
    var v3 uint8 = 127
    v4 := make(map[uint32]string)
    var v5 int = 9223372036854775807
    v6 := []uint32{42}
    var v7 bool = false
    switch v1 {
    case "":
        var v8 int = (v5 / ((v5) | 1))
        _ = v8
    case "a":
        var v9 bool = true
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 2; fi11++ {
                v5 += (fi10 + (fi11 % ((fi10) | 1)))
                v3 -= (v3 | v3)
                var v12 int32 = -2147483648
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    default:
        fmt.Println(v3)
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        if v7 {
            var v14 uint16 = 1
            _ = v14
        }
        fmt.Println(fi13)
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 uint64 = 1
            switch "\n\t" {
            case "hello":
                _ = len(v6)
            case "a":
                var v17 float32 = 3.14159
                _ = v17
            case "foo bar":
                switch "foo bar" {
                case "a":
                    v1 = v1
                case "foo bar":
                    if false {
                        switch v5 {
                        case 660:
                            v18 := v2[("foo bar" + v1)]
                            _ = v18
                        case -2147483648:
                            _ = len(v2)
                        }
                        if v7 {
                            var v19 uint = 0
                            _ = v19
                        }
                        var v20 uint64 = (v16 % ((4294967295) | 1))
                        _ = v20
                    }
                }
            default:
                delete(v2, (v1 + v1))
            }
            _ = fi15
            _ = v16
        }
        _ = fi13
    }
    v21 := GenS1{F0: (4294967295 != 18446744073709551615), F1: v3, F2: -265}
    var v22 uint8 = v3
    if (v1 != v1) {
        v6 = append(v6, 2147483647)
        var v23 uint32 = 65535
        _ = v23
    }
    v22 = (^(v3))
    var v24 string = v1
    if len(v6) > 0 {
        v6[0] = 726
    }
    var v25 int8 = -128
    switch (v5 - v5) {
    case -257:
        var v26 uint = 18446744073709551615
        _ = v26
    default:
        if true {
            v3 -= v22
            delete(v4, 2147483648)
            var v27 uint32 = 1
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
