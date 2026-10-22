package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int64) uint {
    v5 := GenS1{F0: -9223372036854775808}
    if true {
        v6 := []int{85, 697}
        p0 -= p0
        v7 := GenS1{F0: -1}
        _ = v6
        _ = v7
    } else {
        var v8 string = "a"
        switch 128 {
        case 9223372036854775807:
            v9 := []int{925, 553, 128}
            _ = v9
        }
        _ = v8
    }
    _ = p0
    _ = v5
    return 4294967295
}
func genFn1(p0 uint64, p1 int64, p2 uint32) int16 {
    if (false || (1 != 2)) {
        var v10 bool = true
        for fi11 := 0; fi11 < 12; fi11++ {
            switch fi11 {
            case 127:
                var v12 int64 = 42
                _ = v12
            case -106:
                v13 := GenS1{F0: fi11}
                _ = v13
            case 32768:
                for fi14 := 0; fi14 < 16; fi14++ {
                    var v15 float64 = (0.0 / float64(-2.25))
                    _ = fi14
                    _ = v15
                }
            default:
                fmt.Println(p1, fi11, p1)
            }
            _ = fi11
        }
        _ = v10
    }
    var v16 int16 = 711
    v17 := GenS1{F0: 4294967295}
    v16 += v16
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return v16
}

func main() {
    var v1 uint32 = 65535
    var v2 uint8 = 221
    var v3 bool = (2 != 887)
    var v4 uint = 324
    v18 := GenS1{F0: 1}
    v19 := []int{9223372036854775807, 4294967295}
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    v1 -= 420
    fmt.Println(v3, v4, v1)
    if (-15 != -128) {
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 uint64 = 939
            if (2147483647 <= 942) {
                var v22 int = 42
                _ = v22
            } else {
                if len(v19) > 0 {
                    v19[0] = 42
                }
            }
            _ = fi20
            _ = v21
        }
        switch 42 {
        case 870:
            var v23 int32 = 2147483647
            _ = v23
        case 32767:
            for fi24 := 0; fi24 < 6; fi24++ {
                if false {
                    v25 := GenS1{F0: (fi24 * (fi24 & (fi24 & (fi24 ^ 9223372036854775807))))}
                    if len(v19) > 0 {
                        v19[0] = (fi24 / (((fi24 - -9223372036854775808)) | 1))
                    }
                    _ = v25
                } else {
                    if (!(float32(float64((-2.25 + 0.0))) == (1e10 * 1.5))) {
                        var v26 uint64 = 18446744073709551615
                        var v27 uint8 = 151
                        _ = v26
                        _ = v27
                    } else {
                        v1 = v1
                    }
                    var v28 int8 = 127
                    v29 := GenS1{F0: -781}
                    _ = v28
                    _ = v29
                }
                v1 = (v1 << ((v4) & 63))
                _ = fi24
            }
        }
        switch 32768 {
        case 127:
            var v30 int64 = 153
            _ = v30
        case -1:
            switch ("" + "foo bar") {
            case "":
                fmt.Println(v1, v4)
            default:
                var v31 int32 = -327
                _ = v31
            }
        case -9223372036854775808:
            var v32 bool = (false || v3)
            _ = v32
        default:
            switch 127 {
            case -641:
                var v33 string = ("" + "a")
                _ = v33
            }
        }
    }
    fmt.Println(v2)
    var v34 int16 = 545
    if true {
        for fi35 := 0; fi35 < 4; fi35++ {
            for fi36 := 0; fi36 < 19; fi36++ {
                var v37 uint8 = v2
                v37 -= (v2 / ((uint8(v4)) | 1))
                _ = fi36
                _ = v37
            }
            v38 := GenS2{F0: 342}
            _ = fi35
            _ = v38
        }
        fmt.Println(v4, v34, v4)
    }
    switch 128 {
    case 495:
        v1 -= v1
    case 127:
        var v39 uint = (v4 * 698)
        _ = v39
    }
    v1 = 809
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
