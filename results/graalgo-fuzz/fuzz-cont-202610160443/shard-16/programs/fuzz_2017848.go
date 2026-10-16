package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{32768, 500, 128, -2147483648}
    v3 := []int{32768, 32767, 4294967295}
    v4 := GenS1{F0: -2147483648, F1: ("a" + "foo bar")}
    var v5 uint16 = 1
    v1[918] = 42
    v5 = v5
    v6 := GenS2{F0: 0, F1: 0}
    switch ((("hello" + ("foo bar" + "a")) + "foo bar") + "foo bar") {
    case "a":
        fmt.Println(v5)
    case "":
        var v7 uint8 = 255
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 11; fi9++ {
                v5 -= v5
                _ = fi9
            }
            fmt.Println(fi8, v5)
            _ = fi8
        }
    }
    v5 = 65535
    for fi10 := 0; fi10 < 12; fi10++ {
        v5 -= v5
        _ = fi10
    }
    v11 := v1[65535]
    for fi12 := 0; fi12 < 15; fi12++ {
        switch "a" {
        case "":
            v5 = (v5 >> ((235) & 63))
        case "a":
            v13 := GenS3{F0: 0}
            _ = v13
        case "foo bar":
            v5 = v5
        default:
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 int16 = 32767
                var v16 int8 = 100
                var v17 float32 = 1.5
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        }
        _ = fi12
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 int32 = -2147483648
        v20 := GenS3{F0: 0}
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        if false {
            var v22 uint16 = (v5 * v5)
            v3 = append(v3, (fi21 << ((18446744073709551615) & 63)))
            v23 := GenS3{F0: 1}
            _ = v22
            _ = v23
        }
        _ = fi21
    }
    switch 0 {
    case 32767:
        var v24 int32 = -1
        _ = v24
    case 4294967295:
        v5 = (v5 + 42)
    case 42:
        var v25 int64 = 9223372036854775807
        _ = v25
    default:
        v11 += (v11 ^ (^(v11)))
    }
    v11 = v11
    switch "a" {
    case "\n\t":
        var v26 uint16 = (v5 >> ((18446744073709551615) & 63))
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
