package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 float32 = 1e10
    var v2 uint16 = 0
    var v3 uint64 = 0
    v3 = v3
    v1 = 1.5
    v3 -= 42
    v4 := GenS1{F0: -1}
    switch 127 {
    case 32767:
        if (!(65535 >= 30)) {
            var v5 bool = true
            var v6 int8 = -59
            _ = v5
            _ = v6
        }
    case 32768:
        switch "\n\t" {
        case "a":
            var v7 uint32 = 2147483647
            _ = v7
        default:
            v8 := GenS1{F0: 0}
            _ = v8
        }
    case -2147483648:
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 int = 127
            _ = fi9
            _ = v10
        }
    default:
        if false {
            var v11 uint32 = 3
            _ = v11
        } else {
            var v12 float32 = v1
            var v13 uint16 = 792
            _ = v12
            _ = v13
        }
    }
    if false {
        switch ("a" + ("foo bar" + "\n\t")) {
        case "foo bar":
            if true {
                for fi14 := 0; fi14 < 6; fi14++ {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        v16 := []int{(fi15 & (-(fi15))), 2147483647, fi14, fi14}
                        _ = fi15
                        _ = v16
                    }
                    switch -273 {
                    case 9223372036854775807:
                        if (2 < uint(v3)) {
                            v3 = 1
                        }
                    case 167:
                        v17 := GenS1{F0: (fi14 / (((fi14 >> ((0) & 63))) | 1))}
                        _ = v17
                    case 2147483647:
                        fmt.Println(v3)
                    default:
                        var v18 bool = true
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 15; fi19++ {
                        v3 -= (^(v3))
                        if true {
                            var v20 bool = (true && ((v1 == v1) && true))
                            _ = v20
                        } else {
                            v21 := []int{(fi14 & (fi14 % (((fi19 * (fi19 | fi19))) | 1))), fi19, fi14}
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            fmt.Println(fi19)
                            _ = v21
                        }
                        if false {
                            v2 = 65535
                            v22 := []int{fi19}
                            fmt.Println(v2, fi14, v3)
                            _ = v22
                        } else {
                            v3 = 1
                        }
                        v23 := GenS2{F0: v2, F1: 127}
                        v24 := GenS2{F0: (v2 * 2), F1: -82}
                        v25 := []int{fi14}
                        fmt.Println(v1, v3)
                        _ = fi19
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    _ = fi14
                }
                for fi26 := 0; fi26 < 12; fi26++ {
                    if (-1 >= 1) {
                        v27 := GenS3{F0: 55, F1: -948}
                        v1 += 1e10
                        var v28 int64 = 9223372036854775807
                        _ = v27
                        _ = v28
                    }
                    _ = fi26
                }
                v29 := GenS4{F0: -32768, F1: 0}
                _ = v29
            }
        case "":
            var v30 string = "a"
            _ = v30
        }
    }
    if false {
        if false {
            v31 := []int{-696}
            v1 = ((1.5 * 0.1) - float32(3.14159))
            if len(v31) > 0 {
                v31[0] = 32768
            }
            _ = v31
        } else {
            v3 = (v3 % ((4294967295) | 1))
            var v32 int16 = -1
            _ = v32
        }
        if false {
            var v33 int32 = -1
            var v34 float32 = v1
            _ = v33
            _ = v34
        } else {
            if false {
                var v35 uint32 = 42
                for fi36 := 0; fi36 < 6; fi36++ {
                    var v37 int8 = -43
                    _ = fi36
                    _ = v37
                }
                if true {
                    v38 := GenS2{F0: v2, F1: 100}
                    _ = v38
                }
                _ = v35
            }
            v2 -= v2
            var v39 uint64 = v3
            _ = v39
        }
    }
    v3 = (v3 >> ((32) & 63))
    v40 := []int{-786, 9223372036854775807, 32768, 0}
    var v41 int16 = -1
    fmt.Println(v41)
    var v42 uint32 = 199
    v41 -= (v41 / ((v41) | 1))
    for fi43 := 0; fi43 < 7; fi43++ {
        if (32767 >= 32767) {
            _ = len(v40)
        } else {
            if true {
                var v44 bool = true
                if (v1 >= 0.0) {
                    v41 = (^(v41))
                    v3 += (v3 | v3)
                }
                if (v41 != (v41 | v41)) {
                    fmt.Println(v1, v1, v41)
                } else {
                    if true {
                        for fi45 := 0; fi45 < 8; fi45++ {
                            var v46 bool = false
                            if v44 {
                                var v47 string = "foo bar"
                                v48 := GenS5{F0: -1}
                                v49 := GenS1{F0: int(v41)}
                                v50 := GenS6{F0: 42, F1: int(v42)}
                                _ = v47
                                _ = v48
                                _ = v49
                                _ = v50
                            } else {
                                v3 = (v3 / ((v3) | 1))
                                v44 = (!(!(v41 < (v41 << ((65535) & 63)))))
                                var v51 int8 = 24
                                _ = v51
                            }
                            v1 -= v1
                            _ = fi45
                            _ = v46
                        }
                        if len(v40) > 0 {
                            v40[0] = (^(fi43))
                        }
                    } else {
                        if false {
                            var v52 int = fi43
                            v53 := GenS6{F0: 42, F1: (fi43 % ((v52) | 1))}
                            _ = v52
                            _ = v53
                        }
                    }
                    if (v44 || false) {
                        switch int(v3) {
                        case 42:
                            v2 += v2
                        default:
                            for fi54 := 0; fi54 < 5; fi54++ {
                                v55 := GenS7{F0: (^(v42)), F1: 80}
                                var v56 float32 = 1e10
                                _ = fi54
                                _ = v55
                                _ = v56
                            }
                        }
                        for fi57 := 0; fi57 < 3; fi57++ {
                            var v58 uint = uint(v3)
                            _ = fi57
                            _ = v58
                        }
                        if len(v40) >= 2 {
                            v40 = v40[1:len(v40)]
                        }
                    } else {
                        fmt.Println(v44)
                        if len(v40) >= 2 {
                            v40 = v40[1:len(v40)]
                        }
                        v59 := GenS4{F0: (v41 - v41), F1: -9223372036854775808}
                        _ = v59
                    }
                }
                _ = v44
            } else {
                switch ("hello" + "hello") {
                case "hello":
                    var v60 int64 = 42
                    _ = v60
                case "\n\t":
                    var v61 int = fi43
                    _ = v61
                case "foo bar":
                    fmt.Println(v41, v42, v42)
                }
            }
            if len(v40) > 0 {
                v40[0] = 42
            }
            var v62 int32 = 2147483647
            _ = v62
        }
        _ = fi43
    }
    for fi63 := 0; fi63 < 19; fi63++ {
        for fi64 := 0; fi64 < 14; fi64++ {
            var v65 int64 = 425
            v66 := GenS6{F0: 20, F1: (fi63 - 32767)}
            _ = fi64
            _ = v65
            _ = v66
        }
        v1 -= float32((3.14159 * 1e10))
        _ = fi63
    }
    var v67 float64 = 3.14159
    v68 := GenS1{F0: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v67
    _ = v68
}
