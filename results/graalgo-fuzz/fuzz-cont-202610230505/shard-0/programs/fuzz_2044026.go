package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint }

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    var v3 uint = 18446744073709551615
    switch 32768 {
    case -399:
        if v1 {
            var v4 uint16 = 0
            _ = len(v2)
            _ = v4
        }
    case 121:
        v5 := GenS1{F0: true}
        _ = v5
    default:
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := []int{fi6, (fi6 - 42), (^(fi6)), (fi6 >> ((42) & 63))}
            _ = fi6
            _ = v7
        }
    }
    var v8 uint64 = 577
    var v9 int64 = 9223372036854775807
    if (!v1) {
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                switch "" {
                case "":
                    if v1 {
                        v12 := GenS2{F0: fi10, F1: fi10}
                        v3 = v3
                        _ = v12
                    } else {
                        v9 = 0
                    }
                default:
                    var v13 int32 = -41
                    _ = v13
                }
                _ = fi11
            }
            switch "\n\t" {
            case "\n\t":
                v14 := []int{}
                _ = v14
            case "a":
                var v15 uint64 = (v8 | (v8 & 1))
                _ = v15
            default:
                var v16 bool = (1.5 != (0.1 * 0.1))
                _ = v16
            }
            _ = fi10
        }
        if v1 {
            v9 = (v9 | 9223372036854775807)
            for fi17 := 0; fi17 < 14; fi17++ {
                if (((((float32(1e10) > 3.14159) || (false || (true && (uint32(v8) > 2)))) || (("foo bar" + "\n\t") >= (("" + "") + "a"))) && v1) && ("a" <= "hello")) {
                    var v18 float32 = (float32(-2.25) / 1.0)
                    var v19 int = (fi17 * fi17)
                    if v1 {
                        fmt.Println(fi17)
                        if (v9 == (-(v9))) {
                            switch ("foo bar" + ("\n\t" + ("hello" + ("a" + "")))) {
                            case "":
                                var v20 uint = (^(v3))
                                _ = v20
                            }
                            fmt.Println(v18, v9)
                            if (!v1) {
                                v21 := GenS1{F0: (3.14159 < float64(v18))}
                                _ = v21
                            } else {
                                for fi22 := 0; fi22 < 10; fi22++ {
                                    v8 = v8
                                    var v23 uint16 = 65535
                                    _ = fi22
                                    _ = v23
                                }
                            }
                        }
                    }
                    _ = v18
                    _ = v19
                } else {
                    if false {
                        switch int(v9) {
                        case 0:
                            var v24 string = "hello"
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 9; fi25++ {
                            fmt.Println(fi17, fi25, v1)
                            v9 = 9223372036854775807
                            _ = fi25
                        }
                    }
                }
                _ = len(v2)
                _ = fi17
            }
            if (!(2147483647 >= 507)) {
                var v26 int64 = v9
                _ = v26
            }
        }
        if v1 {
            if false {
                v27 := GenS2{F0: 127, F1: 0}
                _ = v27
            } else {
                var v28 int32 = 2147483647
                if v1 {
                    v3 = v3
                    var v29 float64 = 0.0
                    _ = v29
                }
                _ = v28
            }
            var v30 uint = 65535
            switch 32767 {
            case 1:
                switch 0 {
                case -841:
                    var v31 int = 127
                    _ = v31
                case -224:
                    if v1 {
                        v32 := GenS3{F0: (v3 & (v30 << ((18446744073709551615) & 63)))}
                        _ = v32
                    }
                case -9223372036854775808:
                    v30 += 255
                }
            }
            _ = v30
        }
    } else {
        var v33 uint = 392
        if true {
            if v1 {
                var v34 bool = v1
                _ = v34
            } else {
                switch ("a" + "") {
                case "\n\t":
                    v8 += v8
                case "a":
                    var v35 uint32 = 1
                    _ = v35
                }
                switch ((("a" + "a") + "\n\t") + "\n\t") {
                case "foo bar":
                    v36 := GenS2{F0: 926, F1: 128}
                    _ = v36
                case "":
                    v8 = 18446744073709551615
                default:
                    var v37 int32 = 776
                    _ = v37
                }
                if (42 != 415) {
                    if false {
                        fmt.Println(v1, v8)
                        var v38 uint8 = 1
                        fmt.Println(v3)
                        _ = v38
                    }
                    v9 += (v9 * (v9 & -9223372036854775808))
                } else {
                    v8 -= (v8 / ((v8) | 1))
                    var v39 int8 = 1
                    v1 = (!v1)
                    _ = v39
                }
            }
            v3 = (^(v3))
            v3 -= (v3 ^ (v33 | (v33 ^ (^(v33)))))
        } else {
            for fi40 := 0; fi40 < 8; fi40++ {
                _ = len(v2)
                _ = fi40
            }
            switch ("foo bar" + "hello") {
            case "hello":
                var v41 uint16 = 65535
                _ = v41
            case "":
                v9 = (v9 ^ v9)
            default:
                for fi42 := 0; fi42 < 17; fi42++ {
                    fmt.Println(v8, fi42, v3)
                    v33 = v33
                    v1 = (1e10 <= (0.0 + 0.1))
                    _ = fi42
                }
            }
        }
        _ = v33
    }
    v2["foo bar"] = ("foo bar" + "a")
    v43 := GenS1{F0: false}
    for fi44 := 0; fi44 < 8; fi44++ {
        var v45 uint64 = 0
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 3; fi46++ {
        if (100 >= -115) {
            for fi47 := 0; fi47 < 1; fi47++ {
                var v48 float32 = 1e10
                fmt.Println(v8, v9)
                if (2147483647 == 0) {
                    v3 -= 428
                    v49 := GenS2{F0: (^(fi47)), F1: 143}
                    _ = v49
                } else {
                    v50 := []int{int(v3), fi46}
                    _ = v50
                }
                v1 = false
                _ = fi47
                _ = v48
            }
            var v51 int = 4294967295
            _ = v51
        } else {
            v52 := []int{-819, fi46, fi46, (fi46 & fi46)}
            if ((((0.1 / float32(3.14159)) - 1e10) * 3.14159) <= 1.5) {
                if (!(65535 < 657)) {
                    if len(v52) >= 2 {
                        v52 = v52[1:len(v52)]
                    }
                    var v53 int = (fi46 & fi46)
                    _ = v53
                } else {
                    if (-90 != (fi46 * (fi46 + fi46))) {
                        switch "hello" {
                        case "a":
                            if v1 {
                                v3 = v3
                            }
                        case "\n\t":
                            v52 = append(v52, 4294967295)
                        case "":
                            v8 -= (v8 + (^(v8)))
                        }
                        var v54 int = fi46
                        _ = len(v52)
                        _ = v54
                    } else {
                        v9 = -9223372036854775808
                    }
                    v8 = v8
                }
                v2[("foo bar" + ("foo bar" + (("foo bar" + (("\n\t" + "a") + "")) + ("foo bar" + ("hello" + "")))))] = ("hello" + "foo bar")
            } else {
                v55 := GenS3{F0: v3}
                _ = v55
            }
            _ = len(v2)
            _ = v52
        }
        v1 = (!(4294967295 != 4294967295))
        switch ("" + (("" + "hello") + (("" + "hello") + "\n\t"))) {
        case "a":
            if v1 {
                delete(v2, ("a" + "\n\t"))
            } else {
                if false {
                    for fi56 := 0; fi56 < 2; fi56++ {
                        v57 := []int{(^(fi56))}
                        _ = fi56
                        _ = v57
                    }
                    for fi58 := 0; fi58 < 5; fi58++ {
                        v2[""] = (("\n\t" + "a") + "foo bar")
                        v59 := v2["a"]
                        var v60 int8 = int8(fi46)
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                }
            }
        case "\n\t":
            fmt.Println(v3, v8)
        case "":
            v9 -= v9
        default:
            v61 := []int{(fi46 + fi46), fi46, fi46, (^(fi46))}
            _ = v61
        }
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v8
    _ = v9
}
