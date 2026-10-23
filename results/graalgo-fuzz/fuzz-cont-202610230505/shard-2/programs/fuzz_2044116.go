package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 string }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint, p1 uint64) uint64 {
    if true {
        for fi4 := 0; fi4 < 8; fi4++ {
            p1 = 193
            _ = fi4
        }
        p1 += 42
        fmt.Println(p0)
    }
    var v5 uint = (p0 + p0)
    if (!((!false) || (-232 < 4294967295))) {
        p1 = 399
    } else {
        v6 := GenS1{F0: 255, F1: "hello", F2: "hello"}
        v5 = v5
        fmt.Println(p1, p1)
        _ = v6
    }
    v7 := GenS2{F0: -128}
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 507
}

func main() {
    v1 := []string{"", ("\n\t" + "foo bar"), "a", ("hello" + ("\n\t" + (((("\n\t" + ("foo bar" + "")) + "a") + "\n\t") + ("hello" + "foo bar"))))}
    var v2 bool = (42 >= 290)
    var v3 float32 = 0.1
    if len(v1) > 0 {
        v1[0] = ("hello" + "")
    }
    v2 = false
    _ = len(v1)
    v2 = v2
    var v8 int8 = 0
    var v9 int = 2147483647
    var v10 uint = 779
    v11 := GenS2{F0: (v8 << ((v10) & 63))}
    var v12 int64 = -1
    if true {
        v8 = (v8 ^ (v8 ^ -54))
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 float32 = (v3 / float32(1.5))
        fmt.Println(v14, v14)
        fmt.Println(fi13, v14)
        _ = fi13
        _ = v14
    }
    var v15 int8 = 1
    var v16 uint8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
