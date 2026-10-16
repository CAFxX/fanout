package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int8 }
type GenS2 struct { F0 string; F1 int; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 int8 }
func genFn0() int32 {
    var v5 uint16 = 32767
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 bool = true
        if (0 <= 104) {
            var v8 int16 = -32768
            v5 = (v5 << ((42) & 63))
            var v9 string = "foo bar"
            _ = v8
            _ = v9
        } else {
            v7 = v7
        }
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := []int{-2147483648, fi6, (fi6 * fi6)}
            _ = fi10
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    _ = v5
    return 42
}
func genFn1() int16 {
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS1{F0: 0, F1: 65535, F2: -1}
        v14 := []int{fi12, fi12, fi12}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 float32 = 1e10
    var v16 uint = 973
    _ = v15
    _ = v16
    return 224
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{("\n\t" + ""), "a", "a", "hello"}
    v3 := []uint32{2147483647, 0}
    var v4 int32 = 314
    fmt.Println(v4)
    var v17 uint16 = 32767
    fmt.Println(v4, v4)
    switch ("" + "a") {
    case "\n\t":
        v4 += (v4 & (v4 | v4))
    case "hello":
        v4 += -2147483648
    }
    var v18 uint16 = (v17 << ((2) & 63))
    switch 1 {
    case 54:
        v19 := GenS2{F0: "\n\t", F1: -9223372036854775808, F2: 328}
        _ = v19
    default:
        v20 := GenS3{F0: 447}
        _ = v20
    }
    v17 = (^(v18))
    v21 := GenS4{F0: 81, F1: 2147483647}
    fmt.Println(v18, v17, v4)
    var v22 uint64 = 1
    fmt.Println(v18, v4)
    if ((902 >= 857) && (!false)) {
        for fi23 := 0; fi23 < 5; fi23++ {
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 int16 = 0
                if ((((false && false) && (false && (1.5 != (float64(-2.25) - 0.1)))) && false) || (0 > (v17 * 697))) {
                    v4 = int32(v17)
                    if (4294967295 < 0) {
                        var v26 uint64 = (v22 >> ((629) & 63))
                        switch int(v4) {
                        case 4294967295:
                            if false {
                                v18 += v18
                                v25 -= v25
                            }
                        case -9223372036854775808:
                            v25 = genFn1()
                        default:
                            fmt.Println(v17, fi24)
                        }
                        v27 := GenS1{F0: 9223372036854775807, F1: 42, F2: -1}
                        _ = v26
                        _ = v27
                    }
                    if (3.14159 <= -2.25) {
                        v25 -= (-(v25))
                        v17 = 32767
                        v2 = append(v2, "")
                    } else {
                        if (128 <= 127) {
                            var v28 uint = 18446744073709551615
                            _ = v28
                        } else {
                            v25 = (^(v25))
                            v29 := GenS5{F0: 2}
                            _ = len(v1)
                            _ = v29
                        }
                        switch ("" + "hello") {
                        case "foo bar":
                            v30 := GenS3{F0: 1}
                            _ = v30
                        case "hello":
                            v4 = v4
                        }
                    }
                } else {
                    var v31 int16 = v25
                    v3 = append(v3, 995)
                    _ = v31
                }
                v32 := GenS5{F0: 0}
                v25 = v25
                _ = fi24
                _ = v25
                _ = v32
            }
            _ = fi23
        }
    }
    v1[4294967295] = 1
    var v33 int16 = 0
    if (63 == 1) {
        switch 127 {
        case 1:
            switch -208 {
            case 1:
                fmt.Println(v22, v18, v22)
            case -9223372036854775808:
                fmt.Println(v4, v33, v18)
            case 128:
                fmt.Println(v17)
            default:
                for fi34 := 0; fi34 < 3; fi34++ {
                    v35 := GenS6{F0: 9223372036854775807, F1: -128}
                    v36 := GenS1{F0: int(v22), F1: 824, F2: 111}
                    v37 := GenS3{F0: 829}
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
            }
        case 0:
            switch 2147483647 {
            case 628:
                v38 := v1[32768]
                _ = v38
            case -140:
                v22 = (v22 + 923)
            default:
                var v39 int16 = v33
                _ = v39
            }
        case 42:
            v18 = v18
        default:
            switch (((("hello" + "foo bar") + ("hello" + "foo bar")) + "foo bar") + ("\n\t" + "a")) {
            case "\n\t":
                v40 := GenS3{F0: 937}
                _ = v40
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v4
}
