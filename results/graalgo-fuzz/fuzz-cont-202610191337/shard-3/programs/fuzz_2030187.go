package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint32 }
func genFn0() uint64 {
    if (0 == 837) {
        var v6 int64 = 9223372036854775807
        _ = v6
    } else {
        var v7 uint16 = 65535
        v7 = (v7 * 2)
        for fi8 := 0; fi8 < 19; fi8++ {
            if (false && true) {
                if true {
                    var v9 int16 = -110
                    v9 = v9
                    if false {
                        v10 := GenS1{F0: 32768}
                        _ = v10
                    }
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 13; fi11++ {
                        if ((("a" + "") + "hello") <= (("a" + ((((((("foo bar" + "") + "\n\t") + "foo bar") + "a") + "a") + "foo bar") + "\n\t")) + "a")) {
                            v12 := []int{-231}
                            _ = v12
                        } else {
                            fmt.Println(fi8)
                            switch "foo bar" {
                            case "foo bar":
                                switch "foo bar" {
                                case "\n\t":
                                    if true {
                                        v7 = v7
                                        var v13 int8 = 0
                                        if false {
                                            switch (fi8 - fi8) {
                                            case -775:
                                                v14 := GenS1{F0: fi8}
                                                _ = v14
                                            case -2147483648:
                                                switch fi11 {
                                                case -556:
                                                    v15 := GenS1{F0: fi8}
                                                    _ = v15
                                                }
                                            }
                                        }
                                        _ = v13
                                    } else {
                                        v7 = v7
                                        var v16 int32 = 165
                                        _ = v16
                                    }
                                case "foo bar":
                                    v7 = v7
                                case "":
                                    var v17 bool = ("\n\t" >= "foo bar")
                                    _ = v17
                                }
                            case "a":
                                v18 := GenS1{F0: 42}
                                _ = v18
                            }
                            v19 := []int{(^(fi8))}
                            _ = v19
                        }
                        v7 = v7
                        v20 := GenS1{F0: 0}
                        _ = fi11
                        _ = v20
                    }
                }
            } else {
                var v21 int16 = 32767
                _ = v21
            }
            var v22 int64 = -9223372036854775808
            for fi23 := 0; fi23 < 20; fi23++ {
                if false {
                    var v24 int64 = v22
                    _ = v24
                } else {
                    v7 = v7
                    v7 = v7
                }
                v25 := GenS1{F0: (fi23 - fi8)}
                if (127 >= 69) {
                    v26 := GenS2{F0: 44}
                    v22 = 0
                    if (((false && (2 > 0)) && (!(v7 >= v7))) || false) {
                        v27 := []int{fi23, 128}
                        v7 = (^(v7))
                        v28 := GenS2{F0: 1}
                        _ = v27
                        _ = v28
                    } else {
                        v29 := GenS2{F0: 75}
                        _ = v29
                    }
                    var v30 int64 = v22
                    _ = v26
                    _ = v30
                }
                if false {
                    var v31 int32 = 0
                    v32 := []int{9223372036854775807}
                    var v33 bool = (true && true)
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = fi23
                _ = v25
            }
            _ = fi8
            _ = v22
        }
        _ = v7
    }
    var v34 int = 1
    v35 := []int{v34, v34, 128, (^(v34))}
    for fi36 := 0; fi36 < 7; fi36++ {
        fmt.Println(fi36, v34)
        _ = fi36
    }
    _ = v34
    _ = v35
    return 4294967295
}
func genFn1(p0 bool, p1 uint64) uint {
    fmt.Println(p1, p0)
    fmt.Println(p1, p0)
    p0 = true
    var v37 int = -179
    _ = p0
    _ = p1
    _ = v37
    return 362
}

func main() {
    var v1 int8 = 112
    v2 := make(map[string]uint32)
    var v3 int16 = -32768
    v4 := make(map[string]int)
    var v5 bool = (!(("a" + ((("\n\t" + "") + "a") + "hello")) <= ""))
    v5 = v5
    for fi38 := 0; fi38 < 16; fi38++ {
        switch (fi38 + (fi38 >> ((203) & 63))) {
        case -443:
            if (1 <= 1) {
                v39 := GenS3{F0: 4294967295, F1: fi38}
                _ = v39
            }
        case 9223372036854775807:
            var v40 float64 = float64(1.5)
            _ = v40
        default:
            for fi41 := 0; fi41 < 12; fi41++ {
                v42 := GenS4{F0: 1, F1: 270}
                var v43 int32 = -2147483648
                _ = fi41
                _ = v42
                _ = v43
            }
        }
        for fi44 := 0; fi44 < 15; fi44++ {
            var v45 uint32 = 42
            _ = fi44
            _ = v45
        }
        var v46 int8 = (v1 % ((int8(v3)) | 1))
        _ = fi38
        _ = v46
    }
    v47 := GenS3{F0: 18446744073709551615, F1: 396}
    if (0 == -1) {
        v4[(("" + "hello") + "")] = 2147483647
    } else {
        delete(v4, "foo bar")
        var v48 uint = 1
        v49 := GenS4{F0: 0, F1: 304}
        _ = v48
        _ = v49
    }
    var v50 float64 = -2.25
    v51 := GenS2{F0: 128}
    var v52 int = -513
    v53 := GenS4{F0: 803, F1: 2147483648}
    for fi54 := 0; fi54 < 18; fi54++ {
        _ = len(v4)
        switch -1 {
        case 42:
            if v5 {
                var v55 float32 = (1e10 + 0.0)
                fmt.Println(v50)
                _ = v55
            } else {
                v5 = false
            }
        default:
            v1 = (v1 % (((v1 >> ((796) & 63))) | 1))
        }
        _ = fi54
    }
    var v56 uint32 = 147
    var v57 uint = 757
    fmt.Println(v3, v57)
    v58 := GenS5{F0: 0}
    v59 := GenS6{F0: genFn0(), F1: v56, F2: (v56 + (v56 << (((v57 | 1)) & 63)))}
    if (2 > 76) {
        var v60 uint8 = 232
        var v61 string = ("a" + ("\n\t" + "hello"))
        for fi62 := 0; fi62 < 20; fi62++ {
            v60 = (^(v60))
            _ = fi62
        }
        _ = v60
        _ = v61
    } else {
        fmt.Println(v3)
        v63 := []int{}
        v5 = true
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
