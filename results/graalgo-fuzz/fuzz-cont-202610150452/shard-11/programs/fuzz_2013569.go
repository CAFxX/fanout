package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint64, p2 uint8) int8 {
    v4 := GenS1{F0: p1}
    for fi5 := 0; fi5 < 1; fi5++ {
        p1 = (p1 + 4294967295)
        v6 := []int{(fi5 | fi5), (fi5 & (fi5 ^ 32767))}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 100
}

func main() {
    var v1 int16 = 42
    var v2 int64 = -9223372036854775808
    var v3 int64 = (v2 * (v2 & (v2 * v2)))
    if true {
        if (!false) {
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 string = (((((("" + "") + "hello") + "a") + ("foo bar" + ("foo bar" + ""))) + "hello") + "\n\t")
                var v9 int8 = -1
                _ = fi7
                _ = v8
                _ = v9
            }
            switch -1 {
            case 42:
                v3 = (v3 & 9223372036854775807)
            }
            switch "a" {
            case "":
                v10 := GenS1{F0: 576}
                _ = v10
            case "hello":
                switch ("a" + "a") {
                case "":
                    v11 := []int{127, 926, 85, 343}
                    _ = v11
                default:
                    switch -770 {
                    case 582:
                        var v12 uint16 = 641
                        _ = v12
                    case 4294967295:
                        v13 := GenS1{F0: 108}
                        _ = v13
                    default:
                        v3 -= (v2 + (v2 - (v2 + v2)))
                    }
                }
            }
        }
        v2 = -1
        v3 = (v3 | -9223372036854775808)
    } else {
        v14 := []int{388}
        v1 += -1
        _ = v14
    }
    var v15 bool = (true && ((false || false) && false))
    switch 278 {
    case 692:
        if (((1e10 * ((1e10 + (1e10 - float64((-2.25 * -2.25)))) + float64(-2.25))) - ((3.14159 + (0.0 / (-2.25 + 0.0))) / -2.25)) == (1.5 + ((0.0 / 1.5) + (1e10 * 0.1)))) {
            var v16 int16 = 1000
            switch "foo bar" {
            case "a":
                for fi17 := 0; fi17 < 4; fi17++ {
                    fmt.Println(fi17)
                    if (85 < 65535) {
                        v2 = v3
                    }
                    _ = fi17
                }
            default:
                if false {
                    v15 = false
                    v18 := []int{-557, -578, 571}
                    _ = v18
                } else {
                    var v19 int64 = -9223372036854775808
                    if true {
                        v20 := GenS1{F0: 2}
                        _ = v20
                    }
                    _ = v19
                }
            }
            for fi21 := 0; fi21 < 4; fi21++ {
                v22 := []int{(fi21 + -9223372036854775808), fi21, int(v3), fi21}
                fmt.Println(v15, v15, v2)
                v23 := GenS1{F0: 643}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v16
        }
    case 128:
        if v15 {
            if ((false || true) || v15) {
                for fi24 := 0; fi24 < 9; fi24++ {
                    var v25 uint64 = 42
                    _ = fi24
                    _ = v25
                }
                v26 := GenS1{F0: 1}
                v1 = 1
                _ = v26
            } else {
                v27 := GenS1{F0: 255}
                switch 2147483647 {
                case -214:
                    fmt.Println(v3)
                default:
                    v28 := GenS1{F0: 792}
                    _ = v28
                }
                _ = v27
            }
        } else {
            for fi29 := 0; fi29 < 15; fi29++ {
                v30 := GenS1{F0: 0}
                if (65535 >= 782) {
                    v31 := []int{fi29, fi29, 1, -2147483648}
                    var v32 uint = 4294967295
                    _ = v31
                    _ = v32
                }
                _ = fi29
                _ = v30
            }
            v33 := []int{-2147483648, 4294967295, -196, 128}
            _ = v33
        }
    }
    if v15 {
        v34 := GenS1{F0: 2}
        v35 := []int{-9223372036854775808, -354, -2147483648, -2147483648}
        _ = v34
        _ = v35
    }
    fmt.Println(v2, v15, v15)
    fmt.Println(v3, v15)
    for fi36 := 0; fi36 < 1; fi36++ {
        fmt.Println(v1, fi36)
        _ = fi36
    }
    v37 := GenS1{F0: 1}
    if ((v1 * v1) < (^(v1))) {
        v3 += v2
        v38 := GenS1{F0: 2}
        _ = v38
    }
    var v39 uint64 = 0
    v40 := []int{-826, 2147483647, 1}
    if v15 {
        v15 = (0 >= 86)
        for fi41 := 0; fi41 < 12; fi41++ {
            if len(v40) >= 2 {
                v40 = v40[1:len(v40)]
            }
            _ = fi41
        }
    }
    fmt.Println(v2)
    v42 := GenS1{F0: (v39 >> ((255) & 63))}
    fmt.Println(v15, v1, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v40
    _ = v42
}
