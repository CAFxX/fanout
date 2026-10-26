package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
func genFn0(p0 string, p1 int8) uint8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        switch (fi6 + (fi6 / ((fi6) | 1))) {
        case -464:
            if (!(-932 <= 0)) {
                var v7 uint32 = 65535
                var v8 uint64 = 2
                _ = v7
                _ = v8
            } else {
                if ((60 <= 4294967295) || false) {
                    fmt.Println(fi6, fi6, fi6)
                    switch fi6 {
                    case 777:
                        p0 = "foo bar"
                    case -539:
                        v9 := []int{fi6}
                        _ = v9
                    default:
                        var v10 uint8 = 127
                        _ = v10
                    }
                    p1 += (p1 | p1)
                }
            }
        case 42:
            if false {
                p1 -= (p1 >> ((1) & 63))
            } else {
                fmt.Println(fi6, p0, fi6)
            }
        case 1:
            var v11 string = (p0 + ("\n\t" + p0))
            _ = v11
        default:
            p0 = p0
        }
        var v12 int32 = 2147483647
        v13 := GenS1{F0: 2}
        _ = fi6
        _ = v12
        _ = v13
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 41
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 3.14159
    var v3 uint = 42
    v4 := make(map[string]string)
    v5 := []uint32{42}
    var v14 float32 = 1.5
    v15 := GenS1{F0: 0}
    var v16 uint = v3
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS1{F0: 2}
        _ = fi17
        _ = v18
    }
    switch -423 {
    case 2147483647:
        v3 = (v16 % (((v16 | 818)) | 1))
    case 9223372036854775807:
        for fi19 := 0; fi19 < 14; fi19++ {
            switch -1 {
            case 4294967295:
                v2 = v2
            case 127:
                v20 := GenS2{F0: 2}
                _ = v20
            case 756:
                for fi21 := 0; fi21 < 10; fi21++ {
                    fmt.Println(v16, v2)
                    v22 := GenS3{F0: 1, F1: 18446744073709551615}
                    if (689 != 513) {
                        v16 = (v3 % ((18446744073709551615) | 1))
                        if (fi19 != fi21) {
                            v3 -= (v3 << (((v16 | v3)) & 63))
                            var v23 int = 32768
                            _ = v23
                        } else {
                            fmt.Println(v14, v3)
                            switch fi21 {
                            case -9223372036854775808:
                                var v24 uint16 = 32767
                                _ = v24
                            default:
                                v25 := v4[((("" + "\n\t") + "foo bar") + ("a" + ("" + "foo bar")))]
                                _ = v25
                            }
                            fmt.Println(v14)
                            v1[0] = 2147483647
                        }
                        if (!((true && false) && (!true))) {
                            var v26 int32 = 0
                            _ = v26
                        } else {
                            v14 = ((0.1 - v14) + 1e10)
                        }
                        var v27 uint8 = 42
                        v16 += (v16 & v3)
                        v28 := GenS3{F0: 385, F1: 18446744073709551615}
                        _ = v27
                        _ = v28
                    } else {
                        v29 := GenS4{F0: -1, F1: 836, F2: true}
                        _ = v29
                    }
                    if len(v5) > 0 {
                        v5[0] = 289
                    }
                    fmt.Println(v3, fi19)
                    _ = fi21
                    _ = v22
                }
            }
            for fi30 := 0; fi30 < 4; fi30++ {
                _ = len(v5)
                var v31 int8 = 100
                v32 := GenS4{F0: 0, F1: 1, F2: (42 == -2147483648)}
                _ = fi30
                _ = v31
                _ = v32
            }
            if ((false && (1 < 171)) || false) {
                var v33 int16 = -940
                _ = v33
            } else {
                delete(v1, 0)
                switch fi19 {
                case -2147483648:
                    var v34 string = ""
                    _ = v34
                case 42:
                    var v35 bool = (fi19 < (fi19 >> ((v16) & 63)))
                    _ = v35
                case 127:
                    v14 -= (((1e10 * v14) - v14) * 3.14159)
                default:
                    var v36 uint8 = uint8(v3)
                    _ = v36
                }
            }
            _ = fi19
        }
    case -2147483648:
        v37 := GenS1{F0: 0}
        _ = v37
    }
    v5 = append(v5, 567)
    if false {
        v2 -= (((3.14159 * 0.1) + v2) / -2.25)
        var v38 uint32 = 312
        _ = v38
    }
    v16 = v16
    for fi39 := 0; fi39 < 9; fi39++ {
        v40 := v1[2147483648]
        _ = fi39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
