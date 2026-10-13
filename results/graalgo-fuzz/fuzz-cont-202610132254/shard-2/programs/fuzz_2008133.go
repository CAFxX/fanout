package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 uint64; F2 uint }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint64; F1 int32; F2 int64 }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 uint32; F1 int64; F2 int8 }
type GenS9 struct { F0 uint64; F1 string; F2 uint16 }
func genFn0(p0 uint16, p1 uint, p2 int32) string {
    v4 := []int{-1, -1, -9223372036854775808}
    v5 := GenS1{F0: false, F1: 1, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return ""
}

func main() {
    v1 := []int{929, -1}
    var v2 uint32 = 367
    var v3 int64 = 42
    var v6 int32 = 1
    var v7 uint8 = 2
    v6 = v6
    var v8 bool = (true || ((2 >= 4294967295) && false))
    fmt.Println(v3, v8)
    v1 = append(v1, -2147483648)
    var v9 int32 = v6
    switch "a" {
    case "foo bar":
        var v10 uint16 = 65535
        _ = v10
    default:
        if v8 {
            if (127 > 119) {
                for fi11 := 0; fi11 < 13; fi11++ {
                    v12 := GenS1{F0: true, F1: 18446744073709551615, F2: v7}
                    fmt.Println(v8, v7)
                    var v13 string = "hello"
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v1 = append(v1, 128)
            } else {
                v14 := GenS1{F0: (true || v8), F1: 4294967295, F2: 42}
                v1 = append(v1, 428)
                _ = v14
            }
        } else {
            _ = len(v1)
            if len(v1) > 0 {
                v1[0] = 741
            }
            v9 = v6
        }
    }
    v6 = (^(v6))
    var v15 int8 = 92
    v16 := GenS2{F0: v3}
    if v8 {
        var v17 int16 = -519
        v18 := GenS1{F0: v8, F1: 1, F2: 1}
        _ = v17
        _ = v18
    } else {
        if (2 < 0) {
            if (!false) {
                if (v8 && (!false)) {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        var v20 uint64 = 619
                        fmt.Println(v8)
                        var v21 int = -1
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    if len(v1) > 0 {
                        v1[0] = 9223372036854775807
                    }
                    v8 = v8
                } else {
                    switch 9223372036854775807 {
                    case -2147483648:
                        _ = len(v1)
                    case 2147483647:
                        switch "" {
                        case "":
                            if v8 {
                                v22 := GenS1{F0: ("foo bar" >= ("" + "")), F1: 421, F2: (^(v7))}
                                switch 330 {
                                case 127:
                                    if v8 {
                                        switch "foo bar" {
                                        case "foo bar":
                                            fmt.Println(v2, v9, v15)
                                        }
                                        fmt.Println(v8, v2)
                                        v9 = v6
                                    } else {
                                        v15 += (v15 * int8(v3))
                                        var v23 uint8 = 255
                                        v24 := GenS3{F0: 2}
                                        _ = v23
                                        _ = v24
                                    }
                                case 32768:
                                    fmt.Println(v6)
                                }
                                v3 += (-(v3))
                                _ = v22
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                for fi25 := 0; fi25 < 11; fi25++ {
                                    fmt.Println(fi25, fi25)
                                    fmt.Println(v6)
                                    _ = fi25
                                }
                                if (!((-1 < v3) && false)) {
                                    switch "" {
                                    case "hello":
                                        v26 := GenS4{F0: (v9 + (v6 >> ((307) & 63))), F1: 942, F2: 255}
                                        _ = v26
                                    case "foo bar":
                                        if (42 <= 1000) {
                                            fmt.Println(v2)
                                            v6 = v6
                                            v9 -= (v6 + (^(v9)))
                                        } else {
                                            if (!(v8 || v8)) {
                                                v6 = v9
                                                v27 := GenS2{F0: 0}
                                                _ = v27
                                            } else {
                                                if v8 {
                                                    if false {
                                                        for fi28 := 0; fi28 < 10; fi28++ {
                                                            fmt.Println(v9)
                                                            v6 = v6
                                                            _ = fi28
                                                        }
                                                        v29 := GenS5{F0: 2, F1: 42}
                                                        _ = v29
                                                    } else {
                                                        var v30 uint16 = 1
                                                        if v8 {
                                                            switch 4294967295 {
                                                            case -399:
                                                                var v31 uint = 0
                                                                _ = v31
                                                            }
                                                            v3 = 1
                                                            fmt.Println(v8, v15, v15)
                                                        } else {
                                                            v2 += (v2 + 2147483648)
                                                            var v32 uint8 = v7
                                                            v33 := GenS4{F0: 0, F1: 4294967295, F2: 65535}
                                                            _ = v32
                                                            _ = v33
                                                        }
                                                        v8 = v8
                                                        _ = v30
                                                    }
                                                    fmt.Println(v9)
                                                    v34 := GenS6{F0: 0, F1: 1, F2: (^(v3))}
                                                    _ = v34
                                                }
                                                if (v8 && false) {
                                                    var v35 uint = 42
                                                    if (201 < 2) {
                                                        if v8 {
                                                            var v36 uint32 = 2
                                                            var v37 uint64 = 658
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            var v38 bool = (v3 >= v3)
                                                            var v39 uint16 = 884
                                                            if len(v1) > 0 {
                                                                v1[0] = 32767
                                                            }
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        fmt.Println(v6)
                                                        var v40 string = (("\n\t" + "a") + (("foo bar" + "hello") + "a"))
                                                        _ = v40
                                                    } else {
                                                        for fi41 := 0; fi41 < 6; fi41++ {
                                                            if len(v1) > 0 {
                                                                v1[0] = fi41
                                                            }
                                                            v1 = append(v1, fi41)
                                                            _ = fi41
                                                        }
                                                        switch -879 {
                                                        case 692:
                                                            _ = len(v1)
                                                        case 128:
                                                            _ = len(v1)
                                                        case 4294967295:
                                                            v1 = append(v1, -2147483648)
                                                        }
                                                        if false {
                                                            for fi42 := 0; fi42 < 15; fi42++ {
                                                                for fi43 := 0; fi43 < 20; fi43++ {
                                                                    v8 = (v7 <= (v7 | 2))
                                                                    _ = fi43
                                                                }
                                                                v2 = v2
                                                                _ = fi42
                                                            }
                                                            var v44 string = "hello"
                                                            _ = v44
                                                        }
                                                    }
                                                    _ = v35
                                                }
                                            }
                                            switch 9223372036854775807 {
                                            case -2147483648:
                                                _ = len(v1)
                                            default:
                                                var v45 int64 = (^(v3))
                                                _ = v45
                                            }
                                            v7 = v7
                                        }
                                    case "a":
                                        switch 32767 {
                                        case -623:
                                            for fi46 := 0; fi46 < 18; fi46++ {
                                                for fi47 := 0; fi47 < 3; fi47++ {
                                                    _ = len(v1)
                                                    _ = fi47
                                                }
                                                _ = fi46
                                            }
                                        case 128:
                                            v1 = append(v1, -2147483648)
                                        }
                                    default:
                                        v48 := GenS5{F0: 796, F1: -32768}
                                        _ = v48
                                    }
                                }
                            }
                        }
                    }
                    v6 -= (^(v6))
                }
                var v49 float32 = (3.14159 * ((1e10 - -2.25) + 0.1))
                _ = v49
            } else {
                v3 = (v3 & (-(v3)))
            }
        }
        var v50 uint = 4294967295
        var v51 int32 = -2147483648
        _ = v50
        _ = v51
    }
    v52 := GenS7{F0: false, F1: (("a" + "hello") + "hello")}
    for fi53 := 0; fi53 < 9; fi53++ {
        _ = len(v1)
        if v8 {
            switch (fi53 + (fi53 ^ (fi53 | fi53))) {
            case 4294967295:
                for fi54 := 0; fi54 < 13; fi54++ {
                    v55 := GenS8{F0: 2147483648, F1: (-(v3)), F2: (-(v15))}
                    var v56 int16 = -32768
                    v7 -= 0
                    _ = fi54
                    _ = v55
                    _ = v56
                }
            case 1:
                var v57 int8 = (v15 & v15)
                _ = v57
            }
            if (v8 || true) {
                v15 = 13
            }
        } else {
            var v58 uint = 18446744073709551615
            if (65535 == 65535) {
                fmt.Println(v2)
                v2 += (v2 * uint32(v58))
            } else {
                v59 := GenS9{F0: 18446744073709551615, F1: "hello", F2: 1}
                if (v8 || (int16(v2) == 985)) {
                    var v60 uint32 = v2
                    for fi61 := 0; fi61 < 19; fi61++ {
                        if v8 {
                            var v62 uint16 = 42
                            _ = v62
                        } else {
                            switch "" {
                            case "":
                                fmt.Println(v60, v58, v9)
                            case "a":
                                if (((!true) && (v8 || v8)) && (42 <= 0)) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            case "\n\t":
                                _ = len(v1)
                            }
                            var v63 int = 127
                            var v64 int = fi53
                            _ = v63
                            _ = v64
                        }
                        switch (("a" + ("" + "foo bar")) + "") {
                        case "hello":
                            v65 := GenS5{F0: 502, F1: -32768}
                            if len(v1) > 0 {
                                v1[0] = int(v2)
                            }
                            _ = v65
                        }
                        var v66 uint8 = (v7 / (((v7 | v7)) | 1))
                        _ = fi61
                        _ = v66
                    }
                    fmt.Println(v3)
                    _ = v60
                }
                var v67 float32 = float32((1.5 / 1.0))
                _ = v59
                _ = v67
            }
            _ = v58
        }
        _ = fi53
    }
    v3 = v3
    if true {
        fmt.Println(v15)
    } else {
        v15 = v15
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v52
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
