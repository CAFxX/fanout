package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]string)
    var v3 uint = 255
    var v4 uint64 = 18446744073709551615
    v5 := make(map[uint32]uint32)
    var v6 string = ("a" + "a")
    if (v6 < ((v6 + v6) + v6)) {
        var v7 uint8 = 42
        _ = v7
    } else {
        var v8 int = -222
        _ = v8
    }
    v4 = 4294967295
    v4 += (v4 * v4)
    if false {
        var v9 uint8 = 255
        _ = v9
    }
    v10 := GenS1{F0: ("foo bar" + "\n\t")}
    fmt.Println(v6, v6, v4)
    if true {
        var v11 uint32 = 540
        v11 += 361
        _ = v11
    } else {
        switch 9223372036854775807 {
        case 127:
            v4 = (v4 - v4)
        case -2147483648:
            if (!(1 == 42)) {
                v12 := GenS1{F0: ""}
                _ = v12
            }
        }
        v6 = v6
    }
    switch 9223372036854775807 {
    case -1:
        v3 = v3
    }
    var v13 int64 = -361
    var v14 int64 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
