package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int16 }
func genFn0(p0 bool, p1 int8) uint16 {
    if p0 {
        v5 := GenS1{F0: -1}
        _ = v5
    } else {
        var v6 int8 = p1
        v7 := []int{4294967295, -2147483648, 973, -71}
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        if (false && (!(-2.25 > 1e10))) {
            var v9 int32 = 283
            _ = v9
        } else {
            p0 = p0
            v10 := []int{}
            _ = v10
        }
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := []int{9223372036854775807}
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            v13 := GenS1{F0: 43}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi8
    }
    var v14 int8 = p1
    p1 = 42
    p1 -= -1
    _ = p0
    _ = p1
    _ = v14
    return 18
}
func genFn1(p0 int64, p1 uint, p2 uint) string {
    v15 := []int{4294967295, -711}
    if true {
        switch "\n\t" {
        case "\n\t":
            p2 = p2
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            v17 := GenS1{F0: -340}
            v15 = append(v15, (fi16 & fi16))
            _ = fi16
            _ = v17
        }
        var v18 string = ""
        _ = v18
    } else {
        v19 := GenS2{F0: 736}
        _ = v19
    }
    var v20 uint16 = 1
    if true {
        for fi21 := 0; fi21 < 15; fi21++ {
            p0 -= 1
            switch "" {
            case "foo bar":
                var v22 uint64 = 0
                _ = v22
            case "\n\t":
                p2 += p1
            case "":
                for fi23 := 0; fi23 < 14; fi23++ {
                    v24 := GenS3{F0: ("a" + "foo bar"), F1: p0}
                    fmt.Println(fi23, p0)
                    var v25 int32 = 42
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            default:
                var v26 uint32 = 525
                _ = v26
            }
            var v27 uint16 = (^(v20))
            _ = fi21
            _ = v27
        }
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        fmt.Println(p1)
    } else {
        var v28 uint8 = 11
        var v29 uint32 = 2
        if (-2.25 >= 3.14159) {
            v30 := GenS4{F0: 2}
            _ = v30
        } else {
            v28 -= v28
            switch (("hello" + "\n\t") + "a") {
            case "a":
                if true {
                    p0 = (p0 - 0)
                } else {
                    v31 := GenS3{F0: "a", F1: p0}
                    v32 := GenS5{F0: true, F1: uint64(v29)}
                    _ = v31
                    _ = v32
                }
            default:
                if false {
                    fmt.Println(p0, p2)
                }
            }
            p0 = p0
        }
        _ = v28
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v20
    return ("hello" + (("hello" + "hello") + "\n\t"))
}

func main() {
    var v1 int16 = 0
    v2 := []uint32{4294967295}
    var v3 uint64 = 0
    var v4 uint8 = 127
    v2 = append(v2, 2)
    switch ("foo bar" + "foo bar") {
    case "hello":
        var v33 uint32 = 596
        _ = v33
    default:
        for fi34 := 0; fi34 < 11; fi34++ {
            var v35 int32 = 1
            _ = fi34
            _ = v35
        }
    }
    v3 = v3
    for fi36 := 0; fi36 < 15; fi36++ {
        v1 += v1
        for fi37 := 0; fi37 < 10; fi37++ {
            v4 = v4
            _ = fi37
        }
        v4 = v4
        _ = fi36
    }
    if (1 > 9223372036854775807) {
        v1 = v1
    }
    v38 := GenS6{F0: (-(v1))}
    v2 = append(v2, 2147483648)
    var v39 string = "a"
    v2 = append(v2, 42)
    fmt.Println(v4)
    switch 0 {
    case -9223372036854775808:
        var v40 float64 = (float64(float32((1e10 - 0.1))) / 3.14159)
        _ = v40
    case 2147483647:
        v41 := GenS6{F0: v1}
        _ = v41
    case 1:
        switch 9223372036854775807 {
        case 611:
            v4 = v4
        }
    }
    var v42 uint = 2
    v43 := GenS5{F0: false, F1: (v3 * (v3 / ((1) | 1)))}
    var v44 uint32 = 156
    v45 := GenS1{F0: -2147483648}
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
