package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    var v3 float64 = 1.5
    var v4 string = ((("foo bar" + "\n\t") + "hello") + "")
    v5 := []int{351, -793, 0}
    v6 := make(map[int]string)
    var v7 uint64 = 454
    _ = len(v1)
    if true {
        v4 = v4
    } else {
        if false {
            v8 := GenS1{F0: (!(true || (false && true)))}
            _ = v8
        }
        var v9 uint64 = 4294967295
        _ = v9
    }
    if (false && (-619 >= -1)) {
        if (0 == 32767) {
            var v10 uint32 = 650
            if len(v5) > 0 {
                v5[0] = 281
            }
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := GenS1{F0: (v10 > (v10 >> ((338) & 63)))}
                _ = fi11
                _ = v12
            }
            _ = v10
        }
    }
    v13 := GenS1{F0: (0.0 != (1e10 - (3.14159 - 0.1)))}
    var v14 string = v4
    switch (v14 + v14) {
    case "":
        var v15 uint = 65535
        _ = v15
    case "a":
        var v16 string = v4
        _ = v16
    default:
        fmt.Println(v7, v7, v14)
    }
    var v17 uint8 = 0
    var v18 float64 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
