package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 uint; F1 bool; F2 int8 }
func genFn0(p0 bool, p1 string) bool {
    p0 = ((2 < 128) || false)
    fmt.Println(p0, p0)
    var v4 float64 = 0.1
    var v5 int16 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 bool, p1 bool) uint8 {
    v6 := GenS1{F0: (692 <= 2)}
    var v7 int32 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 277
    var v3 bool = false
    v2 = (^(v2))
    v8 := GenS1{F0: false}
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 float32 = 1e10
        if (4294967295 == 234) {
            fmt.Println(fi9, v10, v2)
        } else {
            var v11 float64 = 1e10
            for fi12 := 0; fi12 < 7; fi12++ {
                v1[1] = false
                var v13 int64 = -9223372036854775808
                var v14 uint32 = 1
                _ = fi12
                _ = v13
                _ = v14
            }
            switch "foo bar" {
            case "\n\t":
                if v3 {
                    fmt.Println(fi9, fi9)
                    if (!(18446744073709551615 <= 4294967295)) {
                        fmt.Println(v3, v10)
                        v3 = v3
                    } else {
                        v3 = true
                        v15 := []int{fi9, 496, (v2 << ((65535) & 63))}
                        _ = v15
                    }
                    if (2147483647 >= 2) {
                        v2 = (fi9 | 127)
                        var v16 uint = 881
                        _ = v16
                    }
                }
            }
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    var v17 int8 = -115
    v18 := []int{v2, 9223372036854775807, v2, v2}
    if v3 {
        switch ("foo bar" + "a") {
        case "a":
            v19 := v1[2147483647]
            _ = v19
        case "foo bar":
            var v20 bool = true
            _ = v20
        case "":
            var v21 uint8 = 0
            _ = v21
        }
        v22 := GenS2{F0: 42, F1: 1}
        _ = v22
    }
    if v3 {
        if true {
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 uint64 = 42
                fmt.Println(v2)
                v2 += 470
                _ = fi23
                _ = v24
            }
            v2 = (v2 >> ((65535) & 63))
        } else {
            v25 := GenS2{F0: 2, F1: 4294967295}
            v2 = 1
            _ = v25
        }
    }
    v2 = -556
    var v26 float32 = -2.25
    switch "hello" {
    case "hello":
        for fi27 := 0; fi27 < 20; fi27++ {
            for fi28 := 0; fi28 < 10; fi28++ {
                if (0 != -9223372036854775808) {
                    v2 -= 128
                    if false {
                        _ = len(v18)
                        var v29 int64 = 1
                        _ = v29
                    }
                    _ = len(v1)
                } else {
                    v2 = (-(fi27))
                    v30 := GenS2{F0: 32767, F1: 2147483647}
                    _ = v30
                }
                if v3 {
                    fmt.Println(fi27)
                    v2 = fi27
                } else {
                    v31 := GenS1{F0: true}
                    v2 = v2
                    fmt.Println(fi28, v3, fi28)
                    _ = v31
                }
                _ = fi28
            }
            switch (fi27 | 4294967295) {
            case -492:
                v1[207] = v3
            case -64:
                var v32 int = (fi27 ^ (^(fi27)))
                _ = v32
            case 32768:
                if v3 {
                    switch (fi27 / ((fi27) | 1)) {
                    case 113:
                        switch "\n\t" {
                        case "a":
                            v33 := GenS2{F0: 1, F1: 858}
                            _ = v33
                        case "\n\t":
                            if v3 {
                                v1[2] = v3
                            } else {
                                v3 = (!(v17 != (-(v17))))
                                v34 := GenS2{F0: 65535, F1: 65535}
                                _ = v34
                            }
                        case "":
                            fmt.Println(v3, v26, v3)
                        default:
                            v17 += v17
                        }
                    case -426:
                        var v35 uint64 = 4294967295
                        _ = v35
                    case 0:
                        var v36 uint16 = 750
                        _ = v36
                    }
                    v37 := GenS2{F0: 0, F1: 446}
                    _ = v37
                } else {
                    delete(v1, 2)
                    var v38 int = fi27
                    _ = len(v18)
                    _ = v38
                }
            default:
                v1[517] = v3
            }
            var v39 int8 = 0
            _ = fi27
            _ = v39
        }
    case "a":
        switch (-(v2)) {
        case 42:
            v40 := GenS3{F0: 78, F1: v17}
            _ = v40
        }
    case "\n\t":
        for fi41 := 0; fi41 < 18; fi41++ {
            for fi42 := 0; fi42 < 10; fi42++ {
                v2 += 32768
                v43 := GenS4{F0: 127, F1: genFn1(v3, v3)}
                _ = fi42
                _ = v43
            }
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            _ = fi41
        }
    default:
        v44 := GenS3{F0: 2, F1: 42}
        _ = v44
    }
    v45 := GenS5{F0: 4294967295, F1: 775}
    v46 := GenS6{F0: 110, F1: (((("" + "foo bar") <= "\n\t") && (0 <= 2147483647)) || (0 >= 1)), F2: v17}
    fmt.Println(v26, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v45
    _ = v46
    _ = v8
}
