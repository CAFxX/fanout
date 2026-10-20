package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS2 struct { F0 int16; F1 int16; F2 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint32) uint16 {
    fmt.Println(p0)
    switch "hello" {
    case "\n\t":
        v6 := []int{9223372036854775807, 2147483647}
        _ = v6
    default:
        switch (("hello" + "foo bar") + "hello") {
        case "hello":
            v7 := GenS1{F0: (p0 * p0), F1: -94, F2: p0}
            _ = v7
        }
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 1
    var v3 int16 = 42
    var v4 uint8 = 255
    v5 := []int{566, 32768}
    for fi8 := 0; fi8 < 1; fi8++ {
        v4 += (^(v4))
        _ = fi8
    }
    v3 += 42
    fmt.Println(v3, v2, v2)
    var v9 int64 = 42
    for fi10 := 0; fi10 < 20; fi10++ {
        v2 = (v2 + 4294967295)
        if false {
            v11 := GenS1{F0: 1, F1: 95, F2: 455}
            v12 := v1["hello"]
            fmt.Println(v12, fi10, v3)
            _ = v11
            _ = v12
        }
        var v13 uint = v2
        _ = fi10
        _ = v13
    }
    v1[("hello" + ("a" + ""))] = false
    v9 -= -9223372036854775808
    for fi14 := 0; fi14 < 12; fi14++ {
        _ = len(v5)
        var v15 uint16 = 255
        _ = fi14
        _ = v15
    }
    switch 32768 {
    case 9223372036854775807:
        var v16 bool = true
        _ = v16
    case -314:
        var v17 uint16 = 65535
        _ = v17
    }
    v2 -= 55
    delete(v1, ("foo bar" + "foo bar"))
    _ = len(v1)
    switch 0 {
    case 487:
        if false {
            fmt.Println(v4, v3, v9)
            if false {
                var v18 int8 = 42
                _ = v18
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    case -9223372036854775808:
        for fi19 := 0; fi19 < 14; fi19++ {
            if true {
                fmt.Println(fi19, v9)
            } else {
                v4 = v4
                var v20 string = ""
                _ = v20
            }
            _ = fi19
        }
    default:
        v21 := GenS2{F0: 512, F1: v3, F2: 255}
        _ = v21
    }
    fmt.Println(v2, v2)
    v4 += (^(v4))
    v22 := GenS3{F0: (0 > int8(v9))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
