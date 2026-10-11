package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }

func main() {
    v1 := []uint32{65535, 176}
    v2 := []string{("\n\t" + ""), ((("\n\t" + "\n\t") + ("" + ("hello" + "a"))) + "a")}
    v3 := make(map[string]uint32)
    v4 := make(map[int]uint32)
    v5 := []uint32{}
    var v6 int8 = 100
    var v7 bool = ((168 == 84) && (511 != 4294967295))
    if v7 {
        v8 := GenS1{F0: 255}
        _ = v8
    } else {
        v6 -= v6
        v6 = (v6 & 1)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v9 := GenS2{F0: 230}
    fmt.Println(v6)
    if ("" <= "hello") {
        v2 = append(v2, ("" + ("hello" + "")))
        v10 := GenS3{F0: 9223372036854775807}
        _ = v10
    } else {
        if (616 < 934) {
            fmt.Println(v7)
            var v11 uint = 4294967295
            _ = v11
        }
    }
    v7 = (-2.25 >= (-2.25 * 1.5))
    for fi12 := 0; fi12 < 18; fi12++ {
        v7 = v7
        _ = fi12
    }
    var v13 uint8 = 230
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
