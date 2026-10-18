package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int8, p1 uint, p2 uint32) uint8 {
    v5 := GenS1{F0: false}
    p0 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 23
}

func main() {
    var v1 uint64 = 643
    var v2 int64 = -815
    var v3 uint64 = (v1 - v1)
    var v4 string = "foo bar"
    v6 := GenS1{F0: ("foo bar" < (v4 + (v4 + v4)))}
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{fi7, fi7}
        switch "\n\t" {
        case "a":
            v1 += (v1 << ((42) & 63))
        case "\n\t":
            var v9 int32 = 0
            _ = v9
        case "":
            v2 = (v2 / (((v2 + v2)) | 1))
        }
        _ = fi7
        _ = v8
    }
    var v10 uint = 1
    fmt.Println(v2, v4, v3)
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := []int{}
        _ = fi11
        _ = v12
    }
    v2 = 580
    var v13 float64 = (0.1 / 1e10)
    for fi14 := 0; fi14 < 20; fi14++ {
        switch v4 {
        case "foo bar":
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 float32 = 0.1
                var v17 int32 = 2147483647
                v18 := GenS1{F0: false}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        case "a":
            v19 := []int{fi14}
            _ = v19
        case "\n\t":
            v20 := GenS1{F0: (!false)}
            _ = v20
        default:
            var v21 uint16 = 406
            _ = v21
        }
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 uint8 = 0
            _ = fi22
            _ = v23
        }
        v3 -= (v3 + (v1 ^ v1))
        _ = fi14
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        if (!(44 != 0)) {
            v25 := []int{17}
            v3 -= v1
            _ = v25
        } else {
            v3 = (v3 | (^(v1)))
            v26 := GenS1{F0: false}
            v27 := GenS1{F0: (v3 <= 805)}
            _ = v26
            _ = v27
        }
        var v28 int64 = (v2 << (((v10 - (v10 ^ (^(v10))))) & 63))
        v29 := GenS1{F0: (!(fi24 == fi24))}
        _ = fi24
        _ = v28
        _ = v29
    }
    var v30 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v6
}
