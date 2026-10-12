package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint; F1 int16; F2 uint }
type GenS3 struct { F0 int8; F1 uint; F2 uint64 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int; F1 uint; F2 int32 }
type GenS6 struct { F0 bool; F1 int8 }
func genFn0() int16 {
    v5 := GenS1{F0: 4294967295, F1: 1, F2: 557}
    v6 := GenS1{F0: 0, F1: 1, F2: 1000}
    var v7 uint16 = 1
    v8 := GenS1{F0: 2147483647, F1: -1, F2: 262}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 12
}
func genFn1(p0 uint32, p1 uint8) int16 {
    var v9 float32 = ((3.14159 * 3.14159) + (0.0 + (0.1 * ((0.0 / 0.1) * 1e10))))
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            fmt.Println(fi10)
            _ = fi11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            p0 = (p0 + p0)
            if false {
                var v13 uint = 1
                _ = v13
            } else {
                fmt.Println(p0)
            }
            _ = fi12
        }
        var v14 int64 = -1
        _ = fi10
        _ = v14
    }
    var v15 uint32 = (p0 + (^(p0)))
    fmt.Println(v15)
    v9 = v9
    _ = p0
    _ = p1
    _ = v15
    _ = v9
    return genFn0()
}

func main() {
    var v1 int8 = 42
    var v2 uint = 65535
    var v3 uint = 781
    v4 := make(map[string]bool)
    v2 += (v3 + (v2 / (((v3 >> ((v2) & 63))) | 1)))
    v16 := GenS2{F0: v2, F1: 42, F2: v2}
    fmt.Println(v3)
    v1 -= v1
    var v17 int64 = 628
    fmt.Println(v3, v2, v17)
    if (0 != 1000) {
        var v18 string = "foo bar"
        _ = v18
    } else {
        for fi19 := 0; fi19 < 13; fi19++ {
            for fi20 := 0; fi20 < 19; fi20++ {
                switch "a" {
                case "a":
                    v2 += v3
                default:
                    var v21 string = "foo bar"
                    _ = v21
                }
                _ = fi20
            }
            for fi22 := 0; fi22 < 4; fi22++ {
                fmt.Println(fi19, v2, v17)
                v23 := v4["a"]
                var v24 int16 = -32768
                switch (fi19 - int(v2)) {
                case 0:
                    switch (fi19 - -1) {
                    case 0:
                        fmt.Println(fi19)
                    case 404:
                        fmt.Println(v1, v2, v24)
                    case -197:
                        v24 -= 508
                    default:
                        var v25 int16 = v24
                        _ = v25
                    }
                default:
                    v24 = v24
                }
                if v23 {
                    fmt.Println(v1, v3, v23)
                    v17 -= int64(fi22)
                } else {
                    fmt.Println(fi22, v1, v24)
                }
                if v23 {
                    if (fi22 != (-(fi22))) {
                        var v26 int = (fi22 / (((fi19 % (((fi22 + fi22)) | 1))) | 1))
                        _ = v26
                    } else {
                        fmt.Println(v23, v17, fi22)
                    }
                    v27 := GenS2{F0: (v3 << ((65535) & 63)), F1: -1, F2: (v2 - v2)}
                    if v23 {
                        delete(v4, "foo bar")
                    }
                    _ = v27
                } else {
                    if (0.1 <= 0.1) {
                        if (42 >= 4294967295) {
                            var v28 uint64 = 960
                            var v29 int = fi22
                            if true {
                                var v30 int8 = 1
                                _ = v30
                            } else {
                                v2 += 18446744073709551615
                                v31 := []int{}
                                if v23 {
                                    switch ("hello" + ("\n\t" + "hello")) {
                                    case "":
                                        if v23 {
                                            if (!v23) {
                                                if (((1 != 128) || v23) && (!v23)) {
                                                    if len(v31) >= 2 {
                                                        v31 = v31[1:len(v31)]
                                                    }
                                                } else {
                                                    switch int(v1) {
                                                    case 750:
                                                        var v32 bool = false
                                                        _ = v32
                                                    }
                                                    var v33 int64 = (v17 + v17)
                                                    _ = v33
                                                }
                                                v1 -= v1
                                            } else {
                                                switch (v29 & (^(fi22))) {
                                                case 974:
                                                    if (!true) {
                                                        var v34 uint8 = 199
                                                        v17 -= (v17 >> (((v2 - v3)) & 63))
                                                        v24 = genFn0()
                                                        if v23 {
                                                            v35 := GenS3{F0: v1, F1: (^(v2)), F2: (v28 & 2)}
                                                            var v36 int32 = -2147483648
                                                            _ = v35
                                                            _ = v36
                                                        } else {
                                                            fmt.Println(fi22, v17, v2)
                                                            fmt.Println(v29, v1, v23)
                                                            delete(v4, "a")
                                                        }
                                                        _ = v34
                                                    }
                                                case -815:
                                                    _ = len(v4)
                                                }
                                                v29 += fi19
                                            }
                                            v37 := GenS3{F0: v1, F1: v3, F2: v28}
                                            var v38 int64 = -288
                                            _ = v37
                                            _ = v38
                                        }
                                    case "foo bar":
                                        if false {
                                            if (262 > 2) {
                                                var v39 string = ("foo bar" + "\n\t")
                                                var v40 int8 = v1
                                                _ = v39
                                                _ = v40
                                            } else {
                                                fmt.Println(fi22)
                                            }
                                            _ = len(v31)
                                            v2 = (v3 - v2)
                                        }
                                    case "a":
                                        if (v1 == -128) {
                                            v41 := GenS4{F0: 2147483647, F1: 42}
                                            v24 = int16(v2)
                                            _ = v41
                                        }
                                    default:
                                        if len(v31) >= 2 {
                                            v31 = v31[1:len(v31)]
                                        }
                                    }
                                } else {
                                    v42 := GenS2{F0: v3, F1: genFn0(), F2: (v3 / ((v3) | 1))}
                                    _ = len(v4)
                                    _ = len(v31)
                                    _ = v42
                                }
                                _ = v31
                            }
                            v43 := GenS5{F0: (-(v29)), F1: (v3 * (v3 - v3)), F2: 42}
                            _ = v28
                            _ = v29
                            _ = v43
                        }
                    }
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            switch ("hello" + "") {
            case "\n\t":
                switch (((("foo bar" + (("foo bar" + "hello") + "")) + "a") + "\n\t") + "") {
                case "\n\t":
                    v44 := GenS6{F0: false, F1: 100}
                    _ = v44
                case "":
                    v2 = v2
                }
            case "":
                v17 -= v17
            case "hello":
                if true {
                    v45 := GenS3{F0: v1, F1: (^(v3)), F2: 369}
                    fmt.Println(v1, v2, v3)
                    _ = v45
                } else {
                    v1 += (v1 - (v1 | (v1 ^ (-(v1)))))
                    var v46 uint64 = 18446744073709551615
                    v17 += v17
                    _ = v46
                }
            default:
                fmt.Println(fi19, fi19, v17)
            }
            _ = fi19
        }
        switch ("foo bar" + "") {
        case "":
            var v47 float64 = 3.14159
            _ = v47
        case "a":
            v48 := GenS5{F0: -195, F1: (v2 ^ v2), F2: -427}
            _ = v48
        case "\n\t":
            var v49 int8 = v1
            _ = v49
        default:
            switch "" {
            case "a":
                for fi50 := 0; fi50 < 2; fi50++ {
                    v51 := GenS3{F0: v1, F1: (v2 & 18446744073709551615), F2: 2}
                    v52 := GenS4{F0: (fi50 << ((v2) & 63)), F1: v1}
                    _ = fi50
                    _ = v51
                    _ = v52
                }
            default:
                delete(v4, "a")
            }
        }
        v53 := GenS6{F0: false, F1: (v1 - (v1 | 42))}
        _ = v53
    }
    v4[("a" + ("a" + "hello"))] = (434 == 2)
    var v54 uint = 4294967295
    if (false && ((false || (false || (255 >= 32767))) || true)) {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v17)
        }
        var v55 bool = (true && true)
        _ = v55
    }
    var v56 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v56
}
