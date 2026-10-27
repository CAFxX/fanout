package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint8; F1 int; F2 string }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 uint; F1 bool; F2 int16 }
func genFn0(p0 int16) int64 {
    if (p0 > -32768) {
        for fi4 := 0; fi4 < 14; fi4++ {
            if false {
                if (!((!("hello" >= (("hello" + ("\n\t" + "foo bar")) + "\n\t"))) || (!true))) {
                    v5 := []int{127, fi4, fi4}
                    _ = v5
                }
            } else {
                switch fi4 {
                case -1:
                    fmt.Println(p0, p0)
                case 946:
                    switch 32767 {
                    case 128:
                        v6 := GenS1{F0: 0, F1: true}
                        _ = v6
                    case 236:
                        var v7 uint = 112
                        _ = v7
                    case 1:
                        v8 := []int{-9223372036854775808, -605}
                        _ = v8
                    }
                }
                var v9 uint8 = 2
                for fi10 := 0; fi10 < 7; fi10++ {
                    if true {
                        v11 := GenS1{F0: 42, F1: ((0 <= 1) || true)}
                        if false {
                            var v12 float32 = 1e10
                            _ = v12
                        }
                        fmt.Println(p0, fi4)
                        _ = v11
                    }
                    var v13 uint = 0
                    _ = fi10
                    _ = v13
                }
                _ = v9
            }
            p0 += -1
            _ = fi4
        }
        var v14 int16 = (p0 >> ((18446744073709551615) & 63))
        v15 := []int{-9223372036854775808, 4294967295}
        _ = v14
        _ = v15
    }
    p0 = p0
    for fi16 := 0; fi16 < 10; fi16++ {
        fmt.Println(fi16, p0)
        for fi17 := 0; fi17 < 6; fi17++ {
            switch "a" {
            case "hello":
                p0 += p0
            case "\n\t":
                var v18 bool = false
                _ = v18
            default:
                v19 := GenS1{F0: 0, F1: false}
                _ = v19
            }
            v20 := GenS2{F0: false, F1: fi17}
            var v21 int64 = 9223372036854775807
            _ = fi17
            _ = v20
            _ = v21
        }
        _ = fi16
    }
    switch -9223372036854775808 {
    case 42:
        var v22 uint16 = 76
        _ = v22
    case 127:
        for fi23 := 0; fi23 < 4; fi23++ {
            if false {
                p0 = -32768
                switch (-(fi23)) {
                case 127:
                    p0 -= p0
                case 32767:
                    v24 := GenS1{F0: 618, F1: true}
                    _ = v24
                case -9223372036854775808:
                    v25 := []int{}
                    _ = v25
                }
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 int16 = (p0 + p0)
                    var v28 bool = (-861 == 2147483647)
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            } else {
                if false {
                    v29 := []int{}
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 14; fi30++ {
                        var v31 int16 = (p0 - -32768)
                        if (0 >= 42) {
                            switch ("foo bar" + "a") {
                            case "foo bar":
                                switch ("\n\t" + "") {
                                case "a":
                                    var v32 uint16 = 2
                                    _ = v32
                                case "\n\t":
                                    fmt.Println(fi23, fi30)
                                case "foo bar":
                                    p0 = (^(v31))
                                }
                            case "hello":
                                var v33 int16 = v31
                                _ = v33
                            default:
                                fmt.Println(p0)
                            }
                            var v34 int8 = -128
                            var v35 int32 = -1
                            _ = v34
                            _ = v35
                        } else {
                            var v36 float64 = -2.25
                            v37 := GenS3{F0: "foo bar", F1: 2147483647}
                            fmt.Println(v36)
                            var v38 uint32 = 65535
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                        v31 += (v31 * (p0 + v31))
                        _ = fi30
                        _ = v31
                    }
                }
                for fi39 := 0; fi39 < 12; fi39++ {
                    var v40 float32 = -2.25
                    switch -2147483648 {
                    case 1:
                        var v41 bool = false
                        _ = v41
                    }
                    v42 := GenS4{F0: 1, F1: fi23, F2: "foo bar"}
                    _ = fi39
                    _ = v40
                    _ = v42
                }
            }
            _ = fi23
        }
    case 0:
        var v43 uint16 = 1
        _ = v43
    default:
        p0 -= (^(p0))
    }
    _ = p0
    return 644
}
func genFn1(p0 uint16, p1 uint) int {
    if false {
        var v44 int64 = -835
        var v45 uint16 = (^(p0))
        fmt.Println(v44)
        _ = v44
        _ = v45
    } else {
        p0 -= 1
    }
    for fi46 := 0; fi46 < 20; fi46++ {
        p0 -= 1
        _ = fi46
    }
    p0 = (p0 >> (((p1 & p1)) & 63))
    for fi47 := 0; fi47 < 5; fi47++ {
        var v48 float32 = -2.25
        var v49 uint64 = 48
        p1 -= p1
        _ = fi47
        _ = v48
        _ = v49
    }
    v50 := GenS5{F0: 65535, F1: (p1 << (((p1 + (p1 ^ p1))) & 63))}
    _ = p0
    _ = p1
    _ = v50
    return 128
}

func main() {
    v1 := []uint32{787, 0}
    var v2 int = -2147483648
    v3 := []int{(v2 + (v2 % ((32768) | 1)))}
    switch -438 {
    case 127:
        var v51 int32 = -2147483648
        _ = v51
    default:
        _ = len(v1)
    }
    v2 += v2
    v52 := GenS6{F0: 1, F1: (!true), F2: -411}
    v2 = (v2 >> ((255) & 63))
    switch v2 {
    case 42:
        v1 = append(v1, 0)
    case 4294967295:
        v2 = 32768
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v2)
    if len(v1) > 0 {
        v1[0] = 2
    }
    if (651 < 1) {
        var v53 uint8 = 1
        _ = v53
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
}
