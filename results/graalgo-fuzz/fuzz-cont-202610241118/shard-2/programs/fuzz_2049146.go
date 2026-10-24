package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
func genFn0() int32 {
    var v5 int8 = 127
    switch "a" {
    case "a":
        v6 := []int{-853, 0, 184}
        _ = v6
    case "hello":
        if true {
            v7 := []int{}
            var v8 uint8 = 128
            _ = v7
            _ = v8
        } else {
            var v9 uint16 = 662
            _ = v9
        }
    case "\n\t":
        v5 = 0
    }
    _ = v5
    return 1
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint8 = 2
    var v3 bool = (!true)
    v4 := make(map[uint32]int)
    if v3 {
        v3 = (!v3)
        fmt.Println(v1)
        fmt.Println(v3, v3)
    }
    var v10 int64 = 42
    v1 -= (v1 % ((42) | 1))
    v10 += 0
    var v11 int64 = (v10 << ((693) & 63))
    fmt.Println(v11)
    var v12 int = -233
    v11 = -9223372036854775808
    switch "hello" {
    case "":
        var v13 int32 = 0
        _ = v13
    case "a":
        if (998 < 18446744073709551615) {
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 bool = true
                v15 = (("\n\t" + "") != (("a" + (("a" + "foo bar") + "hello")) + ""))
                if (true || v15) {
                    for fi16 := 0; fi16 < 2; fi16++ {
                        switch "hello" {
                        case "\n\t":
                            delete(v4, 2147483648)
                            v17 := GenS1{F0: 2}
                            var v18 int8 = int8(v2)
                            _ = v17
                            _ = v18
                        }
                        v2 += (v2 & uint8(v1))
                        v3 = (true && ((^(v2)) > 42))
                        _ = fi16
                    }
                    v19 := v4[98]
                    _ = len(v4)
                    _ = v19
                }
                _ = fi14
                _ = v15
            }
            fmt.Println(v1, v11, v2)
            var v20 int16 = -1
            _ = v20
        } else {
            fmt.Println(v10, v2, v12)
            var v21 int8 = 127
            if false {
                var v22 uint32 = 70
                _ = v22
            }
            _ = v21
        }
    }
    var v23 float64 = ((0.0 / 1.5) + (0.0 * ((1.5 * 0.0) - 0.1)))
    switch "foo bar" {
    case "":
        v24 := GenS2{F0: v1}
        _ = v24
    }
    v1 = (v1 << ((2) & 63))
    var v25 uint32 = 42
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 int8 = 127
        switch v12 {
        case 957:
            v28 := GenS3{F0: v27}
            _ = v28
        case -1:
            var v29 bool = (false && (!v3))
            _ = v29
        default:
            delete(v4, (^(v25)))
        }
        _ = fi26
        _ = v27
    }
    var v30 float64 = float64(float32(v23))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v30
    _ = v4
}
