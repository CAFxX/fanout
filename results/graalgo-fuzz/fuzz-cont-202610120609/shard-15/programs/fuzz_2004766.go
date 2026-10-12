package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 bool }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int64; F2 uint8 }
func genFn0() uint32 {
    var v4 uint16 = 336
    var v5 uint8 = 128
    var v6 float32 = -2.25
    if true {
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := []int{fi7, (fi7 - (fi7 << ((89) & 63))), (fi7 ^ (fi7 % (((fi7 >> ((65535) & 63))) | 1))), -34}
            _ = fi7
            _ = v8
        }
        if false {
            var v9 int8 = 118
            if true {
                if true {
                    v10 := GenS1{F0: v5, F1: v5, F2: false}
                    var v11 uint32 = 2
                    _ = v10
                    _ = v11
                }
                var v12 uint = 4294967295
                var v13 int16 = 1
                _ = v12
                _ = v13
            } else {
                v14 := []int{216}
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    v15 := GenS2{F0: 42, F1: (v5 ^ 0)}
                    _ = v15
                }
                _ = v14
            }
            v16 := []int{32768, 1, -1, 831}
            _ = v16
            _ = v9
        } else {
            if ((1.5 * (((float64(3.14159) * -2.25) * 3.14159) / -2.25)) != ((3.14159 - -2.25) * (-2.25 / 1e10))) {
                var v17 int16 = 855
                if true {
                    fmt.Println(v6, v5)
                } else {
                    if ((v17 ^ -32768) < (v17 / ((v17) | 1))) {
                        v5 = v5
                    }
                    var v18 int = -9223372036854775808
                    _ = v18
                }
                v19 := []int{127, 0}
                _ = v17
                _ = v19
            }
            v4 = (v4 >> ((46) & 63))
        }
        v20 := GenS2{F0: 145, F1: 65}
        _ = v20
    }
    _ = v4
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 int64, p1 int16) uint16 {
    if true {
        v21 := []int{751, -377, 32768}
        _ = v21
    } else {
        if (p0 >= 0) {
            if (true || (!true)) {
                if false {
                    p0 += 42
                    var v22 uint8 = 0
                    _ = v22
                } else {
                    v23 := []int{}
                    _ = v23
                }
            }
            switch (("hello" + "a") + "hello") {
            case "":
                for fi24 := 0; fi24 < 10; fi24++ {
                    p0 -= p0
                    for fi25 := 0; fi25 < 15; fi25++ {
                        var v26 bool = (!true)
                        v27 := []int{}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    _ = fi24
                }
            case "\n\t":
                var v28 float32 = (((3.14159 * (-2.25 * 1e10)) - (((3.14159 * -2.25) + ((-2.25 / float32(-2.25)) - 0.1)) * float32((float64(1.5) * (-2.25 - (1e10 * (-2.25 + (1.5 - -2.25)))))))) - 1.5)
                _ = v28
            }
            p0 -= -974
        } else {
            var v29 int64 = int64(p1)
            v30 := GenS1{F0: 2, F1: 255, F2: false}
            fmt.Println(v29, p0)
            _ = v29
            _ = v30
        }
    }
    var v31 int16 = 0
    if (!true) {
        var v32 bool = (!false)
        if v32 {
            v33 := []int{723, -484}
            if v32 {
                if ((v32 || false) && (1 == 65535)) {
                    var v34 uint8 = 255
                    for fi35 := 0; fi35 < 10; fi35++ {
                        switch fi35 {
                        case -570:
                            for fi36 := 0; fi36 < 17; fi36++ {
                                var v37 uint32 = genFn0()
                                _ = fi36
                                _ = v37
                            }
                        case 128:
                            fmt.Println(p0, v32, p1)
                        case 699:
                            if v32 {
                                p0 -= p0
                                var v38 int = (^(fi35))
                                var v39 int32 = 0
                                _ = v38
                                _ = v39
                            }
                        default:
                            switch "hello" {
                            case "a":
                                v32 = (!((18446744073709551615 != 2) && true))
                            case "hello":
                                v40 := GenS2{F0: (v31 + v31), F1: v34}
                                _ = v40
                            }
                        }
                        var v41 uint64 = 42
                        _ = fi35
                        _ = v41
                    }
                    var v42 bool = (!(84 == 42))
                    _ = v34
                    _ = v42
                }
            }
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            _ = v33
        }
        var v43 float64 = ((1e10 + 0.1) / (1.5 + float64(float32(0.0))))
        _ = v32
        _ = v43
    }
    switch "a" {
    case "a":
        v44 := GenS1{F0: 0, F1: 142, F2: (1 == 1)}
        _ = v44
    default:
        for fi45 := 0; fi45 < 3; fi45++ {
            if ((!(32767 >= 842)) && ("" < "hello")) {
                var v46 int64 = p0
                _ = v46
            }
            v47 := GenS2{F0: (p1 % (((p1 + p1)) | 1)), F1: 0}
            _ = fi45
            _ = v47
        }
    }
    v48 := GenS3{F0: 116, F1: 4294967295}
    _ = p0
    _ = p1
    _ = v31
    _ = v48
    return 32767
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 float32 = (0.0 / 1.5)
    v3 = (v3 - (1e10 - v3))
    fmt.Println(v3)
    v3 = ((v3 - v3) + (0.0 * v3))
    if false {
        v49 := GenS4{F0: 136}
        var v50 uint32 = 2147483647
        var v51 int16 = 792
        _ = v49
        _ = v50
        _ = v51
    } else {
        v3 -= 3.14159
    }
    v52 := GenS5{F0: (("foo bar" + "\n\t") + "a"), F1: -517, F2: 128}
    switch 32768 {
    case -975:
        v3 += v3
    case 1:
        fmt.Println(v3)
    case 0:
        v53 := []int{-301, 2147483647, -9223372036854775808, 127}
        _ = v53
    default:
        var v54 int16 = 79
        _ = v54
    }
    v1["foo bar"] = 1
    var v55 bool = (2 <= 2)
    v56 := v2[("hello" + "")]
    var v57 int = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v55
    _ = v56
    _ = v57
}
