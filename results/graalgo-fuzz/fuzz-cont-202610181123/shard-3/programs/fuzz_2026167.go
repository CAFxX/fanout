package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint32) int8 {
    if (("a" < "foo bar") || true) {
        var v6 float64 = float64((float32((0.1 + 0.0)) + 3.14159))
        v7 := GenS1{F0: (454 == 807)}
        p0 = (p0 - 1)
        _ = v6
        _ = v7
    }
    var v8 uint16 = 2
    _ = p0
    _ = v8
    return 35
}
func genFn1(p0 int) uint8 {
    for fi9 := 0; fi9 < 11; fi9++ {
        if false {
            v10 := []int{-2147483648, (p0 >> ((255) & 63)), (p0 | 4294967295), 32768}
            var v11 int64 = 1
            v11 = v11
            _ = v10
            _ = v11
        } else {
            if false {
                switch "a" {
                case "foo bar":
                    v12 := []int{(fi9 % ((fi9) | 1)), p0}
                    _ = v12
                case "\n\t":
                    p0 += fi9
                }
                var v13 uint8 = 0
                _ = v13
            } else {
                p0 = (p0 >> ((2) & 63))
                var v14 int8 = genFn0(572)
                if (4294967295 > 42) {
                    var v15 int32 = 2147483647
                    _ = v15
                } else {
                    var v16 float32 = (0.0 * ((0.0 - -2.25) / 0.1))
                    _ = v16
                }
                _ = v14
            }
            v17 := []int{p0, 582, fi9, fi9}
            fmt.Println(fi9, fi9)
            _ = v17
        }
        p0 = 4294967295
        var v18 int = (fi9 & (p0 << ((18446744073709551615) & 63)))
        _ = fi9
        _ = v18
    }
    if (!((482 >= 1000) && true)) {
        v19 := GenS2{F0: 359}
        var v20 uint64 = 294
        _ = v19
        _ = v20
    }
    v21 := []int{(^(p0)), p0, p0, p0}
    _ = p0
    _ = v21
    return 22
}

func main() {
    var v1 int = 42
    var v2 uint8 = 2
    var v3 int8 = 42
    var v4 int32 = -1
    var v5 uint16 = 2
    if true {
        fmt.Println(v1, v1)
        if (((("a" + "\n\t") + "hello") > "foo bar") && (true && (-792 == 0))) {
            v5 += (^(v5))
            var v22 int16 = -475
            var v23 int64 = 42
            _ = v22
            _ = v23
        }
    }
    var v24 uint32 = 2147483647
    v25 := GenS1{F0: true}
    v26 := []int{(v1 + (v1 ^ (v1 >> ((829) & 63)))), (v1 >> ((264) & 63)), v1, (v1 + v1)}
    switch int(v5) {
    case 127:
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 int16 = 32767
            for fi29 := 0; fi29 < 2; fi29++ {
                v30 := GenS1{F0: true}
                _ = fi29
                _ = v30
            }
            _ = fi27
            _ = v28
        }
    default:
        v5 = (^(v5))
    }
    if len(v26) > 0 {
        v26[0] = int(v4)
    }
    v4 -= (v4 | v4)
    var v31 int16 = 42
    var v32 int8 = (v3 / (((-(v3))) | 1))
    v1 = v1
    v33 := GenS3{F0: 156}
    for fi34 := 0; fi34 < 16; fi34++ {
        if true {
            v2 -= (v2 % ((v2) | 1))
        }
        fmt.Println(v31)
        _ = len(v26)
        _ = fi34
    }
    v35 := GenS1{F0: (100 < (v3 >> ((4294967295) & 63)))}
    for fi36 := 0; fi36 < 16; fi36++ {
        if false {
            var v37 uint32 = v24
            switch fi36 {
            case 1:
                v38 := GenS2{F0: 2147483647}
                _ = v38
            case 336:
                var v39 uint16 = v5
                _ = v39
            }
            _ = v37
        } else {
            v5 -= 42
        }
        fmt.Println(v5, v31)
        for fi40 := 0; fi40 < 17; fi40++ {
            var v41 float32 = 0.1
            var v42 uint = 590
            var v43 int64 = 42
            _ = fi40
            _ = v41
            _ = v42
            _ = v43
        }
        _ = fi36
    }
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v5
}
