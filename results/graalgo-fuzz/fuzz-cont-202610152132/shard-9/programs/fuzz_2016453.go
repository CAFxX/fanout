package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int = 4294967295
    v2 := make(map[string]string)
    v3 := []uint32{673, 65535, 65535, 0}
    var v4 uint = 988
    var v5 uint8 = 1
    v6 := GenS1{F0: (false || (2147483648 >= 2147483648))}
    v7 := GenS1{F0: (127 < 1)}
    fmt.Println(v5)
    v4 = (v4 - (v4 ^ uint(v5)))
    v8 := GenS1{F0: (v4 < (v4 >> ((v4) & 63)))}
    var v9 int = (v1 % (((v1 * v1)) | 1))
    v5 -= v5
    if true {
        var v10 uint16 = 21
        switch "\n\t" {
        case "\n\t":
            v9 = v1
        case "":
            v1 -= -1
        case "foo bar":
            v11 := v2["a"]
            _ = v11
        default:
            v5 += (v5 & (v5 + v5))
        }
        _ = v10
    } else {
        v9 += (-(v1))
    }
    delete(v2, ("hello" + ("a" + "hello")))
    v12 := GenS1{F0: (1 >= 9223372036854775807)}
    fmt.Println(v5, v1, v1)
    if (true || (!(!(2 == 42)))) {
        if (!(true || (161 >= 42))) {
            if ((255 != 255) || false) {
                var v13 int64 = -1
                v13 -= (^(v13))
                _ = v13
            }
            v4 = v4
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
