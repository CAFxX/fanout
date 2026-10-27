package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint16; F1 int32; F2 int16 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int32; F1 string; F2 uint32 }
func genFn0() int8 {
    v7 := GenS1{F0: 998}
    v8 := GenS1{F0: -9223372036854775808}
    v9 := []int{128, 127, -13, -398}
    _ = v7
    _ = v8
    _ = v9
    return -109
}
func genFn1() uint8 {
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int16 = 42
            switch 32767 {
            case -566:
                var v13 string = ("hello" + "foo bar")
                _ = v13
            case 2147483647:
                switch "a" {
                case "hello":
                    fmt.Println(v12)
                case "foo bar":
                    v12 = (^(v12))
                default:
                    v14 := GenS2{F0: (588 == 0), F1: 42}
                    _ = v14
                }
                v15 := GenS3{F0: 65535, F1: 1, F2: v12}
                var v16 uint = 42
                _ = v15
                _ = v16
            }
            _ = fi11
            _ = v12
        }
        v17 := []int{(^(fi10))}
        var v18 int32 = -206
        _ = fi10
        _ = v17
        _ = v18
    }
    var v19 int8 = -128
    if (127 != 2) {
        v19 += v19
    }
    if false {
        for fi20 := 0; fi20 < 2; fi20++ {
            switch fi20 {
            case 32767:
                v21 := GenS3{F0: 255, F1: -2147483648, F2: 166}
                _ = v21
            case -33:
                fmt.Println(fi20, v19)
            }
            var v22 uint16 = 15
            _ = fi20
            _ = v22
        }
    }
    _ = v19
    return 128
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 255
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    var v5 int64 = 0
    var v6 float32 = (((1.5 / ((1.5 / 1.0) + ((0.1 / ((-2.25 - (0.0 * (1.5 - (1.5 + (-2.25 * 1.5))))) * (3.14159 / (1.5 * -2.25)))) / 1e10))) / 1e10) + (0.1 * (-2.25 / float32(1.5))))
    fmt.Println(v5)
    v2 -= 0
    var v23 uint32 = 1
    _ = len(v1)
    var v24 uint64 = 42
    if false {
        if (!true) {
            v25 := GenS2{F0: (284 == -1), F1: v5}
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 string = ("hello" + "foo bar")
                for fi28 := 0; fi28 < 8; fi28++ {
                    v29 := GenS4{F0: v2, F1: ("foo bar" + v27)}
                    _ = len(v4)
                    v1[("hello" + v27)] = fi26
                    _ = fi28
                    _ = v29
                }
                v30 := []int{0, (fi26 >> ((2) & 63))}
                _ = fi26
                _ = v27
                _ = v30
            }
            _ = v25
        } else {
            v3[4294967295] = true
        }
    } else {
        switch -640 {
        case 32767:
            for fi31 := 0; fi31 < 3; fi31++ {
                if ((false && ((!false) || true)) && false) {
                    var v32 int = fi31
                    var v33 uint = 881
                    var v34 string = ("a" + "foo bar")
                    _ = v32
                    _ = v33
                    _ = v34
                }
                _ = fi31
            }
        case 173:
            fmt.Println(v2, v23, v6)
        }
        v35 := []int{32767, 128, 128, -9223372036854775808}
        if true {
            for fi36 := 0; fi36 < 3; fi36++ {
                v37 := GenS5{F0: 113, F1: ("\n\t" + ((("\n\t" + ("hello" + "\n\t")) + "hello") + "")), F2: (v23 | uint32(v24))}
                _ = fi36
                _ = v37
            }
            v23 += 2147483647
        }
        _ = v35
    }
    fmt.Println(v24, v23)
    var v38 uint32 = (v23 * v23)
    var v39 uint8 = uint8(v24)
    var v40 int32 = 176
    fmt.Println(v40, v23, v5)
    switch 32768 {
    case -332:
        v41 := v4[-9223372036854775808]
        _ = v41
    default:
        if true {
            var v42 int8 = 42
            for fi43 := 0; fi43 < 10; fi43++ {
                v42 += (v42 % (((v42 | (v42 & genFn0()))) | 1))
                v2 = (v2 % (((v2 << ((uint(v2)) & 63))) | 1))
                for fi44 := 0; fi44 < 18; fi44++ {
                    v45 := []int{fi44}
                    _ = fi44
                    _ = v45
                }
                _ = fi43
            }
            switch 4294967295 {
            case 1:
                var v46 uint64 = (v24 >> ((4294967295) & 63))
                _ = v46
            default:
                for fi47 := 0; fi47 < 15; fi47++ {
                    var v48 int64 = int64(v40)
                    _ = fi47
                    _ = v48
                }
            }
            _ = v42
        }
    }
    for fi49 := 0; fi49 < 15; fi49++ {
        var v50 uint = 2
        _ = fi49
        _ = v50
    }
    v24 = 859
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
