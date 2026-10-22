package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 float64 = 3.14159
    var v2 string = ("" + "hello")
    v3 := make(map[uint32]bool)
    v4 := GenS1{F0: 1}
    v3[65535] = (-54 > -86)
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := v3[134]
        _ = fi5
        _ = v6
    }
    v1 -= 0.0
    var v7 int16 = -479
    if (int(v7) >= 127) {
        if true {
            v2 = v2
            v2 = ""
        }
        fmt.Println(v7)
        v8 := v3[2]
        _ = v8
    }
    v2 = v2
    v3[635] = true
    v1 += -2.25
    v2 = ("\n\t" + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
