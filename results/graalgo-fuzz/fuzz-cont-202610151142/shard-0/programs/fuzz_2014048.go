package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int }
type GenS2 struct { F0 int8; F1 int8; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 bool; F1 int16; F2 int32 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint16, p1 int64) uint32 {
    var v4 float64 = 3.14159
    switch -2147483648 {
    case 0:
        fmt.Println(v4, p1)
    }
    _ = p0
    _ = p1
    _ = v4
    return 4294967295
}
func genFn1(p0 string, p1 uint32) uint32 {
    v5 := GenS1{F0: p1, F1: 306, F2: 128}
    v6 := GenS1{F0: genFn0(65535, 1), F1: -782, F2: 127}
    switch "\n\t" {
    case "hello":
        v7 := GenS2{F0: -1, F1: -29, F2: -622}
        _ = v7
    case "a":
        var v8 bool = true
        _ = v8
    case "\n\t":
        var v9 uint32 = p1
        _ = v9
    default:
        var v10 int8 = 127
        _ = v10
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    var v1 int16 = -32768
    var v2 uint = 0
    var v3 uint64 = 18446744073709551615
    var v11 int32 = 190
    var v12 int16 = (v1 << ((v2) & 63))
    if ((32767 != 0) || true) {
        v2 -= v2
        v13 := GenS2{F0: 1, F1: -31, F2: 42}
        v14 := GenS2{F0: 127, F1: 1, F2: -149}
        _ = v13
        _ = v14
    } else {
        v15 := []int{-1}
        v1 = (v1 >> ((4294967295) & 63))
        v16 := GenS2{F0: -42, F1: -114, F2: -998}
        _ = v15
        _ = v16
    }
    var v17 string = "a"
    var v18 int64 = 9223372036854775807
    var v19 string = (v17 + v17)
    v20 := []int{}
    if false {
        v3 += v3
    } else {
        v20 = append(v20, 335)
        if true {
            for fi21 := 0; fi21 < 11; fi21++ {
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 float32 = (((((-2.25 + 0.0) * -2.25) * 1.5) / 1.5) + 1e10)
                    v24 := GenS3{F0: 2}
                    fmt.Println(v2)
                    v19 = v17
                    var v25 uint16 = 0
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                v20 = append(v20, (^(fi21)))
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 string = "\n\t"
                    _ = fi26
                    _ = v27
                }
                _ = fi21
            }
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            v12 = v12
        } else {
            for fi28 := 0; fi28 < 2; fi28++ {
                if true {
                    switch (^(fi28)) {
                    case 127:
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    case 32768:
                        v11 += (v11 % ((v11) | 1))
                    default:
                        for fi29 := 0; fi29 < 4; fi29++ {
                            v30 := GenS4{F0: -337, F1: 0}
                            _ = fi29
                            _ = v30
                        }
                    }
                    var v31 int64 = -9223372036854775808
                    switch v19 {
                    case "a":
                        switch int(v31) {
                        case 128:
                            if (false && true) {
                                fmt.Println(v31, v3)
                            }
                        case 763:
                            var v32 int = (fi28 - (fi28 & fi28))
                            _ = v32
                        }
                    case "hello":
                        if false {
                            v33 := GenS5{F0: (v2 * v2), F1: v3}
                            _ = v33
                        } else {
                            v17 = "a"
                        }
                    case "foo bar":
                        if false {
                            v34 := GenS6{F0: true, F1: v1, F2: v11}
                            var v35 int = fi28
                            v17 = v19
                            _ = v34
                            _ = v35
                        }
                    }
                    _ = v31
                }
                v12 = 1
                _ = fi28
            }
            var v36 float64 = 1.5
            if (v11 <= (^(v11))) {
                var v37 string = v17
                fmt.Println(v1, v3)
                _ = v37
            } else {
                fmt.Println(v2, v36)
                if ((^(v2)) >= v2) {
                    for fi38 := 0; fi38 < 19; fi38++ {
                        var v39 int = fi38
                        if true {
                            if true {
                                v11 = (v11 << (((v2 << ((1) & 63))) & 63))
                                v12 -= -533
                            } else {
                                fmt.Println(v1, v39, v11)
                                v11 = (v11 >> (((v2 & (^(v2)))) & 63))
                                switch (fi38 << (((v2 ^ 0)) & 63)) {
                                case 9223372036854775807:
                                    if (true && ((v2 << ((65) & 63)) > 0)) {
                                        if true {
                                            _ = len(v20)
                                        }
                                        for fi40 := 0; fi40 < 6; fi40++ {
                                            if false {
                                                v41 := GenS1{F0: 2147483647, F1: -274, F2: (fi40 + (-(fi38)))}
                                                v42 := GenS6{F0: (("" + "\n\t") > v19), F1: (v12 << (((v2 >> (((v2 & v2)) & 63))) & 63)), F2: v11}
                                                _ = v41
                                                _ = v42
                                            } else {
                                                switch "hello" {
                                                case "hello":
                                                    v3 -= v3
                                                case "\n\t":
                                                    switch "" {
                                                    case "\n\t":
                                                        var v43 uint8 = 127
                                                        _ = v43
                                                    case "hello":
                                                        v44 := GenS7{F0: v18, F1: v18}
                                                        _ = v44
                                                    default:
                                                        if ("" <= v19) {
                                                            var v45 uint64 = 223
                                                            if len(v20) >= 2 {
                                                                v20 = v20[1:len(v20)]
                                                            }
                                                            switch fi38 {
                                                            case 32767:
                                                                v46 := GenS3{F0: 1}
                                                                _ = v46
                                                            case 32768:
                                                                if len(v20) >= 2 {
                                                                    v20 = v20[1:len(v20)]
                                                                }
                                                            default:
                                                                v2 -= v2
                                                            }
                                                            fmt.Println(v19, v17)
                                                            v18 += (v18 ^ 42)
                                                            _ = v45
                                                        } else {
                                                            var v47 int16 = (-(v12))
                                                            if (!(!false)) {
                                                                v17 = "hello"
                                                            }
                                                            v18 -= 9223372036854775807
                                                            _ = v47
                                                        }
                                                    }
                                                }
                                            }
                                            fmt.Println(v1, v39, v17)
                                            _ = fi40
                                        }
                                        var v48 bool = (false && (((733 != (-(v1))) || (!(v19 < ((v19 + v17) + v17)))) && (!(v3 >= v3))))
                                        _ = v48
                                    }
                                case -337:
                                    v49 := GenS7{F0: 1, F1: -1}
                                    _ = v49
                                default:
                                    if len(v20) > 0 {
                                        v20[0] = 2147483647
                                    }
                                }
                            }
                        } else {
                            var v50 float32 = (1e10 - 0.1)
                            _ = v50
                        }
                        v20 = append(v20, (v39 | 127))
                        _ = fi38
                        _ = v39
                    }
                    if false {
                        v12 = (^(v12))
                    } else {
                        v12 += v12
                        v51 := GenS4{F0: (v18 - 1), F1: 42}
                        for fi52 := 0; fi52 < 20; fi52++ {
                            fmt.Println(v18, v36)
                            var v53 int8 = 9
                            if (!(fi52 <= fi52)) {
                                var v54 bool = ((v18 >> ((1) & 63)) == -1)
                                v1 -= (v12 << ((v2) & 63))
                                fmt.Println(v12)
                                _ = v54
                            } else {
                                for fi55 := 0; fi55 < 10; fi55++ {
                                    v1 = v12
                                    if (v2 == v2) {
                                        var v56 int = fi52
                                        v20 = append(v20, fi55)
                                        _ = v56
                                    } else {
                                        if (int32(v3) <= v11) {
                                            if true {
                                                v53 += (v53 * v53)
                                            } else {
                                                if (true && (2 > 763)) {
                                                    if false {
                                                        v57 := GenS8{F0: ((-2.25 < 1e10) || ((4294967295 < 2) && false))}
                                                        fmt.Println(v53, fi55)
                                                        var v58 bool = false
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                } else {
                                                    v1 += int16(fi55)
                                                }
                                            }
                                            if len(v20) > 0 {
                                                v20[0] = fi55
                                            }
                                            switch int(v12) {
                                            case -452:
                                                if (!(3.14159 != (1e10 + float32(3.14159)))) {
                                                    v20 = append(v20, fi52)
                                                    v19 = v19
                                                } else {
                                                    v2 = (v2 & (v2 & 255))
                                                }
                                            case 42:
                                                var v59 uint = v2
                                                _ = v59
                                            }
                                        }
                                        switch ("a" + v17) {
                                        case "\n\t":
                                            v36 = 1.5
                                        default:
                                            if (670 < (v3 >> (((v2 % ((65535) | 1))) & 63))) {
                                                var v60 string = (("" + "\n\t") + (v19 + ""))
                                                fmt.Println(v11, v53, v19)
                                                fmt.Println(v53, fi55)
                                                _ = v60
                                            }
                                        }
                                        v61 := GenS8{F0: (!(1e10 <= float64(0.1)))}
                                        _ = v61
                                    }
                                    _ = fi55
                                }
                                fmt.Println(v17, v53)
                            }
                            _ = fi52
                            _ = v53
                        }
                        _ = v51
                    }
                    v17 = "foo bar"
                } else {
                    v62 := GenS9{F0: v11, F1: 255}
                    _ = v62
                }
                var v63 float32 = float32(1e10)
                _ = v63
            }
            _ = v36
        }
    }
    fmt.Println(v12, v2, v18)
    v1 = (-(v1))
    for fi64 := 0; fi64 < 16; fi64++ {
        v20 = append(v20, (fi64 - -664))
        fmt.Println(v19)
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
