package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 string; F1 uint; F2 int16 }
type GenS5 struct { F0 int32; F1 uint64; F2 uint16 }
func genFn0(p0 int32, p1 int8) uint16 {
    var v4 uint = 18446744073709551615
    v5 := GenS1{F0: 2147483648}
    switch "hello" {
    case "":
        v4 = v4
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 679
}

func main() {
    var v1 uint16 = 255
    var v2 string = ("a" + ((("hello" + "") + "") + "\n\t"))
    var v3 int16 = 0
    v3 += v3
    if (1000 <= v3) {
        v1 -= 65535
    }
    v6 := GenS1{F0: 1}
    var v7 bool = false
    var v8 string = (v2 + v2)
    switch 128 {
    case 9223372036854775807:
        v9 := GenS1{F0: 862}
        _ = v9
    case 32768:
        var v10 int = 32767
        _ = v10
    }
    if (!(!((-2.25 - float64(1.5)) < (1e10 - 0.1)))) {
        v11 := GenS1{F0: 2147483647}
        fmt.Println(v1)
        v7 = v7
        _ = v11
    } else {
        v3 -= (v3 + (v3 + (v3 << ((18446744073709551615) & 63))))
        v12 := GenS1{F0: 2147483648}
        _ = v12
    }
    var v13 uint32 = 2147483648
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := []int{}
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        if true {
            var v16 uint64 = 402
            _ = v16
        } else {
            switch "foo bar" {
            case "\n\t":
                v17 := GenS2{F0: -1}
                _ = v17
            default:
                v2 = ("" + ("a" + v2))
            }
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        }
        _ = fi14
        _ = v15
    }
    v3 = 42
    switch 296 {
    case 4294967295:
        if (!(!false)) {
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 int8 = 27
                _ = fi18
                _ = v19
            }
            v3 -= (v3 >> ((149) & 63))
            var v20 uint16 = genFn0(-155, 0)
            _ = v20
        }
    }
    v21 := GenS3{F0: (("" + (("foo bar" + "\n\t") + v8)) + "\n\t"), F1: 1}
    for fi22 := 0; fi22 < 13; fi22++ {
        v1 = genFn0(1, -128)
        v1 += (^(v1))
        v23 := GenS4{F0: (v2 + ""), F1: 4294967295, F2: v3}
        _ = fi22
        _ = v23
    }
    v13 += (v13 % (((^(v13))) | 1))
    v2 = v2
    switch int(v3) {
    case -9223372036854775808:
        var v24 int16 = v3
        _ = v24
    case -2147483648:
        for fi25 := 0; fi25 < 6; fi25++ {
            v26 := GenS5{F0: 0, F1: 0, F2: (v1 - 1)}
            var v27 uint32 = 534
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    var v28 int16 = (v3 / ((v3) | 1))
    v29 := GenS3{F0: "a", F1: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
