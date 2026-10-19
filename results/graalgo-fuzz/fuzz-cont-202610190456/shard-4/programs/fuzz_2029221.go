package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int; F2 int16 }
type GenS4 struct { F0 int; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 string; F1 int32; F2 int16 }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int32, p1 int16) int {
    var v5 bool = true
    var v6 bool = ((-2.25 - (0.1 * 1.5)) != float32(3.14159))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 9
}
func genFn1() uint32 {
    switch "a" {
    case "foo bar":
        v7 := GenS1{F0: -32768}
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 5; fi8++ {
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := GenS2{F0: 0, F1: fi9, F2: 0}
                v11 := GenS1{F0: 260}
                _ = fi9
                _ = v10
                _ = v11
            }
            fmt.Println(fi8)
            _ = fi8
        }
    case "a":
        switch "\n\t" {
        case "a":
            var v12 uint8 = 246
            _ = v12
        case "\n\t":
            var v13 int32 = -838
            _ = v13
        case "hello":
            switch -78 {
            case 128:
                var v14 bool = (1 >= 2)
                _ = v14
            case 4294967295:
                var v15 int64 = 9223372036854775807
                _ = v15
            case 127:
                switch 128 {
                case 455:
                    var v16 int = genFn0(-576, 605)
                    _ = v16
                case 32767:
                    switch 127 {
                    case 1:
                        var v17 float32 = 1.5
                        _ = v17
                    case -983:
                        var v18 uint16 = 42
                        _ = v18
                    }
                case -9223372036854775808:
                    v19 := []int{}
                    _ = v19
                default:
                    var v20 bool = ("\n\t" >= (((("a" + "hello") + "") + "") + ("\n\t" + ("" + "a"))))
                    _ = v20
                }
            }
        default:
            var v21 int8 = -3
            _ = v21
        }
    }
    if (true || (42 != 0)) {
        v22 := GenS3{F0: 1, F1: -811, F2: 0}
        var v23 uint64 = 18446744073709551615
        _ = v22
        _ = v23
    } else {
        var v24 int8 = 100
        var v25 float64 = ((0.0 / 1.0) * 0.0)
        if true {
            for fi26 := 0; fi26 < 17; fi26++ {
                switch ("hello" + "hello") {
                case "":
                    if false {
                        for fi27 := 0; fi27 < 4; fi27++ {
                            v25 += v25
                            if (!(1.5 != 0.0)) {
                                if true {
                                    v28 := []int{(fi26 % ((fi27) | 1))}
                                    v25 = v25
                                    _ = v28
                                }
                                if true {
                                    var v29 uint = 1
                                    v24 = 1
                                    v29 += v29
                                    _ = v29
                                } else {
                                    switch "\n\t" {
                                    case "foo bar":
                                        if true {
                                            if false {
                                                var v30 uint = 65535
                                                _ = v30
                                            }
                                        }
                                        v31 := GenS4{F0: genFn0(-1, 42), F1: 2, F2: 65535}
                                        v24 = -1
                                        _ = v31
                                    }
                                    v32 := GenS3{F0: 175, F1: 404, F2: 893}
                                    v33 := GenS1{F0: 596}
                                    _ = v32
                                    _ = v33
                                }
                                v34 := []int{genFn0(1, 32767), (fi27 + fi27), 1}
                                _ = v34
                            }
                            v35 := GenS1{F0: 32767}
                            _ = fi27
                            _ = v35
                        }
                        v24 = (v24 - v24)
                    } else {
                        var v36 int = fi26
                        v37 := GenS5{F0: 812}
                        v25 += v25
                        _ = v36
                        _ = v37
                    }
                default:
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(fi26, v25, v25)
                    default:
                        if (fi26 != (fi26 - 9223372036854775807)) {
                            v25 = v25
                            for fi38 := 0; fi38 < 4; fi38++ {
                                v25 = float64(1e10)
                                v25 = v25
                                if (42 < -1) {
                                    var v39 float32 = 3.14159
                                    switch "foo bar" {
                                    case "\n\t":
                                        v40 := GenS1{F0: 578}
                                        _ = v40
                                    default:
                                        fmt.Println(v39, v24, v39)
                                    }
                                    fmt.Println(v39, v25)
                                    _ = v39
                                }
                                _ = fi38
                            }
                        } else {
                            v25 = (0.0 - v25)
                            switch "foo bar" {
                            case "a":
                                var v41 uint = 2
                                _ = v41
                            }
                        }
                    }
                }
                switch ("\n\t" + ("hello" + (("foo bar" + "a") + "foo bar"))) {
                case "foo bar":
                    v42 := []int{(fi26 & -2147483648)}
                    _ = v42
                case "hello":
                    v24 -= (^(v24))
                case "\n\t":
                    var v43 float32 = 0.1
                    _ = v43
                }
                var v44 uint8 = 128
                _ = fi26
                _ = v44
            }
        } else {
            v25 = (v25 / float64(-2.25))
            for fi45 := 0; fi45 < 9; fi45++ {
                v46 := GenS2{F0: 32767, F1: (fi45 + (fi45 >> ((217) & 63))), F2: 0}
                v47 := []int{(fi45 >> ((129) & 63))}
                v47 = append(v47, fi45)
                _ = fi45
                _ = v46
                _ = v47
            }
        }
        _ = v24
        _ = v25
    }
    var v48 int8 = 42
    v49 := GenS6{F0: 0, F1: 0}
    var v50 uint32 = 306
    _ = v48
    _ = v49
    _ = v50
    return v50
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]bool)
    var v3 int = -9223372036854775808
    var v4 int8 = 1
    v51 := GenS7{F0: "a", F1: 1, F2: -212}
    switch 32767 {
    case -9223372036854775808:
        v52 := GenS8{F0: 410, F1: 42}
        _ = v52
    default:
        v1 = 0.0
    }
    v1 -= float32(0.0)
    var v53 int = 127
    var v54 uint64 = 0
    for fi55 := 0; fi55 < 17; fi55++ {
        fmt.Println(v4, v1, v1)
        v56 := GenS3{F0: 117, F1: v3, F2: 32767}
        _ = fi55
        _ = v56
    }
    v1 = 3.14159
    v3 -= v3
    var v57 uint32 = 1
    switch v3 {
    case 1:
        for fi58 := 0; fi58 < 4; fi58++ {
            v59 := GenS9{F0: 31, F1: uint16(v57)}
            v60 := []int{genFn0(-666, -997), (fi58 & int(v54))}
            _ = fi58
            _ = v59
            _ = v60
        }
    case -9223372036854775808:
        v57 -= 1
    }
    switch ("foo bar" + ((("hello" + ("foo bar" + "hello")) + "a") + "a")) {
    case "\n\t":
        switch (-(v3)) {
        case 4294967295:
            fmt.Println(v57)
        case -1:
            fmt.Println(v57, v1)
        case 128:
            v53 = genFn0(-2147483648, -1)
        default:
            v61 := []int{(v53 / ((v3) | 1)), (v53 - genFn0(-2147483648, 1))}
            _ = v61
        }
    default:
        var v62 uint16 = 0
        _ = v62
    }
    v3 -= (v3 * (v3 >> ((0) & 63)))
    var v63 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v53
    _ = v54
    _ = v57
    _ = v63
}
