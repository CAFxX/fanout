package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 bool) uint32 {
    var v4 int = -875
    if (0 > -598) {
        var v5 string = "foo bar"
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := GenS1{F0: 679, F1: 65535, F2: (!p0)}
            fmt.Println(v5, v4, fi6)
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        v4 += (v4 - 42)
        for fi8 := 0; fi8 < 14; fi8++ {
            fmt.Println(fi8, fi8)
            if p0 {
                v9 := GenS1{F0: 736, F1: 2147483647, F2: false}
                _ = v9
            }
            v10 := GenS1{F0: 32767, F1: 4294967295, F2: p0}
            _ = fi8
            _ = v10
        }
    }
    var v11 int16 = -763
    var v12 float32 = 0.1
    _ = p0
    _ = v11
    _ = v12
    _ = v4
    return 65535
}
func genFn1(p0 int64, p1 uint64, p2 uint16) int32 {
    var v13 uint8 = 128
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            switch ("" + ((("\n\t" + "a") + "\n\t") + "\n\t")) {
            case "":
                v16 := GenS1{F0: p2, F1: uint32(v13), F2: true}
                _ = v16
            case "a":
                var v17 float32 = 1.5
                _ = v17
            case "\n\t":
                if false {
                    var v18 uint64 = p1
                    _ = v18
                }
            }
            v13 -= v13
            p1 = p1
            _ = fi15
        }
        _ = fi14
    }
    v19 := []int{0, 1, 32767, 32768}
    v20 := GenS1{F0: (p2 - 0), F1: 65535, F2: true}
    if len(v19) > 0 {
        v19[0] = -581
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v19
    _ = v20
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 65535
    v3 := make(map[uint32]uint32)
    var v21 string = "foo bar"
    var v22 float64 = (float64((-2.25 - 1e10)) / -2.25)
    v22 = float64(1e10)
    var v23 int = -332
    switch -1 {
    case 32768:
        v21 = ("" + v21)
    default:
        for fi24 := 0; fi24 < 19; fi24++ {
            switch v21 {
            case "foo bar":
                if (false && false) {
                    fmt.Println(v23, v2)
                    fmt.Println(v22)
                }
            }
            var v25 int = (fi24 & v23)
            _ = fi24
            _ = v25
        }
    }
    v26 := GenS1{F0: 1, F1: 610, F2: (32767 >= 32767)}
    v27 := GenS1{F0: 2, F1: 532, F2: true}
    var v28 uint64 = 2
    if false {
        v23 += (v23 * (v23 ^ 4294967295))
        var v29 int16 = 42
        _ = v29
    } else {
        for fi30 := 0; fi30 < 13; fi30++ {
            if true {
                v28 += (v28 << ((593) & 63))
                v31 := []int{-2147483648, 853, (fi30 | v23), -392}
                var v32 uint = 255
                _ = v31
                _ = v32
            }
            var v33 float64 = v22
            var v34 uint = 42
            _ = fi30
            _ = v33
            _ = v34
        }
        v1[(v2 ^ (v2 ^ genFn0(true)))] = v21
    }
    v21 = ""
    v35 := GenS1{F0: 255, F1: v2, F2: (1e10 != -2.25)}
    fmt.Println(v28, v28)
    fmt.Println(v23, v23, v22)
    v22 -= 0.1
    if (true || ((1.5 / 1.0) < 1e10)) {
        var v36 int = -2147483648
        v37 := GenS1{F0: 1, F1: (v2 % ((v2) | 1)), F2: true}
        _ = v36
        _ = v37
    }
    if (1.5 >= v22) {
        if (42 == 482) {
            v38 := []int{}
            _ = v38
        }
        fmt.Println(v28, v21, v23)
    } else {
        switch (^(v23)) {
        case -701:
            v39 := GenS1{F0: 32767, F1: v2, F2: false}
            _ = v39
        case -9223372036854775808:
            for fi40 := 0; fi40 < 4; fi40++ {
                v41 := v3[genFn0(((v22 - -2.25) >= float64((1.5 / -2.25))))]
                _ = fi40
                _ = v41
            }
        case 9223372036854775807:
            if true {
                if false {
                    v42 := GenS2{F0: (^(v2))}
                    var v43 bool = (((42 == v28) && (1 < uint16(v2))) || false)
                    _ = v42
                    _ = v43
                } else {
                    v28 += v28
                }
                v44 := GenS2{F0: (^(v2))}
                _ = v44
            } else {
                v28 -= v28
                v21 = "foo bar"
                v45 := []int{v23, (v23 + (v23 ^ int(v28))), (v23 >> ((291) & 63)), v23}
                _ = v45
            }
        default:
            if false {
                for fi46 := 0; fi46 < 19; fi46++ {
                    v21 = v21
                    fmt.Println(v23, v22, v2)
                    _ = fi46
                }
                v28 -= v28
                fmt.Println(v22, v2, v2)
            } else {
                var v47 int16 = 1
                v48 := GenS3{F0: 0}
                fmt.Println(v2, v22, v23)
                _ = v47
                _ = v48
            }
        }
    }
    var v49 uint64 = (v28 ^ 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v49
}
