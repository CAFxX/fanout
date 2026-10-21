package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint64 }
func genFn0(p0 int32, p1 uint32, p2 int32) uint32 {
    p2 = p2
    var v5 int32 = p0
    for fi6 := 0; fi6 < 8; fi6++ {
        switch (fi6 % ((fi6) | 1)) {
        case -812:
            if (9223372036854775807 > 0) {
                v7 := []int{(fi6 - 2147483647), fi6, fi6, (fi6 ^ fi6)}
                v8 := GenS1{F0: 331}
                var v9 int = fi6
                _ = v7
                _ = v8
                _ = v9
            }
        case 0:
            fmt.Println(fi6)
        }
        var v10 float32 = ((0.1 / 1.0) * 1e10)
        p1 -= 1
        _ = fi6
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p1 * p1)
}
func genFn1() uint16 {
    var v11 int = -718
    var v12 int64 = 625
    v12 = (v12 & v12)
    v11 = v11
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 float32 = 1.5
    var v2 int32 = -45
    v3 := make(map[uint32]bool)
    var v4 uint32 = 1
    fmt.Println(v4)
    v2 = v2
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 bool = (127 == 42)
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 int = fi15
            v2 += -1
            v16 -= fi15
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    v2 = 2147483647
    fmt.Println(v2, v4, v2)
    var v17 float32 = (v1 * (1e10 + 3.14159))
    for fi18 := 0; fi18 < 11; fi18++ {
        fmt.Println(v2, v17, v4)
        fmt.Println(v17, v2, v4)
        switch fi18 {
        case -9223372036854775808:
            v19 := GenS1{F0: 1000}
            _ = v19
        case 42:
            v17 = (v1 + v17)
        }
        _ = fi18
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := []int{fi20, fi20}
        var v22 int64 = 9223372036854775807
        v23 := v3[v4]
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 int8 = 64
    if (350 > 0) {
        var v25 uint32 = genFn0(42, (v4 << ((65535) & 63)), v2)
        if (!true) {
            v26 := GenS1{F0: 32767}
            v24 = (v24 << ((99) & 63))
            var v27 bool = (60 < 42)
            _ = v26
            _ = v27
        }
        v28 := GenS2{F0: 2, F1: -1}
        _ = v25
        _ = v28
    }
    if (0 != -346) {
        if false {
            var v29 int8 = (v24 >> ((255) & 63))
            switch "a" {
            case "foo bar":
                v2 = v2
            case "hello":
                if (1 == v2) {
                    v30 := GenS1{F0: 633}
                    v17 -= (float32(3.14159) + 3.14159)
                    fmt.Println(v1)
                    _ = v30
                } else {
                    var v31 bool = true
                    v32 := GenS3{F0: 2, F1: (v2 / ((v2) | 1)), F2: 42}
                    var v33 int32 = v2
                    _ = v31
                    _ = v32
                    _ = v33
                }
            case "a":
                var v34 string = "a"
                _ = v34
            }
            _ = v29
        }
        switch "hello" {
        case "a":
            var v35 uint8 = 128
            _ = v35
        case "hello":
            v36 := []int{4294967295, -5, 220}
            _ = v36
        case "foo bar":
            for fi37 := 0; fi37 < 12; fi37++ {
                switch fi37 {
                case 127:
                    fmt.Println(v1, v2)
                }
                fmt.Println(v2, v24, v4)
                v24 = v24
                _ = fi37
            }
        }
        v1 += v17
    }
    delete(v3, v4)
    if false {
        v24 = 0
        var v38 uint32 = (v4 & (^(v4)))
        _ = v38
    } else {
        switch -588 {
        case 64:
            fmt.Println(v17, v1, v4)
        case -2147483648:
            for fi39 := 0; fi39 < 11; fi39++ {
                if (18446744073709551615 > 0) {
                    for fi40 := 0; fi40 < 20; fi40++ {
                        fmt.Println(fi40, v4)
                        fmt.Println(fi39, v4)
                        _ = fi40
                    }
                    v3[(v4 + (v4 | (v4 | (^(v4)))))] = (1 == 42)
                    v24 = v24
                }
                for fi41 := 0; fi41 < 2; fi41++ {
                    var v42 int = -782
                    v24 -= -1
                    var v43 uint = 1
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                fmt.Println(v2, fi39)
                _ = fi39
            }
        default:
            var v44 int8 = (v24 | -128)
            _ = v44
        }
    }
    var v45 uint32 = genFn0(v2, (v4 | v4), v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v45
}
