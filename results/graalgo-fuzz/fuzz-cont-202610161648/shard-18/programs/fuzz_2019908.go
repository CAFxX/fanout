package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int16 = -32768
    var v2 int64 = 0
    var v3 bool = false
    var v4 uint64 = 0
    var v5 uint8 = 128
    v6 := GenS1{F0: 0}
    var v7 int16 = v1
    v8 := []int{128, -9223372036854775808, -1}
    if (!false) {
        for fi9 := 0; fi9 < 5; fi9++ {
            if (v3 && true) {
                v8 = append(v8, (fi9 - (fi9 * (fi9 ^ (fi9 >> ((736) & 63))))))
            } else {
                var v10 int64 = int64(v4)
                fmt.Println(v5)
                _ = v10
            }
            _ = fi9
        }
        v11 := GenS1{F0: 2147483647}
        v7 -= (v7 & (v7 / ((v1) | 1)))
        _ = v11
    } else {
        for fi12 := 0; fi12 < 9; fi12++ {
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := GenS1{F0: 2147483647}
                var v15 int16 = 0
                v2 = v2
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
        var v16 int16 = v1
        _ = v16
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if len(v8) > 0 {
        v8[0] = 42
    }
    switch "a" {
    case "foo bar":
        var v17 uint16 = 32767
        _ = v17
    case "a":
        switch ("foo bar" + "hello") {
        case "foo bar":
            v7 = (v1 ^ v1)
        case "a":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        default:
            if (v3 && true) {
                _ = len(v8)
                for fi18 := 0; fi18 < 8; fi18++ {
                    v8 = append(v8, -1)
                    fmt.Println(v7)
                    _ = fi18
                }
                var v19 uint32 = 2
                _ = v19
            }
        }
    }
    if (!(int32(v1) == 1)) {
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := GenS1{F0: -2147483648}
            _ = fi20
            _ = v21
        }
        v22 := GenS1{F0: 817}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 5; fi23++ {
            v2 = v2
            for fi24 := 0; fi24 < 10; fi24++ {
                fmt.Println(fi23, fi24, v3)
                var v25 int8 = -1
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
