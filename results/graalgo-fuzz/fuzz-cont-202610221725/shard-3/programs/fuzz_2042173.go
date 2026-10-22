package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint64 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    v3 := []uint32{0, 0, 288}
    v4 := make(map[string]string)
    var v5 uint = 2
    var v6 uint16 = 42
    v2[-106] = -983
    var v7 uint16 = v6
    var v8 int32 = -601
    var v9 uint16 = (v6 | 813)
    for fi10 := 0; fi10 < 19; fi10++ {
        if (false && true) {
            v6 -= (v9 + (v9 + (v7 >> (((v5 + (v5 ^ (v5 + v5)))) & 63))))
        } else {
            v8 = -141
            v11 := v1[32767]
            _ = v11
        }
        if false {
            for fi12 := 0; fi12 < 3; fi12++ {
                v9 = (v6 >> ((v5) & 63))
                switch "foo bar" {
                case "":
                    var v13 uint64 = uint64(v5)
                    _ = v13
                }
                _ = fi12
            }
            v9 = (v9 << (((v5 << (((v5 * v5)) & 63))) & 63))
        }
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 float32 = 1.5
            _ = fi14
            _ = v15
        }
        _ = fi10
    }
    v9 -= 0
    if false {
        for fi16 := 0; fi16 < 1; fi16++ {
            v1[724] = true
            if (42 == 1) {
                for fi17 := 0; fi17 < 13; fi17++ {
                    v18 := GenS1{F0: v5, F1: 1, F2: 553}
                    _ = fi17
                    _ = v18
                }
            }
            fmt.Println(v8, v5)
            _ = fi16
        }
        v9 = (v7 >> (((v5 * (v5 << (((v5 * (v5 * v5))) & 63)))) & 63))
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v9 = (^(v6))
        var v20 int16 = 0
        _ = fi19
        _ = v20
    }
    var v21 uint8 = 116
    var v22 int8 = 127
    switch "hello" {
    case "\n\t":
        var v23 int16 = 32767
        _ = v23
    }
    v6 += uint16(v5)
    var v24 uint = (v5 | (v5 ^ 818))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
