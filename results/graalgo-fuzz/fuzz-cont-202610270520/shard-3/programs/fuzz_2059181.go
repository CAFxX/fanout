package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint64 }
func genFn0(p0 uint32) bool {
    var v6 float32 = ((3.14159 - 0.1) * 3.14159)
    p0 = 2147483647
    _ = p0
    _ = v6
    return true
}

func main() {
    v1 := []string{"foo bar"}
    var v2 int64 = -328
    var v3 uint64 = 2
    var v4 uint = 65535
    v5 := make(map[string]bool)
    v4 = (v4 << ((v4) & 63))
    v2 = (-(v2))
    if false {
        var v7 uint = 366
        if true {
            v4 += v7
            v4 -= (v4 | 968)
        } else {
            if true {
                v8 := GenS1{F0: 846, F1: 42, F2: v3}
                for fi9 := 0; fi9 < 1; fi9++ {
                    if true {
                        v10 := GenS1{F0: 668, F1: 63, F2: uint64(v4)}
                        v7 -= v7
                        v3 = v3
                        _ = v10
                    } else {
                        v2 = 612
                        v1 = append(v1, "foo bar")
                        switch "foo bar" {
                        case "foo bar":
                            var v11 int64 = v2
                            _ = v11
                        default:
                            var v12 bool = (109 >= 42)
                            _ = v12
                        }
                    }
                    fmt.Println(v7, v2, v7)
                    _ = fi9
                }
                fmt.Println(v3, v2)
                _ = v8
            } else {
                v13 := GenS1{F0: 801, F1: 255, F2: (v3 & 42)}
                v2 = v2
                _ = v13
            }
            fmt.Println(v4)
        }
        _ = v7
    }
    v14 := GenS1{F0: 1, F1: 127, F2: 54}
    var v15 bool = (255 > uint16(v4))
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := v5["foo bar"]
        v18 := v5[(((("\n\t" + "a") + "foo bar") + "foo bar") + (("a" + (("hello" + "\n\t") + "\n\t")) + "a"))]
        v2 -= (v2 / (((v2 & v2)) | 1))
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 float64 = 1.5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v20 int16 = -26
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
