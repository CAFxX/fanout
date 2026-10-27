package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint16 }

func main() {
    v1 := []string{}
    v2 := []uint32{966, 2, 1, 42}
    var v3 float32 = 3.14159
    var v4 uint = 0
    v5 := []string{(("foo bar" + "\n\t") + (("a" + "a") + "\n\t")), "hello"}
    var v6 uint = v4
    if true {
        v7 := GenS1{F0: 4294967295, F1: (v6 + 29), F2: 32767}
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 int8 = 66
            _ = fi8
            _ = v9
        }
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        _ = v7
    } else {
        v6 = 4294967295
    }
    v3 = (v3 / ((v3 - v3) + 1.5))
    v6 = 18446744073709551615
    var v10 uint = 65535
    v6 -= v10
    var v11 string = (("hello" + (("foo bar" + "hello") + ("foo bar" + "foo bar"))) + "foo bar")
    v6 -= 65535
    v10 -= (v4 & v6)
    var v12 int16 = 0
    var v13 float64 = -2.25
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v14 uint = v10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
