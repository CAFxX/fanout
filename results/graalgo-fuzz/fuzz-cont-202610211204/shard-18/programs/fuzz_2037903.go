package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 string; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int64; F2 uint32 }

func main() {
    var v1 float64 = 3.14159
    var v2 bool = true
    v3 := []uint32{2147483647, 1, 4294967295}
    var v4 int64 = -1
    var v5 int16 = -32768
    var v6 uint16 = 65535
    var v7 uint64 = 42
    var v8 int8 = -1
    _ = len(v3)
    var v9 int32 = 42
    switch "" {
    case "hello":
        switch 9223372036854775807 {
        case 2147483647:
            _ = len(v3)
        case 32767:
            v9 += 2147483647
        default:
            v8 -= -128
        }
    case "a":
        if true {
            if v2 {
                if (v2 || v2) {
                    switch 1 {
                    case 1:
                        v7 -= 1
                    case -2147483648:
                        v6 -= v6
                    case 4294967295:
                        if v2 {
                            var v10 int32 = int32(v6)
                            var v11 uint64 = (v7 % (((v7 - (v7 & v7))) | 1))
                            switch 414 {
                            case -707:
                                fmt.Println(v7, v2)
                            }
                            _ = v10
                            _ = v11
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    default:
                        switch "foo bar" {
                        case "hello":
                            var v12 bool = true
                            _ = v12
                        case "a":
                            var v13 bool = (42 == 0)
                            _ = v13
                        }
                    }
                    var v14 int32 = (v9 | v9)
                    var v15 uint8 = 146
                    _ = v14
                    _ = v15
                } else {
                    if (42 >= v7) {
                        fmt.Println(v7)
                    }
                }
                switch "foo bar" {
                case "\n\t":
                    v5 = 1
                default:
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v17 := GenS1{F0: (v7 >> ((65535) & 63))}
                        _ = fi16
                        _ = v17
                    }
                }
                v18 := GenS1{F0: (v7 | 1)}
                _ = v18
            } else {
                v5 -= -32768
            }
        } else {
            v6 += (^(v6))
        }
    default:
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 string = "\n\t"
            v8 += v8
            _ = fi19
            _ = v20
        }
    }
    if false {
        if (!(v8 != (^(v8)))) {
            v5 -= int16(v8)
            var v21 bool = ((!true) && v2)
            for fi22 := 0; fi22 < 5; fi22++ {
                v9 = v9
                _ = fi22
            }
            _ = v21
        } else {
            fmt.Println(v2, v7, v6)
            var v23 int16 = (^(v5))
            v24 := GenS1{F0: v7}
            _ = v23
            _ = v24
        }
        switch 1 {
        case 32767:
            if (true && ((1 < 127) || v2)) {
                fmt.Println(v2)
                for fi25 := 0; fi25 < 16; fi25++ {
                    v3 = append(v3, 4294967295)
                    _ = fi25
                }
            } else {
                fmt.Println(v4, v2)
            }
        case 127:
            v7 += v7
        case 1:
            if (0 >= 2) {
                fmt.Println(v8, v5, v2)
                fmt.Println(v9, v2, v8)
            }
        }
        v26 := GenS2{F0: (v5 % (((v5 ^ 1)) | 1)), F1: ("a" + "\n\t"), F2: 255}
        _ = v26
    }
    if false {
        v3 = append(v3, 299)
        v27 := GenS1{F0: 789}
        var v28 float64 = (v1 * 0.1)
        _ = v27
        _ = v28
    }
    fmt.Println(v7, v2, v7)
    var v29 int64 = (v4 % (((v4 - v4)) | 1))
    var v30 uint16 = v6
    v31 := GenS1{F0: 4294967295}
    for fi32 := 0; fi32 < 5; fi32++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if v2 {
            var v33 int16 = v5
            switch fi32 {
            case 42:
                if v2 {
                    var v34 string = "foo bar"
                    var v35 int16 = v5
                    _ = v34
                    _ = v35
                } else {
                    switch (fi32 - (^(fi32))) {
                    case 213:
                        v4 = (v4 * 42)
                    case 0:
                        v8 = v8
                    default:
                        for fi36 := 0; fi36 < 6; fi36++ {
                            v37 := GenS2{F0: (v33 | (v33 & 1)), F1: "hello", F2: 4294967295}
                            v6 = 32767
                            v38 := GenS1{F0: v7}
                            var v39 string = "hello"
                            v5 = (v5 * (v5 * v33))
                            v29 = v29
                            v8 -= 102
                            v40 := GenS2{F0: (^(v5)), F1: v39, F2: 697}
                            if true {
                                v39 = v39
                                _ = len(v3)
                            } else {
                                var v41 int32 = 42
                                v5 = v5
                                var v42 int8 = -128
                                _ = v41
                                _ = v42
                            }
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                    }
                    v1 = v1
                    for fi43 := 0; fi43 < 9; fi43++ {
                        v44 := GenS1{F0: 4294967295}
                        v8 = (v8 ^ -5)
                        var v45 int8 = v8
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                }
            }
            if len(v3) > 0 {
                v3[0] = 857
            }
            _ = v33
        } else {
            for fi46 := 0; fi46 < 3; fi46++ {
                v9 += (v9 ^ -1)
                v3 = append(v3, 2)
                _ = len(v3)
                _ = fi46
            }
            for fi47 := 0; fi47 < 9; fi47++ {
                v8 -= 100
                var v48 uint16 = (v30 ^ v30)
                _ = fi47
                _ = v48
            }
            var v49 int32 = -2147483648
            _ = v49
        }
        _ = fi32
    }
    for fi50 := 0; fi50 < 17; fi50++ {
        v3 = append(v3, 506)
        if (v2 && v2) {
            _ = len(v3)
            var v51 string = "foo bar"
            if len(v3) > 0 {
                v3[0] = 989
            }
            _ = v51
        } else {
            if (fi50 < (fi50 + fi50)) {
                for fi52 := 0; fi52 < 19; fi52++ {
                    if (v2 || true) {
                        if len(v3) > 0 {
                            v3[0] = 2147483648
                        }
                        var v53 int8 = (-(v8))
                        _ = v53
                    } else {
                        v54 := GenS2{F0: (v5 | v5), F1: "a", F2: 205}
                        var v55 bool = v2
                        _ = v54
                        _ = v55
                    }
                    v8 -= (v8 & (^(v8)))
                    var v56 int8 = 1
                    _ = fi52
                    _ = v56
                }
                v30 -= v6
                var v57 int8 = (v8 << ((190) & 63))
                _ = v57
            } else {
                if true {
                    var v58 bool = (!(true && (v2 && v2)))
                    _ = v58
                } else {
                    fmt.Println(v7, v5, v5)
                    for fi59 := 0; fi59 < 5; fi59++ {
                        var v60 uint64 = 42
                        v61 := GenS3{F0: 480}
                        v1 = ((v1 - 1.5) + v1)
                        _ = fi59
                        _ = v60
                        _ = v61
                    }
                }
            }
            v30 += v6
            switch fi50 {
            case -171:
                _ = len(v3)
            case 0:
                var v62 uint64 = v7
                _ = v62
            case 32768:
                switch "foo bar" {
                case "":
                    for fi63 := 0; fi63 < 8; fi63++ {
                        if (v29 <= v4) {
                            v64 := GenS4{F0: (!v2), F1: 42, F2: 544}
                            if v2 {
                                v8 = (v8 & 127)
                            }
                            var v65 float64 = (((v1 * (v1 - float64((0.0 + 1e10)))) * ((1.5 / (1.5 - v1)) + v1)) + v1)
                            _ = v64
                            _ = v65
                        } else {
                            switch "" {
                            case "\n\t":
                                v4 = (v29 << ((18446744073709551615) & 63))
                            default:
                                var v66 int = fi50
                                _ = v66
                            }
                        }
                        v9 = (v9 - 1)
                        fmt.Println(v1, v30)
                        _ = fi63
                    }
                }
            }
        }
        v29 += (-(v4))
        _ = fi50
    }
    v67 := GenS4{F0: v2, F1: (v29 >> ((uint(v30)) & 63)), F2: 2147483647}
    var v68 int64 = (^(v4))
    var v69 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v69
    _ = v7
    _ = v8
    _ = v9
}
