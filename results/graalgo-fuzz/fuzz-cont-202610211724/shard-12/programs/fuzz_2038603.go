package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int; F2 uint8 }
type GenS3 struct { F0 int32; F1 int16 }

func main() {
    v1 := []uint32{876}
    v2 := make(map[string]string)
    var v3 uint32 = 65535
    var v4 int16 = -838
    if (42 == 512) {
        if false {
            v5 := GenS1{F0: 789}
            _ = v5
        }
    }
    v6 := GenS2{F0: ("foo bar" + ((("foo bar" + "hello") + ("" + "")) + "\n\t")), F1: 32768, F2: 70}
    if (763 >= 80) {
        v4 -= 1000
    } else {
        v3 -= 0
        if false {
            var v7 int8 = -1
            if true {
                v7 = (^(v7))
            }
            if (-1 == 4294967295) {
                var v8 uint = 54
                v3 = v3
                var v9 uint16 = 32767
                _ = v8
                _ = v9
            } else {
                if false {
                    for fi10 := 0; fi10 < 2; fi10++ {
                        v7 = (v7 | 46)
                        _ = fi10
                    }
                    var v11 int64 = -9223372036854775808
                    _ = len(v2)
                    _ = v11
                }
                fmt.Println(v7)
            }
            _ = v7
        } else {
            fmt.Println(v3)
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            _ = len(v1)
            _ = fi12
        }
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS3{F0: 0, F1: v4}
        fmt.Println(v3)
        v15 := GenS3{F0: 42, F1: v4}
        _ = fi13
        _ = v14
        _ = v15
    }
    v4 += (v4 >> ((782) & 63))
    if true {
        v2[("a" + (("" + "hello") + ""))] = (("a" + "a") + ("a" + "\n\t"))
        switch 0 {
        case 9223372036854775807:
            var v16 int16 = v4
            _ = v16
        case -1:
            v17 := GenS3{F0: -1, F1: v4}
            _ = v17
        }
        v4 = 1000
    }
    fmt.Println(v3, v4)
    v3 = 0
    var v18 int64 = -1
    v4 -= (v4 / (((-(v4))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
