package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int32; F1 uint64; F2 int32 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 int8 = 42
    var v2 bool = false
    var v3 int16 = 1000
    v4 := []int{127, 128}
    fmt.Println(v2, v2)
    for fi5 := 0; fi5 < 13; fi5++ {
        v2 = ((3.14159 - 1.5) != float32(3.14159))
        _ = fi5
    }
    var v6 float64 = ((1e10 + 1.5) / 0.1)
    v3 -= (v3 << ((42) & 63))
    var v7 bool = (v3 <= (v3 >> ((1) & 63)))
    v4 = append(v4, 128)
    if len(v4) > 0 {
        v4[0] = 1
    }
    v8 := GenS1{F0: "\n\t", F1: 255}
    v9 := GenS2{F0: 405, F1: v1}
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v3, v6)
        _ = fi10
    }
    v11 := GenS3{F0: 0, F1: 123, F2: 2147483647}
    v12 := GenS4{F0: 42}
    v13 := GenS3{F0: 1, F1: 2, F2: -2147483648}
    for fi14 := 0; fi14 < 19; fi14++ {
        switch 2147483647 {
        case 32767:
            if (("foo bar" + ("" + "\n\t")) > (((("\n\t" + "\n\t") + "\n\t") + "foo bar") + ("\n\t" + "foo bar"))) {
                if ((!true) && (!(!(int(v1) != fi14)))) {
                    v1 = (^(v1))
                    v2 = v2
                    fmt.Println(v6, v1, v1)
                }
            } else {
                v7 = true
            }
        }
        v3 = (v3 | v3)
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
