package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0(p0 bool, p1 uint16) uint {
    var v6 float32 = 1.5
    fmt.Println(p1, p0)
    switch -434 {
    case 2147483647:
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 15; fi8++ {
                v6 -= ((-2.25 + -2.25) * v6)
                v6 = v6
                _ = fi8
            }
            if false {
                if (1 > 65535) {
                    var v9 uint64 = 1
                    if (p0 && ("" < ("" + "a"))) {
                        for fi10 := 0; fi10 < 14; fi10++ {
                            var v11 int64 = 1
                            _ = fi10
                            _ = v11
                        }
                    } else {
                        if (64 > 1) {
                            fmt.Println(p0)
                        }
                        var v12 int64 = -1
                        v6 = (v6 + 1e10)
                        _ = v12
                    }
                    v13 := []int{(fi7 - -653), 2147483647, (fi7 - int(p1)), fi7}
                    _ = v13
                    _ = v9
                } else {
                    for fi14 := 0; fi14 < 6; fi14++ {
                        v15 := GenS1{F0: 0, F1: 255}
                        var v16 int32 = -545
                        v16 = (v16 & (v16 - (v16 + (v16 ^ v16))))
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    p1 = 255
                    p0 = (2 == 0)
                }
            } else {
                switch ("\n\t" + "") {
                case "foo bar":
                    switch (^(fi7)) {
                    case -1:
                        switch fi7 {
                        case -694:
                            var v17 bool = (2 >= 155)
                            _ = v17
                        default:
                            v6 = v6
                        }
                    default:
                        fmt.Println(fi7, p0)
                    }
                case "hello":
                    if p0 {
                        if p0 {
                            v6 = v6
                            for fi18 := 0; fi18 < 3; fi18++ {
                                v19 := []int{-9223372036854775808}
                                _ = fi18
                                _ = v19
                            }
                        }
                        if p0 {
                            if p0 {
                                var v20 bool = (45 <= 861)
                                _ = v20
                            } else {
                                var v21 uint = 18446744073709551615
                                _ = v21
                            }
                            v22 := GenS2{F0: -32, F1: 465}
                            p1 = 90
                            _ = v22
                        }
                    } else {
                        v6 = (3.14159 / float32(-2.25))
                        p1 = p1
                    }
                case "\n\t":
                    v23 := []int{fi7, fi7, fi7, (fi7 & (fi7 ^ (fi7 / ((fi7) | 1))))}
                    _ = v23
                }
                var v24 float32 = 3.14159
                _ = v24
            }
            v6 += 1.5
            _ = fi7
        }
    default:
        if p0 {
            var v25 int = 4294967295
            fmt.Println(p1)
            _ = v25
        } else {
            if ((!(("\n\t" > "\n\t") && p0)) && false) {
                v26 := GenS2{F0: -1, F1: 2}
                switch 9223372036854775807 {
                case -656:
                    p0 = (827 < 0)
                case -674:
                    var v27 bool = p0
                    _ = v27
                }
                _ = v26
            } else {
                v28 := GenS3{F0: -1, F1: 65535}
                fmt.Println(v6, p1, p0)
                _ = v28
            }
        }
    }
    p0 = p0
    p1 += 32767
    _ = p0
    _ = p1
    _ = v6
    return 255
}
func genFn1(p0 bool, p1 int64, p2 int8) uint {
    var v29 int16 = -533
    var v30 int32 = -2147483648
    p0 = false
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v30
    return 494
}

func main() {
    var v1 int8 = 1
    var v2 float64 = ((1e10 * float64(-2.25)) * 3.14159)
    var v3 int = -1
    var v4 int8 = v1
    v5 := []int{v3}
    _ = len(v5)
    v4 = (v1 ^ (v1 >> ((65535) & 63)))
    v1 -= (^(v1))
    var v31 string = "\n\t"
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 int64 = 1
        v34 := GenS1{F0: genFn1((42 > 65535), (-(v33)), v1), F1: 1}
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println(v3, v3)
    fmt.Println(v4)
    var v35 float32 = -2.25
    fmt.Println(v2, v31)
    fmt.Println(v35)
    var v36 int32 = -1
    v36 += 2147483647
    v37 := GenS1{F0: 175, F1: 1}
    v36 -= v36
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
