package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int32; F2 int16 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint32 }
func genFn0(p0 string, p1 int16, p2 int) uint16 {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int = (fi6 << ((65535) & 63))
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint32 = 0
            switch ((p0 + p0) + p0) {
            case "hello":
                v10 := []int{2147483647, (fi8 % (((v7 + v7)) | 1)), (v7 % ((32768) | 1))}
                _ = v10
            case "a":
                v7 += fi8
            case "foo bar":
                if (true || (1e10 == 3.14159)) {
                    if false {
                        v11 := []int{-266, 0, 0, (fi6 >> ((42) & 63))}
                        if false {
                            switch 32767 {
                            case -2147483648:
                                v12 := GenS1{F0: 255, F1: -128}
                                _ = v12
                            case 73:
                                var v13 uint = 4294967295
                                switch p0 {
                                case "\n\t":
                                    p1 = 1
                                case "a":
                                    if ((0.1 + 0.0) == (float32((3.14159 + 0.1)) * 1.5)) {
                                        var v14 uint = uint(v9)
                                        p2 = (fi8 + (fi8 / ((v7) | 1)))
                                        p2 = p2
                                        v15 := GenS1{F0: 1, F1: 1}
                                        _ = v14
                                        _ = v15
                                    }
                                default:
                                    var v16 float64 = 0.0
                                    _ = v16
                                }
                                _ = v13
                            case 32767:
                                switch (v7 | 988) {
                                case 2147483647:
                                    var v17 uint = 638
                                    _ = v17
                                }
                            }
                            var v18 uint16 = 558
                            if (!(1 <= 1)) {
                                var v19 uint16 = (v18 % (((v18 ^ v18)) | 1))
                                var v20 uint16 = v18
                                var v21 float32 = -2.25
                                _ = v19
                                _ = v20
                                _ = v21
                            } else {
                                v11 = append(v11, (p2 >> ((0) & 63)))
                                v18 -= (v18 * v18)
                                p1 = p1
                                v22 := GenS1{F0: 87, F1: 73}
                                p1 += (p1 / ((p1) | 1))
                                _ = v22
                            }
                            v9 = v9
                            v23 := GenS2{F0: (p1 & p1)}
                            _ = v18
                            _ = v23
                        } else {
                            fmt.Println(v9)
                            p2 = (p2 - fi8)
                        }
                        _ = v11
                    } else {
                        fmt.Println(v7, p1)
                        p1 = 1000
                        var v24 uint32 = v9
                        _ = v24
                    }
                    p0 = p0
                    var v25 uint64 = 141
                    if true {
                        v26 := GenS2{F0: (p1 << ((42) & 63))}
                        _ = v26
                    }
                    _ = v25
                }
            default:
                v27 := []int{(fi6 ^ p2)}
                _ = v27
            }
            var v28 uint32 = v9
            _ = fi8
            _ = v28
            _ = v9
        }
        if (true && false) {
            if (true || false) {
                var v29 int16 = p1
                var v30 uint8 = 128
                _ = v29
                _ = v30
            }
            switch (p0 + (p0 + "foo bar")) {
            case "foo bar":
                var v31 bool = (2147483648 >= 115)
                _ = v31
            case "hello":
                for fi32 := 0; fi32 < 15; fi32++ {
                    v33 := []int{(v7 % ((v7) | 1))}
                    _ = fi32
                    _ = v33
                }
            case "":
                var v34 int64 = 0
                _ = v34
            default:
                if false {
                    for fi35 := 0; fi35 < 14; fi35++ {
                        v7 = v7
                        var v36 uint = 4294967295
                        v37 := []int{(v7 % ((fi35) | 1)), v7}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    switch p0 {
                    case "hello":
                        if true {
                            p0 = ((p0 + "") + ((("\n\t" + p0) + ("hello" + p0)) + p0))
                        }
                    }
                    p2 = p2
                }
            }
        }
        _ = fi6
        _ = v7
    }
    if true {
        p2 = 433
        var v38 uint32 = 670
        v39 := GenS2{F0: p1}
        _ = v38
        _ = v39
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 222
}
func genFn1() bool {
    var v40 float32 = 1.5
    for fi41 := 0; fi41 < 1; fi41++ {
        v42 := []int{(fi41 & fi41), fi41}
        fmt.Println(fi41, v40)
        _ = len(v42)
        _ = fi41
        _ = v42
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        switch "foo bar" {
        case "\n\t":
            for fi44 := 0; fi44 < 8; fi44++ {
                if false {
                    v40 = v40
                    var v45 int8 = 1
                    v46 := []int{}
                    _ = v45
                    _ = v46
                }
                _ = fi44
            }
        case "":
            v47 := GenS3{F0: (2 != 0), F1: -2147483648, F2: -590}
            _ = v47
        }
        v40 -= v40
        var v48 int = fi43
        _ = fi43
        _ = v48
    }
    _ = v40
    return (!(-2147483648 < 42))
}

func main() {
    var v1 float32 = (3.14159 - 0.1)
    v2 := []string{("" + "\n\t"), "\n\t", "hello", "foo bar"}
    var v3 float64 = (float64(-2.25) * (0.0 - 1e10))
    v4 := []int{1, -9223372036854775808}
    var v5 uint32 = 65535
    v1 = v1
    v3 += (v3 - (1e10 + 3.14159))
    v3 += 0.1
    if ((1 != 32767) && ("a" != ("a" + "\n\t"))) {
        var v49 uint64 = 125
        for fi50 := 0; fi50 < 3; fi50++ {
            v51 := GenS2{F0: 32767}
            v3 = (v3 / (0.1 + float64((v1 + v1))))
            _ = fi50
            _ = v51
        }
        v52 := GenS3{F0: true, F1: -2147483648, F2: -32768}
        _ = v49
        _ = v52
    } else {
        if true {
            if true {
                var v53 float32 = v1
                _ = v53
            } else {
                for fi54 := 0; fi54 < 3; fi54++ {
                    v5 = 42
                    var v55 uint8 = 128
                    _ = fi54
                    _ = v55
                }
                if len(v4) > 0 {
                    v4[0] = 42
                }
                v5 -= v5
            }
            var v56 int32 = 2147483647
            v56 = v56
            _ = v56
        } else {
            if true {
                v57 := GenS3{F0: (-825 > 906), F1: -1, F2: 42}
                if false {
                    for fi58 := 0; fi58 < 20; fi58++ {
                        var v59 float32 = v1
                        _ = fi58
                        _ = v59
                    }
                    var v60 int32 = -1
                    var v61 uint = 42
                    _ = v60
                    _ = v61
                }
                _ = v57
            }
            if false {
                v3 += float64(3.14159)
            }
        }
    }
    var v62 int8 = 100
    v1 = v1
    for fi63 := 0; fi63 < 6; fi63++ {
        v3 -= 1.5
        v64 := GenS2{F0: 32767}
        _ = fi63
        _ = v64
    }
    v65 := GenS4{F0: 0, F1: 564, F2: 999}
    var v66 float64 = float64((v1 - -2.25))
    if true {
        v3 += v3
        v62 = (v62 << ((1) & 63))
        switch 347 {
        case 532:
            v67 := GenS5{F0: 255, F1: 255, F2: (v5 << ((230) & 63))}
            _ = v67
        default:
            var v68 int32 = 2147483647
            _ = v68
        }
    }
    for fi69 := 0; fi69 < 8; fi69++ {
        v3 += (-2.25 / 0.1)
        fmt.Println(v62, v3, v5)
        var v70 int16 = 0
        _ = fi69
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v65
    _ = v66
}
