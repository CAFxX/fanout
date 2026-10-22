package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int16, p1 bool) string {
    var v7 uint = 4294967295
    fmt.Println(p0, p1)
    switch 32767 {
    case -1:
        for fi8 := 0; fi8 < 9; fi8++ {
            if (true && false) {
                v9 := GenS1{F0: (fi8 << ((v7) & 63)), F1: 0}
                if p1 {
                    var v10 bool = (1.5 != 0.0)
                    _ = v10
                } else {
                    if p1 {
                        var v11 bool = (9223372036854775807 > 127)
                        _ = v11
                    }
                    v12 := []int{(fi8 << (((v7 >> ((1) & 63))) & 63)), fi8, 1, 9223372036854775807}
                    _ = v12
                }
                _ = v9
            } else {
                var v13 int8 = -101
                v14 := GenS1{F0: (^(fi8)), F1: -60}
                _ = v13
                _ = v14
            }
            v15 := []int{fi8, -612, -9223372036854775808, (fi8 & (fi8 + fi8))}
            p1 = p1
            _ = fi8
            _ = v15
        }
    case -9223372036854775808:
        if p1 {
            var v16 bool = true
            if v16 {
                switch "" {
                case "\n\t":
                    v7 += v7
                case "foo bar":
                    var v17 int = -2147483648
                    _ = v17
                case "a":
                    switch 456 {
                    case -643:
                        v16 = true
                    }
                }
            }
            fmt.Println(p1, p1)
            _ = v16
        } else {
            for fi18 := 0; fi18 < 19; fi18++ {
                for fi19 := 0; fi19 < 20; fi19++ {
                    p0 = p0
                    _ = fi19
                }
                v20 := []int{fi18, -9223372036854775808}
                var v21 uint8 = 35
                _ = fi18
                _ = v20
                _ = v21
            }
        }
    default:
        for fi22 := 0; fi22 < 10; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                if true {
                    var v24 int = (fi22 & fi22)
                    fmt.Println(v7, p0, v24)
                    p1 = (!(1.5 != 1.5))
                    _ = v24
                }
                if (!(!p1)) {
                    v25 := []int{fi23}
                    var v26 int16 = (-(p0))
                    _ = v25
                    _ = v26
                }
                fmt.Println(fi22, p1, p0)
                if false {
                    v27 := []int{}
                    _ = v27
                } else {
                    switch fi22 {
                    case -412:
                        if (3.14159 < 0.1) {
                            switch (fi23 | (^(fi22))) {
                            case 42:
                                p0 = (p0 << ((4294967295) & 63))
                            case -596:
                                switch "" {
                                case "hello":
                                    v28 := GenS1{F0: -9223372036854775808, F1: 1}
                                    _ = v28
                                default:
                                    var v29 uint32 = 2147483648
                                    _ = v29
                                }
                            case 1:
                                p1 = true
                            }
                        } else {
                            var v30 uint32 = 529
                            _ = v30
                        }
                    }
                }
                v31 := []int{(fi22 & (^(fi22))), 1}
                _ = fi23
                _ = v31
            }
            _ = fi22
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return "hello"
}
func genFn1(p0 uint8, p1 uint16) uint {
    switch 0 {
    case 32768:
        v32 := []int{}
        _ = v32
    case 987:
        v33 := []int{}
        _ = v33
    case 4294967295:
        var v34 int16 = 806
        _ = v34
    default:
        v35 := []int{9223372036854775807, -150}
        _ = v35
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        fmt.Println(p1)
        switch fi36 {
        case 128:
            v37 := GenS1{F0: fi36, F1: 0}
            _ = v37
        case -362:
            var v38 float64 = ((((1.5 + 1e10) - 1e10) - 0.1) / 3.14159)
            _ = v38
        default:
            var v39 bool = (255 < 378)
            _ = v39
        }
        _ = fi36
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 uint32 = 2
    v2 := []string{"foo bar"}
    v3 := []string{("foo bar" + ("\n\t" + (("" + "\n\t") + ("\n\t" + "foo bar")))), "foo bar", "hello", "foo bar"}
    var v4 uint8 = 0
    var v5 int16 = 1000
    v6 := make(map[string]uint32)
    v4 = 127
    v5 += 0
    for fi40 := 0; fi40 < 5; fi40++ {
        fmt.Println(v4, v5)
        _ = fi40
    }
    if len(v3) > 0 {
        v3[0] = ""
    }
    _ = len(v3)
    if true {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        _ = len(v6)
        v1 += v1
    } else {
        v4 -= v4
    }
    var v41 float64 = float64((0.0 + 0.0))
    if (((65535 < uint32(v4)) || (100 < -64)) && true) {
        fmt.Println(v41, v41)
        v4 += (v4 << ((42) & 63))
    } else {
        v5 += v5
        if ((true && false) && (!false)) {
            var v42 string = "foo bar"
            _ = v42
        }
        _ = len(v2)
    }
    v43 := GenS1{F0: 721, F1: -1}
    v44 := GenS2{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
