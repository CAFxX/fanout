package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint32) int8 {
    p0 = 1
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    return 100
}

func main() {
    v1 := []string{"", "\n\t", "foo bar"}
    v2 := make(map[int]bool)
    var v3 uint16 = 1
    var v4 bool = false
    var v5 uint = 0
    switch ("foo bar" + (("\n\t" + "hello") + (("a" + "\n\t") + "a"))) {
    case "":
        v4 = false
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "foo bar":
        for fi6 := 0; fi6 < 19; fi6++ {
            switch -687 {
            case 4294967295:
                v7 := GenS1{F0: 0}
                _ = v7
            case -4:
                v5 = (v5 ^ v5)
            default:
                v8 := GenS1{F0: 617}
                _ = v8
            }
            v9 := GenS1{F0: -32768}
            _ = fi6
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 15; fi10++ {
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 int32 = 79
                v13 := GenS2{F0: 149}
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    }
    v4 = (-1 >= -1)
    v3 = 1
    fmt.Println(v4, v3)
    switch "\n\t" {
    case "foo bar":
        v5 += (v5 << (((v5 | (v5 | (v5 >> ((42) & 63))))) & 63))
    case "a":
        for fi14 := 0; fi14 < 18; fi14++ {
            v2[(fi14 / (((fi14 ^ (fi14 + fi14))) | 1))] = (127 > 1)
            _ = fi14
        }
    }
    if ("\n\t" >= "\n\t") {
        var v15 int64 = 42
        _ = v15
    } else {
        v16 := GenS3{F0: 50, F1: -32768}
        _ = v16
    }
    if true {
        v4 = v4
    } else {
        v4 = false
        v5 = (^(v5))
        v4 = v4
    }
    var v17 float64 = 0.0
    var v18 uint8 = 255
    var v19 int8 = genFn0(2)
    v18 = (v18 ^ (v18 - (v18 + 103)))
    v2[int(v18)] = (0 >= 42)
    var v20 int8 = (v19 & int8(v3))
    v21 := GenS3{F0: v18, F1: 868}
    _ = len(v1)
    v4 = (v20 <= (v20 - int8(v5)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
