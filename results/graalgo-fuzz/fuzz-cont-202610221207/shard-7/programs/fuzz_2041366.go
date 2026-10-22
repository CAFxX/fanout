package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint8; F1 uint; F2 int64 }
type GenS5 struct { F0 string; F1 bool; F2 uint16 }
func genFn0(p0 bool, p1 int32) uint16 {
    var v6 string = "a"
    v6 = v6
    if (265 != -1) {
        for fi7 := 0; fi7 < 18; fi7++ {
            if false {
                switch fi7 {
                case 0:
                    for fi8 := 0; fi8 < 6; fi8++ {
                        v9 := []int{}
                        v10 := GenS1{F0: 885, F1: 0}
                        fmt.Println(p1, p1)
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                case 2147483647:
                    var v11 uint32 = 363
                    _ = v11
                case -1:
                    var v12 bool = p0
                    _ = v12
                }
            }
            if false {
                for fi13 := 0; fi13 < 6; fi13++ {
                    p1 = 156
                    var v14 uint8 = 0
                    var v15 uint64 = 578
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            } else {
                var v16 float64 = 1e10
                var v17 uint64 = 164
                for fi18 := 0; fi18 < 6; fi18++ {
                    v19 := GenS2{F0: -1, F1: 1}
                    v20 := []int{fi7}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = v16
                _ = v17
            }
            v6 = v6
            _ = fi7
        }
        if (!(!p0)) {
            v6 = v6
        }
    } else {
        switch v6 {
        case "a":
            v21 := GenS2{F0: -128, F1: 18446744073709551615}
            _ = v21
        case "\n\t":
            v6 = "a"
        case "":
            p1 = p1
        }
        v22 := GenS3{F0: 482, F1: 528}
        _ = v22
    }
    fmt.Println(v6)
    v23 := GenS1{F0: 4294967295, F1: 2}
    _ = p0
    _ = p1
    _ = v23
    _ = v6
    return 867
}

func main() {
    var v1 int64 = 42
    v2 := []int{9223372036854775807, 1, 2147483647}
    var v3 uint64 = 739
    var v4 int16 = 0
    v5 := []uint32{211}
    v24 := GenS1{F0: 2147483648, F1: 4294967295}
    if len(v2) > 0 {
        v2[0] = 32768
    }
    v25 := GenS2{F0: 42, F1: 306}
    fmt.Println(v1)
    v3 -= v3
    v26 := GenS1{F0: 884, F1: 18446744073709551615}
    fmt.Println(v4, v1)
    v27 := GenS4{F0: 2, F1: 255, F2: v1}
    v5 = append(v5, 693)
    fmt.Println(v1, v3, v3)
    v28 := GenS2{F0: 100, F1: 42}
    var v29 uint32 = 704
    switch -9223372036854775808 {
    case 1:
        var v30 string = ("foo bar" + "foo bar")
        _ = v30
    case 32768:
        fmt.Println(v1, v1, v29)
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        for fi32 := 0; fi32 < 7; fi32++ {
            v33 := GenS3{F0: 18446744073709551615, F1: fi32}
            v29 += (v29 % (((v29 | v29)) | 1))
            if (false && false) {
                switch (fi32 + int(v29)) {
                case 127:
                    if (true || true) {
                        v3 += 2
                        fmt.Println(v29, v4)
                    } else {
                        if true {
                            if false {
                                switch (-(fi31)) {
                                case -194:
                                    var v34 uint8 = 2
                                    var v35 int8 = -27
                                    v4 = (v4 | -590)
                                    _ = v34
                                    _ = v35
                                }
                                v36 := GenS4{F0: 2, F1: 4294967295, F2: v1}
                                v5 = append(v5, v29)
                                v1 += v1
                                if ((-(fi32)) >= (fi31 - 32767)) {
                                    fmt.Println(fi32, fi32, fi32)
                                    if false {
                                        var v37 int = fi32
                                        _ = v37
                                    }
                                    v38 := GenS5{F0: "\n\t", F1: (1e10 <= -2.25), F2: 222}
                                    _ = v38
                                }
                                _ = v36
                            }
                            switch fi31 {
                            case -28:
                                v4 = 1
                            case 128:
                                var v39 float32 = float32(float64((((1.5 * -2.25) * 1.5) - (-2.25 / ((((3.14159 + (1.5 * (((3.14159 + 1e10) - (1.5 * 1.5)) / (1e10 - 3.14159)))) / -2.25) + 1e10) - 0.0)))))
                                _ = v39
                            case -1:
                                var v40 uint = 65535
                                _ = v40
                            default:
                                if (!(v1 != 1)) {
                                    v1 = int64(v3)
                                    fmt.Println(fi31)
                                }
                            }
                        } else {
                            var v41 int = (fi32 >> ((2) & 63))
                            v4 += (-(v4))
                            _ = v41
                        }
                        v1 = (v1 + (v1 >> ((uint(v29)) & 63)))
                    }
                default:
                    if (false && (1 <= -1)) {
                        v4 = v4
                    }
                }
            }
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
