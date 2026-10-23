package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int32; F1 int32; F2 int }
type GenS4 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int64; F1 string; F2 int32 }
type GenS6 struct { F0 int64; F1 string; F2 int32 }
type GenS7 struct { F0 int64; F1 int16 }
func genFn0() int64 {
    var v5 float64 = 3.14159
    v6 := GenS1{F0: 127}
    if (!true) {
        v7 := GenS1{F0: 127}
        _ = v7
    }
    v8 := []int{}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = v5
    _ = v6
    _ = v8
    return -396
}
func genFn1() string {
    v9 := []int{1, 127}
    v9 = append(v9, 128)
    switch "hello" {
    case "hello":
        if false {
            switch "" {
            case "":
                var v10 uint16 = 65535
                _ = v10
            default:
                if (false || true) {
                    v11 := GenS1{F0: 127}
                    for fi12 := 0; fi12 < 9; fi12++ {
                        var v13 uint16 = 32767
                        var v14 uint16 = 255
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    var v15 string = "foo bar"
                    _ = v11
                    _ = v15
                }
            }
            var v16 uint32 = 2
            for fi17 := 0; fi17 < 4; fi17++ {
                v16 = v16
                fmt.Println(v16)
                fmt.Println(v16, v16)
                _ = fi17
            }
            _ = v16
        } else {
            for fi18 := 0; fi18 < 1; fi18++ {
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 uint16 = 0
                    v21 := GenS1{F0: 255}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                v22 := GenS2{F0: fi18, F1: 42}
                if true {
                    if true {
                        var v23 int32 = 248
                        v23 = 42
                        _ = v23
                    } else {
                        if false {
                            fmt.Println(fi18)
                            var v24 int32 = 2147483647
                            if ((657 != 1) || false) {
                                var v25 uint16 = 32767
                                v25 = v25
                                _ = v25
                            } else {
                                var v26 int64 = -1
                                _ = v26
                            }
                            _ = v24
                        } else {
                            if (!true) {
                                if (952 != 687) {
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                }
                                fmt.Println(fi18)
                                var v27 string = "hello"
                                _ = v27
                            }
                            v28 := GenS1{F0: 1}
                            if len(v9) > 0 {
                                v9[0] = (-(fi18))
                            }
                            _ = v28
                        }
                    }
                    for fi29 := 0; fi29 < 5; fi29++ {
                        v30 := GenS2{F0: -934, F1: -45}
                        _ = fi29
                        _ = v30
                    }
                } else {
                    v31 := GenS3{F0: 2147483647, F1: -2147483648, F2: (fi18 + 32767)}
                    var v32 string = "a"
                    v33 := GenS3{F0: -589, F1: 852, F2: fi18}
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = fi18
                _ = v22
            }
        }
    default:
        var v34 uint = 255
        _ = v34
    }
    _ = v9
    return "a"
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    v4 := make(map[int]uint32)
    v1 = (^(v1))
    switch 32767 {
    case -9223372036854775808:
        fmt.Println(v1)
    default:
        switch 32767 {
        case 2147483647:
            v35 := GenS1{F0: v1}
            _ = v35
        case 128:
            var v36 int16 = -1
            _ = v36
        }
    }
    var v37 string = "\n\t"
    for fi38 := 0; fi38 < 9; fi38++ {
        switch -578 {
        case 468:
            fmt.Println(v37, v1, v1)
        case 714:
            fmt.Println(v1, v37, v1)
        default:
            if false {
                switch v37 {
                case "\n\t":
                    var v39 int64 = 9223372036854775807
                    _ = v39
                }
            }
        }
        var v40 float64 = (0.1 + -2.25)
        _ = fi38
        _ = v40
    }
    fmt.Println(v1, v1)
    fmt.Println(v1, v37)
    for fi41 := 0; fi41 < 14; fi41++ {
        switch (fi41 << ((737) & 63)) {
        case 0:
            var v42 uint32 = 2147483648
            _ = v42
        case -9223372036854775808:
            v37 = (v37 + (v37 + (v37 + v37)))
        default:
            for fi43 := 0; fi43 < 15; fi43++ {
                var v44 uint8 = 2
                v45 := GenS1{F0: 2}
                fmt.Println(v44, v44)
                _ = fi43
                _ = v44
                _ = v45
            }
        }
        for fi46 := 0; fi46 < 12; fi46++ {
            switch "" {
            case "\n\t":
                v4[fi46] = 551
                var v47 uint64 = 1
                _ = v47
            case "":
                if ((!true) || (!false)) {
                    fmt.Println(v37)
                    v48 := []int{1, 128, fi46, fi41}
                    var v49 int64 = 42
                    _ = v48
                    _ = v49
                } else {
                    var v50 uint64 = 146
                    var v51 int = fi41
                    _ = v50
                    _ = v51
                }
            case "a":
                _ = len(v3)
            }
            if false {
                switch fi46 {
                case 9223372036854775807:
                    v1 -= v1
                    var v52 uint8 = (v1 & (^(v1)))
                    var v53 uint16 = 2
                    _ = v52
                    _ = v53
                case 891:
                    v54 := GenS4{F0: 100, F1: 251, F2: v1}
                    _ = v54
                case 2147483647:
                    v55 := GenS1{F0: 127}
                    _ = v55
                }
                v56 := GenS5{F0: 0, F1: v37, F2: 84}
                fmt.Println(fi41)
                _ = v56
            }
            _ = fi46
        }
        v57 := GenS2{F0: -634, F1: 1000}
        _ = fi41
        _ = v57
    }
    for fi58 := 0; fi58 < 14; fi58++ {
        v59 := []int{}
        if false {
            v1 -= v1
            if true {
                switch v37 {
                case "":
                    v60 := GenS6{F0: genFn0(), F1: ((v37 + (((v37 + ((v37 + "foo bar") + (v37 + "hello"))) + "hello") + (v37 + "a"))) + ((v37 + (("" + ("" + v37)) + v37)) + v37)), F2: -879}
                    _ = v60
                case "a":
                    v37 = v37
                case "hello":
                    for fi61 := 0; fi61 < 10; fi61++ {
                        if ((v1 << ((4294967295) & 63)) > 0) {
                            v62 := GenS7{F0: genFn0(), F1: -1}
                            _ = v62
                        }
                        if len(v59) > 0 {
                            v59[0] = fi58
                        }
                        v63 := v4[128]
                        delete(v2, fi61)
                        var v64 float32 = 0.0
                        fmt.Println(fi61, v64)
                        _ = fi61
                        _ = v63
                        _ = v64
                    }
                }
            } else {
                v37 = (v37 + "hello")
                v1 += v1
            }
        }
        _ = fi58
        _ = v59
    }
    var v65 int16 = 32767
    if (347 > 778) {
        v65 = v65
    }
    v66 := GenS1{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v65
    _ = v66
}
