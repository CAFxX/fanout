package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint64 }
func genFn0(p0 int32, p1 uint64, p2 bool) uint {
    fmt.Println(p1, p1, p1)
    if p2 {
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 uint8 = 59
            _ = fi5
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int = (fi7 - -222)
        v9 := []int{(-(v8)), 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 string = "hello"
    var v2 float32 = 3.14159
    var v3 uint = 281
    var v4 int16 = -32768
    switch ("\n\t" + ((v1 + ((v1 + v1) + v1)) + "")) {
    case "hello":
        var v10 int8 = 0
        _ = v10
    case "":
        v11 := GenS1{F0: 255, F1: 735, F2: 2}
        _ = v11
    case "foo bar":
        var v12 int = 272
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := []int{}
            _ = len(v15)
            _ = fi14
            _ = v15
        }
        if ((fi13 >= fi13) || (true || true)) {
            var v16 uint32 = 65535
            v1 = (v1 + ((("foo bar" + v1) + v1) + (v1 + ("hello" + v1))))
            _ = v16
        } else {
            v4 += -521
        }
        _ = fi13
    }
    var v17 float64 = 3.14159
    v18 := GenS1{F0: 162, F1: 128, F2: 279}
    v4 += 0
    var v19 bool = (255 == 5)
    v2 = v2
    v20 := []int{502, 127}
    if v19 {
        _ = len(v20)
        v20 = append(v20, -444)
    } else {
        var v21 uint64 = 4294967295
        switch v1 {
        case "a":
            v2 = ((3.14159 / 1.0) - v2)
        default:
            v4 = (-(v4))
        }
        _ = v21
    }
    if len(v20) > 0 {
        v20[0] = 584
    }
    var v22 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
}
