package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint16 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 string; F1 int16; F2 int16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint64, p1 int64, p2 uint32) int16 {
    switch int(p0) {
    case 327:
        var v7 bool = true
        _ = v7
    default:
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := []int{(fi8 - fi8), (fi8 / ((fi8) | 1)), 42, 225}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            var v10 string = "foo bar"
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    p1 = p1
    p1 = -977
    var v11 int = 32768
    v12 := GenS1{F0: v11, F1: ("\n\t" + "foo bar"), F2: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -32768
}
func genFn1(p0 int64) uint32 {
    fmt.Println(p0)
    for fi13 := 0; fi13 < 13; fi13++ {
        fmt.Println(fi13)
        _ = fi13
    }
    switch 4294967295 {
    case 659:
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := GenS1{F0: (fi14 >> ((65535) & 63)), F1: ("foo bar" + "a"), F2: 65535}
            _ = fi14
            _ = v15
        }
    case -944:
        p0 = 9223372036854775807
    case -80:
        var v16 bool = true
        _ = v16
    default:
        if false {
            p0 += p0
            if false {
                var v17 int = 128
                v17 -= (-(v17))
                if (!(float64((0.0 - float32(0.1))) > float64((0.0 * -2.25)))) {
                    if true {
                        switch "a" {
                        case "":
                            var v18 int32 = 0
                            _ = v18
                        default:
                            v17 -= (v17 / (((v17 % ((v17) | 1))) | 1))
                        }
                        v19 := []int{}
                        var v20 uint16 = 0
                        _ = v19
                        _ = v20
                    }
                }
                _ = v17
            } else {
                fmt.Println(p0)
                var v21 bool = false
                _ = v21
            }
            if true {
                v22 := GenS2{F0: "", F1: 1}
                _ = v22
            } else {
                switch "" {
                case "a":
                    p0 = p0
                }
                v23 := []int{42, 145, 0}
                var v24 string = "foo bar"
                _ = v23
                _ = v24
            }
        }
    }
    var v25 string = ""
    _ = p0
    _ = v25
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 0.1
    v3 := make(map[string]int)
    var v4 uint = 0
    var v5 float32 = 1.5
    v6 := []string{"", "", ("\n\t" + "")}
    if len(v6) > 0 {
        v6[0] = ("\n\t" + ("" + "hello"))
    }
    v4 = (v4 << ((65535) & 63))
    switch "hello" {
    case "foo bar":
        switch "foo bar" {
        case "foo bar":
            var v26 uint8 = 2
            _ = v26
        default:
            var v27 uint64 = 42
            _ = v27
        }
    default:
        v5 = (3.14159 - (v5 - (((float32((0.1 + (v2 - v2))) - float32(v2)) + (v5 / v5)) + v5)))
    }
    if (genFn0(474, 42, 2147483647) < 1) {
        if (42 > 42) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v28 := GenS3{F0: "", F1: 0, F2: 32767}
            v29 := GenS2{F0: "foo bar", F1: 2147483648}
            _ = v28
            _ = v29
        }
        v30 := v1[506]
        for fi31 := 0; fi31 < 3; fi31++ {
            for fi32 := 0; fi32 < 20; fi32++ {
                _ = len(v1)
                if (2 == 83) {
                    v4 = (v4 - (v4 ^ (^(v4))))
                }
                v33 := GenS1{F0: (^(v30)), F1: "foo bar", F2: 432}
                _ = fi32
                _ = v33
            }
            fmt.Println(v4, fi31)
            _ = fi31
        }
        _ = v30
    }
    v5 += (((((v5 * v5) + ((float32(v2) / v5) * ((v5 + 3.14159) + v5))) / (v5 - float32(1.5))) * v5) / v5)
    switch ("" + "hello") {
    case "foo bar":
        for fi34 := 0; fi34 < 8; fi34++ {
            var v35 uint = (^(v4))
            v4 = v4
            _ = fi34
            _ = v35
        }
    case "\n\t":
        fmt.Println(v5, v5)
    default:
        v5 = (v5 * v5)
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        for fi37 := 0; fi37 < 16; fi37++ {
            var v38 int = fi37
            switch fi36 {
            case -9223372036854775808:
                v4 = (v4 - v4)
            case 0:
                var v39 int32 = 1
                _ = v39
            default:
                v38 += -2147483648
            }
            var v40 uint = (v4 | (v4 * v4))
            _ = fi37
            _ = v38
            _ = v40
        }
        _ = fi36
    }
    v6 = append(v6, "\n\t")
    var v41 string = "foo bar"
    var v42 uint64 = 42
    var v43 float64 = 0.0
    var v44 uint32 = 65535
    v45 := GenS4{F0: 1000}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
