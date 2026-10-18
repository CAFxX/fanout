package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 1
    v2 := make(map[string]string)
    var v3 uint64 = 4294967295
    v4 := []int{4294967295}
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    var v5 uint64 = (^(v3))
    var v6 int32 = -30
    delete(v2, (("\n\t" + "a") + "\n\t"))
    var v7 uint64 = (v5 | (^(v5)))
    var v8 uint = 4294967295
    if false {
        _ = len(v4)
    } else {
        var v9 int16 = -1
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v3 = v3
        delete(v2, ("hello" + "a"))
        switch "a" {
        case "hello":
            v7 -= (^(v7))
        case "a":
            var v11 int16 = -32768
            _ = v11
        }
        _ = fi10
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 uint32 = 361
        _ = fi12
        _ = v13
    }
    var v14 uint64 = (^(v5))
    if len(v4) > 0 {
        v4[0] = int(v6)
    }
    v1 = 15
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
