package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint {
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 127}
        _ = fi5
        _ = v6
    }
    var v7 int16 = -1
    var v8 string = "\n\t"
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint, p1 string) uint16 {
    p1 = (p1 + "hello")
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v9
    return 65535
}

func main() {
    v1 := []uint32{50, 42, 0}
    var v2 int16 = 0
    var v3 uint16 = 2
    v4 := make(map[string]string)
    fmt.Println(v2)
    v3 = (v3 - 1)
    if true {
        v4[("\n\t" + ("" + ("foo bar" + ("foo bar" + "foo bar"))))] = "hello"
        switch "" {
        case "a":
            if true {
                fmt.Println(v3, v3)
                v10 := GenS1{F0: 100}
                _ = v10
            }
        default:
            fmt.Println(v3)
        }
        v1 = append(v1, 42)
    } else {
        _ = len(v4)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v11 := v4["foo bar"]
    var v12 float32 = float32(3.14159)
    fmt.Println(v2, v11)
    var v13 bool = true
    if v13 {
        v13 = true
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v3 = v3
        _ = fi14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
