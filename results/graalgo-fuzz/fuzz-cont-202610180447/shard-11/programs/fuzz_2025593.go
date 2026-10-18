package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint8; F2 uint }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 uint16; F2 int }
type GenS8 struct { F0 uint16; F1 bool }
type GenS9 struct { F0 int16; F1 int16; F2 uint32 }

func main() {
    var v1 uint16 = 42
    v2 := []uint32{1}
    v3 := make(map[string]bool)
    var v4 int64 = 9223372036854775807
    var v5 uint64 = 322
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 int8 = 0
        if (!(1000 <= 32767)) {
            v3["hello"] = (!true)
            v5 = v5
            v8 := GenS1{F0: 822, F1: v5, F2: (v5 << ((2) & 63))}
            _ = v8
        }
        var v9 uint64 = (^(v5))
        _ = fi6
        _ = v7
        _ = v9
    }
    v2 = append(v2, 4)
    v10 := GenS2{F0: 4294967295}
    var v11 int16 = 231
    var v12 int8 = 127
    v4 = (v4 - int64(v1))
    v13 := GenS1{F0: 42, F1: 4294967295, F2: v5}
    switch 32767 {
    case -2147483648:
        if (false || true) {
            v14 := GenS3{F0: v4, F1: 0, F2: 42}
            v1 = (v1 ^ v1)
            _ = len(v3)
            _ = v14
        }
    default:
        switch "a" {
        case "a":
            v12 -= -68
        case "foo bar":
            if (1 < 2147483647) {
                v15 := v3["foo bar"]
                v11 = v11
                _ = v15
            } else {
                fmt.Println(v5, v4, v5)
                for fi16 := 0; fi16 < 2; fi16++ {
                    if (((1 != (v5 | v5)) && false) || false) {
                        fmt.Println(fi16, v1)
                    }
                    _ = fi16
                }
                fmt.Println(v4, v1)
            }
        case "":
            for fi17 := 0; fi17 < 9; fi17++ {
                v12 = v12
                v12 += 127
                _ = fi17
            }
        default:
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS4{F0: (fi18 + fi18), F1: -1}
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 int = 257
                    v22 := GenS5{F0: v5, F1: 574}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = fi18
                _ = v19
            }
        }
    }
    if false {
        _ = len(v2)
        if true {
            for fi23 := 0; fi23 < 4; fi23++ {
                if ((v11 | v11) >= v11) {
                    var v24 uint16 = v1
                    for fi25 := 0; fi25 < 6; fi25++ {
                        var v26 int = 128
                        var v27 int32 = 1
                        v28 := GenS6{F0: 65535}
                        _ = fi25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    v29 := GenS7{F0: 0, F1: (v1 >> ((4294967295) & 63)), F2: (fi23 & -285)}
                    _ = v24
                    _ = v29
                } else {
                    var v30 int16 = (^(v11))
                    if false {
                        v11 = (v30 * (-(v11)))
                    }
                    _ = v30
                }
                _ = fi23
            }
        }
        switch 42 {
        case 128:
            switch "\n\t" {
            case "a":
                var v31 int64 = v4
                _ = v31
            default:
                for fi32 := 0; fi32 < 10; fi32++ {
                    switch (-(fi32)) {
                    case -67:
                        if ((fi32 & (fi32 ^ int(v12))) > -307) {
                            for fi33 := 0; fi33 < 4; fi33++ {
                                v12 += (v12 / (((v12 << ((47) & 63))) | 1))
                                v4 = 0
                                _ = fi33
                            }
                            v34 := GenS8{F0: 42, F1: false}
                            _ = v34
                        } else {
                            var v35 string = "\n\t"
                            _ = v35
                        }
                    default:
                        if ((v11 != (v11 + v11)) && (true || false)) {
                            switch int(v1) {
                            case -9223372036854775808:
                                v36 := GenS1{F0: v11, F1: v5, F2: v5}
                                _ = v36
                            case 0:
                                var v37 uint64 = (v5 & (v5 ^ v5))
                                _ = v37
                            default:
                                var v38 bool = false
                                _ = v38
                            }
                        } else {
                            v39 := GenS8{F0: 32767, F1: false}
                            var v40 uint16 = (v1 >> ((856) & 63))
                            _ = v39
                            _ = v40
                        }
                    }
                    for fi41 := 0; fi41 < 20; fi41++ {
                        var v42 uint = 6
                        _ = fi41
                        _ = v42
                    }
                    v5 = (v5 >> ((4294967295) & 63))
                    _ = fi32
                }
            }
        case 32768:
            if true {
                if false {
                    switch "hello" {
                    case "hello":
                        if true {
                            v12 = (v12 & v12)
                            for fi43 := 0; fi43 < 20; fi43++ {
                                var v44 string = ""
                                v45 := GenS8{F0: v1, F1: false}
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v4 -= 0
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                    var v46 int16 = (v11 ^ (v11 ^ (v11 | (^(v11)))))
                    _ = v46
                }
                if (true || false) {
                    var v47 uint = 895
                    var v48 uint8 = 42
                    _ = v47
                    _ = v48
                }
                var v49 int8 = v12
                _ = v49
            }
        default:
            if false {
                if (!((-(v11)) > v11)) {
                    var v50 float32 = ((3.14159 * float32(0.0)) + (0.1 / 1.0))
                    _ = v50
                } else {
                    var v51 int = -587
                    switch (-(v51)) {
                    case 1:
                        v1 = (v1 + v1)
                    case -899:
                        v5 = v5
                    default:
                        var v52 uint64 = (v5 + 42)
                        _ = v52
                    }
                    _ = v51
                }
                if (0 == (v11 * (v11 << ((42) & 63)))) {
                    for fi53 := 0; fi53 < 8; fi53++ {
                        v2 = append(v2, 298)
                        _ = fi53
                    }
                    for fi54 := 0; fi54 < 19; fi54++ {
                        var v55 uint = 0
                        v56 := GenS5{F0: (v5 / ((v5) | 1)), F1: int(v11)}
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                    v57 := GenS6{F0: 379}
                    _ = v57
                } else {
                    var v58 uint = 451
                    var v59 uint = v58
                    _ = v58
                    _ = v59
                }
                v60 := GenS3{F0: (v4 * (v4 | 1)), F1: 200, F2: 0}
                _ = v60
            } else {
                v61 := GenS9{F0: -32768, F1: v11, F2: 295}
                v11 += v11
                _ = len(v2)
                _ = v61
            }
        }
    } else {
        fmt.Println(v1, v11)
    }
    if false {
        var v62 int64 = v4
        _ = v62
    } else {
        fmt.Println(v12, v1)
        if true {
            for fi63 := 0; fi63 < 13; fi63++ {
                var v64 float64 = -2.25
                v1 += v1
                _ = fi63
                _ = v64
            }
            v11 -= -73
        }
        fmt.Println(v1, v11)
    }
    v4 += (-(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
