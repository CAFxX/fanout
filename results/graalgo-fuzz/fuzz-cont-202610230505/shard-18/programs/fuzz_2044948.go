package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 bool }
func genFn0(p0 uint16, p1 uint8, p2 int8) uint64 {
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 uint32 = 1
        if true {
            p2 = (p2 % ((42) | 1))
        }
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(p1, p1, p0)
        var v7 float64 = (-2.25 - 0.1)
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 671
}
func genFn1(p0 int16, p1 uint32, p2 uint) uint64 {
    var v8 float64 = ((0.0 - 3.14159) * 1e10)
    v9 := GenS1{F0: 0}
    if false {
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS1{F0: p2}
            v8 = v8
            if true {
                fmt.Println(p1, p1)
                v12 := []int{(fi10 >> (((p2 | p2)) & 63)), -913, fi10, (fi10 + (fi10 - fi10))}
                _ = v12
            } else {
                switch (fi10 * fi10) {
                case -91:
                    v13 := GenS1{F0: (p2 | p2)}
                    _ = v13
                }
                var v14 float32 = 0.1
                _ = v14
            }
            _ = fi10
            _ = v11
        }
    } else {
        v15 := GenS2{F0: -658}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return genFn0(2, 128, -128)
}

func main() {
    v1 := []uint32{}
    var v2 int32 = -1
    var v3 int64 = -9223372036854775808
    var v16 uint16 = 32767
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 float64 = 1.5
        if (-8 > -62) {
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 int32 = 0
                _ = fi19
                _ = v20
            }
            v21 := GenS3{F0: fi17, F1: 127, F2: 42}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v21
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v3, v16)
    v3 = (v3 / ((v3) | 1))
    v22 := GenS4{F0: (v16 | v16)}
    if false {
        v23 := GenS2{F0: (v2 * int32(v3))}
        var v24 uint64 = 42
        switch 4294967295 {
        case 32768:
            fmt.Println(v3, v24, v3)
        case 42:
            var v25 uint8 = 0
            _ = v25
        }
        _ = v23
        _ = v24
    } else {
        var v26 int = 128
        v2 += 0
        var v27 string = ("foo bar" + "foo bar")
        _ = v26
        _ = v27
    }
    var v28 string = "foo bar"
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    var v29 bool = true
    var v30 uint16 = (v16 - v16)
    var v31 uint = 770
    if (!v29) {
        for fi32 := 0; fi32 < 15; fi32++ {
            for fi33 := 0; fi33 < 8; fi33++ {
                if (255 > v16) {
                    v34 := GenS5{F0: v2}
                    switch (fi32 + -782) {
                    case -1:
                        switch v28 {
                        case "\n\t":
                            var v35 uint = (^(v31))
                            _ = v35
                        case "hello":
                            if v29 {
                                switch fi33 {
                                case -2147483648:
                                    fmt.Println(v30, v31)
                                case 128:
                                    v3 -= -1
                                default:
                                    var v36 float32 = float32(1.5)
                                    _ = v36
                                }
                                v37 := GenS6{F0: fi33, F1: (-1 == 67)}
                                _ = v37
                            }
                        }
                    }
                    v1 = append(v1, 65535)
                    _ = v34
                } else {
                    v28 = ((v28 + (v28 + "")) + v28)
                    v38 := GenS4{F0: v16}
                    _ = v38
                }
                var v39 bool = v29
                _ = fi33
                _ = v39
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi32
        }
        if (!v29) {
            _ = len(v1)
        }
        for fi40 := 0; fi40 < 15; fi40++ {
            _ = len(v1)
            for fi41 := 0; fi41 < 20; fi41++ {
                if len(v1) > 0 {
                    v1[0] = 65535
                }
                _ = fi41
            }
            for fi42 := 0; fi42 < 8; fi42++ {
                v43 := GenS5{F0: (v2 / ((v2) | 1))}
                v28 = v28
                var v44 uint16 = (v16 % (((v30 >> (((v31 << (((v31 | v31)) & 63))) & 63))) | 1))
                _ = fi42
                _ = v43
                _ = v44
            }
            _ = fi40
        }
    }
    switch -9223372036854775808 {
    case 2147483647:
        v29 = v29
    case 9223372036854775807:
        _ = len(v1)
    case -320:
        if len(v1) > 0 {
            v1[0] = 1
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
