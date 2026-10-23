package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int16 }
type GenS2 struct { F0 uint; F1 bool }

func main() {
    var v1 int16 = -647
    v2 := []string{}
    v3 := make(map[int]uint32)
    v4 := GenS1{F0: 4294967295, F1: -62, F2: (v1 >> ((280) & 63))}
    for fi5 := 0; fi5 < 10; fi5++ {
        if true {
            v6 := GenS1{F0: (fi5 % ((fi5) | 1)), F1: 0, F2: v1}
            fmt.Println(fi5, v1)
            if (!(274 >= 65535)) {
                var v7 int16 = (v1 ^ v1)
                _ = v7
            } else {
                fmt.Println(fi5, fi5)
                for fi8 := 0; fi8 < 14; fi8++ {
                    _ = len(v3)
                    _ = fi8
                }
            }
            _ = v6
        }
        _ = fi5
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    var v9 float32 = 0.0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3[-1] = 65535
    switch -2147483648 {
    case -175:
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS2{F0: 100, F1: ("" != "foo bar")}
            _ = fi10
            _ = v11
        }
    case 32767:
        v9 -= v9
    case 9223372036854775807:
        var v12 int8 = -1
        _ = v12
    }
    v13 := GenS2{F0: 2, F1: false}
    var v14 int = 894
    v14 = (v14 >> ((65535) & 63))
    v14 = 128
    var v15 uint = 255
    _ = len(v3)
    fmt.Println(v14)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v16 uint8 = 185
    var v17 int8 = 1
    v16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
