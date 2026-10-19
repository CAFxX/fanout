package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int; F2 int16 }
type GenS4 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 int16; F1 int8; F2 int32 }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 uint; F1 int64 }
type GenS11 struct { F0 uint8; F1 int8; F2 int64 }
func genFn0(p0 int, p1 uint8) int8 {
    var v4 string = ("\n\t" + ("hello" + "foo bar"))
    fmt.Println(v4)
    var v5 uint16 = 32767
    v6 := GenS1{F0: 170}
    if (false || false) {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 int64 = -216
            var v9 uint16 = v5
            if (3.14159 == 0.1) {
                var v10 uint8 = p1
                var v11 float64 = 0.1
                switch (v4 + "a") {
                case "":
                    for fi12 := 0; fi12 < 15; fi12++ {
                        fmt.Println(fi7)
                        var v13 float32 = (((0.1 - 1e10) + (0.1 - 1e10)) + 0.0)
                        _ = fi12
                        _ = v13
                    }
                case "foo bar":
                    v14 := []int{fi7, fi7, (-(p0))}
                    _ = v14
                case "\n\t":
                    switch fi7 {
                    case 0:
                        v9 += v9
                    case 2147483647:
                        var v15 string = "foo bar"
                        _ = v15
                    case -827:
                        var v16 int8 = 127
                        _ = v16
                    default:
                        var v17 string = "hello"
                        _ = v17
                    }
                default:
                    for fi18 := 0; fi18 < 2; fi18++ {
                        v9 += 255
                        p0 = 2147483647
                        v5 = v5
                        _ = fi18
                    }
                }
                _ = v10
                _ = v11
            } else {
                v19 := []int{(p0 ^ fi7), p0, int(v9)}
                fmt.Println(p1, v8, v9)
                _ = v19
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        if true {
            var v20 int32 = -2147483648
            v21 := GenS2{F0: 0, F1: -77, F2: 196}
            v22 := GenS3{F0: 2, F1: 42, F2: int16(p1)}
            _ = v20
            _ = v21
            _ = v22
        }
    } else {
        for fi23 := 0; fi23 < 20; fi23++ {
            fmt.Println(v4, p0, v5)
            v24 := GenS4{F0: uint8(v5), F1: v5, F2: -128}
            _ = fi23
            _ = v24
        }
        if true {
            p0 += p0
        } else {
            p1 = p1
            for fi25 := 0; fi25 < 16; fi25++ {
                switch fi25 {
                case 2147483647:
                    switch v4 {
                    case "\n\t":
                        if (false || (645 >= 42)) {
                            switch v4 {
                            case "foo bar":
                                for fi26 := 0; fi26 < 19; fi26++ {
                                    var v27 uint16 = 0
                                    v28 := GenS3{F0: 0, F1: -24, F2: 32767}
                                    v29 := GenS4{F0: 42, F1: v5, F2: 42}
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                            case "a":
                                if true {
                                    var v30 uint16 = (^(v5))
                                    _ = v30
                                } else {
                                    var v31 int16 = -32768
                                    v32 := []int{p0}
                                    _ = v31
                                    _ = v32
                                }
                            }
                        } else {
                            var v33 uint16 = (v5 >> ((113) & 63))
                            _ = v33
                        }
                    case "hello":
                        v34 := []int{-11, 1, p0, p0}
                        _ = v34
                    default:
                        v5 -= 255
                    }
                case -976:
                    if (!((1 <= -867) || (3.14159 != (-2.25 / 0.1)))) {
                        for fi35 := 0; fi35 < 11; fi35++ {
                            switch (fi35 - fi25) {
                            case -9223372036854775808:
                                fmt.Println(p1, p1, p1)
                            case 485:
                                p1 += p1
                            case 210:
                                if (1e10 <= 1.5) {
                                    var v36 uint64 = 0
                                    var v37 float32 = (3.14159 + float32(1.5))
                                    _ = v36
                                    _ = v37
                                }
                            default:
                                v38 := []int{(fi35 % ((9223372036854775807) | 1))}
                                _ = v38
                            }
                            fmt.Println(p0, fi35)
                            var v39 int = p0
                            _ = fi35
                            _ = v39
                        }
                    } else {
                        var v40 float32 = -2.25
                        v41 := GenS5{F0: (true || true), F1: 4294967295}
                        fmt.Println(v5)
                        _ = v40
                        _ = v41
                    }
                case 32768:
                    v42 := []int{p0, 787}
                    _ = v42
                default:
                    switch "" {
                    case "":
                        p0 = (p0 * (-(p0)))
                    default:
                        var v43 int = (-(p0))
                        _ = v43
                    }
                }
                v44 := GenS6{F0: 4294967295}
                _ = fi25
                _ = v44
            }
        }
        if true {
            fmt.Println(p0, p1, p0)
            v5 = 80
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int64) uint64 {
    v45 := GenS6{F0: 4294967295}
    v46 := GenS7{F0: 65535, F1: 877}
    v47 := GenS8{F0: -32768, F1: 42, F2: 0}
    var v48 bool = false
    if v48 {
        if v48 {
            var v49 string = ""
            fmt.Println(v48, p0)
            var v50 int16 = 0
            _ = v49
            _ = v50
        }
    }
    _ = p0
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    return 2
}

func main() {
    var v1 int8 = 42
    var v2 int8 = (v1 << ((255) & 63))
    v3 := make(map[string]int)
    v51 := GenS6{F0: 2}
    switch "a" {
    case "\n\t":
        if true {
            v52 := v3["foo bar"]
            if (0 >= 59) {
                v52 = -883
            }
            _ = v52
        }
    case "hello":
        switch ("\n\t" + ("hello" + ("hello" + (("\n\t" + "foo bar") + ("foo bar" + "a"))))) {
        case "foo bar":
            if true {
                for fi53 := 0; fi53 < 16; fi53++ {
                    v54 := v3["hello"]
                    _ = fi53
                    _ = v54
                }
                var v55 string = "hello"
                v56 := GenS4{F0: 77, F1: 42, F2: -1}
                _ = v55
                _ = v56
            } else {
                fmt.Println(v2, v2)
            }
        default:
            var v57 string = "foo bar"
            _ = v57
        }
    }
    v2 = (v1 - 42)
    switch 326 {
    case 9223372036854775807:
        v1 += 28
    default:
        var v58 int64 = 42
        _ = v58
    }
    for fi59 := 0; fi59 < 11; fi59++ {
        v60 := []int{fi59, fi59}
        var v61 int64 = -1
        _ = fi59
        _ = v60
        _ = v61
    }
    switch -2147483648 {
    case -12:
        v1 -= genFn0(1, 2)
    }
    for fi62 := 0; fi62 < 16; fi62++ {
        var v63 bool = true
        _ = fi62
        _ = v63
    }
    switch -922 {
    case -883:
        v2 = genFn0(-2147483648, 195)
    default:
        fmt.Println(v2, v1)
    }
    var v64 int32 = 2147483647
    var v65 int8 = v2
    _ = len(v3)
    if false {
        v65 -= v65
        v66 := GenS3{F0: 2, F1: -513, F2: 890}
        _ = v66
    } else {
        v2 = (v2 ^ (-(v65)))
        if (255 > 238) {
            v67 := GenS5{F0: true, F1: 0}
            _ = v67
        } else {
            v3[(((("\n\t" + (("foo bar" + "foo bar") + "\n\t")) + "\n\t") + (("\n\t" + "foo bar") + "a")) + "\n\t")] = -690
        }
    }
    if (!(0.1 >= (0.1 + 1e10))) {
        v68 := GenS1{F0: 1}
        v69 := GenS9{F0: 255, F1: 4294967295}
        var v70 int16 = 0
        _ = v68
        _ = v69
        _ = v70
    } else {
        v71 := GenS3{F0: 4294967295, F1: int(v64), F2: 1}
        var v72 float32 = ((0.0 / float32(0.1)) / (1.5 * -2.25))
        v73 := GenS10{F0: 537, F1: -597}
        _ = v71
        _ = v72
        _ = v73
    }
    var v74 string = (("foo bar" + ("hello" + "\n\t")) + ("foo bar" + "foo bar"))
    var v75 uint32 = 2147483647
    v76 := GenS11{F0: 128, F1: genFn0(1, 255), F2: -993}
    var v77 int8 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
    _ = v64
    _ = v65
    _ = v74
    _ = v75
    _ = v76
    _ = v77
}
