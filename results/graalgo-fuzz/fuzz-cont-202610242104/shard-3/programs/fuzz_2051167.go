package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[int]uint32)
    var v3 float32 = 0.0
    var v4 int = 32767
    var v5 int64 = 904
    if true {
        fmt.Println(v3, v3)
    } else {
        v2[v4] = 970
        if (148 >= 2) {
            var v6 int64 = (v5 >> ((65535) & 63))
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 int16 = -894
            v5 += -1
            delete(v2, fi8)
            _ = fi8
            _ = v9
        }
        v1 = append(v1, "\n\t")
        _ = fi7
    }
    v3 = float32(((float64(v3) - -2.25) - (0.0 * 0.0)))
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 float32 = v3
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 uint = 18446744073709551615
        delete(v2, (fi12 + (fi12 & v4)))
        fmt.Println(fi12)
        _ = fi12
        _ = v13
    }
    switch ((("foo bar" + "a") + "") + "a") {
    case "foo bar":
        v5 -= v5
    case "hello":
        for fi14 := 0; fi14 < 20; fi14++ {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            _ = fi14
        }
    }
    fmt.Println(v4)
    delete(v2, v4)
    v15 := GenS1{F0: 944}
    v5 = v5
    v3 = (1e10 - float32((0.1 * -2.25)))
    switch "a" {
    case "":
        for fi16 := 0; fi16 < 13; fi16++ {
            fmt.Println(v3, v4)
            _ = len(v1)
            _ = fi16
        }
    default:
        v17 := GenS1{F0: 360}
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
