package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }

func main() {
    var v1 float64 = -2.25
    var v2 float64 = 3.14159
    var v3 uint8 = 242
    v4 := []int{}
    for fi5 := 0; fi5 < 3; fi5++ {
        if len(v4) > 0 {
            v4[0] = fi5
        }
        var v6 uint8 = 42
        if len(v4) > 0 {
            v4[0] = 32767
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(v1, v1, v3)
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            if false {
                var v9 uint64 = 9
                _ = v9
            }
            fmt.Println(fi7, fi7, fi8)
            _ = fi8
        }
        _ = fi7
    }
    v1 += 1.5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        v3 = 68
    }
    v2 += (v2 - v1)
    var v10 uint = 255
    v11 := GenS1{F0: "\n\t"}
    if len(v4) > 0 {
        v4[0] = 32767
    }
    var v12 uint16 = 125
    v3 -= (v3 / (((v3 & v3)) | 1))
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            fmt.Println(fi14)
            _ = fi14
        }
        _ = fi13
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        if false {
            for fi16 := 0; fi16 < 17; fi16++ {
                v2 = v1
                _ = fi16
            }
            v1 = (1e10 * float64(1e10))
            for fi17 := 0; fi17 < 2; fi17++ {
                v18 := GenS2{F0: 2147483647}
                v19 := GenS1{F0: ("" + "a")}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
