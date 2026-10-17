package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint8) int16 {
    var v4 int32 = 1
    var v5 uint8 = 2
    var v6 int64 = -310
    v6 = -976
    if (!(false || false)) {
        switch (("foo bar" + "") + ("a" + "hello")) {
        case "hello":
            var v7 uint64 = 632
            _ = v7
        case "a":
            var v8 uint32 = 0
            _ = v8
        case "\n\t":
            v9 := GenS1{F0: 78, F1: 4294967295, F2: 4294967295}
            _ = v9
        default:
            var v10 uint = 65535
            _ = v10
        }
        var v11 uint64 = 969
        _ = v11
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return int16(v4)
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint16 = 255
    var v3 int16 = 398
    v12 := GenS2{F0: (42 > 2), F1: v2}
    v13 := GenS2{F0: (-700 == 9223372036854775807), F1: v2}
    v14 := GenS1{F0: -3, F1: 0, F2: v1}
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := []int{fi15}
        switch (fi15 >> ((18446744073709551615) & 63)) {
        case 1:
            var v17 uint = 690
            _ = v17
        default:
            var v18 uint16 = v2
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    var v19 string = "foo bar"
    if true {
        var v20 uint = 42
        v21 := GenS3{F0: v19}
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 17; fi22++ {
            switch "foo bar" {
            case "\n\t":
                var v23 int16 = (v3 & v3)
                _ = v23
            case "":
                v19 = v19
            case "foo bar":
                if (473 == 4294967295) {
                    var v24 int64 = 202
                    _ = v24
                } else {
                    if (v2 >= (v2 / ((188) | 1))) {
                        v25 := []int{127, -1}
                        _ = v25
                    } else {
                        v1 -= 721
                        switch fi22 {
                        case 1:
                            var v26 bool = (float64(-2.25) <= (float64(float32((3.14159 * 3.14159))) + (1.5 + 1e10)))
                            _ = v26
                        case -891:
                            v27 := []int{fi22, fi22}
                            _ = v27
                        default:
                            v1 -= 2
                        }
                    }
                    v28 := []int{fi22, (fi22 >> ((18446744073709551615) & 63))}
                    _ = v28
                }
            }
            _ = fi22
        }
        v29 := []int{-2147483648, 2147483647, 609}
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
        _ = v29
    }
    switch -739 {
    case 107:
        var v30 bool = true
        _ = v30
    case 619:
        var v31 int32 = 0
        _ = v31
    }
    v32 := []int{}
    var v33 int32 = 330
    v34 := GenS4{F0: (^(v2)), F1: true}
    switch ((((v19 + v19) + v19) + "foo bar") + (v19 + v19)) {
    case "a":
        switch v19 {
        case "\n\t":
            v1 += v1
        case "a":
            v32 = append(v32, 42)
        default:
            v35 := GenS4{F0: 32767, F1: (139 < uint8(v2))}
            _ = v35
        }
    case "hello":
        v33 += v33
    }
    v36 := GenS5{F0: (v1 << ((42) & 63))}
    var v37 bool = true
    fmt.Println(v37, v33, v37)
    fmt.Println(v1, v33)
    var v38 int = -503
    for fi39 := 0; fi39 < 5; fi39++ {
        for fi40 := 0; fi40 < 20; fi40++ {
            if (!v37) {
                v3 = (^(v3))
                switch fi40 {
                case 42:
                    fmt.Println(v33, v3)
                default:
                    if v37 {
                        switch v19 {
                        case "":
                            if len(v32) >= 2 {
                                v32 = v32[1:len(v32)]
                            }
                        }
                    } else {
                        if (v37 && v37) {
                            _ = len(v32)
                            v41 := GenS3{F0: v19}
                            _ = v41
                        } else {
                            fmt.Println(v19, v2)
                            var v42 uint32 = 642
                            _ = v42
                        }
                        var v43 float64 = ((float64(3.14159) - 1.5) * 0.0)
                        var v44 int8 = 100
                        if (57 < v3) {
                            v45 := GenS5{F0: v1}
                            var v46 uint = 336
                            fmt.Println(v2)
                            _ = v45
                            _ = v46
                        }
                        _ = v43
                        _ = v44
                    }
                    var v47 uint32 = (v1 & v1)
                    fmt.Println(v37)
                    _ = v47
                }
                if len(v32) > 0 {
                    v32[0] = v38
                }
            } else {
                var v48 int8 = 91
                _ = v48
            }
            var v49 int16 = 0
            fmt.Println(v37)
            _ = fi40
            _ = v49
        }
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v38
}
