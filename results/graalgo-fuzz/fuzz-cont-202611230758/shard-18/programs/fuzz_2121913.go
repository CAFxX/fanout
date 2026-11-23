package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 0
    v3 := make(map[int]string)
    var v4 int32 = -1
    var v5 int8 = 100
    var v6 int16 = -868
    v2 = (v2 + 833)
    v5 = 73
    if (false && true) {
        v5 = v5
    }
    v5 += v5
    fmt.Println(v5, v2)
    v3[1] = ("\n\t" + "")
    var v7 int16 = 42
    var v8 int8 = 100
    v7 = 1000
    switch (("" + "\n\t") + (("foo bar" + "a") + ("a" + ("foo bar" + "")))) {
    case "foo bar":
        var v9 uint = 765
        _ = v9
    case "\n\t":
        v5 = (v8 / ((v5) | 1))
    case "a":
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 int16 = v7
            if (!false) {
                var v12 int = (-(fi10))
                for fi13 := 0; fi13 < 20; fi13++ {
                    v11 += v6
                    _ = fi13
                }
                var v14 int32 = int32(v2)
                _ = v12
                _ = v14
            }
            fmt.Println(v4, v4, v4)
            _ = fi10
            _ = v11
        }
    default:
        _ = len(v1)
    }
    v4 += int32(v8)
    v4 = (v4 | -1)
    v4 -= 450
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
