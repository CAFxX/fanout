package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int64 }
type GenS2 struct { F0 int16; F1 int; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64; F1 bool; F2 int32 }
func genFn0(p0 uint8) string {
    var v6 uint8 = (p0 & (p0 << ((42) & 63)))
    v7 := []int{32767, 2147483647, 4294967295, -650}
    _ = p0
    _ = v6
    _ = v7
    return ("\n\t" + "a")
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 int16 = 727
    v4 := make(map[string]uint32)
    var v5 uint8 = 255
    v5 += v5
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := []int{127, fi8, (fi8 + fi8), (fi8 / (((fi8 / (((fi8 - fi8)) | 1))) | 1))}
        _ = fi8
        _ = v9
    }
    v5 = (v5 >> ((2) & 63))
    v10 := GenS1{F0: 42, F1: 0, F2: -9223372036854775808}
    v11 := GenS2{F0: v3, F1: 9223372036854775807, F2: 287}
    v12 := GenS1{F0: 42, F1: 127, F2: 0}
    v5 = v5
    if (float32((-2.25 - ((1.5 - float64(3.14159)) - (float64(3.14159) - 1e10)))) == 1e10) {
        v13 := GenS3{F0: 42, F1: v3}
        _ = v13
    } else {
        if ((-1 > -354) || true) {
            for fi14 := 0; fi14 < 13; fi14++ {
                if (18446744073709551615 <= 42) {
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 string = ("foo bar" + (("a" + "hello") + (("hello" + "") + (("hello" + "a") + (("" + ("a" + "\n\t")) + "foo bar")))))
                        var v17 uint64 = 712
                        switch (v16 + "foo bar") {
                        case "":
                            v18 := GenS1{F0: 2147483648, F1: -1, F2: -804}
                            _ = v18
                        case "\n\t":
                            switch fi14 {
                            case 127:
                                if (!(!(v17 >= v17))) {
                                    var v19 bool = false
                                    if true {
                                        v3 = 245
                                        v3 -= 0
                                    } else {
                                        if true {
                                            var v20 int16 = (^(v3))
                                            if true {
                                                if (v5 != (v5 * (v5 + v5))) {
                                                    _ = len(v4)
                                                    if true {
                                                        if v19 {
                                                            var v21 int64 = 0
                                                            var v22 uint64 = (v17 | v17)
                                                            switch (v16 + ("a" + v16)) {
                                                            case "":
                                                                if v19 {
                                                                    if (!(!true)) {
                                                                        var v23 uint32 = 2147483648
                                                                        _ = v23
                                                                    }
                                                                    var v24 float32 = (1e10 - 1e10)
                                                                    v5 = 1
                                                                    _ = v24
                                                                } else {
                                                                    v17 -= 42
                                                                }
                                                            }
                                                            _ = v21
                                                            _ = v22
                                                        }
                                                    } else {
                                                        v16 = ((v16 + v16) + "")
                                                    }
                                                    if true {
                                                        if (!(-2147483648 <= 42)) {
                                                            var v25 uint = 0
                                                            v26 := GenS1{F0: 686, F1: 100, F2: 9223372036854775807}
                                                            v17 -= (v17 % ((v17) | 1))
                                                            _ = v25
                                                            _ = v26
                                                        } else {
                                                            var v27 string = "foo bar"
                                                            _ = v27
                                                        }
                                                        var v28 uint = 647
                                                        _ = v28
                                                    } else {
                                                        v29 := []int{(fi14 - fi14), -2147483648}
                                                        v5 += (v5 & v5)
                                                        v30 := GenS1{F0: 528, F1: 1, F2: -1}
                                                        switch "hello" {
                                                        case "hello":
                                                            v29 = append(v29, 32767)
                                                        case "foo bar":
                                                            if (1.5 <= -2.25) {
                                                                v19 = (!(true || (1 == 0)))
                                                            }
                                                        }
                                                        v16 = "\n\t"
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                } else {
                                                    v17 += v17
                                                }
                                            } else {
                                                v19 = v19
                                            }
                                            _ = v20
                                        } else {
                                            v31 := GenS3{F0: 4294967295, F1: v3}
                                            _ = v31
                                        }
                                    }
                                    if false {
                                        var v32 int8 = 0
                                        _ = v32
                                    } else {
                                        var v33 int8 = 85
                                        v34 := GenS4{F0: 0}
                                        _ = v33
                                        _ = v34
                                    }
                                    var v35 uint8 = (v5 / ((255) | 1))
                                    _ = v19
                                    _ = v35
                                }
                            default:
                                if true {
                                    switch 0 {
                                    case 709:
                                        v36 := GenS5{F0: 9223372036854775807}
                                        var v37 string = (v16 + v16)
                                        v38 := GenS6{F0: v5, F1: -75}
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    default:
                                        switch v16 {
                                        case "":
                                            var v39 int64 = 170
                                            _ = v39
                                        default:
                                            var v40 int8 = int8(fi15)
                                            _ = v40
                                        }
                                    }
                                    v5 = v5
                                    var v41 int8 = 42
                                    _ = v41
                                } else {
                                    v42 := []int{}
                                    delete(v2, fi14)
                                    var v43 int16 = (-(v3))
                                    _ = v42
                                    _ = v43
                                }
                            }
                        default:
                            var v44 int64 = -9223372036854775808
                            _ = v44
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    v45 := GenS7{F0: 712}
                    for fi46 := 0; fi46 < 17; fi46++ {
                        if (884 < 0) {
                            if (v5 >= 25) {
                                v5 += 2
                                v5 = (v5 + 163)
                            } else {
                                v5 -= (v5 << ((4294967295) & 63))
                                v47 := GenS7{F0: 42}
                                var v48 bool = true
                                _ = v47
                                _ = v48
                            }
                        }
                        if (2147483647 >= 2147483648) {
                            switch "foo bar" {
                            case "foo bar":
                                if false {
                                    fmt.Println(v3)
                                    if true {
                                        v3 = (v3 + (v3 >> ((2) & 63)))
                                        fmt.Println(fi46)
                                        if true {
                                            fmt.Println(v3, fi14, v3)
                                        }
                                    }
                                } else {
                                    var v49 uint = 4294967295
                                    _ = v49
                                }
                            }
                            switch "\n\t" {
                            case "":
                                v3 -= 1000
                            }
                        }
                        _ = fi46
                    }
                    _ = v45
                } else {
                    if ("hello" >= "\n\t") {
                        v50 := GenS6{F0: v5, F1: 0}
                        fmt.Println(v5, v5, fi14)
                        _ = v50
                    } else {
                        if (42 >= -97) {
                            v51 := v4["\n\t"]
                            var v52 string = "hello"
                            _ = v51
                            _ = v52
                        } else {
                            var v53 bool = (true || false)
                            v5 = (v5 / ((v5) | 1))
                            v53 = (515 >= 65535)
                            _ = v53
                        }
                        for fi54 := 0; fi54 < 2; fi54++ {
                            var v55 int16 = -401
                            var v56 uint32 = 65535
                            var v57 int32 = -1
                            _ = fi54
                            _ = v55
                            _ = v56
                            _ = v57
                        }
                        fmt.Println(fi14, fi14, v5)
                    }
                    v58 := GenS3{F0: 377, F1: v3}
                    _ = v58
                }
                v5 = (v5 << ((4294967295) & 63))
                _ = fi14
            }
        }
        v5 -= v5
        var v59 string = "\n\t"
        _ = v59
    }
    switch "" {
    case "a":
        _ = len(v2)
    case "":
        v60 := GenS8{F0: 9223372036854775807, F1: (3.14159 != ((1e10 / 1.5) + -2.25)), F2: 42}
        _ = v60
    case "hello":
        v61 := []int{42}
        _ = v61
    default:
        var v62 int8 = 100
        _ = v62
    }
    if false {
        v3 = v3
    } else {
        v5 = v5
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
