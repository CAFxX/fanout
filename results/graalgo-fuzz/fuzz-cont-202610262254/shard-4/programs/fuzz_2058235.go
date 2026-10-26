package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 uint; F2 uint32 }

func main() {
    v1 := []uint32{65535, 12, 727, 336}
    var v2 string = "foo bar"
    v3 := []uint32{0}
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: 32768, F1: -823}
    var v6 string = "\n\t"
    var v7 int32 = 562
    v8 := GenS2{F0: 32767, F1: 42, F2: (v2 + v6)}
    fmt.Println(v7, v6)
    var v9 bool = true
    switch 1 {
    case -787:
        v6 = (v2 + "")
    }
    v6 = v2
    for fi10 := 0; fi10 < 15; fi10++ {
        if ((0.0 * -2.25) < -2.25) {
            for fi11 := 0; fi11 < 15; fi11++ {
                v6 = v2
                _ = fi11
            }
            v12 := GenS3{F0: 599}
            _ = v12
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 uint16 = 2
        _ = fi13
        _ = v14
    }
    v15 := GenS4{F0: 9223372036854775807}
    switch (v6 + v6) {
    case "hello":
        if v9 {
            var v16 string = (v2 + "hello")
            switch -737 {
            case 127:
                for fi17 := 0; fi17 < 8; fi17++ {
                    v18 := GenS1{F0: (fi17 | fi17), F1: -204}
                    _ = fi17
                    _ = v18
                }
            case 743:
                v16 = v6
            case 9223372036854775807:
                switch v6 {
                case "a":
                    fmt.Println(v16, v2)
                default:
                    fmt.Println(v16)
                }
            }
            _ = v16
        } else {
            v6 = "a"
            var v19 uint16 = 2
            for fi20 := 0; fi20 < 12; fi20++ {
                var v21 uint8 = 0
                for fi22 := 0; fi22 < 17; fi22++ {
                    switch v2 {
                    case "":
                        v23 := GenS5{F0: v7, F1: 560, F2: 2}
                        _ = v23
                    case "hello":
                        if true {
                            var v24 uint8 = (v21 ^ v21)
                            _ = v24
                        }
                    case "a":
                        fmt.Println(v9, v6, v7)
                    default:
                        if ((2 == 0) && ((v19 ^ (v19 ^ (v19 | 177))) == v19)) {
                            v7 = (v7 >> ((2) & 63))
                            v25 := GenS1{F0: (fi20 | 916), F1: 32767}
                            _ = v25
                        }
                    }
                    v7 -= v7
                    _ = fi22
                }
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    case "\n\t":
        fmt.Println(v7, v9)
    default:
        fmt.Println(v6)
    }
    var v26 int8 = -128
    for fi27 := 0; fi27 < 1; fi27++ {
        for fi28 := 0; fi28 < 18; fi28++ {
            if v9 {
                if ((-1 != 0) && v9) {
                    var v29 uint16 = 42
                    var v30 int = fi28
                    fmt.Println(v30)
                    _ = v29
                    _ = v30
                }
                v31 := GenS5{F0: v7, F1: 18446744073709551615, F2: 2}
                _ = v31
            } else {
                v32 := v4[42]
                _ = v32
            }
            _ = fi28
        }
        var v33 int = (^(fi27))
        _ = fi27
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
