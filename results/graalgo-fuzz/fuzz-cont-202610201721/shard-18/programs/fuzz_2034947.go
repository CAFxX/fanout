package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int8; F2 uint8 }

func main() {
    var v1 uint16 = 32767
    var v2 int64 = 9223372036854775807
    var v3 bool = (!(67 <= 1))
    v4 := make(map[uint32]int)
    var v5 float64 = 1.5
    fmt.Println(v5)
    var v6 float32 = 0.1
    for fi7 := 0; fi7 < 6; fi7++ {
        if v3 {
            v1 = v1
        } else {
            var v8 uint = 483
            var v9 float64 = (float64(0.0) + 1e10)
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS1{F0: (fi10 * (-(fi10)))}
        v1 = v1
        _ = fi10
        _ = v11
    }
    fmt.Println(v2, v5, v2)
    var v12 int64 = (v2 - v2)
    v4[2] = 632
    switch "hello" {
    case "foo bar":
        v6 += (v6 + ((v6 * float32(v5)) - (v6 + 1e10)))
    default:
        v13 := GenS2{F0: 2147483647, F1: -1, F2: 128}
        _ = v13
    }
    v14 := []int{2147483647}
    var v15 uint64 = 19
    var v16 int = -17
    switch "hello" {
    case "\n\t":
        switch (((("" + ((("" + "") + ("a" + "\n\t")) + "a")) + (("\n\t" + "hello") + "")) + "hello") + ("hello" + "foo bar")) {
        case "":
            for fi17 := 0; fi17 < 7; fi17++ {
                fmt.Println(v5, v2)
                v4[uint32(v1)] = v16
                switch "hello" {
                case "a":
                    var v18 uint8 = 2
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 19; fi19++ {
                        fmt.Println(v3, v3, v6)
                        switch 9223372036854775807 {
                        case 9223372036854775807:
                            v20 := GenS1{F0: 128}
                            _ = v20
                        case 128:
                            switch 32767 {
                            case -1:
                                if false {
                                    var v21 int = fi17
                                    v5 += 1e10
                                    _ = v21
                                }
                            case -975:
                                v22 := v4[4294967295]
                                _ = v22
                            case -282:
                                v2 += (v2 & 0)
                            default:
                                var v23 bool = v3
                                _ = v23
                            }
                        }
                        _ = fi19
                    }
                }
                _ = fi17
            }
        }
    case "foo bar":
        v24 := GenS2{F0: (v16 >> ((320) & 63)), F1: 42, F2: 0}
        _ = v24
    case "hello":
        v16 = v16
    default:
        var v25 uint16 = 925
        _ = v25
    }
    fmt.Println(v2)
    _ = len(v14)
    var v26 int16 = -884
    var v27 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
