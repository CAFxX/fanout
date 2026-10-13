package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[int]uint32)
    v3 := []int{32768}
    v4 := []uint32{634}
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := GenS1{F0: 1, F1: 289, F2: 2}
            _ = fi6
            _ = v7
        }
        switch fi5 {
        case 2147483647:
            var v8 uint32 = 1
            _ = v8
        case 431:
            var v9 string = "a"
            _ = v9
        case -20:
            switch ((("\n\t" + "\n\t") + "") + "foo bar") {
            case "":
                v2[1] = 473
            default:
                fmt.Println(fi5, v1)
            }
        }
        _ = fi5
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 int = fi10
        _ = fi10
        _ = v11
    }
    switch 32768 {
    case 32768:
        if true {
            v1 = (v1 / v1)
            v2[32768] = 548
            v1 = v1
        }
    case 127:
        v12 := GenS1{F0: -1, F1: 65535, F2: 390}
        _ = v12
    default:
        v1 = (0.0 * v1)
    }
    var v13 uint64 = 2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v14 := GenS1{F0: 42, F1: 2, F2: 7}
    v15 := GenS2{F0: (2147483648 == 4294967295), F1: 255, F2: 65535}
    v13 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
