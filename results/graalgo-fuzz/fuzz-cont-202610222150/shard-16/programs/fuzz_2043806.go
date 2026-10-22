package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]int)
    v3 := []string{""}
    var v4 int16 = 540
    for fi5 := 0; fi5 < 20; fi5++ {
        if v1 {
            fmt.Println(v1, v1)
            for fi6 := 0; fi6 < 19; fi6++ {
                fmt.Println(v4, fi5, fi5)
                v4 = v4
                _ = len(v3)
                _ = fi6
            }
        }
        for fi7 := 0; fi7 < 2; fi7++ {
            if len(v3) > 0 {
                v3[0] = ("a" + "a")
            }
            _ = fi7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: v4, F1: 183}
        var v10 bool = v1
        _ = fi8
        _ = v10
        _ = v9
    }
    if (1 == 32767) {
        v11 := GenS2{F0: 42}
        _ = v11
    }
    var v12 uint16 = 42
    var v13 uint8 = 1
    v14 := GenS1{F0: (v4 % ((-1) | 1)), F1: (v13 + 255)}
    var v15 string = ("\n\t" + ("a" + "foo bar"))
    v3 = append(v3, ((((v15 + (((v15 + (((((v15 + v15) + v15) + v15) + "") + "hello")) + "hello") + v15)) + v15) + v15) + (v15 + (((v15 + "") + (v15 + "hello")) + v15))))
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 string = v15
        v17 = v15
        _ = fi16
        _ = v17
    }
    switch v15 {
    case "":
        switch 42 {
        case 506:
            _ = len(v3)
        case -477:
            v2[2147483647] = -595
        case 2147483647:
            v13 = v13
        }
    case "\n\t":
        v15 = v15
    }
    v18 := GenS3{F0: 947}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
