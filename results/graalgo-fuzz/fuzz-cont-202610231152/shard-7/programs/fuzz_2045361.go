package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int; F1 uint64; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int; F1 int16; F2 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string }
func genFn0() int16 {
    for fi4 := 0; fi4 < 6; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    for fi5 := 0; fi5 < 16; fi5++ {
        if (1 >= 32767) {
            var v6 int = 1
            v7 := []int{v6, 32767, fi5}
            v6 += v6
            _ = v6
            _ = v7
        } else {
            v8 := []int{(fi5 | 0)}
            var v9 float32 = 1.5
            _ = v8
            _ = v9
        }
        _ = fi5
    }
    var v10 float32 = 0.0
    v11 := []int{-115, 127}
    v10 = v10
    _ = v10
    _ = v11
    return 780
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{2147483647, 364}
    var v3 int = 843
    var v12 uint8 = 2
    if (2 <= 42) {
        v12 -= v12
        var v13 uint8 = v12
        _ = v13
    } else {
        var v14 bool = false
        if ((v12 | v12) != (v12 >> ((1) & 63))) {
            v15 := GenS1{F0: 859}
            var v16 uint64 = 939
            var v17 uint16 = 650
            _ = v15
            _ = v16
            _ = v17
        }
        if v14 {
            v18 := GenS1{F0: 65535}
            if (255 != 0) {
                v19 := GenS1{F0: 1}
                if len(v2) > 0 {
                    v2[0] = 889
                }
                v20 := GenS2{F0: 0}
                _ = v19
                _ = v20
            } else {
                if (((v3 | v3) <= v3) && ("\n\t" <= ("a" + "hello"))) {
                    if (111 >= v12) {
                        for fi21 := 0; fi21 < 6; fi21++ {
                            v22 := GenS3{F0: 1, F1: ""}
                            var v23 int32 = 1
                            for fi24 := 0; fi24 < 4; fi24++ {
                                var v25 uint16 = 1
                                fmt.Println(v3)
                                _ = fi24
                                _ = v25
                            }
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        _ = len(v1)
                    } else {
                        if v14 {
                            var v26 uint8 = v12
                            var v27 uint16 = 969
                            for fi28 := 0; fi28 < 19; fi28++ {
                                v27 -= (v27 << ((539) & 63))
                                v14 = true
                                _ = fi28
                            }
                            _ = v26
                            _ = v27
                        } else {
                            var v29 string = "hello"
                            _ = v29
                        }
                    }
                }
                for fi30 := 0; fi30 < 20; fi30++ {
                    var v31 uint8 = (v12 ^ v12)
                    var v32 uint8 = 128
                    switch "hello" {
                    case "hello":
                        var v33 uint8 = (^(v32))
                        _ = v33
                    case "a":
                        var v34 int64 = -1
                        _ = v34
                    }
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                switch (v3 * -2147483648) {
                case -9223372036854775808:
                    if v14 {
                        if (0 >= 892) {
                            v35 := GenS4{F0: v3, F1: 21, F2: v3}
                            if v14 {
                                switch ("a" + "foo bar") {
                                case "a":
                                    var v36 bool = true
                                    _ = v36
                                case "hello":
                                    var v37 int = (v3 | (v3 / (((v3 | v3)) | 1)))
                                    _ = v37
                                }
                                for fi38 := 0; fi38 < 16; fi38++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v39 uint64 = 0
                                    var v40 uint8 = 35
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                            }
                            v41 := GenS1{F0: 0}
                            _ = v35
                            _ = v41
                        }
                    } else {
                        for fi42 := 0; fi42 < 13; fi42++ {
                            switch (fi42 << ((255) & 63)) {
                            case -919:
                                fmt.Println(v3, v12)
                            case -2147483648:
                                delete(v1, (v3 ^ 42))
                            default:
                                v14 = (v3 <= (fi42 << ((1) & 63)))
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            switch (fi42 % ((42) | 1)) {
                            case 127:
                                var v43 string = ""
                                _ = v43
                            case 48:
                                v44 := GenS4{F0: (v3 / ((9223372036854775807) | 1)), F1: 4294967295, F2: fi42}
                                _ = v44
                            }
                            _ = fi42
                        }
                        v45 := GenS1{F0: 440}
                        var v46 int32 = 2147483647
                        _ = v45
                        _ = v46
                    }
                case 54:
                    v14 = (((v14 || ((v14 || v14) && v14)) || (2 <= 4294967295)) || (74 < v3))
                case -609:
                    v12 = 2
                default:
                    v47 := GenS5{F0: 4294967295}
                    _ = v47
                }
            }
            _ = v18
        }
        _ = v14
    }
    var v48 uint64 = 2
    fmt.Println(v48, v3, v48)
    if false {
        if (42 >= genFn0()) {
            if false {
                if (!(3.14159 < 0.1)) {
                    switch ("\n\t" + ("a" + "hello")) {
                    case "hello":
                        fmt.Println(v3, v12)
                    default:
                        if (-1 == -81) {
                            if len(v2) > 0 {
                                v2[0] = 2147483647
                            }
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                }
                if (604 > 2) {
                    var v49 uint32 = 343
                    fmt.Println(v12)
                    _ = v49
                } else {
                    v12 += v12
                    v48 -= (v48 << ((772) & 63))
                    var v50 bool = (float64((1e10 - 3.14159)) <= 1.5)
                    _ = v50
                }
                for fi51 := 0; fi51 < 8; fi51++ {
                    fmt.Println(v48)
                    v3 += (-(v3))
                    _ = fi51
                }
            } else {
                for fi52 := 0; fi52 < 6; fi52++ {
                    switch (("hello" + ("hello" + ("foo bar" + ("" + ("\n\t" + (((("foo bar" + "hello") + ("\n\t" + "")) + "") + (("foo bar" + (("foo bar" + "") + "\n\t")) + "foo bar"))))))) + "foo bar") {
                    case "foo bar":
                        var v53 uint64 = 42
                        _ = v53
                    }
                    _ = fi52
                }
            }
            var v54 float32 = (-2.25 + -2.25)
            fmt.Println(v3, v48, v48)
            _ = v54
        } else {
            v2 = append(v2, 623)
        }
        v1[v3] = 2
    } else {
        var v55 int = (-(v3))
        delete(v1, v55)
        _ = v55
    }
    v56 := GenS6{F0: (v3 / (((v3 << ((255) & 63))) | 1))}
    v57 := GenS7{F0: 42}
    var v58 uint = 2
    var v59 uint32 = 2147483648
    fmt.Println(v12, v58, v48)
    v3 -= (v3 ^ v3)
    var v60 int32 = 0
    if false {
        var v61 uint = (v58 + 255)
        var v62 float64 = 1.5
        if true {
            _ = len(v2)
            v63 := GenS8{F0: v3, F1: 1, F2: 1}
            fmt.Println(v60)
            _ = v63
        } else {
            v64 := GenS8{F0: (^(v3)), F1: -264, F2: -1}
            var v65 float64 = v62
            if (42 < int64(v60)) {
                v59 = 2
                v66 := GenS9{F0: 32767}
                _ = v66
            } else {
                v58 = v61
                if (42 != 9223372036854775807) {
                    if true {
                        var v67 uint16 = 65535
                        for fi68 := 0; fi68 < 12; fi68++ {
                            v69 := GenS10{F0: "hello"}
                            for fi70 := 0; fi70 < 4; fi70++ {
                                switch "hello" {
                                case "foo bar":
                                    v71 := GenS3{F0: 1, F1: (("foo bar" + "foo bar") + "\n\t")}
                                    _ = v71
                                case "hello":
                                    fmt.Println(v62, v48, v58)
                                case "":
                                    var v72 uint64 = v48
                                    _ = v72
                                }
                                var v73 int64 = 1
                                _ = fi70
                                _ = v73
                            }
                            for fi74 := 0; fi74 < 16; fi74++ {
                                var v75 int = (^(v3))
                                var v76 int = fi74
                                _ = fi74
                                _ = v75
                                _ = v76
                            }
                            _ = fi68
                            _ = v69
                        }
                        _ = v67
                    }
                    v12 = (^(v12))
                    if false {
                        delete(v1, (-(v3)))
                        v58 = 621
                    }
                } else {
                    if (uint64(v61) <= v48) {
                        v2 = append(v2, (v59 << ((0) & 63)))
                    } else {
                        v77 := GenS7{F0: (v59 | v59)}
                        _ = v77
                    }
                    var v78 uint32 = uint32(v58)
                    var v79 float64 = v62
                    _ = v78
                    _ = v79
                }
                fmt.Println(v61, v58)
            }
            _ = v64
            _ = v65
        }
        _ = v61
        _ = v62
    }
    var v80 bool = (1 >= 42)
    if false {
        v60 -= v60
    }
    if v80 {
        v60 = v60
        var v81 uint8 = (^(v12))
        _ = v81
    }
    fmt.Println(v60, v80)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v48
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v80
}
