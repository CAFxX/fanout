package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }

func main() {
    var v1 int = 9223372036854775807
    var v2 int = (v1 << ((702) & 63))
    var v3 uint16 = 2
    var v4 uint8 = 2
    var v5 int = (v1 + (v2 | 127))
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: v2, F1: 127}
        v8 := []int{v2, 0}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 string = "hello"
        for fi11 := 0; fi11 < 9; fi11++ {
            v10 = (("foo bar" + "a") + "foo bar")
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    v12 := GenS1{F0: (^(v1)), F1: -1}
    for fi13 := 0; fi13 < 7; fi13++ {
        v5 = (v5 << ((2) & 63))
        _ = fi13
    }
    switch v2 {
    case 66:
        if (false && (-100 <= 0)) {
            for fi14 := 0; fi14 < 7; fi14++ {
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 bool = true
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
            var v17 uint64 = 42
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 1; fi18++ {
            switch fi18 {
            case -665:
                fmt.Println(v3, v2)
            }
            v3 = v3
            if false {
                v19 := []int{}
                v2 += 128
                _ = v19
            } else {
                v5 = -620
                v3 = 255
            }
            _ = fi18
        }
    }
    var v20 uint16 = (v3 ^ (^(v3)))
    var v21 int8 = -105
    for fi22 := 0; fi22 < 19; fi22++ {
        v1 -= (v1 ^ v2)
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
