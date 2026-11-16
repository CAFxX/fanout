package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint8; F1 bool; F2 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 int }
func genFn0(p0 uint32, p1 uint) int8 {
    var v7 uint32 = (p0 / (((p0 - p0)) | 1))
    v8 := GenS1{F0: 1}
    var v9 int64 = 7
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int16 = 1
    v2 := make(map[string]uint32)
    var v3 int64 = 191
    var v4 bool = ((270 > 4294967295) || (315 == (^(v1))))
    var v5 uint64 = 1
    var v6 float64 = 1e10
    if true {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int16 = v1
            _ = fi10
            _ = v11
        }
    } else {
        v12 := v2["a"]
        for fi13 := 0; fi13 < 13; fi13++ {
            switch fi13 {
            case 711:
                fmt.Println(fi13)
            case 32767:
                v14 := []int{(fi13 % ((fi13) | 1)), (^(fi13)), int(v3)}
                _ = v14
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                v1 -= 390
                if v4 {
                    var v16 uint64 = v5
                    var v17 int8 = 1
                    _ = v16
                    _ = v17
                } else {
                    v1 += v1
                    var v18 float32 = 1e10
                    v19 := GenS2{F0: (v12 | v12), F1: (v1 << ((255) & 63))}
                    _ = v18
                    _ = v19
                }
                var v20 float32 = 3.14159
                _ = fi15
                _ = v20
            }
            _ = fi13
        }
        v4 = (!(v5 >= (v5 & (v5 / (((v5 % ((v5) | 1))) | 1)))))
        _ = v12
    }
    v1 += (v1 ^ 1)
    if ((int(v1) >= 910) || v4) {
        switch "hello" {
        case "hello":
            v6 -= ((0.1 - v6) / v6)
        case "foo bar":
            for fi21 := 0; fi21 < 9; fi21++ {
                v5 -= v5
                switch (fi21 ^ (fi21 >> ((1) & 63))) {
                case 128:
                    if (!((!v4) || v4)) {
                        var v22 uint16 = 0
                        _ = v22
                    }
                case 727:
                    v23 := GenS3{F0: 128, F1: false, F2: 100}
                    _ = v23
                case 0:
                    var v24 bool = v4
                    _ = v24
                default:
                    v1 -= 42
                }
                var v25 int8 = 100
                _ = fi21
                _ = v25
            }
        }
        if (!(!false)) {
            if v4 {
                if (((0.0 + (1.5 * 0.0)) != 1.5) || true) {
                    v5 -= (v5 | (v5 + (v5 >> ((293) & 63))))
                }
                fmt.Println(v3, v6, v5)
                var v26 bool = true
                _ = v26
            }
            var v27 uint8 = 255
            delete(v2, ("a" + "foo bar"))
            _ = v27
        }
    }
    v28 := GenS4{F0: v5}
    switch 9223372036854775807 {
    case -9223372036854775808:
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            switch ((((("foo bar" + (("a" + "a") + (("foo bar" + "") + "\n\t"))) + "a") + "a") + "") + "a") {
            case "a":
                var v29 int16 = (v1 | (-(v1)))
                _ = v29
            case "hello":
                fmt.Println(v4, v3, v4)
            }
        default:
            switch "" {
            case "a":
                for fi30 := 0; fi30 < 11; fi30++ {
                    if v4 {
                        v31 := []int{fi30, 396}
                        if (!true) {
                            for fi32 := 0; fi32 < 3; fi32++ {
                                var v33 int32 = -2147483648
                                v4 = false
                                if true {
                                    fmt.Println(v4, v3, fi32)
                                    var v34 uint16 = 153
                                    v35 := GenS3{F0: 228, F1: v4, F2: 0}
                                    _ = v34
                                    _ = v35
                                }
                                _ = fi32
                                _ = v33
                            }
                            var v36 int16 = v1
                            switch "foo bar" {
                            case "\n\t":
                                _ = len(v31)
                            case "a":
                                v5 = v5
                            case "hello":
                                fmt.Println(v4)
                            }
                            _ = v36
                        } else {
                            fmt.Println(v1)
                        }
                        v1 -= v1
                        _ = v31
                    }
                    _ = len(v2)
                    v4 = (1 != 127)
                    _ = fi30
                }
            case "foo bar":
                for fi37 := 0; fi37 < 20; fi37++ {
                    for fi38 := 0; fi38 < 20; fi38++ {
                        var v39 uint16 = 186
                        _ = fi38
                        _ = v39
                    }
                    _ = fi37
                }
            }
        }
    case 697:
        var v40 string = "\n\t"
        _ = v40
    default:
        for fi41 := 0; fi41 < 3; fi41++ {
            v5 -= (v5 * 4294967295)
            _ = fi41
        }
    }
    switch -135 {
    case -330:
        if v4 {
            var v42 int32 = 1
            _ = v42
        }
    case 9223372036854775807:
        for fi43 := 0; fi43 < 13; fi43++ {
            var v44 uint8 = 1
            var v45 uint8 = 128
            v1 -= 32767
            _ = fi43
            _ = v44
            _ = v45
        }
    case -453:
        for fi46 := 0; fi46 < 5; fi46++ {
            fmt.Println(v5, v3)
            if v4 {
                v3 += (v3 ^ (v3 - 9223372036854775807))
                for fi47 := 0; fi47 < 15; fi47++ {
                    v3 += (v3 / ((v3) | 1))
                    if (fi46 != fi46) {
                        v5 += 18446744073709551615
                        v48 := GenS2{F0: 1, F1: (-(v1))}
                        var v49 uint32 = 2147483647
                        _ = v48
                        _ = v49
                    }
                    v1 += v1
                    _ = fi47
                }
            }
            var v50 int8 = -1
            _ = fi46
            _ = v50
        }
    default:
        v4 = (764 >= int32(v3))
    }
    v1 += int16(v5)
    switch "\n\t" {
    case "a":
        v51 := GenS3{F0: 1, F1: ((v4 && (!false)) || v4), F2: genFn0(2147483647, 1)}
        _ = v51
    default:
        for fi52 := 0; fi52 < 16; fi52++ {
            var v53 uint8 = 2
            v6 -= ((v6 / 3.14159) - 3.14159)
            v54 := v2["a"]
            _ = fi52
            _ = v53
            _ = v54
        }
    }
    v55 := GenS2{F0: 4294967295, F1: (v1 - (v1 ^ 1000))}
    switch 127 {
    case 32768:
        var v56 uint8 = uint8(v5)
        _ = v56
    case 509:
        v57 := GenS5{F0: "\n\t", F1: int(v1)}
        _ = v57
    case -406:
        var v58 int = 314
        _ = v58
    }
    for fi59 := 0; fi59 < 16; fi59++ {
        fmt.Println(v5)
        _ = fi59
    }
    var v60 int64 = -369
    v1 += 217
    v5 = 18446744073709551615
    fmt.Println(v4, v1)
    v6 += float64(1e10)
    var v61 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v60
    _ = v61
}
