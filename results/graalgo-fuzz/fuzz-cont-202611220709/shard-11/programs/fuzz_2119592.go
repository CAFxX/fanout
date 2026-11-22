package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int; F1 int8; F2 uint16 }
func genFn0(p0 int, p1 int16) uint32 {
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            p1 -= 42
            _ = fi8
        }
        _ = fi7
    }
    p0 += (^(p0))
    if (1 <= 4294967295) {
        v9 := GenS1{F0: 1}
        p1 -= (p1 % ((32767) | 1))
        p0 = p0
        _ = v9
    } else {
        switch -1 {
        case -573:
            fmt.Println(p0, p0)
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := []int{}
                var v13 uint = 255
                p0 -= 0
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
        v14 := []int{}
        _ = v14
    }
    _ = p0
    _ = p1
    return 233
}

func main() {
    var v1 float32 = float32(((0.1 * ((1.5 + 3.14159) * float64(float32(3.14159)))) / float64(1.5)))
    v2 := make(map[string]int)
    var v3 string = "foo bar"
    var v4 uint64 = 1
    var v5 bool = ((0 != -365) || (!false))
    v6 := make(map[int]bool)
    fmt.Println(v4)
    v2[v3] = -2147483648
    fmt.Println(v4, v3, v1)
    var v15 int16 = -1
    var v16 uint8 = 1
    v3 = ""
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := []int{fi17, fi17}
        var v19 uint64 = v4
        v20 := GenS2{F0: (v15 >> ((1) & 63)), F1: true}
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    v21 := []int{9223372036854775807, -463, -9223372036854775808, 32767}
    if false {
        switch v3 {
        case "a":
            var v22 uint64 = (v4 ^ v4)
            _ = v22
        case "\n\t":
            for fi23 := 0; fi23 < 17; fi23++ {
                v3 = v3
                if v5 {
                    v4 = (v4 % ((0) | 1))
                }
                for fi24 := 0; fi24 < 3; fi24++ {
                    v25 := v2[""]
                    var v26 int16 = v15
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
        case "hello":
            v4 = (v4 / (((v4 << ((uint(v16)) & 63))) | 1))
        }
        if v5 {
            v5 = (!(924 <= 4294967295))
        } else {
            for fi27 := 0; fi27 < 14; fi27++ {
                fmt.Println(fi27, v1)
                v5 = (9223372036854775807 <= 9223372036854775807)
                _ = fi27
            }
            v16 -= 127
        }
    } else {
        v28 := GenS1{F0: (^(v16))}
        _ = v28
    }
    v29 := GenS2{F0: (v15 - (v15 | (v15 ^ v15))), F1: (2147483647 > 27)}
    var v30 float32 = v1
    v4 += (v4 | v4)
    if false {
        var v31 int = -1
        for fi32 := 0; fi32 < 20; fi32++ {
            _ = len(v21)
            _ = fi32
        }
        _ = v31
    } else {
        v33 := GenS1{F0: v16}
        var v34 uint64 = (^(v4))
        for fi35 := 0; fi35 < 16; fi35++ {
            var v36 uint32 = 65535
            for fi37 := 0; fi37 < 19; fi37++ {
                if v5 {
                    switch "\n\t" {
                    case "foo bar":
                        v3 = v3
                    case "":
                        _ = len(v21)
                    case "hello":
                        v15 += (v15 / (((v15 + (v15 >> ((255) & 63)))) | 1))
                    default:
                        v15 = v15
                    }
                } else {
                    switch ((v3 + ("" + v3)) + ((v3 + v3) + v3)) {
                    case "\n\t":
                        var v38 uint32 = v36
                        _ = v38
                    case "a":
                        var v39 float32 = v1
                        _ = v39
                    case "":
                        var v40 uint64 = 0
                        _ = v40
                    }
                }
                _ = fi37
            }
            v41 := v6[fi35]
            _ = fi35
            _ = v36
            _ = v41
        }
        _ = v33
        _ = v34
    }
    fmt.Println(v15, v30)
    var v42 int64 = 42
    for fi43 := 0; fi43 < 15; fi43++ {
        if false {
            var v44 int16 = (v15 * (v15 >> ((253) & 63)))
            v3 = v3
            v44 += -32768
            _ = v44
        } else {
            for fi45 := 0; fi45 < 16; fi45++ {
                if (-2.25 != 1e10) {
                    v4 = (v4 >> ((2) & 63))
                    fmt.Println(v4)
                    v46 := GenS3{F0: v42, F1: v15}
                    _ = v46
                } else {
                    v47 := GenS2{F0: (^(v15)), F1: ("foo bar" > "a")}
                    if (!v5) {
                        var v48 uint32 = 4294967295
                        v49 := GenS2{F0: v15, F1: ((v5 || false) || (v30 == -2.25))}
                        _ = v48
                        _ = v49
                    } else {
                        v30 -= v30
                        v50 := GenS2{F0: (v15 + (v15 << ((0) & 63))), F1: (v16 >= v16)}
                        _ = v50
                    }
                    _ = v47
                }
                delete(v6, fi45)
                _ = fi45
            }
            v1 += (v30 + (((3.14159 * v30) * ((v30 * 0.0) - 1e10)) / 3.14159))
            v51 := GenS4{F0: (fi43 | 128), F1: 0, F2: 42}
            _ = v51
        }
        _ = len(v2)
        _ = fi43
    }
    v1 -= float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
