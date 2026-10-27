package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = ("a" < "foo bar")
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]string)
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(fi6, fi6, v3)
        _ = fi6
    }
    v7 := GenS1{F0: 340}
    if (501 <= 2) {
        for fi8 := 0; fi8 < 10; fi8++ {
            if (0 > 255) {
                var v9 int8 = -1
                v9 = v9
                _ = v9
            }
            _ = fi8
        }
    } else {
        var v10 float64 = 3.14159
        _ = v10
    }
    v2 = (-108 <= -2147483648)
    v2 = ((666 != -1) || v2)
    v5["foo bar"] = (("foo bar" + "\n\t") + ("" + ""))
    var v11 int64 = 42
    v12 := v1[32767]
    v13 := v5["\n\t"]
    _ = len(v1)
    v11 = -764
    delete(v5, "hello")
    var v14 int16 = -1
    var v15 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
