package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []string{}
    var v2 bool = ((((3.14159 - float64(3.14159)) >= 3.14159) && true) && ((243 >= 2147483648) || true))
    var v3 int = -2147483648
    var v4 int32 = int32(v3)
    v2 = (205 == 0)
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: (v2 || false)}
        _ = len(v1)
        var v7 int64 = 9223372036854775807
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 bool = (v2 || v2)
    if (true && v2) {
        _ = len(v1)
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 bool = (v2 && v8)
        v11 := GenS1{F0: v10}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v4 = v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    var v12 int32 = (v4 | v4)
    if v2 {
        var v13 uint8 = 128
        fmt.Println(v2, v2)
        if v2 {
            var v14 int = (v3 << ((477) & 63))
            var v15 bool = (!true)
            _ = v14
            _ = v15
        }
        _ = v13
    } else {
        for fi16 := 0; fi16 < 17; fi16++ {
            if (!true) {
                v8 = (-548 <= 42)
            } else {
                fmt.Println(fi16)
            }
            var v17 int = (-(fi16))
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v2 = false
        _ = fi18
    }
    var v19 int8 = -107
    var v20 int = (v3 * v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
