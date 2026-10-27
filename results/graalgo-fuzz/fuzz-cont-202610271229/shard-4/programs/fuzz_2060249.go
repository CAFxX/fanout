package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 bool; F2 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 uint16; F1 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64; F1 uint16; F2 uint64 }
func genFn0(p0 uint16, p1 uint16, p2 int32) uint8 {
    var v4 int8 = 0
    p0 -= (p0 ^ 0)
    switch -480 {
    case 2147483647:
        if true {
            if true {
                if false {
                    fmt.Println(p0, p1, p1)
                } else {
                    fmt.Println(v4, p2)
                }
            } else {
                var v5 uint16 = 169
                if false {
                    var v6 int32 = p2
                    _ = v6
                } else {
                    if true {
                        for fi7 := 0; fi7 < 20; fi7++ {
                            switch 1 {
                            case 364:
                                v4 -= 100
                            case -2147483648:
                                var v8 int16 = 32767
                                _ = v8
                            case -889:
                                if false {
                                    v5 = v5
                                    v4 -= (v4 << ((4294967295) & 63))
                                }
                            }
                            _ = fi7
                        }
                        if false {
                            fmt.Println(p0, p2)
                            v4 = (-(v4))
                            var v9 int16 = 664
                            _ = v9
                        } else {
                            switch "a" {
                            case "":
                                var v10 int = -9223372036854775808
                                _ = v10
                            default:
                                switch 337 {
                                case 42:
                                    var v11 int64 = -1
                                    _ = v11
                                default:
                                    v12 := GenS1{F0: uint32(p1)}
                                    _ = v12
                                }
                            }
                            var v13 bool = true
                            var v14 int = 1
                            _ = v13
                            _ = v14
                        }
                    } else {
                        v15 := []int{-1}
                        _ = v15
                    }
                    v16 := []int{401}
                    _ = v16
                }
                _ = v5
            }
            switch "\n\t" {
            case "hello":
                for fi17 := 0; fi17 < 11; fi17++ {
                    var v18 uint32 = 2
                    if false {
                        if true {
                            fmt.Println(fi17)
                            var v19 uint64 = 42
                            _ = v19
                        }
                        fmt.Println(p2, p0, p1)
                        if false {
                            p0 -= (p1 + p1)
                        }
                    }
                    _ = fi17
                    _ = v18
                }
            case "foo bar":
                v20 := GenS2{F0: 13}
                _ = v20
            }
        } else {
            var v21 uint32 = 4294967295
            _ = v21
        }
    case 128:
        v4 += (-(v4))
    default:
        v22 := []int{0, 327}
        _ = v22
    }
    fmt.Println(v4, p1)
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 uint64 = 42
        for fi25 := 0; fi25 < 2; fi25++ {
            p1 += p0
            v26 := []int{int(p2)}
            fmt.Println(p1)
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 20; fi27++ {
            p2 = (p2 & (p2 * -256))
            v4 += v4
            var v28 uint32 = 2147483648
            _ = fi27
            _ = v28
        }
        _ = fi23
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}
func genFn1(p0 uint8, p1 string, p2 bool) uint32 {
    switch -1 {
    case -252:
        if p2 {
            v29 := []int{32768, 379, -744, 773}
            p0 = p0
            _ = v29
        }
    case -2147483648:
        v30 := []int{}
        _ = v30
    }
    v31 := GenS1{F0: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    return 229
}

func main() {
    v1 := []int{2147483647}
    v2 := make(map[int]string)
    var v3 string = ("\n\t" + (("\n\t" + "a") + ("\n\t" + ("\n\t" + ""))))
    v3 = ("hello" + v3)
    var v32 uint = 691
    for fi33 := 0; fi33 < 4; fi33++ {
        v3 = ((v3 + v3) + v3)
        for fi34 := 0; fi34 < 16; fi34++ {
            var v35 int8 = 0
            v32 = v32
            switch (fi33 - fi33) {
            case 127:
                if (734 == 0) {
                    if (false || ((fi33 + (fi33 % (((fi33 - 32767)) | 1))) > fi33)) {
                        _ = len(v1)
                        v36 := GenS3{F0: -2147483648, F1: false, F2: 127}
                        var v37 uint = (v32 + v32)
                        _ = v36
                        _ = v37
                    } else {
                        var v38 int = fi34
                        v39 := GenS3{F0: 247, F1: false, F2: genFn0(663, uint16(v32), 758)}
                        v40 := GenS4{F0: (^(v35))}
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                    fmt.Println(fi33)
                }
            default:
                if false {
                    v2[fi33] = v3
                }
            }
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    fmt.Println(v32, v3)
    switch v3 {
    case "foo bar":
        v32 -= 18446744073709551615
    case "":
        v32 -= v32
    case "hello":
        v32 = v32
    }
    v32 = (^(v32))
    fmt.Println(v3)
    for fi41 := 0; fi41 < 7; fi41++ {
        delete(v2, fi41)
        _ = fi41
    }
    if (0 >= -85) {
        if len(v1) > 0 {
            v1[0] = 32768
        }
        fmt.Println(v3, v3)
        switch 1 {
        case 42:
            fmt.Println(v3, v32)
        case 282:
            switch 128 {
            case 59:
                for fi42 := 0; fi42 < 18; fi42++ {
                    fmt.Println(v3, v32)
                    for fi43 := 0; fi43 < 8; fi43++ {
                        if true {
                            v3 = (v3 + "hello")
                            if (42 <= 127) {
                                v44 := GenS5{F0: fi42}
                                v32 = 42
                                if false {
                                    v3 = v3
                                    v45 := GenS6{F0: (fi42 | 32768), F1: -128}
                                    var v46 int = -2147483648
                                    _ = v45
                                    _ = v46
                                } else {
                                    fmt.Println(v3)
                                }
                                _ = v44
                            } else {
                                v47 := GenS6{F0: (fi43 | 0), F1: 42}
                                fmt.Println(fi43)
                                v3 = v3
                                switch (v3 + "foo bar") {
                                case "a":
                                    var v48 uint32 = 286
                                    _ = v48
                                case "\n\t":
                                    v3 = ("" + v3)
                                case "":
                                    if (false && true) {
                                        v49 := GenS7{F0: (v3 + v3), F1: -1}
                                        v32 = 4294967295
                                        if (-2147483648 > -929) {
                                            var v50 int64 = 1
                                            _ = v50
                                        } else {
                                            v51 := GenS2{F0: 128}
                                            var v52 float32 = 0.0
                                            switch ((v3 + v3) + v3) {
                                            case "":
                                                var v53 int16 = 1000
                                                _ = v53
                                            case "\n\t":
                                                fmt.Println(v3, fi43)
                                                if (!false) {
                                                    var v54 uint = v32
                                                    _ = v54
                                                }
                                            default:
                                                fmt.Println(fi42, v32)
                                            }
                                            _ = v51
                                            _ = v52
                                        }
                                        _ = v49
                                    } else {
                                        fmt.Println(fi43)
                                        v55 := GenS8{F0: 1, F1: 652}
                                        _ = v55
                                    }
                                default:
                                    if false {
                                        v56 := GenS5{F0: fi43}
                                        if (!true) {
                                            if (v3 != "a") {
                                                v57 := GenS9{F0: "\n\t"}
                                                v58 := GenS9{F0: "foo bar"}
                                                if true {
                                                    v32 = 255
                                                }
                                                _ = v57
                                                _ = v58
                                            }
                                        }
                                        _ = v56
                                    } else {
                                        v59 := GenS8{F0: 65535, F1: 1}
                                        v32 += (v32 + v32)
                                        _ = v59
                                    }
                                }
                                v60 := GenS3{F0: -646, F1: true, F2: genFn0(32767, 42, 2147483647)}
                                _ = v47
                                _ = v60
                            }
                            v32 += v32
                            if true {
                                fmt.Println(v3, fi43, fi42)
                                _ = len(v1)
                                fmt.Println(v3, fi43)
                            } else {
                                var v61 string = ("hello" + ((v3 + v3) + "\n\t"))
                                if true {
                                    v1 = append(v1, (fi42 & fi43))
                                    var v62 uint64 = 1
                                    _ = v62
                                }
                                v63 := GenS10{F0: 4294967295, F1: 0, F2: 408}
                                _ = v61
                                _ = v63
                            }
                        }
                        v3 = "\n\t"
                        _ = fi43
                    }
                    _ = fi42
                }
            case -759:
                var v64 uint64 = 0
                _ = v64
            }
        case 32767:
            v1 = append(v1, -2147483648)
        }
    }
    v3 = (v3 + "")
    v1 = append(v1, 127)
    fmt.Println(v3, v32)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
}
