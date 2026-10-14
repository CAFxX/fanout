package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 uint32; F2 int8 }
func genFn0() int {
    var v6 int = -9223372036854775808
    fmt.Println(v6)
    v7 := []int{127, v6, (v6 | 32767)}
    var v8 int8 = 1
    _ = v6
    _ = v7
    _ = v8
    return (-(v6))
}
func genFn1(p0 bool, p1 int32, p2 int32) int32 {
    p1 = (-(p1))
    if true {
        p1 = (p1 << ((18446744073709551615) & 63))
        p0 = false
        p0 = (0.0 != ((0.0 * 0.1) + 1e10))
    } else {
        p1 -= -963
        v9 := []int{1, 9223372036854775807, 127, -1}
        p2 = (p2 >> ((4294967295) & 63))
        _ = v9
    }
    if (!((517 == 128) && false)) {
        p2 -= p2
        if false {
            v10 := []int{}
            _ = v10
        } else {
            var v11 float32 = ((((((1e10 + 0.0) - 3.14159) + ((0.1 / (0.1 * ((1e10 - ((3.14159 - (-2.25 * ((1.5 * -2.25) * 0.0))) * 0.0)) * 0.1))) - 1e10)) / (1e10 + (0.0 + 1.5))) * 3.14159) + float32((float64(1.5) / -2.25)))
            p2 = (p2 >> ((2) & 63))
            _ = v11
        }
    } else {
        var v12 uint = 799
        _ = v12
    }
    if p0 {
        v13 := GenS1{F0: 1}
        p1 += p2
        _ = v13
    }
    v14 := GenS2{F0: 2147483647, F1: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return p1
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483648, 4294967295}
    var v3 float64 = 1.5
    var v4 string = "hello"
    v5 := make(map[int]bool)
    fmt.Println(v3, v3)
    if (-746 <= 9223372036854775807) {
        if (!(true || (42 < 1))) {
            fmt.Println(v3)
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v16 int8 = 127
            _ = fi15
            _ = v16
        }
        if false {
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 uint16 = 32767
                v19 := GenS1{F0: (fi17 / (((^(fi17))) | 1))}
                fmt.Println(fi17)
                _ = fi17
                _ = v18
                _ = v19
            }
            v4 = v4
            _ = len(v2)
        } else {
            if true {
                v1[9223372036854775807] = genFn0()
                var v20 int64 = -851
                var v21 int16 = 5
                _ = v20
                _ = v21
            }
            v22 := GenS1{F0: 91}
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        v3 -= v3
        for fi24 := 0; fi24 < 19; fi24++ {
            if (98 == 65535) {
                fmt.Println(fi23)
                switch 32768 {
                case 599:
                    var v25 uint = 255
                    _ = v25
                case -9223372036854775808:
                    v3 = v3
                case -878:
                    var v26 int8 = -128
                    _ = v26
                default:
                    _ = len(v2)
                }
                fmt.Println(fi24, fi23)
                var v27 int64 = 1
                switch (v4 + v4) {
                case "hello":
                    v28 := GenS1{F0: 9223372036854775807}
                    _ = v28
                case "a":
                    var v29 int16 = 1
                    _ = v29
                case "":
                    v5[(fi23 ^ genFn0())] = true
                default:
                    var v30 uint8 = 218
                    _ = v30
                }
                _ = v27
            } else {
                var v31 uint32 = 1
                _ = v31
            }
            var v32 int16 = 0
            switch (-(fi24)) {
            case 4294967295:
                switch -1 {
                case 32767:
                    _ = len(v2)
                case -819:
                    if len(v2) > 0 {
                        v2[0] = 590
                    }
                case -157:
                    var v33 uint8 = 128
                    switch 982 {
                    case -1:
                        var v34 int = fi24
                        _ = v34
                    default:
                        var v35 float64 = v3
                        _ = v35
                    }
                    v36 := GenS3{F0: 86}
                    _ = v33
                    _ = v36
                }
            default:
                switch v4 {
                case "\n\t":
                    var v37 int8 = 127
                    _ = v37
                case "":
                    v38 := GenS1{F0: (fi23 + fi24)}
                    _ = v38
                case "foo bar":
                    v3 = v3
                }
            }
            _ = fi24
            _ = v32
        }
        if (!true) {
            v3 -= (1e10 / 1.0)
            switch "" {
            case "a":
                var v39 uint8 = 2
                _ = v39
            case "":
                var v40 uint16 = 65535
                _ = v40
            case "hello":
                v41 := GenS1{F0: 0}
                _ = v41
            default:
                if true {
                    if (-122 != 0) {
                        v4 = (((v4 + ((("\n\t" + v4) + "a") + (v4 + ""))) + (v4 + "\n\t")) + ("" + v4))
                    }
                    v4 = ("foo bar" + v4)
                    v3 = 1.5
                }
            }
        }
        _ = fi23
    }
    var v42 uint = 42
    v4 = (v4 + v4)
    var v43 int32 = -1
    if false {
        switch v4 {
        case "foo bar":
            for fi44 := 0; fi44 < 2; fi44++ {
                v45 := GenS3{F0: 42}
                if false {
                    fmt.Println(fi44, v4, v4)
                } else {
                    var v46 uint8 = 40
                    _ = v46
                }
                _ = fi44
                _ = v45
            }
        case "\n\t":
            if (835 > v43) {
                var v47 float64 = ((v3 + ((v3 / v3) + -2.25)) + v3)
                _ = v47
            } else {
                var v48 int32 = v43
                fmt.Println(v4, v43, v48)
                for fi49 := 0; fi49 < 19; fi49++ {
                    v50 := GenS4{F0: (fi49 ^ 42)}
                    _ = fi49
                    _ = v50
                }
                _ = v48
            }
        default:
            fmt.Println(v42)
        }
        var v51 int64 = 9223372036854775807
        fmt.Println(v43)
        _ = v51
    } else {
        for fi52 := 0; fi52 < 3; fi52++ {
            _ = len(v2)
            for fi53 := 0; fi53 < 13; fi53++ {
                var v54 bool = (42 > 642)
                _ = fi53
                _ = v54
            }
            var v55 uint8 = 37
            _ = fi52
            _ = v55
        }
    }
    if (42 >= 42) {
        var v56 int32 = (v43 * -2147483648)
        _ = v56
    } else {
        v57 := GenS5{F0: 120, F1: 2147483648, F2: -1}
        _ = v57
    }
    v4 = v4
    var v58 uint = (v42 - v42)
    var v59 bool = (65535 == 951)
    var v60 uint32 = 777
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v58
    _ = v59
    _ = v60
}
