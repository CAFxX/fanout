package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int16 = 481
    v2 := make(map[string]string)
    var v3 uint16 = 2
    v4 := []string{((("" + "") + "foo bar") + "\n\t"), "a"}
    var v5 int16 = v1
    fmt.Println(v3)
    var v6 string = "foo bar"
    fmt.Println(v1, v3, v3)
    var v7 int = 127
    fmt.Println(v3, v5, v6)
    fmt.Println(v5, v7)
    v8 := GenS1{F0: 0}
    if (!true) {
        for fi9 := 0; fi9 < 16; fi9++ {
            switch -1 {
            case -106:
                var v10 float64 = 0.0
                _ = v10
            case -2147483648:
                var v11 uint8 = 42
                _ = v11
            case -852:
                v12 := GenS2{F0: 864}
                _ = v12
            default:
                var v13 int32 = 1
                _ = v13
            }
            _ = fi9
        }
        v6 = (v6 + (v6 + (v6 + ((("hello" + "a") + ("a" + (v6 + v6))) + v6))))
    }
    var v14 uint64 = 171
    var v15 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
