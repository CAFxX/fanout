package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint }
func genFn0(p0 int, p1 int64) int16 {
    for fi4 := 0; fi4 < 15; fi4++ {
        var v5 int64 = p1
        for fi6 := 0; fi6 < 5; fi6++ {
            p0 = (^(fi6))
            _ = fi6
        }
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: "a", F1: 127, F2: 42}
            var v9 int8 = 127
            p1 = 0
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi4
        _ = v5
    }
    var v10 int16 = -151
    _ = p0
    _ = p1
    _ = v10
    return (v10 & -210)
}

func main() {
    var v1 uint = 636
    var v2 uint8 = 169
    var v3 int64 = 156
    fmt.Println(v3, v3, v2)
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := []int{(fi11 % ((fi12) | 1)), fi12}
            v3 -= (v3 & (v3 - v3))
            v3 = (v3 - (v3 << (((^(v1))) & 63)))
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v14 := GenS2{F0: genFn0(1, v3)}
    switch 0 {
    case 148:
        for fi15 := 0; fi15 < 18; fi15++ {
            for fi16 := 0; fi16 < 13; fi16++ {
                v17 := []int{}
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    case 656:
        v3 = (v3 | v3)
    case -460:
        var v18 uint32 = 2
        _ = v18
    }
    var v19 string = "a"
    if true {
        v19 = v19
        v2 = v2
    } else {
        v20 := GenS2{F0: 0}
        switch 127 {
        case 9223372036854775807:
            v2 -= (v2 | 2)
        case 32767:
            for fi21 := 0; fi21 < 15; fi21++ {
                v22 := []int{4294967295, fi21, (fi21 + fi21)}
                _ = fi21
                _ = v22
            }
        case 32768:
            for fi23 := 0; fi23 < 11; fi23++ {
                v3 = v3
                _ = fi23
            }
        default:
            fmt.Println(v19)
        }
        var v24 int32 = -1
        _ = v20
        _ = v24
    }
    if (-128 < 0) {
        var v25 int8 = int8(v1)
        _ = v25
    } else {
        fmt.Println(v19)
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS1{F0: (v19 + (v19 + v19)), F1: 75, F2: v2}
            switch (-(fi26)) {
            case -2147483648:
                v28 := GenS2{F0: 958}
                _ = v28
            case 42:
                var v29 int32 = 1
                _ = v29
            default:
                v30 := []int{fi26, fi26, (fi26 / ((fi26) | 1))}
                _ = v30
            }
            _ = fi26
            _ = v27
        }
        for fi31 := 0; fi31 < 8; fi31++ {
            var v32 int = 9223372036854775807
            if (!false) {
                v33 := []int{}
                _ = v33
            } else {
                var v34 int16 = 32767
                v32 = (v32 & int(v3))
                _ = v34
            }
            _ = fi31
            _ = v32
        }
    }
    var v35 string = (v19 + v19)
    v1 = (v1 - 571)
    v36 := GenS2{F0: -1}
    if false {
        v37 := []int{-2147483648, -1, 4294967295}
        v38 := GenS3{F0: (v3 - v3), F1: 42, F2: (v1 - v1)}
        fmt.Println(v35, v1)
        _ = v37
        _ = v38
    }
    var v39 int64 = v3
    var v40 uint32 = 746
    var v41 int = 9223372036854775807
    if true {
        var v42 uint = (v1 & (v1 - (v1 % ((42) | 1))))
        v1 = (^(v1))
        var v43 int32 = 42
        _ = v42
        _ = v43
    }
    v40 = 2147483648
    if (1 < 1) {
        v3 += v39
    } else {
        v2 = v2
        v2 += v2
    }
    if (v3 != (v39 << ((18446744073709551615) & 63))) {
        v44 := []int{(v41 & v41), v41, (v41 + -9223372036854775808), v41}
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v40
    _ = v41
}
