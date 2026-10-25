package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 string }
type GenS2 struct { F0 int64; F1 int; F2 int32 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 uint16; F2 string }
type GenS7 struct { F0 uint8 }
func genFn0() int64 {
    if true {
        var v5 int32 = 163
        fmt.Println(v5)
        switch "foo bar" {
        case "\n\t":
            v6 := GenS1{F0: 2, F1: (-(v5)), F2: ("hello" + "")}
            _ = v6
        case "foo bar":
            v5 = (v5 * (^(v5)))
        default:
            v7 := GenS1{F0: 731, F1: (^(v5)), F2: "foo bar"}
            _ = v7
        }
        _ = v5
    }
    v8 := GenS1{F0: 255, F1: 1, F2: "hello"}
    for fi9 := 0; fi9 < 3; fi9++ {
        if (18446744073709551615 <= 161) {
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 uint8 = 0
                v12 := GenS1{F0: 255, F1: -521, F2: "hello"}
                fmt.Println(fi9, fi10, v11)
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        _ = fi9
    }
    var v13 int16 = 0
    for fi14 := 0; fi14 < 17; fi14++ {
        if false {
            if (fi14 <= (fi14 / ((128) | 1))) {
                var v15 uint16 = 449
                _ = v15
            }
            for fi16 := 0; fi16 < 7; fi16++ {
                if false {
                    fmt.Println(v13, fi16, v13)
                    var v17 int8 = 1
                    _ = v17
                } else {
                    if true {
                        if false {
                            var v18 int32 = 42
                            var v19 uint = 18446744073709551615
                            v20 := []int{fi14, fi16}
                            var v21 uint = v19
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            v22 := GenS1{F0: 42, F1: -1, F2: "a"}
                            _ = v22
                        }
                        var v23 float32 = 1.5
                        _ = v23
                    }
                    v24 := []int{(fi16 % ((fi14) | 1)), (fi16 << ((42) & 63))}
                    if ((1 > 2) || ((fi14 << ((112) & 63)) >= fi14)) {
                        switch fi14 {
                        case 922:
                            fmt.Println(fi16, fi16, fi16)
                        case 2147483647:
                            v25 := GenS2{F0: -458, F1: fi16, F2: 0}
                            _ = v25
                        default:
                            var v26 float32 = (1e10 * -2.25)
                            _ = v26
                        }
                        v27 := GenS3{F0: 65535, F1: 4294967295}
                        switch "" {
                        case "a":
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                        case "":
                            var v28 uint = 526
                            _ = v28
                        }
                        _ = v27
                    }
                    _ = v24
                }
                fmt.Println(fi16, fi14)
                v13 = (v13 / (((v13 & v13)) | 1))
                switch -432 {
                case 4294967295:
                    v29 := GenS4{F0: "a", F1: -799}
                    _ = v29
                default:
                    var v30 float32 = 0.1
                    _ = v30
                }
                v31 := GenS3{F0: 1, F1: 965}
                _ = fi16
                _ = v31
            }
        } else {
            var v32 float64 = float64(0.1)
            _ = v32
        }
        v33 := []int{fi14, fi14, fi14, (fi14 * -1)}
        var v34 int32 = 301
        _ = fi14
        _ = v33
        _ = v34
    }
    _ = v13
    _ = v8
    return -674
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = -2.25
    v3 := make(map[string]uint32)
    v4 := make(map[int]string)
    v2 = float64(1e10)
    v35 := GenS5{F0: "hello"}
    var v36 uint8 = 232
    var v37 uint8 = v36
    v2 += float64((-2.25 * 1.5))
    var v38 uint = 42
    v39 := GenS6{F0: 4294967295, F1: 32767, F2: "a"}
    switch 9223372036854775807 {
    case 305:
        var v40 float32 = float32(float64(float32(v2)))
        _ = v40
    case 4294967295:
        v2 = v2
    case -428:
        fmt.Println(v37, v37)
    default:
        for fi41 := 0; fi41 < 14; fi41++ {
            v38 -= (v38 % ((uint(v37)) | 1))
            for fi42 := 0; fi42 < 1; fi42++ {
                fmt.Println(v2, fi41)
                v37 += (v36 - v36)
                _ = fi42
            }
            v38 = v38
            _ = fi41
        }
    }
    var v43 uint64 = uint64(v38)
    if (-339 <= 4294967295) {
        var v44 int = -931
        v45 := GenS7{F0: v37}
        v2 = float64(0.0)
        _ = v44
        _ = v45
    } else {
        var v46 int16 = 42
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v43
}
