package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0() bool {
    var v7 float64 = 0.0
    switch 0 {
    case 4294967295:
        var v8 uint16 = 1
        _ = v8
    case 747:
        var v9 bool = (!((-2147483648 >= 2147483647) && true))
        _ = v9
    default:
        var v10 uint32 = 416
        _ = v10
    }
    fmt.Println(v7)
    v7 -= 1e10
    switch "hello" {
    case "":
        if (621 <= -2147483648) {
            v7 += v7
            var v11 uint16 = 2
            fmt.Println(v11, v11)
            _ = v11
        } else {
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 uint32 = 285
                v14 := []int{}
                v13 = (v13 % (((v13 * (v13 / (((^(v13))) | 1)))) | 1))
                _ = fi12
                _ = v13
                _ = v14
            }
            v15 := []int{-431}
            _ = v15
        }
    default:
        var v16 uint8 = 128
        _ = v16
    }
    _ = v7
    return ((true && (("hello" <= "foo bar") && (("\n\t" == "") || false))) || (733 != 2))
}
func genFn1(p0 int8) uint64 {
    for fi17 := 0; fi17 < 6; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            if true {
                p0 += (p0 >> ((4294967295) & 63))
                p0 -= (p0 >> ((65535) & 63))
                var v19 int8 = (p0 * p0)
                _ = v19
            }
            v20 := GenS1{F0: 0, F1: (p0 - p0)}
            p0 = (p0 + -128)
            v21 := []int{(-(fi17)), fi18}
            _ = fi18
            _ = v20
            _ = v21
        }
        _ = fi17
    }
    switch 2147483647 {
    case 32768:
        p0 += p0
    case -9223372036854775808:
        for fi22 := 0; fi22 < 16; fi22++ {
            p0 -= 42
            if true {
                if true {
                    var v23 float32 = (1.5 / -2.25)
                    _ = v23
                }
                switch 769 {
                case -1:
                    p0 += p0
                case 42:
                    var v24 int = 4294967295
                    _ = v24
                }
                var v25 string = ((("foo bar" + (((("hello" + ("" + "hello")) + "foo bar") + (("hello" + "hello") + "")) + ((("" + ("" + "\n\t")) + ((((((("" + (("\n\t" + ("a" + ("hello" + ("a" + "hello")))) + "hello")) + "") + "") + "a") + "hello") + "foo bar") + "hello")) + "\n\t"))) + "a") + "")
                _ = v25
            }
            var v26 uint = 809
            _ = fi22
            _ = v26
        }
    default:
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 float64 = -2.25
            switch "\n\t" {
            case "a":
                v29 := []int{fi27}
                _ = v29
            default:
                var v30 bool = true
                _ = v30
            }
            _ = fi27
            _ = v28
        }
    }
    switch ("" + "") {
    case "hello":
        v31 := GenS1{F0: 9223372036854775807, F1: (p0 >> ((42) & 63))}
        _ = v31
    }
    v32 := GenS1{F0: 9223372036854775807, F1: p0}
    v33 := []int{128, 578, 9223372036854775807, 2147483647}
    _ = p0
    _ = v32
    _ = v33
    return 1
}

func main() {
    var v1 int8 = 1
    var v2 uint8 = 94
    var v3 string = "hello"
    var v4 uint8 = (v2 - v2)
    v5 := []int{}
    var v6 float64 = 0.1
    fmt.Println(v3)
    v34 := GenS1{F0: 42, F1: v1}
    if false {
        for fi35 := 0; fi35 < 13; fi35++ {
            v5 = append(v5, (^(fi35)))
            fmt.Println(fi35)
            _ = fi35
        }
        fmt.Println(v6, v4, v3)
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        for fi37 := 0; fi37 < 20; fi37++ {
            v6 += float64((1e10 * 0.1))
            if false {
                _ = len(v5)
            }
            _ = fi37
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v38 := GenS2{F0: 1, F1: "a"}
        _ = fi36
        _ = v38
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        for fi40 := 0; fi40 < 1; fi40++ {
            var v41 bool = (!true)
            fmt.Println(v1)
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    v42 := GenS1{F0: 9223372036854775807, F1: (-(v1))}
    v43 := GenS1{F0: -927, F1: -1}
    var v44 bool = ((true || false) && true)
    if (413 == -2147483648) {
        v4 = (^(v2))
    } else {
        v45 := GenS3{F0: -59, F1: 924}
        var v46 int32 = -625
        _ = v45
        _ = v46
    }
    fmt.Println(v4)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
