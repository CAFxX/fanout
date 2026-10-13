package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 1
    v3 := []string{"a", "hello"}
    v4 := make(map[uint32]uint32)
    v5 := []int{378}
    v6 := GenS1{F0: 326}
    fmt.Println(v2)
    var v7 int8 = (v2 * 0)
    var v8 int8 = v2
    v9 := GenS1{F0: 0}
    delete(v4, 833)
    var v10 int64 = 3
    var v11 int16 = -269
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    switch -691 {
    case 1:
        v12 := GenS1{F0: 1}
        _ = v12
    case 127:
        var v13 bool = (float64((0.1 + 0.0)) != 3.14159)
        _ = v13
    case 0:
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS1{F0: -1}
            var v16 string = ("\n\t" + "\n\t")
            switch fi14 {
            case 0:
                v17 := GenS1{F0: 0}
                _ = v17
            case 1:
                v18 := GenS1{F0: 1}
                _ = v18
            default:
                v2 += (v8 % ((v7) | 1))
            }
            _ = fi14
            _ = v15
            _ = v16
        }
    default:
        var v19 float32 = (0.1 * float32(0.1))
        _ = v19
    }
    if (!true) {
        v7 = int8(v11)
        v2 = -1
        var v20 bool = false
        _ = v20
    } else {
        if true {
            fmt.Println(v10)
            switch ("a" + "a") {
            case "foo bar":
                var v21 int8 = v8
                _ = v21
            default:
                switch "\n\t" {
                case "\n\t":
                    switch "foo bar" {
                    case "\n\t":
                        var v22 float32 = 1e10
                        _ = v22
                    case "":
                        fmt.Println(v7, v7, v11)
                    default:
                        if len(v5) > 0 {
                            v5[0] = 4294967295
                        }
                    }
                }
            }
            var v23 uint32 = 42
            _ = v23
        }
        v10 = v10
        v8 -= v7
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
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
    _ = v9
}
