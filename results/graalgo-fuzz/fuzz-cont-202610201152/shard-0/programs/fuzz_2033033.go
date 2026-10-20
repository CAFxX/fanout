package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint64 }

func main() {
    var v1 int16 = -300
    v2 := make(map[uint32]int)
    var v3 float32 = 1e10
    v4 := make(map[string]uint32)
    v5 := []int{4294967295, -9223372036854775808, -326}
    var v6 int8 = 1
    _ = len(v5)
    v3 += (v3 / (v3 * v3))
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint8 = 177
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int32 = 2147483647
            var v11 int = fi7
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        if (v3 > v3) {
            switch fi12 {
            case 466:
                var v13 float32 = float32((0.1 + float64(v3)))
                _ = v13
            case 4294967295:
                _ = len(v4)
            default:
                fmt.Println(v6, v3)
            }
            if (false && true) {
                var v14 uint16 = 952
                _ = v14
            } else {
                v6 += v6
                var v15 int8 = (^(v6))
                switch fi12 {
                case 32767:
                    if false {
                        switch (fi12 | (^(fi12))) {
                        case 1:
                            var v16 uint16 = 2
                            _ = v16
                        case -86:
                            v15 -= v6
                        case 32767:
                            var v17 uint64 = 396
                            _ = v17
                        default:
                            v1 -= v1
                        }
                        switch "hello" {
                        case "hello":
                            var v18 float64 = 1e10
                            _ = v18
                        }
                        var v19 int64 = 9223372036854775807
                        _ = v19
                    } else {
                        var v20 string = "\n\t"
                        for fi21 := 0; fi21 < 13; fi21++ {
                            var v22 uint8 = 133
                            _ = fi21
                            _ = v22
                        }
                        _ = v20
                    }
                case 699:
                    var v23 int64 = 9223372036854775807
                    _ = v23
                }
                _ = v15
            }
            v3 -= 3.14159
        }
        var v24 float64 = 1e10
        _ = len(v4)
        _ = fi12
        _ = v24
    }
    v6 -= v6
    v1 = v1
    for fi25 := 0; fi25 < 4; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            switch ("foo bar" + ("foo bar" + (("\n\t" + "") + "a"))) {
            case "a":
                v27 := GenS1{F0: 0, F1: 127, F2: 595}
                _ = v27
            default:
                if (!(186 >= 1)) {
                    if (0 == 42) {
                        fmt.Println(v6, fi26)
                    }
                    if true {
                        v5 = append(v5, -9223372036854775808)
                        if len(v5) > 0 {
                            v5[0] = 32768
                        }
                    } else {
                        fmt.Println(fi26)
                        _ = len(v5)
                    }
                    _ = len(v4)
                } else {
                    v6 = (v6 - (v6 & (v6 - -128)))
                }
            }
            if (982 > 65535) {
                v28 := v4[("hello" + (("hello" + "a") + ("hello" + "a")))]
                _ = v28
            } else {
                if true {
                    var v29 int64 = -734
                    var v30 uint32 = 248
                    if true {
                        v6 += -128
                    } else {
                        switch 4294967295 {
                        case 423:
                            v5 = append(v5, fi25)
                        case 4294967295:
                            fmt.Println(v6)
                        case -1:
                            if false {
                                v1 -= v1
                                var v31 int16 = int16(v6)
                                _ = v31
                            }
                        }
                        v29 -= -19
                    }
                    _ = v29
                    _ = v30
                }
                var v32 uint32 = 941
                if false {
                    var v33 uint32 = 384
                    v34 := GenS1{F0: fi26, F1: 128, F2: 4294967295}
                    _ = v33
                    _ = v34
                }
                switch "" {
                case "\n\t":
                    fmt.Println(v1, fi25)
                default:
                    var v35 uint16 = 946
                    _ = v35
                }
                fmt.Println(v3)
                v6 += (v6 >> ((42) & 63))
                _ = v32
            }
            _ = fi26
        }
        v6 -= 113
        _ = fi25
    }
    v5 = append(v5, 32767)
    fmt.Println(v1, v3)
    var v36 uint = 2
    if len(v5) > 0 {
        v5[0] = 32767
    }
    fmt.Println(v6, v36)
    if false {
        v37 := GenS1{F0: 326, F1: 0, F2: 4294967295}
        var v38 string = "\n\t"
        v39 := GenS2{F0: 314, F1: -1, F2: 104}
        _ = v37
        _ = v38
        _ = v39
    }
    v2[4294967295] = 0
    v3 -= float32(3.14159)
    v36 = 65535
    switch "a" {
    case "foo bar":
        v40 := GenS1{F0: -1, F1: 9, F2: 1}
        _ = v40
    case "":
        var v41 int16 = (-(v1))
        _ = v41
    default:
        if false {
            if len(v5) > 0 {
                v5[0] = 1
            }
            var v42 float32 = v3
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
