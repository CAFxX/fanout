package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{536, 13, 0, 4294967295}
    var v3 int = -9223372036854775808
    v4 := []string{("\n\t" + ("\n\t" + ("foo bar" + "a")))}
    var v5 int32 = 42
    var v6 uint8 = 2
    var v7 float64 = 1.5
    var v8 float64 = float64(1.5)
    v5 = v5
    delete(v1, 65535)
    var v9 string = ("\n\t" + "a")
    v3 -= (v3 + (v3 | 2147483647))
    if len(v2) > 0 {
        v2[0] = 1
    }
    if true {
        if (true || true) {
            if false {
                for fi10 := 0; fi10 < 2; fi10++ {
                    v11 := GenS1{F0: 0}
                    if false {
                        _ = len(v4)
                    } else {
                        v12 := GenS1{F0: -32768}
                        fmt.Println(v6, v8)
                        _ = v12
                    }
                    _ = fi10
                    _ = v11
                }
                fmt.Println(v6, v6, v7)
            }
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 uint = 0
            _ = fi13
            _ = v14
        }
    } else {
        v15 := GenS2{F0: 805, F1: v6}
        var v16 float64 = (0.1 + (v8 - v8))
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 240}
    switch (v3 & v3) {
    case 4294967295:
        var v18 float32 = 0.1
        _ = v18
    }
    _ = len(v4)
    v8 += float64((1e10 / 1.5))
    v7 = (v8 * float64(3.14159))
    var v19 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
