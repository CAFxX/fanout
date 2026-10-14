package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 2
    var v2 uint32 = v1
    var v3 float32 = -2.25
    var v4 uint = 711
    v5 := make(map[uint32]string)
    var v6 int8 = 100
    fmt.Println(v1, v6, v3)
    fmt.Println(v3)
    var v7 int8 = v6
    var v8 uint16 = 255
    if (((3.14159 * 0.0) + 3.14159) >= v3) {
        _ = len(v5)
    } else {
        var v9 int = -1
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 string = ("a" + ("" + "\n\t"))
        var v12 uint64 = 1
        v7 += 0
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        var v13 int32 = 42
        switch -1 {
        case 32768:
            for fi14 := 0; fi14 < 10; fi14++ {
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 int8 = -89
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
        case 2147483647:
            v3 = v3
        }
        v3 = 0.1
        _ = v13
    } else {
        v6 = 127
    }
    v4 += (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
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
