package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int8; F1 uint; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 uint64; F1 int }
func genFn0(p0 int8, p1 bool) uint16 {
    var v5 int16 = 257
    for fi6 := 0; fi6 < 2; fi6++ {
        if (p1 && p1) {
            switch "foo bar" {
            case "":
                for fi7 := 0; fi7 < 7; fi7++ {
                    if false {
                        v8 := []int{(fi7 - (fi6 & 392)), 177, (fi6 - (fi6 << ((4294967295) & 63))), (fi6 % ((-121) | 1))}
                        _ = len(v8)
                        var v9 int32 = 647
                        var v10 float32 = 3.14159
                        if (-9223372036854775808 != 42) {
                            v11 := GenS1{F0: 2, F1: p0}
                            var v12 bool = p1
                            var v13 int32 = int32(fi6)
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        _ = v10
                        _ = v8
                        _ = v9
                    } else {
                        if ((!(-2.25 < 1e10)) || p1) {
                            v14 := []int{}
                            _ = v14
                        } else {
                            v15 := GenS1{F0: 2, F1: (p0 + -1)}
                            v16 := []int{}
                            var v17 int8 = (p0 % (((p0 * p0)) | 1))
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        v18 := GenS1{F0: 1, F1: p0}
                        v19 := []int{(fi6 / (((fi6 & (fi6 << ((18446744073709551615) & 63)))) | 1)), fi6, fi7}
                        var v20 uint8 = 2
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    _ = fi7
                }
            default:
                var v21 uint8 = 127
                _ = v21
            }
            v22 := GenS2{F0: 65535, F1: (fi6 - fi6)}
            _ = v22
        } else {
            var v23 string = ("foo bar" + "hello")
            v23 = v23
            _ = v23
        }
        p0 += p0
        _ = fi6
    }
    var v24 uint64 = 1
    switch 837 {
    case 791:
        v24 = v24
    case 280:
        if p1 {
            p0 -= p0
            var v25 int32 = -592
            var v26 int8 = 114
            _ = v25
            _ = v26
        }
    default:
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := GenS1{F0: 178, F1: p0}
            _ = fi27
            _ = v28
        }
    }
    p1 = (32767 == 42)
    _ = p0
    _ = p1
    _ = v24
    _ = v5
    return 42
}
func genFn1(p0 uint8, p1 bool, p2 uint8) uint {
    v29 := GenS3{F0: 102, F1: 0, F2: 2}
    if (42 != 0) {
        v30 := GenS1{F0: 42, F1: 0}
        _ = v30
    } else {
        v31 := GenS3{F0: -1, F1: 4294967295, F2: 42}
        p2 += (p2 * (p2 >> ((42) & 63)))
        if p1 {
            if (true || (441 != 0)) {
                v32 := []int{32767}
                if len(v32) >= 2 {
                    v32 = v32[1:len(v32)]
                }
                var v33 uint64 = 2
                _ = v32
                _ = v33
            } else {
                var v34 uint64 = 0
                _ = v34
            }
        }
        _ = v31
    }
    if (!(!((!p1) && false))) {
        v35 := GenS4{F0: -2147483648, F1: 2}
        p1 = p1
        v36 := []int{}
        _ = v35
        _ = v36
    } else {
        var v37 uint32 = 695
        _ = v37
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    return 4294967295
}

func main() {
    v1 := []string{"foo bar", (((("\n\t" + "a") + "hello") + "foo bar") + "foo bar"), (("hello" + "hello") + "hello"), "hello"}
    v2 := make(map[int]int)
    v3 := []int{-836, 127, 2147483647, -66}
    var v4 uint = 65535
    var v38 uint8 = 2
    v4 -= genFn1(v38, (float32((0.0 - -2.25)) > -2.25), (v38 % ((v38) | 1)))
    v3 = append(v3, 42)
    if (-2.25 < float64(0.0)) {
        var v39 uint32 = 0
        _ = v39
    }
    var v40 int = -9223372036854775808
    fmt.Println(v40)
    v4 += 42
    v4 = genFn1(83, false, (v38 | v38))
    v1 = append(v1, ("\n\t" + ""))
    fmt.Println(v38)
    if true {
        var v41 int16 = -261
        if true {
            fmt.Println(v41, v38, v38)
            v41 = v41
            fmt.Println(v40)
        } else {
            switch v40 {
            case -2147483648:
                for fi42 := 0; fi42 < 1; fi42++ {
                    fmt.Println(v41, v40, v4)
                    _ = fi42
                }
            case 128:
                var v43 uint16 = genFn0(-128, false)
                _ = v43
            default:
                if true {
                    for fi44 := 0; fi44 < 2; fi44++ {
                        v41 -= (-(v41))
                        _ = fi44
                    }
                }
            }
        }
        fmt.Println(v38)
        _ = v41
    } else {
        v40 += (v40 << ((v4) & 63))
    }
    for fi45 := 0; fi45 < 14; fi45++ {
        v46 := GenS3{F0: -65, F1: v4, F2: 255}
        for fi47 := 0; fi47 < 13; fi47++ {
            v48 := GenS5{F0: 18446744073709551615, F1: v40}
            _ = fi47
            _ = v48
        }
        v49 := GenS4{F0: -2147483648, F1: genFn0(-1, false)}
        _ = fi45
        _ = v46
        _ = v49
    }
    v40 += v40
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v40
}
