package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }

func main() {
    v1 := []string{""}
    var v2 bool = true
    var v3 uint16 = 1
    v4 := make(map[int]bool)
    v2 = false
    if (v2 || (2 >= 0)) {
        if v2 {
            var v5 uint8 = 128
            for fi6 := 0; fi6 < 7; fi6++ {
                v2 = (248 == 1)
                _ = fi6
            }
            _ = v5
        } else {
            if len(v1) > 0 {
                v1[0] = ("foo bar" + "")
            }
        }
        var v7 uint = 955
        _ = v7
    }
    var v8 uint = 65535
    var v9 int32 = 42
    v2 = (!(v2 && (!false)))
    for fi10 := 0; fi10 < 19; fi10++ {
        v8 += (^(v8))
        switch (fi10 ^ 550) {
        case -2147483648:
            v9 = int32(fi10)
        case 127:
            var v11 string = "a"
            _ = v11
        case 542:
            delete(v4, (fi10 / ((0) | 1)))
        }
        _ = fi10
    }
    fmt.Println(v9)
    v12 := GenS1{F0: 1, F1: v9}
    var v13 uint = (v8 | uint(v3))
    var v14 int16 = 32767
    var v15 int64 = -809
    fmt.Println(v13, v2)
    for fi16 := 0; fi16 < 13; fi16++ {
        v3 = 32767
        fmt.Println(v15)
        delete(v4, fi16)
        _ = fi16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
