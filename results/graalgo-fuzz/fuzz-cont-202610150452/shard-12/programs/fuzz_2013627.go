package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }

func main() {
    var v1 string = ("foo bar" + (("\n\t" + "") + "a"))
    var v2 int16 = 55
    var v3 uint32 = 1
    var v4 int32 = 512
    v5 := GenS1{F0: v3, F1: (v2 >> ((2) & 63))}
    v3 = (v3 << ((102) & 63))
    v6 := GenS1{F0: (^(v3)), F1: v2}
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int = (fi7 | (fi7 | fi7))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: (v3 | 2147483647), F1: (v2 >> ((42) & 63))}
        _ = fi9
        _ = v10
    }
    v11 := []int{32768, 0, -611, -196}
    switch (v1 + "\n\t") {
    case "a":
        fmt.Println(v4, v4)
    case "":
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    var v12 uint64 = 678
    v4 = 0
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v13 int8 = 0
    fmt.Println(v2, v2, v3)
    v4 = (-(v4))
    var v14 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
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
