package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int16 }
type GenS2 struct { F0 int; F1 int32; F2 int }
type GenS3 struct { F0 int; F1 int; F2 bool }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint, p2 uint8) int32 {
    v4 := []int{32768, 4294967295, 1, 128}
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int64 = -29
        switch (fi5 << ((219) & 63)) {
        case -2147483648:
            v7 := GenS1{F0: 976, F1: 0, F2: int16(v6)}
            _ = v7
        case 128:
            if (((!(float32(float64(float32(float64(0.1)))) <= ((1.5 / 1.0) / 1.0))) || (!((fi5 / ((fi5) | 1)) > -846))) && false) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(p0, v6)
                fmt.Println(p1)
            } else {
                if len(v4) > 0 {
                    v4[0] = (fi5 << (((p1 - uint(p0))) & 63))
                }
            }
        }
        switch "foo bar" {
        case "hello":
            v8 := GenS1{F0: 850, F1: 42, F2: 1}
            _ = v8
        default:
            p0 = 2
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}
func genFn1(p0 int64) bool {
    for fi9 := 0; fi9 < 10; fi9++ {
        p0 += (p0 + p0)
        if (true && true) {
            var v10 int = 42
            var v11 uint16 = 806
            fmt.Println(v11, fi9, fi9)
            _ = v10
            _ = v11
        } else {
            v12 := GenS2{F0: (fi9 * fi9), F1: 0, F2: (fi9 & fi9)}
            _ = v12
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            p0 = p0
            p0 -= -282
            v14 := GenS3{F0: (-(fi9)), F1: 32767, F2: (-2147483648 == fi9)}
            _ = fi13
            _ = v14
        }
        _ = fi9
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS4{F0: 2147483647}
        p0 -= 1
        var v17 int = 9223372036854775807
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: 32768, F1: genFn0(772, 746, 76), F2: 42}
    for fi19 := 0; fi19 < 1; fi19++ {
        p0 = p0
        _ = fi19
    }
    p0 = p0
    _ = p0
    _ = v18
    return false
}

func main() {
    var v1 int8 = 0
    var v2 uint8 = 1
    v3 := []string{"\n\t", "", "foo bar", ""}
    for fi20 := 0; fi20 < 12; fi20++ {
        _ = len(v3)
        if false {
            _ = len(v3)
            var v21 uint32 = 65535
            var v22 float32 = 1.5
            _ = v21
            _ = v22
        }
        v2 += v2
        _ = fi20
    }
    _ = len(v3)
    v1 += (v1 | (-(v1)))
    _ = len(v3)
    var v23 int16 = 0
    var v24 int32 = 1
    fmt.Println(v23, v23)
    var v25 uint32 = 413
    if (659 <= 4294967295) {
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 uint = uint(v2)
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 3; fi28++ {
            if true {
                v29 := GenS1{F0: 4294967295, F1: v23, F2: -417}
                if ("\n\t" == "\n\t") {
                    v30 := GenS4{F0: v25}
                    _ = v30
                }
                _ = v29
            } else {
                var v31 float64 = 0.0
                _ = v31
            }
            if (65535 <= 390) {
                fmt.Println(v25)
            }
            _ = fi28
        }
    } else {
        v32 := GenS2{F0: 0, F1: (^(v24)), F2: 462}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
