package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 42
    v3 := []uint32{}
    switch "" {
    case "":
        v2 += v2
    case "\n\t":
        v4 := GenS1{F0: -34, F1: "hello", F2: 1000}
        _ = v4
    default:
        v2 = v2
    }
    switch -2147483648 {
    case 565:
        var v5 uint8 = 2
        _ = v5
    case -2147483648:
        _ = len(v1)
    case 0:
        v2 -= (v2 ^ v2)
    }
    fmt.Println(v2)
    for fi6 := 0; fi6 < 18; fi6++ {
        fmt.Println(v2)
        _ = fi6
    }
    v7 := GenS1{F0: 127, F1: "\n\t", F2: 1000}
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS2{F0: 255, F1: 1}
        _ = fi8
        _ = v9
    }
    v10 := GenS3{F0: 42, F1: 65535}
    if false {
        v2 = (v2 + v2)
    } else {
        if (true || (32767 == 65535)) {
            fmt.Println(v2)
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            v2 += (^(v2))
            var v12 int = fi11
            switch 993 {
            case 128:
                var v13 uint64 = 18446744073709551615
                _ = v13
            case 2147483647:
                v3 = append(v3, 4294967295)
            case -1:
                fmt.Println(v2, v2, v12)
            }
            _ = fi11
            _ = v12
        }
        v2 = (v2 + v2)
    }
    delete(v1, "a")
    var v14 bool = false
    v2 += v2
    v15 := GenS4{F0: ("hello" + "hello")}
    v14 = v14
    v14 = v14
    var v16 string = ("\n\t" + "foo bar")
    v3 = append(v3, 42)
    if v14 {
        var v17 uint64 = 0
        _ = v17
    } else {
        v18 := GenS1{F0: 100, F1: ("hello" + (v16 + (((v16 + ((("\n\t" + v16) + ((v16 + ((v16 + (v16 + v16)) + v16)) + (v16 + "a"))) + (v16 + v16))) + v16) + (v16 + (v16 + "a"))))), F2: 1}
        v3 = append(v3, 2147483648)
        switch 0 {
        case 9223372036854775807:
            v19 := GenS2{F0: 1, F1: 128}
            _ = v19
        case -283:
            for fi20 := 0; fi20 < 3; fi20++ {
                delete(v1, (v16 + "a"))
                _ = fi20
            }
        case 481:
            var v21 int = 2147483647
            _ = v21
        }
        _ = v18
    }
    switch (v16 + v16) {
    case "\n\t":
        v22 := GenS4{F0: v16}
        _ = v22
    default:
        var v23 int16 = -1
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
}
