package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 bool }
type GenS2 struct { F0 bool; F1 uint }
func genFn0(p0 bool) bool {
    switch 655 {
    case 1:
        var v4 float32 = -2.25
        _ = v4
    default:
        if true {
            for fi5 := 0; fi5 < 15; fi5++ {
                for fi6 := 0; fi6 < 9; fi6++ {
                    var v7 uint16 = 674
                    var v8 uint32 = 4294967295
                    var v9 float32 = 1e10
                    _ = fi6
                    _ = v7
                    _ = v8
                    _ = v9
                }
                switch (fi5 ^ (fi5 & (fi5 - fi5))) {
                case -9223372036854775808:
                    var v10 uint32 = 4294967295
                    _ = v10
                }
                fmt.Println(p0, p0)
                _ = fi5
            }
            v11 := GenS1{F0: 127, F1: 0, F2: p0}
            for fi12 := 0; fi12 < 18; fi12++ {
                fmt.Println(fi12)
                _ = fi12
            }
            _ = v11
        }
    }
    p0 = (p0 && (!(65535 != 32767)))
    switch "foo bar" {
    case "\n\t":
        fmt.Println(p0)
    }
    _ = p0
    return p0
}

func main() {
    var v1 int = 0
    var v2 uint64 = 42
    var v3 uint = 1
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v2, fi13, v3)
        switch "a" {
        case "foo bar":
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 uint32 = 2
                v16 := []int{fi14, v1, -1}
                _ = fi14
                _ = v15
                _ = v16
            }
        case "":
            fmt.Println(v2, v3)
        }
        v17 := []int{v1, (fi13 - (fi13 << (((v3 / ((v3) | 1))) & 63)))}
        _ = fi13
        _ = v17
    }
    if ((!(181 >= 0)) && false) {
        v18 := GenS1{F0: v1, F1: uint8(v3), F2: (!true)}
        var v19 int16 = 1000
        _ = v18
        _ = v19
    }
    var v20 uint8 = 42
    v21 := []int{-178, 127, v1}
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 string = "hello"
        switch (v23 + (v23 + (v23 + (v23 + (v23 + "foo bar"))))) {
        case "hello":
            v1 += 128
        default:
            v1 += (fi22 | 32767)
        }
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 uint16 = 684
            _ = fi24
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    switch (v1 / ((v1) | 1)) {
    case -9223372036854775808:
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 uint16 = 812
            _ = fi26
            _ = v27
        }
    default:
        var v28 string = "hello"
        _ = v28
    }
    if (355 >= 65535) {
        var v29 string = (("hello" + (("\n\t" + "") + ("" + "\n\t"))) + "")
        v20 -= 2
        if false {
            _ = len(v21)
        } else {
            v1 -= -447
        }
        _ = v29
    } else {
        for fi30 := 0; fi30 < 17; fi30++ {
            v1 = (-(v1))
            var v31 uint8 = v20
            _ = fi30
            _ = v31
        }
    }
    if false {
        var v32 uint64 = 0
        _ = v32
    } else {
        var v33 uint64 = (v2 ^ (v2 >> (((^(v3))) & 63)))
        _ = v33
    }
    if false {
        var v34 uint16 = 255
        v35 := GenS1{F0: v1, F1: (v20 >> ((v3) & 63)), F2: (!(-990 <= 413))}
        var v36 bool = (v34 >= 700)
        _ = v34
        _ = v35
        _ = v36
    } else {
        v37 := GenS2{F0: ("foo bar" >= ("foo bar" + "\n\t")), F1: v3}
        _ = v37
    }
    var v38 uint32 = 42
    if true {
        v2 -= (v2 * (v2 * (^(v2))))
        var v39 string = ""
        _ = v39
    }
    if (v3 <= v3) {
        var v40 int8 = 127
        _ = v40
    }
    _ = len(v21)
    var v41 int8 = -62
    var v42 uint64 = (^(v2))
    var v43 float64 = 0.1
    v1 = v1
    var v44 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v38
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
