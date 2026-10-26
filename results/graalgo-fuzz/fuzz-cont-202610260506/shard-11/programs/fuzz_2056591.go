package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int64; F1 string; F2 int8 }
type GenS4 struct { F0 int; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint; F1 uint; F2 int16 }

func main() {
    var v1 uint64 = 2
    var v2 int = 492
    var v3 uint64 = (v1 >> ((1) & 63))
    v4 := make(map[string]uint32)
    var v5 int = (v2 + v2)
    if false {
        v2 = 32768
    } else {
        var v6 int8 = 0
        _ = v6
    }
    if (!(float32(0.1) <= 0.1)) {
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := []int{fi7, (v5 ^ v2)}
            if (!false) {
                v3 -= v1
            }
            _ = fi7
            _ = v8
        }
    } else {
        v9 := GenS1{F0: 0}
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 int32 = -1
            v3 += (v3 * (v3 ^ v1))
            var v12 float32 = float32(1.5)
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 20; fi13++ {
            switch fi13 {
            case 667:
                var v14 uint8 = 190
                _ = v14
            }
            _ = fi13
        }
        _ = v9
    }
    switch "" {
    case "foo bar":
        if true {
            for fi15 := 0; fi15 < 17; fi15++ {
                v1 = v3
                _ = fi15
            }
        }
    case "\n\t":
        v16 := []int{(v2 & v5), (^(v5)), v2}
        _ = v16
    case "a":
        v5 -= v2
    default:
        v17 := GenS1{F0: 46}
        _ = v17
    }
    v18 := GenS2{F0: 516, F1: v3, F2: 496}
    v19 := GenS3{F0: -9223372036854775808, F1: ("a" + ("a" + "")), F2: -128}
    switch "hello" {
    case "":
        v2 += v5
    case "hello":
        for fi20 := 0; fi20 < 6; fi20++ {
            v21 := v4["hello"]
            v22 := GenS3{F0: 9223372036854775807, F1: "hello", F2: -1}
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    var v23 bool = ((true || true) || true)
    v2 = v2
    fmt.Println(v3, v23)
    v24 := GenS1{F0: -2147483648}
    for fi25 := 0; fi25 < 15; fi25++ {
        if (!v23) {
            v26 := GenS4{F0: v5, F1: (^(v3)), F2: 52}
            if v23 {
                v5 += v2
                fmt.Println(v5, v1)
            }
            _ = v26
        }
        v3 = (v1 - (v1 % (((v1 * (^(v3)))) | 1)))
        _ = fi25
    }
    switch (v5 | (^(v5))) {
    case -718:
        v27 := GenS5{F0: 255, F1: 65535, F2: 3}
        _ = v27
    case 492:
        fmt.Println(v5, v1, v5)
    case 1000:
        var v28 int64 = -195
        _ = v28
    default:
        var v29 uint16 = 42
        _ = v29
    }
    var v30 uint = 42
    for fi31 := 0; fi31 < 8; fi31++ {
        fmt.Println(v2, v5)
        _ = fi31
    }
    v5 = v2
    switch (v5 * (-(v5))) {
    case -507:
        fmt.Println(v3, v30, v2)
    default:
        var v32 bool = true
        _ = v32
    }
    fmt.Println(v23, v23, v30)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
