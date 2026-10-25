package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint32; F2 bool }
type GenS5 struct { F0 uint; F1 uint64; F2 uint32 }
func genFn0(p0 int64) int32 {
    var v6 uint64 = 42
    var v7 int8 = 42
    var v8 uint32 = 65535
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := []int{fi9, fi9, 32768, (fi9 >> ((18446744073709551615) & 63))}
        _ = fi9
        _ = v10
    }
    if (!true) {
        var v11 uint64 = (v6 - 4294967295)
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := GenS1{F0: 42, F1: (1 > 0)}
            for fi14 := 0; fi14 < 2; fi14++ {
                v15 := []int{}
                v16 := GenS1{F0: -405, F1: false}
                v8 += 42
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1() uint8 {
    var v17 uint8 = 127
    var v18 int = 1
    v19 := GenS2{F0: true}
    for fi20 := 0; fi20 < 9; fi20++ {
        if (-147 != 0) {
            v21 := GenS3{F0: (v17 << ((65535) & 63))}
            _ = v21
        }
        _ = fi20
    }
    _ = v17
    _ = v18
    _ = v19
    return (v17 << ((409) & 63))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (((!false) || true) && (!(2 != 0)))
    var v3 int = 1
    v4 := []string{"foo bar"}
    var v5 float64 = (-2.25 / 1.5)
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    fmt.Println(v2, v3, v2)
    fmt.Println(v2, v2)
    if true {
        if v2 {
            if (65535 <= 1) {
                v22 := v1[65535]
                if len(v4) > 0 {
                    v4[0] = "a"
                }
                _ = v22
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch "" {
                case "foo bar":
                    v23 := GenS4{F0: 241, F1: 863, F2: false}
                    _ = v23
                }
            }
        } else {
            switch -884 {
            case 4294967295:
                var v24 uint32 = 42
                _ = v24
            case 454:
                v2 = v2
            }
            fmt.Println(v3, v2, v3)
        }
    }
    v5 -= ((v5 + 1.5) - 1e10)
    _ = len(v1)
    v25 := GenS1{F0: 42, F1: true}
    if (0 < 32767) {
        var v26 int8 = -67
        delete(v1, 594)
        if false {
            fmt.Println(v5)
            v3 = -9223372036854775808
            var v27 uint8 = 255
            _ = v27
        }
        _ = v26
    } else {
        fmt.Println(v3, v5)
        if v2 {
            var v28 bool = true
            v29 := GenS2{F0: ((255 < 1) || (true && false))}
            _ = v28
            _ = v29
        }
    }
    v30 := GenS2{F0: false}
    if len(v4) > 0 {
        v4[0] = ""
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        fmt.Println(v2)
        v32 := GenS2{F0: v2}
        _ = fi31
        _ = v32
    }
    v2 = (!(!(42 > 1)))
    v33 := v1[2]
    var v34 uint32 = 65535
    var v35 uint8 = 128
    if ("\n\t" != "") {
        v34 -= 65535
    } else {
        v2 = (false && (!(100 <= -128)))
        v34 -= 728
        fmt.Println(v35, v5)
    }
    if (1 > 42) {
        v35 += (v35 << ((0) & 63))
    } else {
        var v36 int16 = -806
        var v37 string = (("" + ("a" + "hello")) + (("\n\t" + ("a" + ("a" + ""))) + "foo bar"))
        v38 := GenS5{F0: 4294967295, F1: 18446744073709551615, F2: (^(v34))}
        _ = v36
        _ = v37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
