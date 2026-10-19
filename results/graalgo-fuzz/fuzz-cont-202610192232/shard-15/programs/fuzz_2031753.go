package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint; F1 string; F2 string }
func genFn0(p0 int, p1 int, p2 uint8) int16 {
    switch p0 {
    case 4294967295:
        var v6 int8 = 0
        _ = v6
    case 2147483647:
        switch "foo bar" {
        case "":
            v7 := GenS1{F0: 48, F1: p1, F2: p1}
            _ = v7
        }
    }
    switch (-(p1)) {
    case 289:
        fmt.Println(p1, p0, p0)
    default:
        if true {
            var v8 uint64 = 18446744073709551615
            _ = v8
        } else {
            switch (p0 + (-(p1))) {
            case 127:
                v9 := GenS2{F0: -705, F1: 42, F2: "a"}
                _ = v9
            }
            v10 := GenS2{F0: 0, F1: 902, F2: ("hello" + "a")}
            switch "foo bar" {
            case "":
                var v11 int16 = 0
                _ = v11
            default:
                if true {
                    if true {
                        for fi12 := 0; fi12 < 3; fi12++ {
                            if true {
                                v13 := GenS2{F0: -2147483648, F1: 776, F2: "\n\t"}
                                fmt.Println(fi12, p0)
                                _ = v13
                            }
                            var v14 uint8 = (p2 ^ 0)
                            v15 := []int{}
                            _ = fi12
                            _ = v14
                            _ = v15
                        }
                        v16 := GenS1{F0: 803, F1: (-(p0)), F2: p0}
                        _ = v16
                    } else {
                        var v17 uint64 = 0
                        v18 := []int{p0, (p0 << ((4294967295) & 63))}
                        _ = v17
                        _ = v18
                    }
                }
            }
            _ = v10
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        p2 -= (p2 % ((1) | 1))
        var v20 float64 = 3.14159
        if true {
            v21 := []int{p0, p0}
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1() uint32 {
    for fi22 := 0; fi22 < 14; fi22++ {
        v23 := []int{fi22, fi22, (fi22 - fi22), fi22}
        _ = fi22
        _ = v23
    }
    var v24 uint8 = 211
    if (15 != 2) {
        for fi25 := 0; fi25 < 10; fi25++ {
            switch (-(fi25)) {
            case 128:
                v26 := GenS3{F0: fi25}
                _ = v26
            case 4294967295:
                for fi27 := 0; fi27 < 3; fi27++ {
                    v24 = 194
                    _ = fi27
                }
            default:
                var v28 int16 = 1000
                _ = v28
            }
            var v29 float32 = (-2.25 - (0.0 * 1.5))
            var v30 int32 = 1
            _ = fi25
            _ = v29
            _ = v30
        }
        var v31 uint8 = v24
        _ = v31
    }
    var v32 uint16 = 0
    var v33 uint32 = 2
    _ = v24
    _ = v32
    _ = v33
    return (^(v33))
}

func main() {
    v1 := []string{"a", "hello", "", ("\n\t" + "")}
    var v2 int8 = -128
    var v3 uint32 = 2147483648
    var v4 int16 = 594
    v5 := make(map[string]uint32)
    v34 := GenS4{F0: (1 > 42)}
    v35 := GenS4{F0: ((^(v3)) != v3)}
    var v36 int16 = v4
    var v37 string = "foo bar"
    var v38 uint16 = 32767
    fmt.Println(v4, v3)
    if false {
        fmt.Println(v38, v3, v4)
    } else {
        for fi39 := 0; fi39 < 10; fi39++ {
            var v40 string = (("hello" + v37) + "a")
            _ = fi39
            _ = v40
        }
        var v41 bool = ((v2 != (v2 / (((v2 ^ v2)) | 1))) && (1e10 <= (1.5 - -2.25)))
        if false {
            var v42 float64 = 1.5
            _ = v42
        } else {
            v43 := GenS5{F0: 96}
            if (1 != 0) {
                v4 += (v36 + (v4 * (v4 * genFn0(-2147483648, -420, 0))))
                v44 := GenS1{F0: 32767, F1: -1, F2: -250}
                if v41 {
                    v3 -= genFn1()
                } else {
                    v36 += genFn0(9223372036854775807, 42, 0)
                    if (0 != 397) {
                        v3 -= genFn1()
                    } else {
                        v38 = v38
                        for fi45 := 0; fi45 < 1; fi45++ {
                            v2 = 1
                            _ = fi45
                        }
                    }
                }
                _ = v44
            }
            var v46 uint8 = 248
            _ = v43
            _ = v46
        }
        _ = v41
    }
    if true {
        for fi47 := 0; fi47 < 17; fi47++ {
            _ = len(v5)
            _ = len(v1)
            if false {
                v3 = (v3 / ((genFn1()) | 1))
                if true {
                    v5["foo bar"] = v3
                    if ((1.5 + 0.1) >= float32(float64(3.14159))) {
                        v38 -= 499
                        var v48 int16 = (v36 << ((65535) & 63))
                        var v49 float32 = ((float32(1e10) * float32((0.0 / 0.1))) - 1e10)
                        _ = v48
                        _ = v49
                    } else {
                        var v50 int32 = -121
                        if len(v1) > 0 {
                            v1[0] = (v37 + v37)
                        }
                        _ = v50
                    }
                    v1 = append(v1, "foo bar")
                }
                v2 = 1
            } else {
                _ = len(v1)
                for fi51 := 0; fi51 < 14; fi51++ {
                    v37 = v37
                    _ = len(v5)
                    var v52 uint = 815
                    if false {
                        fmt.Println(v52)
                        if false {
                            _ = len(v1)
                            if false {
                                var v53 string = ("\n\t" + "foo bar")
                                if true {
                                    fmt.Println(v2)
                                    if true {
                                        v54 := GenS6{F0: -32768, F1: (^(v3))}
                                        _ = v54
                                    } else {
                                        var v55 string = (("foo bar" + v37) + "")
                                        _ = v55
                                    }
                                    var v56 float32 = (float32(1.5) / float32(1e10))
                                    _ = v56
                                }
                                _ = v53
                            } else {
                                v52 += (v52 ^ (v52 ^ v52))
                                if true {
                                    v57 := GenS7{F0: v4}
                                    var v58 string = ("a" + (v37 + (v37 + "hello")))
                                    if true {
                                        delete(v5, v37)
                                        var v59 int8 = 21
                                        v1 = append(v1, (((v58 + "\n\t") + v37) + v58))
                                        _ = v59
                                    } else {
                                        var v60 string = ""
                                        var v61 float32 = (0.0 / 3.14159)
                                        if true {
                                            v62 := GenS5{F0: 0}
                                            if false {
                                                if true {
                                                    v63 := GenS7{F0: (v36 * (v36 * v4))}
                                                    _ = v63
                                                }
                                                if false {
                                                    v60 = "foo bar"
                                                    if ((!(v2 != (v2 - 0))) && (!(152 < 2))) {
                                                        var v64 uint = (v52 % ((255) | 1))
                                                        _ = v64
                                                    } else {
                                                        v5[v37] = (v3 * v3)
                                                    }
                                                }
                                            }
                                            if false {
                                                v4 = v4
                                                var v65 uint8 = 255
                                                _ = v65
                                            }
                                            _ = v62
                                        } else {
                                            v66 := GenS8{F0: 2, F1: v37, F2: ""}
                                            var v67 int64 = -1
                                            _ = v66
                                            _ = v67
                                        }
                                        _ = v60
                                        _ = v61
                                    }
                                    _ = v57
                                    _ = v58
                                }
                            }
                            v2 -= (v2 - (v2 / ((v2) | 1)))
                        }
                        var v68 int = fi51
                        if (155 == 255) {
                            var v69 uint8 = 0
                            _ = v69
                        } else {
                            v38 -= (v38 ^ v38)
                        }
                        _ = v68
                    } else {
                        var v70 float64 = (0.1 * (1e10 - 1e10))
                        _ = v70
                    }
                    delete(v5, "a")
                    _ = fi51
                    _ = v52
                }
            }
            _ = fi47
        }
    } else {
        v37 = ((((("" + ("" + "\n\t")) + (("foo bar" + (v37 + v37)) + "\n\t")) + "a") + ((((("" + "hello") + ((v37 + "\n\t") + (v37 + v37))) + "foo bar") + "a") + v37)) + ("" + ""))
        v71 := GenS7{F0: (v36 + (v4 >> ((18446744073709551615) & 63)))}
        var v72 int8 = (v2 << ((uint(v38)) & 63))
        _ = v71
        _ = v72
    }
    var v73 int16 = (v36 & v4)
    var v74 string = "foo bar"
    switch (v37 + "\n\t") {
    case "hello":
        var v75 int64 = -1
        _ = v75
    default:
        v5[("hello" + v74)] = v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v73
    _ = v74
}
