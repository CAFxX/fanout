package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint16 = 42
    var v3 int32 = 810
    v4 := []string{}
    var v5 float32 = 0.1
    v6 := []string{("" + "hello")}
    v7 := GenS1{F0: 2}
    if len(v4) > 0 {
        v4[0] = (("a" + ((("foo bar" + "") + "hello") + (("\n\t" + ("a" + ("a" + "hello"))) + ""))) + ("\n\t" + "hello"))
    }
    if len(v4) > 0 {
        v4[0] = "a"
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v8 bool = false
    v5 = 3.14159
    v6 = append(v6, "a")
    v2 = (v2 + 65535)
    fmt.Println(v8)
    v3 = v3
    fmt.Println(v2, v1, v1)
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 float64 = float64(v5)
            var v12 uint32 = 2147483647
            fmt.Println(v2)
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 float64 = -2.25
        _ = fi9
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
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
