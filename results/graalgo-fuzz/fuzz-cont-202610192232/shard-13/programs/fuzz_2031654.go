package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint64; F2 int16 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint; F1 int8; F2 int16 }
func genFn0() uint8 {
    v5 := []int{636, 42}
    if (((("" + ("hello" + "foo bar")) + ("foo bar" + "a")) >= (("\n\t" + "") + "foo bar")) || false) {
        for fi6 := 0; fi6 < 8; fi6++ {
            if true {
                for fi7 := 0; fi7 < 16; fi7++ {
                    v8 := GenS1{F0: -9223372036854775808, F1: 2}
                    _ = fi7
                    _ = v8
                }
                fmt.Println(fi6)
                v9 := GenS1{F0: (fi6 | 1), F1: 42}
                _ = v9
            }
            _ = fi6
        }
        var v10 uint8 = 42
        _ = v10
    }
    _ = v5
    return 0
}
func genFn1() uint8 {
    switch "hello" {
    case "hello":
        switch 2147483647 {
        case 930:
            var v11 uint8 = genFn0()
            _ = v11
        case 42:
            v12 := []int{90, 127, 9223372036854775807, 0}
            _ = v12
        case 0:
            for fi13 := 0; fi13 < 14; fi13++ {
                if true {
                    switch "\n\t" {
                    case "a":
                        v14 := []int{fi13}
                        _ = v14
                    }
                }
                _ = fi13
            }
        }
    case "\n\t":
        var v15 uint = 18446744073709551615
        _ = v15
    default:
        var v16 uint64 = 205
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 int = -87
            if true {
                fmt.Println(fi17, fi18, fi17)
                if false {
                    var v20 float64 = ((3.14159 * (1.5 + 1e10)) / (1.5 / 1e10))
                    v19 += -1
                    _ = v20
                } else {
                    if ((!false) || false) {
                        v21 := GenS1{F0: (-(fi18)), F1: 325}
                        _ = v21
                    }
                    v22 := GenS2{F0: 705, F1: 18446744073709551615, F2: 32767}
                    v23 := GenS2{F0: 187, F1: 525, F2: 42}
                    _ = v22
                    _ = v23
                }
            } else {
                v24 := GenS3{F0: 2147483647, F1: v19}
                var v25 uint32 = 327
                switch "\n\t" {
                case "foo bar":
                    if ("a" < "foo bar") {
                        v25 -= (v25 << ((629) & 63))
                        var v26 int = v19
                        if true {
                            v19 += v26
                            switch "foo bar" {
                            case "hello":
                                v25 += v25
                            }
                        } else {
                            switch "hello" {
                            case "hello":
                                fmt.Println(fi17, fi18)
                            case "a":
                                v19 = (v26 + 128)
                            default:
                                v27 := GenS2{F0: 0, F1: 42, F2: -963}
                                _ = v27
                            }
                        }
                        if ((!false) && false) {
                            v26 = (fi18 % (((^(fi18))) | 1))
                            v19 = 0
                            v28 := []int{v26}
                            _ = v28
                        } else {
                            v19 += 4294967295
                            v19 = v19
                        }
                        _ = v26
                    } else {
                        var v29 uint64 = 42
                        v25 += uint32(v29)
                        _ = v29
                    }
                case "hello":
                    fmt.Println(v19, fi18, v25)
                default:
                    var v30 int = v19
                    _ = v30
                }
                _ = v24
                _ = v25
            }
            _ = fi18
            _ = v19
        }
        if true {
            switch "" {
            case "\n\t":
                var v31 int64 = int64(fi17)
                _ = v31
            }
            for fi32 := 0; fi32 < 15; fi32++ {
                if (32767 != 1000) {
                    fmt.Println(fi17)
                    var v33 uint64 = 1
                    v33 += (v33 + v33)
                    var v34 int16 = 0
                    _ = v33
                    _ = v34
                }
                if false {
                    fmt.Println(fi32, fi32)
                }
                _ = fi32
            }
            for fi35 := 0; fi35 < 11; fi35++ {
                fmt.Println(fi17, fi17)
                _ = fi35
            }
        }
        _ = fi17
    }
    if true {
        for fi36 := 0; fi36 < 16; fi36++ {
            var v37 string = "hello"
            _ = fi36
            _ = v37
        }
        var v38 int32 = 0
        v38 = (^(v38))
        _ = v38
    }
    var v39 float64 = 0.1
    v39 = -2.25
    _ = v39
    return genFn0()
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    var v4 int8 = 42
    var v40 float32 = (3.14159 + float32(1.5))
    _ = len(v1)
    if false {
        fmt.Println(v40, v4)
        for fi41 := 0; fi41 < 3; fi41++ {
            for fi42 := 0; fi42 < 16; fi42++ {
                switch -1 {
                case 2147483647:
                    v43 := GenS3{F0: 0, F1: fi41}
                    _ = v43
                case 1:
                    v44 := GenS3{F0: 42, F1: fi42}
                    _ = v44
                default:
                    v4 = 127
                }
                var v45 int8 = (v4 >> ((4294967295) & 63))
                v46 := []int{}
                _ = fi42
                _ = v45
                _ = v46
            }
            v40 += v40
            _ = fi41
        }
        fmt.Println(v4, v40)
    }
    for fi47 := 0; fi47 < 15; fi47++ {
        switch ("hello" + "foo bar") {
        case "foo bar":
            var v48 int = fi47
            _ = v48
        case "hello":
            v4 -= 42
        case "":
            var v49 uint = 42
            _ = v49
        default:
            delete(v2, 839)
        }
        v40 = (v40 + (((((v40 * -2.25) - v40) + (v40 / v40)) - v40) + v40))
        var v50 float64 = (1.5 / 3.14159)
        _ = fi47
        _ = v50
    }
    v4 = (v4 << ((830) & 63))
    v51 := GenS3{F0: -1, F1: 32768}
    var v52 bool = true
    if v52 {
        fmt.Println(v4, v52, v40)
        v53 := GenS3{F0: 42, F1: 42}
        var v54 bool = (-2147483648 >= 2147483647)
        _ = v53
        _ = v54
    }
    var v55 int8 = v4
    if (32767 <= -136) {
        v52 = ((!false) && v52)
        if (((v52 || v52) && ((float32(1e10) / v40) < float32(1e10))) || v52) {
            var v56 int64 = 1
            var v57 int16 = 1
            _ = v56
            _ = v57
        }
        v40 -= ((0.0 * v40) - 0.1)
    }
    if v52 {
        v58 := GenS4{F0: 255, F1: (v4 + v4), F2: -798}
        if (1 != 42) {
            var v59 uint = 282
            _ = v59
        }
        for fi60 := 0; fi60 < 20; fi60++ {
            v61 := []int{fi60, fi60}
            v55 = 100
            _ = fi60
            _ = v61
        }
        _ = v58
    } else {
        var v62 int8 = 42
        _ = v62
    }
    v40 += v40
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v51
    _ = v52
    _ = v55
}
