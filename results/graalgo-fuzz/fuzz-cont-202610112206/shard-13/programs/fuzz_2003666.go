package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 int16; F2 uint64 }
type GenS6 struct { F0 int16; F1 uint32 }
func genFn0() bool {
    if (true && true) {
        v7 := GenS1{F0: 205}
        switch 0 {
        case 127:
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 int16 = 1000
                _ = fi8
                _ = v9
            }
        }
        var v10 uint64 = 42
        _ = v10
        _ = v7
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int8 = -46
        v13 := GenS1{F0: (fi11 + (fi11 + fi11))}
        var v14 uint64 = 4294967295
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 float32 = 0.0
    _ = v15
    return (!true)
}
func genFn1() int64 {
    var v16 uint = 526
    v16 -= v16
    if (127 <= uint8(v16)) {
        switch (("hello" + "a") + "") {
        case "\n\t":
            var v17 float32 = 1e10
            _ = v17
        case "a":
            v18 := GenS2{F0: 201, F1: 131}
            _ = v18
        case "hello":
            fmt.Println(v16)
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS1{F0: fi19}
        v21 := GenS3{F0: 1, F1: 2, F2: 42}
        _ = fi19
        _ = v20
        _ = v21
    }
    _ = v16
    return 42
}

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 158
    v3 := make(map[int]uint32)
    var v4 uint32 = 87
    v5 := make(map[uint32]int)
    v6 := make(map[int]bool)
    fmt.Println(v2)
    fmt.Println(v1, v1, v1)
    var v22 uint = 4294967295
    v23 := GenS4{F0: -9223372036854775808}
    v22 = (v22 * v22)
    fmt.Println(v22)
    fmt.Println(v2, v22, v1)
    v24 := []int{127, 42}
    if true {
        v25 := v5[(v4 % ((uint32(v2)) | 1))]
        switch (v25 | v25) {
        case 4294967295:
            var v26 uint32 = 410
            _ = v26
        }
        switch v25 {
        case 32768:
            v25 += v25
        default:
            v22 = (v22 / (((v22 & v22)) | 1))
        }
        _ = v25
    }
    v22 = 458
    for fi27 := 0; fi27 < 16; fi27++ {
        if ((fi27 >= (fi27 & (-(fi27)))) && false) {
            v22 = v22
        } else {
            if false {
                v28 := v5[v4]
                v29 := GenS3{F0: (v2 | (^(v2))), F1: (v4 / (((^(v4))) | 1)), F2: 39}
                if ((114 != int32(v1)) || false) {
                    v4 += (v4 % (((^(v4))) | 1))
                    switch (("foo bar" + ("" + "foo bar")) + "hello") {
                    case "a":
                        _ = len(v6)
                    case "":
                        var v30 bool = false
                        _ = v30
                    default:
                        switch "hello" {
                        case "hello":
                            v31 := GenS3{F0: uint8(v1), F1: 1, F2: 45}
                            _ = v31
                        case "\n\t":
                            v4 += (v4 * (v4 << (((^(v22))) & 63)))
                        case "foo bar":
                            v32 := GenS3{F0: (^(v2)), F1: v4, F2: v2}
                            _ = v32
                        }
                    }
                } else {
                    switch 32768 {
                    case 9223372036854775807:
                        v1 = (v1 - 4294967295)
                    case -627:
                        var v33 uint = v22
                        _ = v33
                    case 86:
                        switch v28 {
                        case 32768:
                            var v34 int64 = -462
                            _ = v34
                        case 1:
                            var v35 string = ""
                            _ = v35
                        default:
                            if (!((4294967295 <= v1) && true)) {
                                v1 = (v1 + 1)
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                                for fi36 := 0; fi36 < 14; fi36++ {
                                    if ((("hello" == "foo bar") && false) || false) {
                                        switch v28 {
                                        case 274:
                                            v2 = v2
                                        case 42:
                                            var v37 int = v28
                                            _ = v37
                                        }
                                        v2 -= (v2 + v2)
                                    }
                                    v22 = (v22 % ((164) | 1))
                                    var v38 uint64 = 0
                                    var v39 int8 = 100
                                    _ = fi36
                                    _ = v38
                                    _ = v39
                                }
                            } else {
                                v1 = v1
                            }
                        }
                    }
                }
                _ = v28
                _ = v29
            }
            for fi40 := 0; fi40 < 9; fi40++ {
                switch 127 {
                case -691:
                    v24 = append(v24, 0)
                }
                fmt.Println(v1, v1)
                switch "a" {
                case "foo bar":
                    var v41 uint8 = 128
                    v42 := GenS5{F0: (v22 / ((772) | 1)), F1: 0, F2: (v1 + v1)}
                    fmt.Println(v4, v4)
                    var v43 bool = ((904 > (^(v4))) && false)
                    _ = v41
                    _ = v42
                    _ = v43
                case "":
                    v22 = (v22 % ((v22) | 1))
                }
                _ = fi40
            }
            if (!false) {
                _ = len(v6)
                v44 := GenS5{F0: v22, F1: 1000, F2: (v1 | 18446744073709551615)}
                _ = v44
            } else {
                for fi45 := 0; fi45 < 16; fi45++ {
                    v4 -= v4
                    _ = fi45
                }
                if ((v4 << ((v22) & 63)) == v4) {
                    switch -909 {
                    case 32768:
                        switch "foo bar" {
                        case "a":
                            if (v2 != v2) {
                                var v46 float32 = (float32((0.0 + float64(1e10))) * ((0.1 * ((1e10 - 3.14159) - (1.5 - 0.1))) / (float32(float64(float32(3.14159))) * 1.5)))
                                v46 = (v46 * 1.5)
                                fmt.Println(fi27, v22, v22)
                                _ = v46
                            }
                        case "":
                            if false {
                                var v47 uint32 = v4
                                _ = v47
                            }
                        default:
                            var v48 int = (fi27 / (((fi27 | fi27)) | 1))
                            _ = v48
                        }
                    default:
                        var v49 uint8 = 42
                        _ = v49
                    }
                    v4 += (v4 & 4294967295)
                    for fi50 := 0; fi50 < 19; fi50++ {
                        var v51 uint32 = (v4 / ((v4) | 1))
                        _ = len(v24)
                        _ = fi50
                        _ = v51
                    }
                } else {
                    var v52 int32 = -1
                    v1 -= (v1 % ((v1) | 1))
                    for fi53 := 0; fi53 < 6; fi53++ {
                        var v54 int64 = -812
                        var v55 uint32 = v4
                        v56 := GenS6{F0: -471, F1: v55}
                        v57 := GenS2{F0: (v54 & int64(v2)), F1: (v52 & 1)}
                        _ = fi53
                        _ = v54
                        _ = v55
                        _ = v56
                        _ = v57
                    }
                    _ = v52
                }
            }
        }
        _ = fi27
    }
    fmt.Println(v4)
    var v58 string = ("" + "foo bar")
    v3[32768] = v4
    if (-950 <= 32767) {
        v2 += (^(v2))
        fmt.Println(v4, v22, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
}
