package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int }
type GenS4 struct { F0 uint64; F1 int; F2 int64 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int8; F1 uint64; F2 string }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int }
func genFn0(p0 int64) int64 {
    p0 += (p0 >> ((65535) & 63))
    p0 = p0
    v7 := GenS1{F0: 690}
    var v8 uint32 = 435
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    return (p0 >> ((4294967295) & 63))
}
func genFn1(p0 uint8, p1 uint16, p2 int16) uint16 {
    switch "hello" {
    case "foo bar":
        p2 = p2
    case "hello":
        v9 := []int{int(p0), 231, 42}
        _ = v9
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -830
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v5 uint = 0
    var v6 bool = true
    v10 := GenS2{F0: 42, F1: 127}
    v11 := GenS1{F0: -776}
    v12 := []int{0, 190, 2147483647, int(v2)}
    fmt.Println(v2)
    if false {
        if (!v6) {
            var v13 int = 618
            _ = v13
        }
        if false {
            var v14 int32 = 174
            _ = v14
        }
    }
    if (128 >= 0) {
        v2 -= (v2 % ((v2) | 1))
    }
    var v15 int32 = 1
    if len(v12) > 0 {
        v12[0] = -219
    }
    var v16 int16 = 975
    fmt.Println(v5)
    for fi17 := 0; fi17 < 6; fi17++ {
        switch ("hello" + "") {
        case "foo bar":
            for fi18 := 0; fi18 < 5; fi18++ {
                if ((!(!(1 >= 393))) || (false && (20 <= 1))) {
                    switch (fi17 | fi17) {
                    case -1:
                        var v19 uint32 = 65535
                        _ = v19
                    case 484:
                        if false {
                            if v6 {
                                v20 := GenS1{F0: -840}
                                fmt.Println(v2)
                                _ = v20
                            } else {
                                if v6 {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    if (-60 <= -52) {
                                        if v6 {
                                            if v6 {
                                                v15 = v15
                                            } else {
                                                if (4294967295 <= 2) {
                                                    v21 := GenS1{F0: (v16 & (-(v16)))}
                                                    _ = v21
                                                } else {
                                                    v22 := GenS3{F0: 42, F1: 2, F2: (fi18 << ((255) & 63))}
                                                    fmt.Println(fi18)
                                                    _ = v22
                                                }
                                                v16 = v16
                                                var v23 float64 = 1e10
                                                _ = v23
                                            }
                                            switch (fi17 & (fi17 | (fi17 & fi18))) {
                                            case 127:
                                                v2 = (v2 * v2)
                                            case -886:
                                                var v24 string = ("" + "foo bar")
                                                _ = v24
                                            case 523:
                                                var v25 int64 = v2
                                                _ = v25
                                            }
                                            v26 := v3[fi17]
                                            _ = v26
                                        } else {
                                            var v27 int64 = (v2 & (v2 & (v2 << ((v5) & 63))))
                                            _ = v27
                                        }
                                    }
                                    v12 = append(v12, (fi17 - 1))
                                }
                                v28 := GenS4{F0: 781, F1: (^(fi18)), F2: (v2 >> ((v5) & 63))}
                                _ = v28
                            }
                            switch (("\n\t" + "\n\t") + (("" + "foo bar") + ("" + ("\n\t" + "")))) {
                            case "hello":
                                var v29 float64 = float64(((-2.25 - 0.0) - 1e10))
                                fmt.Println(v16)
                                fmt.Println(v5, v6)
                                fmt.Println(v6)
                                if v6 {
                                    if v6 {
                                        if (v6 && v6) {
                                            v30 := GenS4{F0: 4294967295, F1: fi17, F2: v2}
                                            fmt.Println(v6, v2)
                                            _ = v30
                                        }
                                        var v31 uint = v5
                                        var v32 int8 = 21
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v33 int8 = 1
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        var v34 uint = (v5 * (v5 + v5))
                                        switch fi17 {
                                        case -9223372036854775808:
                                            delete(v1, 0)
                                        case -786:
                                            v35 := GenS5{F0: true, F1: fi18}
                                            _ = v35
                                        case 0:
                                            var v36 string = (("foo bar" + "") + "")
                                            v36 = ("\n\t" + "")
                                            v37 := GenS1{F0: v16}
                                            v12 = append(v12, (fi17 & 4294967295))
                                            if len(v12) > 0 {
                                                v12[0] = (fi18 << (((^(v5))) & 63))
                                            }
                                            _ = v36
                                            _ = v37
                                        default:
                                            switch (fi17 / (((fi17 >> (((v5 % ((791) | 1))) & 63))) | 1)) {
                                            case 42:
                                                var v38 uint64 = 18446744073709551615
                                                _ = v38
                                            case -9223372036854775808:
                                                if v6 {
                                                    if (v6 || v6) {
                                                        var v39 float64 = 1e10
                                                        if len(v12) >= 2 {
                                                            v12 = v12[1:len(v12)]
                                                        }
                                                        _ = v39
                                                    }
                                                }
                                            case 305:
                                                v40 := v3[(fi18 >> ((36) & 63))]
                                                _ = v40
                                            default:
                                                var v41 uint64 = 509
                                                _ = v41
                                            }
                                        }
                                        _ = v33
                                        _ = v34
                                    }
                                    var v42 uint64 = 18446744073709551615
                                    var v43 bool = (!v6)
                                    fmt.Println(v16)
                                    _ = v42
                                    _ = v43
                                }
                                v1[35] = ""
                                _ = v29
                            case "":
                                if v6 {
                                    _ = len(v12)
                                    var v44 bool = (!(42 == (v15 | v15)))
                                    v45 := GenS6{F0: 1, F1: 89, F2: "\n\t"}
                                    var v46 uint32 = 1
                                    _ = len(v1)
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                } else {
                                    v2 += -901
                                }
                            case "foo bar":
                                var v47 uint8 = 255
                                var v48 int32 = (v15 | v15)
                                fmt.Println(v6, v15)
                                fmt.Println(v5)
                                if (642 < 4294967295) {
                                    v15 -= -1
                                    _ = len(v3)
                                }
                                v15 += (v48 << (((v5 * (v5 & (^(v5))))) & 63))
                                _ = v47
                                _ = v48
                            default:
                                var v49 int16 = v16
                                _ = v49
                            }
                            switch fi17 {
                            case 32767:
                                v50 := GenS7{F0: 242, F1: v2}
                                _ = v50
                            case 0:
                                fmt.Println(v6, v2, fi18)
                            }
                        } else {
                            v5 += 722
                            switch ("" + "\n\t") {
                            case "":
                                if (284 != v5) {
                                    v2 += (v2 ^ genFn0(v2))
                                } else {
                                    v51 := GenS8{F0: -1}
                                    fmt.Println(v16)
                                    var v52 int32 = -1
                                    var v53 uint = v5
                                    v54 := GenS9{F0: v5}
                                    delete(v3, (fi18 >> ((v53) & 63)))
                                    if (!(!(fi17 < (fi17 + (fi18 | fi17))))) {
                                        v55 := GenS6{F0: 127, F1: 518, F2: ("hello" + "hello")}
                                        fmt.Println(v16)
                                        _ = v55
                                    } else {
                                        var v56 uint64 = 0
                                        _ = v56
                                    }
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                            }
                            delete(v3, fi18)
                            _ = len(v12)
                            v6 = (1.5 >= 0.0)
                            if v6 {
                                v16 -= (v16 % ((382) | 1))
                                v57 := GenS5{F0: v6, F1: int(v16)}
                                _ = v57
                            }
                        }
                    case -9223372036854775808:
                        switch (((("hello" + (("" + (((("hello" + ("hello" + ("foo bar" + "hello"))) + "foo bar") + "a") + "")) + "\n\t")) + "") + "a") + "a") {
                        case "foo bar":
                            v5 = (v5 | v5)
                        case "":
                            v15 = (v15 - v15)
                        default:
                            if len(v12) > 0 {
                                v12[0] = -2147483648
                            }
                            v3[fi17] = fi17
                            var v58 uint8 = 2
                            _ = v58
                        }
                        fmt.Println(v2, fi17, v15)
                        v6 = true
                        switch ("\n\t" + "") {
                        case "hello":
                            fmt.Println(v15, v5, fi18)
                        default:
                            var v59 float64 = (((0.0 - -2.25) + 0.0) * 1e10)
                            _ = v59
                        }
                        v16 = v16
                    default:
                        var v60 uint64 = 2
                        _ = v60
                    }
                    switch (("a" + "foo bar") + ("hello" + "\n\t")) {
                    case "a":
                        if (1 != 18446744073709551615) {
                            v2 = genFn0((v2 + 9223372036854775807))
                            v61 := GenS10{F0: fi17}
                            var v62 bool = (v6 && (127 <= 121))
                            _ = v61
                            _ = v62
                        }
                    }
                    fmt.Println(fi18, v16, v5)
                }
                _ = fi18
            }
        default:
            v63 := GenS8{F0: 100}
            _ = v63
        }
        _ = fi17
    }
    var v64 uint = (v5 / ((v5) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
