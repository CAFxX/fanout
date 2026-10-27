package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16; F1 uint32 }
func genFn0() uint8 {
    var v6 uint64 = 4294967295
    var v7 int64 = 42
    v8 := []int{45, 0}
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int16, p1 int32) uint64 {
    for fi9 := 0; fi9 < 17; fi9++ {
        p0 = p0
        _ = fi9
    }
    var v10 int = 42
    if false {
        var v11 float64 = 1e10
        fmt.Println(p1)
        var v12 int16 = (p0 >> ((0) & 63))
        _ = v11
        _ = v12
    } else {
        v13 := GenS1{F0: "a"}
        v14 := GenS2{F0: 4294967295, F1: false}
        v15 := []int{}
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v10
    _ = v16
    return 299
}

func main() {
    v1 := []uint32{2147483647, 624, 65535}
    var v2 uint = 255
    var v3 string = "hello"
    var v4 uint32 = 42
    v5 := make(map[string]int)
    v17 := GenS1{F0: ((((v3 + "hello") + "foo bar") + ("\n\t" + "\n\t")) + "foo bar")}
    v18 := v5["a"]
    v18 = v18
    if (42 < 739) {
        v19 := GenS1{F0: (v3 + ("hello" + ("hello" + "hello")))}
        if true {
            var v20 int8 = 42
            _ = v20
        }
        _ = v19
    }
    var v21 int8 = 1
    v18 = -463
    for fi22 := 0; fi22 < 8; fi22++ {
        fmt.Println(v18, v18)
        if (127 > 1) {
            switch (v18 - (-(fi22))) {
            case 4294967295:
                switch (v18 ^ 9223372036854775807) {
                case -542:
                    v21 += (v21 % (((v21 % ((v21) | 1))) | 1))
                case 1:
                    fmt.Println(v2, v3, v21)
                default:
                    switch (v18 - v18) {
                    case 128:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case -1:
                        if (v18 <= (v18 - (fi22 & 2147483647))) {
                            for fi23 := 0; fi23 < 10; fi23++ {
                                v24 := GenS3{F0: -1}
                                var v25 uint64 = genFn1(-511, int32(v18))
                                var v26 uint = 42
                                _ = fi23
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                            var v27 int64 = 0
                            _ = v27
                        } else {
                            v28 := GenS4{F0: 1, F1: 1000}
                            if true {
                                var v29 int8 = (v21 + v21)
                                v21 += v21
                                if (true && (uint16(v2) != 32767)) {
                                    v30 := GenS5{F0: v18}
                                    v29 += (v29 % ((v21) | 1))
                                    _ = v30
                                } else {
                                    v29 = -3
                                    if false {
                                        var v31 string = "foo bar"
                                        fmt.Println(v18)
                                        _ = v31
                                    } else {
                                        var v32 uint8 = 0
                                        _ = v32
                                    }
                                    v3 = v3
                                }
                                _ = v29
                            } else {
                                if true {
                                    v3 = ("" + ((v3 + ("a" + "\n\t")) + ((v3 + (v3 + ("a" + v3))) + (v3 + ("\n\t" + "\n\t")))))
                                    v4 = (v4 * v4)
                                } else {
                                    v33 := v5[(v3 + ("a" + (("" + "") + "")))]
                                    _ = v33
                                }
                            }
                            _ = v28
                        }
                    case -206:
                        var v34 uint64 = genFn1(-975, 1)
                        _ = v34
                    default:
                        var v35 int8 = int8(v4)
                        _ = v35
                    }
                }
            case -956:
                if (!(!true)) {
                    v2 -= 1
                    if (uint(v4) == v2) {
                        switch 9223372036854775807 {
                        case 42:
                            v2 -= 255
                        }
                    } else {
                        fmt.Println(v18)
                        fmt.Println(v4, fi22)
                    }
                    switch fi22 {
                    case 0:
                        v1 = append(v1, 1)
                    case -932:
                        for fi36 := 0; fi36 < 14; fi36++ {
                            var v37 float32 = -2.25
                            switch (v3 + v3) {
                            case "":
                                v38 := v5[("" + "a")]
                                var v39 uint8 = 42
                                _ = v38
                                _ = v39
                            case "hello":
                                v40 := GenS4{F0: 984, F1: 956}
                                _ = v40
                            case "foo bar":
                                var v41 int = (fi22 >> (((v2 << (((v2 - 65535)) & 63))) & 63))
                                _ = v41
                            }
                            var v42 string = ("\n\t" + v3)
                            _ = fi36
                            _ = v37
                            _ = v42
                        }
                    case 546:
                        fmt.Println(v21, fi22, v3)
                    }
                }
            }
            var v43 uint8 = 42
            v44 := GenS5{F0: int(v43)}
            _ = v43
            _ = v44
        }
        for fi45 := 0; fi45 < 17; fi45++ {
            var v46 int32 = 1
            if false {
                var v47 int32 = v46
                v2 -= v2
                _ = v47
            } else {
                var v48 string = v3
                v48 = "foo bar"
                var v49 int32 = 2147483647
                _ = v48
                _ = v49
            }
            v46 += -2147483648
            _ = fi45
            _ = v46
        }
        _ = fi22
    }
    if (245 == 966) {
        delete(v5, v3)
    }
    var v50 float32 = 1.5
    switch v3 {
    case "foo bar":
        switch v3 {
        case "\n\t":
            var v51 uint8 = 2
            _ = v51
        case "hello":
            var v52 string = v3
            _ = v52
        case "a":
            if true {
                if (true && false) {
                    v21 = 42
                }
            }
        default:
            switch v3 {
            case "":
                v1 = append(v1, v4)
            case "hello":
                var v53 int16 = 0
                _ = v53
            default:
                v54 := GenS6{F0: genFn1(32767, 0), F1: 666, F2: 127}
                _ = v54
            }
        }
    }
    var v55 float32 = (v50 / 1.0)
    switch (v18 | (v18 % (((v18 >> (((^(v2))) & 63))) | 1))) {
    case -9223372036854775808:
        if false {
            fmt.Println(v3, v50)
            fmt.Println(v3)
        } else {
            switch ("foo bar" + v3) {
            case "hello":
                for fi56 := 0; fi56 < 4; fi56++ {
                    delete(v5, "foo bar")
                    if (2 >= 1) {
                        for fi57 := 0; fi57 < 14; fi57++ {
                            v58 := GenS7{F0: true}
                            v21 = v21
                            fmt.Println(v21, v18)
                            _ = fi57
                            _ = v58
                        }
                        delete(v5, "")
                    }
                    var v59 string = ("hello" + "\n\t")
                    _ = fi56
                    _ = v59
                }
            }
            fmt.Println(v4)
        }
    default:
        var v60 int64 = 0
        _ = v60
    }
    var v61 uint16 = 32767
    for fi62 := 0; fi62 < 18; fi62++ {
        if false {
            v2 -= 65535
            var v63 uint8 = 127
            var v64 uint = 255
            _ = v63
            _ = v64
        } else {
            switch v3 {
            case "foo bar":
                v65 := GenS6{F0: genFn1(42, -1), F1: -1, F2: 114}
                _ = v65
            }
        }
        if ((v4 & (^(v4))) == (^(v4))) {
            for fi66 := 0; fi66 < 6; fi66++ {
                fmt.Println(v61)
                v67 := GenS8{F0: v61, F1: v4}
                var v68 string = v3
                _ = fi66
                _ = v67
                _ = v68
            }
            delete(v5, v3)
            delete(v5, v3)
        }
        v18 = v18
        _ = fi62
    }
    v69 := GenS4{F0: 634, F1: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v55:", v55)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v55
    _ = v61
    _ = v69
}
