package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint8, p1 uint32, p2 uint8) uint16 {
    if (0 >= -9223372036854775808) {
        fmt.Println(p0)
        p2 = (p0 + p0)
    }
    v7 := GenS1{F0: p1, F1: -128, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}
func genFn1(p0 uint8, p1 string) int32 {
    for fi8 := 0; fi8 < 5; fi8++ {
        p0 -= p0
        var v9 int16 = -558
        var v10 bool = (1e10 <= float32(float64(-2.25)))
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 float64 = 3.14159
        v13 := []int{(fi11 * (fi11 + fi11)), fi11, (fi11 / ((fi11) | 1)), (-(fi11))}
        _ = fi11
        _ = v12
        _ = v13
    }
    switch ("hello" + p1) {
    case "hello":
        if true {
            v14 := []int{127, 0}
            p0 -= 255
            _ = v14
        } else {
            v15 := []int{9223372036854775807, 884}
            p1 = p1
            v16 := GenS2{F0: 0}
            _ = v15
            _ = v16
        }
    case "foo bar":
        v17 := []int{-1, 32768, -2147483648}
        _ = v17
    case "":
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 uint = 253
            if (42 >= p0) {
                switch (p1 + (p1 + (p1 + "foo bar"))) {
                case "a":
                    fmt.Println(fi18, p0)
                default:
                    if true {
                        for fi20 := 0; fi20 < 3; fi20++ {
                            if true {
                                switch (fi18 % ((fi20) | 1)) {
                                case 1:
                                    if false {
                                        v21 := GenS1{F0: 2147483648, F1: -122, F2: -1}
                                        var v22 string = p1
                                        v23 := []int{fi18, fi20}
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    } else {
                                        if false {
                                            if false {
                                                switch p1 {
                                                case "a":
                                                    p0 = p0
                                                case "\n\t":
                                                    p1 = "\n\t"
                                                default:
                                                    var v24 string = (p1 + (("hello" + ("foo bar" + p1)) + (p1 + "a")))
                                                    _ = v24
                                                }
                                            }
                                            p0 += p0
                                        } else {
                                            var v25 string = ""
                                            var v26 int = -312
                                            _ = v25
                                            _ = v26
                                        }
                                        fmt.Println(v19, p0, p0)
                                        switch ((("\n\t" + "\n\t") + "hello") + p1) {
                                        case "foo bar":
                                            var v27 uint64 = 533
                                            _ = v27
                                        case "\n\t":
                                            switch p1 {
                                            case "foo bar":
                                                p1 = ("hello" + "\n\t")
                                            case "\n\t":
                                                var v28 uint64 = 869
                                                _ = v28
                                            case "a":
                                                v29 := []int{(^(fi18))}
                                                _ = v29
                                            default:
                                                var v30 float32 = 0.0
                                                _ = v30
                                            }
                                        }
                                    }
                                case -1:
                                    var v31 int32 = -1
                                    _ = v31
                                default:
                                    v32 := GenS3{F0: 4294967295, F1: "hello"}
                                    _ = v32
                                }
                            } else {
                                v33 := []int{fi18, fi18}
                                _ = v33
                            }
                            var v34 int16 = -956
                            v35 := []int{(^(fi18))}
                            switch p1 {
                            case "\n\t":
                                if true {
                                    fmt.Println(fi20)
                                    p1 = p1
                                    var v36 int32 = -104
                                    fmt.Println(v34, v34, fi20)
                                    _ = v36
                                }
                            }
                            p1 = p1
                            _ = fi20
                            _ = v34
                            _ = v35
                        }
                        var v37 int = fi18
                        _ = v37
                    }
                }
            }
            _ = fi18
            _ = v19
        }
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 uint16 = 65535
    var v2 bool = (true && true)
    v3 := []uint32{755}
    v4 := make(map[uint32]string)
    v5 := []int{0, -1, 32767}
    v6 := make(map[string]string)
    v38 := GenS2{F0: genFn1(128, "")}
    fmt.Println(v1, v2)
    var v39 uint = 65535
    fmt.Println(v2)
    if v2 {
        var v40 int16 = -215
        var v41 int32 = -157
        _ = v40
        _ = v41
    }
    var v42 int64 = 42
    v1 -= 65535
    var v43 uint = (v39 - 255)
    for fi44 := 0; fi44 < 5; fi44++ {
        v45 := GenS4{F0: fi44}
        _ = fi44
        _ = v45
    }
    v43 += 544
    switch "" {
    case "":
        for fi46 := 0; fi46 < 3; fi46++ {
            for fi47 := 0; fi47 < 17; fi47++ {
                var v48 bool = v2
                var v49 int8 = 100
                _ = fi47
                _ = v48
                _ = v49
            }
            fmt.Println(v39)
            _ = fi46
        }
    case "foo bar":
        for fi50 := 0; fi50 < 1; fi50++ {
            fmt.Println(v43, v43)
            if ((v1 | 2) >= 170) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            v2 = (fi50 >= 42)
            _ = fi50
        }
    }
    v1 = (v1 % ((uint16(v43)) | 1))
    if v2 {
        v51 := v4[2147483647]
        if v2 {
            fmt.Println(v43)
        } else {
            v1 = v1
        }
        _ = v51
    }
    for fi52 := 0; fi52 < 11; fi52++ {
        if v2 {
            switch "hello" {
            case "\n\t":
                v43 = v43
            case "":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "hello":
                for fi53 := 0; fi53 < 8; fi53++ {
                    v54 := GenS5{F0: "", F1: (255 == 77)}
                    v6["hello"] = ("" + ("hello" + "foo bar"))
                    _ = fi53
                    _ = v54
                }
            }
        }
        _ = fi52
    }
    for fi55 := 0; fi55 < 6; fi55++ {
        v43 += v43
        v56 := GenS4{F0: 9223372036854775807}
        _ = fi55
        _ = v56
    }
    var v57 int = -9223372036854775808
    switch -666 {
    case 139:
        var v58 uint = (v39 << ((2) & 63))
        _ = v58
    case -694:
        var v59 int32 = -2147483648
        _ = v59
    }
    switch (^(v57)) {
    case 4294967295:
        var v60 float32 = 1e10
        _ = v60
    case 9223372036854775807:
        v61 := GenS5{F0: ("a" + ""), F1: (100 < 1)}
        _ = v61
    default:
        switch "\n\t" {
        case "hello":
            if false {
                v62 := GenS6{F0: 65535, F1: 255}
                var v63 uint32 = 768
                _ = v62
                _ = v63
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v57
    _ = v6
}
