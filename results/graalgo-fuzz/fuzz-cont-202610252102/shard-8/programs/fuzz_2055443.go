package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 bool }
func genFn0(p0 uint8, p1 uint8) uint16 {
    fmt.Println(p1, p0)
    p0 = p0
    var v4 uint8 = 127
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    var v1 int32 = -327
    v2 := make(map[string]int)
    var v3 int8 = 127
    if true {
        for fi5 := 0; fi5 < 16; fi5++ {
            v6 := []int{}
            v7 := GenS1{F0: 2}
            _ = fi5
            _ = v6
            _ = v7
        }
        fmt.Println(v1)
    }
    v1 += v1
    switch "foo bar" {
    case "foo bar":
        for fi8 := 0; fi8 < 1; fi8++ {
            switch (fi8 ^ fi8) {
            case -1:
                for fi9 := 0; fi9 < 6; fi9++ {
                    v3 = (v3 - v3)
                    if (true && true) {
                        if false {
                            delete(v2, "")
                        } else {
                            switch fi9 {
                            case 127:
                                if ((1 >= 538) || (false && ((fi8 - 18) <= fi9))) {
                                    _ = len(v2)
                                } else {
                                    var v10 uint = 1
                                    v1 = 1
                                    v2[("foo bar" + ("\n\t" + "\n\t"))] = (fi9 >> ((v10) & 63))
                                    _ = v10
                                }
                            case 32767:
                                if (!(42 < 42)) {
                                    v1 -= v1
                                }
                            case 128:
                                var v11 string = "foo bar"
                                _ = v11
                            default:
                                v1 = v1
                            }
                        }
                    } else {
                        switch int(v1) {
                        case -2147483648:
                            v1 -= 1
                        case 127:
                            var v12 int16 = -32768
                            _ = v12
                        case -1:
                            v3 += (^(v3))
                        }
                        v13 := v2["\n\t"]
                        fmt.Println(fi8, v1, fi9)
                        _ = v13
                    }
                    v2[("foo bar" + "\n\t")] = fi8
                    _ = fi9
                }
            }
            _ = fi8
        }
    default:
        if true {
            v1 -= v1
            v1 = (v1 % (((v1 >> ((38) & 63))) | 1))
            v3 -= 127
        }
    }
    v3 += -101
    v3 = 0
    for fi14 := 0; fi14 < 14; fi14++ {
        if (255 != genFn0(187, 128)) {
            v15 := []int{(fi14 << ((149) & 63)), fi14}
            v3 = 100
            fmt.Println(v3)
            _ = v15
        }
        v16 := GenS1{F0: 65535}
        fmt.Println(v1)
        _ = fi14
        _ = v16
    }
    var v17 uint64 = 0
    if false {
        v18 := GenS2{F0: 2, F1: ("" + "foo bar"), F2: (255 != 1)}
        for fi19 := 0; fi19 < 10; fi19++ {
            v3 += v3
            _ = fi19
        }
        _ = v18
    } else {
        v20 := []int{127, int(v3)}
        switch 42 {
        case 42:
            var v21 int32 = v1
            _ = v21
        case -1:
            fmt.Println(v17)
        case 32767:
            v17 -= v17
        default:
            var v22 int8 = 105
            _ = v22
        }
        _ = v20
    }
    v23 := []int{2147483647}
    for fi24 := 0; fi24 < 18; fi24++ {
        fmt.Println(v3, fi24)
        _ = fi24
    }
    if (2147483647 > v1) {
        var v25 bool = (!false)
        _ = v25
    }
    var v26 string = ("hello" + "\n\t")
    v17 -= v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v26
    _ = v3
}
