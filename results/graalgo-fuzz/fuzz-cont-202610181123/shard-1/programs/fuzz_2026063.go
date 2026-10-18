package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int; F1 uint8 }
func genFn0() uint {
    var v4 int = 32768
    var v5 uint16 = 0
    v6 := GenS1{F0: (!true)}
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint8 = 128
    var v2 string = ("a" + "\n\t")
    var v3 bool = false
    v2 = v2
    var v7 bool = v3
    v3 = v3
    var v8 int16 = -789
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: v7}
        var v11 int8 = 82
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int16 = (^(v8))
    switch v2 {
    case "hello":
        v13 := GenS1{F0: (true || (18446744073709551615 > 4294967295))}
        _ = v13
    case "a":
        var v14 int = -1
        _ = v14
    case "":
        v2 = (v2 + v2)
    default:
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 int64 = 0
                fmt.Println(fi15, v3)
                var v18 uint64 = 42
                _ = fi16
                _ = v17
                _ = v18
            }
            if false {
                if (v7 || ("" >= "a")) {
                    v19 := []int{1, fi15, (fi15 & 32768)}
                    v20 := GenS2{F0: 18446744073709551615, F1: (v1 % ((v1) | 1)), F2: 0}
                    var v21 string = v2
                    _ = v19
                    _ = v20
                    _ = v21
                }
            } else {
                if v7 {
                    if (!(genFn0() < 317)) {
                        var v22 uint16 = 65535
                        v7 = true
                        v23 := []int{fi15}
                        _ = v22
                        _ = v23
                    }
                    var v24 float32 = (3.14159 + (1e10 / 3.14159))
                    _ = v24
                } else {
                    var v25 float64 = (1e10 + 3.14159)
                    if (127 != (-(fi15))) {
                        var v26 int32 = 935
                        var v27 int32 = -1
                        for fi28 := 0; fi28 < 15; fi28++ {
                            switch (fi28 - (fi15 << ((509) & 63))) {
                            case 0:
                                if v3 {
                                    v29 := GenS3{F0: 0, F1: 65535}
                                    v30 := []int{fi28}
                                    _ = v29
                                    _ = v30
                                }
                            default:
                                if (!false) {
                                    switch fi15 {
                                    case 4294967295:
                                        switch fi28 {
                                        case -265:
                                            if ((^(fi15)) > (fi15 + (fi28 + 0))) {
                                                switch v2 {
                                                case "a":
                                                    if (true && (!true)) {
                                                        if (0 >= 65535) {
                                                            v31 := []int{(fi28 - (fi28 / ((-1) | 1))), 2147483647, 2147483647, fi15}
                                                            v32 := GenS4{F0: int(v27), F1: (v1 * 136)}
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                        switch ("a" + "hello") {
                                                        case "hello":
                                                            v26 = v26
                                                        case "a":
                                                            var v33 float64 = v25
                                                            _ = v33
                                                        }
                                                        if true {
                                                            v7 = false
                                                            var v34 int64 = -9223372036854775808
                                                            var v35 int = -2147483648
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                    }
                                                default:
                                                    v36 := []int{127}
                                                    _ = v36
                                                }
                                            } else {
                                                if false {
                                                    v37 := []int{(fi28 ^ fi15), fi28, -1, 280}
                                                    var v38 int = fi15
                                                    var v39 uint16 = 2
                                                    _ = v37
                                                    _ = v38
                                                    _ = v39
                                                }
                                            }
                                        }
                                    case 2147483647:
                                        v40 := []int{4294967295}
                                        _ = v40
                                    case 39:
                                        var v41 int8 = -128
                                        _ = v41
                                    }
                                    v2 = ("" + (v2 + v2))
                                }
                            }
                            _ = fi28
                        }
                        _ = v26
                        _ = v27
                    } else {
                        var v42 float64 = v25
                        if (v3 && (1.5 == v42)) {
                            v43 := []int{(fi15 - 2147483647), (fi15 | (^(fi15))), 1}
                            switch (fi15 * fi15) {
                            case 127:
                                for fi44 := 0; fi44 < 9; fi44++ {
                                    var v45 float32 = (((3.14159 * 0.0) / (1.5 * 1e10)) + 1.5)
                                    if false {
                                        v7 = true
                                        var v46 int8 = -59
                                        var v47 int8 = v46
                                        _ = v46
                                        _ = v47
                                    }
                                    _ = fi44
                                    _ = v45
                                }
                            }
                            _ = v43
                        } else {
                            v7 = (((0.0 / v25) >= 0.0) || ((!v3) || (866 > 680)))
                        }
                        _ = v42
                    }
                    var v48 uint64 = 733
                    _ = v25
                    _ = v48
                }
            }
            _ = fi15
        }
    }
    switch 0 {
    case 0:
        v12 = v12
    }
    fmt.Println(v12, v8, v3)
    v3 = (true && true)
    v12 += (v12 >> ((85) & 63))
    v12 -= (^(v12))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
