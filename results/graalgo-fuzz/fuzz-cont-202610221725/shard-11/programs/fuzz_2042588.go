package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint32 }

func main() {
    v1 := []uint32{439, 1, 2}
    var v2 uint = 65535
    v3 := make(map[string]uint32)
    var v4 int32 = -1
    v2 = 2
    for fi5 := 0; fi5 < 8; fi5++ {
        v4 = v4
        var v6 int = (^(fi5))
        var v7 uint8 = 128
        _ = fi5
        _ = v6
        _ = v7
    }
    switch ("\n\t" + "hello") {
    case "hello":
        var v8 int64 = 1
        _ = v8
    case "a":
        if (((-(v4)) >= (v4 << ((v2) & 63))) || (("hello" + ("\n\t" + "hello")) == "a")) {
            v3[(("" + ("foo bar" + "")) + "\n\t")] = 300
            var v9 int16 = 1000
            switch int(v4) {
            case -886:
                if false {
                    switch -1 {
                    case 127:
                        var v10 int8 = 100
                        _ = v10
                    }
                }
            case 128:
                v4 = v4
            case -893:
                var v11 int64 = 0
                _ = v11
            default:
                var v12 uint32 = 200
                _ = v12
            }
            _ = v9
        } else {
            v4 = (^(v4))
        }
    default:
        v2 += (v2 % (((v2 << ((225) & 63))) | 1))
    }
    v13 := GenS1{F0: 1, F1: 24, F2: 1}
    switch ("" + ("foo bar" + "foo bar")) {
    case "foo bar":
        v4 += 2147483647
    case "hello":
        if true {
            v4 = v4
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            if true {
                if (1.5 == float32((0.1 + 1e10))) {
                    v14 := GenS1{F0: -52, F1: -128, F2: 2147483648}
                    _ = v14
                }
            } else {
                v4 += (v4 & v4)
            }
        }
    case "\n\t":
        if true {
            v4 -= v4
        }
    default:
        fmt.Println(v4, v4)
    }
    fmt.Println(v2)
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS1{F0: 35, F1: 0, F2: 42}
        fmt.Println(v4, fi15, v2)
        _ = fi15
        _ = v16
    }
    v2 += 65535
    v17 := GenS1{F0: 42, F1: 127, F2: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
