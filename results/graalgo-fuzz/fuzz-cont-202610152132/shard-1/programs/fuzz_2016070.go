package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int64; F2 int8 }
func genFn0() int16 {
    if true {
        v6 := []int{127, 9223372036854775807, 128, -2147483648}
        if true {
            if len(v6) > 0 {
                v6[0] = 42
            }
            var v7 int64 = 9223372036854775807
            _ = v7
        } else {
            if true {
                var v8 float32 = 1.5
                _ = v8
            }
            var v9 int = 9223372036854775807
            v9 -= v9
            _ = v9
        }
        _ = v6
    }
    if false {
        var v10 float64 = (0.1 + 3.14159)
        var v11 bool = true
        _ = v10
        _ = v11
    } else {
        if (!(905 > -32768)) {
            var v12 uint8 = 0
            var v13 string = ("\n\t" + "hello")
            _ = v12
            _ = v13
        }
        var v14 float32 = 1.5
        var v15 int = 128
        _ = v14
        _ = v15
    }
    var v16 int16 = -32768
    var v17 int64 = -97
    if (3.14159 == 1.5) {
        v17 -= -659
        var v18 bool = (42 == 4294967295)
        _ = v18
    }
    _ = v16
    _ = v17
    return (^(v16))
}
func genFn1() bool {
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 int64 = -9223372036854775808
        v21 := []int{(fi19 & -219), 0, (-(fi19)), -2147483648}
        v21 = append(v21, (fi19 & (fi19 % (((fi19 * (fi19 + fi19))) | 1))))
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 float64 = 0.0
        v23 = (v23 + (v23 + (1e10 - -2.25)))
        if true {
            for fi24 := 0; fi24 < 19; fi24++ {
                v25 := []int{(fi22 & fi24), (-(fi24))}
                v26 := GenS1{F0: -534, F1: 2}
                _ = fi24
                _ = v25
                _ = v26
            }
            v23 += float64(1.5)
        }
        _ = fi22
        _ = v23
    }
    return true
}

func main() {
    var v1 bool = (-1 == -9223372036854775808)
    var v2 string = ((("hello" + "\n\t") + ("a" + (("a" + "a") + ("" + "")))) + "\n\t")
    var v3 int32 = -258
    var v4 uint16 = 1
    var v5 int16 = -1
    v27 := []int{}
    switch -697 {
    case -135:
        if (genFn0() <= v5) {
            switch 261 {
            case 42:
                var v28 int32 = -1
                _ = v28
            case 196:
                v29 := GenS1{F0: int16(v4), F1: 738}
                _ = v29
            case 127:
                if len(v27) > 0 {
                    v27[0] = 4294967295
                }
            default:
                v30 := GenS1{F0: genFn0(), F1: 944}
                _ = v30
            }
        }
    default:
        switch -795 {
        case -577:
            v31 := GenS1{F0: 522, F1: 0}
            _ = v31
        case -686:
            v3 -= int32(v4)
        default:
            v2 = "foo bar"
        }
    }
    var v32 uint64 = 4294967295
    if (-19 >= 28) {
        v33 := GenS2{F0: 2, F1: 131, F2: -1}
        var v34 float32 = -2.25
        _ = v33
        _ = v34
    }
    switch ((v2 + v2) + v2) {
    case "a":
        if len(v27) >= 2 {
            v27 = v27[1:len(v27)]
        }
    case "":
        for fi35 := 0; fi35 < 10; fi35++ {
            if v1 {
                if len(v27) > 0 {
                    v27[0] = (fi35 ^ int(v32))
                }
                var v36 string = v2
                _ = v36
            }
            var v37 uint8 = 212
            _ = fi35
            _ = v37
        }
    }
    var v38 int64 = -321
    var v39 int32 = (v3 ^ -745)
    v2 = (v2 + v2)
    v5 -= (^(v5))
    var v40 uint8 = 127
    v32 = uint64(v40)
    var v41 int64 = -1
    var v42 uint16 = (v4 * (v4 >> ((65535) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
