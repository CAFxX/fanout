package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint32, p1 int64) uint8 {
    v6 := GenS1{F0: 4294967295}
    fmt.Println(p1, p1)
    var v7 float64 = 0.1
    var v8 uint64 = 18446744073709551615
    v9 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 239
}
func genFn1(p0 int64) uint64 {
    v10 := GenS1{F0: -1}
    var v11 int8 = 70
    var v12 int = 486
    v13 := GenS1{F0: v12}
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int = v12
        v11 -= 83
        var v16 int64 = (-(p0))
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 480
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = (1e10 - (1.5 + 1e10))
    var v3 float64 = -2.25
    v4 := make(map[int]bool)
    var v5 uint8 = 233
    v17 := []int{-9223372036854775808, -2147483648, 2147483647, -659}
    fmt.Println(v2, v2)
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 uint32 = 0
        for fi20 := 0; fi20 < 7; fi20++ {
            switch ("foo bar" + "hello") {
            case "\n\t":
                var v21 int16 = -32768
                _ = v21
            case "foo bar":
                v2 = 1.5
            }
            v3 = v3
            switch -1 {
            case 42:
                v5 = (v5 + v5)
            case 9223372036854775807:
                if true {
                    switch ((("" + "") + "\n\t") + "hello") {
                    case "hello":
                        fmt.Println(fi18)
                    }
                }
            }
            v22 := GenS1{F0: 127}
            _ = fi20
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    v3 -= (v3 / (v3 - (3.14159 - (1e10 * v3))))
    if false {
        var v23 int = -2147483648
        var v24 int = (v23 & 127)
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: 0}
    if (((1.5 + 0.0) - float64(v2)) >= v3) {
        for fi26 := 0; fi26 < 2; fi26++ {
            switch fi26 {
            case 4294967295:
                if true {
                    var v27 uint64 = 0
                    v5 += genFn0(2147483647, -1)
                    _ = len(v17)
                    _ = v27
                }
            case 32768:
                if false {
                    switch fi26 {
                    case 127:
                        fmt.Println(fi26, v2, v2)
                    case 32767:
                        v1["foo bar"] = 724
                    case 2147483647:
                        var v28 uint = 409
                        _ = v28
                    default:
                        var v29 uint8 = (v5 ^ v5)
                        _ = v29
                    }
                } else {
                    var v30 int8 = 42
                    _ = v30
                }
            }
            _ = fi26
        }
        v2 += float32(v3)
    }
    v3 += 1e10
    for fi31 := 0; fi31 < 3; fi31++ {
        for fi32 := 0; fi32 < 18; fi32++ {
            switch fi31 {
            case 32768:
                if (0 <= 0) {
                    v2 = v2
                }
            case 42:
                if (!true) {
                    delete(v1, "")
                    if false {
                        if false {
                            var v33 uint = 65535
                            switch "" {
                            case "":
                                delete(v4, (fi32 >> (((v33 << ((v33) & 63))) & 63)))
                            case "\n\t":
                                switch ("a" + ((((("\n\t" + "hello") + "hello") + "") + (("" + ("\n\t" + "hello")) + "")) + "a")) {
                                case "":
                                    var v34 uint64 = 84
                                    _ = v34
                                case "hello":
                                    var v35 float64 = float64((v2 + float32(((v3 * 1.5) * v3))))
                                    _ = v35
                                case "foo bar":
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(v5, fi32)
                                    }
                                default:
                                    fmt.Println(v33, v3)
                                }
                            default:
                                v36 := GenS1{F0: 712}
                                fmt.Println(v2, v2)
                                v33 -= v33
                                _ = v36
                            }
                            switch ("\n\t" + "") {
                            case "hello":
                                fmt.Println(v2, v33, fi31)
                            case "foo bar":
                                var v37 int32 = -855
                                _ = v37
                            }
                            _ = len(v1)
                            switch fi31 {
                            case 32768:
                                var v38 int32 = 0
                                _ = v38
                            }
                            _ = v33
                        } else {
                            v5 = 168
                            switch (("hello" + "\n\t") + ("a" + "foo bar")) {
                            case "":
                                v2 += -2.25
                            case "foo bar":
                                delete(v4, fi32)
                            default:
                                var v39 uint16 = 42
                                _ = v39
                            }
                        }
                        _ = len(v1)
                    }
                } else {
                    switch fi31 {
                    case 32767:
                        fmt.Println(v3, fi31, fi32)
                    }
                }
            case -993:
                if true {
                    fmt.Println(v5)
                }
                switch (fi32 % (((fi31 / (((fi32 << ((0) & 63))) | 1))) | 1)) {
                case 2147483647:
                    v3 = v3
                default:
                    v17 = append(v17, (^(fi31)))
                    var v40 float32 = v2
                    _ = v40
                }
                delete(v1, "")
            default:
                v2 = v2
            }
            v5 = 255
            _ = fi32
        }
        for fi41 := 0; fi41 < 14; fi41++ {
            var v42 int16 = 1
            _ = fi41
            _ = v42
        }
        if true {
            var v43 uint16 = 42
            _ = v43
        } else {
            var v44 uint32 = 967
            _ = v44
        }
        _ = fi31
    }
    if ("\n\t" >= "hello") {
        switch "a" {
        case "":
            var v45 float32 = v2
            _ = v45
        case "hello":
            v2 -= ((0.0 * 1.5) / ((v2 * (1e10 + ((0.1 / (0.1 * v2)) * 0.1))) + ((((v2 / (v2 - 1.5)) * v2) - v2) / float32(v3))))
        default:
            if true {
                var v46 int32 = -2147483648
                v17 = append(v17, 891)
                _ = v46
            } else {
                for fi47 := 0; fi47 < 16; fi47++ {
                    if (-721 > 407) {
                        switch "" {
                        case "a":
                            v48 := GenS1{F0: (fi47 << ((902) & 63))}
                            _ = v48
                        case "foo bar":
                            var v49 uint32 = 513
                            _ = v49
                        default:
                            v50 := GenS1{F0: (fi47 ^ 432)}
                            _ = v50
                        }
                        delete(v4, fi47)
                        var v51 int8 = -1
                        _ = v51
                    } else {
                        switch 338 {
                        case 9223372036854775807:
                            if true {
                                v2 -= v2
                                v52 := GenS1{F0: 9223372036854775807}
                                switch "\n\t" {
                                case "foo bar":
                                    fmt.Println(v3, v2)
                                }
                                _ = v52
                            }
                        case 2147483647:
                            if false {
                                _ = len(v17)
                                v53 := GenS1{F0: 42}
                                _ = v53
                            } else {
                                var v54 float32 = 1.5
                                if true {
                                    var v55 int32 = -36
                                    var v56 int16 = 1000
                                    v17 = append(v17, fi47)
                                    _ = v55
                                    _ = v56
                                } else {
                                    var v57 uint = 42
                                    _ = v57
                                }
                                _ = v54
                            }
                        default:
                            if false {
                                fmt.Println(v2, v5)
                                v3 = v3
                            } else {
                                var v58 bool = (!false)
                                _ = v58
                            }
                        }
                        fmt.Println(v5)
                        _ = len(v17)
                    }
                    v59 := GenS1{F0: (fi47 ^ 4294967295)}
                    _ = fi47
                    _ = v59
                }
                switch 2147483647 {
                case 32767:
                    var v60 int8 = int8(v5)
                    _ = v60
                case -17:
                    fmt.Println(v2, v3, v2)
                case 922:
                    if (2147483647 > 4294967295) {
                        v61 := v1[("\n\t" + (("hello" + "foo bar") + ""))]
                        _ = v61
                    }
                }
            }
        }
        var v62 int8 = -93
        _ = v62
    } else {
        switch 1 {
        case 9223372036854775807:
            switch "" {
            case "foo bar":
                var v63 int16 = 32767
                _ = v63
            default:
                if true {
                    v2 += 0.0
                    _ = len(v17)
                }
            }
        case 0:
            var v64 bool = (2 >= 42)
            _ = v64
        }
        v17 = append(v17, 777)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
