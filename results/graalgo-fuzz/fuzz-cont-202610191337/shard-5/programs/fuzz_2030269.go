package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int; F1 uint8 }

func main() {
    var v1 int16 = 1
    var v2 int = 42
    v3 := []string{("hello" + "hello"), "", (("a" + "\n\t") + "hello")}
    v4 := []uint32{0, 0, 2147483648, 2147483647}
    v5 := make(map[int]string)
    _ = len(v5)
    var v6 int = v2
    switch "a" {
    case "hello":
        fmt.Println(v2)
    }
    v6 -= (v6 - -2147483648)
    v6 += v2
    if (v2 != (^(v6))) {
        var v7 int8 = -113
        v7 -= (^(v7))
        _ = v7
    } else {
        fmt.Println(v1)
        v8 := GenS1{F0: (^(v6)), F1: 341}
        for fi9 := 0; fi9 < 6; fi9++ {
            if ((!(!(322 == 727))) && (!(100 < 1))) {
                for fi10 := 0; fi10 < 16; fi10++ {
                    v1 = (v1 << ((4294967295) & 63))
                    var v11 uint64 = 1
                    var v12 uint8 = 2
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                fmt.Println(v1, fi9)
            } else {
                if (-2.25 < 1.5) {
                    switch "" {
                    case "foo bar":
                        for fi13 := 0; fi13 < 11; fi13++ {
                            v14 := GenS2{F0: (fi9 / ((fi9) | 1)), F1: 255}
                            var v15 float64 = (float64((3.14159 - -2.25)) * -2.25)
                            v15 += (0.0 - (0.0 - 1e10))
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    }
                    v2 -= fi9
                }
                v6 = -2147483648
            }
            if true {
                v1 -= (v1 << ((464) & 63))
                if false {
                    if true {
                        v4 = append(v4, 2147483647)
                    }
                }
            }
            switch "a" {
            case "":
                if (42 > 128) {
                    var v16 float32 = (0.0 - 1.5)
                    _ = v16
                } else {
                    v4 = append(v4, 4294967295)
                }
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi9
        }
        _ = v8
    }
    delete(v5, 9223372036854775807)
    delete(v5, 571)
    switch (("" + ("hello" + "")) + "foo bar") {
    case "":
        var v17 int32 = 2147483647
        _ = v17
    case "foo bar":
        fmt.Println(v1, v2, v2)
    case "hello":
        v18 := GenS1{F0: (v2 << ((6) & 63)), F1: 1}
        _ = v18
    default:
        for fi19 := 0; fi19 < 16; fi19++ {
            v1 -= 42
            v6 -= (v2 * (v6 | 1))
            if true {
                if false {
                    var v20 bool = (!false)
                    fmt.Println(v2)
                    v2 = 4294967295
                    _ = v20
                }
            } else {
                if (1e10 == 1e10) {
                    switch -438 {
                    case 978:
                        _ = len(v5)
                    default:
                        var v21 uint = 438
                        _ = v21
                    }
                    var v22 uint64 = 505
                    _ = v22
                }
            }
            _ = fi19
        }
    }
    var v23 uint16 = 2
    v3 = append(v3, "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
