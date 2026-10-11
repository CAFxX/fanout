package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int; F1 int16; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 uint8; F1 int16; F2 uint64 }
func genFn0(p0 uint64) int64 {
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 uint32 = 1
            if (!(65535 < 697)) {
                v7 := GenS1{F0: 0, F1: 127, F2: (((65535 <= 65535) && (fi4 != (fi4 % (((^(fi5))) | 1)))) || true)}
                if (!(!true)) {
                    v6 += 2
                    v8 := []int{204, 9223372036854775807}
                    v8 = append(v8, fi5)
                    _ = v8
                } else {
                    v9 := []int{}
                    v10 := GenS2{F0: (p0 | (p0 << ((4294967295) & 63))), F1: 48, F2: 610}
                    p0 = p0
                    _ = v10
                    _ = v9
                }
                _ = v7
            } else {
                v11 := GenS3{F0: fi4, F1: 1, F2: 18446744073709551615}
                _ = v11
            }
            var v12 int = (-(fi5))
            _ = fi5
            _ = v12
            _ = v6
        }
        _ = fi4
    }
    if true {
        p0 = (p0 % (((p0 * p0)) | 1))
        p0 = p0
        v13 := GenS4{F0: 42}
        _ = v13
    }
    switch 445 {
    case 128:
        if (!true) {
            v14 := GenS3{F0: -472, F1: 732, F2: 1}
            p0 = 1
            _ = v14
        } else {
            var v15 int64 = 9223372036854775807
            v16 := []int{}
            for fi17 := 0; fi17 < 18; fi17++ {
                fmt.Println(v15)
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = fi17
            }
            _ = v15
            _ = v16
        }
    case 6:
        if (2147483647 >= 0) {
            var v18 uint32 = 98
            _ = v18
        }
    case -1:
        switch "foo bar" {
        case "a":
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 uint = 18446744073709551615
                for fi21 := 0; fi21 < 9; fi21++ {
                    switch ("\n\t" + ("" + ("\n\t" + "hello"))) {
                    case "hello":
                        v22 := []int{}
                        _ = v22
                    case "a":
                        v20 -= v20
                        var v23 uint16 = 65535
                        _ = v23
                    case "foo bar":
                        switch ("hello" + "a") {
                        case "a":
                            v20 = 4294967295
                        default:
                            if (0.1 == 0.0) {
                                switch (fi19 / ((fi21) | 1)) {
                                case 4294967295:
                                    p0 += (p0 ^ (^(p0)))
                                case -1:
                                    v24 := []int{(fi21 / ((fi19) | 1)), (fi19 + (fi19 & 2147483647))}
                                    _ = v24
                                case -9223372036854775808:
                                    v20 += (v20 - v20)
                                default:
                                    p0 = p0
                                }
                                switch (fi19 | fi21) {
                                case 4294967295:
                                    fmt.Println(p0, fi19, fi21)
                                }
                                p0 -= (p0 + p0)
                                fmt.Println(v20)
                            }
                        }
                    }
                    if false {
                        if (0 == 1000) {
                            v25 := GenS5{F0: (^(v20)), F1: 255, F2: 4294967295}
                            fmt.Println(v20, fi19, p0)
                            switch (fi19 & (fi21 ^ -615)) {
                            case 42:
                                var v26 bool = (!true)
                                _ = v26
                            case -142:
                                v27 := GenS3{F0: 2147483647, F1: int16(fi19), F2: 65535}
                                _ = v27
                            default:
                                v28 := GenS3{F0: (fi21 % ((fi21) | 1)), F1: -292, F2: (v20 & 42)}
                                _ = v28
                            }
                            _ = v25
                        } else {
                            v29 := []int{(fi21 | 4294967295), (^(fi19)), fi21, (-(fi21))}
                            if true {
                                var v30 int8 = -128
                                _ = v30
                            }
                            _ = v29
                        }
                        switch "a" {
                        case "foo bar":
                            var v31 int64 = -90
                            fmt.Println(fi21, fi21, fi19)
                            var v32 uint32 = uint32(v20)
                            _ = v31
                            _ = v32
                        case "a":
                            v33 := GenS5{F0: (^(v20)), F1: 127, F2: 42}
                            _ = v33
                        case "":
                            p0 += 2
                        default:
                            if ((p0 == 420) && (!((("\n\t" + (("\n\t" + "a") + (("" + "") + ("\n\t" + "hello")))) + "a") != (("foo bar" + "foo bar") + "a")))) {
                                fmt.Println(fi21, fi21, p0)
                                var v34 int8 = 0
                                var v35 float32 = 1.5
                                _ = v34
                                _ = v35
                            } else {
                                v36 := []int{fi21, fi19, 4294967295, (fi21 % (((fi19 + (fi19 + fi19))) | 1))}
                                _ = v36
                            }
                        }
                    } else {
                        p0 = 428
                    }
                    p0 = p0
                    p0 = p0
                    _ = fi21
                }
                _ = fi19
                _ = v20
            }
        default:
            if false {
                for fi37 := 0; fi37 < 7; fi37++ {
                    v38 := GenS6{F0: fi37, F1: p0}
                    _ = fi37
                    _ = v38
                }
                var v39 int64 = int64(p0)
                p0 += 633
                _ = v39
            } else {
                var v40 int8 = -128
                var v41 float32 = 1.5
                p0 = 2
                _ = v40
                _ = v41
            }
        }
    }
    _ = p0
    return 1
}
func genFn1(p0 int64, p1 int) string {
    for fi42 := 0; fi42 < 1; fi42++ {
        var v43 bool = false
        _ = fi42
        _ = v43
    }
    p0 = (p0 << ((4294967295) & 63))
    if true {
        p0 = (p0 / (((p0 | (p0 ^ 9223372036854775807))) | 1))
        v44 := GenS1{F0: 0, F1: 100, F2: (!false)}
        _ = v44
    }
    _ = p0
    _ = p1
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 1
    var v3 uint16 = 255
    v45 := GenS3{F0: -1, F1: 32767, F2: 2}
    var v46 uint64 = 1
    v47 := []int{1}
    switch "\n\t" {
    case "":
        v2 = -837
    default:
        for fi48 := 0; fi48 < 9; fi48++ {
            var v49 float64 = ((0.1 + 1e10) - 1e10)
            v46 = (^(v46))
            _ = fi48
            _ = v49
        }
    }
    v2 = (v2 << ((2) & 63))
    v3 = v3
    fmt.Println(v3)
    for fi50 := 0; fi50 < 8; fi50++ {
        if false {
            v51 := GenS7{F0: 121, F1: 42, F2: v46}
            _ = v51
        }
        _ = fi50
    }
    var v52 uint64 = (v46 + (v46 / (((^(v46))) | 1)))
    switch 833 {
    case 424:
        v46 = v52
    case 545:
        if (!false) {
            v46 = (^(v52))
            switch -181 {
            case -2147483648:
                v2 = (v2 / ((v2) | 1))
            default:
                v3 = (^(v3))
            }
        }
    default:
        v46 = (v46 << ((18446744073709551615) & 63))
    }
    v2 = 0
    v53 := GenS4{F0: 42}
    v2 = v2
    v1[659] = true
    if len(v47) > 0 {
        v47[0] = 4294967295
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v47
    _ = v52
    _ = v53
}
