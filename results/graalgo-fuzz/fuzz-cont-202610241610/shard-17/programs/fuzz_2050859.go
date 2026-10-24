package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }

func main() {
    var v1 int16 = -870
    var v2 string = ("" + "foo bar")
    var v3 uint16 = 1
    var v4 uint8 = 0
    switch -9223372036854775808 {
    case 2147483647:
        switch ("foo bar" + v2) {
        case "\n\t":
            var v5 int64 = 9223372036854775807
            _ = v5
        case "":
            v6 := []int{9223372036854775807, 128, 42, -2147483648}
            _ = v6
        case "hello":
            var v7 uint32 = 111
            _ = v7
        default:
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 int32 = -63
                _ = fi8
                _ = v9
            }
        }
    }
    var v10 uint = 687
    if (-556 > -9223372036854775808) {
        v1 += v1
    }
    fmt.Println(v10)
    v1 = v1
    if (v10 > 91) {
        if false {
            v11 := GenS1{F0: (^(v10)), F1: true}
            var v12 int32 = -1
            v1 += (v1 + (v1 - (v1 | (v1 | (-(v1))))))
            _ = v11
            _ = v12
        } else {
            var v13 uint8 = v4
            _ = v13
        }
        v14 := []int{-2147483648, 792, 487}
        for fi15 := 0; fi15 < 2; fi15++ {
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 uint32 = 540
                v17 -= (^(v17))
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 2; fi18++ {
                v1 += (^(v1))
                _ = fi18
            }
            fmt.Println(fi15, v3)
            _ = fi15
        }
        _ = v14
    }
    v1 += v1
    v19 := GenS1{F0: v10, F1: false}
    var v20 uint = v10
    if true {
        v21 := []int{-30, 32767, 32767, 9223372036854775807}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
