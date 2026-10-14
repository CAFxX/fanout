package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int32; F2 bool }
type GenS4 struct { F0 uint; F1 uint; F2 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 uint32; F1 uint32 }
func genFn0(p0 string, p1 uint16, p2 uint64) uint32 {
    v7 := []int{963, 560, 32768, 923}
    for fi8 := 0; fi8 < 4; fi8++ {
        if true {
            fmt.Println(p0, fi8, p1)
            for fi9 := 0; fi9 < 14; fi9++ {
                if true {
                    var v10 uint64 = (p2 << ((65535) & 63))
                    fmt.Println(v10)
                    p1 = (p1 + p1)
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    var v11 int = 42
                    _ = v10
                    _ = v11
                } else {
                    var v12 uint = 703
                    if (0 >= 74) {
                        v13 := GenS1{F0: 1}
                        if len(v7) > 0 {
                            v7[0] = (fi8 / ((fi8) | 1))
                        }
                        _ = v13
                    } else {
                        switch fi9 {
                        case 208:
                            p1 -= (p1 >> (((v12 - v12)) & 63))
                        default:
                            p1 -= 151
                        }
                        var v14 uint32 = uint32(v12)
                        var v15 uint = v12
                        _ = v14
                        _ = v15
                    }
                    if (734 > -2147483648) {
                        var v16 int8 = 0
                        fmt.Println(p1, fi9, p2)
                        fmt.Println(p2, fi8, v12)
                        switch (p0 + p0) {
                        case "a":
                            if (false || true) {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            } else {
                                if len(v7) > 0 {
                                    v7[0] = (fi8 ^ -762)
                                }
                            }
                        case "foo bar":
                            fmt.Println(v16, p0)
                        }
                        _ = v16
                    } else {
                        v17 := GenS1{F0: 743}
                        switch (fi9 / (((^(fi9))) | 1)) {
                        case 0:
                            p0 = p0
                        case 4294967295:
                            v12 = (^(v12))
                        }
                        var v18 bool = false
                        _ = v17
                        _ = v18
                    }
                    _ = v12
                }
                p0 = p0
                if ((("foo bar" != ("a" + "foo bar")) || (((p1 % (((p1 << ((42) & 63))) | 1)) <= (^(p1))) && false)) || ((fi9 | (fi9 << ((681) & 63))) != int(p1))) {
                    if true {
                        var v19 float32 = 0.0
                        p2 = p2
                        v20 := GenS2{F0: (p2 / ((p2) | 1))}
                        var v21 string = ("hello" + "\n\t")
                        p0 = v21
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    var v22 uint16 = (p1 & p1)
                    _ = v22
                }
                _ = fi9
            }
        } else {
            var v23 float64 = (1.5 + 1.5)
            p1 = 781
            _ = v23
        }
        _ = fi8
    }
    var v24 float64 = 0.1
    switch "\n\t" {
    case "foo bar":
        for fi25 := 0; fi25 < 7; fi25++ {
            for fi26 := 0; fi26 < 14; fi26++ {
                p2 = (p2 >> ((1) & 63))
                _ = fi26
            }
            if true {
                fmt.Println(v24, v24, p1)
            } else {
                var v27 string = ("" + "")
                _ = v27
            }
            if (!true) {
                var v28 int8 = -1
                _ = v28
            } else {
                fmt.Println(fi25, p1)
                if (1 > 42) {
                    switch p0 {
                    case "\n\t":
                        var v29 uint32 = uint32(p1)
                        _ = v29
                    case "foo bar":
                        v30 := GenS3{F0: 0, F1: -349, F2: (!false)}
                        _ = v30
                    default:
                        v31 := GenS2{F0: 2}
                        _ = v31
                    }
                    p1 += p1
                } else {
                    var v32 uint = 4294967295
                    _ = v32
                }
            }
            _ = fi25
        }
    case "\n\t":
        var v33 uint32 = 4294967295
        _ = v33
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v7
    return 42
}
func genFn1() string {
    if true {
        v34 := GenS1{F0: 108}
        var v35 int64 = 0
        _ = v34
        _ = v35
    }
    var v36 int16 = 948
    for fi37 := 0; fi37 < 15; fi37++ {
        v36 = v36
        _ = fi37
    }
    _ = v36
    return "foo bar"
}

func main() {
    var v1 uint32 = 0
    var v2 uint16 = 2
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    var v5 string = ("foo bar" + "\n\t")
    var v6 int16 = 1
    for fi38 := 0; fi38 < 15; fi38++ {
        v6 = 0
        _ = fi38
    }
    for fi39 := 0; fi39 < 5; fi39++ {
        v6 += 42
        v40 := []int{fi39, fi39, -208, fi39}
        v2 += 0
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        var v42 int64 = 0
        _ = fi41
        _ = v42
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        var v44 uint = 4294967295
        _ = fi43
        _ = v44
    }
    var v45 int16 = -1
    v46 := GenS4{F0: 42, F1: 1, F2: (v1 >> ((2) & 63))}
    v47 := GenS1{F0: (v1 ^ (^(v1)))}
    if true {
        if false {
            v45 = (v45 - v6)
        } else {
            switch -9223372036854775808 {
            case 127:
                if false {
                    v2 += 494
                    var v48 int = 2147483647
                    fmt.Println(v48, v45)
                    _ = v48
                }
            case 2147483647:
                if false {
                    var v49 int16 = 32767
                    _ = v49
                }
            case 42:
                v1 -= v1
            }
        }
        v50 := GenS3{F0: 225, F1: -1, F2: false}
        for fi51 := 0; fi51 < 2; fi51++ {
            v2 -= 255
            v2 = v2
            _ = fi51
        }
        _ = v50
    }
    if false {
        v52 := GenS1{F0: (v1 * v1)}
        for fi53 := 0; fi53 < 15; fi53++ {
            for fi54 := 0; fi54 < 7; fi54++ {
                v55 := GenS5{F0: "foo bar"}
                switch v5 {
                case "hello":
                    var v56 bool = true
                    _ = v56
                case "\n\t":
                    v6 = 0
                case "foo bar":
                    var v57 float64 = (0.1 + (0.0 / float64((float32(1.5) - 3.14159))))
                    _ = v57
                }
                v58 := []int{(fi53 + (fi53 + fi54)), (-(fi53))}
                _ = fi54
                _ = v55
                _ = v58
            }
            switch "foo bar" {
            case "hello":
                delete(v3, fi53)
            default:
                v2 -= (v2 / (((v2 | v2)) | 1))
            }
            v59 := GenS6{F0: 688, F1: fi53}
            _ = fi53
            _ = v59
        }
        _ = v52
    }
    v60 := []int{42}
    if false {
        var v61 uint64 = 1
        _ = v61
    }
    v45 -= int16(v2)
    v62 := GenS7{F0: v1, F1: v1}
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v60:", len(v60), v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v60
    _ = v62
}
