package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int; F2 string }
type GenS4 struct { F0 int16; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int16 }
func genFn0() uint16 {
    var v6 uint8 = 183
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint16 = 65535
        var v9 float64 = 0.1
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{-2147483648, 32767, 4294967295}
    for fi11 := 0; fi11 < 6; fi11++ {
        v6 -= (v6 / ((v6) | 1))
        if (-2.25 >= 0.0) {
            v10 = append(v10, (fi11 + (fi11 >> ((42) & 63))))
            if len(v10) > 0 {
                v10[0] = -242
            }
        } else {
            v6 -= v6
            var v12 uint16 = 2
            _ = v12
        }
        _ = fi11
    }
    switch 2147483647 {
    case -765:
        v6 = 0
    case 0:
        v6 = v6
    case -19:
        if false {
            var v13 int = 605
            _ = v13
        } else {
            v6 = 255
        }
    }
    _ = v10
    _ = v6
    return 32767
}
func genFn1(p0 int8) uint64 {
    if false {
        if (183 == 78) {
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 bool = false
                fmt.Println(v15, p0, p0)
                _ = fi14
                _ = v15
            }
            switch 128 {
            case 32767:
                if (!false) {
                    for fi16 := 0; fi16 < 19; fi16++ {
                        v17 := []int{(fi16 + fi16), (fi16 ^ fi16), -9223372036854775808}
                        _ = fi16
                        _ = v17
                    }
                }
            case 9223372036854775807:
                for fi18 := 0; fi18 < 19; fi18++ {
                    switch (fi18 | (fi18 % ((4294967295) | 1))) {
                    case -9223372036854775808:
                        var v19 int32 = -2147483648
                        _ = v19
                    case 433:
                        var v20 int = (-(fi18))
                        _ = v20
                    }
                    _ = fi18
                }
            }
            if true {
                p0 = p0
            } else {
                var v21 uint16 = 165
                _ = v21
            }
        }
        v22 := GenS1{F0: "foo bar", F1: 128, F2: 42}
        _ = v22
    }
    if (-817 == 42) {
        fmt.Println(p0)
    } else {
        v23 := GenS2{F0: 210}
        switch ("hello" + ("" + "hello")) {
        case "a":
            v24 := []int{32768, 32768, 9223372036854775807}
            _ = v24
        case "\n\t":
            for fi25 := 0; fi25 < 9; fi25++ {
                v26 := GenS3{F0: (372 == 42), F1: -144, F2: "a"}
                var v27 uint32 = 0
                _ = fi25
                _ = v26
                _ = v27
            }
        default:
            fmt.Println(p0)
        }
        _ = v23
    }
    v28 := []int{}
    v29 := GenS4{F0: 1, F1: 2147483647, F2: 65535}
    _ = p0
    _ = v28
    _ = v29
    return 794
}

func main() {
    var v1 uint16 = 1
    var v2 string = ("hello" + (("hello" + ("\n\t" + "\n\t")) + "foo bar"))
    var v3 bool = true
    var v4 bool = v3
    var v5 float64 = (1.5 / (0.0 + (0.1 + -2.25)))
    for fi30 := 0; fi30 < 20; fi30++ {
        for fi31 := 0; fi31 < 11; fi31++ {
            v32 := GenS5{F0: 128, F1: 2}
            var v33 uint16 = 65535
            _ = fi31
            _ = v32
            _ = v33
        }
        fmt.Println(v2, v5)
        var v34 uint32 = 2147483648
        _ = fi30
        _ = v34
    }
    v35 := []int{9223372036854775807}
    switch v2 {
    case "hello":
        v5 = v5
    default:
        v36 := GenS5{F0: 0, F1: 65535}
        _ = v36
    }
    var v37 uint64 = 480
    var v38 uint8 = 246
    fmt.Println(v3)
    _ = len(v35)
    v39 := GenS5{F0: 128, F1: 255}
    var v40 uint32 = 42
    v35 = append(v35, 32768)
    for fi41 := 0; fi41 < 7; fi41++ {
        v42 := GenS6{F0: 32767}
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
