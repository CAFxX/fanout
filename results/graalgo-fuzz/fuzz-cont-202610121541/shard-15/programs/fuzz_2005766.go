package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 bool; F2 uint }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint }
func genFn0() uint16 {
    v7 := GenS1{F0: 17}
    var v8 int32 = 451
    var v9 bool = false
    var v10 int64 = 857
    v11 := GenS1{F0: 2}
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 1e10
    var v3 uint32 = 4294967295
    var v4 string = "foo bar"
    v5 := make(map[uint32]bool)
    var v6 bool = (109 > 162)
    for fi12 := 0; fi12 < 12; fi12++ {
        v6 = (genFn0() < 2)
        _ = fi12
    }
    if (65535 < 42) {
        switch -650 {
        case -9223372036854775808:
            var v13 string = (v4 + ("a" + ""))
            _ = v13
        case 2147483647:
            v14 := []int{660}
            _ = v14
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            switch v4 {
            case "":
                var v16 int32 = 42
                _ = v16
            case "foo bar":
                if (!(!v6)) {
                    fmt.Println(v2, v4, v4)
                    v17 := []int{(fi15 ^ (fi15 + (fi15 | (fi15 * (fi15 << ((279) & 63)))))), fi15, (^(fi15))}
                    _ = v17
                }
            case "hello":
                var v18 float64 = ((0.1 / (1e10 + float64(v2))) - ((3.14159 * (1e10 / 0.1)) * ((((3.14159 + (0.0 - 0.1)) + (1.5 / 3.14159)) * (float64((3.14159 + -2.25)) + (((3.14159 - (1e10 - float64(float32(-2.25)))) + 1e10) * 1.5))) / float64((3.14159 + v2)))))
                _ = v18
            }
            switch (fi15 >> ((4294967295) & 63)) {
            case -9223372036854775808:
                switch "hello" {
                case "a":
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v2 = v2
                        v2 += v2
                        var v20 int64 = 42
                        v21 := GenS2{F0: 32767}
                        switch (v4 + "\n\t") {
                        case "hello":
                            v6 = v6
                        case "foo bar":
                            if v6 {
                                if (!false) {
                                    v22 := GenS2{F0: 42}
                                    _ = v22
                                }
                                v23 := []int{390, (fi15 << ((434) & 63))}
                                if ((fi15 * fi15) != (fi15 - (fi15 << ((42) & 63)))) {
                                    v4 = ("foo bar" + ("a" + ((v4 + v4) + (((v4 + v4) + v4) + v4))))
                                    var v24 float64 = 3.14159
                                    _ = v24
                                }
                                _ = len(v23)
                                v25 := GenS1{F0: 932}
                                _ = v23
                                _ = v25
                            }
                        }
                        v26 := []int{(-(fi15)), (fi19 / ((fi19) | 1)), fi15, fi19}
                        var v27 uint32 = v3
                        fmt.Println(v27, v2)
                        var v28 bool = (v6 || (v6 || false))
                        v6 = (-2.25 != (1.5 + (3.14159 + v2)))
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                }
            default:
                v4 = v4
            }
            if ((v3 * 42) <= (v3 / (((v3 - (^(v3)))) | 1))) {
                switch "" {
                case "\n\t":
                    var v29 int8 = -87
                    _ = v29
                default:
                    var v30 uint8 = 42
                    _ = v30
                }
                v31 := []int{(fi15 * fi15), 32768, 0, fi15}
                if v6 {
                    var v32 uint64 = 383
                    _ = v32
                } else {
                    v6 = v6
                }
                _ = v31
            }
            _ = fi15
        }
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        v34 := GenS3{F0: (fi33 >> ((18446744073709551615) & 63)), F1: (-1 < 100), F2: 2}
        _ = fi33
        _ = v34
    }
    if (!v6) {
        v2 -= v2
        v4 = v4
        v35 := []int{0, 32768, int(v3)}
        _ = v35
    }
    switch 15 {
    case 9223372036854775807:
        v6 = (-1 >= 42)
    case 2147483647:
        v36 := []int{0, -2147483648}
        _ = v36
    }
    fmt.Println(v4)
    var v37 uint16 = 1
    if true {
        v38 := GenS3{F0: 1, F1: v6, F2: 1}
        var v39 int32 = 376
        if false {
            if (-410 >= 1000) {
                v40 := GenS4{F0: v37, F1: 0}
                _ = v40
            } else {
                var v41 int8 = 1
                v42 := GenS5{F0: 650, F1: true}
                switch (v4 + "a") {
                case "hello":
                    fmt.Println(v2, v6, v2)
                case "":
                    v43 := []int{int(v37), 32767, -1, 32768}
                    _ = v43
                case "\n\t":
                    v44 := []int{-1, -351}
                    _ = v44
                default:
                    v6 = (!(3.14159 > float64(((v2 - 3.14159) - (float32((3.14159 / 3.14159)) * v2)))))
                }
                _ = v41
                _ = v42
            }
            var v45 int8 = 127
            switch 42 {
            case -2147483648:
                var v46 bool = true
                _ = v46
            case 32768:
                fmt.Println(v2, v37, v45)
            case 1:
                if false {
                    v47 := GenS4{F0: v37, F1: 812}
                    _ = v47
                }
            default:
                v48 := []int{int(v3), 32767, 42, 1}
                _ = v48
            }
            _ = v45
        } else {
            v49 := GenS1{F0: (^(v37))}
            for fi50 := 0; fi50 < 19; fi50++ {
                v51 := GenS5{F0: 9223372036854775807, F1: false}
                v52 := GenS6{F0: v4}
                for fi53 := 0; fi53 < 13; fi53++ {
                    fmt.Println(v39, v6, fi50)
                    _ = fi53
                }
                _ = fi50
                _ = v51
                _ = v52
            }
            _ = v49
        }
        _ = v38
        _ = v39
    }
    for fi54 := 0; fi54 < 14; fi54++ {
        v55 := v5[(v3 / ((1) | 1))]
        _ = fi54
        _ = v55
    }
    v3 = (^(v3))
    _ = len(v1)
    v56 := GenS3{F0: 2147483647, F1: (v37 <= v37), F2: 474}
    v57 := GenS7{F0: 18446744073709551615}
    var v58 uint = 685
    var v59 int8 = 100
    var v60 int32 = -903
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
