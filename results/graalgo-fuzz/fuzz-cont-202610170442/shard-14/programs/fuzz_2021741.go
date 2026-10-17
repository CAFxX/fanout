package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 bool; F2 bool }
type GenS3 struct { F0 uint32 }
func genFn0() int64 {
    v6 := []int{4294967295, -1, 4294967295}
    switch ("foo bar" + ((("\n\t" + "hello") + ("hello" + "foo bar")) + ("\n\t" + (("hello" + "\n\t") + "\n\t")))) {
    case "":
        var v7 float32 = 1.5
        _ = v7
    case "\n\t":
        v8 := GenS1{F0: 4294967295}
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: 2}
            if (!(false || (false || (true || true)))) {
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 float64 = 0.1
                    v12 = (v12 / float64(1e10))
                    _ = fi11
                    _ = v12
                }
            } else {
                switch "a" {
                case "":
                    var v13 uint = 18446744073709551615
                    _ = v13
                case "hello":
                    v14 := GenS1{F0: 255}
                    _ = v14
                }
                if (("\n\t" + ("a" + "foo bar")) != "foo bar") {
                    _ = len(v6)
                    for fi15 := 0; fi15 < 6; fi15++ {
                        fmt.Println(fi15, fi15)
                        fmt.Println(fi9, fi9)
                        _ = fi15
                    }
                    var v16 uint32 = 860
                    _ = v16
                }
            }
            fmt.Println(fi9)
            _ = fi9
            _ = v10
        }
    default:
        var v17 uint64 = 799
        _ = v17
    }
    var v18 uint16 = 65535
    v18 = 255
    _ = v18
    _ = v6
    return -443
}
func genFn1(p0 int16, p1 int64, p2 bool) uint8 {
    var v19 float64 = 0.0
    var v20 float32 = -2.25
    v21 := []int{32767, 32768}
    fmt.Println(v19)
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 6; fi23++ {
            fmt.Println(fi23)
            p0 = (p0 & int16(p1))
            p0 += (p0 ^ (p0 | -1))
            var v24 string = "foo bar"
            p0 = (p0 + 610)
            _ = fi23
            _ = v24
        }
        if (3.14159 > (v20 - (0.1 - (v20 / 1.5)))) {
            var v25 uint64 = 4294967295
            switch -714 {
            case 9223372036854775807:
                v26 := GenS1{F0: 724}
                _ = v26
            case 127:
                switch (-(fi22)) {
                case -970:
                    for fi27 := 0; fi27 < 18; fi27++ {
                        fmt.Println(fi22)
                        var v28 bool = (p2 || (3.14159 > 3.14159))
                        _ = fi27
                        _ = v28
                    }
                case 32768:
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                case -1:
                    for fi29 := 0; fi29 < 5; fi29++ {
                        v25 = 0
                        var v30 int32 = -1
                        var v31 int16 = p0
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                }
            default:
                for fi32 := 0; fi32 < 19; fi32++ {
                    if len(v21) > 0 {
                        v21[0] = 27
                    }
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    p2 = (!((float64(v20) - float64((v20 - v20))) <= v19))
                    if true {
                        if p2 {
                            var v33 int = fi32
                            p2 = (("a" + "a") > "")
                            fmt.Println(p1)
                            _ = v33
                        } else {
                            if (774 > 2147483647) {
                                var v34 string = ("\n\t" + (("" + "\n\t") + ("\n\t" + "\n\t")))
                                _ = v34
                            } else {
                                if (("hello" + (("a" + "") + "foo bar")) != ("foo bar" + "hello")) {
                                    if (false || ((127 < 160) || (("\n\t" + (("" + ("hello" + (("hello" + "a") + ""))) + "a")) < "foo bar"))) {
                                        v35 := GenS1{F0: 266}
                                        var v36 uint16 = 32767
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                        var v37 int8 = -104
                                        p0 += int16(fi32)
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                    var v38 int16 = 0
                                    if p2 {
                                        v39 := GenS1{F0: 18446744073709551615}
                                        _ = v39
                                    } else {
                                        v19 += (float64(v20) - (v19 + v19))
                                        if len(v21) > 0 {
                                            v21[0] = (fi22 % ((int(v25)) | 1))
                                        }
                                    }
                                    _ = v38
                                }
                            }
                            switch "a" {
                            case "a":
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                            case "":
                                p1 -= p1
                            }
                            v25 = 0
                        }
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                    } else {
                        p1 = (p1 * (p1 << ((360) & 63)))
                    }
                    _ = fi32
                }
            }
            p1 = genFn0()
            _ = v25
        }
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v21
    return 83
}

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]int)
    v3 := []uint32{2}
    var v4 int32 = 1
    v5 := []int{-1}
    switch ((("\n\t" + "") + "a") + ("hello" + "hello")) {
    case "foo bar":
        if true {
            for fi40 := 0; fi40 < 14; fi40++ {
                fmt.Println(v4, v4)
                _ = fi40
            }
            v41 := GenS1{F0: 18446744073709551615}
            _ = v41
        } else {
            v42 := GenS1{F0: 3}
            for fi43 := 0; fi43 < 19; fi43++ {
                var v44 int8 = (-(v1))
                if false {
                    if false {
                        for fi45 := 0; fi45 < 7; fi45++ {
                            var v46 uint = 255
                            _ = fi45
                            _ = v46
                        }
                        var v47 bool = true
                        _ = v47
                    }
                    v48 := GenS2{F0: fi43, F1: true, F2: (9223372036854775807 == fi43)}
                    if (true || (739 == 255)) {
                        v49 := GenS2{F0: (fi43 / ((fi43) | 1)), F1: (!true), F2: ((fi43 / ((fi43) | 1)) != 42)}
                        v50 := GenS3{F0: 159}
                        if true {
                            v51 := v2[2147483647]
                            _ = v51
                        }
                        _ = v49
                        _ = v50
                    } else {
                        for fi52 := 0; fi52 < 5; fi52++ {
                            if len(v3) > 0 {
                                v3[0] = 42
                            }
                            v1 = (v1 << ((65535) & 63))
                            _ = fi52
                        }
                    }
                    _ = v48
                }
                delete(v2, 42)
                _ = fi43
                _ = v44
            }
            _ = v42
        }
    case "hello":
        v4 += v4
    }
    var v53 uint64 = 1
    var v54 uint = 1
    var v55 int8 = 117
    v55 = 1
    for fi56 := 0; fi56 < 15; fi56++ {
        if true {
            fmt.Println(v54, v54, v1)
            var v57 float32 = float32((1.5 * 1e10))
            _ = v57
        }
        var v58 uint16 = 342
        v55 = int8(v58)
        _ = fi56
        _ = v58
    }
    var v59 string = ("" + (("" + "foo bar") + ("foo bar" + "")))
    var v60 int16 = 32767
    var v61 int = 32
    _ = len(v2)
    var v62 uint64 = v53
    v61 += v61
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v59
    _ = v60
    _ = v61
    _ = v62
}
