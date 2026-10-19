package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int8; F1 uint8; F2 string }
type GenS3 struct { F0 uint8; F1 uint; F2 int8 }
func genFn0(p0 uint8, p1 int8) int16 {
    for fi7 := 0; fi7 < 3; fi7++ {
        switch (fi7 & -388) {
        case 42:
            for fi8 := 0; fi8 < 4; fi8++ {
                if (p0 != (p0 - p0)) {
                    v9 := []int{(^(fi7)), (fi8 | fi8), fi7, 649}
                    p0 -= p0
                    var v10 string = "foo bar"
                    _ = v10
                    _ = v9
                }
                var v11 float64 = 0.1
                _ = fi8
                _ = v11
            }
        }
        if (true && false) {
            var v12 uint64 = 1
            _ = v12
        }
        p0 = p0
        _ = fi7
    }
    p0 += p0
    p0 = (p0 & p0)
    _ = p0
    _ = p1
    return 32767
}
func genFn1(p0 uint32, p1 int16) int {
    if ((0 != 690) || ((!true) && false)) {
        v13 := GenS1{F0: 78, F1: "foo bar"}
        for fi14 := 0; fi14 < 4; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                if (!(!(0 == 2147483647))) {
                    if ((!true) && true) {
                        var v16 int64 = -706
                        var v17 int64 = (v16 ^ v16)
                        switch fi14 {
                        case 550:
                            switch 757 {
                            case 1:
                                fmt.Println(v16, v17, p1)
                            }
                        }
                        var v18 int8 = 1
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        var v19 uint8 = 50
                        v20 := GenS1{F0: -2147483648, F1: "\n\t"}
                        _ = v19
                        _ = v20
                    }
                }
                fmt.Println(p1, fi15, p0)
                _ = fi15
            }
            _ = fi14
        }
        var v21 int64 = 9223372036854775807
        _ = v13
        _ = v21
    } else {
        for fi22 := 0; fi22 < 13; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                fmt.Println(p0, p0)
                p1 = genFn0(42, -1)
                v24 := []int{}
                _ = fi23
                _ = v24
            }
            var v25 int64 = -1
            _ = fi22
            _ = v25
        }
        for fi26 := 0; fi26 < 1; fi26++ {
            v27 := GenS2{F0: 100, F1: 42, F2: "a"}
            p0 = 2147483648
            _ = fi26
            _ = v27
        }
        switch 4294967295 {
        case -2147483648:
            p1 = -716
        }
    }
    p1 = genFn0(255, 42)
    if true {
        fmt.Println(p0, p0)
        v28 := GenS3{F0: 221, F1: 18446744073709551615, F2: 107}
        _ = v28
    }
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    v1 := []uint32{4294967295, 1}
    v2 := []string{"hello"}
    var v3 string = ("hello" + "\n\t")
    v4 := []uint32{65535, 364}
    v5 := make(map[uint32]uint32)
    var v6 int32 = -125
    switch v3 {
    case "":
        v29 := GenS3{F0: 127, F1: 65535, F2: 33}
        _ = v29
    case "a":
        var v30 uint32 = 664
        _ = v30
    case "hello":
        switch 254 {
        case 32767:
            v3 = "\n\t"
        default:
            var v31 string = "\n\t"
            _ = v31
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 4294967295 {
    case 32768:
        var v32 uint64 = 2
        _ = v32
    case -1:
        for fi33 := 0; fi33 < 6; fi33++ {
            v3 = v3
            _ = fi33
        }
    case 81:
        v6 += v6
    }
    _ = len(v4)
    var v34 float64 = (0.1 * 1.5)
    var v35 bool = true
    v36 := v5[376]
    v4 = append(v4, (v36 & v36))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
