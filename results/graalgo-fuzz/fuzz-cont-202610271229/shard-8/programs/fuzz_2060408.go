package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint, p1 string, p2 int64) int64 {
    p1 = p1
    p0 -= 255
    var v6 uint32 = 277
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}

func main() {
    var v1 int8 = 1
    v2 := make(map[int]uint32)
    var v3 int32 = 2147483647
    var v4 int32 = (v3 >> ((520) & 63))
    var v5 int = 9223372036854775807
    var v7 float64 = (3.14159 - float64((1e10 / (-2.25 / (0.0 - 3.14159)))))
    var v8 int32 = v4
    v7 -= (-2.25 * ((v7 - ((v7 * v7) + (v7 - float64(1.5)))) - 1e10))
    switch ("foo bar" + "") {
    case "\n\t":
        var v9 int64 = 1
        _ = v9
    case "":
        v10 := GenS1{F0: "\n\t"}
        _ = v10
    case "foo bar":
        fmt.Println(v5)
    }
    var v11 int16 = -1
    var v12 int8 = (v1 % (((v1 >> ((264) & 63))) | 1))
    v13 := []int{v5}
    for fi14 := 0; fi14 < 17; fi14++ {
        if ((float64(((1.5 * 1.5) - -2.25)) + 0.0) == (((float64((((-2.25 - 3.14159) / (1e10 - 0.1)) - 0.0)) - (v7 + 1.5)) * ((((v7 / float64(((1.5 / ((1.5 + 0.0) / 3.14159)) + 3.14159))) / ((((v7 / 1.5) + v7) + v7) * 1e10)) / -2.25) + (v7 / 1.5))) + 0.1)) {
            switch -265 {
            case -9223372036854775808:
                delete(v2, -9223372036854775808)
            case 2147483647:
                v3 = int32(v5)
            default:
                if true {
                    var v15 int = (^(fi14))
                    _ = v15
                }
            }
        } else {
            if (65535 >= 1) {
                v5 += (v5 >> ((146) & 63))
                fmt.Println(v1, v12)
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            } else {
                var v16 float32 = -2.25
                v17 := v2[fi14]
                v18 := v2[9223372036854775807]
                _ = v16
                _ = v17
                _ = v18
            }
            if (2 > 378) {
                delete(v2, -2147483648)
            } else {
                v12 += v12
                v5 -= (v5 >> ((65535) & 63))
                if (!((false && false) || true)) {
                    v13 = append(v13, (fi14 & 9223372036854775807))
                }
            }
        }
        _ = fi14
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v4, v8, v5)
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 int8 = (v1 & v12)
            if false {
                switch 32768 {
                case 2147483647:
                    v22 := GenS1{F0: "hello"}
                    _ = v22
                default:
                    switch "\n\t" {
                    case "\n\t":
                        if len(v13) > 0 {
                            v13[0] = fi20
                        }
                    }
                    v23 := GenS1{F0: "foo bar"}
                    var v24 uint16 = 581
                    _ = v23
                    _ = v24
                }
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    fmt.Println(v4, v5)
                default:
                    if false {
                        if (!false) {
                            v3 = (v4 * -274)
                        } else {
                            if (!((fi20 >> ((42) & 63)) <= (v5 | (fi19 / ((1) | 1))))) {
                                if true {
                                    v25 := GenS2{F0: -57, F1: fi20, F2: 127}
                                    if true {
                                        var v26 float64 = ((v7 + 0.1) / (1e10 + 0.0))
                                        var v27 uint16 = 42
                                        _ = v26
                                        _ = v27
                                    } else {
                                        v28 := GenS3{F0: ("a" + "\n\t")}
                                        var v29 uint32 = 2147483648
                                        fmt.Println(v12, v12)
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = v25
                                }
                                v5 = 32767
                                if (!true) {
                                    v5 += 212
                                    delete(v2, (v5 / ((533) | 1)))
                                    v30 := v2[(-(fi19))]
                                    _ = v30
                                }
                            } else {
                                v21 -= (^(v21))
                                fmt.Println(v7)
                                v31 := GenS3{F0: (("\n\t" + "a") + ("\n\t" + ((((((("" + ("\n\t" + ("a" + ("foo bar" + "foo bar")))) + "a") + ("a" + "foo bar")) + "hello") + "hello") + "hello") + ("foo bar" + ("hello" + "\n\t")))))}
                                _ = v31
                            }
                            fmt.Println(fi19)
                        }
                        v32 := GenS3{F0: "a"}
                        v4 += (v3 / ((1) | 1))
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = v32
                    } else {
                        fmt.Println(fi19, fi19, v3)
                        fmt.Println(fi19)
                        v13 = append(v13, fi20)
                    }
                }
            }
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            v33 := GenS1{F0: "hello"}
            var v34 int = 32767
            _ = fi20
            _ = v21
            _ = v33
            _ = v34
        }
        _ = fi19
    }
    switch "foo bar" {
    case "":
        for fi35 := 0; fi35 < 8; fi35++ {
            for fi36 := 0; fi36 < 9; fi36++ {
                fmt.Println(v8, v5)
                var v37 int32 = 2147483647
                _ = fi36
                _ = v37
            }
            for fi38 := 0; fi38 < 4; fi38++ {
                var v39 float32 = 0.1
                switch -170 {
                case 898:
                    var v40 bool = (0 <= (v1 | (v1 ^ int8(v11))))
                    _ = v40
                case 127:
                    v5 += (fi35 ^ 4294967295)
                }
                v11 = v11
                _ = fi38
                _ = v39
            }
            _ = fi35
        }
    case "\n\t":
        var v41 int16 = (v11 & v11)
        _ = v41
    default:
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        for fi43 := 0; fi43 < 6; fi43++ {
            if (9223372036854775807 < 42) {
                v5 += (v5 | (v5 << ((42) & 63)))
            } else {
                _ = len(v13)
                v7 = (v7 * v7)
                switch (fi43 * (v5 << ((0) & 63))) {
                case -2147483648:
                    v3 = 1
                    v44 := GenS1{F0: ("a" + (("hello" + "hello") + "\n\t"))}
                    _ = v44
                case 32768:
                    v7 = v7
                case 9223372036854775807:
                    v11 = -32768
                default:
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v8, v11)
                    default:
                        v45 := GenS2{F0: int64(v11), F1: v5, F2: (fi42 * int(v1))}
                        _ = v45
                    }
                }
                if true {
                    if true {
                        var v46 int64 = -580
                        v47 := GenS3{F0: "hello"}
                        fmt.Println(v11, v8, v11)
                        _ = v46
                        _ = v47
                    }
                    var v48 int64 = 9223372036854775807
                    _ = v48
                } else {
                    v49 := GenS2{F0: 42, F1: (fi43 + fi43), F2: (fi43 * v5)}
                    _ = v49
                }
            }
            _ = fi43
        }
        for fi50 := 0; fi50 < 20; fi50++ {
            v1 = v12
            v51 := v2[(fi42 % (((fi50 % (((fi50 % (((v5 % ((fi42) | 1))) | 1))) | 1))) | 1))]
            var v52 int16 = v11
            _ = fi50
            _ = v51
            _ = v52
        }
        _ = fi42
    }
    _ = len(v2)
    if ((!false) || (32767 > 71)) {
        for fi53 := 0; fi53 < 17; fi53++ {
            v13 = append(v13, v5)
            var v54 string = "a"
            _ = fi53
            _ = v54
        }
    } else {
        v55 := GenS4{F0: -1, F1: 4294967295}
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
