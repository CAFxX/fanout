package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int; F1 uint8; F2 int64 }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 int64, p1 string) uint64 {
    var v4 uint8 = 93
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 int8 = -128
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: 1, F1: 856, F2: 279}
            var v9 float64 = 0.0
            p1 = p1
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 580
}
func genFn1() int64 {
    var v10 float64 = float64(1e10)
    fmt.Println(v10)
    for fi11 := 0; fi11 < 4; fi11++ {
        if false {
            var v12 bool = (!(!(4294967295 >= 334)))
            _ = v12
        }
        _ = fi11
    }
    v10 += (((-2.25 * (v10 - v10)) - 3.14159) + 0.1)
    if (42 < -1) {
        switch 32768 {
        case -945:
            var v13 int = -2147483648
            _ = v13
        case 834:
            v10 = v10
        }
    } else {
        var v14 string = ("" + "\n\t")
        var v15 uint = 0
        if (32767 == 2) {
            v15 -= 2
            if (1.5 != float32((v10 * 0.1))) {
                v14 = ""
                v16 := GenS2{F0: 57, F1: 137, F2: 657}
                _ = v16
            }
        }
        _ = v14
        _ = v15
    }
    _ = v10
    return 0
}

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{42, 568, 924, 2}
    v3 := make(map[int]int)
    var v17 int64 = 9223372036854775807
    fmt.Println(v1, v1)
    if (v17 < (v17 / ((9223372036854775807) | 1))) {
        delete(v3, 32768)
        var v18 uint16 = 32767
        switch v1 {
        case "\n\t":
            for fi19 := 0; fi19 < 11; fi19++ {
                v20 := GenS2{F0: (fi19 / ((fi19) | 1)), F1: 127, F2: 423}
                v17 = (v17 - (-(v17)))
                if false {
                    v17 = (v17 & (v17 / ((genFn1()) | 1)))
                }
                _ = fi19
                _ = v20
            }
        }
        _ = v18
    } else {
        if len(v2) > 0 {
            v2[0] = 246
        }
        var v21 int16 = 42
        _ = v21
    }
    v17 -= genFn1()
    var v22 int64 = (-(v17))
    v23 := GenS2{F0: 194, F1: 45, F2: -9223372036854775808}
    for fi24 := 0; fi24 < 20; fi24++ {
        for fi25 := 0; fi25 < 16; fi25++ {
            _ = len(v2)
            v1 = v1
            v26 := GenS3{F0: fi24, F1: 128}
            _ = fi25
            _ = v26
        }
        if (1000 < -989) {
            fmt.Println(v1, v1)
            v3[745] = (fi24 / (((-(fi24))) | 1))
            var v27 string = ((v1 + "hello") + "foo bar")
            _ = v27
        }
        switch (fi24 * 127) {
        case 688:
            fmt.Println(fi24)
        default:
            var v28 int = -850
            _ = v28
        }
        _ = fi24
    }
    v17 -= (v17 & (v17 | v22))
    var v29 bool = (2 > 0)
    for fi30 := 0; fi30 < 2; fi30++ {
        v31 := GenS1{F0: 2147483647, F1: 2147483647, F2: 32767}
        var v32 int = fi30
        _ = fi30
        _ = v31
        _ = v32
    }
    v33 := GenS1{F0: 42, F1: 539, F2: 32767}
    v17 += -9223372036854775808
    if v29 {
        v34 := GenS1{F0: 0, F1: 940, F2: 2}
        _ = v34
    } else {
        var v35 float32 = ((0.1 - (-2.25 / 1.5)) - 0.0)
        _ = v35
    }
    v36 := GenS2{F0: 127, F1: 195, F2: genFn1()}
    if v29 {
        var v37 int8 = 0
        var v38 uint32 = 737
        _ = v37
        _ = v38
    } else {
        v3[4294967295] = -659
    }
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v33
    _ = v36
}
