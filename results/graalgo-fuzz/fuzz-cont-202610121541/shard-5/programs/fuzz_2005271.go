package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 int }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint8 }
func genFn0(p0 uint8, p1 uint, p2 bool) uint8 {
    if (42 > 9223372036854775807) {
        if (!((255 > 32767) || (p1 < (p1 | p1)))) {
            var v7 uint32 = 65535
            fmt.Println(v7, p1)
            _ = v7
        }
        var v8 uint = (^(p1))
        switch 9223372036854775807 {
        case -730:
            var v9 int16 = 0
            _ = v9
        }
        _ = v8
    } else {
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := []int{(fi10 >> (((^(p1))) & 63)), (fi10 % ((-2147483648) | 1)), fi10}
            if (!true) {
                v12 := GenS1{F0: uint16(p1), F1: int(p0), F2: uint64(p1)}
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        var v13 float32 = float32(0.1)
        _ = v13
    }
    if (((!true) && false) && (p0 == 255)) {
        p0 = (p0 % ((p0) | 1))
        var v14 int32 = -744
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return 52
}
func genFn1() uint32 {
    var v15 string = ""
    v15 = (("a" + v15) + "a")
    v16 := GenS1{F0: 574, F1: 835, F2: 535}
    v15 = ("hello" + v15)
    var v17 uint = 715
    _ = v15
    _ = v16
    _ = v17
    return 192
}

func main() {
    var v1 uint16 = 1
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]bool)
    var v4 int = 1
    v5 := make(map[int]int)
    v6 := make(map[uint32]bool)
    var v18 uint32 = 1
    v19 := []int{}
    v4 = v4
    v20 := GenS1{F0: v1, F1: v4, F2: 4294967295}
    fmt.Println(v1)
    v21 := v6[(v18 | 475)]
    var v22 string = ("hello" + "a")
    if ("hello" == v22) {
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 bool = v21
            _ = fi23
            _ = v24
        }
        if v21 {
            for fi25 := 0; fi25 < 3; fi25++ {
                if false {
                    v6[(v18 % ((0) | 1))] = (genFn1() >= (v18 * (^(v18))))
                    v26 := GenS1{F0: v1, F1: 2147483647, F2: 121}
                    var v27 int = (-(v4))
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 16; fi28++ {
                    v21 = ((v2 << ((65535) & 63)) == v2)
                    var v29 string = v22
                    v5[fi28] = fi28
                    _ = fi28
                    _ = v29
                }
                v4 += 740
                _ = fi25
            }
            var v30 uint8 = 2
            v21 = v21
            _ = v30
        }
    } else {
        if false {
            v31 := GenS2{F0: -47, F1: (v4 & v4), F2: (-(v4))}
            fmt.Println(v22)
            delete(v5, 2147483647)
            _ = v31
        }
        var v32 int8 = 127
        _ = v32
    }
    switch "a" {
    case "a":
        if (float32(0.1) >= (1.5 - float32((1e10 + 1e10)))) {
            for fi33 := 0; fi33 < 1; fi33++ {
                switch v22 {
                case "a":
                    if len(v19) > 0 {
                        v19[0] = 32768
                    }
                case "":
                    var v34 float32 = -2.25
                    _ = v34
                case "foo bar":
                    for fi35 := 0; fi35 < 6; fi35++ {
                        v36 := GenS1{F0: 1, F1: (fi35 / (((fi35 % (((fi33 | (fi35 + (-(fi35))))) | 1))) | 1)), F2: 1}
                        _ = fi35
                        _ = v36
                    }
                }
                v2 += int64(fi33)
                _ = fi33
            }
            var v37 int8 = -1
            _ = v37
        }
    case "foo bar":
        _ = len(v19)
    default:
        for fi38 := 0; fi38 < 17; fi38++ {
            var v39 uint = 252
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            _ = fi38
            _ = v39
        }
    }
    switch "" {
    case "":
        fmt.Println(v21, v2)
    default:
        v4 += v4
    }
    if len(v19) > 0 {
        v19[0] = v4
    }
    if false {
        for fi40 := 0; fi40 < 14; fi40++ {
            if (1 > -543) {
                v41 := GenS3{F0: (v1 & v1), F1: 108, F2: 207}
                _ = v41
            } else {
                if (v21 && (!v21)) {
                    if (127 > genFn0(2, 42, v21)) {
                        var v42 int64 = v2
                        _ = v42
                    }
                }
                var v43 int8 = 66
                _ = len(v5)
                _ = v43
            }
            var v44 uint64 = 371
            var v45 uint = 0
            _ = fi40
            _ = v44
            _ = v45
        }
        v46 := GenS3{F0: 42, F1: 4294967295, F2: 0}
        v19 = append(v19, (v4 ^ v4))
        _ = v46
    }
    var v47 float64 = (0.1 * 3.14159)
    v18 -= v18
    var v48 uint16 = v1
    var v49 string = (v22 + ("\n\t" + "a"))
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
