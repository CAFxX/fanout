package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0() bool {
    switch ("" + ("a" + ("a" + "a"))) {
    case "a":
        var v6 bool = true
        _ = v6
    }
    var v7 int8 = 100
    _ = v7
    return false
}
func genFn1(p0 int8, p1 int8) int {
    v8 := []int{861, 122, -9223372036854775808}
    v9 := GenS1{F0: 255, F1: (0 > 223), F2: p0}
    _ = len(v8)
    p0 -= p0
    for fi10 := 0; fi10 < 2; fi10++ {
        if len(v8) > 0 {
            v8[0] = 127
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 32768
}

func main() {
    var v1 uint64 = 319
    var v2 int8 = -128
    var v3 int32 = 2147483647
    v4 := make(map[int]uint32)
    var v5 string = (("\n\t" + "") + "hello")
    v3 = v3
    var v11 uint16 = 495
    for fi12 := 0; fi12 < 1; fi12++ {
        switch (("foo bar" + v5) + v5) {
        case "\n\t":
            switch v5 {
            case "hello":
                switch fi12 {
                case 32767:
                    fmt.Println(v2, v3)
                case -2147483648:
                    fmt.Println(v2, fi12)
                }
            case "foo bar":
                var v13 int16 = 42
                _ = v13
            case "a":
                v5 = v5
            }
        default:
            var v14 uint64 = v1
            _ = v14
        }
        _ = fi12
    }
    v15 := GenS2{F0: uint8(v1), F1: 753}
    v16 := []int{int(v1), genFn1(100, (v2 + v2)), -659}
    for fi17 := 0; fi17 < 6; fi17++ {
        _ = len(v4)
        if ((v3 >> ((260) & 63)) > v3) {
            var v18 uint16 = v11
            var v19 uint64 = (v1 + (v1 + v1))
            switch "\n\t" {
            case "a":
                switch (^(fi17)) {
                case -657:
                    v2 -= -128
                case 0:
                    switch v5 {
                    case "\n\t":
                        var v20 int8 = 100
                        _ = v20
                    case "a":
                        fmt.Println(v11, v19)
                    case "foo bar":
                        v19 = (v19 / ((v1) | 1))
                    }
                }
            case "":
                v4[fi17] = uint32(v19)
            case "foo bar":
                if true {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        switch (fi21 >> ((255) & 63)) {
                        case -9223372036854775808:
                            if ((fi17 >> ((4294967295) & 63)) <= 890) {
                                v5 = ("" + v5)
                                v18 = 1
                                v22 := GenS3{F0: 65535, F1: (fi21 >> ((2) & 63))}
                                _ = v22
                            } else {
                                v5 = v5
                            }
                        case -1:
                            if len(v16) > 0 {
                                v16[0] = 354
                            }
                        case -2147483648:
                            v23 := GenS2{F0: 1, F1: 746}
                            _ = v23
                        }
                        v1 = (v1 << ((255) & 63))
                        delete(v4, (fi21 / ((genFn1(100, -76)) | 1)))
                        switch (fi17 | int(v3)) {
                        case -909:
                            v18 += 538
                        default:
                            switch genFn1((v2 + (v2 / ((42) | 1))), (-(v2))) {
                            case 36:
                                v1 = (^(v1))
                            case 127:
                                fmt.Println(v5, v18)
                                v2 = v2
                                var v24 int = fi21
                                _ = v24
                            case 987:
                                v2 = (v2 + 100)
                                if true {
                                    var v25 uint64 = 768
                                    _ = v25
                                }
                            default:
                                v19 = 768
                            }
                        }
                        _ = fi21
                    }
                    var v26 int = fi17
                    _ = v26
                } else {
                    v27 := GenS3{F0: 2, F1: 4294967295}
                    _ = v27
                }
            default:
                for fi28 := 0; fi28 < 13; fi28++ {
                    v11 = (v18 / ((v18) | 1))
                    _ = fi28
                }
            }
            _ = v18
            _ = v19
        } else {
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 int8 = v2
                v3 = (v3 + 0)
                if ((1e10 - (3.14159 - float32(0.1))) > (-2.25 - (((-2.25 / -2.25) - (0.1 * 1.5)) * 0.0))) {
                    v31 := GenS2{F0: 230, F1: 2}
                    _ = v31
                }
                fmt.Println(v30, v2, fi29)
                v30 = (v30 / ((int8(v11)) | 1))
                _ = fi29
                _ = v30
            }
            var v32 int64 = -566
            for fi33 := 0; fi33 < 2; fi33++ {
                var v34 int = (fi33 % (((-(fi33))) | 1))
                if true {
                    v5 = v5
                    v32 = v32
                    if true {
                        var v35 string = "hello"
                        _ = v35
                    } else {
                        v1 = v1
                        var v36 float32 = 1.5
                        var v37 uint = 18446744073709551615
                        v36 = v36
                        _ = v36
                        _ = v37
                    }
                } else {
                    if ((v11 % (((^(v11))) | 1)) > (^(v11))) {
                        var v38 uint8 = 42
                        _ = v38
                    }
                    fmt.Println(v11, v3)
                    v39 := GenS3{F0: 2, F1: (v34 + (fi17 + fi33))}
                    _ = v39
                }
                v40 := v4[(fi33 * 1)]
                _ = fi33
                _ = v34
                _ = v40
            }
            _ = v32
        }
        var v41 uint32 = 2147483648
        _ = fi17
        _ = v41
    }
    _ = len(v16)
    for fi42 := 0; fi42 < 13; fi42++ {
        var v43 uint32 = 65535
        for fi44 := 0; fi44 < 10; fi44++ {
            if (!((v5 + "") >= v5)) {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                if len(v16) > 0 {
                    v16[0] = genFn1((v2 ^ v2), v2)
                }
            }
            _ = fi44
        }
        _ = fi42
        _ = v43
    }
    fmt.Println(v2)
    var v45 uint16 = (^(v11))
    for fi46 := 0; fi46 < 6; fi46++ {
        var v47 int64 = -9223372036854775808
        _ = fi46
        _ = v47
    }
    v5 = v5
    v48 := GenS4{F0: v5, F1: (v45 >> ((0) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v5
}
