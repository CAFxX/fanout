package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
func genFn0(p0 int, p1 uint64, p2 int64) int {
    v5 := []int{(-(p0)), (p0 * 127)}
    v6 := GenS1{F0: p0}
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: (p0 / ((-523) | 1))}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (^(p0))
}

func main() {
    v1 := []uint32{1, 517, 346}
    v2 := make(map[int]string)
    var v3 uint16 = 255
    var v4 int = -651
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint32 = 724
        v10 = (v10 & v10)
        _ = len(v1)
        _ = fi9
        _ = v10
    }
    v3 -= 674
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 int32 = -178
    v2[v4] = ("" + "foo bar")
    var v12 uint16 = v3
    var v13 string = "\n\t"
    var v14 uint32 = 2147483647
    if (18446744073709551615 >= 731) {
        for fi15 := 0; fi15 < 12; fi15++ {
            for fi16 := 0; fi16 < 18; fi16++ {
                v17 := GenS1{F0: (fi16 + genFn0(42, 0, 1))}
                v18 := v2[genFn0(844, 42, 1)]
                v19 := GenS1{F0: 2147483647}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            fmt.Println(v12, v4, v13)
            _ = fi15
        }
        v20 := GenS1{F0: int(v3)}
        switch v13 {
        case "foo bar":
            if true {
                var v21 float32 = ((0.0 + 0.0) * 1.5)
                for fi22 := 0; fi22 < 20; fi22++ {
                    v14 = 923
                    v2[(v4 - (-(fi22)))] = v13
                    v23 := GenS2{F0: (("a" > ("foo bar" + v13)) || true)}
                    _ = fi22
                    _ = v23
                }
                _ = v21
            }
        case "a":
            v24 := GenS2{F0: false}
            _ = v24
        case "":
            v4 = (v4 & genFn0(genFn0((v4 * (v4 & -1)), 18446744073709551615, 1), 0, -740))
        default:
            fmt.Println(v12, v13)
        }
        _ = v20
    }
    if false {
        v14 = v14
    } else {
        v14 += (v14 / (((v14 % (((v14 + v14)) | 1))) | 1))
        for fi25 := 0; fi25 < 15; fi25++ {
            if (1 < -14) {
                fmt.Println(v13, v4)
                switch "\n\t" {
                case "a":
                    v14 = v14
                default:
                    v3 = 42
                }
            } else {
                if (0 > 0) {
                    v26 := GenS2{F0: true}
                    var v27 uint = 4294967295
                    fmt.Println(v13, v13, v14)
                    _ = v26
                    _ = v27
                }
                switch (v13 + "foo bar") {
                case "foo bar":
                    v11 -= 2147483647
                case "a":
                    for fi28 := 0; fi28 < 18; fi28++ {
                        var v29 string = (v13 + ((v13 + v13) + ""))
                        switch fi25 {
                        case 32768:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case -733:
                            fmt.Println(fi28, v4, v4)
                        default:
                            v2[fi25] = "\n\t"
                        }
                        _ = fi28
                        _ = v29
                    }
                case "":
                    var v30 int = genFn0(v4, 115, -1)
                    _ = v30
                }
            }
            _ = fi25
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
