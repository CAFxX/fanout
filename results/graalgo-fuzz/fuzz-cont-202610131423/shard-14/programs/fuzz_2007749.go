package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 bool }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []string{"a", "foo bar", "foo bar", "hello"}
    var v2 float64 = 0.0
    v3 := []uint32{}
    v4 := []uint32{996, 349, 65535, 194}
    var v5 bool = (-935 <= 513)
    var v6 string = "hello"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v7 := GenS1{F0: 0, F1: ((v6 + "hello") + (v6 + v6)), F2: (v5 || (false && (!(true && (1 < 42)))))}
    var v8 uint32 = 630
    v9 := GenS1{F0: 306, F1: "", F2: (v6 >= v6)}
    for fi10 := 0; fi10 < 14; fi10++ {
        if v5 {
            if len(v4) > 0 {
                v4[0] = (v8 << ((255) & 63))
            }
        }
        var v11 bool = (!true)
        _ = fi10
        _ = v11
    }
    var v12 uint8 = 2
    var v13 int32 = -272
    v13 = (v13 << ((42) & 63))
    v14 := GenS1{F0: 7, F1: v6, F2: (v8 <= (v8 >> ((18446744073709551615) & 63)))}
    fmt.Println(v8, v12)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v15 uint16 = 985
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v16 := GenS2{F0: uint32(v12)}
    v4 = append(v4, v8)
    v2 = v2
    var v17 int16 = 401
    var v18 uint = 482
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
