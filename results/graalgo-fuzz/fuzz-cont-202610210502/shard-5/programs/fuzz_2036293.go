package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int; F1 uint }

func main() {
    v1 := []uint32{850, 916, 65535}
    var v2 int16 = 42
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: 18446744073709551615, F1: 0}
    fmt.Println(v2)
    fmt.Println(v2)
    if ((((((!true) && (0 != 2)) || (1 != 1)) && (-9223372036854775808 != 32767)) && (!false)) || false) {
        if false {
            v2 = v2
            if ("foo bar" <= "\n\t") {
                v6 := GenS1{F0: 42, F1: 0}
                v2 = (v2 - v2)
                _ = v6
            }
            var v7 int8 = -1
            _ = v7
        } else {
            v1 = append(v1, 2)
            v2 += (v2 << ((44) & 63))
        }
    } else {
        var v8 float32 = ((1e10 - (1e10 - -2.25)) - float32((3.14159 + (1.5 - 1.5))))
        v8 -= float32(1e10)
        _ = v8
    }
    v2 = 1
    v9 := GenS1{F0: 0, F1: 2}
    switch 42 {
    case 2147483647:
        for fi10 := 0; fi10 < 7; fi10++ {
            _ = len(v4)
            v11 := GenS1{F0: 554, F1: 101}
            v12 := GenS1{F0: 436, F1: 255}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v1 = append(v1, 1)
        _ = fi13
    }
    v2 = 32767
    var v14 float64 = 1.5
    var v15 int8 = 46
    switch 147 {
    case 32767:
        if true {
            switch -1 {
            case 127:
                if false {
                    switch (("a" + "\n\t") + "") {
                    case "":
                        v16 := GenS1{F0: 255, F1: 42}
                        _ = v16
                    case "hello":
                        v15 += -90
                    case "a":
                        if true {
                            v17 := GenS1{F0: 65535, F1: 35}
                            _ = v17
                        }
                    default:
                        v18 := GenS2{F0: 42, F1: 848}
                        _ = v18
                    }
                    v19 := GenS1{F0: 892, F1: 42}
                    var v20 uint8 = 227
                    _ = v19
                    _ = v20
                }
            default:
                if true {
                    v15 -= (-(v15))
                    var v21 uint = 242
                    for fi22 := 0; fi22 < 14; fi22++ {
                        if len(v1) > 0 {
                            v1[0] = 328
                        }
                        _ = fi22
                    }
                    _ = v21
                }
            }
        }
    case 2147483647:
        var v23 float64 = (v14 / (v14 - -2.25))
        _ = v23
    case 128:
        var v24 int16 = v2
        _ = v24
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        v1 = append(v1, 664)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
