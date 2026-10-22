package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 string = "hello"
    v2 := []uint32{671}
    v3 := []int{0, -9223372036854775808}
    v4 := []int{9223372036854775807, 127}
    var v5 bool = (!(1 <= 998))
    v6 := make(map[int]string)
    var v7 bool = true
    if (42 > 282) {
        _ = len(v4)
    }
    var v8 uint = 255
    var v9 int16 = 42
    switch "\n\t" {
    case "foo bar":
        if (6 < 377) {
            v10 := GenS1{F0: (v9 + v9)}
            for fi11 := 0; fi11 < 1; fi11++ {
                v7 = ((255 < 74) || (("a" + "a") == (v1 + v1)))
                _ = len(v3)
                _ = fi11
            }
            v12 := GenS1{F0: (v9 + (^(v9)))}
            _ = v10
            _ = v12
        }
    case "hello":
        v13 := GenS1{F0: (v9 * v9)}
        _ = v13
    }
    switch ("a" + "\n\t") {
    case "hello":
        v14 := v6[1]
        _ = v14
    case "":
        v1 = "hello"
    }
    var v15 int = -9223372036854775808
    v16 := GenS1{F0: v9}
    if (v15 != v15) {
        var v17 int16 = (v9 >> (((^(v8))) & 63))
        for fi18 := 0; fi18 < 12; fi18++ {
            fmt.Println(v5)
            var v19 uint8 = 127
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    if (!((3.14159 <= (-2.25 / 1.0)) && (571 == 768))) {
        var v20 int64 = -822
        fmt.Println(v20)
        _ = v20
    } else {
        v5 = v5
        v8 -= 0
        delete(v6, v15)
    }
    var v21 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
