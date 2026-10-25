package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int16, p1 uint16) int16 {
    p1 -= 2
    for fi4 := 0; fi4 < 12; fi4++ {
        if false {
            fmt.Println(p1, p1)
            v5 := GenS1{F0: 4294967295}
            p0 = (p0 | p0)
            _ = v5
        }
        _ = fi4
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    v1 := []string{}
    v2 := make(map[string]int)
    var v3 float32 = 0.0
    v3 = float32(0.1)
    _ = len(v2)
    delete(v2, ("a" + "foo bar"))
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: 255}
        switch "a" {
        case "":
            v8 := GenS1{F0: 1}
            _ = v8
        default:
            var v9 int = (fi6 << ((255) & 63))
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v3 = 3.14159
    v3 -= 3.14159
    fmt.Println(v3)
    var v10 uint16 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
