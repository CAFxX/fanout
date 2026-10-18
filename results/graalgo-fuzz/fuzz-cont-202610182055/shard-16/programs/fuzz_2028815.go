package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 bool }
type GenS2 struct { F0 uint; F1 uint8 }

func main() {
    v1 := []uint32{2, 4294967295}
    var v2 string = "hello"
    v3 := []string{"hello", "a", ""}
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: 51, F1: 255, F2: (3.14159 > (-2.25 + (-2.25 * 0.0)))}
    fmt.Println(v2)
    switch v2 {
    case "a":
        switch v2 {
        case "foo bar":
            var v6 float32 = (-2.25 / ((1.5 / 1.0) * 0.1))
            _ = v6
        case "a":
            fmt.Println(v2)
        case "hello":
            for fi7 := 0; fi7 < 17; fi7++ {
                v2 = v2
                v2 = ((((((v2 + "a") + ((v2 + "a") + (v2 + ("\n\t" + (("hello" + "hello") + ""))))) + v2) + v2) + v2) + "hello")
                v2 = ""
                _ = fi7
            }
        default:
            v8 := GenS2{F0: 95, F1: 42}
            _ = v8
        }
    }
    v2 = "foo bar"
    var v9 int16 = 1
    v2 = v2
    for fi10 := 0; fi10 < 2; fi10++ {
        if false {
            v11 := v4[2147483648]
            if ((42 >= 4294967295) && false) {
                v12 := GenS1{F0: 706, F1: 140, F2: (true && (((!(59 < 127)) && ((true || ((fi10 & fi10) >= (^(fi10)))) || true)) || true))}
                _ = v12
            }
            _ = v11
        } else {
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint16 = 0
                var v15 int32 = -310
                fmt.Println(v15)
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        var v16 int32 = 1
        _ = fi10
        _ = v16
    }
    var v17 uint16 = 1
    _ = len(v4)
    var v18 float64 = (1.5 / -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
