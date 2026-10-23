package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }

func main() {
    var v1 string = "a"
    var v2 float32 = 1e10
    var v3 uint8 = 128
    var v4 uint16 = 255
    var v5 uint64 = 960
    v6 := GenS1{F0: 0}
    var v7 int32 = -1
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            fmt.Println(v3)
            if false {
                v10 := GenS2{F0: (v7 * 42)}
                fmt.Println(v4, fi9)
                _ = v10
            } else {
                var v11 int16 = 32767
                v12 := []int{fi8, fi8}
                _ = len(v12)
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        switch (v1 + ("" + "a")) {
        case "":
            v2 = ((v2 / v2) * float32((-2.25 / (1e10 + 3.14159))))
        default:
            if (v2 > (1.5 + v2)) {
                var v13 uint = 466
                v14 := GenS2{F0: (^(v7))}
                _ = v13
                _ = v14
            } else {
                if false {
                    var v15 bool = (42 != v3)
                    var v16 int = fi8
                    switch (v16 * (v16 << ((436) & 63))) {
                    case -419:
                        var v17 float64 = 0.0
                        _ = v17
                    case 128:
                        switch -483 {
                        case 90:
                            if v15 {
                                if (3.14159 >= float64(v2)) {
                                    v18 := []int{(v16 - fi8), 32767, 42}
                                    v5 -= (v5 << ((329) & 63))
                                    if (884 > -1) {
                                        v5 = (v5 * v5)
                                    } else {
                                        v1 = "\n\t"
                                        _ = len(v18)
                                        var v19 int = v16
                                        _ = v19
                                    }
                                    _ = v18
                                } else {
                                    fmt.Println(v4, fi8, v1)
                                    v20 := []int{(fi8 * v16)}
                                    for fi21 := 0; fi21 < 4; fi21++ {
                                        v15 = (v3 >= v3)
                                        _ = fi21
                                    }
                                    _ = v20
                                }
                            } else {
                                v3 += (v3 % (((^(v3))) | 1))
                                var v22 uint32 = uint32(v3)
                                _ = v22
                            }
                        case -1:
                            fmt.Println(v16, fi8)
                        case 596:
                            switch fi8 {
                            case 128:
                                if ((v5 * v5) > uint64(v3)) {
                                    var v23 int8 = -128
                                    if (v4 == v4) {
                                        v23 = v23
                                        if ((v1 + (v1 + (v1 + v1))) < (v1 + ((("" + v1) + (v1 + "a")) + (v1 + (v1 + v1))))) {
                                            v24 := GenS2{F0: (-(v7))}
                                            v2 = (v2 * (float32(float64(float32(float64(v2)))) - (v2 - (-2.25 / ((v2 / v2) + v2)))))
                                            _ = v24
                                        } else {
                                            var v25 int8 = v23
                                            v26 := GenS1{F0: 42}
                                            _ = v25
                                            _ = v26
                                        }
                                    } else {
                                        var v27 int8 = (v23 - (v23 % (((v23 + (v23 + int8(v3)))) | 1)))
                                        _ = v27
                                    }
                                    _ = v23
                                } else {
                                    fmt.Println(v2, v7, v1)
                                    for fi28 := 0; fi28 < 20; fi28++ {
                                        v4 -= v4
                                        _ = fi28
                                    }
                                    v29 := []int{fi8}
                                    _ = v29
                                }
                            case -919:
                                v1 = ("foo bar" + "foo bar")
                            default:
                                fmt.Println(fi8)
                            }
                        default:
                            v30 := GenS2{F0: (v7 << ((2) & 63))}
                            _ = v30
                        }
                    default:
                        v2 = 0.0
                    }
                    _ = v15
                    _ = v16
                }
                v5 -= 18446744073709551615
            }
        }
        _ = fi8
    }
    v4 -= (v4 >> ((2) & 63))
    var v31 float64 = 0.0
    for fi32 := 0; fi32 < 12; fi32++ {
        fmt.Println(v4, v2, v4)
        _ = fi32
    }
    v33 := GenS2{F0: 873}
    v34 := GenS2{F0: (v7 + (v7 ^ -1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
