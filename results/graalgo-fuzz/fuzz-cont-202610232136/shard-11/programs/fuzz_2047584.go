package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int) string {
    p0 -= p0
    fmt.Println(p0)
    var v5 uint = 65535
    var v6 uint = v5
    v5 = (v6 + v6)
    _ = p0
    _ = v5
    _ = v6
    return "\n\t"
}
func genFn1(p0 int, p1 int32, p2 int8) uint16 {
    var v7 uint8 = 127
    fmt.Println(v7, p2)
    v8 := GenS1{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 282
}

func main() {
    var v1 int16 = 42
    var v2 uint64 = 4294967295
    var v3 uint = 18446744073709551615
    var v4 int64 = 42
    fmt.Println(v2)
    switch (("" + "\n\t") + "foo bar") {
    case "hello":
        var v9 uint32 = 2147483648
        _ = v9
    }
    v10 := []int{127, 625}
    switch ("a" + ("" + "hello")) {
    case "a":
        var v11 int16 = (v1 << ((v3) & 63))
        _ = v11
    }
    var v12 uint64 = v2
    fmt.Println(v1, v4)
    switch "a" {
    case "":
        for fi13 := 0; fi13 < 14; fi13++ {
            switch (("hello" + "") + ("a" + ("hello" + "a"))) {
            case "":
                if len(v10) > 0 {
                    v10[0] = fi13
                }
            case "hello":
                var v14 string = (((((("hello" + "a") + "") + "a") + ("foo bar" + (("\n\t" + "hello") + "\n\t"))) + "") + "")
                _ = v14
            default:
                fmt.Println(v12, v2)
            }
            v2 += 289
            for fi15 := 0; fi15 < 11; fi15++ {
                switch "" {
                case "":
                    if false {
                        if (false && true) {
                            switch fi13 {
                            case -1:
                                switch ("" + ("hello" + "")) {
                                case "foo bar":
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                default:
                                    v16 := GenS2{F0: fi15, F1: -1}
                                    _ = v16
                                }
                            }
                            v4 = 193
                        }
                    } else {
                        v4 -= v4
                        switch "foo bar" {
                        case "hello":
                            if ((!(false && false)) && false) {
                                var v17 uint8 = 255
                                v18 := GenS2{F0: 4294967295, F1: 609}
                                _ = v17
                                _ = v18
                            }
                        }
                        v19 := GenS3{F0: 255}
                        _ = v19
                    }
                case "hello":
                    v4 = v4
                case "\n\t":
                    var v20 string = "foo bar"
                    _ = v20
                default:
                    v2 = (v2 % ((v2) | 1))
                }
                _ = fi15
            }
            _ = fi13
        }
    case "foo bar":
        var v21 int64 = (v4 >> ((uint(v12)) & 63))
        _ = v21
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 int32 = -382
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            var v25 int32 = 1
            _ = fi23
            _ = v24
            _ = v25
        }
        if len(v10) > 0 {
            v10[0] = -754
        }
        var v26 int64 = v4
        _ = fi22
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
