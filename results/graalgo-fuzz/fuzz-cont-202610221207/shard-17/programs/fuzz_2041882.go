package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int32 = -2147483648
    var v2 int64 = 9223372036854775807
    var v3 int8 = 12
    var v4 uint64 = 4294967295
    if false {
        switch 1 {
        case 24:
            var v5 uint32 = 343
            _ = v5
        case -2147483648:
            v6 := GenS1{F0: ("" >= "hello")}
            _ = v6
        }
        if false {
            v7 := []int{831, -697}
            _ = v7
        }
        if (false || (-517 > v1)) {
            if (true || (659 < v2)) {
                for fi8 := 0; fi8 < 11; fi8++ {
                    v1 += (v1 & 41)
                    var v9 float64 = 1.5
                    _ = fi8
                    _ = v9
                }
                var v10 uint8 = 0
                v4 = 1
                _ = v10
            }
            var v11 int8 = v3
            var v12 int8 = v3
            _ = v11
            _ = v12
        }
    } else {
        var v13 int8 = (v3 - (v3 | (v3 >> ((255) & 63))))
        switch "a" {
        case "hello":
            switch "a" {
            case "hello":
                fmt.Println(v3, v3, v3)
            }
        }
        switch "foo bar" {
        case "foo bar":
            v14 := []int{}
            _ = v14
        }
        _ = v13
    }
    v15 := GenS1{F0: true}
    var v16 float32 = (1e10 * 0.1)
    for fi17 := 0; fi17 < 19; fi17++ {
        fmt.Println(v4, v4, fi17)
        _ = fi17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        v2 -= v2
        _ = fi18
    }
    var v19 int16 = 628
    fmt.Println(v19, v16)
    var v20 int = 506
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
