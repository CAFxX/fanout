package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 bool }
type GenS2 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS3 struct { F0 int; F1 uint8; F2 string }
type GenS4 struct { F0 int32; F1 int16 }
func genFn0() string {
    if false {
        v6 := GenS1{F0: 421, F1: (false || true), F2: true}
        if false {
            var v7 uint64 = 1
            _ = v7
        }
        if (-2147483648 != 0) {
            v8 := []int{127, -9223372036854775808}
            _ = v8
        } else {
            var v9 uint8 = 128
            switch 4294967295 {
            case 652:
                for fi10 := 0; fi10 < 4; fi10++ {
                    if ((2147483648 != 637) || ((880 < 65535) && true)) {
                        for fi11 := 0; fi11 < 9; fi11++ {
                            if (127 < 8) {
                                var v12 int32 = -1
                                fmt.Println(v9, fi10, fi11)
                                var v13 uint = 2
                                _ = v12
                                _ = v13
                            } else {
                                fmt.Println(v9, fi11, fi10)
                                if false {
                                    v9 = v9
                                }
                                v14 := []int{fi11, (fi11 | 9223372036854775807), fi10}
                                _ = v14
                            }
                            _ = fi11
                        }
                    }
                    _ = fi10
                }
            case 4294967295:
                if (!(false && true)) {
                    v9 += v9
                    v9 = 60
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 string = "hello"
                        for fi17 := 0; fi17 < 14; fi17++ {
                            if (!(0 >= 1)) {
                                v16 = "\n\t"
                            } else {
                                var v18 float32 = 0.0
                                _ = v18
                            }
                            switch fi17 {
                            case 4294967295:
                                var v19 uint16 = 65535
                                _ = v19
                            case -2147483648:
                                if true {
                                    switch v16 {
                                    case "hello":
                                        v16 = v16
                                        var v20 uint16 = 2
                                        v21 := []int{532}
                                        _ = v20
                                        _ = v21
                                    case "a":
                                        v22 := []int{}
                                        _ = v22
                                    }
                                    v23 := []int{fi15, -976, fi17}
                                    var v24 int16 = 404
                                    if (!(!true)) {
                                        var v25 int = fi17
                                        v16 = "hello"
                                        var v26 int32 = 2147483647
                                        _ = v25
                                        _ = v26
                                    }
                                    v24 = v24
                                    _ = v23
                                    _ = v24
                                }
                            default:
                                if false {
                                    var v27 uint8 = (^(v9))
                                    v28 := []int{fi15}
                                    fmt.Println(v16, v27, v16)
                                    _ = v27
                                    _ = v28
                                }
                            }
                            _ = fi17
                        }
                        _ = fi15
                        _ = v16
                    }
                } else {
                    var v29 int32 = 2147483647
                    var v30 string = "foo bar"
                    _ = v29
                    _ = v30
                }
            default:
                fmt.Println(v9)
            }
            _ = v9
        }
        _ = v6
    } else {
        if true {
            v31 := GenS2{F0: 18, F1: 563, F2: 65535}
            _ = v31
        }
        switch -423 {
        case 4294967295:
            v32 := []int{434, 880, 4294967295, -649}
            _ = v32
        case 32768:
            if true {
                var v33 int16 = 32767
                var v34 uint8 = 255
                _ = v33
                _ = v34
            }
        }
    }
    var v35 float32 = 1e10
    if (18446744073709551615 == 985) {
        var v36 uint64 = 0
        v37 := []int{-803}
        _ = v36
        _ = v37
    } else {
        var v38 uint32 = 796
        var v39 float32 = 1.5
        fmt.Println(v39, v39, v39)
        _ = v38
        _ = v39
    }
    var v40 int64 = -750
    _ = v35
    _ = v40
    return "a"
}
func genFn1(p0 uint8, p1 uint64) uint8 {
    v41 := GenS1{F0: (p1 % ((4294967295) | 1)), F1: false, F2: ((3.14159 / 1.0) <= 0.0)}
    v42 := GenS1{F0: (p1 ^ p1), F1: true, F2: false}
    for fi43 := 0; fi43 < 17; fi43++ {
        var v44 int32 = 42
        _ = fi43
        _ = v44
    }
    p0 = uint8(p1)
    _ = p0
    _ = p1
    _ = v41
    _ = v42
    return 20
}

func main() {
    var v1 int32 = 42
    v2 := []string{"a"}
    var v3 string = "a"
    var v4 int32 = 2147483647
    var v5 uint32 = 548
    for fi45 := 0; fi45 < 9; fi45++ {
        v4 -= (v1 ^ v1)
        var v46 float64 = 0.1
        _ = fi45
        _ = v46
    }
    var v47 float64 = (3.14159 / 1.0)
    for fi48 := 0; fi48 < 7; fi48++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi48
    }
    switch -1 {
    case 32768:
        var v49 float32 = (0.1 + ((1.5 - 3.14159) + (1.5 / 3.14159)))
        _ = v49
    case 829:
        v50 := GenS3{F0: 127, F1: genFn1(87, 2), F2: ("" + "a")}
        _ = v50
    case -326:
        if (0 <= 0) {
            v4 = v1
            var v51 int32 = (v1 / ((v1) | 1))
            if (((65535 == 18446744073709551615) && (0 != 65535)) || true) {
                for fi52 := 0; fi52 < 2; fi52++ {
                    v47 += float64(1.5)
                    _ = fi52
                }
                var v53 float32 = (3.14159 + float32(v47))
                _ = v53
            } else {
                v1 = 0
                _ = len(v2)
                var v54 string = v3
                _ = v54
            }
            _ = v51
        } else {
            v4 += 2147483647
            v3 = "\n\t"
            var v55 float32 = 0.1
            _ = v55
        }
    }
    if (127 > 42) {
        fmt.Println(v1, v4, v1)
        v56 := GenS4{F0: -2147483648, F1: 32767}
        _ = v56
    }
    _ = len(v2)
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v5 = v5
        var v57 uint8 = 22
        v5 -= 0
        _ = v57
    }
    fmt.Println(v47, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
