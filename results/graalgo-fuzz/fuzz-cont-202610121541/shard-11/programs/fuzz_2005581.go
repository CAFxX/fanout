package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16 }
func genFn0() uint16 {
    v6 := []int{61, 32767, 128}
    switch 963 {
    case 32767:
        v7 := GenS1{F0: 255, F1: 0}
        _ = v7
    case -2147483648:
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 6; fi9++ {
                if ((!true) || true) {
                    var v10 uint = 989
                    if true {
                        v10 -= v10
                        var v11 int32 = 180
                        _ = v11
                    } else {
                        var v12 int64 = 9223372036854775807
                        var v13 string = ("" + "hello")
                        _ = v12
                        _ = v13
                    }
                    _ = v10
                }
                var v14 int32 = 0
                _ = fi9
                _ = v14
            }
            if false {
                var v15 float32 = 3.14159
                var v16 int32 = 2147483647
                _ = v15
                _ = v16
            }
            fmt.Println(fi8)
            _ = fi8
        }
    case 1:
        var v17 int = 0
        _ = v17
    }
    if true {
        if false {
            switch -1 {
            case 1:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case -77:
                _ = len(v6)
            }
            var v18 uint64 = 2
            _ = v18
        }
        var v19 bool = true
        for fi20 := 0; fi20 < 9; fi20++ {
            _ = len(v6)
            for fi21 := 0; fi21 < 9; fi21++ {
                v22 := GenS2{F0: 0}
                _ = fi21
                _ = v22
            }
            if (-1 >= 161) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if len(v6) > 0 {
                    v6[0] = (fi20 ^ 128)
                }
                if ((614 < 65535) || ((3.14159 + 0.0) < 1.5)) {
                    v19 = true
                } else {
                    switch "" {
                    case "a":
                        var v23 int64 = 660
                        _ = v23
                    default:
                        v19 = v19
                    }
                    if false {
                        var v24 int64 = -9223372036854775808
                        _ = len(v6)
                        _ = v24
                    }
                }
            }
            _ = fi20
        }
        _ = v19
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v6 = append(v6, 32768)
        var v25 string = "a"
        _ = v25
    }
    _ = v6
    return 51
}
func genFn1() int32 {
    var v26 uint64 = 0
    fmt.Println(v26)
    var v27 uint16 = 65535
    switch 127 {
    case 0:
        for fi28 := 0; fi28 < 9; fi28++ {
            var v29 uint16 = genFn0()
            _ = fi28
            _ = v29
        }
    case 127:
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 uint = 377
            _ = fi30
            _ = v31
        }
    }
    _ = v26
    _ = v27
    return -2147483648
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{65535, 65535, 42}
    v3 := []uint32{2147483647, 2, 2147483647}
    v4 := []int{32768, 1, -684}
    var v5 uint64 = 980
    v32 := GenS1{F0: 255, F1: 4294967295}
    fmt.Println(v5)
    fmt.Println(v5)
    var v33 uint32 = 858
    var v34 int32 = -1
    fmt.Println(v33, v33, v5)
    if (127 > 128) {
        switch 32768 {
        case -1:
            var v35 uint = 0
            _ = v35
        case 372:
            switch 1 {
            case 32768:
                for fi36 := 0; fi36 < 15; fi36++ {
                    v34 -= 1
                    switch fi36 {
                    case 4294967295:
                        v34 -= v34
                    default:
                        fmt.Println(v33, v34)
                    }
                    if false {
                        _ = len(v4)
                        v37 := GenS3{F0: 186}
                        if false {
                            var v38 int64 = -9223372036854775808
                            fmt.Println(v38)
                            v5 = v5
                            _ = v38
                        }
                        _ = v37
                    } else {
                        var v39 int64 = 42
                        _ = v39
                    }
                    _ = fi36
                }
            case 2147483647:
                if ((^(v33)) == v33) {
                    v4 = append(v4, -9223372036854775808)
                }
            case 9223372036854775807:
                var v40 uint32 = 42
                _ = v40
            }
        default:
            for fi41 := 0; fi41 < 11; fi41++ {
                if false {
                    v1[-477] = 891
                    var v42 uint64 = (v5 / (((v5 << ((255) & 63))) | 1))
                    _ = v42
                }
                _ = fi41
            }
        }
        var v43 uint32 = v33
        _ = v43
    } else {
        v33 = (v33 * (v33 - 2147483647))
    }
    if true {
        for fi44 := 0; fi44 < 8; fi44++ {
            v5 = (v5 >> ((253) & 63))
            v45 := GenS4{F0: genFn1(), F1: 42, F2: v33}
            _ = fi44
            _ = v45
        }
        for fi46 := 0; fi46 < 6; fi46++ {
            switch ((("hello" + ("a" + "foo bar")) + "hello") + ("" + "")) {
            case "foo bar":
                var v47 int32 = (v34 & (v34 - int32(v5)))
                _ = v47
            default:
                var v48 uint8 = 1
                _ = v48
            }
            v49 := GenS2{F0: 42}
            _ = fi46
            _ = v49
        }
        for fi50 := 0; fi50 < 11; fi50++ {
            var v51 int32 = (v34 * (v34 - v34))
            for fi52 := 0; fi52 < 13; fi52++ {
                if len(v4) > 0 {
                    v4[0] = -9223372036854775808
                }
                if len(v4) > 0 {
                    v4[0] = 77
                }
                _ = fi52
            }
            _ = fi50
            _ = v51
        }
    } else {
        v5 = (v5 << ((42) & 63))
        v53 := GenS5{F0: 2}
        _ = v53
    }
    delete(v1, 0)
    for fi54 := 0; fi54 < 12; fi54++ {
        var v55 int64 = 9223372036854775807
        var v56 uint8 = 1
        _ = fi54
        _ = v55
        _ = v56
    }
    v57 := GenS6{F0: 294}
    fmt.Println(v33)
    var v58 int8 = 100
    v58 = 127
    fmt.Println(v5, v5)
    if false {
        v34 -= (v34 ^ genFn1())
        v59 := GenS2{F0: v5}
        _ = v59
    } else {
        var v60 float32 = (-2.25 - 0.0)
        v5 += v5
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v57
    _ = v58
}
