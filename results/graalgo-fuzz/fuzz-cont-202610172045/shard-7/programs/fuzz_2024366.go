package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint }

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 153
    var v3 int8 = 100
    v4 := make(map[string]uint32)
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 int = (-(fi5))
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: v3, F1: true, F2: 938}
    for fi8 := 0; fi8 < 20; fi8++ {
        if true {
            v1 -= 2
        }
        _ = fi8
    }
    switch 0 {
    case -1:
        for fi9 := 0; fi9 < 6; fi9++ {
            if true {
                var v10 string = ((("foo bar" + "hello") + ((("foo bar" + "foo bar") + "\n\t") + "\n\t")) + "")
                var v11 int8 = v3
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := []int{int(v2)}
                if true {
                    fmt.Println(v2, fi9)
                } else {
                    var v14 bool = true
                    _ = v14
                }
                var v15 int = (fi9 | 961)
                fmt.Println(fi9, v1)
                var v16 uint64 = (v1 ^ v1)
                _ = fi12
                _ = v13
                _ = v15
                _ = v16
            }
            _ = fi9
        }
    case -180:
        v2 += v2
    case 9223372036854775807:
        v2 -= 32767
    default:
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            var v17 int = -1
            _ = v17
        case "a":
            var v18 int16 = 32767
            _ = v18
        case "":
            var v19 uint16 = (^(v2))
            _ = v19
        }
    }
    var v20 uint16 = 42
    v1 -= 72
    var v21 uint = 1
    var v22 uint = 4294967295
    var v23 float32 = (3.14159 - float32(3.14159))
    fmt.Println(v20, v20)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v7
}
