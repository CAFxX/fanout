package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint8; F1 int8 }
func genFn0() uint16 {
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := GenS1{F0: 989, F1: 42, F2: 42}
        var v6 float64 = (float64(1e10) + 1e10)
        fmt.Println(fi4)
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 float64 = 1e10
    v8 := GenS2{F0: true}
    if (true || false) {
        var v9 int64 = 0
        var v10 uint64 = 18446744073709551615
        _ = v10
        _ = v9
    }
    _ = v7
    _ = v8
    return 200
}

func main() {
    var v1 uint16 = 42
    var v2 uint = 0
    v3 := []int{-9223372036854775808, -732, 32767}
    var v11 int8 = 100
    switch "\n\t" {
    case "foo bar":
        var v12 uint8 = 0
        _ = v12
    }
    v2 = (v2 << (((v2 & (v2 % (((^(v2))) | 1)))) & 63))
    v11 += v11
    v13 := GenS3{F0: true, F1: 719}
    v1 = (v1 | genFn0())
    if (1 >= 2147483647) {
        switch 2147483647 {
        case 112:
            v2 += (v2 & (v2 - (^(v2))))
        case 1:
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
        case -9223372036854775808:
            var v14 string = ("hello" + "a")
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 19; fi15++ {
            v11 = 1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi15
        }
    }
    switch "\n\t" {
    case "a":
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 uint16 = (v1 % ((65535) | 1))
                _ = fi17
                _ = v18
            }
            v11 = v11
            _ = fi16
        }
    case "":
        var v19 bool = (false || true)
        _ = v19
    }
    if true {
        fmt.Println(v2, v11)
        v1 = v1
        var v20 uint32 = 49
        _ = v20
    }
    fmt.Println(v11)
    v21 := GenS2{F0: false}
    var v22 bool = true
    _ = len(v3)
    switch ("a" + ("\n\t" + "foo bar")) {
    case "hello":
        switch (("a" + ("a" + "foo bar")) + "\n\t") {
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            v23 := GenS1{F0: -2147483648, F1: -277, F2: 42}
            _ = v23
        }
    case "foo bar":
        var v24 uint32 = uint32(v1)
        _ = v24
    case "a":
        var v25 float64 = (3.14159 + (0.1 - 0.0))
        _ = v25
    default:
        for fi26 := 0; fi26 < 6; fi26++ {
            v27 := GenS4{F0: 1, F1: v11}
            var v28 uint64 = 0
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    var v29 int32 = 2147483647
    for fi30 := 0; fi30 < 2; fi30++ {
        v1 = v1
        switch (fi30 - (^(fi30))) {
        case 9223372036854775807:
            v1 += (v1 ^ 648)
        default:
            v1 = v1
        }
        _ = fi30
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
}
