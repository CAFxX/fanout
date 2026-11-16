package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int16 }

func main() {
    var v1 int8 = 81
    var v2 uint8 = 11
    var v3 int16 = 32767
    v4 := []int{48, -33, 42, 4294967295}
    if len(v4) > 0 {
        v4[0] = 32767
    }
    var v5 uint8 = v2
    v3 = 32767
    v4 = append(v4, -1)
    for fi6 := 0; fi6 < 4; fi6++ {
        if true {
            v3 = (v3 << ((4294967295) & 63))
        }
        _ = fi6
    }
    v1 -= 100
    switch -568 {
    case -950:
        var v7 bool = (!((1e10 == -2.25) && true))
        _ = v7
    }
    _ = len(v4)
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 9; fi9++ {
        if (463 == 2) {
            fmt.Println(v1, fi9, v3)
            v5 += (v2 >> ((631) & 63))
            fmt.Println(fi9, fi9, v2)
        } else {
            var v10 int64 = 42
            if len(v4) > 0 {
                v4[0] = (fi9 * (^(fi9)))
            }
            _ = v10
        }
        _ = len(v4)
        var v11 uint = 363
        _ = fi9
        _ = v11
    }
    var v12 int32 = -831
    if false {
        v1 += (^(v1))
        switch ("hello" + ("hello" + "")) {
        case "":
            fmt.Println(v5)
        case "a":
            for fi13 := 0; fi13 < 11; fi13++ {
                switch 525 {
                case -916:
                    v2 = (v5 + 255)
                }
                var v14 uint32 = 687
                v15 := GenS2{F0: 338, F1: (-(v3))}
                _ = fi13
                _ = v14
                _ = v15
            }
        case "foo bar":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
