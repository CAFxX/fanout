package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int; F2 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int32; F2 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 uint8) string {
    if true {
        var v6 uint32 = 65535
        var v7 int64 = -9223372036854775808
        var v8 uint8 = p0
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 int16 = 42
    _ = p0
    _ = v9
    return "foo bar"
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -1
    var v3 uint16 = 1
    var v4 bool = (0.0 >= (((3.14159 + 0.0) - (1.5 + (-2.25 - -2.25))) / 3.14159))
    var v5 bool = v4
    v10 := v1[88]
    v11 := []int{(v10 % (((v10 * v10)) | 1)), (-(v10)), v10, v10}
    var v12 int8 = 1
    var v13 uint16 = 627
    switch ("hello" + "a") {
    case "foo bar":
        if true {
            if v4 {
                var v14 int64 = -1
                v2 += (v2 * v2)
                _ = v14
            }
            v1[(v10 * 4294967295)] = (v10 | (^(v10)))
            v3 = (v3 % ((v3) | 1))
        } else {
            v4 = (!(v4 && v4))
        }
    default:
        var v15 string = "foo bar"
        _ = v15
    }
    var v16 int32 = v2
    if len(v11) > 0 {
        v11[0] = (v10 % (((v10 | int(v12))) | 1))
    }
    if len(v11) > 0 {
        v11[0] = (v10 >> ((65535) & 63))
    }
    fmt.Println(v4)
    if false {
        switch v10 {
        case -9223372036854775808:
            v3 -= (v3 % ((459) | 1))
        case 32767:
            if (!(862 < 42)) {
                _ = len(v11)
                if ((1e10 + -2.25) <= ((1.5 / 3.14159) + 3.14159)) {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        if false {
                            if len(v11) > 0 {
                                v11[0] = (^(fi17))
                            }
                            v5 = v5
                        }
                        v18 := GenS1{F0: (^(v16))}
                        _ = fi17
                        _ = v18
                    }
                } else {
                    v4 = ((!v4) && (!false))
                    if (-1 < 0) {
                        v19 := GenS1{F0: (v16 + 2147483647)}
                        for fi20 := 0; fi20 < 2; fi20++ {
                            var v21 uint8 = 219
                            if len(v11) > 0 {
                                v11[0] = (^(v10))
                            }
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    } else {
                        if v5 {
                            var v22 bool = false
                            v23 := GenS2{F0: 286, F1: 1}
                            _ = v22
                            _ = v23
                        } else {
                            if true {
                                var v24 bool = v4
                                if v5 {
                                    v24 = ("\n\t" <= "\n\t")
                                    if v5 {
                                        v25 := GenS1{F0: v2}
                                        if true {
                                            v26 := GenS2{F0: -189, F1: 1}
                                            if false {
                                                for fi27 := 0; fi27 < 18; fi27++ {
                                                    if (("hello" == "foo bar") || v24) {
                                                        var v28 uint8 = 246
                                                        if (!(1.5 >= -2.25)) {
                                                            v12 = -1
                                                        }
                                                        _ = v28
                                                    } else {
                                                        var v29 uint32 = 2147483648
                                                        _ = v29
                                                    }
                                                    var v30 float32 = 1.5
                                                    _ = fi27
                                                    _ = v30
                                                }
                                                for fi31 := 0; fi31 < 7; fi31++ {
                                                    switch ("a" + "\n\t") {
                                                    case "":
                                                        v10 += 32767
                                                    case "hello":
                                                        if len(v11) >= 2 {
                                                            v11 = v11[1:len(v11)]
                                                        }
                                                    }
                                                    v13 += v3
                                                    _ = fi31
                                                }
                                            } else {
                                                switch "" {
                                                case "a":
                                                    fmt.Println(v2)
                                                case "\n\t":
                                                    v32 := v1[v10]
                                                    _ = v32
                                                case "":
                                                    var v33 uint16 = v3
                                                    _ = v33
                                                default:
                                                    var v34 string = "a"
                                                    _ = v34
                                                }
                                                v35 := GenS3{F0: 280, F1: int(v13), F2: 4294967295}
                                                _ = v35
                                            }
                                            _ = v26
                                        }
                                        if true {
                                            v36 := GenS4{F0: 42}
                                            v37 := GenS5{F0: (-(v10)), F1: -2147483648, F2: 0}
                                            _ = v36
                                            _ = v37
                                        }
                                        _ = v25
                                    }
                                } else {
                                    v2 = (v2 & v16)
                                    if len(v11) > 0 {
                                        v11[0] = (v10 & (-(v10)))
                                    }
                                }
                                _ = v24
                            } else {
                                v38 := v1[(v10 % (((v10 + v10)) | 1))]
                                v39 := v1[v38]
                                var v40 int = (v10 ^ 127)
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                        }
                    }
                }
            } else {
                for fi41 := 0; fi41 < 2; fi41++ {
                    v42 := GenS6{F0: 333}
                    _ = fi41
                    _ = v42
                }
            }
        case -625:
            v12 -= 42
        default:
            v5 = (!v5)
        }
        v11 = append(v11, v10)
        var v43 int16 = 32767
        _ = v43
    }
    v16 -= (-(v16))
    switch (("\n\t" + "") + "a") {
    case "hello":
        var v44 uint32 = 4294967295
        _ = v44
    }
    var v45 uint64 = 691
    var v46 int64 = -9223372036854775808
    for fi47 := 0; fi47 < 11; fi47++ {
        v4 = (fi47 != fi47)
        if (-364 != -378) {
            var v48 uint16 = 255
            _ = v48
        }
        _ = fi47
    }
    v49 := GenS7{F0: 507}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v49
    _ = v5
}
