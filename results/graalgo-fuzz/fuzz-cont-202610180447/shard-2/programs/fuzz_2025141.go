package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint32, p2 uint8) uint32 {
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := []int{fi4}
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 uint = 65535
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    var v8 int32 = 981
    p1 -= 989
    if (true || (0 != int8(p1))) {
        v9 := []int{32768, 9223372036854775807}
        for fi10 := 0; fi10 < 1; fi10++ {
            switch "" {
            case "a":
                v11 := GenS1{F0: 9223372036854775807, F1: 127, F2: (p2 | (^(p2)))}
                _ = v11
            case "hello":
                var v12 uint64 = 4294967295
                _ = v12
            }
            _ = fi10
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p1 << ((255) & 63))
}
func genFn1(p0 bool) uint {
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS2{F0: 1}
        v15 := GenS3{F0: (("\n\t" + "\n\t") + "hello")}
        var v16 uint = 281
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := GenS4{F0: 127}
    if (p0 && (!(p0 && false))) {
        if ((0 > 2) || ((-9223372036854775808 == 390) || true)) {
            var v18 string = "hello"
            if (123 != 1) {
                var v19 float64 = 1e10
                v18 = v18
                _ = v19
            }
            _ = v18
        }
        switch ("foo bar" + ((("hello" + ("hello" + ("foo bar" + ("foo bar" + "\n\t")))) + ("\n\t" + "a")) + "hello")) {
        case "":
            var v20 string = ("a" + "foo bar")
            _ = v20
        case "\n\t":
            v21 := []int{}
            _ = v21
        case "hello":
            var v22 uint = 537
            _ = v22
        default:
            v23 := GenS5{F0: 0, F1: 0}
            _ = v23
        }
        p0 = true
    } else {
        var v24 string = "foo bar"
        v25 := []int{}
        if p0 {
            p0 = (1 > 1000)
            for fi26 := 0; fi26 < 15; fi26++ {
                fmt.Println(v24)
                var v27 int8 = -18
                _ = fi26
                _ = v27
            }
            switch v24 {
            case "\n\t":
                v28 := GenS3{F0: v24}
                _ = v28
            }
        }
        _ = v24
        _ = v25
    }
    _ = p0
    _ = v17
    return 4294967295
}

func main() {
    var v1 string = ("" + ("\n\t" + ("a" + "")))
    var v2 uint8 = 61
    var v3 int8 = 0
    v29 := []int{int(v3), -2147483648, -9223372036854775808, 620}
    var v30 float64 = 3.14159
    switch 894 {
    case 127:
        if true {
            if true {
                v30 += ((0.1 - -2.25) - v30)
                fmt.Println(v30)
            } else {
                v3 += v3
            }
            if (3.14159 > 0.0) {
                for fi31 := 0; fi31 < 14; fi31++ {
                    for fi32 := 0; fi32 < 13; fi32++ {
                        var v33 float32 = 0.1
                        v2 += 255
                        _ = fi32
                        _ = v33
                    }
                    v2 = (v2 - v2)
                    _ = fi31
                }
            } else {
                v34 := GenS2{F0: 255}
                var v35 int8 = (v3 % ((v3) | 1))
                _ = v34
                _ = v35
            }
            fmt.Println(v2, v3)
        } else {
            v36 := GenS1{F0: -987, F1: 1, F2: (^(v2))}
            if ((-(v3)) != -54) {
                v37 := GenS6{F0: "\n\t", F1: 948}
                _ = v37
            } else {
                v38 := GenS6{F0: "a", F1: 1}
                for fi39 := 0; fi39 < 7; fi39++ {
                    switch -2147483648 {
                    case 4294967295:
                        v30 = 1e10
                    case 601:
                        fmt.Println(v2)
                    case 2147483647:
                        v40 := GenS7{F0: 42}
                        _ = v40
                    default:
                        var v41 float32 = float32((v30 - v30))
                        _ = v41
                    }
                    _ = fi39
                }
                _ = v38
            }
            _ = v36
        }
    default:
        for fi42 := 0; fi42 < 1; fi42++ {
            if true {
                v30 = v30
            } else {
                switch v1 {
                case "a":
                    var v43 int16 = -1
                    _ = v43
                case "foo bar":
                    v3 += v3
                case "\n\t":
                    if len(v29) > 0 {
                        v29[0] = fi42
                    }
                default:
                    switch (-(fi42)) {
                    case -727:
                        var v44 string = v1
                        _ = v44
                    default:
                        _ = len(v29)
                    }
                }
                v29 = append(v29, fi42)
                var v45 string = (v1 + ("\n\t" + v1))
                _ = v45
            }
            var v46 uint64 = 2
            v47 := GenS3{F0: (v1 + (v1 + (v1 + (("foo bar" + "\n\t") + ((("a" + ("hello" + "hello")) + "\n\t") + "")))))}
            _ = fi42
            _ = v46
            _ = v47
        }
    }
    v3 = 57
    if true {
        fmt.Println(v1, v30, v1)
    }
    v48 := GenS5{F0: 1, F1: 434}
    var v49 string = v1
    if (!(!(!(false || (v30 > 3.14159))))) {
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
        if false {
            if len(v29) > 0 {
                v29[0] = 1
            }
        }
    }
    if false {
        v50 := GenS2{F0: 255}
        var v51 uint16 = 663
        v30 -= (0.1 + (0.0 + (v30 * v30)))
        _ = v50
        _ = v51
    }
    var v52 uint = 42
    var v53 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v49:", v49)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v48
    _ = v49
    _ = v52
    _ = v53
}
