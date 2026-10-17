package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint32 }
func genFn0(p0 uint, p1 int) string {
    var v4 uint64 = 681
    var v5 float32 = (0.0 - 1e10)
    switch p1 {
    case -63:
        var v6 string = ""
        _ = v6
    default:
        v7 := GenS1{F0: 1, F1: ((!((v4 <= 0) || true)) || false), F2: 1}
        _ = v7
    }
    var v8 float64 = 0.1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v8
    return ""
}
func genFn1(p0 uint64) uint64 {
    v9 := GenS1{F0: 548, F1: true, F2: 693}
    var v10 uint16 = uint16(p0)
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint32 = 474
        switch "foo bar" {
        case "foo bar":
            fmt.Println(fi11)
        default:
            v13 := GenS2{F0: (p0 | p0)}
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v9
    return p0
}

func main() {
    var v1 string = "hello"
    var v2 int16 = 0
    v3 := []uint32{0, 231, 2147483647, 65535}
    if (127 < 212) {
        var v14 uint32 = 58
        _ = v14
    }
    switch -142 {
    case 128:
        var v15 uint64 = 940
        _ = v15
    case 32768:
        var v16 uint64 = 319
        _ = v16
    case -9223372036854775808:
        if len(v3) > 0 {
            v3[0] = 1
        }
    default:
        _ = len(v3)
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            if len(v3) > 0 {
                v3[0] = 405
            }
            switch v1 {
            case "":
                fmt.Println(v1, v1)
            case "foo bar":
                v3 = append(v3, 42)
            case "a":
                v1 = v1
                var v19 int = (fi18 & 2147483647)
                v1 = ((v1 + "a") + ((v1 + ((((v1 + v1) + (("\n\t" + v1) + v1)) + "a") + (v1 + (v1 + ("foo bar" + ""))))) + v1))
                _ = v19
            }
            v1 = "a"
            _ = fi18
        }
        if true {
            for fi20 := 0; fi20 < 4; fi20++ {
                fmt.Println(v1, v2)
                _ = fi20
            }
            v21 := GenS1{F0: 461, F1: (421 == 1), F2: 80}
            _ = v21
        }
        _ = fi17
    }
    var v22 uint = 2
    if true {
        var v23 int64 = 9223372036854775807
        v23 = (v23 | (v23 << ((v22) & 63)))
        fmt.Println(v22)
        _ = v23
    }
    v24 := GenS3{F0: 1, F1: v22, F2: 861}
    var v25 uint32 = 65535
    v3 = append(v3, 768)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
}
