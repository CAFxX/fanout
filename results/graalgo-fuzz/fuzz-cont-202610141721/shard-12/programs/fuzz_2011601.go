package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint }
type GenS2 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint16 }
func genFn0() uint64 {
    v6 := []int{}
    if true {
        var v7 int = 128
        fmt.Println(v7)
        _ = v7
    } else {
        if (0 <= -82) {
            var v8 uint = 4294967295
            var v9 uint16 = 1
            var v10 int8 = 0
            _ = v10
            _ = v8
            _ = v9
        } else {
            var v11 bool = (22 <= 42)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v11
        }
        switch 1 {
        case 32767:
            v6 = append(v6, 32768)
        case 127:
            switch -1 {
            case -365:
                _ = len(v6)
            case 127:
                var v12 string = ("" + "\n\t")
                _ = v12
            }
        case 42:
            var v13 uint = 42
            _ = v13
        default:
            switch 2147483647 {
            case 32768:
                var v14 uint = 65535
                _ = v14
            }
        }
        var v15 string = "foo bar"
        _ = v15
    }
    _ = v6
    return 4294967295
}

func main() {
    v1 := []int{808, -426, -9223372036854775808, -2147483648}
    v2 := []string{"", "a", ((("\n\t" + ("hello" + "a")) + "") + "a"), "hello"}
    var v3 uint16 = 1
    v4 := make(map[uint32]int)
    var v5 int8 = -128
    fmt.Println(v3)
    v3 += (v3 << ((4294967295) & 63))
    var v16 float32 = 0.0
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 uint32 = 2
        _ = fi17
        _ = v18
    }
    v3 -= v3
    v16 -= float32(0.0)
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 bool = (2147483648 < 1)
        delete(v4, 4294967295)
        if v20 {
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 uint = 65535
                var v23 int8 = v5
                v4[737] = (^(fi21))
                if v20 {
                    switch "hello" {
                    case "":
                        fmt.Println(v16, fi19, v23)
                    case "hello":
                        v24 := GenS1{F0: 567, F1: -367, F2: v22}
                        _ = v24
                    case "a":
                        if false {
                            fmt.Println(v16, v20, v3)
                        } else {
                            v25 := v4[898]
                            v26 := GenS2{F0: 1000, F1: 0, F2: uint64(v22)}
                            _ = v25
                            _ = v26
                        }
                    default:
                        var v27 bool = v20
                        _ = v27
                    }
                    var v28 int64 = -321
                    v28 = 360
                    var v29 string = "foo bar"
                    _ = v28
                    _ = v29
                } else {
                    switch ((("" + "\n\t") + "hello") + ("" + "hello")) {
                    case "a":
                        var v30 uint16 = (^(v3))
                        _ = v30
                    }
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            v5 = 100
        } else {
            var v31 uint = 42
            if (v20 && ((v5 * (-(v5))) <= v5)) {
                for fi32 := 0; fi32 < 1; fi32++ {
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        v4[65535] = (fi19 % ((32767) | 1))
                    case "a":
                        v2 = append(v2, ("hello" + "hello"))
                    }
                    if v20 {
                        if true {
                            v33 := GenS1{F0: 1, F1: 32767, F2: (v31 >> ((v31) & 63))}
                            var v34 uint64 = 920
                            v31 = 255
                            if (!(!v20)) {
                                v35 := GenS2{F0: -1, F1: (v34 ^ v34), F2: genFn0()}
                                var v36 string = ""
                                v20 = (2 > 169)
                                _ = v35
                                _ = v36
                            } else {
                                fmt.Println(v3, v5)
                                v37 := GenS3{F0: v20}
                                _ = v37
                            }
                            if v20 {
                                v38 := GenS3{F0: v20}
                                if true {
                                    v5 = 42
                                }
                                var v39 uint = (v31 ^ v31)
                                if (((("a" + "foo bar") + (("hello" + "") + "foo bar")) + (("foo bar" + "foo bar") + "hello")) >= "\n\t") {
                                    var v40 int32 = 2147483647
                                    v1 = append(v1, (fi19 / (((fi19 + -129)) | 1)))
                                    v5 -= v5
                                    v4[1] = (fi32 | fi32)
                                    var v41 string = "\n\t"
                                    var v42 uint16 = (v3 & v3)
                                    var v43 uint = (v31 + uint(v42))
                                    v44 := GenS4{F0: fi19, F1: (v40 / ((v40) | 1))}
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                }
                                _ = v38
                                _ = v39
                            }
                            v3 = (v3 * v3)
                            switch "" {
                            case "hello":
                                var v45 uint16 = v3
                                _ = v45
                            case "\n\t":
                                switch "a" {
                                case "":
                                    var v46 uint32 = 0
                                    _ = v46
                                case "hello":
                                    switch fi19 {
                                    case 4294967295:
                                        v47 := GenS2{F0: 1, F1: v34, F2: 4294967295}
                                        _ = v47
                                    }
                                }
                            }
                            v34 = (^(v34))
                            _ = v33
                            _ = v34
                        }
                        if v20 {
                            var v48 uint16 = v3
                            _ = v48
                        } else {
                            switch fi19 {
                            case 0:
                                v49 := GenS2{F0: -1, F1: 42, F2: 678}
                                _ = v49
                            case 2147483647:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case 468:
                                var v50 uint = (v31 >> ((v31) & 63))
                                _ = v50
                            default:
                                v51 := GenS5{F0: 1}
                                _ = v51
                            }
                            var v52 bool = true
                            var v53 uint = (v31 / ((v31) | 1))
                            _ = v52
                            _ = v53
                        }
                    } else {
                        v54 := GenS1{F0: 746, F1: -32768, F2: 18446744073709551615}
                        v31 = 56
                        switch 9223372036854775807 {
                        case 128:
                            v16 += -2.25
                        case 536:
                            v5 -= 42
                        }
                        v31 = v31
                        _ = v54
                    }
                    _ = fi32
                }
            }
            _ = v31
        }
        _ = fi19
        _ = v20
    }
    v16 = (v16 * float32(1e10))
    v5 += v5
    switch -655 {
    case 2147483647:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v55 float32 = v16
    v3 = (v3 | (v3 << ((225) & 63)))
    v56 := v4[878]
    v56 -= 32768
    var v57 float64 = 1e10
    fmt.Println(v16)
    v16 += v55
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
}
