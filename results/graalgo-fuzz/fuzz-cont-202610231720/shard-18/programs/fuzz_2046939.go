package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 string }
func genFn0(p0 uint, p1 int16) uint {
    fmt.Println(p1, p1)
    if (!(108 <= 177)) {
        p0 = p0
    } else {
        for fi4 := 0; fi4 < 10; fi4++ {
            if true {
                p1 -= (^(p1))
                if false {
                    fmt.Println(p0, p0)
                    p0 += p0
                }
            }
            v5 := []int{}
            switch (fi4 & 2147483647) {
            case -750:
                p1 += 960
            default:
                p1 = p1
            }
            _ = fi4
            _ = v5
        }
    }
    var v6 uint64 = 883
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint32 = 2147483648
            var v10 uint32 = v9
            _ = fi8
            _ = v10
            _ = v9
        }
        p1 = (p1 % ((p1) | 1))
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return (p0 - (^(p0)))
}
func genFn1() uint64 {
    if true {
        if true {
            var v11 string = "\n\t"
            v12 := []int{4294967295, -82, 32768, -888}
            v11 = ("\n\t" + "hello")
            _ = v11
            _ = v12
        } else {
            var v13 int64 = 1
            v13 = v13
            _ = v13
        }
    }
    var v14 uint8 = 243
    v14 = (^(v14))
    v15 := GenS1{F0: true, F1: "hello", F2: v14}
    var v16 uint64 = 4294967295
    _ = v14
    _ = v15
    _ = v16
    return v16
}

func main() {
    var v1 int8 = 1
    v2 := []int{-1, -2147483648}
    var v3 int16 = -32768
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v17 bool = (!(0 >= 0))
    switch -422 {
    case -480:
        v2 = append(v2, 128)
    case 32767:
        for fi18 := 0; fi18 < 2; fi18++ {
            if v17 {
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 int8 = v1
                    fmt.Println(v1, fi19)
                    if len(v2) > 0 {
                        v2[0] = fi18
                    }
                    _ = fi19
                    _ = v20
                }
            } else {
                if len(v2) > 0 {
                    v2[0] = fi18
                }
                _ = len(v2)
            }
            var v21 uint32 = 220
            if len(v2) > 0 {
                v2[0] = fi18
            }
            _ = fi18
            _ = v21
        }
    }
    v2 = append(v2, -921)
    if (!(1e10 > 0.1)) {
        if false {
            var v22 uint32 = 2147483647
            switch ("foo bar" + (("\n\t" + ("foo bar" + ((("\n\t" + "a") + "\n\t") + ("" + ("a" + "foo bar"))))) + (("a" + "") + ("foo bar" + ("a" + ("a" + ("foo bar" + ("hello" + ((("hello" + "\n\t") + "\n\t") + ""))))))))) {
            case "a":
                switch ("hello" + "hello") {
                case "\n\t":
                    v23 := GenS2{F0: 1000}
                    _ = v23
                case "":
                    v1 = v1
                case "foo bar":
                    switch -1 {
                    case -428:
                        v3 = v3
                    }
                default:
                    fmt.Println(v3)
                }
            case "":
                v24 := GenS3{F0: "\n\t", F1: 432}
                _ = v24
            default:
                var v25 uint8 = 99
                _ = v25
            }
            _ = len(v2)
            _ = v22
        }
    } else {
        v1 -= (-(v1))
        var v26 uint = 255
        _ = v26
    }
    if (65535 >= 329) {
        fmt.Println(v3, v17, v3)
        if (!(((v3 * (v3 | -32768)) >= v3) && v17)) {
            v27 := GenS4{F0: 65535, F1: 569}
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 uint16 = 1
                _ = fi28
                _ = v29
            }
            _ = len(v2)
            _ = v27
        }
        v30 := GenS1{F0: v17, F1: "a", F2: 42}
        _ = v30
    } else {
        v1 -= v1
        if len(v2) > 0 {
            v2[0] = -699
        }
        for fi31 := 0; fi31 < 5; fi31++ {
            v32 := GenS2{F0: (v3 / (((v3 >> ((449) & 63))) | 1))}
            var v33 uint32 = 94
            v3 -= (-(v3))
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        v35 := GenS5{F0: "hello"}
        v36 := GenS4{F0: 42, F1: 42}
        _ = fi34
        _ = v35
        _ = v36
    }
    switch "" {
    case "":
        var v37 uint32 = 65535
        _ = v37
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        v2 = append(v2, -706)
    default:
        v3 = 1000
    }
    var v38 float64 = (-2.25 * 0.0)
    v3 -= v3
    fmt.Println(v3, v38)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v38
}
