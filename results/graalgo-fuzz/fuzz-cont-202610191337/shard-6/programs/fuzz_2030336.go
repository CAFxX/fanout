package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint16, p1 uint16, p2 string) string {
    v4 := GenS1{F0: (p1 >> ((1) & 63))}
    v5 := []int{9223372036854775807, -879, 42}
    switch 42 {
    case 127:
        p2 = (p2 + "hello")
    case -708:
        p1 += 237
    case -1:
        var v6 int8 = 127
        _ = v6
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return ("\n\t" + "a")
}

func main() {
    var v1 float64 = ((0.0 * (0.0 - 1.5)) - (-2.25 * (3.14159 / 0.1)))
    v2 := []string{("foo bar" + "a")}
    var v3 int8 = 83
    if (4294967295 != 1) {
        var v7 uint16 = 65535
        _ = v7
    } else {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint16 = 32767
            v9 += 32767
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 uint8 = 255
            _ = fi10
            _ = v11
        }
        var v12 uint8 = 12
        _ = v12
    }
    v1 = (v1 / v1)
    fmt.Println(v3, v3)
    for fi13 := 0; fi13 < 11; fi13++ {
        fmt.Println(v3, v1)
        _ = fi13
    }
    v14 := GenS1{F0: 32767}
    var v15 int16 = 393
    fmt.Println(v15, v15, v1)
    var v16 float64 = v1
    v15 = v15
    fmt.Println(v3, v15, v3)
    switch -173 {
    case -858:
        for fi17 := 0; fi17 < 14; fi17++ {
            if len(v2) > 0 {
                v2[0] = ("foo bar" + "")
            }
            var v18 uint32 = 344
            _ = fi17
            _ = v18
        }
    case -488:
        v19 := GenS1{F0: 473}
        _ = v19
    default:
        v15 -= (v15 << ((18446744073709551615) & 63))
    }
    if false {
        v20 := GenS1{F0: 95}
        v21 := GenS2{F0: 0, F1: 2, F2: 4294967295}
        switch "\n\t" {
        case "foo bar":
            for fi22 := 0; fi22 < 13; fi22++ {
                v1 += 0.1
                _ = fi22
            }
        case "hello":
            v23 := GenS3{F0: 255}
            _ = v23
        }
        _ = v20
        _ = v21
    } else {
        var v24 int64 = 1
        if len(v2) > 0 {
            v2[0] = ""
        }
        var v25 uint64 = 1
        _ = v24
        _ = v25
    }
    v15 = v15
    switch ("hello" + "foo bar") {
    case "hello":
        if (float64(3.14159) != v16) {
            v26 := GenS4{F0: (v3 % ((42) | 1))}
            var v27 string = ("foo bar" + "hello")
            _ = v26
            _ = v27
        }
    case "a":
        for fi28 := 0; fi28 < 12; fi28++ {
            fmt.Println(v1)
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            _ = fi28
        }
    default:
        v15 = -942
    }
    var v29 float64 = v16
    var v30 uint64 = 0
    for fi31 := 0; fi31 < 5; fi31++ {
        fmt.Println(v1, v16)
        var v32 int8 = (v3 | (v3 - int8(v30)))
        var v33 float32 = 1.5
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
