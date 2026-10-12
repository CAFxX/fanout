package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint; F2 string }
type GenS4 struct { F0 uint16; F1 int8; F2 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint, p1 uint64, p2 uint) uint64 {
    v7 := GenS1{F0: 32767}
    p1 = 691
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := []int{(fi8 | fi8), (fi8 << ((p0) & 63)), fi8, fi8}
        var v10 uint8 = 2
        for fi11 := 0; fi11 < 9; fi11++ {
            switch "a" {
            case "":
                if ((-2.25 - 0.1) == 1e10) {
                    fmt.Println(p0, fi11)
                } else {
                    v12 := GenS2{F0: (false || true)}
                    var v13 float32 = -2.25
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = v12
                    _ = v13
                }
            case "a":
                v9 = append(v9, fi8)
            case "foo bar":
                var v14 int8 = -76
                _ = v14
            }
            var v15 string = "\n\t"
            if (1 != 0) {
                v9 = append(v9, (fi8 & -418))
                if false {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v16 int8 = -99
                    _ = v16
                }
            }
            _ = fi11
            _ = v15
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v17 := GenS3{F0: 127, F1: 255, F2: ("\n\t" + ("hello" + "\n\t"))}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v7
    return (^(p1))
}

func main() {
    var v1 uint = 255
    v2 := make(map[int]int)
    v3 := []string{"foo bar", "a", ("hello" + ("a" + ("foo bar" + "foo bar")))}
    var v4 uint8 = 42
    var v5 string = "foo bar"
    var v6 uint32 = uint32(v1)
    if false {
        for fi18 := 0; fi18 < 5; fi18++ {
            for fi19 := 0; fi19 < 20; fi19++ {
                v4 += v4
                _ = fi19
            }
            v20 := v2[fi18]
            _ = fi18
            _ = v20
        }
        v21 := GenS4{F0: 886, F1: 100, F2: v1}
        _ = v21
    }
    if (32767 >= -94) {
        fmt.Println(v5)
    } else {
        var v22 bool = true
        v5 = "\n\t"
        _ = v22
    }
    fmt.Println(v5, v5, v4)
    if (1.5 == (((float64(((1.5 - 0.0) / 1e10)) - float64(0.1)) + (((3.14159 + 0.0) + 3.14159) / ((3.14159 - (1e10 + (((((-2.25 - float64(((float32(float64(-2.25)) - (1e10 / 1.0)) * 1e10))) + float64((float32(3.14159) * (float32(3.14159) * 3.14159)))) / float64(1.5)) / 1.0) + 1e10))) + float64(-2.25)))) * 0.1)) {
        for fi23 := 0; fi23 < 12; fi23++ {
            for fi24 := 0; fi24 < 6; fi24++ {
                switch fi24 {
                case -212:
                    v5 = v5
                case 0:
                    v5 = v5
                }
                v25 := GenS2{F0: false}
                _ = fi24
                _ = v25
            }
            v26 := v2[128]
            _ = fi23
            _ = v26
        }
        var v27 uint32 = 462
        v28 := GenS5{F0: false}
        _ = v27
        _ = v28
    }
    if (v4 == v4) {
        var v29 int16 = -635
        if ((32767 < 32767) && false) {
            fmt.Println(v5, v1)
        } else {
            v6 = (v6 + 632)
            var v30 int32 = -1
            _ = v30
        }
        v6 = (v6 & 759)
        _ = v29
    } else {
        var v31 float32 = (3.14159 - 0.0)
        _ = v31
    }
    v1 = 812
    v6 -= uint32(v1)
    v32 := GenS6{F0: 127, F1: 0}
    v1 += (v1 - 0)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
