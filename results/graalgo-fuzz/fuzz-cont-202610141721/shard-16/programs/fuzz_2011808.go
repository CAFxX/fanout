package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int16 = -1
    var v2 bool = false
    v3 := []int{127, 441, 1}
    v4 := make(map[string]uint32)
    var v5 bool = (((v1 / ((32767) | 1)) == 32767) || v2)
    var v6 float32 = (-2.25 + 3.14159)
    var v7 uint16 = 2
    var v8 uint64 = 4294967295
    var v9 uint8 = 255
    if (0.1 > v6) {
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: -101}
            var v12 int16 = (v1 * -435)
            v4[(((("a" + "hello") + "hello") + "") + "foo bar")] = 119
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 uint32 = 1
    if len(v3) > 0 {
        v3[0] = 1
    }
    v14 := GenS2{F0: (v8 + v8)}
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            switch "a" {
            case "":
                var v17 float64 = float64(0.1)
                _ = v17
            }
            _ = fi16
        }
        v2 = ("hello" >= (((("" + ("hello" + "hello")) + "\n\t") + "hello") + "\n\t"))
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
