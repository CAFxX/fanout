package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int8; F1 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 866
    var v3 int32 = -1
    var v4 uint32 = 557
    switch (("\n\t" + "a") + "foo bar") {
    case "":
        switch (v2 + v2) {
        case 9223372036854775807:
            fmt.Println(v4, v2, v2)
        }
    case "foo bar":
        v5 := GenS1{F0: -889, F1: 1}
        _ = v5
    default:
        if false {
            var v6 bool = false
            _ = v6
        }
    }
    switch v2 {
    case -9223372036854775808:
        fmt.Println(v3, v2, v4)
    case -2147483648:
        fmt.Println(v2, v2)
    case -1:
        for fi7 := 0; fi7 < 11; fi7++ {
            if true {
                for fi8 := 0; fi8 < 11; fi8++ {
                    v9 := GenS2{F0: 42, F1: 483}
                    _ = fi8
                    _ = v9
                }
                fmt.Println(v2, fi7, v2)
                for fi10 := 0; fi10 < 20; fi10++ {
                    if false {
                        v11 := []int{651, (fi10 * fi10), fi10}
                        v12 := GenS1{F0: 883, F1: 0}
                        v2 = v2
                        _ = v11
                        _ = v12
                    }
                    _ = fi10
                }
            }
            v3 = (v3 | -440)
            _ = fi7
        }
    default:
        fmt.Println(v3)
    }
    v4 = v4
    v2 = int(v3)
    v3 += v3
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := v1[65535]
        if false {
            if false {
                delete(v1, (v4 - 178))
                var v15 uint = 338
                fmt.Println(v14)
                _ = v15
            }
        } else {
            v14 = (v14 + v14)
        }
        v16 := GenS3{F0: 127, F1: v14}
        _ = fi13
        _ = v14
        _ = v16
    }
    v4 = 4294967295
    var v17 uint32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
