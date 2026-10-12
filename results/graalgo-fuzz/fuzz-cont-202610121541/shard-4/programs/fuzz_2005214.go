package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint }

func main() {
    var v1 string = ""
    var v2 int16 = 109
    v3 := []int{4294967295, 0, -2147483648}
    v4 := make(map[string]string)
    var v5 int64 = -1
    var v6 float32 = -2.25
    v7 := GenS1{F0: 582}
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint8 = 127
            v11 := GenS1{F0: 65535}
            v6 = v6
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: 65535}
        _ = fi8
        _ = v12
    }
    var v13 uint16 = 255
    fmt.Println(v13, v1, v6)
    v13 = v13
    v5 = 0
    switch v1 {
    case "a":
        v1 = v1
    }
    switch -852 {
    case -780:
        v14 := GenS2{F0: 0, F1: 65535}
        _ = v14
    case 32768:
        if (180 > 0) {
            v15 := v4[(((v1 + ((v1 + v1) + v1)) + (v1 + v1)) + (v1 + (v1 + ("hello" + ("foo bar" + v1)))))]
            v16 := v4[("hello" + (v15 + v15))]
            _ = v15
            _ = v16
        } else {
            v6 = v6
            v4[v1] = v1
            v2 = 0
        }
    case 128:
        delete(v4, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
