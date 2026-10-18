package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 983
    var v3 int16 = 0
    v4 := make(map[string]string)
    v5 := []string{"", "", "foo bar", "a"}
    var v6 int8 = 0
    v7 := GenS1{F0: -1}
    var v8 bool = true
    switch 4294967295 {
    case 9223372036854775807:
        var v9 int = -1
        _ = v9
    case 32767:
        if false {
            var v10 int32 = -2147483648
            v11 := GenS1{F0: 42}
            _ = v10
            _ = v11
        } else {
            v8 = true
        }
    case -1:
        for fi12 := 0; fi12 < 15; fi12++ {
            if len(v5) > 0 {
                v5[0] = ("foo bar" + "a")
            }
            v2 = v2
            _ = fi12
        }
    }
    v3 = (v3 % ((v3) | 1))
    fmt.Println(v8, v8, v6)
    var v13 uint16 = v2
    v2 -= 65535
    fmt.Println(v3, v3, v8)
    var v14 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
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
}
