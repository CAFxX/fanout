package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int8 = -128
    var v2 uint64 = 4294967295
    v3 := []uint32{42, 2147483647, 2, 1}
    var v4 uint64 = 18446744073709551615
    var v5 int = 296
    var v6 string = "foo bar"
    v7 := GenS1{F0: 0, F1: (v1 + v1)}
    if true {
        fmt.Println(v2, v4)
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint16 = 255
            _ = fi8
            _ = v9
        }
        var v10 int64 = 170
        _ = v10
    } else {
        for fi11 := 0; fi11 < 20; fi11++ {
            for fi12 := 0; fi12 < 3; fi12++ {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                if len(v3) > 0 {
                    v3[0] = 2
                }
                _ = fi12
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                v4 -= 0
                _ = fi13
            }
            switch (v5 / ((-44) | 1)) {
            case 913:
                v2 = 18446744073709551615
            case 128:
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
            }
            _ = fi11
        }
        switch v5 {
        case 32768:
            for fi14 := 0; fi14 < 1; fi14++ {
                switch v6 {
                case "foo bar":
                    for fi15 := 0; fi15 < 9; fi15++ {
                        var v16 uint32 = 2147483648
                        v1 = (v1 >> ((255) & 63))
                        fmt.Println(v2, fi15)
                        if len(v3) > 0 {
                            v3[0] = v16
                        }
                        v17 := GenS2{F0: (v2 ^ (^(v2))), F1: -1000}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case "a":
                    var v18 uint8 = 42
                    _ = v18
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v2 += 0
                        _ = fi19
                    }
                }
                _ = fi14
            }
        case 128:
            if true {
                v4 = v4
                var v20 int8 = v1
                _ = v20
            }
        case -1:
            v2 = v4
        }
    }
    v4 -= 4294967295
    var v21 uint32 = 2147483648
    var v22 int64 = 42
    if (!((v4 ^ 4294967295) > (v2 / ((v2) | 1)))) {
        var v23 float32 = 1.5
        switch "hello" {
        case "foo bar":
            v4 = v2
        default:
            v2 += 1
        }
        fmt.Println(v1, v1, v2)
        _ = v23
    }
    fmt.Println(v21)
    for fi24 := 0; fi24 < 19; fi24++ {
        if true {
            v25 := GenS2{F0: 243, F1: 2147483647}
            for fi26 := 0; fi26 < 7; fi26++ {
                if len(v3) > 0 {
                    v3[0] = (v21 % ((2147483647) | 1))
                }
                v27 := GenS1{F0: (^(v5)), F1: (v1 >> ((18446744073709551615) & 63))}
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 2; fi28++ {
                switch "hello" {
                case "\n\t":
                    v29 := GenS1{F0: fi24, F1: (v1 & (v1 ^ (v1 & (v1 >> ((0) & 63)))))}
                    _ = v29
                case "a":
                    v30 := GenS1{F0: (-(v5)), F1: (v1 % (((v1 * (v1 | (v1 - 0)))) | 1))}
                    _ = v30
                case "hello":
                    var v31 uint8 = 127
                    _ = v31
                default:
                    var v32 uint8 = 128
                    _ = v32
                }
                _ = fi28
            }
            _ = v25
        }
        if len(v3) > 0 {
            v3[0] = v21
        }
        if ((fi24 | (fi24 << ((255) & 63))) != 2147483647) {
            var v33 string = v6
            v5 += 544
            _ = v33
        }
        _ = fi24
    }
    var v34 uint16 = 255
    v1 = (v1 >> ((823) & 63))
    var v35 uint = 65535
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v5, v4)
    v1 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
