package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int; F2 int }
type GenS5 struct { F0 int32; F1 int }
func genFn0(p0 bool, p1 int32, p2 uint8) uint {
    v7 := []int{127, 9223372036854775807, 2147483647, 4294967295}
    fmt.Println(p0, p2)
    if p0 {
        if len(v7) > 0 {
            v7[0] = 675
        }
    }
    if (0.1 < (-2.25 + 1e10)) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 4294967295
}
func genFn1() int16 {
    var v8 uint32 = 4294967295
    v8 += (v8 / ((v8) | 1))
    v9 := GenS1{F0: uint16(v8)}
    var v10 int = 32767
    var v11 uint32 = (v8 >> ((255) & 63))
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 bool = (false && ((-273 != -1) || true))
    v2 := make(map[string]uint32)
    var v3 int64 = 0
    var v4 uint16 = 293
    v5 := make(map[int]string)
    v6 := []string{"a"}
    v12 := GenS1{F0: (^(v4))}
    v13 := GenS2{F0: 0, F1: 9223372036854775807}
    v14 := GenS2{F0: 334, F1: (v3 & v3)}
    var v15 bool = false
    fmt.Println(v1)
    v16 := GenS3{F0: 0, F1: 255}
    switch ("a" + ("\n\t" + "hello")) {
    case "hello":
        v15 = ((!false) || (true || false))
    case "\n\t":
        v1 = (1.5 < (-2.25 + (-2.25 * (-2.25 / (float64((0.1 + 0.1)) - 1e10)))))
    case "a":
        v17 := GenS1{F0: 328}
        _ = v17
    default:
        fmt.Println(v15, v3, v15)
    }
    if ((652 < -9223372036854775808) || false) {
        if (-9223372036854775808 <= -1) {
            var v18 uint8 = 127
            if v15 {
                for fi19 := 0; fi19 < 6; fi19++ {
                    fmt.Println(v3, v18)
                    var v20 string = (("a" + (((("foo bar" + ("\n\t" + "hello")) + "foo bar") + "") + (("foo bar" + ("hello" + ("" + "foo bar"))) + ("" + ("foo bar" + "hello"))))) + "a")
                    v18 = v18
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 2; fi21++ {
                    if false {
                        v22 := GenS1{F0: v4}
                        fmt.Println(v1)
                        _ = v22
                    } else {
                        fmt.Println(v3, fi21, v15)
                        for fi23 := 0; fi23 < 1; fi23++ {
                            if false {
                                delete(v2, "hello")
                                var v24 uint16 = v4
                                _ = v24
                            } else {
                                var v25 uint16 = (v4 | uint16(v18))
                                if (fi21 < (fi23 >> ((65535) & 63))) {
                                    v26 := GenS3{F0: fi23, F1: v18}
                                    switch ("\n\t" + "") {
                                    case "\n\t":
                                        var v27 int64 = v3
                                        _ = v27
                                    case "hello":
                                        switch fi23 {
                                        case 9223372036854775807:
                                            v28 := GenS2{F0: 646, F1: 1}
                                            _ = v28
                                        case 32768:
                                            var v29 uint16 = 0
                                            _ = v29
                                        case 334:
                                            v4 = (v25 & (v4 * v4))
                                        }
                                    }
                                    v4 += 528
                                    _ = v26
                                } else {
                                    v3 = 56
                                }
                                var v30 int32 = 0
                                var v31 int32 = 2147483647
                                switch "a" {
                                case "hello":
                                    var v32 uint32 = 991
                                    _ = v32
                                case "\n\t":
                                    var v33 bool = (!((false && (0 <= (v18 >> ((18446744073709551615) & 63)))) && (true && (1 <= 2))))
                                    _ = v33
                                }
                                var v34 string = "hello"
                                _ = v25
                                _ = v30
                                _ = v31
                                _ = v34
                            }
                            if (1 == 4294967295) {
                                var v35 string = "\n\t"
                                var v36 uint32 = 816
                                v36 += (v36 % ((uint32(v18)) | 1))
                                _ = v35
                                _ = v36
                            } else {
                                if v15 {
                                    var v37 int64 = v3
                                    var v38 uint16 = (v4 ^ 816)
                                    switch 32768 {
                                    case 126:
                                        v39 := GenS1{F0: v4}
                                        _ = v39
                                    case 1:
                                        if ("hello" <= ("foo bar" + "a")) {
                                            var v40 int32 = 2147483647
                                            if true {
                                                v41 := GenS1{F0: (v4 >> ((255) & 63))}
                                                _ = v41
                                            }
                                            delete(v5, (fi21 % (((fi21 ^ (fi21 / (((fi23 - fi23)) | 1)))) | 1)))
                                            v1 = v15
                                            _ = v40
                                        } else {
                                            v1 = v15
                                        }
                                    default:
                                        var v42 uint = 157
                                        _ = v42
                                    }
                                    _ = v37
                                    _ = v38
                                }
                                _ = len(v6)
                                if v15 {
                                    switch fi23 {
                                    case -9223372036854775808:
                                        delete(v2, "")
                                    }
                                } else {
                                    var v43 uint = uint(v4)
                                    fmt.Println(v43, v15)
                                    var v44 int8 = 42
                                    v3 -= v3
                                    _ = v43
                                    _ = v44
                                }
                            }
                            fmt.Println(v3, v18, v1)
                            fmt.Println(fi23)
                            if (v1 && v1) {
                                var v45 int16 = 42
                                _ = v45
                            } else {
                                if v1 {
                                    var v46 uint8 = v18
                                    fmt.Println(v4)
                                    v3 += v3
                                    _ = v46
                                } else {
                                    if ("" <= "\n\t") {
                                        if v1 {
                                            if (!v15) {
                                                v18 = (^(v18))
                                                v18 += v18
                                            }
                                            var v47 int32 = -1
                                            var v48 string = "foo bar"
                                            switch v48 {
                                            case "hello":
                                                fmt.Println(v15, v1)
                                            case "":
                                                if true {
                                                    v49 := GenS2{F0: 1, F1: v3}
                                                    _ = v49
                                                } else {
                                                    switch -248 {
                                                    case 4294967295:
                                                        if len(v6) > 0 {
                                                            v6[0] = v48
                                                        }
                                                    default:
                                                        var v50 float64 = float64((3.14159 / 1.0))
                                                        _ = v50
                                                    }
                                                }
                                                fmt.Println(fi21, fi21)
                                            case "a":
                                                switch (fi21 << ((18446744073709551615) & 63)) {
                                                case 1:
                                                    fmt.Println(v15)
                                                case -9223372036854775808:
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v48 = v48
                                                    }
                                                    var v51 int16 = 1000
                                                    v6 = append(v6, ((v48 + ((v48 + "\n\t") + (v48 + "foo bar"))) + v48))
                                                    _ = v51
                                                }
                                            }
                                            _ = v47
                                            _ = v48
                                        } else {
                                            var v52 int64 = (v3 << ((2) & 63))
                                            v3 = v52
                                            _ = v52
                                        }
                                        if false {
                                            delete(v5, (fi23 << ((genFn0(true, 83, (v18 & v18))) & 63)))
                                            v53 := GenS4{F0: 65535, F1: (fi21 ^ 1), F2: (-(fi23))}
                                            v1 = v15
                                            _ = v53
                                        }
                                        v18 -= (v18 | v18)
                                    }
                                    v3 += v3
                                    var v54 uint = 65535
                                    _ = v54
                                }
                                var v55 int64 = v3
                                v56 := v2["\n\t"]
                                _ = v55
                                _ = v56
                            }
                            _ = fi23
                        }
                    }
                    v1 = v15
                    _ = fi21
                }
                switch 128 {
                case -2147483648:
                    v57 := GenS5{F0: 42, F1: -957}
                    _ = v57
                case -9223372036854775808:
                    v3 -= v3
                case 32767:
                    fmt.Println(v4, v4)
                }
            }
            v4 += (v4 - v4)
            _ = v18
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v58 bool = ((1.5 / 0.1) < (0.1 / 1e10))
    fmt.Println(v3, v58, v3)
    var v59 float32 = ((float32(3.14159) + float32(float64((-2.25 - 1.5)))) + -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
}
