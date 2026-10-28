package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint8 }
func genFn0(p0 int) int16 {
    fmt.Println(p0)
    var v6 uint16 = 0
    _ = p0
    _ = v6
    return -959
}
func genFn1(p0 int64, p1 int64, p2 uint16) int {
    v7 := GenS1{F0: 42, F1: (-32768 < 473), F2: 255}
    var v8 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]bool)
    var v3 bool = (42 >= 255)
    var v4 int = -951
    var v5 float32 = 1.5
    fmt.Println(v1, v3, v1)
    var v9 int32 = -1
    v3 = ("a" < ("foo bar" + ("hello" + "foo bar")))
    fmt.Println(v1, v4, v1)
    v1 = v1
    for fi10 := 0; fi10 < 2; fi10++ {
        v5 = 3.14159
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 uint16 = 32767
            fmt.Println(v5)
            var v13 bool = true
            var v14 float64 = (0.1 * v1)
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int = (v4 - fi15)
        _ = fi15
        _ = v16
    }
    v3 = ((1 >= 774) && ((4294967295 < 0) || true))
    v17 := GenS1{F0: 2147483647, F1: v3, F2: 0}
    v18 := v2[597]
    fmt.Println(v9, v1, v3)
    _ = len(v2)
    for fi19 := 0; fi19 < 6; fi19++ {
        fmt.Println(v18, fi19)
        switch (fi19 / (((^(v4))) | 1)) {
        case 731:
            fmt.Println(v5, fi19)
        case 9223372036854775807:
            v20 := v2[(-(fi19))]
            _ = v20
        }
        _ = fi19
    }
    fmt.Println(v18, v9, v3)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
