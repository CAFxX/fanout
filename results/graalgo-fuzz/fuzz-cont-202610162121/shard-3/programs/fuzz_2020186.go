package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 string = (("" + ("foo bar" + "\n\t")) + (("hello" + "foo bar") + (((("" + "\n\t") + "foo bar") + ("hello" + "foo bar")) + ("foo bar" + (("a" + "\n\t") + "foo bar")))))
    var v2 float32 = 0.1
    v3 := []uint32{2147483648}
    var v4 uint16 = 0
    var v5 uint64 = 1
    v6 := make(map[uint32]bool)
    v7 := GenS1{F0: 2147483648}
    var v8 int32 = 42
    switch "hello" {
    case "hello":
        if (-153 != -618) {
            fmt.Println(v4, v2, v8)
            v5 = v5
        } else {
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 uint = 4294967295
                var v11 uint8 = 40
                fmt.Println(v11, v10, fi9)
                _ = fi9
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 int = (fi12 / ((fi12) | 1))
                fmt.Println(v8)
                _ = fi12
                _ = v13
            }
        }
    case "":
        v14 := GenS2{F0: 32767}
        _ = v14
    }
    var v15 int16 = 1000
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int16 = (v15 | (v15 >> ((1) & 63)))
        _ = fi16
        _ = v17
    }
    v4 = (^(v4))
    v2 = float32((1e10 * 3.14159))
    if (1 >= (^(v4))) {
        v3 = append(v3, 2)
        fmt.Println(v2, v1, v4)
    }
    v18 := GenS1{F0: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
