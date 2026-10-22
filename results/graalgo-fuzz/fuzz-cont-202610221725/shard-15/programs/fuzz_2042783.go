package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int16, p1 int) uint32 {
    p0 = 944
    if false {
        v6 := GenS1{F0: 32767}
        var v7 uint = 766
        v8 := []int{p1, (p1 >> (((v7 / ((v7) | 1))) & 63)), p1}
        _ = v6
        _ = v7
        _ = v8
    } else {
        switch "hello" {
        case "foo bar":
            for fi9 := 0; fi9 < 10; fi9++ {
                p0 = (p0 & p0)
                _ = fi9
            }
        case "a":
            var v10 int64 = int64(p1)
            _ = v10
        case "hello":
            for fi11 := 0; fi11 < 1; fi11++ {
                for fi12 := 0; fi12 < 12; fi12++ {
                    p0 -= p0
                    _ = fi12
                }
                _ = fi11
            }
        default:
            if true {
                v13 := []int{p1, (p1 & (^(p1))), p1}
                for fi14 := 0; fi14 < 18; fi14++ {
                    p0 = p0
                    var v15 bool = false
                    _ = fi14
                    _ = v15
                }
                v16 := GenS1{F0: p1}
                _ = v13
                _ = v16
            }
        }
        v17 := []int{972, 32767}
        _ = v17
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 uint = 255
    var v2 uint = v1
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    var v5 int8 = -128
    v18 := GenS1{F0: 977}
    if true {
        if true {
            var v19 uint = (v2 * 1)
            switch ((("" + "foo bar") + ("" + "hello")) + "\n\t") {
            case "\n\t":
                v19 += v1
            default:
                fmt.Println(v2, v1)
            }
            var v20 int32 = 2147483647
            _ = v19
            _ = v20
        }
        var v21 uint64 = 1
        for fi22 := 0; fi22 < 5; fi22++ {
            for fi23 := 0; fi23 < 17; fi23++ {
                v24 := GenS1{F0: int(v2)}
                if false {
                    switch 32767 {
                    case 32768:
                        switch 455 {
                        case -1:
                            v25 := GenS1{F0: (fi22 | fi22)}
                            _ = v25
                        case 9223372036854775807:
                            if (fi22 <= fi22) {
                                v26 := GenS1{F0: fi22}
                                _ = v26
                            } else {
                                if false {
                                    fmt.Println(v2, v2)
                                } else {
                                    if (false || (409 < uint16(v21))) {
                                        if true {
                                            var v27 int32 = 664
                                            delete(v3, (fi22 / ((fi22) | 1)))
                                            _ = v27
                                        }
                                    }
                                    _ = len(v3)
                                }
                            }
                        case 0:
                            delete(v3, 4294967295)
                        default:
                            v28 := v4[(("a" + ("foo bar" + "")) + "foo bar")]
                            _ = v28
                        }
                        var v29 int = (fi23 * fi23)
                        v30 := []int{fi22, (v29 / ((42) | 1)), (v29 + fi23)}
                        _ = v29
                        _ = v30
                    case 2147483647:
                        v2 = v2
                    case 32767:
                        var v31 float64 = 0.1
                        _ = v31
                    default:
                        v32 := []int{2147483647, (fi23 * fi22), fi23, (fi22 / (((fi23 - fi23)) | 1))}
                        _ = v32
                    }
                    var v33 float64 = float64(((3.14159 + 3.14159) + 0.1))
                    var v34 uint16 = 941
                    _ = v33
                    _ = v34
                }
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        _ = v21
    } else {
        if (!true) {
            v35 := []int{-292}
            for fi36 := 0; fi36 < 5; fi36++ {
                var v37 float64 = ((3.14159 * ((3.14159 * (1e10 + 0.0)) * -2.25)) * 3.14159)
                if len(v35) >= 2 {
                    v35 = v35[1:len(v35)]
                }
                _ = fi36
                _ = v37
            }
            _ = v35
        }
        for fi38 := 0; fi38 < 12; fi38++ {
            var v39 float64 = ((0.0 + -2.25) / 1.0)
            fmt.Println(v5, v39)
            _ = fi38
            _ = v39
        }
    }
    var v40 int32 = 42
    v5 = (v5 - -1)
    if (9223372036854775807 >= 403) {
        v40 -= v40
        v41 := GenS1{F0: 2147483647}
        _ = v41
    }
    var v42 bool = false
    v2 += v1
    v43 := GenS1{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v5
}
