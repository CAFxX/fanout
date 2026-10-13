package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 int16) int {
    var v7 int8 = 100
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int32 = -964
        var v10 int = -127
        switch "hello" {
        case "":
            v11 := GenS1{F0: p0, F1: false, F2: v10}
            _ = v11
        case "hello":
            v10 += (-(fi8))
        default:
            v12 := []int{-789, v10}
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v7
    return 11
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 153
    var v3 int = 32767
    var v4 float64 = (((3.14159 + (3.14159 + 0.1)) - 3.14159) + -2.25)
    var v5 string = "hello"
    var v6 bool = (!(0.0 == (0.1 * 0.0)))
    var v13 bool = (42 >= 532)
    var v14 string = (v5 + v5)
    var v15 uint16 = 1
    for fi16 := 0; fi16 < 3; fi16++ {
        v15 = v15
        _ = fi16
    }
    v17 := GenS2{F0: (v15 % ((v15) | 1))}
    if v6 {
        var v18 uint16 = v15
        fmt.Println(v13, v3, v4)
        var v19 uint8 = uint8(v2)
        _ = v18
        _ = v19
    }
    var v20 int = v3
    if v13 {
        v21 := GenS2{F0: uint16(v2)}
        v3 = genFn0(int16(v20))
        _ = v21
    } else {
        switch (v5 + "foo bar") {
        case "foo bar":
            if v13 {
                for fi22 := 0; fi22 < 18; fi22++ {
                    switch -2147483648 {
                    case -658:
                        switch (fi22 - -9223372036854775808) {
                        case 646:
                            v1[2147483648] = v20
                        case -682:
                            v23 := GenS1{F0: 32767, F1: ((float64((((-2.25 + 1.5) * (-2.25 * 0.1)) - float32(v4))) - 0.1) < 1.5), F2: v20}
                            _ = v23
                        case 42:
                            var v24 uint16 = v15
                            _ = v24
                        default:
                            v25 := GenS2{F0: v15}
                            _ = v25
                        }
                    case -1:
                        var v26 uint16 = v15
                        _ = v26
                    case -209:
                        var v27 int = genFn0(32767)
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 15; fi28++ {
                        var v29 int = fi22
                        v6 = v6
                        fmt.Println(v14)
                        _ = fi28
                        _ = v29
                    }
                    _ = fi22
                }
            }
        case "":
            var v30 uint64 = 18446744073709551615
            _ = v30
        case "hello":
            v20 += (v3 - (-(v20)))
        default:
            var v31 uint8 = 42
            _ = v31
        }
        v32 := GenS3{F0: 245, F1: -1, F2: (^(v2))}
        _ = v32
    }
    var v33 uint8 = uint8(v2)
    if (((v15 % ((v15) | 1)) < (^(v15))) || v13) {
        var v34 int64 = 42
        v35 := GenS2{F0: (v15 | (^(v15)))}
        var v36 string = (v14 + ((v5 + ((v14 + "hello") + v14)) + v5))
        _ = v34
        _ = v35
        _ = v36
    } else {
        _ = len(v1)
    }
    var v37 uint64 = 815
    v38 := GenS4{F0: false, F1: v33}
    if (2 == 190) {
        if false {
            switch 42 {
            case -2147483648:
                var v39 int = (v20 << (((v2 + 0)) & 63))
                _ = v39
            case 931:
                if true {
                    for fi40 := 0; fi40 < 6; fi40++ {
                        for fi41 := 0; fi41 < 18; fi41++ {
                            fmt.Println(v13, v33, fi41)
                            var v42 int = fi41
                            _ = fi41
                            _ = v42
                        }
                        var v43 string = ((("foo bar" + (v5 + ("" + (("foo bar" + ((v5 + v14) + v5)) + "")))) + "") + v5)
                        v44 := []int{v3}
                        _ = fi40
                        _ = v43
                        _ = v44
                    }
                } else {
                    fmt.Println(v14, v4, v14)
                    v37 = (v37 % ((4294967295) | 1))
                }
            case 360:
                v3 += (v20 & -1)
            }
        } else {
            var v45 uint8 = 0
            switch "\n\t" {
            case "\n\t":
                v46 := GenS4{F0: v6, F1: (v45 << ((844) & 63))}
                _ = v46
            default:
                v3 = -608
            }
            _ = v45
        }
    } else {
        fmt.Println(v6, v15)
        switch ("" + v5) {
        case "\n\t":
            var v47 uint64 = (v37 & 491)
            _ = v47
        case "":
            switch v3 {
            case 42:
                v14 = ((v5 + v5) + (v5 + "foo bar"))
            default:
                var v48 uint8 = (^(v33))
                _ = v48
            }
        default:
            if (v6 && false) {
                var v49 uint32 = 1
                _ = v49
            } else {
                switch (v3 | (v20 << (((v2 ^ (v2 / (((v2 ^ 4294967295)) | 1)))) & 63))) {
                case 0:
                    fmt.Println(v13)
                case 2147483647:
                    v5 = (v5 + "hello")
                }
                for fi50 := 0; fi50 < 16; fi50++ {
                    if (!v13) {
                        var v51 bool = (-951 > (fi50 + 32768))
                        v5 = (("hello" + v5) + v5)
                        _ = v51
                    }
                    v14 = "foo bar"
                    v52 := []int{}
                    _ = fi50
                    _ = v52
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
