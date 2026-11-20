package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int32; F1 string; F2 uint }

func main() {
    v1 := []int{538}
    v2 := make(map[int]int)
    v3 := []string{("\n\t" + "hello"), "", "a", ""}
    var v4 int16 = 32767
    var v5 uint32 = 2147483648
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 32767}
        switch -1 {
        case 0:
            switch fi6 {
            case -94:
                fmt.Println(v5, v4)
            case 32768:
                var v8 uint = 65535
                _ = v8
            case 2147483647:
                if false {
                    var v9 uint = 2
                    if (!((("" + "hello") > "a") || (!(973 <= (v9 | (v9 % ((0) | 1))))))) {
                        v5 = (v5 & (v5 ^ uint32(v9)))
                        v4 += v4
                        v4 = (v4 << ((v9) & 63))
                    }
                    _ = v9
                }
            default:
                var v10 uint32 = 536
                _ = v10
            }
        default:
            var v11 int16 = (-(v4))
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    var v12 int8 = -92
    switch -982 {
    case 42:
        var v13 uint16 = 0
        _ = v13
    case -1:
        var v14 float32 = ((float32((((-2.25 - (-2.25 + float64(-2.25))) / (float64(float32(float64(0.1))) - float64((float32(1e10) / -2.25)))) + (1.5 / 0.1))) * 1.5) - ((1e10 + float32(float64(1e10))) / -2.25))
        _ = v14
    default:
        switch 1 {
        case 32768:
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
        case 2147483647:
            _ = len(v2)
        case -208:
            var v15 uint8 = 2
            _ = v15
        default:
            for fi16 := 0; fi16 < 1; fi16++ {
                if (127 < 42) {
                    v4 += (v4 - (v4 << ((42) & 63)))
                    switch -343 {
                    case 32767:
                        for fi17 := 0; fi17 < 5; fi17++ {
                            v18 := GenS2{F0: ""}
                            v19 := GenS3{F0: "hello", F1: -2147483648}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                }
                for fi20 := 0; fi20 < 6; fi20++ {
                    var v21 int64 = 0
                    _ = fi20
                    _ = v21
                }
                var v22 int32 = -2147483648
                _ = fi16
                _ = v22
            }
        }
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        v4 = 80
        var v24 uint32 = v5
        fmt.Println(v12)
        _ = fi23
        _ = v24
    }
    if ("a" >= "a") {
        for fi25 := 0; fi25 < 5; fi25++ {
            switch int(v4) {
            case 1:
                v4 += (v4 * v4)
            default:
                v4 = 32767
            }
            v2[fi25] = fi25
            _ = fi25
        }
        if (312 == 429) {
            v4 -= (v4 ^ v4)
            for fi26 := 0; fi26 < 16; fi26++ {
                v27 := GenS4{F0: 618, F1: int16(v12)}
                for fi28 := 0; fi28 < 19; fi28++ {
                    var v29 float64 = 0.1
                    _ = fi28
                    _ = v29
                }
                _ = fi26
                _ = v27
            }
            if true {
                fmt.Println(v12, v5, v4)
                for fi30 := 0; fi30 < 3; fi30++ {
                    switch fi30 {
                    case 391:
                        switch fi30 {
                        case 127:
                            var v31 float64 = (0.0 - 0.0)
                            _ = v31
                        case -2147483648:
                            for fi32 := 0; fi32 < 2; fi32++ {
                                var v33 uint8 = 128
                                var v34 int32 = -2147483648
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                        case -721:
                            v5 = v5
                        }
                    case 4294967295:
                        if false {
                            var v35 string = ((("" + ("a" + "hello")) + ("\n\t" + "a")) + ("a" + ("hello" + ("hello" + "foo bar"))))
                            _ = v35
                        }
                    case 127:
                        v36 := GenS5{F0: 513, F1: "\n\t", F2: 0}
                        _ = v36
                    default:
                        v37 := GenS3{F0: "\n\t", F1: 103}
                        _ = v37
                    }
                    var v38 uint8 = 0
                    fmt.Println(v4)
                    _ = fi30
                    _ = v38
                }
            } else {
                if false {
                    v3 = append(v3, "foo bar")
                    v39 := GenS2{F0: "a"}
                    delete(v2, 127)
                    _ = v39
                } else {
                    var v40 uint = 2
                    _ = v40
                }
            }
        }
        for fi41 := 0; fi41 < 1; fi41++ {
            v42 := v2[(fi41 << ((111) & 63))]
            v42 -= fi41
            v43 := GenS4{F0: 1, F1: (^(v4))}
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    var v44 uint32 = (v5 / (((v5 * v5)) | 1))
    var v45 float64 = (1e10 / ((0.0 - 1e10) / 3.14159))
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    for fi46 := 0; fi46 < 10; fi46++ {
        v44 = (v5 % (((v44 + (v44 + 2147483648))) | 1))
        var v47 int8 = v12
        _ = fi46
        _ = v47
    }
    v45 += (3.14159 - float64(1e10))
    var v48 int32 = -2147483648
    fmt.Println(v12, v12)
    switch 9223372036854775807 {
    case -1:
        var v49 int16 = (v4 + v4)
        _ = v49
    case 1:
        _ = len(v2)
    default:
        v4 += v4
    }
    var v50 int8 = (v12 + (^(v12)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v48
    _ = v5
    _ = v50
}
