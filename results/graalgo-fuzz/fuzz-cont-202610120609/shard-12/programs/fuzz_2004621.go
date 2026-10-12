package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }
type GenS2 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint; F2 int }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint8 }
func genFn0(p0 int8, p1 int16) uint16 {
    p1 += (p1 / (((p1 << ((4294967295) & 63))) | 1))
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{-1, fi7}
        _ = fi7
        _ = v8
    }
    if true {
        var v9 int = 128
        p1 = 32767
        if (false || true) {
            if false {
                var v10 uint16 = 0
                var v11 float32 = (0.1 / 1.0)
                v9 = (v9 | v9)
                _ = v10
                _ = v11
            }
        } else {
            p0 -= -79
        }
        _ = v9
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 float64 = -2.25
        v14 := GenS1{F0: 0, F1: 2147483647, F2: fi12}
        switch (fi12 >> ((688) & 63)) {
        case -9223372036854775808:
            var v15 uint8 = 42
            _ = v15
        case 1:
            p1 = p1
        default:
            switch fi12 {
            case 1:
                var v16 uint32 = 450
                _ = v16
            case 32768:
                var v17 float64 = 0.0
                _ = v17
            case 43:
                if false {
                    v13 += float64(1.5)
                    for fi18 := 0; fi18 < 9; fi18++ {
                        v13 = (((v13 + ((v13 + -2.25) / float64(1.5))) - ((v13 - (0.1 + v13)) / v13)) * (0.1 + v13))
                        _ = fi18
                    }
                } else {
                    v19 := GenS1{F0: 243, F1: 105, F2: (fi12 / (((fi12 % ((fi12) | 1))) | 1))}
                    switch (fi12 >> ((4294967295) & 63)) {
                    case 4294967295:
                        p1 = -32768
                    default:
                        for fi20 := 0; fi20 < 20; fi20++ {
                            v21 := []int{-646, (fi12 - (fi12 ^ (-(fi12))))}
                            _ = fi20
                            _ = v21
                        }
                    }
                    var v22 bool = ((-105 >= -629) && true)
                    _ = v19
                    _ = v22
                }
            }
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v23 int8 = (p0 | 1)
    _ = p0
    _ = p1
    _ = v23
    return 65535
}

func main() {
    var v1 int16 = 42
    var v2 int32 = -2147483648
    var v3 int32 = (v2 ^ (v2 & (^(v2))))
    var v4 uint32 = 4294967295
    v5 := make(map[uint32]uint32)
    var v6 uint64 = 4294967295
    v24 := v5[(v4 << ((18446744073709551615) & 63))]
    v24 += uint32(v6)
    switch (("a" + ("hello" + ("\n\t" + "a"))) + "hello") {
    case "foo bar":
        for fi25 := 0; fi25 < 9; fi25++ {
            if (((v24 & (v24 | (v4 % ((v24) | 1)))) < 2) || ((3.14159 - 1.5) > ((0.1 * 0.1) / 1.5))) {
                v26 := []int{fi25, fi25}
                _ = v26
            }
            _ = fi25
        }
    case "":
        delete(v5, (v24 % (((v4 + 812)) | 1)))
    }
    if true {
        v27 := GenS2{F0: (!false), F1: 2, F2: (v6 << ((2) & 63))}
        if false {
            v28 := []int{32767, -9223372036854775808}
            if len(v28) > 0 {
                v28[0] = 233
            }
            _ = v28
        } else {
            for fi29 := 0; fi29 < 17; fi29++ {
                v30 := GenS3{F0: "\n\t"}
                _ = fi29
                _ = v30
            }
            for fi31 := 0; fi31 < 16; fi31++ {
                v24 -= (v4 + v4)
                fmt.Println(v2, v6)
                _ = fi31
            }
        }
        _ = v27
    }
    var v32 int16 = (v1 >> ((65535) & 63))
    var v33 int8 = -59
    for fi34 := 0; fi34 < 11; fi34++ {
        v35 := GenS1{F0: 255, F1: (^(v4)), F2: (^(fi34))}
        v4 = (v4 | (v4 ^ (^(v24))))
        var v36 uint8 = 42
        _ = fi34
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        v4 = (v24 / (((v4 << ((255) & 63))) | 1))
        _ = fi37
    }
    switch "" {
    case "a":
        var v38 uint64 = 886
        _ = v38
    case "foo bar":
        for fi39 := 0; fi39 < 5; fi39++ {
            if (("foo bar" + ("hello" + ("hello" + ("\n\t" + "a")))) > ("hello" + (("foo bar" + ((("a" + "\n\t") + "a") + (("hello" + "foo bar") + "foo bar"))) + ("" + "foo bar")))) {
                var v40 float64 = (((((0.1 * ((1.5 + ((((-2.25 / (-2.25 + 0.1)) + (1e10 + float64(((1.5 * float32(1e10)) + ((3.14159 + 1e10) / (float32((float64(-2.25) - (((1e10 * ((3.14159 * (1e10 + 0.1)) * 1.5)) * ((-2.25 / (0.0 + (float64((-2.25 / 1.5)) - 0.1))) * 3.14159)) - float64(0.0)))) / (1.5 / 0.1))))))) / float64(0.1)) - 0.1)) / 1.5)) + (0.0 + (0.0 / float64(3.14159)))) + ((0.1 + 1.5) * -2.25)) / 1.0) + 3.14159)
                v41 := []int{(fi39 << ((308) & 63)), fi39, (fi39 - -1), fi39}
                v42 := GenS2{F0: ((1e10 + v40) < 0.1), F1: 32767, F2: (v6 & v6)}
                _ = v40
                _ = v41
                _ = v42
            } else {
                switch (("foo bar" + "foo bar") + "a") {
                case "\n\t":
                    var v43 int64 = 42
                    _ = v43
                default:
                    v24 += 65535
                }
                v44 := GenS4{F0: 42, F1: 729, F2: fi39}
                fmt.Println(v3)
                _ = v44
            }
            _ = fi39
        }
    }
    var v45 float32 = (((1.5 - 1.5) / 1.0) / -2.25)
    fmt.Println(v1, v4, v24)
    fmt.Println(v2)
    v45 = v45
    v46 := []int{127, int(v6)}
    for fi47 := 0; fi47 < 4; fi47++ {
        switch fi47 {
        case 32767:
            v5[v24] = v24
        case 882:
            v32 = (^(v1))
        case 423:
            if (false && (v1 == v1)) {
                v3 = v2
                if false {
                    if true {
                        v33 -= (^(v33))
                        for fi48 := 0; fi48 < 7; fi48++ {
                            v2 += -2147483648
                            _ = fi48
                        }
                        for fi49 := 0; fi49 < 2; fi49++ {
                            v50 := GenS5{F0: (^(v4)), F1: v3}
                            var v51 string = ((((("a" + ("\n\t" + ("" + "a"))) + "hello") + "") + ("\n\t" + (("" + "a") + "a"))) + "foo bar")
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                    }
                    if true {
                        if (true && false) {
                            v1 = v32
                            v6 += (v6 / (((v6 & 42)) | 1))
                            v52 := GenS1{F0: 122, F1: v24, F2: 32767}
                            _ = v52
                        }
                        v32 += v32
                        var v53 uint64 = v6
                        _ = v53
                    }
                    v4 = (^(v24))
                } else {
                    v5[(v24 ^ 2)] = (v4 >> ((42) & 63))
                    _ = len(v46)
                }
            } else {
                fmt.Println(v45, v32)
            }
        default:
            for fi54 := 0; fi54 < 10; fi54++ {
                var v55 int = fi47
                v3 = v3
                _ = fi54
                _ = v55
            }
        }
        _ = fi47
    }
    var v56 int32 = 1
    for fi57 := 0; fi57 < 3; fi57++ {
        switch "\n\t" {
        case "foo bar":
            var v58 uint32 = v24
            _ = v58
        case "\n\t":
            fmt.Println(v45)
        }
        for fi59 := 0; fi59 < 13; fi59++ {
            if false {
                v60 := GenS6{F0: 2, F1: 4294967295, F2: 167}
                v61 := GenS2{F0: ((float64(v45) * float64(float32((1.5 - (0.1 - 3.14159))))) < ((3.14159 + float64(0.1)) / ((-2.25 + 0.0) * 1.5))), F1: 65535, F2: 668}
                _ = v60
                _ = v61
            }
            var v62 uint = 4294967295
            _ = fi59
            _ = v62
        }
        _ = fi57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v56
    _ = v6
}
