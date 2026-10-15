package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint8 = 2
    v2 := []string{("foo bar" + ("hello" + "foo bar"))}
    v3 := make(map[string]string)
    var v4 int32 = -1
    var v5 int64 = -116
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := v3[("\n\t" + (("hello" + ("\n\t" + "")) + (("foo bar" + "") + ((((("a" + "") + "hello") + "") + "a") + ("" + "a")))))]
        _ = fi6
        _ = v7
    }
    switch 32768 {
    case -9223372036854775808:
        v2 = append(v2, "\n\t")
    default:
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    }
    var v8 uint = 255
    v8 = (v8 ^ v8)
    var v9 float32 = (-2.25 / 0.1)
    for fi10 := 0; fi10 < 18; fi10++ {
        v4 = (^(v4))
        if (2 <= 65535) {
            var v11 float64 = float64(1.5)
            _ = v11
        } else {
            var v12 uint64 = 436
            fmt.Println(fi10, v4)
            _ = v12
        }
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS1{F0: (^(v5))}
        var v15 int16 = 1000
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v8)
    v9 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
