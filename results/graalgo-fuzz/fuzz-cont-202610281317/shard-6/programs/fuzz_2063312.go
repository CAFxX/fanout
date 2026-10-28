package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 string }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 string; F1 bool; F2 uint8 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 uint; F1 int64; F2 int16 }
type GenS9 struct { F0 uint32 }
func genFn0() bool {
    for fi4 := 0; fi4 < 16; fi4++ {
        for fi5 := 0; fi5 < 8; fi5++ {
            if false {
                if (32767 < 427) {
                    if true {
                        fmt.Println(fi5, fi4)
                        var v6 bool = false
                        if v6 {
                            v7 := GenS1{F0: fi5, F1: false, F2: "a"}
                            _ = v7
                        }
                        v6 = ((0.1 - 3.14159) >= 1.5)
                        v8 := GenS2{F0: 42, F1: 255}
                        v6 = false
                        var v9 float32 = (0.1 + (float32(float64((((float32(1.5) + 1e10) - (0.1 * (3.14159 - 0.0))) / 1e10))) + 0.0))
                        v9 = -2.25
                        _ = v6
                        _ = v8
                        _ = v9
                    }
                    var v10 int32 = 42
                    v10 -= (v10 / ((v10) | 1))
                    _ = v10
                } else {
                    v11 := []int{fi4}
                    _ = v11
                }
            } else {
                v12 := GenS3{F0: 198, F1: ""}
                var v13 uint8 = 255
                v14 := GenS2{F0: 169, F1: 2}
                v15 := []int{128, fi5, fi5}
                var v16 uint32 = 807
                if (42 <= 4294967295) {
                    if ("foo bar" > "\n\t") {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(fi4, fi5)
                        default:
                            var v17 string = ""
                            _ = v17
                        }
                        _ = len(v15)
                        switch ("a" + "foo bar") {
                        case "":
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        case "foo bar":
                            v18 := GenS3{F0: uint8(v16), F1: "a"}
                            _ = v18
                        case "\n\t":
                            v19 := GenS4{F0: -2147483648, F1: v13}
                            _ = v19
                        }
                    }
                    v13 = (v13 / ((128) | 1))
                    fmt.Println(fi4, fi5, v13)
                } else {
                    switch (fi5 - -9223372036854775808) {
                    case -1:
                        _ = len(v15)
                    default:
                        v20 := GenS5{F0: "", F1: false, F2: (^(v13))}
                        var v21 int64 = 1
                        fmt.Println(fi4, v21)
                        var v22 uint64 = 18446744073709551615
                        if (!false) {
                            var v23 bool = false
                            if v23 {
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                fmt.Println(fi5, v22)
                            }
                            v16 = 495
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            _ = v23
                        }
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    var v24 int64 = 1
                    var v25 uint8 = (v13 % (((v13 - 0)) | 1))
                    _ = v24
                    _ = v25
                }
                _ = v12
                _ = v13
                _ = v14
                _ = v15
                _ = v16
            }
            fmt.Println(fi4)
            v26 := []int{(-(fi5)), (-(fi5))}
            _ = fi5
            _ = v26
        }
        v27 := []int{(fi4 % ((fi4) | 1)), (fi4 % ((fi4) | 1)), -9223372036854775808}
        _ = fi4
        _ = v27
    }
    var v28 int8 = 0
    v28 = (v28 << ((148) & 63))
    for fi29 := 0; fi29 < 5; fi29++ {
        v28 += -1
        _ = fi29
    }
    switch -2147483648 {
    case 32768:
        var v30 uint = 0
        _ = v30
    case 744:
        v28 -= v28
    }
    _ = v28
    return false
}
func genFn1() bool {
    v31 := []int{}
    var v32 int32 = 42
    var v33 int16 = 882
    _ = v31
    _ = v32
    _ = v33
    return false
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "\n\t"
    var v3 int8 = 42
    v2 = "a"
    var v34 int16 = -584
    v2 = v2
    if true {
        delete(v1, 32767)
    }
    var v35 int32 = -1
    for fi36 := 0; fi36 < 3; fi36++ {
        v37 := []int{}
        v38 := GenS6{F0: (v2 + (v2 + v2)), F1: 128}
        _ = fi36
        _ = v37
        _ = v38
    }
    if true {
        fmt.Println(v35, v2)
        switch "hello" {
        case "":
            v3 += (v3 << ((500) & 63))
        }
    } else {
        fmt.Println(v2, v2, v3)
        var v39 uint64 = 18446744073709551615
        v35 -= v35
        _ = v39
    }
    v3 = (v3 & 42)
    v2 = (v2 + v2)
    var v40 uint16 = 32767
    fmt.Println(v40, v34, v34)
    for fi41 := 0; fi41 < 4; fi41++ {
        fmt.Println(v2, v40)
        v1[(fi41 & fi41)] = 4294967295
        for fi42 := 0; fi42 < 1; fi42++ {
            switch fi41 {
            case 9223372036854775807:
                v34 -= (v34 | v34)
            case 32767:
                v43 := GenS7{F0: 1, F1: 42}
                _ = v43
            }
            switch "foo bar" {
            case "\n\t":
                v35 = 0
            case "hello":
                v34 += 42
            default:
                v44 := GenS3{F0: 242, F1: "hello"}
                _ = v44
            }
            switch -1 {
            case 9223372036854775807:
                v2 = v2
            case 32767:
                v40 += v40
            case -9223372036854775808:
                v45 := GenS8{F0: 18446744073709551615, F1: -9223372036854775808, F2: v34}
                _ = v45
            default:
                v46 := []int{(fi42 / (((fi42 << ((42) & 63))) | 1)), (fi41 / ((fi41) | 1)), fi41, (fi41 | -272)}
                _ = v46
            }
            var v47 string = (v2 + v2)
            v48 := GenS9{F0: 2147483648}
            _ = fi42
            _ = v47
            _ = v48
        }
        _ = fi41
    }
    v49 := GenS6{F0: v2, F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v40
    _ = v49
}
