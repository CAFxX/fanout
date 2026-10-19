package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }
type GenS2 struct { F0 bool; F1 int8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    v3 := []string{"a", "a"}
    var v4 bool = false
    switch "a" {
    case "hello":
        fmt.Println(v4)
    case "a":
        for fi5 := 0; fi5 < 4; fi5++ {
            v4 = (-2.25 != (1e10 - 0.0))
            if v4 {
                v4 = ("foo bar" > (("a" + ("a" + ("\n\t" + ("a" + "a")))) + "\n\t"))
            } else {
                for fi6 := 0; fi6 < 11; fi6++ {
                    if (1e10 < 0.0) {
                        if v4 {
                            var v7 int = fi5
                            _ = v7
                        } else {
                            v4 = v4
                        }
                        v8 := GenS1{F0: 1, F1: 842, F2: 748}
                        v4 = (0 != 128)
                        _ = v8
                    }
                    _ = fi6
                }
                v9 := GenS2{F0: ((^(fi5)) < -329), F1: -8}
                _ = v9
            }
            v10 := GenS2{F0: false, F1: 127}
            _ = fi5
            _ = v10
        }
    default:
        for fi11 := 0; fi11 < 12; fi11++ {
            v4 = false
            v4 = (0.1 <= float32(3.14159))
            _ = fi11
        }
    }
    v4 = v4
    v4 = (!v4)
    v12 := GenS1{F0: 127, F1: 42, F2: 1}
    fmt.Println(v4)
    var v13 float64 = 3.14159
    if (32767 > 0) {
        v4 = (!true)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
