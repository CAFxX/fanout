package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int16 }
type GenS2 struct { F0 int64; F1 int }

func main() {
    var v1 int8 = 99
    var v2 string = ""
    var v3 string = ""
    v4 := []int{-653, 128, -2147483648}
    v5 := GenS1{F0: v2, F1: 255, F2: -141}
    var v6 uint32 = 2
    var v7 bool = (false || true)
    if (false && v7) {
        var v8 string = v2
        v4 = append(v4, 9223372036854775807)
        v1 = v1
        _ = v8
    } else {
        v9 := GenS2{F0: 42, F1: 32768}
        _ = v9
    }
    v4 = append(v4, 200)
    for fi10 := 0; fi10 < 4; fi10++ {
        v7 = v7
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint = 890
            _ = fi11
            _ = v12
        }
        var v13 uint16 = 1
        _ = fi10
        _ = v13
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (!v7) {
        var v14 uint16 = 235
        _ = v14
    } else {
        v4 = append(v4, int(v1))
        switch (("\n\t" + (v2 + v2)) + v3) {
        case "a":
            v7 = (v6 >= v6)
        }
        var v15 int32 = -2147483648
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(v1)
        v17 := GenS1{F0: "", F1: 2, F2: 1000}
        if ((v6 < v6) && ((v6 - v6) <= 2)) {
            v4 = append(v4, fi16)
        } else {
            v1 += 59
        }
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 string = "foo bar"
        v3 = v19
        _ = fi18
        _ = v19
    }
    fmt.Println(v6, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
