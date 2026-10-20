package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint8; F1 uint16; F2 uint32 }
func genFn0(p0 string) uint64 {
    fmt.Println(p0)
    switch 32767 {
    case 2147483647:
        var v7 uint = 255
        _ = v7
    case 190:
        p0 = p0
    case 127:
        v8 := []int{}
        _ = v8
    }
    switch (("foo bar" + (p0 + p0)) + (p0 + (("foo bar" + (p0 + p0)) + p0))) {
    case "hello":
        var v9 string = (((p0 + "a") + (p0 + "hello")) + "hello")
        _ = v9
    default:
        if ("foo bar" < "a") {
            v10 := []int{2147483647, -237}
            var v11 string = ((p0 + "") + p0)
            _ = v10
            _ = v11
        } else {
            var v12 int8 = -1
            v13 := GenS1{F0: 4294967295, F1: 0, F2: "a"}
            _ = v12
            _ = v13
        }
    }
    _ = p0
    return 4294967295
}

func main() {
    var v1 bool = true
    v2 := make(map[string]int)
    v3 := []uint32{65535, 131, 65535, 2147483648}
    var v4 int32 = 1
    var v5 int = -349
    var v6 uint64 = 18446744073709551615
    var v14 int8 = 117
    var v15 int64 = -935
    v4 = (v4 >> ((3) & 63))
    v16 := GenS1{F0: genFn0("\n\t"), F1: (v4 / ((-623) | 1)), F2: "\n\t"}
    for fi17 := 0; fi17 < 3; fi17++ {
        v6 -= v6
        switch (fi17 >> ((4294967295) & 63)) {
        case -797:
            v18 := v2["foo bar"]
            _ = v18
        }
        v19 := GenS1{F0: 2, F1: (v4 / ((v4) | 1)), F2: ("\n\t" + ("foo bar" + "\n\t"))}
        _ = fi17
        _ = v19
    }
    var v20 float32 = ((1.5 + (0.0 / 3.14159)) + 1.5)
    if ("hello" != "") {
        v20 = v20
    } else {
        for fi21 := 0; fi21 < 10; fi21++ {
            fmt.Println(v14, v4, v15)
            var v22 uint64 = genFn0(("hello" + ("" + (("\n\t" + "foo bar") + "hello"))))
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 uint32 = 4294967295
            if len(v3) > 0 {
                v3[0] = 2
            }
            switch ("\n\t" + "\n\t") {
            case "a":
                if len(v3) > 0 {
                    v3[0] = 435
                }
            case "hello":
                if false {
                    var v25 uint = 255
                    if v1 {
                        if (!((32767 <= 1000) || v1)) {
                            _ = len(v3)
                            var v26 bool = (-103 != -32768)
                            _ = v26
                        } else {
                            if (v1 && v1) {
                                v27 := GenS2{F0: (v6 ^ (v6 + v6)), F1: 202, F2: 461}
                                var v28 int16 = -1
                                v29 := GenS3{F0: (-(v15)), F1: (true || (708 > (v6 << ((4294967295) & 63))))}
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            if ((v15 ^ (v15 & -1)) < -831) {
                                v30 := GenS2{F0: v6, F1: 2, F2: v15}
                                switch (-(fi23)) {
                                case -1:
                                    delete(v2, "hello")
                                default:
                                    v31 := GenS4{F0: 81}
                                    _ = v31
                                }
                                v24 = 435
                                _ = v30
                            }
                            v1 = (!false)
                        }
                    }
                    switch fi23 {
                    case 0:
                        for fi32 := 0; fi32 < 10; fi32++ {
                            v14 = v14
                            delete(v2, "")
                            if v1 {
                                v20 += -2.25
                                v1 = true
                                if len(v3) > 0 {
                                    v3[0] = v24
                                }
                                v6 += uint64(v24)
                            }
                            if v1 {
                                v15 = 9223372036854775807
                                var v33 string = "a"
                                var v34 uint64 = v6
                                if (!false) {
                                    v25 -= v25
                                    if (!(-702 < -1)) {
                                        switch 1 {
                                        case 563:
                                            var v35 string = (v33 + "")
                                            _ = v35
                                        case -2147483648:
                                            v36 := v2[v33]
                                            _ = v36
                                        default:
                                            fmt.Println(v15)
                                        }
                                    } else {
                                        switch (fi32 << ((v25) & 63)) {
                                        case -222:
                                            switch 885 {
                                            case 127:
                                                if v1 {
                                                    v37 := GenS5{F0: ("foo bar" + "hello"), F1: true}
                                                    fmt.Println(fi23, v25, v15)
                                                    _ = v37
                                                } else {
                                                    fmt.Println(v6)
                                                    if false {
                                                        var v38 float32 = (v20 + (v20 + (v20 + v20)))
                                                        v1 = (v34 >= (v6 << ((uint(v6)) & 63)))
                                                        v39 := GenS6{F0: true}
                                                        _ = v38
                                                        _ = v39
                                                    } else {
                                                        var v40 float32 = (float32(0.1) * (v20 * 0.0))
                                                        v3 = append(v3, v24)
                                                        v15 -= v15
                                                        _ = v40
                                                    }
                                                    _ = len(v2)
                                                }
                                            case 343:
                                                var v41 int32 = v4
                                                _ = v41
                                            default:
                                                var v42 int = (v5 % (((fi32 * (fi32 >> ((v25) & 63)))) | 1))
                                                _ = v42
                                            }
                                        case 81:
                                            _ = len(v2)
                                        default:
                                            v15 -= -888
                                        }
                                        var v43 int = (fi23 * fi23)
                                        var v44 int16 = 972
                                        _ = v43
                                        _ = v44
                                    }
                                    var v45 int32 = 497
                                    if true {
                                        v1 = ((v45 / ((v45) | 1)) <= (v45 | v4))
                                    } else {
                                        if v1 {
                                            fmt.Println(v33, v15, v45)
                                        } else {
                                            fmt.Println(v4, v1)
                                            fmt.Println(v33, v14)
                                            if ((!(-32768 < 1000)) || (!(!(-32768 > 1)))) {
                                                v14 -= (v14 | v14)
                                                v46 := GenS7{F0: 0}
                                                _ = v46
                                            }
                                        }
                                        v47 := GenS1{F0: 158, F1: v45, F2: ((v33 + ("" + v33)) + (v33 + (v33 + v33)))}
                                        var v48 int32 = 0
                                        _ = v47
                                        _ = v48
                                    }
                                    fmt.Println(v24, v4, v33)
                                    _ = v45
                                } else {
                                    if v1 {
                                        var v49 int32 = v4
                                        _ = v49
                                    } else {
                                        var v50 int64 = (-(v15))
                                        v51 := GenS1{F0: v34, F1: int32(fi23), F2: ("\n\t" + "a")}
                                        _ = v50
                                        _ = v51
                                    }
                                    var v52 int32 = (v4 - (v4 + 42))
                                    var v53 int = (v5 & fi23)
                                    _ = v52
                                    _ = v53
                                }
                                var v54 bool = v1
                                _ = v33
                                _ = v34
                                _ = v54
                            }
                            _ = fi32
                        }
                    }
                    _ = v25
                }
            case "":
                var v55 int = (fi23 * int(v4))
                _ = v55
            default:
                v4 -= (v4 | v4)
            }
            _ = fi23
            _ = v24
        }
        if (42 == (^(v5))) {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            fmt.Println(v4, v6)
        }
    }
    var v56 uint8 = 236
    v56 -= v56
    var v57 uint16 = 255
    v58 := GenS1{F0: 0, F1: (v4 ^ 0), F2: "\n\t"}
    v56 = v56
    v59 := GenS8{F0: v56, F1: uint16(v56), F2: uint32(v6)}
    v56 += v56
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
}
