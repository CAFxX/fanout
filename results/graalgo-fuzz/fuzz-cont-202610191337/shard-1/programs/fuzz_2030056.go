package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int32 }

func main() {
    var v1 uint64 = 862
    var v2 uint = 743
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    v1 += v1
    v5 := GenS1{F0: ("a" + "foo bar"), F1: true, F2: 848}
    v1 = uint64(v2)
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 float64 = 0.1
        fmt.Println(v7, fi6)
        v1 = v1
        _ = fi6
        _ = v7
    }
    fmt.Println(v1, v2)
    fmt.Println(v1, v2)
    v2 = (v2 >> ((v2) & 63))
    fmt.Println(v2, v1)
    v2 = 972
    var v8 uint16 = 42
    v9 := GenS1{F0: "foo bar", F1: false, F2: v8}
    for fi10 := 0; fi10 < 15; fi10++ {
        if false {
            var v11 uint = v2
            _ = v11
        } else {
            var v12 uint16 = v8
            fmt.Println(v1)
            _ = v12
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            if true {
                switch (fi14 + fi14) {
                case 9223372036854775807:
                    v15 := GenS2{F0: 2, F1: 127, F2: -405}
                    _ = v15
                case -1:
                    v16 := v4[uint32(v1)]
                    _ = v16
                }
                v2 += uint(v1)
            }
            _ = fi14
        }
        _ = fi13
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
