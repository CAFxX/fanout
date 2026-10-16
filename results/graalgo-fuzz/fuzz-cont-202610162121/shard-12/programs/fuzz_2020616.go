package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int32 }
type GenS2 struct { F0 bool }

func main() {
    var v1 float64 = 3.14159
    var v2 string = ""
    var v3 int = 631
    var v4 int16 = -32768
    v5 := []uint32{}
    var v6 string = (v2 + ((v2 + v2) + ((v2 + "\n\t") + "hello")))
    v4 += v4
    var v7 int32 = 1
    fmt.Println(v4)
    v8 := GenS1{F0: (v4 % (((v4 & 1)) | 1)), F1: v2, F2: (v7 ^ (v7 << ((4294967295) & 63)))}
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            v3 -= fi10
            var v11 uint64 = 654
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v12 := GenS2{F0: (v7 < (v7 << ((65535) & 63)))}
    v5 = append(v5, 855)
    switch (v2 + ("hello" + "\n\t")) {
    case "a":
        v13 := GenS2{F0: true}
        _ = v13
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v14 float64 = 0.1
    v3 -= int(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
