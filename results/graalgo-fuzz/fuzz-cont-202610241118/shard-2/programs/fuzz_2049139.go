package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint8 }

func main() {
    var v1 uint = 759
    var v2 uint8 = 187
    var v3 uint = (v1 | (v1 & 65535))
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    var v6 bool = true
    for fi7 := 0; fi7 < 10; fi7++ {
        v6 = ((192 == 127) && (!true))
        _ = fi7
    }
    var v8 int16 = 0
    var v9 uint8 = 128
    v5[0] = "foo bar"
    if v6 {
        var v10 float32 = 0.1
        v11 := GenS1{F0: 32767, F1: 65535, F2: 692}
        var v12 bool = true
        _ = v10
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 2; fi13++ {
            switch ("\n\t" + ("\n\t" + ((((("foo bar" + "foo bar") + "foo bar") + "a") + "a") + (("" + "foo bar") + "a")))) {
            case "\n\t":
                for fi14 := 0; fi14 < 3; fi14++ {
                    v8 += 951
                    if (v1 != (v3 / (((v1 + (v1 - 775))) | 1))) {
                        var v15 uint = 4294967295
                        _ = v15
                    } else {
                        if false {
                            var v16 int16 = v8
                            v16 += v16
                            var v17 int8 = -1
                            _ = v16
                            _ = v17
                        } else {
                            var v18 uint64 = 78
                            var v19 float64 = -2.25
                            _ = v18
                            _ = v19
                        }
                    }
                    _ = fi14
                }
            case "foo bar":
                v20 := GenS2{F0: 65535, F1: 171, F2: 88}
                _ = v20
            case "hello":
                v8 -= (v8 + (v8 | v8))
            }
            v8 = v8
            var v21 int8 = 42
            _ = fi13
            _ = v21
        }
        v6 = (42 < -64)
    }
    fmt.Println(v9, v2, v6)
    switch ("" + ("\n\t" + "hello")) {
    case "a":
        v22 := v5[4294967295]
        _ = v22
    case "\n\t":
        v9 = (v9 + v2)
    case "":
        fmt.Println(v3)
    default:
        v4[""] = 0
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        fmt.Println(v3, v6)
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
