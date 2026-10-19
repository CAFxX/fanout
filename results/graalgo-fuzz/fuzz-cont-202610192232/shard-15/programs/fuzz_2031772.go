package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint16; F2 string }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int16; F1 int16; F2 uint32 }
func genFn0(p0 uint) int {
    v4 := GenS1{F0: "foo bar"}
    p0 = (^(p0))
    _ = p0
    _ = v4
    return 936
}

func main() {
    var v1 uint32 = 42
    v2 := []uint32{}
    v3 := make(map[int]int)
    v1 = (^(v1))
    var v5 int8 = 100
    v5 = (v5 << ((255) & 63))
    v6 := GenS1{F0: (("" + "a") + "")}
    for fi7 := 0; fi7 < 17; fi7++ {
        v1 = (v1 + v1)
        if len(v2) > 0 {
            v2[0] = v1
        }
        _ = fi7
    }
    v3[304] = -1
    v5 -= v5
    v5 -= v5
    v8 := GenS2{F0: (v1 + (v1 + (^(v1)))), F1: 0, F2: ("a" + "a")}
    v5 -= -1
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 int32 = -2147483648
        var v11 uint = 404
        switch "foo bar" {
        case "foo bar":
            v12 := GenS2{F0: v1, F1: 513, F2: ("hello" + "hello")}
            _ = v12
        case "hello":
            switch (("hello" + "a") + ("" + "\n\t")) {
            case "a":
                for fi13 := 0; fi13 < 4; fi13++ {
                    v14 := GenS3{F0: v1, F1: 65535}
                    v1 -= (v1 >> (((v11 - v11)) & 63))
                    var v15 int32 = int32(v5)
                    fmt.Println(fi13, v1, v11)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case "foo bar":
                if (0 > 1) {
                    v16 := GenS3{F0: (v1 - (^(v1))), F1: 0}
                    v1 -= v1
                    for fi17 := 0; fi17 < 13; fi17++ {
                        v18 := GenS4{F0: -197, F1: -1, F2: 2147483648}
                        _ = fi17
                        _ = v18
                    }
                    _ = v16
                }
            default:
                if len(v2) > 0 {
                    v2[0] = 295
                }
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v8
}
