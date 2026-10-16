package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{0, 12, 42}
    v2 := []string{"", ("a" + "a"), (("a" + "hello") + "\n\t"), ("\n\t" + "hello")}
    v3 := make(map[string]int)
    var v4 bool = true
    var v5 uint8 = 0
    v4 = false
    switch 0 {
    case 214:
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 uint16 = 1
            _ = len(v1)
            var v8 uint16 = 853
            _ = fi6
            _ = v7
            _ = v8
        }
    case 32768:
        var v9 int64 = 348
        _ = v9
    case 4294967295:
        var v10 uint8 = (v5 % ((v5) | 1))
        _ = v10
    default:
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := v3[(((((("" + "foo bar") + "hello") + "foo bar") + "a") + "\n\t") + ((("foo bar" + "foo bar") + "") + ""))]
            var v13 uint64 = 223
            v5 += 62
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    var v14 int8 = -1
    var v15 int8 = v14
    v15 -= -1
    v4 = true
    v14 = -1
    v14 += v14
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
