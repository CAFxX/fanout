package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 string; F1 int16; F2 uint }
type GenS4 struct { F0 int; F1 string; F2 uint64 }
func genFn0() uint64 {
    switch 823 {
    case 1:
        var v4 int32 = -2147483648
        _ = v4
    case 393:
        var v5 int8 = 0
        _ = v5
    case -2147483648:
        var v6 uint32 = 1
        _ = v6
    default:
        v7 := []int{35}
        _ = v7
    }
    var v8 uint = 554
    _ = v8
    return 2
}

func main() {
    var v1 float64 = float64(1e10)
    var v2 int64 = 42
    v3 := make(map[uint32]uint32)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int8 = -105
        _ = fi9
        _ = v10
    }
    fmt.Println(v2, v2)
    var v11 int8 = 0
    var v12 bool = (930 == 65535)
    var v13 uint8 = 217
    switch ("hello" + "a") {
    case "a":
        if false {
            v14 := []int{32768, -9223372036854775808, -9223372036854775808}
            var v15 float64 = v1
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 6; fi16++ {
                v12 = false
                _ = fi16
            }
            v2 -= 42
        }
    default:
        v17 := GenS1{F0: (0.0 > 0.0), F1: 2}
        _ = v17
    }
    switch ("" + ("foo bar" + "hello")) {
    case "hello":
        for fi18 := 0; fi18 < 17; fi18++ {
            fmt.Println(fi18, v1, v1)
            if (false || true) {
                var v19 int64 = v2
                _ = v19
            } else {
                var v20 uint8 = (v13 * v13)
                var v21 uint16 = 174
                _ = v20
                _ = v21
            }
            switch (-(fi18)) {
            case 148:
                switch fi18 {
                case 128:
                    var v22 int32 = 83
                    _ = v22
                case 42:
                    v13 += v13
                case -1:
                    var v23 uint32 = 4294967295
                    _ = v23
                default:
                    v24 := []int{fi18, (fi18 + (fi18 >> ((18446744073709551615) & 63))), 4294967295}
                    _ = v24
                }
            case -9223372036854775808:
                var v25 uint = 1
                _ = v25
            case 32767:
                if (v12 || ((-(v11)) <= v11)) {
                    var v26 int8 = 0
                    v13 -= 255
                    _ = v26
                }
            default:
                if v12 {
                    fmt.Println(v1, v12, v2)
                    for fi27 := 0; fi27 < 3; fi27++ {
                        v28 := GenS2{F0: 42, F1: (-(v11))}
                        v3[2] = 4294967295
                        _ = fi27
                        _ = v28
                    }
                    if v12 {
                        for fi29 := 0; fi29 < 20; fi29++ {
                            fmt.Println(v2)
                            if (163 < v13) {
                                v30 := GenS1{F0: (!(fi29 != (fi29 >> ((981) & 63)))), F1: 1}
                                v31 := []int{0}
                                _ = v30
                                _ = v31
                            }
                            _ = fi29
                        }
                        for fi32 := 0; fi32 < 20; fi32++ {
                            v33 := GenS3{F0: "\n\t", F1: 32767, F2: 0}
                            _ = fi32
                            _ = v33
                        }
                        switch 32768 {
                        case 32767:
                            var v34 string = "hello"
                            _ = v34
                        case 128:
                            var v35 uint8 = (^(v13))
                            _ = v35
                        }
                    } else {
                        for fi36 := 0; fi36 < 11; fi36++ {
                            var v37 float64 = float64(float32((v1 - 0.0)))
                            v37 += v37
                            switch fi36 {
                            case -980:
                                fmt.Println(fi18, v13)
                            default:
                                v1 = v37
                            }
                            _ = fi36
                            _ = v37
                        }
                        var v38 string = "foo bar"
                        _ = v38
                    }
                }
            }
            _ = fi18
        }
    case "":
        v39 := GenS4{F0: 32767, F1: ("\n\t" + ("\n\t" + "\n\t")), F2: 42}
        _ = v39
    default:
        v13 = v13
    }
    var v40 int16 = 0
    v41 := []int{562, -2147483648, -2147483648}
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v40
    _ = v41
}
