package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int; F2 int16 }

func main() {
    var v1 int = 1
    v2 := []string{}
    var v3 float32 = 1e10
    var v4 int = (v1 * v1)
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := GenS1{F0: 0, F1: 1, F2: 18446744073709551615}
            v8 := GenS1{F0: 32767, F1: -964, F2: 18446744073709551615}
            if false {
                v2 = append(v2, ("\n\t" + (("" + "hello") + "a")))
                v3 = float32(-2.25)
            } else {
                var v9 uint32 = 89
                v4 = v1
                if len(v2) > 0 {
                    v2[0] = ""
                }
                _ = v9
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        v1 += v4
        fmt.Println(v1)
        _ = fi5
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v1 = (v4 >> ((4294967295) & 63))
        var v11 int32 = 42
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi10
        _ = v11
    }
    switch ("" + "hello") {
    case "":
        v3 = 0.1
    default:
        v4 = v4
    }
    v1 -= (v1 * v1)
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 string = ("\n\t" + "a")
            _ = fi13
            _ = v14
        }
        var v15 bool = true
        _ = fi12
        _ = v15
    }
    v16 := GenS2{F0: 255, F1: (^(v1)), F2: 0}
    v17 := GenS2{F0: 1, F1: v1, F2: 42}
    var v18 uint64 = 0
    var v19 float32 = 1e10
    var v20 int = v4
    var v21 uint = 255
    for fi22 := 0; fi22 < 18; fi22++ {
        v3 = 0.1
        fmt.Println(v3, v3)
        _ = fi22
    }
    var v23 float32 = ((3.14159 + 3.14159) * v19)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
}
