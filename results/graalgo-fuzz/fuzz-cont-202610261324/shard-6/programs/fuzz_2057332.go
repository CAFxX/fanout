package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 string; F2 int }
type GenS5 struct { F0 int32; F1 int64; F2 int32 }

func main() {
    var v1 bool = false
    var v2 bool = v1
    var v3 int = 42
    var v4 int8 = 127
    var v5 uint = 167
    var v6 bool = ((2 > 200) && true)
    var v7 bool = v6
    v4 += 27
    for fi8 := 0; fi8 < 20; fi8++ {
        if (1 == 42) {
            v9 := []int{fi8, (fi8 * (v3 << (((v5 ^ (^(v5)))) & 63))), (^(v3))}
            if (v5 < v5) {
                fmt.Println(v7)
                _ = len(v9)
            }
            _ = v9
        } else {
            var v10 int32 = -2147483648
            _ = v10
        }
        switch 4294967295 {
        case -9223372036854775808:
            var v11 uint8 = 128
            _ = v11
        case -2147483648:
            v12 := []int{fi8, (fi8 - v3), v3, 42}
            _ = v12
        case -910:
            v13 := []int{9223372036854775807, fi8, (fi8 * (fi8 | v3)), (fi8 | fi8)}
            _ = v13
        default:
            if true {
                var v14 uint32 = 87
                switch ((("foo bar" + "hello") + ("\n\t" + "")) + ("hello" + ("hello" + ""))) {
                case "hello":
                    if (("hello" + "\n\t") <= (("foo bar" + ("\n\t" + "\n\t")) + "a")) {
                        v15 := []int{-9223372036854775808, fi8, int(v14), v3}
                        _ = v15
                    }
                default:
                    for fi16 := 0; fi16 < 6; fi16++ {
                        fmt.Println(fi16, v2)
                        _ = fi16
                    }
                }
                v6 = false
                _ = v14
            } else {
                v1 = (65535 < 65535)
            }
        }
        _ = fi8
    }
    v7 = (!v6)
    switch "\n\t" {
    case "\n\t":
        v17 := []int{-219, v3, v3}
        _ = v17
    }
    v1 = v1
    if (127 < 250) {
        v18 := GenS1{F0: 529, F1: true}
        if v6 {
            for fi19 := 0; fi19 < 5; fi19++ {
                if v6 {
                    fmt.Println(fi19, v6, v3)
                }
                fmt.Println(v6)
                v3 -= (fi19 & (v3 & -1))
                _ = fi19
            }
            for fi20 := 0; fi20 < 4; fi20++ {
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 uint = 0
                    _ = fi21
                    _ = v22
                }
                switch 1 {
                case 4294967295:
                    if v6 {
                        switch fi20 {
                        case 2147483647:
                            v23 := []int{fi20, fi20, -233, (fi20 + 12)}
                            _ = v23
                        case -136:
                            v24 := []int{(v3 | v3), -1, (v3 ^ (fi20 << ((v5) & 63))), v3}
                            _ = v24
                        case 0:
                            v2 = v2
                        default:
                            fmt.Println(v4)
                        }
                        for fi25 := 0; fi25 < 3; fi25++ {
                            v1 = false
                            v6 = (1.5 > (1e10 + ((float64(0.0) + 0.0) - (1.5 - float64((0.1 + -2.25))))))
                            v26 := GenS2{F0: 1000, F1: 2147483647}
                            _ = fi25
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 3; fi27++ {
                            var v28 bool = (int32(v5) <= 0)
                            if v1 {
                                var v29 uint = (v5 / (((v5 | v5)) | 1))
                                v30 := []int{(v3 + -9223372036854775808), fi27, (-(fi27)), (fi20 % (((fi20 / (((fi20 % (((v3 & (fi27 * (fi20 + fi27)))) | 1))) | 1))) | 1))}
                                if true {
                                    v2 = (fi27 <= fi20)
                                    v7 = (0.1 <= float32(-2.25))
                                    var v31 int = fi20
                                    fmt.Println(v28, v5, v31)
                                    _ = v31
                                } else {
                                    v32 := GenS1{F0: 0, F1: v7}
                                    _ = v32
                                }
                                _ = v29
                                _ = v30
                            }
                            _ = fi27
                            _ = v28
                        }
                    }
                case -147:
                    var v33 uint16 = 143
                    _ = v33
                }
                _ = fi20
            }
            for fi34 := 0; fi34 < 8; fi34++ {
                v6 = v6
                v35 := GenS3{F0: (v4 & (v4 & 1))}
                _ = fi34
                _ = v35
            }
        }
        if (350 != 255) {
            var v36 int32 = -1
            switch v3 {
            case 854:
                v1 = (255 < 1)
            }
            _ = v36
        } else {
            v7 = (float64(3.14159) < ((3.14159 - -2.25) + float64(1e10)))
            v37 := GenS3{F0: (v4 & v4)}
            fmt.Println(v3)
            _ = v37
        }
        _ = v18
    } else {
        v38 := GenS4{F0: v1, F1: "foo bar", F2: v3}
        fmt.Println(v5, v5)
        _ = v38
    }
    v1 = (v4 == (-(v4)))
    v5 -= v5
    if (32767 > -227) {
        v39 := GenS5{F0: 527, F1: 294, F2: -1}
        switch ("a" + "hello") {
        case "foo bar":
            v40 := GenS2{F0: 1000, F1: -1}
            _ = v40
        }
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
