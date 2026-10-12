package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint, p1 int, p2 int8) uint64 {
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int8 = 1
        _ = fi5
        _ = v6
    }
    p0 += p0
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: p0, F1: int32(p0), F2: 631}
        var v9 uint8 = 83
        v9 += v9
        _ = fi7
        _ = v8
        _ = v9
    }
    switch p1 {
    case 0:
        p1 += p1
    case 32768:
        if (false || true) {
            p2 = 26
        } else {
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 uint16 = 42
                for fi12 := 0; fi12 < 19; fi12++ {
                    v11 -= v11
                    p2 -= (p2 % (((p2 / ((int8(fi12)) | 1))) | 1))
                    fmt.Println(p2, fi12)
                    _ = fi12
                }
                _ = fi10
                _ = v11
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int16 = 1000
    var v2 uint16 = 0
    var v3 uint32 = 2147483647
    v4 := make(map[string]bool)
    fmt.Println(v1)
    if false {
        switch 32767 {
        case -1:
            v13 := []int{4294967295, 1, 128}
            _ = v13
        case 32767:
            var v14 uint8 = 255
            _ = v14
        }
    }
    switch -2147483648 {
    case 32768:
        v15 := v4[""]
        _ = v15
    case 42:
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 int64 = 0
            switch (("a" + "") + "a") {
            case "\n\t":
                var v18 float32 = 1.5
                _ = v18
            case "hello":
                fmt.Println(v1, fi16, v1)
            case "":
                if ((v2 / ((32767) | 1)) >= uint16(v3)) {
                    if (("\n\t" + "foo bar") < ("a" + "hello")) {
                        v19 := GenS2{F0: (v3 & v3), F1: v2}
                        var v20 uint32 = (v3 + (v3 * (v3 + (v3 >> ((255) & 63)))))
                        v21 := GenS1{F0: 4294967295, F1: 1, F2: (v3 ^ v3)}
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        fmt.Println(v3)
                    }
                }
            default:
                v1 = (v1 >> ((0) & 63))
            }
            _ = fi16
            _ = v17
        }
    }
    switch 470 {
    case 0:
        v22 := GenS2{F0: (v3 + uint32(v2)), F1: 826}
        _ = v22
    case 9223372036854775807:
        v2 = (v2 - (^(v2)))
    case -556:
        v23 := []int{9223372036854775807, 556, 733, 311}
        _ = v23
    default:
        var v24 int64 = 42
        _ = v24
    }
    if false {
        v25 := GenS3{F0: 244}
        var v26 uint = 0
        v27 := []int{128, 32768, 0, 42}
        _ = v25
        _ = v26
        _ = v27
    }
    switch 2147483647 {
    case 0:
        var v28 int16 = 0
        _ = v28
    }
    var v29 uint32 = v3
    v3 = (v29 + (v29 / ((634) | 1)))
    fmt.Println(v3)
    _ = len(v4)
    v29 += v3
    var v30 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
