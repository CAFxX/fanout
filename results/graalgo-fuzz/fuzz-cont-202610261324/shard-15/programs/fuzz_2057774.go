package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 string = "hello"
    v2 := make(map[int]uint32)
    var v3 string = (v1 + "")
    v3 = v1
    switch (v3 + ("" + "")) {
    case "foo bar":
        v4 := GenS1{F0: 421}
        _ = v4
    case "hello":
        v2[-9223372036854775808] = 2
    }
    fmt.Println(v1, v1)
    var v5 string = (((("a" + v1) + v1) + "\n\t") + v1)
    v6 := GenS1{F0: 42}
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(v5, v3)
        _ = fi7
    }
    fmt.Println(v5)
    var v8 int = -1
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(fi9, v8, v1)
        _ = fi9
    }
    v2[v8] = 65535
    switch ((v1 + "\n\t") + v3) {
    case "\n\t":
        var v10 int64 = -9223372036854775808
        _ = v10
    case "a":
        if false {
            v5 = (v3 + (v5 + v3))
        } else {
            switch (-(v8)) {
            case 128:
                var v11 uint16 = 87
                _ = v11
            case -59:
                v5 = "a"
            case -1:
                for fi12 := 0; fi12 < 8; fi12++ {
                    v8 += v8
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v3 = ("a" + v3)
                        fmt.Println(fi13, v3, v3)
                        _ = fi13
                    }
                    v14 := GenS1{F0: 427}
                    _ = fi12
                    _ = v14
                }
            default:
                v5 = ((v3 + v3) + "foo bar")
            }
        }
    default:
        v2[127] = 2
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v8
}
