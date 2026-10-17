package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint32, p1 bool) bool {
    var v7 string = "foo bar"
    var v8 bool = p1
    fmt.Println(p1, v8)
    fmt.Println(p1, v8, v7)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (407 >= -76)
}
func genFn1(p0 uint16) uint {
    fmt.Println(p0)
    p0 = p0
    fmt.Println(p0)
    var v9 uint16 = (p0 & p0)
    _ = p0
    _ = v9
    return 561
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 4294967295
    var v3 float32 = 3.14159
    var v4 bool = (587 <= 1)
    v5 := make(map[uint32]uint32)
    v6 := []int{128, 32768}
    if v4 {
        if len(v6) > 0 {
            v6[0] = 127
        }
        v10 := GenS1{F0: 0, F1: 444}
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int16 = 32767
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    v5[2147483648] = 65535
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch "a" {
    case "":
        switch "\n\t" {
        case "a":
            var v13 int8 = 47
            _ = v13
        case "foo bar":
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS1{F0: 42, F1: 0}
                _ = fi14
                _ = v15
            }
        }
    case "hello":
        switch -2147483648 {
        case -412:
            switch ("" + "") {
            case "a":
                v2 += genFn1(2)
            case "":
                var v16 uint = 720
                _ = v16
            default:
                fmt.Println(v3)
            }
        case 128:
            for fi17 := 0; fi17 < 4; fi17++ {
                for fi18 := 0; fi18 < 8; fi18++ {
                    if v4 {
                        switch ("a" + ("hello" + "")) {
                        case "a":
                            if v4 {
                                var v19 int32 = 391
                                if (!false) {
                                    var v20 uint8 = 128
                                    v21 := GenS1{F0: 1000, F1: 784}
                                    if false {
                                        switch ("foo bar" + "") {
                                        case "\n\t":
                                            if v4 {
                                                v22 := GenS2{F0: 0}
                                                _ = v22
                                            }
                                        case "":
                                            fmt.Println(v3, v19, fi17)
                                        }
                                    } else {
                                        v3 -= 1e10
                                    }
                                    _ = v20
                                    _ = v21
                                } else {
                                    v19 = (v19 << ((65535) & 63))
                                    v23 := GenS1{F0: 1000, F1: 338}
                                    v1[fi17] = ""
                                    v3 = 1.5
                                    _ = v23
                                }
                                switch fi18 {
                                case -2147483648:
                                    fmt.Println(v19, v2, v2)
                                case 1:
                                    _ = len(v6)
                                case 128:
                                    if true {
                                        v24 := v5[2]
                                        var v25 string = "foo bar"
                                        if false {
                                            v2 -= v2
                                        }
                                        v26 := GenS1{F0: -527, F1: uint16(v24)}
                                        v3 += (1e10 + ((0.1 / v3) * 3.14159))
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    } else {
                                        _ = len(v6)
                                        v2 -= (^(v2))
                                    }
                                }
                                _ = v19
                            } else {
                                v2 -= (v2 << (((^(v2))) & 63))
                                switch "foo bar" {
                                case "a":
                                    v27 := GenS3{F0: (^(v2)), F1: 714}
                                    _ = v27
                                case "":
                                    v28 := GenS4{F0: 1}
                                    _ = v28
                                default:
                                    if (v4 || false) {
                                        v2 -= (^(v2))
                                        fmt.Println(fi18)
                                        var v29 uint16 = 255
                                        _ = v29
                                    }
                                }
                                switch 32768 {
                                case 861:
                                    if (!v4) {
                                        switch ("foo bar" + "a") {
                                        case "a":
                                            var v30 uint8 = 1
                                            _ = v30
                                        }
                                        v4 = (0.1 != 1.5)
                                    } else {
                                        if false {
                                            var v31 uint16 = 65535
                                            _ = v31
                                        }
                                        v2 -= v2
                                        v32 := GenS3{F0: 4294967295, F1: 1}
                                        _ = v32
                                    }
                                case 127:
                                    delete(v1, (fi17 | int(v2)))
                                default:
                                    v33 := v5[4294967295]
                                    _ = v33
                                }
                            }
                        }
                        fmt.Println(v4, v3, v3)
                    } else {
                        if v4 {
                            v34 := GenS1{F0: -1, F1: 345}
                            var v35 int64 = 0
                            _ = v34
                            _ = v35
                        }
                        if false {
                            v1[(fi17 % ((-747) | 1))] = "foo bar"
                            fmt.Println(fi18)
                            var v36 uint = v2
                            _ = v36
                        } else {
                            if (!true) {
                                fmt.Println(v3)
                                v3 = (((v3 / v3) * v3) + 0.1)
                                var v37 uint32 = 915
                                _ = v37
                            }
                            fmt.Println(v3, fi18, v2)
                        }
                        if (v4 || false) {
                            fmt.Println(fi18, fi17, v3)
                            fmt.Println(v4)
                            v38 := GenS1{F0: -32768, F1: 65535}
                            _ = v38
                        }
                        fmt.Println(v3, v3, fi17)
                        fmt.Println(v4)
                    }
                    switch (fi18 ^ (fi18 - -116)) {
                    case -1:
                        var v39 uint = (v2 << (((v2 << (((v2 >> ((genFn1(1)) & 63))) & 63))) & 63))
                        _ = v39
                    case -51:
                        v2 += v2
                    default:
                        v40 := GenS1{F0: 0, F1: 32767}
                        _ = v40
                    }
                    var v41 int32 = -2147483648
                    _ = fi18
                    _ = v41
                }
                for fi42 := 0; fi42 < 13; fi42++ {
                    if true {
                        v6 = append(v6, -453)
                    }
                    _ = fi42
                }
                _ = fi17
            }
        default:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
    }
    var v43 uint16 = 32767
    if v4 {
        switch -186 {
        case 128:
            v6 = append(v6, 32768)
        default:
            for fi44 := 0; fi44 < 3; fi44++ {
                var v45 int = fi44
                _ = fi44
                _ = v45
            }
        }
        delete(v1, 32768)
        var v46 bool = ((!((v2 >> ((v2) & 63)) >= (v2 | v2))) && (42 <= 2))
        _ = v46
    }
    v6 = append(v6, 128)
    v47 := GenS3{F0: (^(v2)), F1: 2}
    fmt.Println(v3, v3, v43)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v47
    _ = v5
    _ = v6
}
