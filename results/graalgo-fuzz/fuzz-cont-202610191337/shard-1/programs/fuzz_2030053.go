package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int32 }

func main() {
    var v1 int = 32767
    var v2 int8 = -1
    v3 := []uint32{816, 65535}
    v4 := GenS1{F0: "", F1: 0, F2: -167}
    var v5 int16 = 574
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(v5, v5)
        fmt.Println(v1, fi6)
        var v7 uint32 = 83
        _ = fi6
        _ = v7
    }
    switch ("" + "hello") {
    case "\n\t":
        v5 += v5
    }
    v8 := GenS1{F0: ("\n\t" + "foo bar"), F1: 2147483647, F2: -805}
    var v9 float32 = -2.25
    var v10 string = "hello"
    for fi11 := 0; fi11 < 8; fi11++ {
        v2 += (v2 >> ((255) & 63))
        _ = fi11
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v5)
    v5 += v5
    if false {
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
