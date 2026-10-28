package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint16 }
type GenS2 struct { F0 int }
func genFn0() bool {
    var v7 string = "a"
    var v8 uint16 = 32767
    var v9 uint8 = 2
    v7 = v7
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 int) uint16 {
    v10 := GenS1{F0: 65535, F1: 32768, F2: 2}
    var v11 uint64 = 44
    v11 = v11
    _ = p0
    _ = v10
    _ = v11
    return 255
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    v3 := []uint32{42, 4294967295, 65535}
    var v4 uint32 = 42
    var v5 int64 = 0
    v6 := make(map[int]uint32)
    var v12 float32 = ((1e10 * (1e10 * (1e10 + (0.1 - 1e10)))) - (0.0 - 0.1))
    var v13 uint8 = 127
    v12 += float32((1.5 - float64((((0.1 / (v12 / v12)) / v12) - float32(1e10)))))
    if false {
        for fi14 := 0; fi14 < 7; fi14++ {
            v12 = 0.0
            switch "" {
            case "\n\t":
                v15 := GenS1{F0: genFn1(fi14), F1: fi14, F2: 32767}
                _ = v15
            case "a":
                v6[(fi14 << ((4294967295) & 63))] = (v4 | v4)
            case "":
                switch "a" {
                case "foo bar":
                    if false {
                        v16 := GenS1{F0: 1, F1: int(v13), F2: 810}
                        _ = v16
                    }
                default:
                    var v17 uint = 65535
                    _ = v17
                }
            default:
                switch ("hello" + "hello") {
                case "foo bar":
                    if ((false && false) || (42 == 1)) {
                        if false {
                            v5 = 9223372036854775807
                        }
                        var v18 int32 = -1
                        v13 -= (v13 ^ v13)
                        _ = v18
                    }
                }
            }
            switch (fi14 | fi14) {
            case -608:
                v19 := GenS2{F0: fi14}
                _ = v19
            case 2147483647:
                if ((v5 % (((v5 % (((v5 >> ((4294967295) & 63))) | 1))) | 1)) < v5) {
                    v5 += (v5 & (v5 / ((9223372036854775807) | 1)))
                    if false {
                        fmt.Println(v4)
                        var v20 uint64 = 0
                        _ = v20
                    }
                    fmt.Println(v13, fi14)
                } else {
                    v12 -= v12
                    v4 += 0
                }
            default:
                var v21 uint32 = 4294967295
                _ = v21
            }
            _ = fi14
        }
    } else {
        for fi22 := 0; fi22 < 11; fi22++ {
            switch fi22 {
            case -598:
                switch fi22 {
                case -2147483648:
                    if false {
                        for fi23 := 0; fi23 < 14; fi23++ {
                            var v24 int16 = int16(v4)
                            v25 := GenS1{F0: 255, F1: fi23, F2: 65535}
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        fmt.Println(fi22, fi22)
                    }
                case 910:
                    fmt.Println(v5)
                }
            case -1:
                if len(v3) > 0 {
                    v3[0] = uint32(v13)
                }
            }
            var v26 uint = 65535
            switch (((("\n\t" + "\n\t") + "foo bar") + (((("a" + "a") + ("foo bar" + "foo bar")) + "\n\t") + ((("\n\t" + "\n\t") + ("\n\t" + "foo bar")) + ("a" + "\n\t")))) + "\n\t") {
            case "a":
                if false {
                    if false {
                        if len(v3) > 0 {
                            v3[0] = v4
                        }
                    }
                    delete(v2, fi22)
                } else {
                    if (-17 < -1) {
                        var v27 string = "hello"
                        v4 = (v4 ^ 4294967295)
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 1; fi28++ {
                            var v29 uint = 42
                            v4 = 42
                            v13 = uint8(v26)
                            _ = fi28
                            _ = v29
                        }
                        var v30 uint32 = 39
                        _ = v30
                    }
                    _ = len(v1)
                }
            }
            _ = fi22
            _ = v26
        }
        v31 := GenS2{F0: 1}
        _ = v31
    }
    var v32 bool = false
    v33 := GenS2{F0: 472}
    var v34 uint64 = uint64(v4)
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 int16 = -1
        var v37 uint8 = (v13 >> ((18446744073709551615) & 63))
        for fi38 := 0; fi38 < 6; fi38++ {
            var v39 uint = 18446744073709551615
            if (false || ((^(v36)) == (^(v36)))) {
                v40 := GenS2{F0: -509}
                _ = v40
            }
            _ = fi38
            _ = v39
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    var v41 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
