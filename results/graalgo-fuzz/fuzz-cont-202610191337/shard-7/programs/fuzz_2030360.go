package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int16 = -32768
    var v2 uint64 = 0
    var v3 uint = 2
    v4 := make(map[int]string)
    if (false && (false && true)) {
        v1 += 1000
        v1 += v1
        fmt.Println(v3, v1)
    } else {
        fmt.Println(v3, v3)
        for fi5 := 0; fi5 < 6; fi5++ {
            v4[fi5] = ""
            fmt.Println(v1, v2, v1)
            var v6 uint64 = v2
            _ = fi5
            _ = v6
        }
    }
    v3 = v3
    v7 := GenS1{F0: 4294967295}
    fmt.Println(v3, v1, v3)
    v2 += (v2 & (v2 << (((v3 << (((v3 ^ (v3 - (v3 - 18446744073709551615)))) & 63))) & 63)))
    v8 := GenS1{F0: (v3 | 65535)}
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := v4[(^(fi9))]
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int8 = 100
            var v13 int = fi9
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    v3 -= 335
    fmt.Println(v1, v3)
    var v14 bool = (!(!false))
    v2 -= (v2 | 285)
    var v15 uint8 = 0
    var v16 uint64 = (v2 ^ (v2 | 4294967295))
    v2 += 464
    v14 = (9223372036854775807 == -1)
    var v17 uint16 = 65535
    v14 = ((v14 && (v17 != (v17 % (((v17 / (((v17 % ((uint16(v3)) | 1))) | 1))) | 1)))) && v14)
    var v18 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
