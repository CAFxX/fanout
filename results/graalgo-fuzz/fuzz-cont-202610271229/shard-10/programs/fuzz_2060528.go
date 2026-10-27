package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 127
    var v3 uint = 18446744073709551615
    var v4 uint16 = 65535
    fmt.Println(v2, v3)
    for fi5 := 0; fi5 < 11; fi5++ {
        v3 = (v3 % (((v3 & v3)) | 1))
        _ = fi5
    }
    v3 -= (v3 >> ((18446744073709551615) & 63))
    v3 += (v3 % (((v3 & (v3 - (v3 | (v3 << ((616) & 63)))))) | 1))
    var v6 uint16 = (v4 ^ (v4 + 188))
    v1[""] = false
    var v7 uint8 = 127
    var v8 int32 = 22
    v8 = -1
    v9 := GenS1{F0: (("foo bar" + "hello") + ("hello" + ("foo bar" + "foo bar"))), F1: 395}
    if (!false) {
        var v10 uint32 = 65535
        var v11 uint8 = uint8(v10)
        fmt.Println(v6, v11)
        _ = v10
        _ = v11
    } else {
        var v12 uint = (v3 ^ uint(v6))
        _ = v12
    }
    var v13 float32 = ((0.0 * 0.0) - (((1.5 - ((0.1 * (((0.1 - 0.0) * (1e10 * (float32((float64(3.14159) + 3.14159)) / ((3.14159 + 3.14159) / ((float32(1e10) + (3.14159 + (((3.14159 / 1.0) - float32(((-2.25 * 3.14159) + (-2.25 - 0.0)))) / 1.0))) - -2.25))))) + 1.5)) * 3.14159)) / (1e10 + ((1e10 / 1.0) + (1e10 + 3.14159)))) - 0.0))
    var v14 int32 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
