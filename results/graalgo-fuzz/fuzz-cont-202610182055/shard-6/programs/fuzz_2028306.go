package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 931
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 5; fi5++ {
        switch (fi5 / ((fi5) | 1)) {
        case 1:
            var v6 float64 = 1e10
            _ = v6
        case 921:
            v7 := GenS1{F0: 32767}
            _ = v7
        default:
            switch "hello" {
            case "":
                var v8 int32 = -1
                _ = v8
            case "foo bar":
                fmt.Println(v2, fi5)
            case "\n\t":
                v9 := GenS2{F0: 0, F1: "\n\t"}
                _ = v9
            }
        }
        v2 = (v2 - 42)
        _ = fi5
    }
    if false {
        var v10 float32 = 1e10
        v10 = float32(1e10)
        _ = v10
    } else {
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 string = "a"
            v12 = v12
            _ = len(v1)
            _ = fi11
            _ = v12
        }
        v2 = v2
        v13 := GenS2{F0: 42, F1: "\n\t"}
        _ = v13
    }
    if (!(2147483647 >= -911)) {
        v3 = (v3 >> ((101) & 63))
    } else {
        v2 -= (-(v2))
    }
    v2 += (v2 + 42)
    fmt.Println(v2, v3)
    var v14 int8 = 100
    v15 := GenS1{F0: 183}
    for fi16 := 0; fi16 < 7; fi16++ {
        if (!false) {
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := GenS3{F0: v3}
                _ = fi17
                _ = v18
            }
        }
        _ = len(v1)
        v19 := v4[1]
        _ = fi16
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 uint = 2
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 float32 = 3.14159
        var v24 uint8 = 1
        _ = fi22
        _ = v23
        _ = v24
    }
    if (-9223372036854775808 <= 1) {
        v25 := v1[uint32(v3)]
        for fi26 := 0; fi26 < 1; fi26++ {
            v25 = (^(v25))
            v2 += (v2 << ((v3) & 63))
            _ = fi26
        }
        _ = v25
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        v3 += v3
        var v28 uint8 = 155
        v29 := GenS4{F0: "hello"}
        _ = fi27
        _ = v28
        _ = v29
    }
    v4[475] = false
    var v30 bool = (false || true)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
