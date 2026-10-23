package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int16 = -890
    var v2 uint = 294
    v3 := make(map[string]string)
    if ((("a" + "a") <= ((("\n\t" + "\n\t") + "foo bar") + "a")) && false) {
        v2 += v2
    } else {
        var v4 uint8 = 0
        switch 9223372036854775807 {
        case 32767:
            v5 := GenS1{F0: 65535}
            _ = v5
        case 12:
            v6 := GenS1{F0: 65535}
            _ = v6
        }
        _ = v4
    }
    var v7 int = -1
    var v8 int16 = v1
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS2{F0: 65535}
        delete(v3, "hello")
        _ = fi9
        _ = v10
    }
    if true {
        v8 = 1000
    }
    if (!true) {
        v11 := GenS2{F0: 2}
        fmt.Println(v7, v1)
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(v1, v8, v2)
            var v13 int32 = 1
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    v14 := GenS3{F0: 65535}
    fmt.Println(v7, v7)
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int64 = 1
        switch "foo bar" {
        case "":
            v7 -= -352
        }
        v8 = (v1 + (v8 - (v8 >> (((v2 << ((v2) & 63))) & 63))))
        _ = fi15
        _ = v16
    }
    v17 := []int{v7, (v7 >> (((v2 ^ v2)) & 63)), v7}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
