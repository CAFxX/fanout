package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 127
    v3 := make(map[string]int)
    var v4 uint32 = 1
    var v5 int = 128
    var v6 uint32 = v4
    if true {
        if (((-(v5)) > v5) && (float64(((((-2.25 / float32(3.14159)) / (0.0 - (3.14159 * 1e10))) - 1.5) * 3.14159)) > 3.14159)) {
            v7 := GenS1{F0: 101}
            _ = v7
        }
        if (false || true) {
            var v8 uint8 = (v2 ^ (v2 & 42))
            for fi9 := 0; fi9 < 6; fi9++ {
                v8 = (v2 + (v2 & (v2 << ((255) & 63))))
                _ = fi9
            }
            _ = v8
        } else {
            v10 := []int{-9}
            v4 = (v6 * v4)
            v11 := GenS1{F0: 0}
            _ = v10
            _ = v11
        }
        fmt.Println(v5, v2)
    }
    var v12 int64 = int64(v6)
    v13 := []int{-233}
    v14 := v1[v5]
    var v15 uint16 = 1
    fmt.Println(v14, v5)
    delete(v3, "foo bar")
    v3[(("hello" + "\n\t") + "\n\t")] = (v14 | v14)
    for fi16 := 0; fi16 < 16; fi16++ {
        for fi17 := 0; fi17 < 11; fi17++ {
            delete(v3, ((("" + "a") + "a") + "foo bar"))
            var v18 int32 = 101
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    switch v5 {
    case 4294967295:
        v2 = (v2 / (((v2 * (v2 & v2))) | 1))
    case 0:
        var v19 int64 = (v12 | (v12 % (((v12 + -9223372036854775808)) | 1)))
        _ = v19
    case -561:
        if (!((((float64(float32(((((float64(1e10) + 1.5) + 1.5) - 0.1) * 3.14159))) / float64(-2.25)) * 1e10) * (((3.14159 / (((0.1 - ((0.1 - float64((3.14159 - 0.0))) + 0.0)) * 1e10) - 1.5)) * 1e10) - 1.5)) <= float64((1e10 * 1.5)))) {
            var v20 uint = 584
            v15 -= (v15 * (v15 & 0))
            v21 := v1[(v14 % (((v14 & v5)) | 1))]
            _ = v20
            _ = v21
        } else {
            var v22 int = -337
            v23 := GenS2{F0: v15}
            _ = v22
            _ = v23
        }
    }
    var v24 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
