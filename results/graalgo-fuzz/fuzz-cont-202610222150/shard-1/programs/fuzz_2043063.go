package main

import (
    "fmt"
)

func genFn0(p0 uint32, p1 int32, p2 string) uint {
    p2 = p2
    fmt.Println(p2, p0)
    p2 = p2
    p1 -= 42
    var v6 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}

func main() {
    v1 := []uint32{2147483647, 65535, 65535, 554}
    var v2 int32 = -567
    var v3 int8 = 42
    var v4 uint64 = 467
    var v5 float32 = 0.0
    var v7 int32 = v2
    v7 -= int32(v4)
    v4 = 4294967295
    var v8 string = ("\n\t" + (("\n\t" + "\n\t") + "foo bar"))
    v7 -= v2
    _ = len(v1)
    switch 32768 {
    case -9223372036854775808:
        for fi9 := 0; fi9 < 18; fi9++ {
            for fi10 := 0; fi10 < 20; fi10++ {
                v8 = v8
                if len(v1) > 0 {
                    v1[0] = 514
                }
                _ = fi10
            }
            _ = len(v1)
            _ = fi9
        }
    }
    var v11 string = ("foo bar" + "foo bar")
    var v12 int32 = (v2 & 42)
    for fi13 := 0; fi13 < 18; fi13++ {
        v5 -= (v5 - v5)
        _ = fi13
    }
    v12 -= v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
