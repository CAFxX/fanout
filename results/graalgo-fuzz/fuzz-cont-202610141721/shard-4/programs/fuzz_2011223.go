package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 string; F1 int; F2 int }
func genFn0(p0 int64, p1 uint32) uint32 {
    fmt.Println(p0, p1)
    v7 := GenS1{F0: 0, F1: 42, F2: 2}
    v8 := GenS1{F0: 1, F1: 2, F2: 42}
    if (!false) {
        for fi9 := 0; fi9 < 18; fi9++ {
            for fi10 := 0; fi10 < 14; fi10++ {
                fmt.Println(p0)
                var v11 int64 = (p0 & p0)
                var v12 string = "hello"
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 uint64 = 0
            _ = fi9
            _ = v13
        }
    }
    if (("" + ("" + ("a" + "hello"))) > ("hello" + "foo bar")) {
        fmt.Println(p0)
        var v14 string = "hello"
        _ = v14
    } else {
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 int16, p1 string) uint {
    var v15 uint64 = 551
    var v16 float64 = 1e10
    if false {
        v17 := []int{}
        if len(v17) > 0 {
            v17[0] = 1
        }
        var v18 int64 = 42
        _ = v17
        _ = v18
    } else {
        p0 -= (p0 & p0)
    }
    if (p1 != p1) {
        fmt.Println(p0, p0)
    } else {
        fmt.Println(v16)
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 1
}

func main() {
    v1 := []uint32{0}
    v2 := make(map[int]uint32)
    var v3 uint = 240
    v4 := make(map[string]bool)
    v5 := make(map[string]bool)
    v6 := []int{32767, 61, -1, 441}
    v3 += v3
    fmt.Println(v3)
    var v19 int32 = -1
    _ = len(v6)
    if true {
        if (!(183 == 1)) {
            fmt.Println(v3, v19)
            var v20 int = -374
            _ = v20
        }
    } else {
        if (((v19 & v19) < 2147483647) && (false || (1000 != -32768))) {
            v3 = 2
            fmt.Println(v3, v19)
            if (!false) {
                v19 = (-(v19))
                if (!true) {
                    switch "a" {
                    case "a":
                        var v21 int = 32767
                        _ = v21
                    default:
                        switch "\n\t" {
                        case "":
                            v22 := GenS1{F0: 398, F1: 0, F2: 626}
                            _ = v22
                        case "\n\t":
                            var v23 bool = (148 <= 1)
                            _ = v23
                        case "foo bar":
                            v24 := GenS1{F0: genFn0(-349, 777), F1: 42, F2: 499}
                            _ = v24
                        default:
                            fmt.Println(v19, v3)
                        }
                    }
                    for fi25 := 0; fi25 < 15; fi25++ {
                        v3 = (v3 & 42)
                        v26 := v2[(fi25 + (fi25 | (^(fi25))))]
                        var v27 bool = ("\n\t" != "a")
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
            }
        } else {
            v19 += v19
            if len(v1) > 0 {
                v1[0] = genFn0(int64(v19), 1)
            }
            v3 = (v3 * 443)
        }
        var v28 int16 = -1
        _ = v28
    }
    if false {
        var v29 int64 = -200
        _ = v29
    } else {
        var v30 uint32 = 242
        var v31 int32 = 1
        if false {
            v31 = 1
            var v32 float64 = (-2.25 + 1.5)
            _ = v32
        } else {
            if (!true) {
                var v33 uint32 = v30
                _ = v33
            }
            fmt.Println(v31)
        }
        _ = v30
        _ = v31
    }
    if (0.1 > (float64(float32(0.0)) - 0.0)) {
        var v34 float32 = (0.1 - -2.25)
        fmt.Println(v3, v19)
        v34 -= 1.5
        _ = v34
    } else {
        for fi35 := 0; fi35 < 16; fi35++ {
            fmt.Println(v19, fi35)
            v36 := GenS1{F0: 1, F1: 18446744073709551615, F2: 699}
            switch 0 {
            case 2147483647:
                var v37 uint = 255
                _ = v37
            case -785:
                switch "hello" {
                case "a":
                    var v38 int64 = -9223372036854775808
                    _ = v38
                case "":
                    fmt.Println(fi35, fi35)
                case "foo bar":
                    fmt.Println(fi35)
                }
            default:
                switch ("foo bar" + "") {
                case "foo bar":
                    fmt.Println(v3, v19)
                default:
                    fmt.Println(v19, v3)
                }
            }
            _ = fi35
            _ = v36
        }
        v19 = (v19 * v19)
    }
    for fi39 := 0; fi39 < 19; fi39++ {
        v3 = genFn1(-1, "foo bar")
        switch "a" {
        case "":
            v40 := GenS2{F0: "a", F1: (fi39 & fi39), F2: fi39}
            _ = v40
        case "\n\t":
            delete(v5, "\n\t")
        case "a":
            if (false && (-9223372036854775808 <= 942)) {
                fmt.Println(v3, v19)
                fmt.Println(v3)
                v3 = v3
            } else {
                _ = len(v4)
            }
        }
        fmt.Println(v19)
        _ = fi39
    }
    var v41 int16 = -1
    var v42 string = "a"
    for fi43 := 0; fi43 < 11; fi43++ {
        var v44 int32 = v19
        _ = fi43
        _ = v44
    }
    v19 = v19
    if (100 <= 1) {
        switch 2147483647 {
        case -1:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 128:
            v45 := GenS1{F0: 0, F1: 2, F2: 382}
            _ = v45
        case 32767:
            if true {
                v42 = ""
            }
        }
    }
    v19 -= (v19 / ((v19) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
