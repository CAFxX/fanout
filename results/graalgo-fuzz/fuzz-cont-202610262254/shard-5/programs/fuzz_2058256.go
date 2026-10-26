package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint }
type GenS2 struct { F0 int; F1 int32; F2 int }
type GenS3 struct { F0 uint }
func genFn0() uint16 {
    var v5 int = 197
    v6 := []int{(v5 | v5), (v5 % ((-123) | 1))}
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: 42, F1: fi7, F2: 336}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    fmt.Println(v5)
    _ = v5
    _ = v6
    return 255
}

func main() {
    v1 := []int{}
    var v2 uint64 = 0
    var v3 int8 = -128
    v4 := make(map[int]string)
    fmt.Println(v3, v3)
    v2 -= (v2 + 42)
    if (v3 >= v3) {
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 uint8 = 42
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 uint8 = 161
            v14 := GenS1{F0: 0, F1: fi12, F2: 255}
            fmt.Println(v13)
            _ = fi12
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 uint8 = 42
            _ = fi15
            _ = v16
        }
    }
    switch ("hello" + "foo bar") {
    case "hello":
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
    case "\n\t":
        var v17 string = "\n\t"
        _ = v17
    default:
        v3 -= (v3 * 1)
    }
    fmt.Println(v3, v3)
    v3 = (v3 | -1)
    v2 = 2
    var v18 string = "a"
    fmt.Println(v2, v3)
    for fi19 := 0; fi19 < 5; fi19++ {
        if (!((!(0 < 1)) || (!(false || false)))) {
            if len(v1) > 0 {
                v1[0] = (fi19 & fi19)
            }
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 uint64 = v2
                if false {
                    switch fi20 {
                    case -440:
                        v2 = v2
                    }
                }
                _ = fi20
                _ = v21
            }
        } else {
            fmt.Println(v18, v3)
            var v22 int = 0
            switch "a" {
            case "foo bar":
                v2 -= (v2 & (v2 ^ 1))
            case "hello":
                fmt.Println(v2, v3, fi19)
            }
            _ = v22
        }
        _ = fi19
    }
    var v23 uint64 = v2
    var v24 int64 = -134
    if false {
        switch v18 {
        case "":
            v25 := GenS1{F0: 255, F1: 0, F2: 708}
            _ = v25
        case "a":
            v24 -= (v24 | v24)
        }
    } else {
        var v26 int16 = -1
        for fi27 := 0; fi27 < 16; fi27++ {
            v2 -= v23
            for fi28 := 0; fi28 < 10; fi28++ {
                v29 := v4[int(v24)]
                v30 := GenS2{F0: fi27, F1: 0, F2: (-(fi27))}
                _ = fi28
                _ = v29
                _ = v30
            }
            v31 := GenS3{F0: 789}
            _ = fi27
            _ = v31
        }
        fmt.Println(v18, v23)
        _ = v26
    }
    fmt.Println(v2)
    for fi32 := 0; fi32 < 11; fi32++ {
        for fi33 := 0; fi33 < 13; fi33++ {
            if ((!(!true)) || (("hello" < ("" + (("a" + ((((("\n\t" + v18) + "foo bar") + ("" + v18)) + "") + ("a" + (v18 + "hello")))) + (((((((v18 + v18) + (v18 + ("\n\t" + v18))) + "hello") + ("a" + v18)) + ((v18 + "") + "foo bar")) + (v18 + "a")) + "a")))) || (128 == 197))) {
                if ((-2147483648 < -100) && (int16(fi33) == 414)) {
                    var v34 uint32 = 2
                    _ = v34
                } else {
                    v18 = "hello"
                }
            } else {
                switch v18 {
                case "hello":
                    v18 = (v18 + "hello")
                case "a":
                    v18 = (v18 + ("\n\t" + v18))
                case "":
                    var v35 float64 = -2.25
                    _ = v35
                default:
                    v36 := v4[fi32]
                    _ = v36
                }
                v23 += 2
                var v37 uint = 406
                _ = v37
            }
            v3 = v3
            switch fi32 {
            case -1:
                fmt.Println(fi32, v23, fi32)
            case 9223372036854775807:
                var v38 uint32 = 415
                _ = v38
            case 635:
                fmt.Println(v24, v18)
            }
            fmt.Println(fi32)
            v23 -= (v2 / (((v23 & 1)) | 1))
            _ = fi33
        }
        _ = fi32
    }
    var v39 int32 = 606
    var v40 int64 = (v24 / (((v24 + (-(v24)))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}
