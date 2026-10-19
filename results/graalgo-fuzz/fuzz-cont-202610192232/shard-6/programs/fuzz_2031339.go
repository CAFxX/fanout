package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int64 }
type GenS2 struct { F0 int16; F1 int16; F2 uint64 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS7 struct { F0 int16; F1 int8 }
func genFn0() int {
    var v5 int64 = 9223372036854775807
    for fi6 := 0; fi6 < 4; fi6++ {
        v5 += v5
        v7 := []int{fi6, (fi6 + (fi6 << ((660) & 63))), (fi6 + 0), (^(fi6))}
        _ = fi6
        _ = v7
    }
    _ = v5
    return 42
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint8 = 221
    var v3 int = (v1 - (v1 >> ((65535) & 63)))
    var v4 string = (("hello" + "") + "\n\t")
    switch "a" {
    case "foo bar":
        v8 := []int{(v1 ^ v1), (^(v1)), v3}
        _ = v8
    case "a":
        v9 := GenS1{F0: -118, F1: 100, F2: -9223372036854775808}
        _ = v9
    case "hello":
        v4 = ("" + ("\n\t" + ""))
    }
    v10 := GenS1{F0: v1, F1: 100, F2: 1}
    v4 = v4
    switch int(v2) {
    case -193:
        var v11 int64 = -1
        _ = v11
    case 127:
        v1 = (v3 >> ((894) & 63))
    default:
        v4 = "foo bar"
    }
    v4 = (v4 + "\n\t")
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS1{F0: (v1 << ((2) & 63)), F1: int8(fi13), F2: -193}
            var v15 uint16 = 2
            switch (v3 + int(v2)) {
            case 42:
                if (726 >= 697) {
                    if (!true) {
                        v16 := GenS1{F0: (v3 >> ((4294967295) & 63)), F1: 100, F2: -1}
                        v15 = v15
                        _ = v16
                    } else {
                        var v17 uint64 = 42
                        switch v3 {
                        case 1:
                            var v18 bool = (false || false)
                            _ = v18
                        case 32767:
                            v4 = v4
                            v15 -= 1
                            if true {
                                if (((2147483647 == 2147483647) || false) && (0 != 2147483647)) {
                                    v19 := GenS2{F0: 1000, F1: 1, F2: 4294967295}
                                    _ = v19
                                }
                                switch v4 {
                                case "hello":
                                    switch (((("a" + v4) + v4) + "foo bar") + (v4 + "foo bar")) {
                                    case "a":
                                        var v20 uint64 = (v17 / (((^(v17))) | 1))
                                        _ = v20
                                    case "hello":
                                        v21 := GenS3{F0: 42, F1: 42}
                                        _ = v21
                                    case "\n\t":
                                        v17 -= v17
                                    default:
                                        var v22 int = 32767
                                        _ = v22
                                    }
                                case "":
                                    var v23 int32 = -1
                                    _ = v23
                                default:
                                    switch v4 {
                                    case "hello":
                                        var v24 bool = (2147483647 != -978)
                                        _ = v24
                                    case "a":
                                        if (true && (false || true)) {
                                            fmt.Println(v1, fi13)
                                            v25 := GenS4{F0: 818, F1: genFn0()}
                                            _ = v25
                                        } else {
                                            v15 += (^(v15))
                                            var v26 int32 = -2147483648
                                            fmt.Println(v26, v3, v1)
                                            fmt.Println(v17)
                                            _ = v26
                                        }
                                    default:
                                        v15 = v15
                                    }
                                }
                                v1 = v1
                            }
                            v27 := GenS2{F0: -477, F1: 32767, F2: 2}
                            v2 += v2
                            v3 -= (v3 & (fi12 >> ((792) & 63)))
                            _ = v27
                        default:
                            var v28 uint = 1
                            _ = v28
                        }
                        var v29 string = "a"
                        _ = v17
                        _ = v29
                    }
                    var v30 uint32 = 0
                    v31 := GenS3{F0: 32767, F1: -1}
                    _ = v30
                    _ = v31
                } else {
                    v3 -= (fi13 ^ genFn0())
                    fmt.Println(v15, v3)
                }
            case 32768:
                var v32 uint32 = 2147483648
                _ = v32
            case 0:
                fmt.Println(v3, v4, v15)
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        v33 := GenS5{F0: (v2 - 42)}
        var v34 uint32 = 594
        _ = fi12
        _ = v33
        _ = v34
    }
    if ((0.0 + ((-2.25 * ((((0.0 * ((0.1 / 1.5) * ((float32(((3.14159 - (1.5 / (1.5 + 3.14159))) - 1.5)) + 1.5) / (-2.25 * -2.25)))) / 1.0) + 0.1) - 1e10)) + (0.0 * 0.0))) == 3.14159) {
        switch v4 {
        case "a":
            if (v4 != v4) {
                for fi35 := 0; fi35 < 2; fi35++ {
                    var v36 string = (v4 + v4)
                    if false {
                        v4 = ((("\n\t" + ("foo bar" + ("a" + v4))) + v36) + v36)
                        if (!true) {
                            v1 = genFn0()
                            v2 = v2
                        } else {
                            v37 := GenS2{F0: 0, F1: 32767, F2: 2}
                            _ = v37
                        }
                    }
                    _ = fi35
                    _ = v36
                }
                fmt.Println(v1)
                v38 := []int{-9223372036854775808}
                _ = v38
            }
        case "\n\t":
            v4 = v4
        case "":
            if (v3 > (v3 >> ((18446744073709551615) & 63))) {
                var v39 float64 = ((-2.25 * 0.1) / 1.5)
                _ = v39
            } else {
                v40 := []int{int(v2)}
                v1 = 42
                for fi41 := 0; fi41 < 14; fi41++ {
                    fmt.Println(v2)
                    v4 = ((v4 + v4) + ((v4 + v4) + v4))
                    _ = fi41
                }
                _ = v40
            }
        default:
            if (882 <= 18446744073709551615) {
                var v42 int32 = 42
                _ = v42
            } else {
                v1 = v3
            }
        }
    } else {
        for fi43 := 0; fi43 < 19; fi43++ {
            if false {
                if (true && (-123 >= 32767)) {
                    if (!(-583 != 0)) {
                        switch (v4 + "hello") {
                        case "hello":
                            v4 = ""
                        case "\n\t":
                            v44 := GenS4{F0: 2147483647, F1: -550}
                            _ = v44
                        case "a":
                            v45 := GenS3{F0: -32768, F1: 0}
                            _ = v45
                        default:
                            fmt.Println(v1, fi43)
                        }
                        for fi46 := 0; fi46 < 1; fi46++ {
                            var v47 int16 = -340
                            var v48 float64 = ((((0.0 - float64(-2.25)) + (-2.25 / 3.14159)) * -2.25) / 3.14159)
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                        v49 := GenS6{F0: -128, F1: 128, F2: 18446744073709551615}
                        _ = v49
                    }
                } else {
                    v50 := []int{v3, 4294967295, v3, (fi43 - genFn0())}
                    if ((-2.25 / 1e10) != (0.1 - (-2.25 + ((-2.25 / 1.0) / 1.0)))) {
                        var v51 uint = 4294967295
                        var v52 string = "foo bar"
                        var v53 string = "\n\t"
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    v4 = "\n\t"
                    _ = v50
                }
                var v54 int32 = -1
                v55 := GenS7{F0: -1, F1: 100}
                _ = v54
                _ = v55
            }
            _ = fi43
        }
        v3 += v3
    }
    var v56 uint32 = 0
    if (-1 <= -1) {
        v3 -= -702
    }
    v4 = (v4 + "")
    var v57 int32 = 202
    v58 := GenS3{F0: -32768, F1: 42}
    var v59 int8 = 127
    v57 = v57
    switch (v4 + v4) {
    case "\n\t":
        var v60 int64 = -9223372036854775808
        _ = v60
    default:
        var v61 float64 = 0.1
        _ = v61
    }
    for fi62 := 0; fi62 < 13; fi62++ {
        v3 += fi62
        _ = fi62
    }
    fmt.Println(v57)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
