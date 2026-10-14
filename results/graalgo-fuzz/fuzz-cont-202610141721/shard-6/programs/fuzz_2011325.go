package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint32 }
func genFn0() int16 {
    switch 127 {
    case 127:
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := []int{fi6, 4294967295, fi6}
            v7 = append(v7, (fi6 >> ((4294967295) & 63)))
            switch fi6 {
            case -9223372036854775808:
                switch fi6 {
                case 836:
                    var v8 int = (fi6 / ((127) | 1))
                    _ = v8
                default:
                    var v9 uint64 = 0
                    _ = v9
                }
            case 128:
                var v10 uint32 = 410
                _ = v10
            case 32768:
                fmt.Println(fi6)
            default:
                switch "" {
                case "":
                    var v11 int32 = 773
                    _ = v11
                case "a":
                    var v12 uint16 = 525
                    _ = v12
                case "foo bar":
                    for fi13 := 0; fi13 < 4; fi13++ {
                        var v14 uint = 1
                        v14 -= 18446744073709551615
                        _ = fi13
                        _ = v14
                    }
                default:
                    var v15 float64 = 0.0
                    _ = v15
                }
            }
            _ = fi6
            _ = v7
        }
    default:
        if (true && (42 > 127)) {
            v16 := []int{1, 1, 278, 713}
            var v17 uint8 = 233
            _ = v16
            _ = v17
        }
    }
    if true {
        if ((!(-9223372036854775808 != 1)) || (18446744073709551615 < 0)) {
            for fi18 := 0; fi18 < 5; fi18++ {
                v19 := []int{}
                _ = fi18
                _ = v19
            }
            v20 := []int{32767}
            _ = v20
        }
        var v21 string = "hello"
        v22 := GenS1{F0: 0}
        _ = v21
        _ = v22
    } else {
        var v23 uint = 0
        _ = v23
    }
    var v24 int64 = 9223372036854775807
    switch 4294967295 {
    case -2147483648:
        var v25 bool = ((true || false) || false)
        _ = v25
    case 128:
        switch -1 {
        case 2147483647:
            var v26 string = "foo bar"
            _ = v26
        case -9223372036854775808:
            for fi27 := 0; fi27 < 8; fi27++ {
                if (-680 < 32767) {
                    var v28 int64 = (v24 | v24)
                    v29 := GenS1{F0: 2}
                    fmt.Println(v28, v28)
                    _ = v28
                    _ = v29
                } else {
                    v30 := GenS1{F0: 4294967295}
                    _ = v30
                }
                v31 := GenS2{F0: 2}
                _ = fi27
                _ = v31
            }
        case -2147483648:
            for fi32 := 0; fi32 < 17; fi32++ {
                v24 -= (v24 % ((v24) | 1))
                _ = fi32
            }
        }
    case 32767:
        var v33 int64 = -9223372036854775808
        _ = v33
    }
    v34 := []int{32767, 2147483647}
    _ = v24
    _ = v34
    return -1
}

func main() {
    var v1 uint32 = 779
    v2 := make(map[int]uint32)
    var v3 string = "a"
    var v4 float32 = float32(float64(3.14159))
    v5 := make(map[uint32]uint32)
    fmt.Println(v4)
    for fi35 := 0; fi35 < 13; fi35++ {
        v4 -= v4
        for fi36 := 0; fi36 < 4; fi36++ {
            var v37 int = fi35
            _ = fi36
            _ = v37
        }
        if true {
            var v38 int64 = -9223372036854775808
            _ = v38
        } else {
            v3 = v3
        }
        _ = fi35
    }
    if (true && true) {
        v3 = ("\n\t" + v3)
        v3 = "foo bar"
    }
    v39 := GenS3{F0: (float64((v4 * v4)) >= ((3.14159 - 1.5) * 3.14159)), F1: 609}
    switch v3 {
    case "hello":
        for fi40 := 0; fi40 < 6; fi40++ {
            v4 = float32(0.0)
            _ = fi40
        }
    default:
        var v41 float64 = -2.25
        _ = v41
    }
    if true {
        var v42 int8 = 100
        v43 := GenS4{F0: v1}
        fmt.Println(v1, v1, v3)
        _ = v42
        _ = v43
    } else {
        v4 = v4
        var v44 int = 4294967295
        _ = v44
    }
    if false {
        var v45 uint32 = (v1 + (^(v1)))
        v3 = ((v3 + (v3 + ("hello" + "a"))) + v3)
        _ = v45
    } else {
        var v46 uint64 = 17
        var v47 uint8 = 198
        v3 = v3
        _ = v46
        _ = v47
    }
    v48 := GenS3{F0: true, F1: 4294967295}
    for fi49 := 0; fi49 < 16; fi49++ {
        if (0.1 <= -2.25) {
            v3 = "hello"
            fmt.Println(v3, v4, v3)
        } else {
            if false {
                if (v4 < (v4 / 1.0)) {
                    fmt.Println(v3)
                } else {
                    fmt.Println(v4)
                    v1 = v1
                    v4 += 1.5
                }
            }
            v4 = float32(-2.25)
            for fi50 := 0; fi50 < 11; fi50++ {
                v1 = (v1 - v1)
                _ = fi50
            }
        }
        _ = fi49
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v48
    _ = v5
}
