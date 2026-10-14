package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int8) uint16 {
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 uint = 578
            var v10 uint32 = 1
            if (!(1 < 0)) {
                p0 = (p0 >> (((v9 * (v9 * v9))) & 63))
                var v11 uint64 = 72
                _ = v11
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            p0 -= 0
            p0 = 100
            fmt.Println(fi12, fi12, fi12)
            _ = fi12
        }
        _ = fi7
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        if true {
            for fi14 := 0; fi14 < 2; fi14++ {
                v15 := []int{fi13, fi13}
                _ = fi14
                _ = v15
            }
            switch fi13 {
            case 128:
                v16 := GenS1{F0: 816, F1: (^(fi13))}
                _ = v16
            case 9223372036854775807:
                var v17 int32 = -2147483648
                _ = v17
            }
        } else {
            if (true || true) {
                var v18 float64 = -2.25
                var v19 float64 = (float64(1.5) - float64((-2.25 - 0.1)))
                _ = v18
                _ = v19
            } else {
                v20 := GenS1{F0: 2, F1: (-(fi13))}
                v21 := GenS1{F0: 1, F1: fi13}
                _ = v20
                _ = v21
            }
            var v22 bool = ((fi13 - (fi13 * (fi13 + fi13))) < -2147483648)
            _ = v22
        }
        _ = fi13
    }
    v23 := []int{4294967295, 9223372036854775807, 4294967295, 9223372036854775807}
    p0 = 0
    var v24 uint = 42
    _ = p0
    _ = v23
    _ = v24
    return 255
}

func main() {
    var v1 int32 = 0
    var v2 uint8 = 193
    var v3 float64 = float64(-2.25)
    var v4 int = 2147483647
    var v5 uint64 = 18446744073709551615
    var v6 uint16 = 285
    v25 := GenS1{F0: v6, F1: (v4 / ((v4) | 1))}
    v3 = v3
    v6 = v6
    if false {
        v1 = (v1 ^ 42)
        var v26 int32 = 1
        _ = v26
    } else {
        v27 := []int{(v4 + (^(v4))), -794, 929}
        v2 -= 127
        v28 := GenS1{F0: (v6 / ((0) | 1)), F1: (v4 % (((v4 >> ((2) & 63))) | 1))}
        _ = v27
        _ = v28
    }
    var v29 int16 = -32768
    var v30 int32 = -2147483648
    if ((-(v29)) >= v29) {
        v31 := GenS1{F0: (v6 >> ((65535) & 63)), F1: v4}
        var v32 string = ""
        _ = v31
        _ = v32
    } else {
        var v33 uint8 = (v2 ^ v2)
        switch "a" {
        case "\n\t":
            if true {
                var v34 int64 = 0
                fmt.Println(v4, v4, v34)
                _ = v34
            } else {
                var v35 uint8 = v33
                v4 = v4
                v36 := []int{v4, (v4 >> ((0) & 63)), v4, 32767}
                _ = v35
                _ = v36
            }
        case "hello":
            v4 = 32768
        case "":
            if ((v4 | v4) >= (v4 % ((v4) | 1))) {
                for fi37 := 0; fi37 < 16; fi37++ {
                    fmt.Println(v33)
                    switch (v4 << ((433) & 63)) {
                    case 0:
                        for fi38 := 0; fi38 < 8; fi38++ {
                            v4 += v4
                            v1 += v1
                            _ = fi38
                        }
                    case 599:
                        switch ("a" + ("\n\t" + ("hello" + (("a" + "hello") + "hello")))) {
                        case "a":
                            var v39 uint = 552
                            _ = v39
                        case "hello":
                            v33 = (v2 << ((65535) & 63))
                        }
                    default:
                        v40 := []int{v4, v4}
                        _ = v40
                    }
                    _ = fi37
                }
            }
        }
        _ = v33
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        fmt.Println(v2, v1, v6)
        for fi42 := 0; fi42 < 13; fi42++ {
            if true {
                var v43 bool = true
                if false {
                    var v44 bool = v43
                    v45 := GenS1{F0: 42, F1: fi41}
                    fmt.Println(v1, v5, v43)
                    _ = v44
                    _ = v45
                }
                if (v43 || ("\n\t" <= "a")) {
                    v46 := []int{}
                    if (!(!(!(!((!((((float64(((((0.0 + 0.0) - ((float32(v3) * float32(3.14159)) / 1.0)) / 1.5) * 1e10)) / -2.25) * v3) == (((v3 + v3) * v3) / 0.1)) || ((v29 ^ (v29 / ((1) | 1))) == v29))) || false))))) {
                        var v47 string = ("foo bar" + "hello")
                        v43 = ((v2 / ((uint8(v5)) | 1)) <= (^(v2)))
                        _ = v47
                    }
                    if v43 {
                        var v48 int64 = -531
                        v1 = 42
                        v46 = append(v46, (fi42 / ((v4) | 1)))
                        _ = v48
                    } else {
                        v49 := GenS1{F0: v6, F1: (-(fi41))}
                        v46 = append(v46, fi42)
                        if false {
                            if v43 {
                                v46 = append(v46, 1)
                                v50 := GenS1{F0: v6, F1: v4}
                                _ = v50
                            }
                            var v51 uint64 = v5
                            var v52 int32 = (-(v30))
                            _ = v51
                            _ = v52
                        } else {
                            v30 += v30
                        }
                        _ = v49
                    }
                    fmt.Println(v5, v6, v1)
                    _ = v46
                } else {
                    v29 += v29
                    v30 += v30
                }
                v4 -= (-(fi42))
                switch 32768 {
                case -854:
                    var v53 uint16 = genFn0(127)
                    _ = v53
                case 127:
                    v54 := GenS2{F0: v30, F1: v5}
                    _ = v54
                case -241:
                    if (v43 && true) {
                        if v43 {
                            var v55 uint = 255
                            _ = v55
                        }
                        var v56 int32 = v30
                        var v57 uint = 18446744073709551615
                        _ = v56
                        _ = v57
                    }
                    if (!(v43 && v43)) {
                        v3 = 0.0
                        var v58 uint8 = uint8(v5)
                        v59 := GenS3{F0: -1}
                        _ = v58
                        _ = v59
                    } else {
                        var v60 int16 = v29
                        var v61 int16 = (^(v60))
                        _ = v60
                        _ = v61
                    }
                }
                var v62 int64 = -280
                _ = v43
                _ = v62
            } else {
                var v63 string = ""
                _ = v63
            }
            v64 := []int{fi41}
            if len(v64) >= 2 {
                v64 = v64[1:len(v64)]
            }
            _ = fi42
            _ = v64
        }
        var v65 int = (v4 ^ (^(v4)))
        _ = fi41
        _ = v65
    }
    v5 += v5
    for fi66 := 0; fi66 < 1; fi66++ {
        switch "" {
        case "a":
            var v67 float64 = 1e10
            _ = v67
        }
        _ = fi66
    }
    var v68 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v68
}
