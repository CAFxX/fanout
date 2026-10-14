package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float64 = 3.14159
    v2 := []string{}
    var v3 uint32 = 65535
    v4 := make(map[int]int)
    var v5 uint8 = 255
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(v3, v5, v1)
        v7 := v4[(-(fi6))]
        v7 += v7
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 0}
    fmt.Println(v1, v5, v5)
    var v9 uint32 = v3
    fmt.Println(v3, v5)
    v10 := v4[32767]
    if true {
        fmt.Println(v1, v3)
        if false {
            v4[(v10 * (v10 * v10))] = 529
            switch (-(v10)) {
            case 32768:
                var v11 int8 = 1
                _ = v11
            case -9223372036854775808:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 1:
                v10 = v10
            }
            var v12 int = (v10 % ((v10) | 1))
            _ = v12
        }
        _ = len(v2)
    } else {
        var v13 float32 = -2.25
        delete(v4, v10)
        var v14 uint32 = v3
        _ = v13
        _ = v14
    }
    v3 += 0
    if len(v2) > 0 {
        v2[0] = "a"
    }
    switch "\n\t" {
    case "hello":
        v15 := GenS1{F0: 32767}
        _ = v15
    }
    var v16 uint = 901
    fmt.Println(v5)
    if len(v2) > 0 {
        v2[0] = ((("a" + "") + "") + "")
    }
    var v17 uint64 = 42
    var v18 uint = (v16 / (((v16 + uint(v5))) | 1))
    var v19 uint64 = 18446744073709551615
    v16 = uint(v19)
    var v20 float64 = ((0.0 + 3.14159) + float64((0.0 / float32(0.1))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
