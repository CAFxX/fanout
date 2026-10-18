package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 string }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 int64; F2 uint8 }
type GenS6 struct { F0 int16; F1 int64; F2 int16 }
type GenS7 struct { F0 uint8 }
func genFn0() int64 {
    var v6 int8 = 1
    var v7 int32 = -2147483648
    if false {
        var v8 int16 = -517
        v7 = v7
        _ = v8
    }
    fmt.Println(v6, v6)
    for fi9 := 0; fi9 < 2; fi9++ {
        if (-1 > -9223372036854775808) {
            if false {
                v10 := GenS1{F0: 709, F1: v6, F2: ("foo bar" + "a")}
                fmt.Println(fi9)
                var v11 uint16 = 0
                _ = v10
                _ = v11
            }
        } else {
            switch "hello" {
            case "a":
                for fi12 := 0; fi12 < 16; fi12++ {
                    v7 += (v7 / ((v7) | 1))
                    switch "a" {
                    case "hello":
                        v7 += (^(v7))
                    case "foo bar":
                        if false {
                            v13 := GenS2{F0: 1, F1: fi9}
                            var v14 bool = (32767 > 687)
                            if ((v14 && (v14 && (v14 && (2 <= 255)))) || ("foo bar" == (("foo bar" + "a") + ("" + "foo bar")))) {
                                fmt.Println(fi12)
                                v15 := []int{fi12, (^(fi9))}
                                v14 = (!(!(4294967295 > 0)))
                                _ = v15
                            } else {
                                if (!(0.0 > 0.0)) {
                                    v14 = v14
                                    v6 = (v6 / ((v6) | 1))
                                }
                            }
                            _ = v13
                            _ = v14
                        } else {
                            if ((^(v7)) >= (v7 & 42)) {
                                v7 = (v7 & v7)
                                v16 := GenS1{F0: 244, F1: v6, F2: "a"}
                                fmt.Println(v7, v6, v7)
                                _ = v16
                            } else {
                                if false {
                                    v17 := []int{}
                                    v6 = -111
                                    switch -863 {
                                    case 128:
                                        var v18 int32 = v7
                                        _ = v18
                                    case -1:
                                        var v19 bool = true
                                        var v20 uint32 = 2147483648
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = v17
                                } else {
                                    v6 = v6
                                    v6 += (^(v6))
                                    v21 := GenS3{F0: 81}
                                    v22 := []int{(fi12 << ((65) & 63)), fi9}
                                    _ = v21
                                    _ = v22
                                }
                                fmt.Println(v7, v6)
                            }
                            fmt.Println(v7)
                            v23 := GenS4{F0: 18446744073709551615}
                            _ = v23
                        }
                    default:
                        if true {
                            v24 := []int{(fi12 & 42), (fi9 & (fi12 ^ fi9)), (-(fi9)), fi12}
                            v25 := GenS3{F0: 127}
                            _ = v24
                            _ = v25
                        }
                        var v26 uint32 = 2
                        fmt.Println(fi12, fi12, v6)
                        _ = v26
                    }
                    var v27 float32 = (3.14159 * float32(3.14159))
                    _ = fi12
                    _ = v27
                }
            default:
                v28 := GenS5{F0: v6, F1: 0, F2: 255}
                _ = v28
            }
        }
        v29 := GenS4{F0: 233}
        _ = fi9
        _ = v29
    }
    _ = v6
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint64) int {
    fmt.Println(p0)
    if true {
        var v30 uint8 = 42
        v31 := GenS6{F0: 0, F1: 1, F2: 0}
        if (!(false && false)) {
            var v32 uint64 = (p0 >> ((128) & 63))
            _ = v32
        }
        _ = v30
        _ = v31
    }
    fmt.Println(p0)
    _ = p0
    return 554
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 0
    var v3 bool = (true && (!false))
    var v4 int16 = 42
    v5 := make(map[string]string)
    var v33 int32 = 1
    fmt.Println(v33, v2)
    delete(v1, 42)
    switch 0 {
    case 1:
        switch "\n\t" {
        case "a":
            v34 := GenS3{F0: v2}
            _ = v34
        case "hello":
            v35 := []int{genFn1(18446744073709551615), genFn1(0)}
            _ = v35
        case "":
            if true {
                v4 -= (-(v4))
                var v36 float32 = float32(((0.0 * -2.25) + 0.1))
                if true {
                    var v37 uint = 18446744073709551615
                    fmt.Println(v33)
                    _ = v37
                }
                _ = v36
            } else {
                var v38 int32 = (v33 & (v33 | -2147483648))
                _ = v38
            }
        }
    case 32767:
        var v39 int16 = (^(v4))
        _ = v39
    }
    v40 := GenS7{F0: (v2 / (((v2 << ((42) & 63))) | 1))}
    switch "" {
    case "\n\t":
        v2 = 1
    case "foo bar":
        if (((!(!false)) && ((-1 > 4294967295) || false)) || (9223372036854775807 > 4294967295)) {
            for fi41 := 0; fi41 < 4; fi41++ {
                fmt.Println(v3)
                if ("foo bar" == "\n\t") {
                    var v42 int8 = 127
                    switch (("\n\t" + "hello") + "a") {
                    case "foo bar":
                        if ((true && false) || v3) {
                            for fi43 := 0; fi43 < 19; fi43++ {
                                var v44 bool = (!v3)
                                v2 = (^(v2))
                                if (v4 != (v4 >> ((876) & 63))) {
                                    fmt.Println(v2, fi43, v3)
                                    fmt.Println(fi43)
                                    v45 := []int{fi43, (^(fi43)), fi41}
                                    _ = v45
                                } else {
                                    var v46 uint32 = 42
                                    switch fi43 {
                                    case 988:
                                        var v47 float64 = 0.1
                                        _ = v47
                                    default:
                                        v46 -= v46
                                    }
                                    if false {
                                        v44 = v3
                                        switch fi43 {
                                        case -527:
                                            v3 = false
                                        case 9223372036854775807:
                                            v48 := GenS3{F0: 2}
                                            _ = v48
                                        default:
                                            fmt.Println(fi41)
                                        }
                                        v49 := GenS7{F0: 0}
                                        _ = v49
                                    }
                                    _ = v46
                                }
                                _ = fi43
                                _ = v44
                            }
                            fmt.Println(v42)
                        }
                    }
                    _ = v42
                }
                _ = fi41
            }
            v50 := []int{}
            _ = v50
        }
    case "":
        v2 -= (v2 & v2)
    }
    var v51 string = "a"
    v4 = (v4 * (v4 - (v4 + (v4 % ((v4) | 1)))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v40
    _ = v5
    _ = v51
}
