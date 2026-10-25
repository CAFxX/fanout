package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint8 }
func genFn0(p0 uint8) uint64 {
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: true}
        _ = fi5
        _ = v6
    }
    var v7 uint8 = 1
    _ = p0
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 4294967295
    var v3 int8 = 1
    var v4 int8 = v3
    var v8 uint32 = 810
    delete(v1, (^(v8)))
    v4 += v4
    v2 += (v2 | 421)
    var v9 int32 = 0
    v2 = (v2 % ((4294967295) | 1))
    var v10 int64 = 42
    var v11 uint16 = 32767
    v12 := v1[2]
    v13 := GenS1{F0: false}
    if false {
        v2 += (v2 - (v2 >> ((4294967295) & 63)))
    } else {
        var v14 bool = (!(float32(1e10) != -2.25))
        _ = v14
    }
    var v15 int64 = v10
    fmt.Println(v4, v2, v12)
    var v16 int32 = 2147483647
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := GenS1{F0: v12}
        if (345 <= fi17) {
            var v19 string = "foo bar"
            v9 -= 550
            _ = v19
        } else {
            v8 = (v8 ^ (v8 << ((2) & 63)))
            v20 := GenS1{F0: v12}
            _ = v20
        }
        var v21 bool = v12
        _ = fi17
        _ = v18
        _ = v21
    }
    if v12 {
        var v22 uint = 42
        v23 := GenS1{F0: (("" >= "hello") || ((^(v22)) <= 255))}
        _ = v22
        _ = v23
    } else {
        var v24 uint = 42
        switch ("hello" + ("hello" + "\n\t")) {
        case "":
            for fi25 := 0; fi25 < 20; fi25++ {
                delete(v1, (v8 << ((v24) & 63)))
                _ = fi25
            }
        case "a":
            fmt.Println(v11, v16)
        case "\n\t":
            for fi26 := 0; fi26 < 12; fi26++ {
                if (true && v12) {
                    for fi27 := 0; fi27 < 18; fi27++ {
                        v15 = v15
                        _ = fi27
                    }
                }
                var v28 uint32 = 84
                if (fi26 <= int(v16)) {
                    for fi29 := 0; fi29 < 5; fi29++ {
                        if (-1 > (v10 % ((v15) | 1))) {
                            if (v24 < v24) {
                                v30 := []int{(fi29 * fi29), 1, 2147483647, (-(fi29))}
                                _ = v30
                            }
                            delete(v1, (v8 / ((1) | 1)))
                            v31 := GenS1{F0: (v12 && true)}
                            _ = v31
                        } else {
                            if v12 {
                                fmt.Println(v28, v10, v3)
                                v3 = (v3 / (((v3 >> (((v24 ^ v24)) & 63))) | 1))
                            } else {
                                var v32 bool = ((v28 << ((v24) & 63)) <= uint32(v2))
                                _ = v32
                            }
                            switch (fi26 % (((fi29 & (fi26 << ((0) & 63)))) | 1)) {
                            case 128:
                                if (float64(1.5) == 1.5) {
                                    v33 := []int{fi26, fi26}
                                    var v34 int8 = (v3 << (((v24 / ((264) | 1))) & 63))
                                    v16 = (v9 - -2147483648)
                                    _ = v33
                                    _ = v34
                                } else {
                                    var v35 int = (^(fi26))
                                    var v36 int64 = v10
                                    v8 += (^(v8))
                                    _ = v35
                                    _ = v36
                                }
                                v12 = v12
                                v37 := GenS2{F0: 255, F1: v11, F2: 1}
                                _ = v37
                            case -2147483648:
                                v1[568] = (42 <= 42)
                            default:
                                v12 = (false || (((v10 & (v15 + 42)) <= 0) || (!v12)))
                            }
                            v10 = v15
                        }
                        if (!true) {
                            _ = len(v1)
                        }
                        _ = fi29
                    }
                }
                _ = fi26
                _ = v28
            }
        }
        if v12 {
            if v12 {
                var v38 uint16 = (v11 & v11)
                for fi39 := 0; fi39 < 12; fi39++ {
                    var v40 bool = (v12 && (!false))
                    if v40 {
                        var v41 uint32 = 4294967295
                        for fi42 := 0; fi42 < 4; fi42++ {
                            var v43 int64 = 42
                            _ = fi42
                            _ = v43
                        }
                        v8 = (v41 - v41)
                        _ = v41
                    }
                    _ = fi39
                    _ = v40
                }
                _ = v38
            }
            for fi44 := 0; fi44 < 17; fi44++ {
                v45 := GenS2{F0: (^(v11)), F1: v11, F2: 127}
                _ = fi44
                _ = v45
            }
        }
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
