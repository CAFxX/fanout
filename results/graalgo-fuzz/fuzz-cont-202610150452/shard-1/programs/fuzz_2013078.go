package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int64; F2 int32 }
type GenS4 struct { F0 int32; F1 int8; F2 uint16 }
type GenS5 struct { F0 int64; F1 uint16; F2 uint8 }

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 uint32 = 445
    var v5 uint = 183
    v5 = 65535
    var v6 int16 = 32767
    var v7 float32 = (1.5 / 1.5)
    for fi8 := 0; fi8 < 17; fi8++ {
        if (true || true) {
            v6 += v6
            for fi9 := 0; fi9 < 12; fi9++ {
                if true {
                    v10 := GenS1{F0: (fi9 * (fi8 & -76))}
                    if true {
                        fmt.Println(v5)
                    }
                    fmt.Println(fi9, fi8)
                    var v11 uint16 = 255
                    fmt.Println(v7, v4)
                    _ = v10
                    _ = v11
                } else {
                    var v12 uint32 = (v4 | v4)
                    var v13 int16 = int16(fi8)
                    _ = v12
                    _ = v13
                }
                v14 := []int{(fi9 + (fi8 >> (((v5 % ((255) | 1))) & 63))), (fi8 % ((fi9) | 1)), (fi9 * (fi9 / (((fi9 & fi9)) | 1))), 861}
                var v15 int32 = (v1 + (^(v1)))
                _ = fi9
                _ = v14
                _ = v15
            }
            v4 -= v4
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            if (2 == 152) {
                v17 := GenS2{F0: -64}
                v1 = 2147483647
                fmt.Println(v4)
                _ = v17
            }
            v18 := []int{fi16, -9223372036854775808, (fi8 & fi16)}
            _ = fi16
            _ = v18
        }
        v5 = v5
        _ = fi8
    }
    v7 -= 0.1
    if false {
        if false {
            v19 := GenS3{F0: v4, F1: 1, F2: (v1 - 2147483647)}
            v20 := v3["a"]
            _ = v19
            _ = v20
        }
        fmt.Println(v1, v5, v7)
    } else {
        if (!(!(!(!false)))) {
            switch 672 {
            case -438:
                fmt.Println(v6, v6)
            case 4294967295:
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v6, v7, v7)
                case "foo bar":
                    for fi21 := 0; fi21 < 2; fi21++ {
                        for fi22 := 0; fi22 < 12; fi22++ {
                            v7 += float32(-2.25)
                            _ = fi22
                        }
                        for fi23 := 0; fi23 < 11; fi23++ {
                            v7 = v7
                            v24 := GenS3{F0: (^(v4)), F1: 1, F2: -106}
                            _ = fi23
                            _ = v24
                        }
                        _ = fi21
                    }
                }
            default:
                for fi25 := 0; fi25 < 9; fi25++ {
                    v26 := []int{(fi25 | -1), (fi25 >> ((v5) & 63))}
                    v27 := GenS4{F0: (-(v1)), F1: 127, F2: 65535}
                    switch -1 {
                    case 9223372036854775807:
                        _ = len(v26)
                    }
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
        } else {
            switch 32768 {
            case -9223372036854775808:
                v28 := GenS1{F0: -2147483648}
                _ = v28
            case 0:
                fmt.Println(v7)
            case 32767:
                for fi29 := 0; fi29 < 17; fi29++ {
                    for fi30 := 0; fi30 < 5; fi30++ {
                        v3[("foo bar" + "\n\t")] = ("\n\t" + "hello")
                        var v31 uint8 = 244
                        _ = fi30
                        _ = v31
                    }
                    var v32 int8 = int8(v4)
                    for fi33 := 0; fi33 < 1; fi33++ {
                        if (42 > int16(v5)) {
                            var v34 float64 = 0.0
                            _ = v34
                        } else {
                            v5 += v5
                        }
                        var v35 int64 = 9223372036854775807
                        fmt.Println(v1, v35, fi29)
                        _ = fi33
                        _ = v35
                    }
                    _ = fi29
                    _ = v32
                }
            default:
                var v36 uint64 = 793
                _ = v36
            }
        }
        v37 := GenS5{F0: 7, F1: 1, F2: 46}
        _ = v37
    }
    v5 += (^(v5))
    fmt.Println(v4, v1)
    v3["a"] = "foo bar"
    fmt.Println(v6, v4)
    var v38 int16 = v6
    for fi39 := 0; fi39 < 9; fi39++ {
        v2[(v4 & (v4 >> ((v5) & 63)))] = (v4 | (v4 * (v4 << ((255) & 63))))
        var v40 bool = ((v5 < 2) || false)
        v41 := []int{(fi39 / (((fi39 + (fi39 << ((65535) & 63)))) | 1)), fi39}
        _ = fi39
        _ = v40
        _ = v41
    }
    switch "\n\t" {
    case "":
        v42 := []int{4294967295, 1, -676}
        _ = v42
    case "foo bar":
        v4 = 0
    default:
        v38 = 1
    }
    var v43 uint8 = 2
    var v44 bool = (!(82 != 5))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
