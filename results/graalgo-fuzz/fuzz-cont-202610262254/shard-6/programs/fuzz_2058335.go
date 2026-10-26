package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 int16; F1 uint64; F2 string }

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]string)
    var v3 int64 = 641
    var v4 bool = (!(false && (364 < (v1 + v1))))
    v5 := make(map[int]int)
    switch 565 {
    case 70:
        for fi6 := 0; fi6 < 2; fi6++ {
            for fi7 := 0; fi7 < 10; fi7++ {
                v8 := []int{(fi7 & (fi6 | fi6)), fi7, (fi7 << ((0) & 63))}
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
    case -9223372036854775808:
        var v9 string = ("foo bar" + "\n\t")
        _ = v9
    case 32767:
        v10 := GenS1{F0: "hello"}
        _ = v10
    default:
        v11 := GenS2{F0: 0, F1: (v1 | (v1 % ((v1) | 1)))}
        _ = v11
    }
    var v12 float32 = 1.5
    v12 = v12
    for fi13 := 0; fi13 < 10; fi13++ {
        if v4 {
            switch "a" {
            case "foo bar":
                v14 := []int{}
                _ = v14
            case "a":
                fmt.Println(v1, fi13)
            default:
                v15 := GenS1{F0: ("foo bar" + "")}
                _ = v15
            }
            var v16 int16 = 632
            _ = v16
        }
        v17 := GenS3{F0: 42, F1: -128}
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 bool = (true && (0 <= -916))
            switch -9223372036854775808 {
            case -1:
                var v20 uint8 = 128
                _ = v20
            case 42:
                v21 := GenS4{F0: 65535, F1: 1}
                _ = v21
            case 32767:
                v19 = true
            }
            v22 := []int{fi18, fi18, fi18}
            _ = fi18
            _ = v19
            _ = v22
        }
        _ = fi13
        _ = v17
    }
    switch -608 {
    case 32767:
        v4 = v4
    case 127:
        switch "\n\t" {
        case "hello":
            var v23 float32 = ((v12 / -2.25) * 0.1)
            _ = v23
        default:
            v24 := GenS5{F0: 1, F1: 852}
            _ = v24
        }
    case 1:
        v25 := GenS1{F0: ("\n\t" + ("a" + ("hello" + "foo bar")))}
        _ = v25
    }
    fmt.Println(v12, v12, v1)
    v4 = v4
    v12 -= 1.5
    v26 := GenS6{F0: 0}
    switch -2147483648 {
    case -2147483648:
        var v27 uint8 = 128
        _ = v27
    case 901:
        v28 := GenS7{F0: "hello", F1: 127}
        _ = v28
    default:
        if true {
            switch 140 {
            case 4294967295:
                for fi29 := 0; fi29 < 4; fi29++ {
                    v12 = float32(float64(((v12 / v12) + float32(float64(v12)))))
                    var v30 uint64 = 0
                    _ = fi29
                    _ = v30
                }
            case 0:
                fmt.Println(v1, v3)
            case 127:
                if v4 {
                    fmt.Println(v3, v1, v12)
                    v31 := []int{}
                    v32 := GenS2{F0: 42, F1: (v1 & (^(v1)))}
                    _ = v31
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 6; fi33++ {
                        if (int16(v1) <= -1) {
                            v4 = (!(934 == 18446744073709551615))
                            v12 = v12
                        }
                        v4 = true
                        var v34 string = "foo bar"
                        _ = fi33
                        _ = v34
                    }
                    if (871 == 971) {
                        var v35 float64 = float64(0.1)
                        _ = v35
                    } else {
                        v1 = (^(v1))
                    }
                    var v36 int = 127
                    _ = v36
                }
            }
            switch -1 {
            case -628:
                v37 := []int{828, 42, int(v3)}
                _ = v37
            default:
                v38 := GenS8{F0: v3, F1: (^(v1))}
                _ = v38
            }
        }
    }
    if (!(21 >= 100)) {
        v39 := v5[32768]
        var v40 uint = 18446744073709551615
        var v41 float32 = 0.1
        _ = v39
        _ = v40
        _ = v41
    } else {
        for fi42 := 0; fi42 < 11; fi42++ {
            switch "foo bar" {
            case "a":
                switch (fi42 & 9223372036854775807) {
                case 72:
                    v1 = (v1 + 869)
                default:
                    v43 := GenS9{F0: true}
                    _ = v43
                }
            case "":
                var v44 string = ("foo bar" + ("a" + ""))
                _ = v44
            case "\n\t":
                if (!v4) {
                    var v45 int = 32767
                    _ = v45
                } else {
                    v46 := GenS10{F0: v1, F1: 4}
                    fmt.Println(v1)
                    _ = v46
                }
            }
            _ = fi42
        }
        delete(v2, -1)
    }
    var v47 int64 = 9223372036854775807
    for fi48 := 0; fi48 < 16; fi48++ {
        var v49 float32 = (3.14159 + (float32((0.1 + 1.5)) * v12))
        switch (fi48 + fi48) {
        case 4294967295:
            v5[(^(fi48))] = -9223372036854775808
        case 32768:
            v50 := GenS9{F0: v4}
            _ = v50
        case 32767:
            v51 := []int{}
            _ = v51
        default:
            switch "hello" {
            case "":
                v49 = ((float32(1.5) - 1e10) / v12)
            case "hello":
                if ((v4 && ((v47 <= 1) && false)) || (100 > 100)) {
                    v52 := v2[-9223372036854775808]
                    v49 += v49
                    _ = v52
                }
            case "\n\t":
                v53 := GenS11{F0: 42, F1: 42, F2: "\n\t"}
                _ = v53
            default:
                for fi54 := 0; fi54 < 12; fi54++ {
                    v55 := GenS11{F0: -651, F1: 1, F2: "\n\t"}
                    _ = fi54
                    _ = v55
                }
            }
        }
        _ = fi48
        _ = v49
    }
    v47 += (-(v47))
    v4 = ((2 != 4294967295) || v4)
    for fi56 := 0; fi56 < 19; fi56++ {
        for fi57 := 0; fi57 < 20; fi57++ {
            var v58 int = (fi57 % ((fi57) | 1))
            v59 := []int{127, (v58 | fi56)}
            v60 := GenS8{F0: (v3 ^ (v47 >> ((701) & 63))), F1: v1}
            _ = fi57
            _ = v58
            _ = v59
            _ = v60
        }
        _ = fi56
    }
    for fi61 := 0; fi61 < 16; fi61++ {
        v62 := GenS9{F0: v4}
        switch "" {
        case "hello":
            v63 := []int{fi61, fi61, fi61, 0}
            _ = v63
        case "a":
            var v64 float64 = 0.1
            _ = v64
        default:
            if v4 {
                v65 := []int{(fi61 >> ((125) & 63)), int(v47)}
                _ = v65
            } else {
                v3 = v3
                fmt.Println(fi61, v47)
            }
        }
        var v66 int32 = 2147483647
        _ = fi61
        _ = v62
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
