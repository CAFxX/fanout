package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 int64, p1 uint64, p2 int) uint8 {
    var v4 int64 = (p0 >> ((1) & 63))
    switch ("" + ("\n\t" + ("\n\t" + (((("a" + "") + (("hello" + "foo bar") + ("a" + "a"))) + "a") + "hello")))) {
    case "a":
        var v5 int32 = 0
        _ = v5
    case "":
        var v6 uint16 = 931
        _ = v6
    case "hello":
        v7 := GenS1{F0: 65535}
        _ = v7
    default:
        switch "a" {
        case "a":
            for fi8 := 0; fi8 < 20; fi8++ {
                for fi9 := 0; fi9 < 17; fi9++ {
                    v10 := []int{}
                    var v11 uint16 = 32767
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                v12 := GenS1{F0: 560}
                v13 := []int{(fi8 ^ (fi8 & (p2 << ((40) & 63))))}
                _ = fi8
                _ = v12
                _ = v13
            }
        }
    }
    var v14 int = 9223372036854775807
    v15 := []int{(-(p2)), v14, (p2 * v14), (p2 - p2)}
    var v16 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    _ = v4
    return 39
}

func main() {
    v1 := []int{-2147483648}
    v2 := []string{"hello"}
    var v3 int8 = 127
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 float32 = 1e10
        fmt.Println(v18, v18)
        _ = fi17
        _ = v18
    }
    var v19 bool = false
    v20 := GenS1{F0: 65535}
    if len(v2) > 0 {
        v2[0] = ("\n\t" + "a")
    }
    fmt.Println(v19)
    var v21 uint8 = 42
    var v22 int = 497
    for fi23 := 0; fi23 < 11; fi23++ {
        _ = len(v1)
        _ = fi23
    }
    switch "hello" {
    case "foo bar":
        if (-832 == 969) {
            if v19 {
                v24 := GenS1{F0: 642}
                _ = v24
            } else {
                if v19 {
                    var v25 int = -2147483648
                    for fi26 := 0; fi26 < 11; fi26++ {
                        for fi27 := 0; fi27 < 4; fi27++ {
                            if v19 {
                                if false {
                                    switch (v22 + fi26) {
                                    case 765:
                                        var v28 int16 = 1000
                                        _ = v28
                                    default:
                                        var v29 int8 = v3
                                        _ = v29
                                    }
                                    v30 := GenS2{F0: 2147483647, F1: (v22 - (-(v25)))}
                                    _ = v30
                                }
                                if (!false) {
                                    v31 := GenS3{F0: (v25 % (((fi27 + fi27)) | 1)), F1: genFn0(-9223372036854775808, 616, v22)}
                                    if ((-2.25 / 3.14159) <= 0.1) {
                                        if true {
                                            if false {
                                                var v32 float64 = ((-2.25 + 3.14159) - float64(1e10))
                                                _ = len(v2)
                                                v21 += (v21 & v21)
                                                var v33 uint64 = 42
                                                switch "hello" {
                                                case "":
                                                    var v34 string = "\n\t"
                                                    _ = v34
                                                default:
                                                    var v35 int64 = -1
                                                    _ = v35
                                                }
                                                _ = v32
                                                _ = v33
                                            }
                                            var v36 uint32 = 1
                                            v21 += v21
                                            _ = v36
                                        } else {
                                            var v37 uint8 = v21
                                            _ = v37
                                        }
                                    }
                                    switch ("foo bar" + "a") {
                                    case "a":
                                        var v38 int8 = 1
                                        if len(v1) > 0 {
                                            v1[0] = 0
                                        }
                                        _ = v38
                                    default:
                                        v19 = (v21 <= (v21 << ((4294967295) & 63)))
                                    }
                                    _ = v31
                                }
                                v39 := GenS1{F0: 255}
                                fmt.Println(v21, v19)
                                _ = v39
                            } else {
                                var v40 int8 = v3
                                _ = v40
                            }
                            _ = fi27
                        }
                        _ = fi26
                    }
                    _ = len(v2)
                    _ = v25
                } else {
                    for fi41 := 0; fi41 < 5; fi41++ {
                        _ = len(v1)
                        _ = fi41
                    }
                }
            }
        }
    case "":
        for fi42 := 0; fi42 < 19; fi42++ {
            var v43 int64 = 601
            var v44 uint32 = 2
            _ = fi42
            _ = v43
            _ = v44
        }
    }
    v22 = 581
    if v19 {
        v3 += -128
        v2 = append(v2, "")
    } else {
        v45 := GenS3{F0: v22, F1: (v21 & genFn0(791, 2, 32768))}
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
