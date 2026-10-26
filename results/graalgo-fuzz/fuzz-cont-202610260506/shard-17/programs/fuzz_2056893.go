package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint16 }
func genFn0() string {
    if false {
        var v5 uint32 = 418
        _ = v5
    }
    var v6 uint = 255
    fmt.Println(v6)
    v7 := GenS1{F0: 65535, F1: -21, F2: 255}
    _ = v6
    _ = v7
    return ("" + "a")
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = -553
    v3 := []string{"foo bar", ("" + (("hello" + "") + "")), "", ("foo bar" + "hello")}
    var v4 uint64 = 2
    if false {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int32 = 42
            var v10 float32 = (((-2.25 - -2.25) + 1.5) / ((0.0 - float32(((0.1 - (3.14159 / 1.0)) + 0.1))) * -2.25))
            _ = fi8
            _ = v10
            _ = v9
        }
        switch 593 {
        case 32767:
            v4 -= v4
        default:
            v2 -= (v2 >> ((484) & 63))
        }
        v3 = append(v3, (("foo bar" + ("hello" + "foo bar")) + (("" + ("foo bar" + "\n\t")) + "\n\t")))
    }
    v1[("hello" + "\n\t")] = 42
    fmt.Println(v2, v2)
    delete(v1, "hello")
    fmt.Println(v2, v4)
    var v11 float64 = 1e10
    var v12 string = "hello"
    var v13 string = ("hello" + (v12 + v12))
    fmt.Println(v4, v12, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
