package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 int; F2 int64 }
type GenS5 struct { F0 uint32; F1 int16 }

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    v4 := []string{("\n\t" + "foo bar"), "foo bar", "foo bar", ("\n\t" + ("" + "a"))}
    delete(v2, 2147483648)
    switch "" {
    case "a":
        var v5 uint32 = 2147483648
        _ = v5
    case "hello":
        fmt.Println(v1)
    case "":
        fmt.Println(v1)
    default:
        var v6 float32 = 0.0
        _ = v6
    }
    var v7 bool = true
    fmt.Println(v7)
    var v8 uint16 = 191
    var v9 uint64 = 0
    for fi10 := 0; fi10 < 19; fi10++ {
        v8 += (^(v8))
        for fi11 := 0; fi11 < 19; fi11++ {
            switch (fi11 & (fi11 % (((-(fi11))) | 1))) {
            case 9223372036854775807:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                fmt.Println(v8)
            }
            _ = fi11
        }
        v7 = (-1 == 1)
        _ = fi10
    }
    if v7 {
        if (!v7) {
            v1 -= 1
        }
        v12 := v3[250]
        if (127 >= 128) {
            var v13 float32 = (1e10 * (float32((float64((3.14159 / 3.14159)) / (((float64(1.5) / (-2.25 * 0.1)) / -2.25) - (0.0 / 3.14159)))) / 1.5))
            _ = v13
        } else {
            if len(v4) > 0 {
                v4[0] = v12
            }
        }
        _ = v12
    }
    var v14 uint8 = 0
    var v15 uint = 628
    v16 := GenS1{F0: (v7 || true), F1: "foo bar"}
    fmt.Println(v1)
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            delete(v2, 177)
            v14 -= (v14 & (v14 << ((18446744073709551615) & 63)))
            _ = fi18
        }
        if len(v4) > 0 {
            v4[0] = (("" + "") + "foo bar")
        }
        fmt.Println(v7)
        _ = fi17
    }
    var v19 string = ((("a" + "a") + "foo bar") + "\n\t")
    if v7 {
        fmt.Println(v9, v9, v19)
        switch -1 {
        case 1:
            v20 := GenS1{F0: v7, F1: v19}
            _ = v20
        case 689:
            for fi21 := 0; fi21 < 16; fi21++ {
                switch (v19 + "") {
                case "":
                    switch (fi21 | 32768) {
                    case -610:
                        v22 := v2[1]
                        _ = v22
                    case -9223372036854775808:
                        v9 += v9
                    case 197:
                        v23 := GenS2{F0: (fi21 << ((v15) & 63)), F1: 42}
                        _ = v23
                    default:
                        for fi24 := 0; fi24 < 3; fi24++ {
                            var v25 float32 = 1.5
                            fmt.Println(fi24, fi24)
                            var v26 uint16 = v8
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                default:
                    var v27 int32 = -377
                    _ = v27
                }
                var v28 uint = (v15 + (v15 >> (((v15 << ((v15) & 63))) & 63)))
                _ = fi21
                _ = v28
            }
        default:
            for fi29 := 0; fi29 < 10; fi29++ {
                if v7 {
                    v9 = (v9 ^ (v9 | v9))
                } else {
                    _ = len(v2)
                    if v7 {
                        if ((^(v15)) < 42) {
                            v9 -= 1
                            v7 = true
                        }
                        switch fi29 {
                        case 0:
                            if (-2.25 >= (((1e10 * 0.1) * float64((-2.25 - float32(1e10)))) - (((3.14159 - 0.0) * ((3.14159 / -2.25) * (0.1 * (((((3.14159 / 1.5) - 0.1) + 1.5) * 0.1) / 1.0)))) / 1.0))) {
                                for fi30 := 0; fi30 < 3; fi30++ {
                                    var v31 float64 = 3.14159
                                    var v32 uint16 = v8
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                                for fi33 := 0; fi33 < 15; fi33++ {
                                    var v34 string = "hello"
                                    var v35 int32 = 425
                                    v1 -= int16(v15)
                                    _ = fi33
                                    _ = v34
                                    _ = v35
                                }
                                v9 -= (v9 * uint64(v8))
                            }
                        case -2147483648:
                            var v36 int8 = -1
                            _ = v36
                        }
                    }
                }
                var v37 int8 = 113
                _ = fi29
                _ = v37
            }
        }
        for fi38 := 0; fi38 < 7; fi38++ {
            v39 := v3[4294967295]
            _ = fi38
            _ = v39
        }
    } else {
        if (!true) {
            v40 := GenS1{F0: v7, F1: "a"}
            _ = v40
        } else {
            v41 := GenS3{F0: 1}
            _ = v41
        }
        var v42 int32 = -103
        _ = v42
    }
    v15 += v15
    v8 -= v8
    if v7 {
        fmt.Println(v1)
        v43 := GenS4{F0: 42, F1: -868, F2: 40}
        var v44 uint8 = (^(v14))
        _ = v43
        _ = v44
    } else {
        v45 := GenS1{F0: (true || (!(uint16(v14) != uint16(v14)))), F1: ""}
        v8 = uint16(v14)
        for fi46 := 0; fi46 < 2; fi46++ {
            var v47 int = (fi46 * 938)
            for fi48 := 0; fi48 < 19; fi48++ {
                v49 := GenS5{F0: 0, F1: 612}
                delete(v3, 0)
                v14 += v14
                v50 := GenS3{F0: 2147483648}
                var v51 int64 = 1
                _ = fi48
                _ = v49
                _ = v50
                _ = v51
            }
            _ = fi46
            _ = v47
        }
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
