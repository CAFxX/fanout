package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int; F1 uint; F2 int8 }
func genFn0(p0 int32) int {
    var v4 bool = false
    var v5 uint = 870
    for fi6 := 0; fi6 < 15; fi6++ {
        v4 = (false && ((773 > 4294967295) && (false && (false || ((true && v4) || false)))))
        var v7 uint64 = 860
        var v8 uint = uint(v7)
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 uint16 = 667
    v2 := make(map[string]uint32)
    var v3 string = "a"
    fmt.Println(v1, v1)
    v3 = ""
    v9 := []int{}
    delete(v2, "hello")
    if false {
        delete(v2, ((v3 + "a") + v3))
        switch 2147483647 {
        case 128:
            v3 = v3
        case 32767:
            v10 := GenS1{F0: 335}
            _ = v10
        case -942:
            var v11 uint64 = 2
            _ = v11
        default:
            if false {
                var v12 uint8 = 2
                _ = v12
            } else {
                if (false && (1 < 800)) {
                    var v13 int16 = 32767
                    for fi14 := 0; fi14 < 1; fi14++ {
                        _ = len(v9)
                        for fi15 := 0; fi15 < 6; fi15++ {
                            v2[(("hello" + ("\n\t" + "\n\t")) + "\n\t")] = 0
                            _ = fi15
                        }
                        for fi16 := 0; fi16 < 13; fi16++ {
                            v9 = append(v9, fi14)
                            v3 = "foo bar"
                            if (!false) {
                                var v17 int16 = (v13 / (((^(v13))) | 1))
                                v3 = (v3 + (v3 + (((v3 + "foo bar") + v3) + "a")))
                                v3 = (v3 + v3)
                                _ = v17
                            }
                            _ = fi16
                        }
                        _ = fi14
                    }
                    _ = v13
                }
                v18 := GenS2{F0: 65535, F1: 697}
                _ = v18
            }
        }
    } else {
        if (((v1 <= 441) || true) && true) {
            if false {
                v1 = (v1 % (((v1 / ((v1) | 1))) | 1))
                if false {
                    fmt.Println(v3, v1)
                } else {
                    var v19 int8 = 21
                    for fi20 := 0; fi20 < 10; fi20++ {
                        if true {
                            _ = len(v9)
                            if (true || (4294967295 != 255)) {
                                v21 := GenS1{F0: 844}
                                v22 := GenS3{F0: fi20, F1: 65535, F2: (v19 << ((785) & 63))}
                                _ = v21
                                _ = v22
                            }
                            v19 = v19
                        }
                        _ = fi20
                    }
                    _ = v19
                }
            }
        }
        switch v3 {
        case "hello":
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := GenS2{F0: 2, F1: 2}
                v1 -= v1
                _ = len(v9)
                _ = fi23
                _ = v24
            }
        case "":
            v3 = v3
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    fmt.Println(v1, v1)
    var v25 float64 = 1.5
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 int = -1
        var v28 int16 = -935
        _ = fi26
        _ = v27
        _ = v28
    }
    v25 -= float64(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v9
}
