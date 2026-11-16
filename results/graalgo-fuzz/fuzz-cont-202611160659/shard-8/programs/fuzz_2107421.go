package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint = 42
    var v2 int8 = 42
    var v3 uint = (v1 - 42)
    v4 := []int{216}
    v5 := GenS1{F0: 0}
    v4 = append(v4, -1)
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint = (^(v1))
        var v8 uint16 = 665
        _ = fi6
        _ = v7
        _ = v8
    }
    if true {
        var v9 uint = v3
        _ = v9
    } else {
        var v10 int8 = int8(v3)
        var v11 string = ""
        var v12 uint32 = 0
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint16 = 65535
        switch 0 {
        case 898:
            v15 := GenS2{F0: (float32((float64(-2.25) / 1.5)) == 1e10)}
            _ = v15
        }
        var v16 float32 = 1e10
        _ = fi13
        _ = v14
        _ = v16
    }
    v1 += v3
    v1 += (v1 ^ (v1 % ((255) | 1)))
    v17 := GenS3{F0: -180}
    v3 = 255
    var v18 bool = true
    if v18 {
        if ((v18 && (904 != 0)) && ((886 > 2147483647) && (true || (!(2 != 4294967295))))) {
            var v19 uint16 = 682
            _ = len(v4)
            _ = v19
        } else {
            v2 = v2
            if (734 <= 563) {
                v3 -= v3
                var v20 uint64 = 0
                v4 = append(v4, 236)
                _ = v20
            }
        }
        v21 := GenS3{F0: 0}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
