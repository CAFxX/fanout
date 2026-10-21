package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int; F2 bool }

func main() {
    var v1 uint32 = 65535
    var v2 float64 = float64((float32(0.0) * 0.1))
    v3 := make(map[string]string)
    var v4 uint = 4294967295
    var v5 int32 = 0
    v6 := make(map[string]bool)
    v7 := GenS1{F0: 18446744073709551615}
    fmt.Println(v2, v1)
    _ = len(v3)
    if true {
        v5 = 2147483647
        delete(v3, "")
        v5 += (-(v5))
    }
    if true {
        if false {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 int64 = 105
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := GenS2{F0: v4, F1: fi8, F2: (v4 > v4)}
                    _ = fi10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
        } else {
            switch "\n\t" {
            case "a":
                switch ("a" + (("a" + "foo bar") + "hello")) {
                case "hello":
                    v12 := []int{4294967295, -106, -56}
                    _ = v12
                default:
                    var v13 int64 = 9223372036854775807
                    _ = v13
                }
            case "hello":
                v5 -= 336
            }
            var v14 int16 = 867
            fmt.Println(v5)
            _ = v14
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            if false {
                var v16 uint64 = uint64(v1)
                for fi17 := 0; fi17 < 11; fi17++ {
                    fmt.Println(v4, fi17, v16)
                    _ = fi17
                }
                _ = v16
            }
            _ = fi15
        }
    }
    v18 := GenS1{F0: 1}
    v5 = -2147483648
    switch 32768 {
    case 2147483647:
        delete(v3, "a")
    }
    v5 -= (v5 | (v5 >> ((255) & 63)))
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 int64 = 0
        v4 += v4
        _ = fi19
        _ = v20
    }
    var v21 float32 = 0.1
    fmt.Println(v21, v5)
    fmt.Println(v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
