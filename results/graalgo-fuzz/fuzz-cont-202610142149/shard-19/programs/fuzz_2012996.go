package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int8 = 64
    v2 := make(map[string]uint32)
    v3 := []int{}
    var v4 uint = 42
    var v5 uint64 = 4294967295
    v6 := []string{("hello" + ("\n\t" + ("\n\t" + ("foo bar" + "foo bar")))), (("\n\t" + "hello") + "foo bar")}
    if true {
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 float32 = (-2.25 / -2.25)
            _ = fi7
            _ = v8
        }
    }
    v1 -= 121
    v5 += 866
    switch 128 {
    case 32767:
        switch 32768 {
        case 32767:
            v9 := GenS1{F0: 128}
            _ = v9
        case -154:
            v4 = 18446744073709551615
        }
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint = v4
        v11 = 2
        v11 += (v11 % ((v4) | 1))
        _ = fi10
        _ = v11
    }
    v12 := v2["a"]
    v4 = 0
    var v13 float32 = float32(3.14159)
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint8 = 1
        var v16 string = ""
        v17 := GenS1{F0: v15}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
