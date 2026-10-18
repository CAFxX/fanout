package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int8; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int }

func main() {
    var v1 uint = 0
    v2 := make(map[string]int)
    var v3 uint16 = 65535
    var v4 uint32 = uint32(v3)
    switch ((((((("a" + "\n\t") + "a") + "a") + "\n\t") + (("" + (("a" + (("a" + "hello") + "")) + (("a" + ("\n\t" + ("a" + "a"))) + "a"))) + (("a" + "") + "hello"))) + "a") + "a") {
    case "hello":
        for fi5 := 0; fi5 < 17; fi5++ {
            for fi6 := 0; fi6 < 17; fi6++ {
                if true {
                    v1 -= uint(v3)
                    var v7 uint = (v1 >> (((^(v1))) & 63))
                    _ = v7
                }
                _ = fi6
            }
            v8 := []int{fi5, fi5}
            v9 := GenS1{F0: 1, F1: v4}
            _ = fi5
            _ = v8
            _ = v9
        }
    case "\n\t":
        v4 -= (v4 ^ v4)
    default:
        var v10 uint64 = 0
        _ = v10
    }
    v11 := GenS2{F0: 342, F1: 2}
    v12 := v2["a"]
    for fi13 := 0; fi13 < 1; fi13++ {
        if (true || (!false)) {
            if false {
                v12 = fi13
                var v14 uint32 = (v4 % (((v4 | v4)) | 1))
                var v15 float32 = 0.0
                _ = v14
                _ = v15
            } else {
                v16 := GenS3{F0: 128, F1: -122, F2: v12}
                for fi17 := 0; fi17 < 18; fi17++ {
                    v4 = 1
                    var v18 int8 = 42
                    _ = fi17
                    _ = v18
                }
                var v19 uint32 = (v4 << ((1) & 63))
                _ = v16
                _ = v19
            }
            v20 := []int{}
            _ = v20
        }
        _ = fi13
    }
    var v21 uint64 = 4294967295
    v22 := GenS4{F0: (v4 / ((v4) | 1))}
    v23 := GenS5{F0: 32767, F1: v1}
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 int8 = 42
        v12 = (fi24 + fi24)
        v26 := GenS6{F0: (fi24 & (-(fi24)))}
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
