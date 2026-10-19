package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float32 = 0.1
    v2 := []string{""}
    v3 := make(map[string]uint32)
    var v4 int64 = -481
    var v5 uint16 = 1
    v5 = (v5 % (((v5 | (v5 - (v5 / ((32767) | 1))))) | 1))
    var v6 uint64 = 1
    v5 += (v5 << ((1) & 63))
    v5 -= (^(v5))
    if true {
        fmt.Println(v5, v1, v6)
        var v7 int64 = (v4 + (v4 / (((v4 ^ (v4 * v4))) | 1)))
        _ = v7
    }
    var v8 uint8 = 0
    delete(v3, ("a" + ("\n\t" + "\n\t")))
    for fi9 := 0; fi9 < 13; fi9++ {
        v5 += 2
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(v4)
            v11 := GenS1{F0: -837}
            _ = fi10
            _ = v11
        }
        switch (fi9 & (fi9 % (((fi9 >> ((2) & 63))) | 1))) {
        case -612:
            var v12 int = 32767
            _ = v12
        case 0:
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := GenS1{F0: -631}
                switch fi9 {
                case 128:
                    v3["foo bar"] = 1
                case 982:
                    v1 += (0.1 / 0.1)
                    switch fi13 {
                    case 0:
                        v15 := GenS1{F0: -1}
                        _ = v15
                    }
                }
                _ = fi13
                _ = v14
            }
        case -1:
            if (("foo bar" + "foo bar") != "\n\t") {
                v16 := GenS1{F0: int16(v8)}
                var v17 uint8 = (v8 / ((v8) | 1))
                _ = v16
                _ = v17
            }
        default:
            var v18 int32 = -2147483648
            _ = v18
        }
        _ = fi9
    }
    fmt.Println(v5)
    v6 -= (v6 >> ((925) & 63))
    v6 -= v6
    fmt.Println(v1, v1, v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
