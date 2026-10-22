package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int; F1 string }
func genFn0(p0 uint64) uint8 {
    var v4 bool = (!(855 >= 2))
    v5 := GenS1{F0: 979, F1: 127}
    var v6 uint16 = 65535
    var v7 uint64 = (p0 << ((4294967295) & 63))
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 bool, p1 uint16) uint32 {
    for fi8 := 0; fi8 < 19; fi8++ {
        if p0 {
            if p0 {
                var v9 uint16 = 65535
                _ = v9
            } else {
                switch fi8 {
                case 127:
                    fmt.Println(p0)
                }
                switch "hello" {
                case "\n\t":
                    if true {
                        v10 := GenS2{F0: (^(fi8)), F1: ("a" + "a")}
                        _ = v10
                    }
                }
                var v11 float64 = (0.0 * 3.14159)
                _ = v11
            }
            if (false || (42 > -245)) {
                if false {
                    if true {
                        v12 := []int{}
                        p0 = false
                        if ((fi8 + (fi8 / (((fi8 - (^(fi8)))) | 1))) < (fi8 << ((255) & 63))) {
                            if (fi8 == (fi8 & fi8)) {
                                var v13 bool = (!((42 > -1) || (1000 > -32768)))
                                for fi14 := 0; fi14 < 15; fi14++ {
                                    p1 = p1
                                    v13 = false
                                    var v15 uint16 = 255
                                    if true {
                                        p1 += (^(p1))
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                    }
                                    _ = fi14
                                    _ = v15
                                }
                                _ = v13
                            }
                            var v16 uint = 1
                            var v17 int16 = 0
                            _ = v16
                            _ = v17
                        }
                        _ = v12
                    } else {
                        for fi18 := 0; fi18 < 19; fi18++ {
                            p0 = (p0 && p0)
                            p1 -= p1
                            var v19 int32 = -2147483648
                            _ = fi18
                            _ = v19
                        }
                        var v20 float64 = (1.5 / (float64((-2.25 / (0.1 + 1.5))) - -2.25))
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 16; fi21++ {
                        p1 = p1
                        v22 := []int{(fi8 << ((505) & 63))}
                        _ = fi21
                        _ = v22
                    }
                } else {
                    p0 = p0
                    if true {
                        for fi23 := 0; fi23 < 8; fi23++ {
                            var v24 float64 = 0.0
                            _ = fi23
                            _ = v24
                        }
                        var v25 float64 = 3.14159
                        _ = v25
                    } else {
                        v26 := []int{fi8, fi8}
                        _ = v26
                    }
                }
                v27 := []int{}
                _ = v27
            } else {
                fmt.Println(p1, fi8, p0)
                p0 = (42 <= 1)
            }
        }
        fmt.Println(p0, p0, p0)
        _ = fi8
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = -136
    v3 := []uint32{4294967295, 771, 703, 2147483647}
    var v28 bool = ((-1 == -177) && true)
    v28 = (1 != -52)
    v28 = v28
    v2 = v2
    v29 := GenS1{F0: 42, F1: 87}
    var v30 uint8 = 87
    fmt.Println(v30, v28, v28)
    switch -1 {
    case -2147483648:
        v30 = (v30 % ((179) | 1))
    case 51:
        v30 = (^(v30))
    case 42:
        fmt.Println(v2, v28)
    }
    var v31 string = ("\n\t" + "a")
    v31 = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
