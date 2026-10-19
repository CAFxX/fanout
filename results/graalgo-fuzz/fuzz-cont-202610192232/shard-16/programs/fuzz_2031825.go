package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint8, p1 bool, p2 uint) uint8 {
    switch 135 {
    case 42:
        switch ("" + "") {
        case "foo bar":
            var v4 bool = (71 <= 42)
            _ = v4
        case "":
            switch 270 {
            case 800:
                v5 := []int{-2147483648, 1}
                _ = v5
            }
        }
    }
    v6 := []int{961, 32767, 4294967295}
    fmt.Println(p0, p0, p2)
    _ = len(v6)
    var v7 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (^(p0))
}

func main() {
    var v1 int8 = -128
    var v2 uint32 = 65535
    v3 := make(map[uint32]uint32)
    switch "" {
    case "":
        if true {
            var v8 int = 9223372036854775807
            v2 = (v2 * 42)
            var v9 string = "a"
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: (-(v1))}
            v11 := v3[(v2 & v2)]
            _ = v10
            _ = v11
        }
    default:
        var v12 string = ("hello" + "foo bar")
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v2 += (v2 * (v2 << ((255) & 63)))
        var v14 string = ((("hello" + ("\n\t" + "hello")) + "hello") + "")
        fmt.Println(fi13)
        _ = fi13
        _ = v14
    }
    switch "" {
    case "foo bar":
        switch 2147483647 {
        case -9223372036854775808:
            if (!false) {
                for fi15 := 0; fi15 < 1; fi15++ {
                    if (2 < 42) {
                        var v16 uint8 = uint8(v2)
                        var v17 int32 = 1
                        if true {
                            for fi18 := 0; fi18 < 9; fi18++ {
                                fmt.Println(fi18)
                                v19 := GenS1{F0: 13}
                                _ = fi18
                                _ = v19
                            }
                        } else {
                            fmt.Println(v2, v1)
                            delete(v3, (v2 ^ v2))
                            if (!true) {
                                for fi20 := 0; fi20 < 9; fi20++ {
                                    var v21 int16 = 32767
                                    if true {
                                        fmt.Println(v1, v21, fi20)
                                        v22 := GenS1{F0: 127}
                                        fmt.Println(v16, v16)
                                        _ = v22
                                    }
                                    _ = fi20
                                    _ = v21
                                }
                            } else {
                                v1 -= int8(v2)
                                var v23 float32 = (3.14159 * -2.25)
                                _ = v23
                            }
                        }
                        _ = v16
                        _ = v17
                    } else {
                        if (false || false) {
                            v24 := GenS1{F0: 0}
                            if (!false) {
                                var v25 int16 = -32768
                                _ = v25
                            } else {
                                var v26 int8 = v1
                                _ = v26
                            }
                            _ = v24
                        } else {
                            var v27 int32 = 389
                            v28 := v3[651]
                            v29 := GenS1{F0: (v1 * v1)}
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        v30 := v3[2147483648]
                        var v31 uint64 = 0
                        _ = v30
                        _ = v31
                    }
                    var v32 int64 = -1
                    _ = fi15
                    _ = v32
                }
            } else {
                v1 -= (-(v1))
            }
        case 2147483647:
            var v33 string = ("" + "hello")
            _ = v33
        default:
            if (0 >= 255) {
                if false {
                    var v34 uint8 = 10
                    _ = v34
                } else {
                    v35 := GenS1{F0: v1}
                    for fi36 := 0; fi36 < 19; fi36++ {
                        v1 -= (v1 & v1)
                        v3[693] = 42
                        v37 := GenS1{F0: 42}
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                }
            } else {
                v1 = 42
                for fi38 := 0; fi38 < 3; fi38++ {
                    delete(v3, v2)
                    var v39 int32 = -1
                    _ = fi38
                    _ = v39
                }
                for fi40 := 0; fi40 < 1; fi40++ {
                    delete(v3, 0)
                    _ = fi40
                }
            }
        }
    case "":
        v2 += v2
    default:
        _ = len(v3)
    }
    _ = len(v3)
    fmt.Println(v2, v1)
    v2 = 0
    _ = len(v3)
    v41 := []int{}
    v1 = v1
    v2 += v2
    for fi42 := 0; fi42 < 20; fi42++ {
        fmt.Println(v2)
        switch (fi42 / ((fi42) | 1)) {
        case -9223372036854775808:
            v1 += v1
        }
        v43 := GenS1{F0: v1}
        _ = fi42
        _ = v43
    }
    v44 := GenS1{F0: v1}
    v45 := GenS1{F0: v1}
    v1 += v1
    v46 := GenS1{F0: v1}
    for fi47 := 0; fi47 < 1; fi47++ {
        v1 = (^(v1))
        _ = fi47
    }
    v1 -= -54
    v2 -= 461
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v44
    _ = v45
    _ = v46
}
