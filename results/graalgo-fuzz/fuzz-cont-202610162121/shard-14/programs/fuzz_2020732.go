package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint64, p1 int8) uint {
    if true {
        if true {
            fmt.Println(p0, p0)
            if false {
                p1 = p1
            } else {
                var v6 uint16 = 0
                if true {
                    v6 -= (v6 << ((0) & 63))
                    p1 = (^(p1))
                } else {
                    for fi7 := 0; fi7 < 3; fi7++ {
                        if false {
                            if (!(231 != 9223372036854775807)) {
                                var v8 int16 = 42
                                if false {
                                    if (255 == 42) {
                                        if ((true && false) || false) {
                                            v9 := GenS1{F0: v8, F1: 4294967295, F2: (("\n\t" + "a") + (("foo bar" + "") + "\n\t"))}
                                            var v10 uint8 = 1
                                            _ = v10
                                            _ = v9
                                        }
                                    } else {
                                        p1 += p1
                                        v11 := []int{fi7, (fi7 - fi7)}
                                        v12 := GenS2{F0: -46}
                                        _ = v11
                                        _ = v12
                                    }
                                }
                                _ = v8
                            } else {
                                for fi13 := 0; fi13 < 1; fi13++ {
                                    fmt.Println(fi7, p1)
                                    _ = fi13
                                }
                                if false {
                                    for fi14 := 0; fi14 < 12; fi14++ {
                                        var v15 int8 = (p1 - 100)
                                        _ = fi14
                                        _ = v15
                                    }
                                }
                            }
                        }
                        var v16 uint = 0
                        _ = fi7
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 20; fi17++ {
                        for fi18 := 0; fi18 < 4; fi18++ {
                            var v19 float64 = ((0.0 - -2.25) - 0.0)
                            v20 := []int{}
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        v21 := GenS1{F0: 0, F1: 18446744073709551615, F2: "foo bar"}
                        v22 := GenS2{F0: 127}
                        _ = fi17
                        _ = v21
                        _ = v22
                    }
                }
                p0 += (p0 - p0)
                _ = v6
            }
        } else {
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 uint = 42
                _ = fi23
                _ = v24
            }
        }
        p0 -= (^(p0))
    } else {
        for fi25 := 0; fi25 < 12; fi25++ {
            fmt.Println(fi25, fi25)
            v26 := GenS2{F0: (p1 / (((p1 << ((643) & 63))) | 1))}
            _ = fi25
            _ = v26
        }
        v27 := []int{-2147483648}
        _ = v27
    }
    var v28 uint8 = 1
    _ = p0
    _ = p1
    _ = v28
    return 42
}

func main() {
    v1 := []string{"a", ""}
    v2 := make(map[string]uint32)
    var v3 float64 = 3.14159
    var v4 uint64 = 1
    v5 := make(map[string]uint32)
    var v29 string = "foo bar"
    if false {
        v29 = v29
        var v30 uint = uint(v4)
        _ = v30
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        v3 = v3
        v32 := v5[v29]
        v33 := GenS3{F0: 1}
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 string = ""
    for fi35 := 0; fi35 < 13; fi35++ {
        fmt.Println(v3)
        _ = fi35
    }
    var v36 int = 127
    for fi37 := 0; fi37 < 10; fi37++ {
        v3 += ((0.1 / (v3 / float64((1e10 / float32(float64(1e10)))))) + v3)
        _ = fi37
    }
    fmt.Println(v29, v34, v4)
    switch "\n\t" {
    case "foo bar":
        for fi38 := 0; fi38 < 8; fi38++ {
            if true {
                var v39 uint = genFn0(v4, -128)
                var v40 float32 = -2.25
                if (v3 <= (3.14159 - v3)) {
                    for fi41 := 0; fi41 < 8; fi41++ {
                        if true {
                            v4 += 0
                            var v42 float32 = v40
                            v43 := GenS1{F0: 1, F1: v39, F2: (v34 + "a")}
                            _ = v42
                            _ = v43
                        }
                        v34 = ("" + "foo bar")
                        var v44 uint = v39
                        _ = fi41
                        _ = v44
                    }
                }
                _ = v39
                _ = v40
            }
            _ = fi38
        }
    case "":
        v3 += v3
    case "a":
        v45 := GenS3{F0: 18446744073709551615}
        _ = v45
    default:
        fmt.Println(v36)
    }
    v4 -= v4
    fmt.Println(v4, v36, v29)
    v4 = (v4 & (v4 << ((18446744073709551615) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v36
    _ = v4
    _ = v5
}
