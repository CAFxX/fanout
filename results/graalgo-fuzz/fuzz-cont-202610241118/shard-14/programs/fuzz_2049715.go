package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint32; F2 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64; F1 bool; F2 uint }
func genFn0(p0 int32) uint64 {
    v5 := []int{0, -9223372036854775808}
    fmt.Println(p0)
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: p0, F1: 159, F2: 62}
        var v8 string = "a"
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: p0, F1: 4294967295, F2: 15}
    _ = p0
    _ = v5
    _ = v9
    return 355
}

func main() {
    v1 := []string{"hello", (("hello" + ("" + ("hello" + "a"))) + "")}
    v2 := make(map[string]int)
    var v3 uint32 = 1
    var v4 uint64 = 1
    if false {
        var v10 uint8 = 1
        _ = v10
    }
    var v11 float64 = 0.0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v11 = v11
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v12 := GenS1{F0: 2147483647, F1: 59, F2: 128}
    var v13 string = ("foo bar" + "\n\t")
    if (false && false) {
        _ = len(v2)
        fmt.Println(v11)
        var v14 uint64 = genFn0(-2147483648)
        _ = v14
    } else {
        _ = len(v2)
    }
    var v15 uint32 = 4294967295
    var v16 uint32 = 710
    _ = len(v1)
    if false {
        for fi17 := 0; fi17 < 2; fi17++ {
            v18 := GenS2{F0: 128}
            var v19 float64 = ((v11 - float64((3.14159 / 0.1))) * float64(1.5))
            _ = fi17
            _ = v18
            _ = v19
        }
        if true {
            v11 = v11
            v20 := GenS3{F0: (!(!true)), F1: v3, F2: -572}
            for fi21 := 0; fi21 < 1; fi21++ {
                v16 += v3
                if true {
                    if (false && ((v4 << ((4294967295) & 63)) > v4)) {
                        if true {
                            for fi22 := 0; fi22 < 2; fi22++ {
                                v16 = (v16 + v3)
                                fmt.Println(fi21, v11, v3)
                                var v23 uint32 = (v15 / ((42) | 1))
                                var v24 float32 = float32(0.1)
                                var v25 int64 = -102
                                _ = fi22
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        }
                        if true {
                            var v26 uint64 = v4
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                for fi27 := 0; fi27 < 13; fi27++ {
                                    v13 = ((v13 + "\n\t") + "a")
                                    var v28 int64 = 42
                                    fmt.Println(v15, v15)
                                    _ = fi27
                                    _ = v28
                                }
                            } else {
                                v16 = v16
                                if false {
                                    var v29 string = v13
                                    var v30 float64 = float64(-2.25)
                                    var v31 float32 = 1e10
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                } else {
                                    _ = len(v2)
                                }
                                fmt.Println(fi21)
                            }
                            fmt.Println(v13)
                            _ = v26
                        } else {
                            _ = len(v2)
                            switch (v13 + v13) {
                            case "":
                                v4 += v4
                            case "\n\t":
                                v11 += ((1.5 + v11) + 0.0)
                            }
                            var v32 int32 = -2147483648
                            _ = v32
                        }
                        v33 := GenS3{F0: true, F1: (v16 % ((v15) | 1)), F2: 2147483647}
                        _ = v33
                    }
                    var v34 int8 = -83
                    _ = v34
                } else {
                    _ = len(v2)
                    if true {
                        for fi35 := 0; fi35 < 11; fi35++ {
                            delete(v2, ((v13 + "a") + ("foo bar" + v13)))
                            switch v13 {
                            case "":
                                v36 := GenS1{F0: 2147483647, F1: 2, F2: 236}
                                _ = v36
                            case "hello":
                                var v37 int8 = -29
                                _ = v37
                            case "\n\t":
                                var v38 float32 = (((1e10 - float32(0.1)) / ((float32(0.0) - (float32(v11) + 1e10)) / float32(float64(3.14159)))) - (((3.14159 / 1e10) * 3.14159) / -2.25))
                                _ = v38
                            default:
                                var v39 uint32 = (v15 << ((575) & 63))
                                _ = v39
                            }
                            v16 -= v15
                            _ = fi35
                        }
                    } else {
                        for fi40 := 0; fi40 < 1; fi40++ {
                            _ = len(v2)
                            v41 := GenS4{F0: v13}
                            v11 -= 0.0
                            _ = fi40
                            _ = v41
                        }
                    }
                    v42 := GenS1{F0: 1, F1: 18446744073709551615, F2: 255}
                    _ = v42
                }
                _ = len(v1)
                _ = fi21
            }
            _ = v20
        }
        for fi43 := 0; fi43 < 14; fi43++ {
            var v44 uint8 = 255
            _ = len(v1)
            for fi45 := 0; fi45 < 4; fi45++ {
                delete(v2, v13)
                v2["\n\t"] = 127
                switch 128 {
                case -473:
                    var v46 uint = 65535
                    _ = v46
                case 32768:
                    switch fi43 {
                    case 0:
                        switch (fi43 << ((4294967295) & 63)) {
                        case 1:
                            v47 := GenS4{F0: (((v13 + "hello") + v13) + "a")}
                            _ = v47
                        case 32768:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 128:
                            if false {
                                switch (-(fi45)) {
                                case -9223372036854775808:
                                    switch v13 {
                                    case "hello":
                                        _ = len(v2)
                                    case "\n\t":
                                        v4 = v4
                                        _ = len(v2)
                                        v16 += (v16 * v3)
                                    default:
                                        v11 = ((((1e10 + 3.14159) / ((v11 - float64((1.5 + 1e10))) - v11)) + 1e10) + float64(0.0))
                                    }
                                case 203:
                                    if (true && false) {
                                        v48 := v2[v13]
                                        var v49 string = v13
                                        _ = v48
                                        _ = v49
                                    }
                                }
                            } else {
                                v4 = (v4 << ((42) & 63))
                            }
                        default:
                            var v50 uint8 = (v44 << ((1) & 63))
                            _ = v50
                        }
                    case -937:
                        if (false && (!(!true))) {
                            fmt.Println(v15, v44, v16)
                            v51 := v2["hello"]
                            fmt.Println(v11, v44)
                            fmt.Println(v3, v51, v13)
                            _ = v51
                        }
                    }
                case 4294967295:
                    var v52 string = v13
                    _ = v52
                default:
                    v53 := GenS2{F0: v44}
                    _ = v53
                }
                _ = fi45
            }
            _ = fi43
            _ = v44
        }
    }
    var v54 bool = false
    var v55 uint = 42
    switch 32768 {
    case 128:
        var v56 bool = (v13 == v13)
        _ = v56
    case -2147483648:
        var v57 string = v13
        _ = v57
    }
    v13 = v13
    switch v13 {
    case "foo bar":
        v58 := GenS5{F0: v13}
        _ = v58
    case "\n\t":
        v59 := GenS6{F0: 65535}
        _ = v59
    case "hello":
        v60 := GenS7{F0: 42, F1: false, F2: (^(v55))}
        _ = v60
    }
    fmt.Println(v3, v55, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
}
