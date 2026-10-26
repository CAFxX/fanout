package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int8; F2 string }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int16; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint32; F1 int; F2 uint }
type GenS8 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS9 struct { F0 uint32 }
func genFn0() int64 {
    if true {
        var v6 bool = ((!true) || true)
        for fi7 := 0; fi7 < 6; fi7++ {
            if (!(0 <= 42)) {
                v8 := []int{(fi7 >> ((18446744073709551615) & 63)), -1, fi7, (fi7 - (fi7 / ((fi7) | 1)))}
                if (442 >= 1) {
                    var v9 uint16 = 1
                    if v6 {
                        for fi10 := 0; fi10 < 10; fi10++ {
                            if (!v6) {
                                var v11 uint16 = (v9 << ((763) & 63))
                                if v6 {
                                    v8 = append(v8, fi10)
                                    var v12 int32 = 1
                                    var v13 int16 = 42
                                    _ = v12
                                    _ = v13
                                } else {
                                    fmt.Println(v11)
                                }
                                if len(v8) > 0 {
                                    v8[0] = -626
                                }
                                fmt.Println(v6)
                                var v14 int = (fi7 ^ fi10)
                                _ = v11
                                _ = v14
                            } else {
                                fmt.Println(fi7)
                                var v15 string = "a"
                                var v16 int8 = 1
                                v8 = append(v8, fi7)
                                v17 := GenS1{F0: 1}
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            _ = fi10
                        }
                        var v18 uint32 = 0
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v20 := GenS1{F0: 0}
                        switch fi7 {
                        case -1:
                            v8 = append(v8, 4294967295)
                        case 2147483647:
                            v6 = v6
                        default:
                            var v21 float32 = -2.25
                            _ = v21
                        }
                        _ = fi19
                        _ = v20
                    }
                    _ = v9
                } else {
                    v22 := GenS2{F0: ""}
                    for fi23 := 0; fi23 < 16; fi23++ {
                        fmt.Println(fi23, fi7, fi7)
                        _ = fi23
                    }
                    _ = v22
                }
                if true {
                    fmt.Println(v6, fi7)
                }
                _ = v8
            }
            _ = fi7
        }
        _ = v6
    }
    v24 := GenS3{F0: 186, F1: 121, F2: "foo bar"}
    for fi25 := 0; fi25 < 4; fi25++ {
        switch "hello" {
        case "":
            var v26 uint64 = 18446744073709551615
            _ = v26
        default:
            if (-128 < 127) {
                var v27 int16 = 1
                if true {
                    v28 := []int{fi25, (fi25 ^ (fi25 & (fi25 * (fi25 - 4294967295)))), (^(fi25)), fi25}
                    if false {
                        fmt.Println(v27, fi25)
                    }
                    fmt.Println(v27, fi25)
                    _ = v28
                } else {
                    if ((2 > 1) && true) {
                        v29 := []int{(fi25 / ((fi25) | 1)), fi25, fi25, (^(fi25))}
                        v30 := GenS4{F0: false, F1: v27}
                        v31 := GenS4{F0: (0 > 9223372036854775807), F1: (v27 + -1)}
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        v27 += (v27 << ((42) & 63))
                        fmt.Println(fi25)
                        var v32 int8 = 1
                        _ = v32
                    }
                }
                _ = v27
            } else {
                var v33 string = "hello"
                v34 := GenS5{F0: 114, F1: -126, F2: 1}
                var v35 int8 = -1
                _ = v33
                _ = v34
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 13; fi36++ {
            var v37 uint16 = 42
            v38 := []int{(fi36 / (((fi25 - (fi36 / ((fi36) | 1)))) | 1)), 906}
            v39 := GenS6{F0: (^(v37)), F1: "a"}
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        var v40 uint32 = 299
        _ = fi25
        _ = v40
    }
    switch "" {
    case "a":
        if false {
            var v41 uint = 18446744073709551615
            if true {
                v42 := GenS7{F0: 42, F1: 0, F2: v41}
                if false {
                    var v43 bool = false
                    var v44 int16 = -591
                    v45 := GenS1{F0: v41}
                    _ = v43
                    _ = v44
                    _ = v45
                } else {
                    var v46 int32 = -809
                    v41 -= (v41 / ((v41) | 1))
                    switch -2147483648 {
                    case 32768:
                        v47 := GenS8{F0: 0, F1: 2, F2: 42}
                        _ = v47
                    case -780:
                        var v48 string = ("\n\t" + ("foo bar" + (("foo bar" + "a") + (("foo bar" + ((((("" + "\n\t") + "a") + "") + "hello") + "foo bar")) + "foo bar"))))
                        _ = v48
                    case 42:
                        var v49 bool = false
                        _ = v49
                    }
                    _ = v46
                }
                v50 := []int{36, -2147483648, 127}
                _ = v42
                _ = v50
            }
            fmt.Println(v41)
            _ = v41
        }
    case "\n\t":
        var v51 uint16 = 32767
        _ = v51
    case "":
        v52 := []int{915, 32767, -2147483648, 1}
        _ = v52
    default:
        v53 := []int{-942}
        _ = v53
    }
    _ = v24
    return 1
}

func main() {
    v1 := []string{("hello" + "hello"), "hello"}
    var v2 uint16 = 125
    var v3 int8 = 42
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    fmt.Println(v2)
    fmt.Println(v3)
    fmt.Println(v2, v3)
    v54 := GenS9{F0: 0}
    if (65535 <= 749) {
        var v55 int8 = v3
        var v56 float64 = -2.25
        var v57 uint16 = (v2 << ((1) & 63))
        _ = v55
        _ = v56
        _ = v57
    } else {
        var v58 uint = 65535
        if (56 <= 2147483647) {
            if true {
                switch "hello" {
                case "foo bar":
                    fmt.Println(v3)
                }
            }
            v58 = v58
        } else {
            delete(v5, "hello")
            var v59 int16 = 1000
            _ = v59
        }
        var v60 float64 = (1e10 * (0.0 - 1.5))
        _ = v58
        _ = v60
    }
    v2 -= v2
    v2 = v2
    fmt.Println(v3, v3)
    v2 = (v2 + 1)
    for fi61 := 0; fi61 < 11; fi61++ {
        v62 := GenS2{F0: "hello"}
        if false {
            v3 = (v3 - (v3 / ((v3) | 1)))
            var v63 bool = false
            _ = v63
        }
        _ = fi61
        _ = v62
    }
    v3 += v3
    var v64 int = -1
    fmt.Println(v2, v2, v64)
    var v65 float64 = -2.25
    v1 = append(v1, "foo bar")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v64
    _ = v65
}
