package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 string; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint64; F2 uint64 }
func genFn0(p0 uint) uint64 {
    v5 := []int{32768, -9223372036854775808, 32768, 32767}
    if true {
        if false {
            var v6 uint = p0
            var v7 int32 = -2147483648
            _ = v6
            _ = v7
        }
        var v8 uint = p0
        if (42 == 65535) {
            switch 2147483647 {
            case 0:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                p0 -= p0
            }
        }
        _ = v8
    } else {
        v9 := GenS1{F0: 608, F1: 729, F2: -205}
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        p0 = p0
        _ = fi10
    }
    _ = p0
    _ = v5
    return 4294967295
}
func genFn1(p0 string, p1 uint32) uint32 {
    var v11 uint16 = 2
    var v12 uint64 = 845
    if false {
        if true {
            fmt.Println(v12, p0)
            for fi13 := 0; fi13 < 3; fi13++ {
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(p0, v11)
                default:
                    v14 := []int{(-(fi13))}
                    _ = v14
                }
                var v15 float32 = 0.0
                _ = fi13
                _ = v15
            }
            v16 := GenS2{F0: (p1 + 1), F1: 42}
            _ = v16
        } else {
            fmt.Println(p1, p0)
        }
        if (genFn0(65535) != v12) {
            if (!(!(9223372036854775807 >= -901))) {
                var v17 int32 = 0
                _ = v17
            } else {
                var v18 int64 = 266
                var v19 int32 = 1
                for fi20 := 0; fi20 < 1; fi20++ {
                    if (v18 > v18) {
                        p0 = (p0 + p0)
                        var v21 string = (p0 + (p0 + p0))
                        v22 := []int{}
                        _ = v21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 17; fi23++ {
                        v24 := GenS2{F0: 1, F1: fi23}
                        _ = fi23
                        _ = v24
                    }
                    _ = fi20
                }
                _ = v18
                _ = v19
            }
            v25 := GenS3{F0: "", F1: uint32(v12), F2: p1}
            fmt.Println(v11, p0, v11)
            _ = v25
        } else {
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := []int{}
                v28 := GenS4{F0: 182, F1: genFn0(629), F2: v12}
                _ = fi26
                _ = v27
                _ = v28
            }
            p1 = (p1 ^ p1)
        }
        for fi29 := 0; fi29 < 7; fi29++ {
            p0 = "\n\t"
            _ = fi29
        }
    } else {
        for fi30 := 0; fi30 < 5; fi30++ {
            fmt.Println(v11, p1, v12)
            v31 := GenS4{F0: 1, F1: 0, F2: genFn0(42)}
            p1 += 2147483648
            _ = fi30
            _ = v31
        }
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 2147483648
}

func main() {
    v1 := []int{2147483647}
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 bool = ((42 == 1) || ((false && false) && true))
    if v4 {
        var v32 int = 299
        v33 := GenS4{F0: 2, F1: 355, F2: 4294967295}
        v4 = (v4 && ((v32 + (v32 | (v32 - (v32 << ((255) & 63))))) <= v32))
        _ = v32
        _ = v33
    } else {
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 uint16 = 381
            switch (fi34 >> ((354) & 63)) {
            case 32767:
                delete(v2, ("hello" + "hello"))
            default:
                var v36 float64 = (((((1.5 + 0.0) + 1e10) - (3.14159 * 3.14159)) * (0.1 * 0.0)) * 0.1)
                _ = v36
            }
            _ = fi34
            _ = v35
        }
        v4 = v4
    }
    fmt.Println(v4)
    v4 = true
    switch ("" + "\n\t") {
    case "hello":
        var v37 int8 = 89
        _ = v37
    }
    switch 32768 {
    case 611:
        if v4 {
            if (((("\n\t" + "a") + "a") + (("" + ("" + "\n\t")) + "hello")) < "hello") {
                fmt.Println(v4)
            } else {
                v38 := GenS4{F0: 42, F1: genFn0(1), F2: 2}
                _ = v38
            }
            switch -1 {
            case 9223372036854775807:
                v39 := GenS2{F0: 0, F1: 2147483647}
                _ = v39
            case -1:
                v40 := GenS2{F0: 797, F1: 178}
                _ = v40
            }
            v4 = v4
        } else {
            var v41 int16 = -953
            _ = v41
        }
    case -2147483648:
        var v42 bool = (false || true)
        _ = v42
    case 1:
        var v43 int32 = 0
        _ = v43
    }
    _ = len(v2)
    var v44 int64 = 465
    fmt.Println(v4, v44)
    fmt.Println(v4, v4)
    v4 = true
    v1 = append(v1, 32767)
    switch "hello" {
    case "\n\t":
        var v45 float32 = float32(0.0)
        _ = v45
    case "hello":
        var v46 int16 = 1
        _ = v46
    case "":
        v47 := GenS3{F0: "foo bar", F1: 2147483647, F2: genFn1((("\n\t" + "") + "foo bar"), 4294967295)}
        _ = v47
    default:
        for fi48 := 0; fi48 < 7; fi48++ {
            switch "\n\t" {
            case "\n\t":
                v44 = v44
            case "hello":
                fmt.Println(v4)
            case "foo bar":
                v4 = true
            }
            var v49 bool = ((-9223372036854775808 > fi48) && (!(406 > 42)))
            var v50 uint16 = 0
            _ = fi48
            _ = v49
            _ = v50
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
}
