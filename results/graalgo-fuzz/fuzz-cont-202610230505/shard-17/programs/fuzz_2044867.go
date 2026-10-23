package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int }
type GenS2 struct { F0 int; F1 int64; F2 bool }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint32; F1 uint; F2 int }
type GenS5 struct { F0 int; F1 int; F2 string }
type GenS6 struct { F0 int; F1 int8; F2 uint32 }
type GenS7 struct { F0 int32 }
func genFn0() uint8 {
    var v5 float64 = -2.25
    v5 = v5
    var v6 int64 = -9223372036854775808
    _ = v5
    _ = v6
    return 222
}

func main() {
    var v1 int64 = -666
    var v2 bool = (567 > 917)
    var v3 float32 = (0.0 / (1.5 - (-2.25 - 0.1)))
    v4 := make(map[int]string)
    var v7 float64 = 1e10
    switch ("hello" + "\n\t") {
    case "":
        v8 := GenS1{F0: (899 == 638), F1: 286, F2: 127}
        _ = v8
    case "hello":
        if v2 {
            switch "hello" {
            case "\n\t":
                fmt.Println(v7, v1)
            }
            v9 := GenS2{F0: 395, F1: v1, F2: (!((!(!v2)) && ("\n\t" != ("hello" + "\n\t"))))}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 19; fi10++ {
                v3 += v3
                switch "hello" {
                case "\n\t":
                    v2 = ((("foo bar" + "a") + ("" + "hello")) == "hello")
                }
                for fi11 := 0; fi11 < 11; fi11++ {
                    switch (-(fi10)) {
                    case 9223372036854775807:
                        v7 = v7
                    case 127:
                        v7 -= -2.25
                    case 128:
                        switch "foo bar" {
                        case "a":
                            v12 := v4[fi11]
                            v4[706] = ((v12 + (v12 + (v12 + v12))) + ((("\n\t" + v12) + v12) + "\n\t"))
                            _ = v12
                        }
                    default:
                        delete(v4, fi10)
                    }
                    _ = fi11
                }
                _ = fi10
            }
        }
    case "foo bar":
        _ = len(v4)
    }
    v13 := GenS1{F0: false, F1: 247, F2: 818}
    switch "a" {
    case "hello":
        v14 := []int{32767, 4294967295}
        _ = v14
    case "foo bar":
        if v2 {
            var v15 int = 42
            delete(v4, 9223372036854775807)
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := []int{-940, (fi16 ^ fi16), 42, fi16}
        v17 = append(v17, (fi16 & fi16))
        if v2 {
            v1 = v1
            v18 := GenS1{F0: false, F1: 18446744073709551615, F2: 127}
            _ = v18
        } else {
            var v19 string = "foo bar"
            v20 := GenS3{F0: 1, F1: (!v2)}
            _ = v19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    if true {
        switch ("foo bar" + "hello") {
        case "foo bar":
            switch ("a" + "") {
            case "":
                for fi21 := 0; fi21 < 19; fi21++ {
                    v22 := []int{}
                    _ = fi21
                    _ = v22
                }
            case "a":
                switch -242 {
                case -591:
                    var v23 int = 42
                    _ = v23
                }
            default:
                switch -1 {
                case -123:
                    switch "foo bar" {
                    case "hello":
                        v1 -= 42
                    case "a":
                        var v24 int16 = -32768
                        _ = v24
                    case "foo bar":
                        fmt.Println(v7, v3)
                    default:
                        for fi25 := 0; fi25 < 16; fi25++ {
                            for fi26 := 0; fi26 < 3; fi26++ {
                                var v27 string = ("hello" + "foo bar")
                                if v2 {
                                    switch ((v27 + (v27 + v27)) + ((v27 + v27) + v27)) {
                                    case "hello":
                                        if true {
                                            fmt.Println(v1, v2, fi25)
                                            v7 -= (1.5 + v7)
                                            var v28 uint16 = 32767
                                            _ = v28
                                        } else {
                                            v2 = ((1 > -1) || v2)
                                        }
                                        fmt.Println(fi25)
                                    case "foo bar":
                                        v29 := GenS1{F0: v2, F1: 571, F2: (fi25 * 0)}
                                        switch (fi26 | fi26) {
                                        case 9223372036854775807:
                                            v7 += v7
                                        default:
                                            fmt.Println(v2, v7, fi26)
                                            v30 := []int{(fi25 * fi25)}
                                            _ = v30
                                        }
                                        delete(v4, 128)
                                        _ = v29
                                    case "":
                                        v31 := GenS1{F0: v2, F1: 1, F2: 2147483647}
                                        _ = v31
                                    default:
                                        if true {
                                            v32 := GenS3{F0: 255, F1: false}
                                            _ = v32
                                        }
                                    }
                                    v1 = (v1 & -270)
                                } else {
                                    v27 = ("hello" + v27)
                                    v2 = (-2147483648 < 42)
                                    v33 := GenS1{F0: true, F1: 4294967295, F2: fi25}
                                    _ = v33
                                }
                                v34 := []int{(fi26 ^ fi26), 1}
                                _ = fi26
                                _ = v27
                                _ = v34
                            }
                            _ = fi25
                        }
                    }
                default:
                    v35 := GenS4{F0: 1, F1: 1, F2: -278}
                    _ = v35
                }
            }
        default:
            switch -1 {
            case -811:
                if (32767 >= 65535) {
                    for fi36 := 0; fi36 < 10; fi36++ {
                        var v37 int16 = 1000
                        v38 := GenS5{F0: (fi36 * (fi36 + 127)), F1: fi36, F2: "foo bar"}
                        switch "hello" {
                        case "\n\t":
                            var v39 int = 771
                            _ = v39
                        case "":
                            v40 := []int{fi36, 4294967295, -645}
                            _ = v40
                        case "hello":
                            if v2 {
                                v41 := []int{(fi36 % (((fi36 % ((127) | 1))) | 1)), fi36, (fi36 % ((fi36) | 1)), (fi36 >> ((0) & 63))}
                                v3 += 3.14159
                                v7 += float64(((-2.25 / 1.5) + (((1e10 - v3) / float32((v7 - v7))) - ((1e10 * 3.14159) * float32(0.1)))))
                                _ = v41
                            } else {
                                v37 += v37
                            }
                        }
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                    fmt.Println(v2, v7)
                } else {
                    v1 -= 1
                    v3 = 1.5
                }
            case 9223372036854775807:
                if false {
                    v42 := v4[9223372036854775807]
                    v43 := []int{2147483647}
                    if len(v43) > 0 {
                        v43[0] = 647
                    }
                    _ = v42
                    _ = v43
                } else {
                    var v44 float64 = float64(1.5)
                    var v45 uint = 0
                    var v46 float32 = 3.14159
                    _ = v44
                    _ = v45
                    _ = v46
                }
            default:
                v47 := GenS3{F0: 0, F1: (2147483648 < 0)}
                _ = v47
            }
        }
        var v48 int16 = 0
        for fi49 := 0; fi49 < 19; fi49++ {
            v50 := GenS4{F0: 322, F1: 42, F2: (^(fi49))}
            delete(v4, (fi49 >> ((65535) & 63)))
            _ = fi49
            _ = v50
        }
        _ = v48
    } else {
        fmt.Println(v3, v3)
    }
    fmt.Println(v2, v1)
    v3 = -2.25
    if (0 > -1) {
        v1 = (v1 / ((-956) | 1))
        v1 -= v1
    } else {
        _ = len(v4)
        v7 = v7
    }
    var v51 uint8 = 249
    for fi52 := 0; fi52 < 6; fi52++ {
        v53 := []int{32768, fi52, 0}
        switch (fi52 % ((32768) | 1)) {
        case 32767:
            v54 := GenS4{F0: 4294967295, F1: 0, F2: fi52}
            _ = v54
        default:
            var v55 uint64 = 1
            _ = v55
        }
        _ = fi52
        _ = v53
    }
    for fi56 := 0; fi56 < 12; fi56++ {
        v57 := GenS6{F0: (fi56 + -325), F1: 42, F2: 521}
        _ = fi56
        _ = v57
    }
    v58 := GenS5{F0: 1, F1: -9223372036854775808, F2: ("foo bar" + "a")}
    v59 := GenS7{F0: -2147483648}
    switch -9223372036854775808 {
    case 42:
        if false {
            v51 = 63
        }
    case 742:
        v1 += (v1 & 0)
    }
    v51 = (^(v51))
    var v60 int32 = 579
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    fmt.Println("v60:", v60)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v58
    _ = v59
    _ = v60
    _ = v7
}
