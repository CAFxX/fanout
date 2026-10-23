package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{921, 2147483647, 735}
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    var v4 float32 = (0.1 - 3.14159)
    var v5 uint64 = 2
    var v6 int = 32767
    if true {
        var v7 int64 = -918
        var v8 bool = ((32767 >= 42) && (102 == 0))
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 uint8 = 127
            for fi11 := 0; fi11 < 3; fi11++ {
                fmt.Println(v4, v6)
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
    }
    var v12 int = 4294967295
    var v13 float32 = 0.0
    var v14 int16 = int16(v6)
    var v15 int32 = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
