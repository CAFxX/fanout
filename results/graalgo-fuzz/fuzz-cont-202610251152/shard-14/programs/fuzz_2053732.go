package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0() int8 {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 uint64 = 430
        v8 := []int{fi6}
        if (0 < 255) {
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 int16 = 42
                _ = fi9
                _ = v10
            }
            if false {
                if true {
                    v11 := GenS1{F0: -9223372036854775808, F1: -1}
                    if true {
                        v12 := GenS2{F0: 0}
                        _ = v12
                    }
                    _ = v11
                } else {
                    switch fi6 {
                    case 32767:
                        v13 := GenS1{F0: -1, F1: 4294967295}
                        _ = v13
                    case 662:
                        var v14 int32 = -682
                        _ = v14
                    }
                    var v15 int8 = -1
                    var v16 uint = 875
                    _ = v15
                    _ = v16
                }
                fmt.Println(v7)
            } else {
                if false {
                    for fi17 := 0; fi17 < 10; fi17++ {
                        var v18 float32 = 1.5
                        switch (fi6 - -377) {
                        case -1:
                            v18 += v18
                        case -9223372036854775808:
                            _ = len(v8)
                        }
                        _ = fi17
                        _ = v18
                    }
                    fmt.Println(v7, v7)
                    v7 = (v7 & v7)
                } else {
                    fmt.Println(v7, fi6)
                }
                if false {
                    var v19 int = fi6
                    v20 := GenS2{F0: 2}
                    var v21 bool = false
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    var v22 uint = 18446744073709551615
                    switch "hello" {
                    case "\n\t":
                        var v23 int64 = 42
                        _ = v23
                    case "hello":
                        v7 += v7
                    default:
                        v24 := GenS3{F0: 825, F1: -494}
                        _ = v24
                    }
                    if len(v8) > 0 {
                        v8[0] = fi6
                    }
                    _ = v22
                }
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v25 string = ("foo bar" + "foo bar")
    v26 := []int{-995, 42, 1}
    _ = v25
    _ = v26
    return 1
}

func main() {
    var v1 int = 2147483647
    var v2 int16 = -32768
    var v3 float64 = 0.1
    var v4 uint32 = 4294967295
    var v5 uint64 = 436
    v1 = (v1 >> ((65535) & 63))
    v27 := GenS3{F0: 878, F1: -2147483648}
    if false {
        v28 := GenS1{F0: 0, F1: 309}
        if (v2 <= (v2 | v2)) {
            fmt.Println(v3)
        }
        _ = v28
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 float64 = float64(0.1)
        for fi31 := 0; fi31 < 3; fi31++ {
            v32 := []int{int(v2), 714, (fi31 / ((v1) | 1)), v1}
            var v33 float32 = float32((v30 * -2.25))
            if true {
                v32 = append(v32, (fi29 % (((fi29 * int(v4))) | 1)))
                v3 = v3
                switch fi31 {
                case 128:
                    if false {
                        var v34 int32 = -1
                        _ = v34
                    }
                case -9223372036854775808:
                    v2 = v2
                case 0:
                    v2 = v2
                default:
                    if false {
                        if len(v32) > 0 {
                            v32[0] = fi31
                        }
                        _ = len(v32)
                    }
                    fmt.Println(v3, v1)
                    v35 := GenS3{F0: 476, F1: -2147483648}
                    var v36 float32 = float32(v3)
                    var v37 int16 = v2
                    _ = v35
                    _ = v36
                    _ = v37
                }
            }
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v5, v3)
    var v38 int = (-(v1))
    if (42 != 1) {
        var v39 int16 = (v2 & -32768)
        _ = v39
    }
    if false {
        var v40 string = ("a" + (("" + "\n\t") + ""))
        var v41 uint64 = v5
        switch (^(v38)) {
        case -387:
            v3 = 0.0
        case 136:
            var v42 int16 = int16(v4)
            _ = v42
        case -955:
            var v43 string = v40
            _ = v43
        }
        _ = v40
        _ = v41
    } else {
        fmt.Println(v5, v4)
        var v44 float64 = 0.1
        if (((!(!(!false))) && (!(false || false))) && true) {
            var v45 int8 = 100
            v38 -= (v38 - (v38 << ((4294967295) & 63)))
            _ = v45
        }
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
