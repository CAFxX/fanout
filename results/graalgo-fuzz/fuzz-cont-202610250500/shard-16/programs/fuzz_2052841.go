package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint; F1 int; F2 int8 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS5 struct { F0 string; F1 uint64; F2 int16 }
type GenS6 struct { F0 int64; F1 int; F2 int8 }
func genFn0(p0 int, p1 uint16) uint8 {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int64 = 0
        var v7 uint64 = 42
        switch "foo bar" {
        case "a":
            fmt.Println(v6, fi5)
        case "\n\t":
            v8 := GenS1{F0: 127, F1: 1, F2: 2147483648}
            _ = v8
        case "":
            p1 = p1
        default:
            if false {
                if false {
                    if (255 > (p1 - p1)) {
                        var v9 int = fi5
                        v10 := []int{fi5, 4294967295, (^(fi5)), v9}
                        _ = v10
                        _ = v9
                    } else {
                        v11 := []int{-1, 32768, (^(p0)), (p0 / ((398) | 1))}
                        _ = v11
                    }
                    v12 := GenS2{F0: 140, F1: 290}
                    if (true || false) {
                        p0 += (^(fi5))
                        v13 := GenS2{F0: 32767, F1: -2147483648}
                        _ = v13
                    }
                    _ = v12
                }
                var v14 float32 = float32(-2.25)
                _ = v14
            } else {
                var v15 uint8 = 128
                for fi16 := 0; fi16 < 9; fi16++ {
                    v15 -= 201
                    _ = fi16
                }
                fmt.Println(v6, p0, fi5)
                _ = v15
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 int64 = 0
        _ = fi17
        _ = v18
    }
    v19 := []int{(p0 + (p0 ^ p0)), (p0 / (((p0 ^ 2147483647)) | 1))}
    var v20 uint = 255
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 0
}
func genFn1() bool {
    var v21 uint64 = 2
    v22 := []int{-1}
    var v23 uint64 = v21
    _ = v21
    _ = v22
    _ = v23
    return true
}

func main() {
    var v1 uint = 255
    var v2 bool = true
    var v3 int8 = 0
    v4 := []int{4294967295, -691}
    v3 += (v3 & -1)
    v3 = (v3 - (v3 / ((127) | 1)))
    var v24 string = "foo bar"
    fmt.Println(v2)
    v2 = (255 >= 127)
    var v25 uint8 = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v1 += 42
        v1 -= v1
        _ = fi26
    }
    var v27 uint8 = (^(v25))
    for fi28 := 0; fi28 < 15; fi28++ {
        fmt.Println(v2, v3, v3)
        if v2 {
            fmt.Println(v27, fi28, v2)
            var v29 uint64 = 18446744073709551615
            _ = v29
        } else {
            v30 := GenS3{F0: v1, F1: (-(fi28)), F2: (-(v3))}
            for fi31 := 0; fi31 < 17; fi31++ {
                var v32 int16 = 0
                v33 := GenS4{F0: 42, F1: 760, F2: (v25 | genFn0(127, 65535))}
                _ = fi31
                _ = v32
                _ = v33
            }
            var v34 int8 = v3
            _ = v30
            _ = v34
        }
        if v2 {
            var v35 int32 = 538
            if (!v2) {
                for fi36 := 0; fi36 < 13; fi36++ {
                    v37 := GenS5{F0: "foo bar", F1: 364, F2: 42}
                    var v38 string = "hello"
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                fmt.Println(v1, v35)
                switch fi28 {
                case 503:
                    for fi39 := 0; fi39 < 15; fi39++ {
                        fmt.Println(v3, v1, v25)
                        var v40 uint64 = 2
                        v3 = (v3 | (v3 - v3))
                        _ = fi39
                        _ = v40
                    }
                default:
                    var v41 uint8 = (v27 % (((^(v27))) | 1))
                    _ = v41
                }
            } else {
                var v42 int32 = int32(fi28)
                _ = v42
            }
            switch -9223372036854775808 {
            case -9223372036854775808:
                v43 := GenS6{F0: 1, F1: 710, F2: (^(v3))}
                _ = v43
            case -2147483648:
                for fi44 := 0; fi44 < 2; fi44++ {
                    fmt.Println(v1, v35)
                    v25 = (^(v25))
                    _ = fi44
                }
            case 32767:
                var v45 float64 = 3.14159
                _ = v45
            default:
                fmt.Println(v3, v25)
            }
            _ = v35
        }
        _ = fi28
    }
    v46 := GenS4{F0: -64, F1: 790, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v46
}
