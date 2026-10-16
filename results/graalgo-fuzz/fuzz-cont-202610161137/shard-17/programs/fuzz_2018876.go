package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint; F2 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int, p1 uint32) int64 {
    p1 = (p1 * (p1 << ((685) & 63)))
    p1 = p1
    v6 := []int{42, (p0 ^ (^(p0))), p0, p0}
    p1 = p1
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(p1, fi8, fi8)
            var v9 string = "\n\t"
            _ = fi8
            _ = v9
        }
        fmt.Println(p0)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return -523
}
func genFn1() uint {
    v10 := []int{-948}
    v11 := GenS1{F0: 639}
    _ = v10
    _ = v11
    return 65535
}

func main() {
    var v1 uint = 527
    var v2 uint = (v1 % (((v1 >> ((v1) & 63))) | 1))
    v3 := make(map[string]string)
    v4 := []int{-3}
    v5 := []uint32{}
    fmt.Println(v1)
    var v12 uint16 = 2
    v13 := GenS2{F0: 250}
    v12 = v12
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 int64 = 42
        _ = fi14
        _ = v15
    }
    v1 = genFn1()
    var v16 int8 = -1
    for fi17 := 0; fi17 < 9; fi17++ {
        v18 := GenS3{F0: (v1 << ((4294967295) & 63))}
        if true {
            var v19 uint = v2
            _ = len(v3)
            v1 = v19
            _ = v19
        } else {
            var v20 bool = (true && false)
            switch (fi17 | fi17) {
            case 1:
                v21 := v3[((("foo bar" + "a") + "\n\t") + "hello")]
                _ = v21
            case 0:
                v1 = (v1 >> ((v1) & 63))
            case 491:
                fmt.Println(v20, v16, v1)
            }
            _ = v20
        }
        v3[("a" + "")] = ((((((("hello" + ("" + ("hello" + "a"))) + "a") + ("\n\t" + (((("a" + "foo bar") + "a") + ((("\n\t" + "") + "foo bar") + "\n\t")) + ("a" + ("foo bar" + "hello"))))) + "foo bar") + ("foo bar" + "")) + "") + "hello")
        _ = fi17
        _ = v18
    }
    var v22 float32 = (3.14159 + 3.14159)
    if (-32768 < 32767) {
        if false {
            _ = len(v5)
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint16 = (v12 % (((v12 >> ((genFn1()) & 63))) | 1))
                _ = fi23
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 1; fi25++ {
                fmt.Println(v22, fi25)
                v1 += (v2 * v1)
                _ = fi25
            }
            v16 = 127
        }
    } else {
        v12 += (v12 - v12)
    }
    v1 -= v1
    if ((!(!(v16 < v16))) || false) {
        if ((v16 - -128) >= v16) {
            v26 := GenS1{F0: -9223372036854775808}
            _ = v26
        } else {
            var v27 int32 = -2147483648
            _ = v27
        }
    } else {
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 float32 = ((1.5 * v22) - v22)
            if false {
                fmt.Println(v29, v1)
            } else {
                switch "" {
                case "hello":
                    switch "a" {
                    case "":
                        if true {
                            for fi30 := 0; fi30 < 20; fi30++ {
                                var v31 int64 = 0
                                v16 += (v16 | -128)
                                var v32 int32 = -1
                                v29 = 3.14159
                                v33 := GenS4{F0: -1, F1: (^(v1)), F2: v2}
                                _ = fi30
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            for fi34 := 0; fi34 < 14; fi34++ {
                                v12 -= (^(v12))
                                v5 = append(v5, 42)
                                switch ("hello" + "a") {
                                case "a":
                                    v29 = 1.5
                                default:
                                    v1 -= (^(v1))
                                    if len(v4) > 0 {
                                        v4[0] = fi28
                                    }
                                    v35 := GenS4{F0: v16, F1: 0, F2: 255}
                                    _ = v35
                                }
                                switch fi34 {
                                case 1:
                                    var v36 uint32 = 557
                                    _ = v36
                                case -812:
                                    if true {
                                        if ("\n\t" == (("" + "hello") + "")) {
                                            var v37 uint8 = 40
                                            var v38 int16 = -32768
                                            _ = v37
                                            _ = v38
                                        } else {
                                            var v39 int16 = -32768
                                            switch fi34 {
                                            case -901:
                                                var v40 string = "a"
                                                _ = v40
                                            case 42:
                                                v16 -= 90
                                            default:
                                                v16 = v16
                                            }
                                            v41 := GenS5{F0: 4294967295}
                                            fmt.Println(v1)
                                            _ = v39
                                            _ = v41
                                        }
                                        v42 := GenS2{F0: 127}
                                        _ = v42
                                    } else {
                                        v43 := GenS2{F0: 32}
                                        v2 = (v1 << ((4294967295) & 63))
                                        _ = v43
                                    }
                                case 32767:
                                    v1 -= (v1 * genFn1())
                                default:
                                    v1 += (^(v1))
                                }
                                switch -2147483648 {
                                case -424:
                                    v44 := GenS6{F0: v12, F1: 836}
                                    _ = v44
                                case 2147483647:
                                    v29 = ((0.0 * v22) * -2.25)
                                case 128:
                                    v16 = (v16 ^ v16)
                                }
                                _ = fi34
                            }
                        }
                    case "foo bar":
                        v22 -= 0.0
                    case "a":
                        if false {
                            var v45 int16 = 1000
                            _ = v45
                        } else {
                            fmt.Println(v29, v22, v12)
                            var v46 float64 = ((((-2.25 * 1.5) + 0.0) * (((1e10 + ((0.0 * 1.5) + float64(float32(0.1)))) + -2.25) * 3.14159)) / -2.25)
                            v4 = append(v4, (fi28 % ((int(v12)) | 1)))
                            _ = v46
                        }
                    default:
                        _ = len(v3)
                    }
                default:
                    if true {
                        if ("a" <= "hello") {
                            if false {
                                for fi47 := 0; fi47 < 4; fi47++ {
                                    fmt.Println(v29)
                                    var v48 int16 = 579
                                    _ = fi47
                                    _ = v48
                                }
                                var v49 int16 = -392
                                if (true && (1 < v1)) {
                                    if true {
                                        var v50 int8 = (-(v16))
                                        _ = v50
                                    }
                                    v29 += v22
                                } else {
                                    var v51 bool = false
                                    _ = v51
                                }
                                _ = v49
                            }
                            fmt.Println(v1, v12)
                            var v52 int = 1
                            _ = v52
                        }
                    }
                }
                v53 := GenS4{F0: v16, F1: v2, F2: 255}
                _ = v53
            }
            fmt.Println(v1, v1)
            _ = fi28
            _ = v29
        }
    }
    if ((v1 ^ v2) > v1) {
        _ = len(v4)
    }
    var v54 float32 = float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v54
}
