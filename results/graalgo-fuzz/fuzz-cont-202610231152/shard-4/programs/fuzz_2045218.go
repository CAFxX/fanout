package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint64) int16 {
    var v7 uint8 = 51
    if (-2147483648 > int32(p0)) {
        var v8 int8 = 48
        fmt.Println(v7, v8)
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return 686
}
func genFn1(p0 string, p1 int8, p2 uint16) uint16 {
    var v9 uint32 = 2147483647
    if ((^(p2)) > (p2 >> ((1) & 63))) {
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 int64 = 817
            v11 = (v11 >> ((1) & 63))
            _ = fi10
            _ = v11
        }
        switch (p0 + "hello") {
        case "foo bar":
            if (!true) {
                switch "foo bar" {
                case "foo bar":
                    v12 := GenS1{F0: (p1 + p1), F1: 2}
                    _ = v12
                case "hello":
                    var v13 int32 = 0
                    _ = v13
                }
                p1 -= (p1 & 100)
                for fi14 := 0; fi14 < 3; fi14++ {
                    if (0.0 != 3.14159) {
                        fmt.Println(p0, v9, v9)
                    } else {
                        for fi15 := 0; fi15 < 20; fi15++ {
                            v16 := GenS1{F0: 100, F1: 1}
                            fmt.Println(v9, fi15)
                            _ = fi15
                            _ = v16
                        }
                        fmt.Println(p0, p0, v9)
                    }
                    if false {
                        p0 = p0
                        switch p0 {
                        case "\n\t":
                            switch fi14 {
                            case 9223372036854775807:
                                var v17 uint32 = (v9 ^ v9)
                                _ = v17
                            case -9223372036854775808:
                                if false {
                                    var v18 int32 = 795
                                    v9 -= 999
                                    _ = v18
                                } else {
                                    for fi19 := 0; fi19 < 14; fi19++ {
                                        switch p0 {
                                        case "":
                                            fmt.Println(v9)
                                        case "a":
                                            fmt.Println(p2, p2)
                                        default:
                                            var v20 string = (p0 + "hello")
                                            _ = v20
                                        }
                                        var v21 int64 = 1
                                        if true {
                                            fmt.Println(v9)
                                            v22 := GenS1{F0: (p1 % ((-11) | 1)), F1: 1}
                                            var v23 float64 = 3.14159
                                            _ = v22
                                            _ = v23
                                        } else {
                                            if ((p0 + p0) <= p0) {
                                                v24 := GenS1{F0: p1, F1: uint64(p2)}
                                                v21 -= v21
                                                p1 -= p1
                                                p0 = ""
                                                var v25 int32 = -2147483648
                                                _ = v24
                                                _ = v25
                                            } else {
                                                v26 := []int{fi14}
                                                fmt.Println(p2)
                                                _ = v26
                                            }
                                            if ((-9223372036854775808 <= (fi19 ^ fi19)) || true) {
                                                v27 := []int{fi14, fi19}
                                                fmt.Println(p0, p1, v9)
                                                if false {
                                                    p0 = "hello"
                                                }
                                                var v28 int8 = p1
                                                _ = v27
                                                _ = v28
                                            } else {
                                                p1 = (p1 | (-(p1)))
                                                var v29 uint8 = 12
                                                v30 := []int{fi14, fi14, fi14}
                                                v21 += -9223372036854775808
                                                if len(v30) > 0 {
                                                    v30[0] = (^(fi14))
                                                }
                                                _ = v29
                                                _ = v30
                                            }
                                        }
                                        _ = fi19
                                        _ = v21
                                    }
                                }
                            }
                        case "hello":
                            var v31 float32 = ((-2.25 + 0.0) - (0.0 + (float32((3.14159 / -2.25)) + (0.0 + 0.1))))
                            _ = v31
                        case "a":
                            if (4294967295 <= 794) {
                                v32 := []int{700, fi14}
                                var v33 int = fi14
                                _ = v32
                                _ = v33
                            }
                        }
                        var v34 int16 = 1000
                        _ = v34
                    } else {
                        v35 := GenS1{F0: (-(p1)), F1: 524}
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 9; fi36++ {
                        v37 := []int{-9223372036854775808}
                        p2 = 255
                        if len(v37) >= 2 {
                            v37 = v37[1:len(v37)]
                        }
                        v38 := GenS2{F0: (65535 == p2)}
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                    _ = fi14
                }
            }
        case "":
            v39 := GenS1{F0: p1, F1: uint64(p2)}
            _ = v39
        case "hello":
            v40 := []int{32767, -146, -9223372036854775808, -2147483648}
            _ = v40
        }
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p2 % ((p2) | 1))
}

func main() {
    var v1 int = -101
    var v2 int32 = 883
    v3 := []int{(v1 * v1), (v1 - v1), (v1 << ((838) & 63))}
    var v4 float32 = float32(0.0)
    v5 := []string{"a"}
    var v6 string = ("a" + (("hello" + "a") + ("foo bar" + (("foo bar" + "a") + ("a" + ("hello" + ("hello" + (("hello" + "\n\t") + ((("hello" + (("hello" + "a") + "foo bar")) + (((((("a" + "\n\t") + "foo bar") + "\n\t") + "") + ("hello" + ("\n\t" + "\n\t"))) + ("\n\t" + ""))) + "")))))))))
    switch (("foo bar" + v6) + "") {
    case "\n\t":
        if false {
            fmt.Println(v1)
        }
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        var v42 bool = false
        v5 = append(v5, "\n\t")
        _ = fi41
        _ = v42
    }
    if false {
        switch ("a" + "foo bar") {
        case "foo bar":
            fmt.Println(v4, v6)
        case "":
            v43 := GenS3{F0: 2}
            _ = v43
        }
        fmt.Println(v6, v4, v6)
        for fi44 := 0; fi44 < 7; fi44++ {
            fmt.Println(fi44, v1)
            _ = fi44
        }
    }
    v4 -= (float32(float64((v4 * float32(float64(float32((1e10 / 1.0))))))) / (-2.25 - 1e10))
    var v45 float32 = v4
    for fi46 := 0; fi46 < 4; fi46++ {
        v47 := GenS2{F0: false}
        _ = fi46
        _ = v47
    }
    var v48 int16 = 0
    switch (^(v1)) {
    case -338:
        v48 = genFn0(0, 235)
    case -945:
        if (!false) {
            v49 := GenS2{F0: (!(270 < 255))}
            if false {
                v50 := GenS3{F0: 4294967295}
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
                var v51 float64 = float64(v4)
                _ = v50
                _ = v51
            }
            fmt.Println(v2, v6)
            _ = v49
        } else {
            var v52 uint16 = 65535
            _ = v52
        }
    }
    var v53 int32 = (v2 << ((4294967295) & 63))
    for fi54 := 0; fi54 < 8; fi54++ {
        if (193 < 143) {
            switch v6 {
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
            case "":
                v55 := GenS3{F0: 42}
                _ = v55
            default:
                switch 128 {
                case 32767:
                    var v56 int = fi54
                    _ = v56
                case 564:
                    for fi57 := 0; fi57 < 6; fi57++ {
                        switch v6 {
                        case "a":
                            v58 := GenS2{F0: (!false)}
                            _ = v58
                        }
                        v1 -= (fi57 >> ((42) & 63))
                        var v59 uint = 18446744073709551615
                        _ = fi57
                        _ = v59
                    }
                case -2147483648:
                    fmt.Println(v2, fi54, v45)
                }
            }
            if (!(0.0 <= float64(((v4 / (v45 - v45)) * v45)))) {
                fmt.Println(v6, v2)
                var v60 int8 = 1
                v53 = v2
                _ = v60
            }
            if len(v3) > 0 {
                v3[0] = v1
            }
        } else {
            var v61 int32 = v53
            _ = v61
        }
        fmt.Println(v2, v6, v6)
        _ = fi54
    }
    if len(v3) > 0 {
        v3[0] = v1
    }
    v3 = append(v3, (v1 ^ (^(v1))))
    if len(v5) > 0 {
        v5[0] = (v6 + v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v5
    _ = v53
    _ = v6
}
