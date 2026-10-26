package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint16 }
type GenS2 struct { F0 int64; F1 bool }
func genFn0(p0 uint8) int16 {
    var v6 int16 = -84
    for fi7 := 0; fi7 < 20; fi7++ {
        switch "hello" {
        case "foo bar":
            var v8 int64 = 0
            _ = v8
        case "a":
            fmt.Println(v6, p0)
        case "hello":
            var v9 int8 = -1
            _ = v9
        default:
            var v10 int16 = v6
            _ = v10
        }
        _ = fi7
    }
    fmt.Println(p0)
    switch 128 {
    case 0:
        if true {
            v6 = 0
        }
    }
    _ = p0
    _ = v6
    return (v6 / ((123) | 1))
}

func main() {
    var v1 uint8 = 255
    v2 := []string{"a", ("\n\t" + "a")}
    v3 := make(map[uint32]int)
    v4 := []uint32{}
    v5 := make(map[string]uint32)
    delete(v5, "foo bar")
    v11 := GenS1{F0: int(v1), F1: -755, F2: 0}
    var v12 uint = 0
    v12 = v12
    v1 = v1
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            fmt.Println(v1)
            _ = fi14
        }
        _ = fi13
    }
    if len(v2) > 0 {
        v2[0] = ("hello" + "\n\t")
    }
    if false {
        if len(v2) > 0 {
            v2[0] = ("" + ("\n\t" + ""))
        }
    }
    switch 916 {
    case -2147483648:
        v12 -= (v12 / ((v12) | 1))
    case 525:
        v15 := GenS2{F0: 9223372036854775807, F1: (true && true)}
        _ = v15
    case 4294967295:
        v12 = (v12 ^ (v12 * (^(v12))))
    }
    if false {
        for fi16 := 0; fi16 < 17; fi16++ {
            for fi17 := 0; fi17 < 9; fi17++ {
                v12 = (v12 / ((v12) | 1))
                if (!false) {
                    delete(v5, "foo bar")
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v1 = v1
                }
                _ = fi17
            }
            fmt.Println(v1, fi16, v1)
            _ = fi16
        }
    } else {
        var v18 bool = (-786 == 2147483647)
        for fi19 := 0; fi19 < 4; fi19++ {
            v12 += 958
            _ = len(v5)
            _ = fi19
        }
        _ = v18
    }
    v20 := GenS2{F0: 0, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
