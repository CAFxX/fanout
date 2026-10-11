package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint }
type GenS2 struct { F0 uint32 }
func genFn0() int8 {
    v6 := GenS1{F0: 1000, F1: 471, F2: 132}
    v7 := GenS1{F0: 0, F1: 4294967295, F2: 222}
    _ = v6
    _ = v7
    return -128
}

func main() {
    var v1 int16 = 168
    var v2 uint32 = 165
    var v3 uint16 = 32767
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]uint32)
    var v8 int64 = 9223372036854775807
    switch "\n\t" {
    case "\n\t":
        var v9 uint = 517
        _ = v9
    case "":
        switch -1 {
        case 0:
            v1 = v1
        case 26:
            if true {
                fmt.Println(v2, v2, v1)
                for fi10 := 0; fi10 < 20; fi10++ {
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v1 = -371
                        _ = fi11
                    }
                    v12 := v4[fi10]
                    _ = fi10
                    _ = v12
                }
            }
        case 4294967295:
            for fi13 := 0; fi13 < 19; fi13++ {
                switch "" {
                case "a":
                    v14 := v5[(v2 >> ((255) & 63))]
                    _ = v14
                default:
                    if true {
                        v3 += v3
                        v2 = v2
                        v15 := GenS2{F0: v2}
                        _ = v15
                    }
                }
                _ = fi13
            }
        default:
            switch 0 {
            case 2147483647:
                var v16 uint = 42
                _ = v16
            default:
                var v17 uint16 = (v3 << ((18446744073709551615) & 63))
                _ = v17
            }
        }
    case "hello":
        fmt.Println(v1, v3, v2)
    }
    delete(v5, 146)
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := v5[(v2 << ((1) & 63))]
        v3 -= v3
        _ = fi18
        _ = v19
    }
    var v20 int32 = -678
    v21 := []int{9223372036854775807}
    var v22 int = -2147483648
    v20 += 42
    var v23 uint = 0
    v1 += (v1 >> ((18446744073709551615) & 63))
    if ((false || ((v23 + uint(v2)) != (v23 ^ (v23 % ((471) | 1))))) && true) {
        switch -9223372036854775808 {
        case -2147483648:
            for fi24 := 0; fi24 < 14; fi24++ {
                var v25 int64 = (v8 << (((v23 << (((v23 - v23)) & 63))) & 63))
                fmt.Println(v22)
                _ = fi24
                _ = v25
            }
        }
    } else {
        v23 = v23
        fmt.Println(v22)
    }
    var v26 uint8 = 1
    var v27 uint16 = uint16(v2)
    for fi28 := 0; fi28 < 17; fi28++ {
        for fi29 := 0; fi29 < 10; fi29++ {
            _ = len(v21)
            var v30 string = ("foo bar" + "a")
            v31 := GenS1{F0: v1, F1: 713, F2: (v23 - (v23 / ((v23) | 1)))}
            var v32 uint32 = v2
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        switch int(v2) {
        case -182:
            var v33 uint = v23
            _ = v33
        case 2147483647:
            var v34 int32 = 42
            _ = v34
        }
        _ = fi28
    }
    if (("hello" + ((("a" + "foo bar") + "") + "\n\t")) != ("a" + "")) {
        v27 += 907
    }
    v8 = v8
    v26 = 6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
