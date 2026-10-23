package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int16; F1 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 1
    var v3 int16 = 1
    if true {
        var v4 int32 = -734
        if true {
            var v5 int32 = v4
            v2 += (v2 + (v5 % (((v4 % (((v5 % ((v4) | 1))) | 1))) | 1)))
            v4 = (v4 * -110)
            _ = v5
        } else {
            var v6 float64 = 0.1
            _ = v6
        }
        v7 := v1[""]
        _ = v4
        _ = v7
    }
    var v8 float32 = -2.25
    v1[""] = 1
    v8 = (float32(((float64(float32(((float64((3.14159 * float32(1.5))) + 0.0) - 0.1))) + ((float64(v8) + 1.5) - 1e10)) * ((1.5 + ((((float64(0.0) + ((0.1 * (-2.25 - ((3.14159 - (-2.25 / (3.14159 / -2.25))) / 0.1))) - float64(v8))) / 1e10) * (0.1 * 1.5)) + 0.1)) * 1e10))) + 3.14159)
    v9 := []int{-372, 9223372036854775807, 32767}
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 uint64 = 42
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        v8 += -2.25
        for fi13 := 0; fi13 < 6; fi13++ {
            v8 = v8
            _ = fi13
        }
        _ = fi12
    }
    v14 := GenS1{F0: ((("a" + "\n\t") + ("foo bar" + "")) + ""), F1: 65535}
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS1{F0: "a", F1: 164}
            v3 = v3
            switch fi16 {
            case 128:
                if true {
                    v9 = append(v9, int(v3))
                } else {
                    var v18 int64 = -1
                    v19 := GenS2{F0: (v3 << ((2) & 63)), F1: (v3 - v3)}
                    v20 := v1[("hello" + ("a" + "hello"))]
                    if len(v9) > 0 {
                        v9[0] = (^(fi15))
                    }
                    var v21 int = v20
                    _ = v18
                    _ = v19
                    _ = v20
                    _ = v21
                }
            case 127:
                v8 -= v8
            case 0:
                v22 := GenS2{F0: v3, F1: v3}
                _ = v22
            }
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v2 = v2
    v1["hello"] = -984
    fmt.Println(v3)
    if (4294967295 < 148) {
        _ = len(v1)
        switch "foo bar" {
        case "hello":
            v3 = 1
        case "foo bar":
            v8 = ((3.14159 / (0.0 - v8)) * v8)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
