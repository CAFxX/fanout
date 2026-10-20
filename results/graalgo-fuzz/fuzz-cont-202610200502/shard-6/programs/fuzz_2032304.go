package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint32; F2 bool }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 bool = true
    v2 := []int{-1, 32768, 127, 685}
    var v3 int = -9223372036854775808
    v4 := make(map[int]string)
    var v5 int = (-(v3))
    v3 = v5
    for fi6 := 0; fi6 < 8; fi6++ {
        if (4294967295 != 1) {
            var v7 uint64 = 2
            if (!false) {
                var v8 float32 = float32(0.1)
                v7 = v7
                _ = v8
            } else {
                var v9 uint64 = v7
                _ = v9
            }
            _ = v7
        }
        if (v1 || (v1 || true)) {
            if (127 < -30) {
                fmt.Println(fi6)
                fmt.Println(v3, v5)
            } else {
                v1 = (!(42 <= 2))
            }
        }
        v5 += 9223372036854775807
        _ = fi6
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v1 {
        v5 -= -530
        v10 := GenS1{F0: "a", F1: 55}
        v1 = v1
        _ = v10
    }
    delete(v4, -226)
    _ = len(v4)
    switch "a" {
    case "hello":
        v11 := v4[v5]
        _ = v11
    }
    v12 := GenS2{F0: 4294967295, F1: 733, F2: (-1 >= 42)}
    v13 := GenS3{F0: 0}
    _ = len(v4)
    switch "foo bar" {
    case "":
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 int8 = 7
            for fi16 := 0; fi16 < 16; fi16++ {
                v15 += (v15 << ((0) & 63))
                _ = fi16
            }
            fmt.Println(v15)
            _ = fi14
            _ = v15
        }
    case "\n\t":
        var v17 int = 0
        _ = v17
    case "foo bar":
        fmt.Println(v1, v5)
    }
    var v18 bool = false
    var v19 uint64 = 740
    if ("" <= ("" + ((("foo bar" + "foo bar") + "hello") + "foo bar"))) {
        for fi20 := 0; fi20 < 5; fi20++ {
            v19 -= 0
            v2 = append(v2, fi20)
            _ = fi20
        }
        v5 = 0
    } else {
        v21 := GenS1{F0: "", F1: 1}
        v5 = (v5 << ((255) & 63))
        _ = v21
    }
    var v22 int = (v5 >> ((336) & 63))
    v3 -= -2147483648
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
