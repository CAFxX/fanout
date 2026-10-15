package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint64 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 uint32; F1 int }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 4; fi4++ {
        if true {
            if true {
                v5 := []int{fi4, fi4, 193}
                v6 := GenS1{F0: 1, F1: ("foo bar" < ("hello" + "a")), F2: 1}
                var v7 int8 = -36
                _ = v5
                _ = v6
                _ = v7
            } else {
                fmt.Println(fi4)
            }
        } else {
            var v8 uint32 = 458
            if false {
                var v9 int8 = -115
                _ = v9
            } else {
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 int16 = -1
                    _ = fi10
                    _ = v11
                }
            }
            _ = v8
        }
        _ = fi4
    }
    v12 := GenS1{F0: -962, F1: false, F2: 445}
    v13 := GenS1{F0: 338, F1: (127 <= 4294967295), F2: 272}
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 bool = (42 > 42)
    if (42 <= 2) {
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int8 = 127
            for fi16 := 0; fi16 < 17; fi16++ {
                v15 -= (v15 - v15)
                v3 = false
                if v3 {
                    if v3 {
                        v17 := GenS2{F0: (^(fi14)), F1: -9223372036854775808}
                        fmt.Println(fi16)
                        _ = v17
                    } else {
                        v18 := GenS3{F0: 558, F1: -43, F2: 318}
                        v15 -= v15
                        switch (fi16 % ((fi14) | 1)) {
                        case 42:
                            var v19 int8 = (^(v15))
                            _ = v19
                        case -914:
                            if v3 {
                                switch fi14 {
                                case 728:
                                    v20 := GenS4{F0: genFn0()}
                                    _ = v20
                                default:
                                    v21 := GenS4{F0: 656}
                                    _ = v21
                                }
                                var v22 int8 = (v15 >> ((0) & 63))
                                _ = v22
                            } else {
                                if true {
                                    var v23 uint16 = 415
                                    _ = v23
                                }
                                v15 -= 61
                            }
                        }
                        _ = v18
                    }
                    var v24 float64 = ((1e10 * 1e10) / 1.0)
                    var v25 uint64 = 2
                    _ = v24
                    _ = v25
                }
                _ = fi16
            }
            switch "a" {
            case "a":
                if (2 <= 4294967295) {
                    fmt.Println(v3)
                }
            case "foo bar":
                var v26 uint64 = 0
                _ = v26
            case "hello":
                v15 = (-(v15))
            default:
                if (776 <= 0) {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v28 := GenS5{F0: 2147483647, F1: fi14}
                        _ = fi27
                        _ = v28
                    }
                } else {
                    v29 := GenS6{F0: 2147483648, F1: int(v15)}
                    _ = v29
                }
            }
            _ = fi14
            _ = v15
        }
    }
    var v30 int64 = 986
    fmt.Println(v30, v30)
    v31 := []int{-9223372036854775808}
    if false {
        if len(v31) >= 2 {
            v31 = v31[1:len(v31)]
        }
        v31 = append(v31, 32768)
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        if ((-27 == -1) && (v3 || (42 < 127))) {
            _ = len(v31)
            var v33 bool = v3
            _ = v33
        }
        _ = fi32
    }
    v30 -= v30
    var v34 bool = true
    var v35 int16 = -49
    var v36 uint8 = 128
    switch "hello" {
    case "hello":
        fmt.Println(v30, v35, v35)
    case "a":
        delete(v2, 42)
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        if false {
            v34 = (v30 > v30)
        } else {
            var v38 uint8 = 63
            _ = v38
        }
        _ = fi37
    }
    var v39 int32 = 245
    var v40 uint8 = v36
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v40
}
