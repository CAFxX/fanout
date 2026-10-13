package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 bool }
type GenS3 struct { F0 uint; F1 uint64; F2 int }
type GenS4 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS5 struct { F0 string; F1 uint32; F2 int32 }
func genFn0(p0 uint32, p1 string, p2 string) uint8 {
    var v5 int = 127
    switch (v5 >> ((456) & 63)) {
    case -1:
        fmt.Println(p0, v5)
    case 944:
        v6 := GenS1{F0: 1}
        _ = v6
    case 1:
        for fi7 := 0; fi7 < 18; fi7++ {
            for fi8 := 0; fi8 < 9; fi8++ {
                v9 := []int{(fi7 % ((fi7) | 1))}
                if (1 == 32768) {
                    v9 = append(v9, (fi8 - fi8))
                    var v10 uint32 = p0
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
            for fi11 := 0; fi11 < 10; fi11++ {
                switch (fi7 ^ (-(fi11))) {
                case 781:
                    if ((true && false) || false) {
                        v12 := []int{}
                        _ = v12
                    } else {
                        var v13 int = fi11
                        p1 = ("a" + "foo bar")
                        p2 = "a"
                        fmt.Println(fi7)
                        var v14 int8 = -1
                        _ = v13
                        _ = v14
                    }
                case 42:
                    var v15 int = 468
                    _ = v15
                case 1:
                    var v16 int64 = -452
                    _ = v16
                default:
                    if false {
                        v17 := GenS2{F0: 263, F1: 179, F2: (false && (!false))}
                        var v18 float32 = ((1.5 * 1.5) / (-2.25 + -2.25))
                        v19 := []int{}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
                v20 := GenS1{F0: 0}
                _ = fi11
                _ = v20
            }
            _ = fi7
        }
    }
    p1 = p2
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := GenS1{F0: 100}
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 130
}
func genFn1() int16 {
    var v23 uint16 = 458
    var v24 bool = (127 <= 1)
    _ = v23
    _ = v24
    return -579
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint16 = 0
    v3 := make(map[int]uint32)
    var v4 uint64 = 42
    v2 += (v2 ^ v2)
    v1 -= (v1 / (((v1 | (v1 % (((v1 | v1)) | 1)))) | 1))
    for fi25 := 0; fi25 < 4; fi25++ {
        v26 := GenS1{F0: 100}
        if (0 >= 1) {
            var v27 float32 = 0.1
            if false {
                v28 := GenS1{F0: -3}
                _ = v28
            }
            _ = v27
        }
        v1 -= (v1 & v1)
        _ = fi25
        _ = v26
    }
    v29 := GenS3{F0: v1, F1: 330, F2: -363}
    for fi30 := 0; fi30 < 20; fi30++ {
        if true {
            for fi31 := 0; fi31 < 17; fi31++ {
                v4 = (v4 & v4)
                fmt.Println(v2, fi31)
                _ = fi31
            }
        } else {
            v32 := GenS3{F0: v1, F1: (v4 & (v4 | (^(v4)))), F2: fi30}
            v33 := GenS4{F0: v1, F1: 18446744073709551615, F2: 255}
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    v2 += v2
    var v34 float64 = 0.0
    _ = len(v3)
    if (!(!(-1 >= 227))) {
        v2 = (^(v2))
    } else {
        switch 32767 {
        case -49:
            if (true || true) {
                var v35 uint = 1
                _ = v35
            }
        case 2147483647:
            v36 := []int{}
            _ = v36
        case -1:
            var v37 uint64 = 2
            _ = v37
        default:
            v34 += v34
        }
        v34 += v34
        var v38 uint = 318
        _ = v38
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        var v40 uint8 = 42
        v4 = (v4 - 537)
        for fi41 := 0; fi41 < 4; fi41++ {
            v42 := GenS3{F0: (v1 + (^(v1))), F1: v4, F2: fi39}
            var v43 int16 = genFn1()
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi39
        _ = v40
    }
    if true {
        v34 += float64(0.1)
    } else {
        if true {
            switch 0 {
            case 9223372036854775807:
                v4 += v4
            case -9223372036854775808:
                v44 := GenS3{F0: (v1 + (v1 / (((v1 >> ((v1) & 63))) | 1))), F1: (v4 % (((^(v4))) | 1)), F2: 42}
                _ = v44
            case 182:
                var v45 float64 = 1e10
                _ = v45
            default:
                for fi46 := 0; fi46 < 3; fi46++ {
                    v1 = v1
                    _ = fi46
                }
            }
            for fi47 := 0; fi47 < 19; fi47++ {
                var v48 uint8 = 128
                if (255 < uint(v4)) {
                    v34 = float64(1.5)
                    var v49 uint = v1
                    delete(v3, (^(fi47)))
                    _ = v49
                }
                for fi50 := 0; fi50 < 8; fi50++ {
                    v51 := v3[(-(fi50))]
                    _ = fi50
                    _ = v51
                }
                _ = fi47
                _ = v48
            }
            v52 := GenS4{F0: uint(v4), F1: (^(v4)), F2: (v2 - v2)}
            _ = v52
        } else {
            switch "foo bar" {
            case "":
                fmt.Println(v34)
            default:
                v53 := GenS1{F0: -128}
                _ = v53
            }
            fmt.Println(v4, v34, v4)
            var v54 bool = false
            _ = v54
        }
        fmt.Println(v2)
    }
    if true {
        switch 9223372036854775807 {
        case 42:
            var v55 int8 = 1
            _ = v55
        }
        v56 := []int{-775}
        if ((v4 << ((2) & 63)) > v4) {
            v2 = 65535
            var v57 int32 = -1
            var v58 int = 586
            _ = v57
            _ = v58
        } else {
            if false {
                for fi59 := 0; fi59 < 2; fi59++ {
                    fmt.Println(v1)
                    v60 := GenS3{F0: v1, F1: (v4 << ((v1) & 63)), F2: fi59}
                    _ = fi59
                    _ = v60
                }
            }
            v34 = v34
        }
        _ = v56
    } else {
        if (!false) {
            if ("\n\t" >= "foo bar") {
                var v61 string = (("" + "") + (("" + "a") + ""))
                var v62 int64 = 1
                _ = v61
                _ = v62
            } else {
                var v63 string = "hello"
                v64 := []int{-9223372036854775808}
                _ = v63
                _ = v64
            }
            v1 = v1
        }
        var v65 bool = ((214 == 2) && true)
        _ = v65
    }
    var v66 float64 = v34
    var v67 uint32 = 2
    for fi68 := 0; fi68 < 1; fi68++ {
        var v69 int8 = 42
        for fi70 := 0; fi70 < 13; fi70++ {
            fmt.Println(fi68, v4, v2)
            _ = fi70
        }
        v1 += (v1 * (v1 % (((v1 + (v1 % (((v1 << ((v1) & 63))) | 1)))) | 1)))
        _ = fi68
        _ = v69
    }
    v71 := GenS5{F0: "\n\t", F1: v67, F2: 303}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v66
    _ = v67
    _ = v71
}
