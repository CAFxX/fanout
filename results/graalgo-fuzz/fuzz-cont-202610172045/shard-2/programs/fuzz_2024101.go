package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint = 65535
    var v2 int32 = 42
    v3 := []string{}
    var v4 int8 = 100
    fmt.Println(v4, v1)
    if len(v3) > 0 {
        v3[0] = ("hello" + "hello")
    }
    v1 = (v1 >> ((18446744073709551615) & 63))
    for fi5 := 0; fi5 < 9; fi5++ {
        switch (fi5 ^ (fi5 * fi5)) {
        case 128:
            var v6 uint = 4294967295
            _ = v6
        default:
            v4 = v4
        }
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int64 = -1
            var v9 uint16 = 2
            var v10 int16 = -32768
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi5
    }
    var v11 string = ""
    v12 := GenS1{F0: 4}
    fmt.Println(v4, v11, v1)
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS2{F0: 42}
        for fi15 := 0; fi15 < 8; fi15++ {
            if true {
                v11 = "a"
                v16 := GenS2{F0: 42}
                v11 = (v11 + ((v11 + v11) + (v11 + v11)))
                v17 := GenS2{F0: 32767}
                var v18 int16 = -1
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi15
        }
        fmt.Println(v1)
        _ = fi13
        _ = v14
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        var v19 uint32 = 2147483647
        _ = v19
    }
    if true {
        var v20 int8 = v4
        _ = v20
    }
    var v21 uint16 = 42
    v22 := GenS2{F0: v21}
    v2 += 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
