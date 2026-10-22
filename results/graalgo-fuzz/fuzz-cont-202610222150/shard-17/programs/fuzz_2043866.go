package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint8; F2 int }
type GenS4 struct { F0 uint16 }
func genFn0(p0 string, p1 uint8, p2 uint8) uint64 {
    var v4 int64 = 0
    v5 := GenS1{F0: 2}
    var v6 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int32 = 42
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    v1 += v1
    if true {
        switch 4294967295 {
        case 0:
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 float64 = -2.25
                fmt.Println(v1, fi7, v8)
                for fi9 := 0; fi9 < 8; fi9++ {
                    fmt.Println(v8, v8, fi9)
                    var v10 uint32 = 42
                    _ = fi9
                    _ = v10
                }
                _ = fi7
                _ = v8
            }
        default:
            if false {
                var v11 int = 32767
                _ = v11
            } else {
                for fi12 := 0; fi12 < 20; fi12++ {
                    v1 = 42
                    _ = fi12
                }
                if false {
                    v2["a"] = "\n\t"
                } else {
                    fmt.Println(v1)
                    var v13 uint32 = 65535
                    for fi14 := 0; fi14 < 15; fi14++ {
                        switch (fi14 >> ((18446744073709551615) & 63)) {
                        case -564:
                            if true {
                                var v15 float64 = (float64(float32(3.14159)) / 1.0)
                                var v16 int8 = 1
                                _ = v15
                                _ = v16
                            } else {
                                var v17 uint32 = v13
                                for fi18 := 0; fi18 < 3; fi18++ {
                                    if true {
                                        v19 := []int{}
                                        var v20 uint16 = 0
                                        _ = v19
                                        _ = v20
                                    } else {
                                        if (127 < 2) {
                                            v13 = 42
                                        } else {
                                            v17 = (^(v17))
                                            switch fi14 {
                                            case 127:
                                                var v21 int = (-(fi18))
                                                _ = v21
                                            default:
                                                _ = len(v3)
                                            }
                                        }
                                    }
                                    if (0 == 127) {
                                        fmt.Println(v1)
                                        switch "" {
                                        case "hello":
                                            var v22 bool = true
                                            _ = v22
                                        }
                                        fmt.Println(v13)
                                    } else {
                                        v23 := GenS1{F0: 0}
                                        if false {
                                            switch "" {
                                            case "":
                                                v24 := GenS1{F0: 1}
                                                _ = v24
                                            case "hello":
                                                var v25 int = (fi14 >> ((255) & 63))
                                                _ = v25
                                            case "\n\t":
                                                var v26 int16 = 1000
                                                _ = v26
                                            }
                                            v27 := GenS2{F0: 0}
                                            _ = v27
                                        } else {
                                            v13 -= v13
                                        }
                                        v17 -= v17
                                        _ = v23
                                    }
                                    if (798 != 255) {
                                        var v28 uint = 29
                                        v13 = (^(v17))
                                        v1 -= int32(v13)
                                        v29 := GenS3{F0: 65535, F1: 1, F2: (^(fi18))}
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = fi18
                                }
                                _ = v17
                            }
                        case 128:
                            for fi30 := 0; fi30 < 17; fi30++ {
                                v13 -= v13
                                _ = fi30
                            }
                        case -2147483648:
                            fmt.Println(v1, v13, v13)
                        }
                        fmt.Println(fi14)
                        var v31 string = (("\n\t" + "\n\t") + ("hello" + "\n\t"))
                        _ = fi14
                        _ = v31
                    }
                    _ = v13
                }
            }
        }
        var v32 uint8 = 128
        if false {
            v1 = v1
        }
        _ = v32
    }
    var v33 uint = 65535
    for fi34 := 0; fi34 < 15; fi34++ {
        switch "a" {
        case "hello":
            fmt.Println(v1, v33, v1)
        case "foo bar":
            if false {
                fmt.Println(v33)
                fmt.Println(v1, v1, fi34)
            } else {
                var v35 int8 = -128
                v35 -= v35
                _ = v35
            }
        default:
            v1 -= v1
        }
        _ = fi34
    }
    v36 := GenS3{F0: 42, F1: 255, F2: 501}
    var v37 int8 = 0
    v1 += v1
    if ((-1 == 1) || true) {
        v38 := []int{-9223372036854775808, 32768, 32768}
        _ = v38
    }
    var v39 int64 = 42
    switch -9223372036854775808 {
    case 4294967295:
        v2["a"] = ("hello" + (("a" + "a") + "\n\t"))
    case 1:
        v40 := GenS2{F0: v33}
        _ = v40
    case -932:
        for fi41 := 0; fi41 < 12; fi41++ {
            if (-2.25 >= -2.25) {
                v39 = v39
            } else {
                for fi42 := 0; fi42 < 3; fi42++ {
                    var v43 uint32 = 2
                    _ = fi42
                    _ = v43
                }
            }
            fmt.Println(fi41, v37, v37)
            _ = fi41
        }
    }
    v44 := GenS4{F0: 2}
    fmt.Println(v33)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v39
    _ = v44
}
