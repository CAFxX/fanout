package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 bool }
type GenS2 struct { F0 int32; F1 uint32; F2 int }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int64; F2 int32 }
func genFn0(p0 string, p1 uint) int32 {
    p0 = "hello"
    p0 = (p0 + "foo bar")
    for fi6 := 0; fi6 < 14; fi6++ {
        if (!(true && true)) {
            v7 := []int{}
            for fi8 := 0; fi8 < 2; fi8++ {
                var v9 float32 = 0.1
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            switch fi6 {
            case 957:
                if false {
                    switch (fi6 << ((0) & 63)) {
                    case 32768:
                        p1 = p1
                    default:
                        for fi10 := 0; fi10 < 18; fi10++ {
                            v11 := GenS1{F0: 42, F1: 482, F2: true}
                            _ = fi10
                            _ = v11
                        }
                    }
                } else {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        v13 := GenS1{F0: 0, F1: 221, F2: true}
                        v14 := []int{(^(fi6)), (fi12 ^ fi6), -9223372036854775808, (fi12 % ((128) | 1))}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    v15 := []int{(fi6 | fi6), 293, fi6, int(p1)}
                    var v16 uint16 = 42
                    _ = v15
                    _ = v16
                }
            case 32767:
                v17 := []int{fi6, (fi6 - (fi6 * (fi6 | fi6))), 2147483647, (fi6 >> ((4294967295) & 63))}
                _ = v17
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int, p1 bool, p2 int32) uint8 {
    v18 := GenS1{F0: -884, F1: (p2 % ((p2) | 1)), F2: p1}
    v19 := []int{32768, (-(p0))}
    v20 := GenS2{F0: (p2 & p2), F1: 1, F2: (p0 >> ((4294967295) & 63))}
    if p1 {
        p1 = (564 <= (p0 >> ((42) & 63)))
        p0 += p0
        for fi21 := 0; fi21 < 17; fi21++ {
            if true {
                for fi22 := 0; fi22 < 2; fi22++ {
                    fmt.Println(fi21, fi22)
                    var v23 uint32 = 0
                    if p1 {
                        var v24 uint16 = 65535
                        var v25 int = (fi21 | 0)
                        _ = v24
                        _ = v25
                    } else {
                        switch (fi22 / ((-188) | 1)) {
                        case 331:
                            switch fi22 {
                            case 128:
                                fmt.Println(p2, p2)
                                p1 = p1
                                v26 := GenS1{F0: -9223372036854775808, F1: p2, F2: p1}
                                var v27 uint16 = 255
                                p0 -= fi22
                                _ = v26
                                _ = v27
                            }
                        default:
                            if (!(139 > 0)) {
                                v28 := GenS1{F0: 0, F1: 42, F2: true}
                                _ = v28
                            }
                        }
                        if ((!p1) && false) {
                            p1 = p1
                        } else {
                            var v29 int16 = 966
                            var v30 int = fi22
                            _ = v29
                            _ = v30
                        }
                        v19 = append(v19, (-(p0)))
                    }
                    v31 := GenS3{F0: 744, F1: 1}
                    _ = fi22
                    _ = v23
                    _ = v31
                }
            } else {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                var v32 int32 = -2147483648
                p0 += 127
                _ = v32
            }
            _ = fi21
        }
    }
    if p1 {
        switch p0 {
        case 32767:
            if ((478 >= 65535) || (!p1)) {
                p1 = (128 != 128)
            } else {
                fmt.Println(p2, p1, p2)
            }
        case 223:
            if true {
                v19 = append(v19, (p0 * (-(p0))))
                switch ("foo bar" + (("hello" + ("" + "hello")) + ("\n\t" + ""))) {
                case "":
                    fmt.Println(p2, p0, p1)
                case "\n\t":
                    if p1 {
                        p2 -= (^(p2))
                    } else {
                        var v33 float32 = 0.0
                        v34 := GenS2{F0: (-(p2)), F1: 999, F2: p0}
                        _ = v33
                        _ = v34
                    }
                }
            } else {
                for fi35 := 0; fi35 < 11; fi35++ {
                    var v36 int64 = 1
                    v37 := GenS2{F0: -1, F1: 2, F2: p0}
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                if (float64(-2.25) != -2.25) {
                    p1 = p1
                    p0 = p0
                }
            }
        default:
            for fi38 := 0; fi38 < 15; fi38++ {
                var v39 int64 = 42
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = fi38
                _ = v39
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v20
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 2
    var v3 string = ("\n\t" + "")
    var v4 uint8 = 2
    v5 := make(map[string]uint32)
    v40 := []int{0, -2147483648, 42, -751}
    delete(v1, 42)
    var v41 bool = false
    switch "hello" {
    case "a":
        v42 := GenS4{F0: 4294967295, F1: 981, F2: 0}
        _ = v42
    default:
        delete(v1, 32768)
    }
    v4 -= v4
    v43 := v1[128]
    for fi44 := 0; fi44 < 8; fi44++ {
        switch (fi44 % ((int(v43)) | 1)) {
        case -810:
            v4 = uint8(v43)
        }
        _ = fi44
    }
    v41 = ((!false) || false)
    v2 += (^(v2))
    var v45 int8 = 127
    fmt.Println(v41)
    v4 = (v4 - genFn1(127, (v41 && (!v41)), -1))
    v43 = 121
    for fi46 := 0; fi46 < 10; fi46++ {
        v2 += v2
        _ = fi46
    }
    for fi47 := 0; fi47 < 18; fi47++ {
        v41 = true
        var v48 float64 = -2.25
        _ = fi47
        _ = v48
    }
    fmt.Println(v4, v43)
    switch 128 {
    case 2147483647:
        if ((1 == 2147483647) && v41) {
            fmt.Println(v43, v45)
            for fi49 := 0; fi49 < 19; fi49++ {
                fmt.Println(v4, v2)
                if len(v40) >= 2 {
                    v40 = v40[1:len(v40)]
                }
                _ = fi49
            }
        }
    case 1:
        v4 = v4
    case 32767:
        v50 := GenS1{F0: -947, F1: 0, F2: v41}
        _ = v50
    default:
        fmt.Println(v41)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v45
    _ = v5
}
