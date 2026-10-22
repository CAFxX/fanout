package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []uint32{853, 0, 891, 343}
    var v2 uint32 = 2
    v3 := make(map[uint32]uint32)
    fmt.Println(v2)
    v2 = (v2 >> ((4294967295) & 63))
    var v4 float64 = 3.14159
    for fi5 := 0; fi5 < 3; fi5++ {
        v2 += 65535
        _ = fi5
    }
    if (v2 == 4294967295) {
        v6 := v3[v2]
        for fi7 := 0; fi7 < 7; fi7++ {
            v4 = v4
            v6 = 65535
            _ = fi7
        }
        _ = v6
    }
    var v8 int16 = -978
    v1 = append(v1, (v2 % ((v2) | 1)))
    v1 = append(v1, (v2 - (^(v2))))
    var v9 uint8 = 128
    v10 := GenS1{F0: (^(v2))}
    if (true || (true || true)) {
        fmt.Println(v4, v4)
    } else {
        delete(v3, 65535)
        v4 = (3.14159 - (v4 + float64(0.0)))
    }
    switch "" {
    case "\n\t":
        fmt.Println(v2, v4, v2)
    case "foo bar":
        v2 = 203
    case "a":
        for fi11 := 0; fi11 < 10; fi11++ {
            v2 = v2
            fmt.Println(fi11, v2)
            var v12 int = 32767
            _ = fi11
            _ = v12
        }
    }
    v4 = v4
    v4 = v4
    switch (("foo bar" + "a") + "a") {
    case "\n\t":
        if (((1 >= 127) || ((v4 / float64(1e10)) >= v4)) || (42 >= 0)) {
            v3[v2] = v2
            v9 += 115
        }
    case "":
        var v13 int = 4294967295
        _ = v13
    case "a":
        v8 += -388
    default:
        v14 := GenS1{F0: (v2 / ((v2) | 1))}
        _ = v14
    }
    if true {
        v4 += v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
