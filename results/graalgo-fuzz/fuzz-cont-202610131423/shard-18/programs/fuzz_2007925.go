package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
type GenS2 struct { F0 int32; F1 uint32; F2 uint64 }

func main() {
    var v1 uint64 = 0
    var v2 float64 = 0.0
    var v3 uint64 = (v1 >> ((777) & 63))
    v1 -= 852
    for fi4 := 0; fi4 < 14; fi4++ {
        v3 = (^(v3))
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 uint8 = 0
            if (!true) {
                v2 += v2
            } else {
                fmt.Println(v6, v6, v3)
                v7 := GenS1{F0: 2, F1: 65535, F2: fi4}
                if false {
                    var v8 string = ("\n\t" + ("hello" + "foo bar"))
                    if (2147483648 < 1) {
                        var v9 uint16 = 883
                        fmt.Println(v1, v8, v9)
                        _ = v9
                    }
                    v10 := []int{}
                    v3 += v1
                    _ = v10
                    _ = v8
                }
                _ = v7
            }
            v3 += v1
            _ = fi5
            _ = v6
        }
        var v11 uint32 = 2147483647
        _ = fi4
        _ = v11
    }
    v12 := GenS1{F0: 239, F1: 2, F2: 42}
    var v13 int64 = 980
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := GenS2{F0: -2147483648, F1: 42, F2: v3}
        if false {
            v2 = v2
            v2 = v2
        }
        _ = fi14
        _ = v15
    }
    v16 := []int{-2147483648}
    if (-32768 > 32767) {
        var v17 uint8 = 0
        _ = v17
    }
    fmt.Println(v13)
    v16 = append(v16, 128)
    v13 -= v13
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    if len(v16) > 0 {
        v16[0] = 42
    }
    var v18 int64 = 118
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}
