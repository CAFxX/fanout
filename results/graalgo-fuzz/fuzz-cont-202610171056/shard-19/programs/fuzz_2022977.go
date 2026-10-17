package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint64; F1 string; F2 int16 }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint) uint8 {
    v4 := GenS1{F0: -1, F1: 32767, F2: 1}
    v5 := GenS1{F0: 9223372036854775807, F1: -236, F2: 42}
    var v6 uint32 = 2
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 189
}
func genFn1(p0 string) uint8 {
    if true {
        if (1 < 394) {
            var v7 uint64 = 2
            for fi8 := 0; fi8 < 15; fi8++ {
                p0 = p0
                var v9 int16 = -140
                var v10 uint32 = 2147483647
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        } else {
            fmt.Println(p0)
            v11 := GenS2{F0: true, F1: 2}
            _ = v11
        }
        v12 := []int{-9223372036854775808, 42}
        p0 = ("" + (p0 + ("a" + "foo bar")))
        _ = v12
    }
    v13 := GenS3{F0: 127, F1: -70}
    for fi14 := 0; fi14 < 12; fi14++ {
        if (42 < -1) {
            fmt.Println(fi14, p0)
            var v15 uint64 = 0
            v16 := []int{fi14}
            _ = v15
            _ = v16
        } else {
            var v17 float32 = 1.5
            p0 = "\n\t"
            _ = v17
        }
        _ = fi14
    }
    if true {
        var v18 int8 = 100
        v19 := GenS2{F0: false, F1: 1}
        _ = v18
        _ = v19
    }
    var v20 int = -1
    _ = p0
    _ = v13
    _ = v20
    return 128
}

func main() {
    var v1 float32 = 1.5
    var v2 string = ""
    var v3 uint32 = 65535
    v21 := GenS4{F0: 4294967295, F1: 4294967295}
    var v22 int = 0
    v23 := []int{}
    if true {
        for fi24 := 0; fi24 < 2; fi24++ {
            v1 = (v1 - v1)
            v22 -= (v22 / (((v22 * v22)) | 1))
            _ = fi24
        }
        var v25 int16 = -914
        _ = v25
    }
    fmt.Println(v22)
    fmt.Println(v1)
    if (!((!(v2 <= v2)) && ((!(!((v22 / (((v22 << ((955) & 63))) | 1)) > v22))) || false))) {
        var v26 int8 = 0
        switch (^(v22)) {
        case -722:
            fmt.Println(v3, v22)
        case 127:
            for fi27 := 0; fi27 < 2; fi27++ {
                var v28 int32 = -1
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                for fi29 := 0; fi29 < 3; fi29++ {
                    v1 = v1
                    v22 = 32767
                    v30 := GenS1{F0: 42, F1: -32768, F2: 42}
                    _ = fi29
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
        default:
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
        }
        _ = v26
    } else {
        fmt.Println(v2)
    }
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    if (v2 <= v2) {
        v2 = (v2 + v2)
        switch v2 {
        case "\n\t":
            v23 = append(v23, v22)
        case "a":
            v23 = append(v23, 322)
        }
        for fi31 := 0; fi31 < 4; fi31++ {
            v22 -= (^(v22))
            if true {
                v32 := GenS1{F0: 9223372036854775807, F1: -140, F2: 1}
                if ((3.14159 / float64(((v1 * v1) / 1.0))) > (1e10 + float64((1e10 - 1.5)))) {
                    var v33 bool = false
                    var v34 int = 0
                    _ = v33
                    _ = v34
                }
                _ = v32
            }
            _ = len(v23)
            _ = fi31
        }
    }
    v22 -= -9223372036854775808
    switch v22 {
    case 9223372036854775807:
        switch v22 {
        case 2147483647:
            for fi35 := 0; fi35 < 18; fi35++ {
                for fi36 := 0; fi36 < 9; fi36++ {
                    switch v2 {
                    case "hello":
                        switch fi36 {
                        case -2147483648:
                            var v37 int16 = -1
                            _ = v37
                        case -187:
                            fmt.Println(v22)
                        case 127:
                            v2 = v2
                        default:
                            var v38 uint32 = 4294967295
                            var v39 int = fi36
                            v2 = (v2 + "\n\t")
                            _ = v38
                            _ = v39
                        }
                        fmt.Println(fi35, v22)
                    case "\n\t":
                        v2 = v2
                    case "foo bar":
                        _ = len(v23)
                    }
                    var v40 int16 = 32767
                    var v41 uint32 = (v3 % (((v3 << ((916) & 63))) | 1))
                    v3 = (v41 & (v41 << ((327) & 63)))
                    _ = fi36
                    _ = v40
                    _ = v41
                }
                v42 := GenS5{F0: 0, F1: v2, F2: -32768}
                v43 := GenS1{F0: 9223372036854775807, F1: -32768, F2: 4294967295}
                _ = fi35
                _ = v42
                _ = v43
            }
        }
    default:
        _ = len(v23)
    }
    v3 = (v3 & (v3 % ((v3) | 1)))
    var v44 uint = 4294967295
    if (!(1 > 1)) {
        switch 0 {
        case 39:
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
        case 32768:
            switch v22 {
            case -9223372036854775808:
                v3 -= 65535
            case -185:
                v45 := GenS6{F0: 32767}
                _ = v45
            case 15:
                if len(v23) > 0 {
                    v23[0] = -9223372036854775808
                }
            }
        default:
            switch ((v2 + v2) + (((((v2 + "\n\t") + ("hello" + "a")) + (v2 + v2)) + "hello") + ("" + (("a" + (v2 + "a")) + v2)))) {
            case "a":
                var v46 uint32 = v3
                _ = v46
            case "hello":
                switch v22 {
                case 32767:
                    var v47 int32 = 2147483647
                    _ = v47
                case 32768:
                    v48 := GenS4{F0: (^(v44)), F1: (v44 - 65535)}
                    _ = v48
                case -2147483648:
                    v23 = append(v23, 328)
                default:
                    fmt.Println(v3, v3)
                }
            default:
                fmt.Println(v3)
            }
        }
    }
    for fi49 := 0; fi49 < 17; fi49++ {
        var v50 float32 = v1
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v44
}
