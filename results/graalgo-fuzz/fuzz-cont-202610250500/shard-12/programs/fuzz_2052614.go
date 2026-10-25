package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 int64; F2 int8 }
type GenS5 struct { F0 int64; F1 bool }
func genFn0() bool {
    v5 := GenS1{F0: 42, F1: ""}
    if (127 >= 59) {
        if (!(!((1 >= -128) || false))) {
            if true {
                for fi6 := 0; fi6 < 1; fi6++ {
                    fmt.Println(fi6)
                    var v7 float32 = 0.0
                    _ = fi6
                    _ = v7
                }
                switch "" {
                case "foo bar":
                    var v8 bool = false
                    _ = v8
                default:
                    if ((940 >= -1) && false) {
                        var v9 uint8 = 128
                        _ = v9
                    } else {
                        switch ("" + "") {
                        case "hello":
                            if (!true) {
                                var v10 float64 = -2.25
                                fmt.Println(v10)
                                _ = v10
                            }
                        case "\n\t":
                            var v11 string = "hello"
                            _ = v11
                        default:
                            var v12 float32 = float32(float64((3.14159 * -2.25)))
                            _ = v12
                        }
                    }
                }
                v13 := []int{-733, 9223372036854775807, -651, 231}
                _ = v13
            } else {
                v14 := []int{32767, 2147483647, -1}
                var v15 int16 = -702
                v15 = 32767
                _ = v14
                _ = v15
            }
            v16 := []int{-1, -9223372036854775808}
            _ = v16
        } else {
            v17 := GenS2{F0: 249, F1: 517}
            switch -2147483648 {
            case -922:
                var v18 int8 = -41
                _ = v18
            case 1:
                var v19 float32 = (1.5 * (-2.25 - 0.0))
                _ = v19
            case -583:
                var v20 uint = 73
                _ = v20
            }
            _ = v17
        }
        var v21 uint64 = 433
        _ = v21
    } else {
        v22 := GenS1{F0: 0, F1: "a"}
        _ = v22
    }
    _ = v5
    return true
}
func genFn1(p0 int, p1 int8) int {
    p0 += p0
    var v23 uint64 = 1
    for fi24 := 0; fi24 < 14; fi24++ {
        for fi25 := 0; fi25 < 18; fi25++ {
            v26 := []int{9223372036854775807}
            if false {
                p0 = -9223372036854775808
                v27 := GenS3{F0: 962}
                if len(v26) > 0 {
                    v26[0] = (-(fi24))
                }
                _ = v27
            } else {
                var v28 uint64 = (^(v23))
                var v29 uint32 = 2
                _ = v28
                _ = v29
            }
            v23 -= (v23 << ((42) & 63))
            _ = fi25
            _ = v26
        }
        var v30 int64 = 0
        if true {
            v31 := []int{p0}
            _ = v31
        } else {
            p1 -= p1
            for fi32 := 0; fi32 < 6; fi32++ {
                v33 := []int{(-(p0)), fi32}
                if false {
                    v34 := GenS4{F0: int32(p1), F1: (-(v30)), F2: (p1 - int8(v23))}
                    fmt.Println(fi32, p1, fi32)
                    fmt.Println(fi32, v30, fi32)
                    _ = v34
                } else {
                    if (1 != 42) {
                        v30 = (v30 / ((v30) | 1))
                    }
                }
                _ = fi32
                _ = v33
            }
            switch (^(p0)) {
            case 32767:
                for fi35 := 0; fi35 < 9; fi35++ {
                    var v36 uint16 = 255
                    _ = fi35
                    _ = v36
                }
            case 148:
                if true {
                    var v37 int32 = 1
                    var v38 int = (fi24 / ((2147483647) | 1))
                    _ = v37
                    _ = v38
                } else {
                    switch (p0 - int(v30)) {
                    case -1:
                        for fi39 := 0; fi39 < 11; fi39++ {
                            var v40 uint64 = 42
                            if (0.0 < 1.5) {
                                p1 += 0
                                var v41 uint8 = 127
                                _ = v41
                            }
                            fmt.Println(fi24)
                            _ = fi39
                            _ = v40
                        }
                    default:
                        p1 += (^(p1))
                    }
                    v42 := []int{p0, (p0 ^ (p0 & (fi24 >> ((18446744073709551615) & 63)))), p0}
                    _ = len(v42)
                    _ = v42
                }
            case 963:
                v43 := []int{(fi24 ^ p0), -107, -554}
                _ = v43
            }
        }
        _ = fi24
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v23
    return p0
}

func main() {
    v1 := []int{32767, -2147483648, 128}
    v2 := make(map[string]uint32)
    v3 := []string{"hello", (("\n\t" + ("hello" + "a")) + "hello"), "a", ("hello" + "\n\t")}
    v4 := make(map[string]string)
    var v44 float32 = (((0.0 - -2.25) / 1e10) - -2.25)
    var v45 int32 = -691
    v44 -= 1.5
    if false {
        if (-31 >= 978) {
            switch -9223372036854775808 {
            case 127:
                var v46 uint64 = 42
                _ = v46
            }
            if true {
                v47 := GenS3{F0: 859}
                switch "hello" {
                case "foo bar":
                    var v48 int = 42
                    _ = v48
                case "\n\t":
                    var v49 int64 = 1
                    _ = v49
                }
                _ = v47
            } else {
                v50 := GenS2{F0: 42, F1: 0}
                var v51 int64 = -774
                v52 := GenS5{F0: (v51 >> ((4294967295) & 63)), F1: (!true)}
                _ = v50
                _ = v51
                _ = v52
            }
            var v53 bool = true
            _ = v53
        } else {
            var v54 int = 128
            for fi55 := 0; fi55 < 7; fi55++ {
                var v56 int = (v54 + 4294967295)
                _ = fi55
                _ = v56
            }
            v57 := GenS2{F0: 82, F1: 0}
            _ = v54
            _ = v57
        }
    }
    for fi58 := 0; fi58 < 19; fi58++ {
        var v59 uint32 = 4294967295
        v1 = append(v1, int(v45))
        _ = fi58
        _ = v59
    }
    switch 619 {
    case -2147483648:
        v4[("\n\t" + "a")] = (("" + "a") + "")
    case 32768:
        v44 -= v44
    case 42:
        var v60 uint32 = 65535
        _ = v60
    default:
        var v61 string = ("foo bar" + (("" + ("hello" + "\n\t")) + "hello"))
        _ = v61
    }
    var v62 string = "foo bar"
    for fi63 := 0; fi63 < 2; fi63++ {
        fmt.Println(fi63, v62, v44)
        v45 = v45
        for fi64 := 0; fi64 < 4; fi64++ {
            var v65 float32 = (1e10 / 0.1)
            v65 = 3.14159
            _ = fi64
            _ = v65
        }
        _ = fi63
    }
    if false {
        var v66 float32 = -2.25
        v67 := GenS1{F0: 65535, F1: v62}
        switch -877 {
        case 4294967295:
            var v68 uint16 = 263
            _ = v68
        case -529:
            v69 := GenS3{F0: 2147483647}
            _ = v69
        }
        _ = v66
        _ = v67
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v70 := GenS1{F0: 2, F1: v62}
    v71 := GenS5{F0: 0, F1: true}
    if (1000 == 1) {
        for fi72 := 0; fi72 < 4; fi72++ {
            var v73 int16 = 42
            _ = fi72
            _ = v73
        }
        fmt.Println(v44, v62, v44)
        v62 = v62
    }
    var v74 int8 = -1
    fmt.Println(v62, v62, v44)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v62:", v62)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v62
    _ = v70
    _ = v71
    _ = v74
}
