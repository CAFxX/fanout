package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int8 = 0
    var v2 string = (("hello" + "foo bar") + ((("" + ("\n\t" + ("a" + "hello"))) + "hello") + (("a" + "hello") + "a")))
    var v3 float32 = float32(-2.25)
    v4 := make(map[uint32]string)
    var v5 uint16 = 65535
    v6 := GenS1{F0: -81}
    var v7 uint32 = 2
    var v8 int8 = v1
    switch 4294967295 {
    case 32768:
        var v9 uint64 = 879
        _ = v9
    case 4294967295:
        var v10 string = v2
        _ = v10
    default:
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 int64 = 291
            for fi13 := 0; fi13 < 8; fi13++ {
                _ = len(v4)
                var v14 int8 = v1
                v15 := GenS1{F0: (v1 % ((-1) | 1))}
                _ = fi13
                _ = v14
                _ = v15
            }
            v16 := GenS1{F0: (v1 + v1)}
            _ = fi11
            _ = v12
            _ = v16
        }
    }
    v3 = v3
    v4[v7] = v2
    switch 4294967295 {
    case 0:
        var v17 int = -9223372036854775808
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
