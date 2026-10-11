package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int32 = -365
    var v2 int = 9223372036854775807
    var v3 bool = true
    v4 := []uint32{42, 4294967295, 307}
    var v5 float64 = 0.0
    if (!(!v3)) {
        v6 := GenS1{F0: ("\n\t" + "a")}
        var v7 float32 = 1.5
        _ = v6
        _ = v7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1)
    _ = len(v4)
    v5 = v5
    var v8 uint16 = 675
    fmt.Println(v5, v1)
    v9 := GenS1{F0: ""}
    if v3 {
        if v3 {
            v1 -= v1
        } else {
            v5 = (3.14159 + v5)
            _ = len(v4)
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 int16 = -1
                for fi12 := 0; fi12 < 13; fi12++ {
                    v8 = 32767
                    var v13 int32 = (v1 | 2147483647)
                    _ = fi12
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
        }
    } else {
        if v3 {
            fmt.Println(v1, v2, v1)
        }
    }
    switch "" {
    case "hello":
        var v14 float32 = 0.0
        _ = v14
    case "a":
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 int16 = -1
            _ = fi15
            _ = v16
        }
    }
    var v17 float64 = float64((-2.25 / float32(1.5)))
    if false {
        if false {
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int16 = -32768
        fmt.Println(v5, v8)
        _ = fi18
        _ = v19
    }
    v2 -= (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
