package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint32; F2 int32 }
type GenS3 struct { F0 int8; F1 int8; F2 int8 }
type GenS4 struct { F0 uint64; F1 int64 }
func genFn0(p0 string) uint64 {
    var v4 float64 = ((0.1 + 0.0) / 1.0)
    p0 = p0
    v5 := []int{2147483647, 32767}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        p0 = p0
        _ = fi6
    }
    _ = p0
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 uint32 = 2
    var v2 int32 = -2147483648
    var v3 uint = 4294967295
    v7 := GenS1{F0: 780}
    if true {
        switch 2147483647 {
        case 0:
            fmt.Println(v1, v3, v2)
        case 1:
            var v8 bool = ((3.14159 <= 1.5) && true)
            _ = v8
        case 9223372036854775807:
            if (-32768 > 806) {
                var v9 float32 = 0.0
                _ = v9
            } else {
                var v10 uint32 = (v1 >> ((v3) & 63))
                if true {
                    var v11 int8 = 10
                    _ = v11
                } else {
                    var v12 int16 = -1
                    _ = v12
                }
                if ((-(v2)) <= (v2 >> ((v3) & 63))) {
                    if false {
                        if true {
                            v13 := []int{127, 128, 210, int(v3)}
                            for fi14 := 0; fi14 < 2; fi14++ {
                                for fi15 := 0; fi15 < 3; fi15++ {
                                    if ("\n\t" >= ("hello" + "\n\t")) {
                                        if true {
                                            fmt.Println(v10, v3)
                                            var v16 uint32 = (v10 >> ((v3) & 63))
                                            v1 += 1
                                            _ = v16
                                        }
                                        v2 = (v2 + (v2 * -2147483648))
                                        var v17 float64 = (((1e10 * 3.14159) * float64(1e10)) / ((1.5 * ((1.5 / 0.1) - (3.14159 + -2.25))) + 3.14159))
                                        _ = v17
                                    } else {
                                        v2 = (v2 - v2)
                                    }
                                    var v18 uint = v3
                                    _ = fi15
                                    _ = v18
                                }
                                if len(v13) > 0 {
                                    v13[0] = fi14
                                }
                                _ = fi14
                            }
                            v3 = (v3 + (v3 << ((v3) & 63)))
                            _ = v13
                        }
                        var v19 uint = v3
                        _ = v19
                    }
                    v10 += 1
                    v10 += v1
                } else {
                    v20 := GenS2{F0: 592, F1: v10, F2: v2}
                    v21 := GenS3{F0: -128, F1: -120, F2: 1}
                    _ = v20
                    _ = v21
                }
                _ = v10
            }
        }
        var v22 int8 = 127
        v23 := []int{-1, 127, 0, -2147483648}
        _ = v22
        _ = v23
    }
    v24 := []int{}
    for fi25 := 0; fi25 < 18; fi25++ {
        for fi26 := 0; fi26 < 12; fi26++ {
            switch "foo bar" {
            case "a":
                var v27 int64 = -1
                _ = v27
            case "\n\t":
                v1 = (v1 - (v1 ^ (v1 & v1)))
            default:
                var v28 string = ""
                _ = v28
            }
            switch "\n\t" {
            case "foo bar":
                if (54 <= 255) {
                    v2 -= (v2 % ((v2) | 1))
                    v29 := GenS1{F0: 42}
                    _ = v29
                } else {
                    if len(v24) > 0 {
                        v24[0] = 4294967295
                    }
                }
            case "a":
                var v30 bool = false
                _ = v30
            default:
                fmt.Println(fi26, fi25, fi26)
            }
            _ = fi26
        }
        _ = fi25
    }
    var v31 float32 = 1e10
    v1 = 65535
    v3 = (v3 << ((42) & 63))
    var v32 int32 = (v2 >> (((v3 | (v3 | v3))) & 63))
    v31 = v31
    if false {
        var v33 uint8 = 42
        switch "a" {
        case "a":
            v34 := GenS4{F0: 42, F1: -9223372036854775808}
            _ = v34
        case "\n\t":
            if len(v24) > 0 {
                v24[0] = 32767
            }
        case "hello":
            var v35 bool = (!(206 != -824))
            _ = v35
        }
        for fi36 := 0; fi36 < 3; fi36++ {
            v37 := GenS2{F0: 2147483647, F1: v1, F2: (v2 + (v32 * (-(v2))))}
            _ = fi36
            _ = v37
        }
        _ = v33
    } else {
        if len(v24) > 0 {
            v24[0] = 4294967295
        }
    }
    fmt.Println(v1, v1, v32)
    fmt.Println(v32, v31, v2)
    var v38 uint16 = 32767
    var v39 uint = 372
    var v40 uint64 = 246
    v38 -= v38
    fmt.Println(v31)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v39
    _ = v40
    _ = v7
}
