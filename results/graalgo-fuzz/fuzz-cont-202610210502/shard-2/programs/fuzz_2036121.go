package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint32; F2 string }
type GenS4 struct { F0 int64; F1 int; F2 int8 }
func genFn0() int8 {
    var v4 uint64 = 4294967295
    v5 := GenS1{F0: 4294967295}
    _ = v4
    _ = v5
    return 127
}
func genFn1(p0 int8, p1 uint16, p2 int8) uint16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int64 = 0
        var v8 int32 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        if (!(!(true && false))) {
            if (!true) {
                if false {
                    var v10 uint8 = 1
                    _ = v10
                } else {
                    fmt.Println(fi9, p2)
                    p2 = p0
                    if true {
                        switch fi9 {
                        case 32767:
                            for fi11 := 0; fi11 < 9; fi11++ {
                                if (!false) {
                                    v12 := []int{fi11, 9223372036854775807, (fi9 + -2147483648)}
                                    if (!(!true)) {
                                        fmt.Println(p1)
                                    } else {
                                        p1 += 1
                                    }
                                    _ = v12
                                } else {
                                    v13 := []int{}
                                    _ = v13
                                }
                                var v14 int32 = -1
                                v14 = (v14 * -756)
                                p0 = p0
                                v15 := []int{fi11, (fi11 & (^(fi9))), (fi9 << ((42) & 63))}
                                _ = fi11
                                _ = v14
                                _ = v15
                            }
                        }
                    }
                }
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := GenS2{F0: 2, F1: (p0 - 0), F2: 128}
                    _ = fi16
                    _ = v17
                }
                v18 := []int{}
                _ = v18
            }
            var v19 int8 = genFn0()
            _ = v19
        } else {
            if true {
                var v20 int8 = -128
                for fi21 := 0; fi21 < 4; fi21++ {
                    var v22 int8 = p0
                    v23 := GenS1{F0: 255}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v20 -= (v20 + genFn0())
                _ = v20
            }
        }
        switch (-(fi9)) {
        case 127:
            v24 := []int{(fi9 - (-(fi9)))}
            _ = v24
        }
        _ = fi9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return (p1 ^ (p1 / (((^(p1))) | 1)))
}

func main() {
    var v1 int64 = 1
    var v2 int64 = -9223372036854775808
    v3 := []string{((((("" + "a") + "\n\t") + ("foo bar" + "\n\t")) + (("\n\t" + "\n\t") + "foo bar")) + "\n\t"), ("foo bar" + "a"), (("foo bar" + "\n\t") + ""), "foo bar"}
    for fi25 := 0; fi25 < 14; fi25++ {
        var v26 int16 = 0
        var v27 uint = 255
        _ = fi25
        _ = v26
        _ = v27
    }
    switch 32767 {
    case 0:
        v28 := GenS1{F0: 18446744073709551615}
        _ = v28
    }
    v1 = v2
    for fi29 := 0; fi29 < 10; fi29++ {
        for fi30 := 0; fi30 < 4; fi30++ {
            switch ("\n\t" + "\n\t") {
            case "hello":
                if (128 <= 97) {
                    fmt.Println(v2, fi30, fi29)
                }
            default:
                if true {
                    v2 = -755
                    v1 = v1
                }
            }
            switch "hello" {
            case "\n\t":
                var v31 int = 4294967295
                _ = v31
            case "a":
                v1 = (v1 | 9223372036854775807)
            default:
                v2 = (v1 ^ (v2 - 42))
            }
            if true {
                var v32 uint64 = 2
                var v33 int8 = 100
                fmt.Println(fi30)
                _ = v32
                _ = v33
            }
            _ = fi30
        }
        var v34 uint16 = 32767
        fmt.Println(fi29)
        _ = fi29
        _ = v34
    }
    v1 = v2
    v35 := GenS3{F0: 478, F1: 2147483647, F2: "\n\t"}
    v36 := GenS2{F0: 421, F1: -128, F2: 127}
    var v37 uint8 = 1
    v38 := GenS1{F0: 65535}
    switch 4294967295 {
    case 42:
        for fi39 := 0; fi39 < 11; fi39++ {
            if false {
                var v40 uint8 = 1
                _ = v40
            } else {
                if (0 <= 794) {
                    var v41 float64 = 0.0
                    if false {
                        v1 = v2
                        var v42 float32 = float32(v41)
                        _ = v42
                    } else {
                        var v43 uint16 = 1
                        v43 += 2
                        v1 = v2
                        _ = v43
                    }
                    _ = v41
                }
                fmt.Println(v2, v2, fi39)
            }
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v2, v1)
            case "a":
                for fi44 := 0; fi44 < 8; fi44++ {
                    var v45 bool = false
                    _ = fi44
                    _ = v45
                }
            }
            _ = fi39
        }
    case 1:
        v46 := GenS4{F0: 9223372036854775807, F1: 9223372036854775807, F2: 100}
        _ = v46
    }
    v1 = 1
    if (((-1 == genFn0()) && (!((v1 % ((937) | 1)) != (v2 * v1)))) || true) {
        v3 = append(v3, "\n\t")
    }
    v3 = append(v3, "\n\t")
    fmt.Println(v37, v37, v1)
    v3 = append(v3, "foo bar")
    for fi47 := 0; fi47 < 20; fi47++ {
        _ = len(v3)
        var v48 int8 = 36
        _ = fi47
        _ = v48
    }
    for fi49 := 0; fi49 < 12; fi49++ {
        fmt.Println(fi49)
        var v50 uint32 = 603
        if (-541 > 803) {
            var v51 int32 = -417
            _ = len(v3)
            _ = v51
        }
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
