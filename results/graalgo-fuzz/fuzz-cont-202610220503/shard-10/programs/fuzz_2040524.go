package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 uint32 = 1
    var v4 uint32 = 2147483648
    if false {
        var v5 int8 = -1
        _ = v5
    } else {
        if (630 >= -1) {
            if true {
                v3 -= (v4 - 2147483647)
                fmt.Println(v4)
            }
        }
    }
    v3 = (v3 | v3)
    v6 := []int{}
    var v7 string = ""
    var v8 int = 4294967295
    v8 = v8
    v3 = v4
    fmt.Println(v8)
    for fi9 := 0; fi9 < 4; fi9++ {
        v3 -= v3
        var v10 float32 = (((((3.14159 + 1.5) * (1.5 * 0.1)) / (1e10 - (float32((3.14159 + (1e10 + 0.0))) * float32(1.5)))) * (0.0 + ((0.1 * (float32(float64(0.0)) * -2.25)) + (float32(-2.25) * float32((1e10 + (1e10 - 0.1))))))) / 1e10)
        _ = fi9
        _ = v10
    }
    v11 := v2[(("" + (v7 + (v7 + "foo bar"))) + ("foo bar" + (v7 + "")))]
    var v12 int64 = -935
    fmt.Println(v8, v3)
    if (((v4 << ((11) & 63)) > v3) || (1000 >= 886)) {
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := GenS1{F0: 127, F1: -128, F2: v12}
            if true {
                v4 -= (v4 * (v3 ^ 4294967295))
                for fi15 := 0; fi15 < 19; fi15++ {
                    if (!false) {
                        switch (v7 + "hello") {
                        case "foo bar":
                            if true {
                                var v16 string = (v7 + v7)
                                _ = v16
                            } else {
                                v17 := GenS1{F0: uint8(v4), F1: -1, F2: v12}
                                fmt.Println(v3, v11)
                                v18 := v2[""]
                                _ = v17
                                _ = v18
                            }
                        }
                    } else {
                        _ = len(v2)
                        v4 = 65535
                    }
                    v19 := GenS1{F0: 127, F1: -1, F2: (v12 - 42)}
                    fmt.Println(fi13, fi13, v12)
                    _ = fi15
                    _ = v19
                }
            } else {
                _ = len(v6)
                switch ("\n\t" + "a") {
                case "hello":
                    fmt.Println(v4)
                case "":
                    v20 := GenS1{F0: 1, F1: 0, F2: v12}
                    _ = v20
                case "\n\t":
                    v21 := v1[(v4 << ((4294967295) & 63))]
                    _ = v21
                default:
                    v22 := GenS1{F0: 2, F1: -88, F2: (v12 | (v12 - int64(v3)))}
                    _ = v22
                }
            }
            _ = fi13
            _ = v14
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
