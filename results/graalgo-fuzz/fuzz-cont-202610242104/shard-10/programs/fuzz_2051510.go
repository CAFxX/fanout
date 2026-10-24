package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint8 = 128
    var v2 int8 = 100
    var v3 int8 = v2
    v2 = (v2 >> ((18446744073709551615) & 63))
    if true {
        var v4 float64 = 0.0
        var v5 bool = false
        _ = v4
        _ = v5
    }
    v6 := GenS1{F0: 42}
    if (!((v2 == -128) && (false && false))) {
        var v7 int64 = 9223372036854775807
        _ = v7
    } else {
        v8 := []int{}
        _ = v8
    }
    v9 := GenS2{F0: 0, F1: -751}
    v1 = v1
    if false {
        fmt.Println(v3)
    }
    var v10 uint64 = 819
    for fi11 := 0; fi11 < 20; fi11++ {
        if (!(163 > 65535)) {
            switch -816 {
            case 127:
                var v12 int16 = 32767
                _ = v12
            case -443:
                v13 := GenS3{F0: "foo bar", F1: 65535}
                _ = v13
            case 371:
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := []int{(fi14 - 1), fi14}
                    _ = fi14
                    _ = v15
                }
            }
            v2 += 42
            v2 -= (v2 << ((2) & 63))
        } else {
            v16 := GenS4{F0: 70}
            var v17 string = ""
            var v18 bool = false
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
