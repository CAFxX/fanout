package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 string; F2 uint8 }
type GenS6 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS7 struct { F0 int64; F1 int32; F2 int64 }
func genFn0(p0 uint, p1 int16) int {
    p0 -= p0
    v4 := GenS1{F0: 2}
    switch 540 {
    case 0:
        var v5 int = 9223372036854775807
        _ = v5
    case -142:
        if false {
            p1 = (p1 + p1)
        }
    default:
        fmt.Println(p0)
    }
    v6 := GenS1{F0: 255}
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 0
}
func genFn1() bool {
    var v7 bool = (true && true)
    v7 = v7
    fmt.Println(v7)
    _ = v7
    return (true && v7)
}

func main() {
    var v1 int32 = 473
    var v2 uint64 = 1
    var v3 bool = ((0.1 < float32(1.5)) && false)
    if false {
        if v3 {
            switch 1 {
            case 1:
                var v8 int64 = -827
                _ = v8
            case -1:
                switch ("\n\t" + "\n\t") {
                case "a":
                    for fi9 := 0; fi9 < 13; fi9++ {
                        v3 = (v3 || (("foo bar" != ("" + "\n\t")) || v3))
                        v10 := GenS2{F0: 2147483647}
                        switch "foo bar" {
                        case "foo bar":
                            var v11 uint16 = 288
                            _ = v11
                        case "":
                            fmt.Println(v3)
                        default:
                            for fi12 := 0; fi12 < 8; fi12++ {
                                switch genFn0(255, 0) {
                                case -663:
                                    var v13 uint32 = 62
                                    _ = v13
                                }
                                var v14 uint32 = uint32(v2)
                                _ = fi12
                                _ = v14
                            }
                        }
                        _ = fi9
                        _ = v10
                    }
                case "foo bar":
                    var v15 float32 = (((1e10 - ((((0.1 + (0.1 / 1.5)) / 1.0) * 1e10) / 1.0)) / 1e10) * (3.14159 + (1e10 * 1.5)))
                    _ = v15
                case "\n\t":
                    if (0.0 < -2.25) {
                        v16 := GenS3{F0: 32767, F1: 100}
                        v17 := []int{}
                        _ = v16
                        _ = v17
                    } else {
                        v18 := []int{genFn0(470, -32768), -9223372036854775808, 4294967295, -547}
                        switch "hello" {
                        case "a":
                            _ = len(v18)
                        case "hello":
                            if false {
                                for fi19 := 0; fi19 < 11; fi19++ {
                                    if v3 {
                                        if (true && v3) {
                                            _ = len(v18)
                                        }
                                    }
                                    var v20 int64 = -146
                                    _ = fi19
                                    _ = v20
                                }
                                v1 = 818
                                _ = len(v18)
                            }
                        }
                        _ = v18
                    }
                default:
                    var v21 int64 = 144
                    _ = v21
                }
            }
        } else {
            v22 := []int{genFn0(42, 42), 128, -284}
            v22 = append(v22, genFn0(uint(v2), 0))
            _ = v22
        }
        switch genFn0(255, 32767) {
        case 9223372036854775807:
            var v23 float32 = (1e10 * 0.1)
            _ = v23
        case -551:
            if (42 >= 179) {
                v3 = false
                var v24 uint32 = 4294967295
                _ = v24
            } else {
                v25 := GenS4{F0: -1}
                v2 += v2
                _ = v25
            }
        }
    } else {
        v3 = v3
        if v3 {
            v1 -= v1
            if v3 {
                var v26 int = int(v1)
                _ = v26
            } else {
                var v27 uint = 0
                if (-270 >= 42) {
                    v27 += (v27 * 468)
                    switch -244 {
                    case 2147483647:
                        v27 += v27
                    case 4294967295:
                        var v28 uint64 = (^(v2))
                        _ = v28
                    default:
                        v1 = 1
                    }
                    v29 := GenS4{F0: 32767}
                    _ = v29
                }
                for fi30 := 0; fi30 < 7; fi30++ {
                    for fi31 := 0; fi31 < 20; fi31++ {
                        var v32 uint64 = 979
                        _ = fi31
                        _ = v32
                    }
                    v1 = 357
                    v33 := GenS3{F0: 1, F1: 100}
                    _ = fi30
                    _ = v33
                }
                _ = v27
            }
            v34 := GenS5{F0: "a", F1: ("a" + "foo bar"), F2: 255}
            _ = v34
        } else {
            v35 := GenS4{F0: 72}
            for fi36 := 0; fi36 < 20; fi36++ {
                v2 = (^(v2))
                for fi37 := 0; fi37 < 7; fi37++ {
                    switch ("\n\t" + "") {
                    case "foo bar":
                        fmt.Println(fi37, fi37)
                    }
                    _ = fi37
                }
                _ = fi36
            }
            for fi38 := 0; fi38 < 2; fi38++ {
                switch (fi38 >> ((18446744073709551615) & 63)) {
                case 127:
                    v39 := GenS6{F0: 32767, F1: v1, F2: 255}
                    _ = v39
                default:
                    for fi40 := 0; fi40 < 16; fi40++ {
                        var v41 uint32 = 65535
                        v3 = (v3 || true)
                        v1 = v1
                        v42 := GenS7{F0: 0, F1: v1, F2: 42}
                        var v43 int = genFn0(2, 1)
                        _ = fi40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                }
                _ = fi38
            }
            _ = v35
        }
        if true {
            for fi44 := 0; fi44 < 13; fi44++ {
                if v3 {
                    switch ((("" + "foo bar") + "hello") + "") {
                    case "foo bar":
                        v45 := []int{genFn0(0, 1000)}
                        _ = v45
                    default:
                        var v46 int64 = 880
                        _ = v46
                    }
                    v3 = (!v3)
                } else {
                    v1 -= v1
                }
                for fi47 := 0; fi47 < 3; fi47++ {
                    v48 := []int{(fi44 >> ((4294967295) & 63)), (^(fi47))}
                    fmt.Println(v3)
                    var v49 bool = v3
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                v50 := GenS1{F0: 2}
                _ = fi44
                _ = v50
            }
        }
    }
    v51 := []int{9223372036854775807, 4294967295}
    v52 := GenS3{F0: 65535, F1: -1}
    fmt.Println(v3)
    var v53 int32 = v1
    var v54 int8 = 100
    v53 -= v53
    var v55 uint32 = 89
    for fi56 := 0; fi56 < 15; fi56++ {
        v53 = int32(fi56)
        _ = fi56
    }
    v51 = append(v51, int(v55))
    switch -675 {
    case -461:
        for fi57 := 0; fi57 < 4; fi57++ {
            for fi58 := 0; fi58 < 17; fi58++ {
                var v59 int8 = (-(v54))
                _ = len(v51)
                _ = fi58
                _ = v59
            }
            v60 := GenS5{F0: "", F1: "foo bar", F2: 0}
            var v61 uint16 = 255
            _ = fi57
            _ = v60
            _ = v61
        }
    }
    if len(v51) >= 2 {
        v51 = v51[1:len(v51)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
}
