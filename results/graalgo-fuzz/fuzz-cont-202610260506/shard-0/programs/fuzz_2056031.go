package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 int64 }
func genFn0(p0 int, p1 uint32) uint32 {
    if false {
        var v4 int = (p0 | (p0 - -2147483648))
        if (!false) {
            v5 := []int{(v4 & (v4 >> ((42) & 63)))}
            var v6 int8 = 42
            var v7 bool = (-2.25 == (0.1 + (-2.25 * (-2.25 + float64(0.0)))))
            _ = v5
            _ = v6
            _ = v7
        }
        v4 -= p0
        _ = v4
    } else {
        var v8 float32 = 1e10
        v9 := []int{(p0 * -529), p0, -1, p0}
        v9 = append(v9, p0)
        _ = v8
        _ = v9
    }
    var v10 int = p0
    fmt.Println(p0, v10, v10)
    v10 = v10
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS1{F0: -9223372036854775808}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return (p1 ^ 345)
}
func genFn1(p0 int16) uint32 {
    p0 = (-(p0))
    v13 := GenS1{F0: 1}
    var v14 bool = false
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS2{F0: false, F1: (-(p0))}
        v17 := GenS3{F0: 1, F1: "a"}
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := []int{9223372036854775807, 42, 128}
    _ = p0
    _ = v13
    _ = v14
    _ = v18
    return 575
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint64 = 2
    var v3 int64 = -642
    if (!true) {
        v19 := GenS2{F0: (v3 < (v3 ^ (-(v3)))), F1: -1}
        v2 += v2
        var v20 uint16 = 1
        _ = v19
        _ = v20
    }
    if true {
        for fi21 := 0; fi21 < 18; fi21++ {
            switch "hello" {
            case "\n\t":
                v3 = v3
            case "hello":
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 int16 = 1000
                    if true {
                        v23 = v23
                        var v24 int64 = (v3 ^ 0)
                        if false {
                            v25 := GenS3{F0: 42, F1: ("\n\t" + "foo bar")}
                            var v26 uint64 = 2
                            _ = v25
                            _ = v26
                        }
                        _ = v24
                    }
                    v27 := []int{4294967295, (fi22 << (((v1 ^ (v1 & v1))) & 63)), (fi21 / (((fi22 >> (((v1 ^ 0)) & 63))) | 1)), fi21}
                    _ = fi22
                    _ = v23
                    _ = v27
                }
            default:
                for fi28 := 0; fi28 < 5; fi28++ {
                    var v29 bool = true
                    _ = fi28
                    _ = v29
                }
            }
            v3 -= 1
            _ = fi21
        }
    }
    var v30 uint16 = 42
    v31 := []int{128, 128, int(v30)}
    var v32 float64 = 0.0
    if ((^(v2)) < (v2 + v2)) {
        fmt.Println(v30, v30, v3)
        v33 := GenS3{F0: -196, F1: "\n\t"}
        switch 32768 {
        case 50:
            v34 := GenS1{F0: v3}
            _ = v34
        case 944:
            if (42 <= (v3 - v3)) {
                var v35 int16 = 873
                v3 += v3
                _ = v35
            }
        }
        _ = v33
    } else {
        if (127 < 188) {
            if len(v31) >= 2 {
                v31 = v31[1:len(v31)]
            }
        }
        var v36 int8 = 42
        _ = v36
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        for fi38 := 0; fi38 < 11; fi38++ {
            var v39 int8 = 42
            v30 += v30
            v40 := GenS3{F0: -2147483648, F1: ""}
            v41 := GenS3{F0: 0, F1: ""}
            v42 := GenS4{F0: v39, F1: -790}
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
            _ = v42
        }
        var v43 string = ("a" + (("foo bar" + "a") + "a"))
        _ = fi37
        _ = v43
    }
    v44 := GenS4{F0: 100, F1: -32768}
    v45 := GenS4{F0: 127, F1: 42}
    var v46 float32 = (1e10 * ((float32(v32) + 0.1) + 1.5))
    _ = len(v31)
    v47 := GenS3{F0: 987, F1: ("a" + "a")}
    for fi48 := 0; fi48 < 2; fi48++ {
        if (true && (1 != 65535)) {
            var v49 int64 = -1
            _ = len(v31)
            switch (fi48 - (fi48 ^ 127)) {
            case 128:
                v50 := GenS5{F0: 239}
                _ = v50
            default:
                v51 := GenS6{F0: v1, F1: 744}
                _ = v51
            }
            _ = v49
        }
        v52 := GenS7{F0: v3}
        if true {
            if ((v30 & v30) > (v30 / ((v30) | 1))) {
                v31 = append(v31, 2147483647)
                var v53 uint8 = 128
                var v54 uint32 = genFn0(fi48, 2147483648)
                _ = v53
                _ = v54
            } else {
                var v55 uint64 = v2
                v55 = (v55 + (v2 * (v55 | 2)))
                v31 = append(v31, (fi48 >> (((v1 - v1)) & 63)))
                _ = v55
            }
            var v56 string = (("hello" + "hello") + "a")
            _ = v56
        } else {
            if true {
                for fi57 := 0; fi57 < 2; fi57++ {
                    if true {
                        v3 = (v3 - (v3 | v3))
                        if true {
                            if len(v31) >= 2 {
                                v31 = v31[1:len(v31)]
                            }
                        }
                        _ = len(v31)
                    } else {
                        if (true || (42 == v1)) {
                            var v58 string = ((("foo bar" + "foo bar") + ("hello" + "")) + "")
                            fmt.Println(v32, v1, v2)
                            _ = v58
                        } else {
                            var v59 bool = false
                            _ = v59
                        }
                    }
                    var v60 uint8 = 103
                    var v61 int = (fi48 << (((^(v1))) & 63))
                    _ = fi57
                    _ = v60
                    _ = v61
                }
                for fi62 := 0; fi62 < 1; fi62++ {
                    fmt.Println(fi62, fi62, fi48)
                    switch ("a" + "a") {
                    case "":
                        v31 = append(v31, -9223372036854775808)
                    case "hello":
                        v30 = 42
                    case "foo bar":
                        var v63 float32 = (1e10 - v46)
                        _ = v63
                    }
                    _ = fi62
                }
            } else {
                var v64 int16 = 0
                switch ((("hello" + "a") + "a") + ("" + "")) {
                case "\n\t":
                    v3 = v3
                case "":
                    if false {
                        v31 = append(v31, int(v2))
                        for fi65 := 0; fi65 < 3; fi65++ {
                            if len(v31) > 0 {
                                v31[0] = (fi65 % ((fi65) | 1))
                            }
                            fmt.Println(v3)
                            _ = fi65
                        }
                        var v66 float64 = 1e10
                        _ = v66
                    } else {
                        v32 = v32
                        v67 := GenS2{F0: (!((true || (20 >= 2)) || false)), F1: (v64 >> ((v1) & 63))}
                        var v68 uint = 512
                        _ = v67
                        _ = v68
                    }
                default:
                    v1 = v1
                }
                var v69 int32 = -1
                _ = v64
                _ = v69
            }
            for fi70 := 0; fi70 < 14; fi70++ {
                fmt.Println(v32, v30, fi48)
                fmt.Println(fi70)
                _ = fi70
            }
            var v71 string = "\n\t"
            _ = v71
        }
        _ = fi48
        _ = v52
    }
    if len(v31) > 0 {
        v31[0] = -744
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
