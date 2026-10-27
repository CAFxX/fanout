package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int64; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 int16; F2 bool }
func genFn0(p0 int8) int8 {
    v4 := GenS1{F0: 18446744073709551615, F1: (!true)}
    switch 744 {
    case -761:
        v5 := []int{32768, 32767}
        _ = v5
    case 4294967295:
        if true {
            v6 := GenS1{F0: 1, F1: ("a" > "hello")}
            _ = v6
        } else {
            p0 += p0
            switch 128 {
            case 944:
                p0 = (p0 + (p0 % ((p0) | 1)))
            case 2147483647:
                fmt.Println(p0)
            case 42:
                if false {
                    v7 := []int{}
                    var v8 int32 = -2147483648
                    _ = v7
                    _ = v8
                }
            }
            switch -399 {
            case -626:
                if (65535 < 65535) {
                    if false {
                        for fi9 := 0; fi9 < 2; fi9++ {
                            p0 += (p0 / (((p0 % (((p0 - 100)) | 1))) | 1))
                            fmt.Println(p0)
                            _ = fi9
                        }
                        for fi10 := 0; fi10 < 17; fi10++ {
                            p0 = -128
                            for fi11 := 0; fi11 < 19; fi11++ {
                                if (738 != 0) {
                                    if true {
                                        var v12 bool = true
                                        _ = v12
                                    }
                                    switch "" {
                                    case "foo bar":
                                        p0 -= 100
                                    case "a":
                                        p0 = p0
                                    }
                                }
                                _ = fi11
                            }
                            for fi13 := 0; fi13 < 16; fi13++ {
                                v14 := []int{fi13, (fi10 << ((665) & 63))}
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                var v15 string = "\n\t"
                                _ = len(v14)
                                v16 := GenS1{F0: 0, F1: (1 == 1)}
                                _ = len(v14)
                                _ = fi13
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                            _ = fi10
                        }
                        p0 -= p0
                    } else {
                        v17 := GenS1{F0: 4294967295, F1: (!true)}
                        p0 -= -81
                        _ = v17
                    }
                } else {
                    v18 := []int{}
                    var v19 uint = 255
                    _ = v18
                    _ = v19
                }
            default:
                var v20 float32 = 0.0
                _ = v20
            }
        }
    case 128:
        v21 := GenS2{F0: -561, F1: 32767, F2: 42}
        _ = v21
    default:
        p0 = (p0 >> ((18446744073709551615) & 63))
    }
    fmt.Println(p0)
    p0 += (p0 % ((p0) | 1))
    var v22 int32 = 0
    _ = p0
    _ = v22
    _ = v4
    return 61
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 0
    var v3 int32 = 2147483647
    v23 := GenS2{F0: (v2 + -748), F1: -1, F2: 42}
    v24 := GenS2{F0: (v2 + (v2 << ((660) & 63))), F1: int16(v2), F2: 23}
    v25 := GenS3{F0: 0, F1: 0, F2: 4294967295}
    for fi26 := 0; fi26 < 18; fi26++ {
        delete(v1, ("" + "\n\t"))
        v3 = v3
        if true {
            v2 -= (v2 * 9223372036854775807)
        } else {
            fmt.Println(v2, v3)
            fmt.Println(v2, v3, v3)
            v27 := GenS4{F0: 1}
            _ = v27
        }
        _ = fi26
    }
    v28 := GenS5{F0: 0}
    fmt.Println(v3)
    v3 -= -101
    v29 := []int{}
    var v30 int64 = v2
    switch 32768 {
    case -36:
        switch "foo bar" {
        case "":
            var v31 bool = (false && (true && true))
            _ = v31
        case "foo bar":
            var v32 int64 = 1
            _ = v32
        }
    }
    var v33 uint64 = 813
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 int16 = int16(fi34)
        v36 := GenS6{F0: 411, F1: v35, F2: (true && false)}
        if (!(127 != 0)) {
            var v37 int64 = (v30 * (v30 & (v2 << ((4294967295) & 63))))
            var v38 int8 = -1
            _ = v37
            _ = v38
        } else {
            fmt.Println(v33)
        }
        _ = fi34
        _ = v35
        _ = v36
    }
    switch "\n\t" {
    case "foo bar":
        v39 := v1[(("a" + (("a" + "a") + "\n\t")) + "\n\t")]
        _ = v39
    }
    v40 := GenS1{F0: 4294967295, F1: (626 < 65535)}
    var v41 int64 = (^(v30))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v40
    _ = v41
}
