package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8, p1 string) uint8 {
    for fi4 := 0; fi4 < 7; fi4++ {
        switch p1 {
        case "":
            p1 = "foo bar"
        default:
            p1 = ""
        }
        v5 := []int{(fi4 << ((18446744073709551615) & 63)), fi4, 9223372036854775807}
        _ = fi4
        _ = v5
    }
    p0 += p0
    var v6 int = 1
    switch -9223372036854775808 {
    case 463:
        if false {
            v7 := GenS1{F0: 1, F1: (-107 >= 61)}
            fmt.Println(p1)
            _ = v7
        }
    default:
        var v8 float64 = 1e10
        _ = v8
    }
    if true {
        var v9 string = p1
        switch 205 {
        case 0:
            v9 = "hello"
        case 537:
            switch (-(v6)) {
            case -295:
                p1 = (("foo bar" + "a") + (p1 + "a"))
            case 0:
                switch v9 {
                case "":
                    var v10 bool = false
                    _ = v10
                case "\n\t":
                    fmt.Println(v6)
                case "foo bar":
                    var v11 float64 = (float64(3.14159) * (((3.14159 * -2.25) + (0.0 / -2.25)) / 1.0))
                    _ = v11
                }
            }
        case 42:
            var v12 float32 = ((0.0 * (0.1 - (((3.14159 / float32(3.14159)) / 0.1) / (float32(-2.25) / 1.5)))) + 0.1)
            _ = v12
        }
        _ = v9
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            p0 = 1
            v6 += -2147483648
            var v14 uint16 = 1
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: 136, F1: (128 >= (p0 & p0))}
        fmt.Println(p0, v6, v6)
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v6
    return p0
}
func genFn1() uint16 {
    var v16 uint8 = genFn0(2, "hello")
    v17 := GenS1{F0: 18446744073709551615, F1: false}
    v18 := GenS1{F0: 412, F1: (1.5 > 1.5)}
    v16 = (v16 - (v16 + (v16 | (v16 << ((910) & 63)))))
    switch 32768 {
    case -9223372036854775808:
        v19 := []int{408, 42}
        _ = v19
    case 1:
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := []int{32767, fi20, fi20}
            _ = fi20
            _ = v21
        }
    }
    _ = v16
    _ = v17
    _ = v18
    return 0
}

func main() {
    v1 := []uint32{2147483647, 2}
    var v2 uint16 = 65535
    var v3 bool = ("" == "")
    v1 = append(v1, 42)
    switch "" {
    case "":
        fmt.Println(v3, v3)
    case "a":
        for fi22 := 0; fi22 < 16; fi22++ {
            for fi23 := 0; fi23 < 17; fi23++ {
                fmt.Println(fi22)
                _ = fi23
            }
            var v24 uint32 = 0
            _ = fi22
            _ = v24
        }
    default:
        if (!v3) {
            v25 := GenS1{F0: 0, F1: false}
            _ = v25
        }
    }
    switch 9223372036854775807 {
    case 813:
        v3 = (!(!false))
    case -722:
        if len(v1) > 0 {
            v1[0] = 1
        }
    case 557:
        for fi26 := 0; fi26 < 3; fi26++ {
            if (!true) {
                if v3 {
                    var v27 int8 = -128
                    v3 = (v3 || true)
                    _ = v27
                }
                if v3 {
                    switch fi26 {
                    case 664:
                        var v28 uint = 4294967295
                        _ = v28
                    }
                } else {
                    v29 := GenS1{F0: 0, F1: v3}
                    var v30 int = fi26
                    _ = v29
                    _ = v30
                }
            }
            _ = fi26
        }
    }
    v2 = 0
    var v31 int64 = 1
    var v32 float32 = (1.5 / (3.14159 * (((1e10 - 3.14159) + float32(float64(0.1))) / (float32((float64(1e10) * 1.5)) / float32((-2.25 + 3.14159))))))
    if (-21 < -40) {
        switch 32768 {
        case -706:
            var v33 int64 = (v31 ^ v31)
            _ = v33
        case 62:
            v34 := GenS2{F0: 0}
            _ = v34
        case 32768:
            _ = len(v1)
        default:
            for fi35 := 0; fi35 < 12; fi35++ {
                v36 := GenS2{F0: 2}
                if v3 {
                    fmt.Println(v3, v3)
                }
                fmt.Println(v3, v2, fi35)
                _ = fi35
                _ = v36
            }
        }
        var v37 uint64 = uint64(v2)
        _ = v37
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
}
