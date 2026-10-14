package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 int }
func genFn0(p0 uint16, p1 int, p2 uint16) uint32 {
    v5 := []int{}
    _ = len(v5)
    p1 = p1
    if len(v5) > 0 {
        v5[0] = (p1 & (^(p1)))
    }
    switch ("" + "hello") {
    case "":
        if false {
            v6 := GenS1{F0: -1}
            _ = v6
        }
    case "hello":
        v7 := GenS1{F0: 42}
        _ = v7
    case "foo bar":
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: 0}
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2147483648
}
func genFn1(p0 uint, p1 int16) string {
    for fi10 := 0; fi10 < 4; fi10++ {
        p1 = p1
        p1 = -1
        _ = fi10
    }
    var v11 uint8 = 254
    var v12 int16 = p1
    v13 := GenS2{F0: (717 < 413), F1: v12}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return "a"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 42
    var v3 uint8 = v2
    var v4 uint = 4294967295
    v14 := GenS3{F0: ""}
    if ((v3 >> ((v4) & 63)) == v2) {
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(v2, fi15, v4)
            v16 := []int{fi15, 71, (fi15 << ((607) & 63))}
            v3 = v3
            _ = fi15
            _ = v16
        }
    } else {
        var v17 string = "hello"
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        fmt.Println(fi18)
        switch ("" + ("foo bar" + ("" + "\n\t"))) {
        case "a":
            var v19 string = "foo bar"
            _ = v19
        }
        _ = fi18
    }
    v2 = v3
    switch 32767 {
    case 1:
        if false {
            var v20 uint16 = 0
            var v21 uint32 = genFn0(109, 42, v20)
            v4 = 255
            _ = v20
            _ = v21
        }
    case 127:
        if (18446744073709551615 > 18446744073709551615) {
            var v22 int8 = 42
            var v23 int8 = v22
            for fi24 := 0; fi24 < 3; fi24++ {
                v25 := GenS4{F0: genFn0(390, fi24, 315), F1: (fi24 % (((fi24 + fi24)) | 1))}
                _ = fi24
                _ = v25
            }
            _ = v22
            _ = v23
        } else {
            v26 := []int{128, 2147483647, 128}
            _ = v26
        }
    case 4294967295:
        for fi27 := 0; fi27 < 11; fi27++ {
            fmt.Println(v4, v3)
            _ = fi27
        }
    default:
        v3 += (^(v2))
    }
    var v28 int32 = 42
    var v29 int32 = 2147483647
    var v30 uint = v4
    var v31 float64 = 0.0
    var v32 float32 = (float32(v31) / 3.14159)
    fmt.Println(v29, v28)
    v28 -= -1
    v30 = (^(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
