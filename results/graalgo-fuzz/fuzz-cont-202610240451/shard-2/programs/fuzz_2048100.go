package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint8; F1 uint; F2 int32 }
func genFn0(p0 int16, p1 uint) int16 {
    v6 := []int{508, -514}
    v6 = append(v6, 127)
    v7 := GenS1{F0: 531, F1: (p0 / ((32767) | 1))}
    v8 := GenS2{F0: 1, F1: true}
    for fi9 := 0; fi9 < 8; fi9++ {
        if len(v6) > 0 {
            v6[0] = (fi9 << ((p1) & 63))
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 uint8 = 160
    v2 := []int{32767, -2147483648}
    v3 := []int{-263, 4294967295, -2147483648, 431}
    var v4 int = -2147483648
    var v5 string = "foo bar"
    v10 := GenS1{F0: (-(v4)), F1: 1000}
    _ = len(v2)
    switch (v5 + "\n\t") {
    case "foo bar":
        var v11 int16 = 1
        _ = v11
    }
    if (((1e10 / ((1.5 - 0.0) / (1e10 / float32(float64(3.14159))))) / float32(-2.25)) != 3.14159) {
        switch "" {
        case "foo bar":
            var v12 bool = false
            _ = v12
        case "hello":
            var v13 uint16 = 32767
            _ = v13
        case "a":
            switch (v4 >> ((398) & 63)) {
            case 0:
                var v14 int32 = 2147483647
                _ = v14
            case 4294967295:
                v1 += 2
            }
        }
        var v15 string = ""
        _ = v15
    }
    if false {
        if (!(false || false)) {
            fmt.Println(v4, v4, v5)
        }
        if ((3.14159 / 0.1) < 0.1) {
            v4 = -881
            if (v4 > (^(v4))) {
                if (-265 > 467) {
                    var v16 uint64 = 973
                    if len(v2) > 0 {
                        v2[0] = 32767
                    }
                    _ = v16
                } else {
                    _ = len(v3)
                }
                v17 := GenS2{F0: 688, F1: true}
                _ = v17
            } else {
                v4 -= (v4 * (-(v4)))
                if true {
                    v1 = v1
                    v2 = append(v2, (v4 >> ((1) & 63)))
                } else {
                    v1 -= (v1 % (((v1 >> ((255) & 63))) | 1))
                    var v18 float64 = 1.5
                    _ = v18
                }
            }
            var v19 uint32 = 408
            _ = v19
        } else {
            var v20 uint32 = 2147483647
            _ = v20
        }
        v1 = 42
    } else {
        if false {
            for fi21 := 0; fi21 < 2; fi21++ {
                switch v4 {
                case 128:
                    for fi22 := 0; fi22 < 11; fi22++ {
                        v2 = append(v2, fi21)
                        _ = fi22
                    }
                case 42:
                    for fi23 := 0; fi23 < 15; fi23++ {
                        fmt.Println(v4)
                        v5 = v5
                        if false {
                            var v24 int16 = 1
                            if false {
                                v1 -= 127
                            }
                            fmt.Println(v4, v5)
                            _ = v24
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v1 = v1
                        }
                        v25 := GenS1{F0: (-(fi23)), F1: 0}
                        fmt.Println(fi23)
                        fmt.Println(v4)
                        _ = fi23
                        _ = v25
                    }
                }
                fmt.Println(v5)
                _ = fi21
            }
            for fi26 := 0; fi26 < 2; fi26++ {
                v5 = (v5 + v5)
                var v27 uint64 = 4294967295
                var v28 uint8 = v1
                _ = fi26
                _ = v27
                _ = v28
            }
            v3 = append(v3, 42)
        }
        for fi29 := 0; fi29 < 8; fi29++ {
            v30 := GenS3{F0: 1, F1: 42}
            for fi31 := 0; fi31 < 6; fi31++ {
                v1 += v1
                _ = fi31
            }
            var v32 float64 = 0.0
            _ = fi29
            _ = v30
            _ = v32
        }
    }
    if false {
        if false {
            fmt.Println(v1, v4)
        } else {
            for fi33 := 0; fi33 < 16; fi33++ {
                if false {
                    var v34 float64 = (1.5 + (1.5 - (float64(1e10) * 0.1)))
                    var v35 int32 = 0
                    _ = v34
                    _ = v35
                } else {
                    var v36 uint = 42
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v36
                }
                _ = fi33
            }
        }
        v37 := GenS4{F0: 140, F1: 650, F2: 0}
        _ = v37
    } else {
        var v38 int8 = -128
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v38
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        var v40 uint32 = 2
        v40 -= (v40 | v40)
        v4 = (fi39 >> ((2) & 63))
        _ = fi39
        _ = v40
    }
    fmt.Println(v1, v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
