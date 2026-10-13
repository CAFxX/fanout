package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int8; F2 int32 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int32; F1 bool; F2 int32 }
type GenS5 struct { F0 int16 }
func genFn0() string {
    if (42 >= -1) {
        var v6 float32 = 0.1
        if true {
            v6 += v6
        } else {
            for fi7 := 0; fi7 < 20; fi7++ {
                v6 -= v6
                for fi8 := 0; fi8 < 20; fi8++ {
                    v9 := GenS1{F0: 1}
                    v10 := GenS2{F0: 42, F1: 0, F2: 2147483647}
                    v11 := []int{fi7, (fi7 ^ fi8), fi8, 32768}
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
                var v12 string = (("\n\t" + ("\n\t" + "hello")) + "hello")
                _ = fi7
                _ = v12
            }
            if false {
                v13 := []int{-2147483648, 4294967295, 1, 1}
                var v14 float32 = v6
                if (true || true) {
                    if (!false) {
                        v6 -= (v14 / 1e10)
                    } else {
                        var v15 int32 = 2147483647
                        var v16 int16 = -666
                        _ = v15
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 20; fi17++ {
                        v18 := GenS2{F0: -1, F1: -1, F2: 386}
                        var v19 int64 = -56
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
                _ = v13
                _ = v14
            } else {
                for fi20 := 0; fi20 < 16; fi20++ {
                    var v21 uint8 = 181
                    _ = fi20
                    _ = v21
                }
                fmt.Println(v6)
            }
            fmt.Println(v6)
        }
        fmt.Println(v6)
        _ = v6
    } else {
        if (42 > 624) {
            for fi22 := 0; fi22 < 12; fi22++ {
                v23 := []int{759, fi22}
                _ = fi22
                _ = v23
            }
        }
        var v24 bool = ("foo bar" > "")
        _ = v24
    }
    v25 := []int{0}
    switch 9223372036854775807 {
    case 32767:
        if true {
            var v26 uint8 = 255
            _ = v26
        } else {
            var v27 int16 = 1000
            fmt.Println(v27)
            _ = v27
        }
    case 373:
        v28 := GenS3{F0: (3.14159 != 1.5), F1: 32767}
        _ = v28
    case -1:
        v29 := GenS4{F0: -627, F1: false, F2: 2147483647}
        _ = v29
    }
    if (!true) {
        v30 := GenS1{F0: 718}
        for fi31 := 0; fi31 < 13; fi31++ {
            if (false || (!false)) {
                var v32 float32 = 0.1
                _ = v32
            } else {
                var v33 int8 = -57
                if false {
                    v33 = (v33 / ((v33) | 1))
                    var v34 float64 = 3.14159
                    _ = v34
                } else {
                    var v35 uint32 = 504
                    var v36 uint = 0
                    _ = v35
                    _ = v36
                }
                _ = v33
            }
            _ = fi31
        }
        switch -222 {
        case 2147483647:
            var v37 float64 = 0.0
            _ = v37
        default:
            var v38 int16 = 1
            _ = v38
        }
        _ = v30
    } else {
        _ = len(v25)
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        fmt.Println(fi39)
        for fi40 := 0; fi40 < 3; fi40++ {
            var v41 float64 = -2.25
            if (!false) {
                if true {
                    fmt.Println(fi40)
                }
                _ = len(v25)
                switch ("\n\t" + "a") {
                case "\n\t":
                    v41 = 0.0
                }
                v41 = v41
            } else {
                if true {
                    var v42 string = "foo bar"
                    _ = v42
                }
                v41 += v41
                switch fi40 {
                case 652:
                    var v43 uint8 = 144
                    var v44 int = (fi39 + (fi39 | fi40))
                    if len(v25) > 0 {
                        v25[0] = fi39
                    }
                    _ = v43
                    _ = v44
                case -550:
                    v41 = v41
                default:
                    var v45 string = (("foo bar" + "\n\t") + "foo bar")
                    _ = v45
                }
                v46 := GenS3{F0: false, F1: 42}
                _ = v46
            }
            v41 = v41
            _ = fi40
            _ = v41
        }
        for fi47 := 0; fi47 < 8; fi47++ {
            v25 = append(v25, (fi47 >> ((65535) & 63)))
            switch "" {
            case "a":
                if len(v25) > 0 {
                    v25[0] = fi47
                }
            }
            _ = fi47
        }
        _ = fi39
    }
    _ = v25
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 1
    var v3 uint32 = 42
    v4 := make(map[uint32]int)
    var v5 int16 = 1000
    var v48 uint8 = 1
    v3 = v3
    v1[2] = "hello"
    v2 -= (v2 * 127)
    var v49 int64 = -691
    switch v2 {
    case -1:
        v3 = (v3 ^ (^(v3)))
    case 9223372036854775807:
        var v50 int16 = v5
        _ = v50
    default:
        v51 := GenS1{F0: 255}
        _ = v51
    }
    v52 := GenS2{F0: (v49 / ((v49) | 1)), F1: 42, F2: 746}
    v2 += (^(v2))
    v53 := []int{}
    var v54 int64 = (-(v49))
    var v55 uint16 = 42
    if true {
        v56 := GenS5{F0: 885}
        _ = len(v4)
        var v57 uint8 = (v48 + v48)
        _ = v56
        _ = v57
    } else {
        v49 -= v49
        v58 := GenS2{F0: v54, F1: 42, F2: 42}
        _ = v58
    }
    switch ("a" + "hello") {
    case "\n\t":
        switch v2 {
        case 128:
            if true {
                _ = len(v4)
            } else {
                var v59 uint = 0
                _ = v59
            }
        }
    case "foo bar":
        switch v2 {
        case 1:
            for fi60 := 0; fi60 < 15; fi60++ {
                v2 = fi60
                var v61 int = 127
                v3 -= v3
                _ = fi60
                _ = v61
            }
        case 9223372036854775807:
            v48 = (v48 ^ v48)
        case 4294967295:
            switch (("\n\t" + "foo bar") + ((((("\n\t" + "hello") + "\n\t") + ("a" + "foo bar")) + "") + ((("" + "foo bar") + ("foo bar" + "\n\t")) + ""))) {
            case "a":
                fmt.Println(v48)
            case "hello":
                if ((613 <= 2147483647) || ((((^(v3)) >= v3) || false) && ((v2 | v2) < int(v3)))) {
                    var v62 float32 = 0.1
                    _ = v62
                }
            default:
                for fi63 := 0; fi63 < 7; fi63++ {
                    v64 := GenS2{F0: v49, F1: -83, F2: -799}
                    v65 := GenS1{F0: 4294967295}
                    _ = fi63
                    _ = v64
                    _ = v65
                }
            }
        default:
            var v66 uint8 = (v48 / ((127) | 1))
            _ = v66
        }
    case "":
        if len(v53) >= 2 {
            v53 = v53[1:len(v53)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
}
