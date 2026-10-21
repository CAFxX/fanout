package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
func genFn0() string {
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 float32 = ((1e10 - ((0.1 + -2.25) / (3.14159 + 1.5))) - 3.14159)
        _ = fi4
        _ = v5
    }
    switch 127 {
    case 42:
        v6 := GenS1{F0: 2147483647, F1: 32767}
        _ = v6
    case -748:
        var v7 float64 = -2.25
        _ = v7
    case 0:
        var v8 uint16 = 1
        _ = v8
    }
    return ("hello" + ("a" + "a"))
}

func main() {
    v1 := []int{-9223372036854775808, -53}
    v2 := make(map[string]bool)
    var v3 uint16 = 42
    v3 += 255
    var v9 uint16 = v3
    v3 = 0
    v10 := GenS1{F0: 227, F1: 32767}
    _ = len(v2)
    if ((0 >= -73) || true) {
        fmt.Println(v3, v9)
        var v11 int8 = 12
        if (true && false) {
            for fi12 := 0; fi12 < 20; fi12++ {
                v11 += -104
                _ = len(v2)
                for fi13 := 0; fi13 < 8; fi13++ {
                    var v14 int32 = 1
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
            switch ("\n\t" + "foo bar") {
            case "":
                fmt.Println(v9, v3)
            case "a":
                fmt.Println(v11, v9, v3)
            case "foo bar":
                _ = len(v2)
            }
            var v15 uint16 = v9
            _ = v15
        }
        _ = v11
    } else {
        if true {
            delete(v2, ("" + "foo bar"))
        }
    }
    v9 -= (v3 << ((1) & 63))
    var v16 int8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
