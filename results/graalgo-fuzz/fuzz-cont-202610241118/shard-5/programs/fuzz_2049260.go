package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 int32; F2 int64 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 string, p1 int, p2 string) uint8 {
    if true {
        fmt.Println(p1, p1, p1)
        for fi7 := 0; fi7 < 17; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                if true {
                    p0 = "foo bar"
                    v9 := GenS1{F0: 734}
                    if (true && (!true)) {
                        switch p1 {
                        case 2147483647:
                            var v10 float32 = 1e10
                            _ = v10
                        default:
                            var v11 uint8 = 0
                            _ = v11
                        }
                        switch p0 {
                        case "hello":
                            fmt.Println(fi7, fi7, fi7)
                        case "\n\t":
                            v12 := []int{(fi8 + fi8)}
                            _ = v12
                        default:
                            v13 := GenS1{F0: -1}
                            _ = v13
                        }
                    } else {
                        var v14 float64 = 0.1
                        v15 := []int{fi8, fi7, -75, fi8}
                        p2 = "a"
                        _ = v14
                        _ = v15
                    }
                    _ = v9
                }
                var v16 int16 = -1
                _ = fi8
                _ = v16
            }
            var v17 int = p1
            _ = fi7
            _ = v17
        }
    } else {
        p2 = p2
        if true {
            if (124 > 2) {
                v18 := []int{(-(p1)), (p1 - 1), 32768, p1}
                if (p1 < (p1 * p1)) {
                    var v19 int = p1
                    p0 = "foo bar"
                    _ = v19
                }
                _ = v18
            } else {
                p1 = 32767
                p1 = (-(p1))
                var v20 float64 = float64(0.1)
                _ = v20
            }
            if true {
                fmt.Println(p1, p2)
                switch p1 {
                case 721:
                    v21 := GenS2{F0: (("" + (p2 + p2)) + "\n\t")}
                    _ = v21
                default:
                    var v22 int = p1
                    _ = v22
                }
            }
        }
        if (-658 > -486) {
            var v23 int64 = -147
            var v24 uint = 10
            _ = v23
            _ = v24
        } else {
            switch p1 {
            case -878:
                v25 := GenS3{F0: 42}
                _ = v25
            default:
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 uint64 = 18446744073709551615
                    switch fi26 {
                    case 1:
                        v28 := GenS3{F0: 177}
                        _ = v28
                    }
                    p0 = "foo bar"
                    _ = fi26
                    _ = v27
                }
            }
        }
    }
    switch p2 {
    case "":
        var v29 string = p0
        _ = v29
    case "foo bar":
        fmt.Println(p0, p2)
    case "\n\t":
        v30 := GenS4{F0: 255, F1: 0, F2: 1}
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}
func genFn1(p0 uint32) string {
    v31 := GenS4{F0: 886, F1: 716, F2: 0}
    switch 429 {
    case -939:
        v32 := GenS2{F0: (("foo bar" + "hello") + "")}
        _ = v32
    case -21:
        if false {
            var v33 bool = true
            _ = v33
        } else {
            var v34 float32 = 0.1
            p0 -= (p0 << ((1) & 63))
            _ = v34
        }
    }
    p0 = p0
    for fi35 := 0; fi35 < 2; fi35++ {
        var v36 int = (fi35 & fi35)
        switch (fi35 | (fi35 >> ((65535) & 63))) {
        case 292:
            p0 = (p0 >> ((2) & 63))
        case -496:
            var v37 float32 = 0.1
            _ = v37
        }
        v38 := []int{(v36 + v36), 146, (fi35 ^ (fi35 << ((255) & 63)))}
        _ = fi35
        _ = v36
        _ = v38
    }
    _ = p0
    _ = v31
    return ""
}

func main() {
    var v1 uint64 = 0
    v2 := []uint32{493, 42}
    var v3 uint16 = uint16(v1)
    var v4 int32 = 2147483647
    v5 := make(map[string]int)
    var v6 bool = false
    var v39 int16 = 42
    var v40 uint32 = 0
    fmt.Println(v1, v4, v40)
    var v41 int16 = (v39 | -41)
    v40 = v40
    for fi42 := 0; fi42 < 3; fi42++ {
        var v43 int64 = -789
        v40 -= 4294967295
        _ = fi42
        _ = v43
    }
    for fi44 := 0; fi44 < 3; fi44++ {
        for fi45 := 0; fi45 < 14; fi45++ {
            if (v6 || ((v41 != v39) || (!v6))) {
                var v46 string = "foo bar"
                if (v6 || ((128 < fi45) && v6)) {
                    v41 += 115
                }
                _ = v46
            }
            _ = fi45
        }
        var v47 float32 = (float32(((1e10 - 1.5) + (float64(0.1) * 1.5))) / 1.5)
        _ = fi44
        _ = v47
    }
    switch "" {
    case "":
        var v48 int64 = 848
        _ = v48
    case "a":
        v40 -= (v40 ^ (v40 ^ uint32(v3)))
    case "foo bar":
        for fi49 := 0; fi49 < 17; fi49++ {
            var v50 float64 = float64(3.14159)
            var v51 uint64 = v1
            v52 := GenS5{F0: (v1 % ((v51) | 1))}
            _ = fi49
            _ = v50
            _ = v51
            _ = v52
        }
    default:
        for fi53 := 0; fi53 < 2; fi53++ {
            v54 := GenS3{F0: 2147483647}
            _ = fi53
            _ = v54
        }
    }
    v1 = (v1 << ((723) & 63))
    if (true && v6) {
        if len(v2) > 0 {
            v2[0] = 1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
