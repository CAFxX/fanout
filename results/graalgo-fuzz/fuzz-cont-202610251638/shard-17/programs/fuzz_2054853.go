package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int16; F1 int16; F2 int64 }
func genFn0(p0 int16, p1 int32, p2 string) int16 {
    switch -1 {
    case -9223372036854775808:
        v5 := GenS1{F0: 42}
        _ = v5
    }
    v6 := []int{}
    var v7 float32 = (1e10 / 1.5)
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint = 927
        v10 := GenS1{F0: 0}
        var v11 int32 = (p1 * 2147483647)
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 157
}
func genFn1(p0 int8, p1 int16) int {
    p1 = 42
    if true {
        p0 = p0
        v12 := []int{}
        _ = v12
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []string{"a"}
    var v2 uint32 = 342
    v3 := make(map[int]string)
    var v4 int8 = 0
    _ = len(v1)
    v13 := GenS1{F0: v4}
    for fi14 := 0; fi14 < 6; fi14++ {
        fmt.Println(fi14, v2)
        v3[fi14] = (("" + "a") + "hello")
        _ = fi14
    }
    var v15 bool = false
    v2 += (v2 << ((784) & 63))
    switch 1 {
    case 1:
        switch 42 {
        case 42:
            v15 = (!(3.14159 != (1e10 / 1e10)))
        case -71:
            var v16 string = "\n\t"
            _ = v16
        case -1:
            for fi17 := 0; fi17 < 8; fi17++ {
                fmt.Println(fi17)
                for fi18 := 0; fi18 < 4; fi18++ {
                    v2 -= 2
                    switch (fi18 ^ (fi17 % ((fi18) | 1))) {
                    case 128:
                        var v19 int = fi17
                        _ = v19
                    case 32768:
                        v3[(fi17 + fi18)] = ((((("" + "foo bar") + "hello") + ("foo bar" + "\n\t")) + ("\n\t" + "foo bar")) + ("" + "a"))
                    case 4294967295:
                        v3[-70] = "a"
                    }
                    _ = fi18
                }
                fmt.Println(v2)
                _ = fi17
            }
        default:
            var v20 float64 = 1e10
            _ = v20
        }
    case 32767:
        if true {
            v21 := GenS2{F0: v2, F1: v2}
            _ = v21
        }
    }
    if v15 {
        var v22 int16 = 32767
        for fi23 := 0; fi23 < 6; fi23++ {
            if len(v1) > 0 {
                v1[0] = ""
            }
            for fi24 := 0; fi24 < 19; fi24++ {
                if (32767 <= 2) {
                    v15 = true
                    delete(v3, (fi23 - fi23))
                }
                fmt.Println(fi24, v2)
                v25 := GenS1{F0: (v4 ^ v4)}
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
        v26 := GenS3{F0: 0, F1: 166}
        _ = v22
        _ = v26
    } else {
        v27 := GenS2{F0: v2, F1: (v2 - (^(v2)))}
        var v28 uint16 = 462
        if ((816 >= 255) && (!(-1 != -33))) {
            v4 -= v4
            v29 := v3[127]
            if v15 {
                fmt.Println(v15, v15, v4)
                var v30 uint = 296
                fmt.Println(v15, v2)
                _ = v30
            }
            _ = v29
        } else {
            var v31 uint16 = (v28 + 0)
            for fi32 := 0; fi32 < 20; fi32++ {
                v2 += 925
                for fi33 := 0; fi33 < 18; fi33++ {
                    var v34 float32 = (0.1 - 1e10)
                    var v35 uint64 = 18446744073709551615
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                _ = fi32
            }
            _ = v31
        }
        _ = v27
        _ = v28
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v36 bool = (0.0 >= -2.25)
    switch 4294967295 {
    case -536:
        for fi37 := 0; fi37 < 2; fi37++ {
            v4 = int8(v2)
            var v38 float64 = -2.25
            _ = fi37
            _ = v38
        }
    case 1:
        switch genFn1((-(v4)), -540) {
        case -1:
            fmt.Println(v36, v2, v36)
        case 173:
            if true {
                var v39 uint32 = v2
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v4, v15, v2)
                _ = v39
            } else {
                v15 = (((0 >= 18446744073709551615) || v36) && false)
                for fi40 := 0; fi40 < 4; fi40++ {
                    var v41 bool = v15
                    var v42 uint64 = 1
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                switch "a" {
                case "foo bar":
                    v15 = (0 > 2)
                case "":
                    for fi43 := 0; fi43 < 18; fi43++ {
                        v4 -= int8(fi43)
                        v15 = (!v15)
                        fmt.Println(v15, v36)
                        _ = fi43
                    }
                }
            }
        default:
            var v44 int32 = -924
            _ = v44
        }
    default:
        v45 := GenS4{F0: -632, F1: -310, F2: -1}
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
