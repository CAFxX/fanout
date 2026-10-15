package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 int8 = -30
    var v4 float64 = ((((float64((-2.25 - 0.1)) / 1e10) / float64(1e10)) * 3.14159) * 0.0)
    var v5 uint32 = 65535
    var v6 uint16 = 42
    if (-32768 < 0) {
        delete(v1, 42)
    }
    _ = len(v2)
    if (65535 == 18446744073709551615) {
        v6 += v6
        v7 := v2["foo bar"]
        _ = v7
    } else {
        v8 := GenS1{F0: 65535, F1: (2 != 1), F2: -1}
        switch "a" {
        case "\n\t":
            var v9 uint16 = v6
            _ = v9
        case "":
            var v10 int8 = -70
            _ = v10
        case "foo bar":
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 int = 32767
                _ = fi11
                _ = v12
            }
        }
        fmt.Println(v5)
        _ = v8
    }
    var v13 uint16 = v6
    var v14 float32 = -2.25
    v15 := GenS2{F0: (v13 % (((v6 | v6)) | 1)), F1: 128}
    var v16 string = "\n\t"
    var v17 uint32 = v5
    fmt.Println(v14, v4)
    var v18 string = v16
    v19 := GenS2{F0: 42, F1: 551}
    v5 -= (v5 & v5)
    v18 = ((v18 + "a") + ("hello" + "a"))
    if (1 != 346) {
        if (v16 >= v16) {
            v20 := GenS3{F0: 107, F1: 841}
            v3 -= v3
            _ = v20
        } else {
            var v21 bool = (true && ((^(v5)) >= (v5 % (((v17 ^ v5)) | 1))))
            _ = v21
        }
    }
    v22 := []int{int(v3), -331, 62, -9223372036854775808}
    v23 := GenS3{F0: v3, F1: 757}
    for fi24 := 0; fi24 < 20; fi24++ {
        v25 := GenS4{F0: true, F1: fi24}
        var v26 uint32 = (v17 << ((255) & 63))
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
