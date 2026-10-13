package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 bool }

func main() {
    var v1 bool = (!(!true))
    v2 := []uint32{263, 42, 2}
    var v3 int16 = -313
    v4 := make(map[uint32]bool)
    var v5 int8 = 0
    var v6 uint32 = 2147483648
    delete(v4, (v6 + 832))
    v6 = v6
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    _ = len(v2)
    if (!v1) {
        _ = len(v4)
        var v7 uint16 = 927
        if (v3 < -1) {
            v6 = v6
        }
        _ = v7
    } else {
        var v8 float64 = float64((0.0 + 3.14159))
        var v9 int8 = (v5 << ((42) & 63))
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 81, F1: "", F2: ((46 > 128) || false)}
    switch ("foo bar" + (("\n\t" + "") + "a")) {
    case "\n\t":
        var v11 bool = (42 > 0)
        _ = v11
    case "":
        if (42 > 1) {
            fmt.Println(v6)
            for fi12 := 0; fi12 < 13; fi12++ {
                v3 -= -32768
                _ = fi12
            }
        }
    case "foo bar":
        v3 -= v3
    }
    v3 -= (v3 >> ((65535) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
