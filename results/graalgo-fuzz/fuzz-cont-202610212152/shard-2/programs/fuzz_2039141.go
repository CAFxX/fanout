package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }
type GenS3 struct { F0 int8; F1 uint16; F2 string }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint16 = 42
    var v2 float32 = 1.5
    v3 := make(map[string]int)
    var v4 float32 = (float32(0.1) - (v2 * v2))
    var v5 int16 = -1
    v6 := []uint32{157}
    if true {
        v6 = append(v6, 4294967295)
    }
    v5 = v5
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v7 := GenS1{F0: v5, F1: 694}
    v5 += -32768
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    delete(v3, "")
    v2 += float32(0.0)
    fmt.Println(v1, v4)
    v6 = append(v6, 42)
    var v8 int = 32768
    for fi9 := 0; fi9 < 15; fi9++ {
        v2 -= v2
        _ = fi9
    }
    var v10 int = (v8 - v8)
    v2 -= v2
    v11 := GenS2{F0: 0, F1: v1, F2: ("\n\t" + ("\n\t" + ("foo bar" + "foo bar")))}
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS3{F0: int8(v1), F1: (^(v1)), F2: ""}
        v14 := GenS4{F0: 100}
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
