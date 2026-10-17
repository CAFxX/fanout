package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{4294967295, 4294967295}
    var v2 float32 = (0.1 / 1.5)
    v3 := make(map[string]string)
    var v4 int64 = 0
    v3[(("foo bar" + (("hello" + "") + ("\n\t" + "hello"))) + "")] = "foo bar"
    fmt.Println(v4)
    if ((false && true) || true) {
        var v5 uint8 = 94
        _ = v5
    }
    var v6 uint64 = 937
    v4 = v4
    if false {
        v4 += (v4 % ((1) | 1))
        if len(v1) > 0 {
            v1[0] = 2
        }
        if false {
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 uint8 = 255
                _ = fi7
                _ = v8
            }
        }
    } else {
        v4 -= (v4 + 193)
        var v9 bool = ((v4 - (^(v4))) != 1)
        _ = v9
    }
    v2 = float32(3.14159)
    v10 := v3["hello"]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
