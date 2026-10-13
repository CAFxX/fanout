package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int; F1 uint64; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int }

func main() {
    var v1 uint64 = 0
    v2 := []uint32{0}
    var v3 int8 = -128
    var v4 int16 = -711
    v5 := make(map[string]int)
    if len(v2) > 0 {
        v2[0] = 1
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    if false {
        var v6 uint8 = 120
        if false {
            for fi7 := 0; fi7 < 9; fi7++ {
                v1 += v1
                v8 := GenS1{F0: 42, F1: v6, F2: uint16(v1)}
                var v9 float32 = ((float32(((3.14159 - ((float64((((0.1 * float32((3.14159 + (1e10 * 1.5)))) + float32(float64(1.5))) + -2.25)) / 1e10) / float64(float32(0.1)))) / 1.0)) / (1.5 - 0.0)) + float32(1e10))
                _ = fi7
                _ = v8
                _ = v9
            }
        } else {
            v3 += int8(v4)
            fmt.Println(v6, v6)
        }
        switch "foo bar" {
        case "hello":
            v6 -= (v6 & (v6 % ((42) | 1)))
        case "\n\t":
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 int16 = (v4 * v4)
                _ = fi10
                _ = v11
            }
        case "foo bar":
            var v12 int32 = -2147483648
            _ = v12
        default:
            if true {
                if true {
                    var v13 uint = 0
                    v6 = v6
                    fmt.Println(v4)
                    _ = v13
                }
                v6 = (v6 & 128)
                switch "foo bar" {
                case "":
                    v14 := GenS2{F0: -2147483648, F1: v1, F2: 1000}
                    _ = v14
                case "a":
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                }
            }
        }
        _ = v6
    }
    switch "hello" {
    case "\n\t":
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 int32 = 61
            _ = fi15
            _ = v16
        }
    default:
        switch -317 {
        case 1:
            if (!true) {
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 int = fi17
                    for fi19 := 0; fi19 < 4; fi19++ {
                        if false {
                            var v20 int32 = 1
                            var v21 int = 380
                            _ = v20
                            _ = v21
                        }
                        _ = fi19
                    }
                    var v22 int8 = int8(v1)
                    _ = fi17
                    _ = v18
                    _ = v22
                }
                var v23 int16 = -1
                _ = len(v2)
                _ = v23
            } else {
                if (v3 == 1) {
                    switch 9223372036854775807 {
                    case 16:
                        var v24 uint64 = v1
                        _ = v24
                    case 32767:
                        switch (("hello" + "") + (("" + "hello") + "foo bar")) {
                        case "hello":
                            var v25 int64 = 9223372036854775807
                            _ = v25
                        case "":
                            var v26 string = ""
                            _ = v26
                        default:
                            for fi27 := 0; fi27 < 19; fi27++ {
                                if false {
                                    v4 += int16(v1)
                                } else {
                                    var v28 uint32 = 373
                                    v29 := v5["\n\t"]
                                    if true {
                                        var v30 int64 = int64(v28)
                                        _ = v30
                                    } else {
                                        v31 := GenS2{F0: 2147483647, F1: v1, F2: v4}
                                        _ = v31
                                    }
                                    _ = v28
                                    _ = v29
                                }
                                _ = fi27
                            }
                        }
                    case -819:
                        for fi32 := 0; fi32 < 17; fi32++ {
                            v3 += (^(v3))
                            v4 = (v4 | (v4 >> ((841) & 63)))
                            _ = fi32
                        }
                    default:
                        v33 := GenS3{F0: 32768}
                        _ = v33
                    }
                }
            }
        case 293:
            v34 := v5["hello"]
            _ = v34
        case -834:
            for fi35 := 0; fi35 < 4; fi35++ {
                var v36 int16 = (v4 - v4)
                for fi37 := 0; fi37 < 20; fi37++ {
                    v3 = -128
                    var v38 string = "hello"
                    v39 := GenS3{F0: fi35}
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                _ = fi35
                _ = v36
            }
        default:
            switch 887 {
            case -9223372036854775808:
                for fi40 := 0; fi40 < 9; fi40++ {
                    var v41 string = ("hello" + ("" + "hello"))
                    _ = fi40
                    _ = v41
                }
            default:
                switch 0 {
                case 0:
                    var v42 float32 = 3.14159
                    _ = v42
                }
            }
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v43 := GenS2{F0: 4294967295, F1: (v1 % (((^(v1))) | 1)), F2: v4}
    v1 += 42
    v3 = v3
    v44 := GenS1{F0: 42, F1: 255, F2: 65535}
    v4 -= (-(v4))
    v45 := GenS1{F0: 4294967295, F1: 1, F2: 42}
    v46 := GenS4{F0: 0, F1: 483}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v4, v1)
    v1 += v1
    switch 720 {
    case -914:
        v3 = v3
    default:
        v4 = v4
    }
    switch "a" {
    case "\n\t":
        _ = len(v5)
    case "foo bar":
        var v47 uint32 = 676
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
