package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int) int64 {
    v6 := GenS1{F0: 255, F1: -128}
    fmt.Println(p0)
    var v7 uint = 666
    v8 := GenS1{F0: 65535, F1: 1}
    p0 += p0
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 bool, p1 uint, p2 uint16) int64 {
    p2 -= 0
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            p0 = p0
            var v11 uint16 = (p2 % (((^(p2))) | 1))
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 float32 = (1.5 + (((0.0 + (-2.25 - ((-2.25 * 3.14159) + -2.25))) * -2.25) - 0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 250
    v2 := []int{0, 32767, 2147483647}
    v3 := make(map[int]uint32)
    var v4 int = 0
    var v5 bool = (!true)
    switch v4 {
    case 857:
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    case 32767:
        v4 = -1
    default:
        switch (v4 >> ((512) & 63)) {
        case 1:
            var v13 uint = 0
            _ = v13
        }
    }
    _ = len(v2)
    v14 := GenS1{F0: 255, F1: -1}
    var v15 float64 = (1.5 * float64((1e10 * 1e10)))
    if true {
        for fi16 := 0; fi16 < 12; fi16++ {
            fmt.Println(v1, v4, v15)
            v4 -= (v4 & (v4 - fi16))
            _ = fi16
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 uint64 = 707
            switch ("a" + "foo bar") {
            case "hello":
                var v19 string = ""
                _ = v19
            }
            if v5 {
                _ = len(v3)
                var v20 uint32 = 0
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        var v21 int = -1
        _ = v21
    }
    var v22 int8 = -1
    v15 = (((0.1 / v15) + (3.14159 + v15)) + 1.5)
    delete(v3, (v4 >> ((509) & 63)))
    if false {
        for fi23 := 0; fi23 < 1; fi23++ {
            v1 = 126
            var v24 float64 = v15
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 uint16 = 2
                switch "a" {
                case "foo bar":
                    v27 := GenS1{F0: v26, F1: -128}
                    _ = v27
                default:
                    v28 := v3[fi25]
                    _ = v28
                }
                var v29 int16 = -32768
                _ = fi25
                _ = v26
                _ = v29
            }
            _ = fi23
            _ = v24
        }
    }
    var v30 int32 = -2147483648
    var v31 uint64 = 1
    v15 -= v15
    if true {
        for fi32 := 0; fi32 < 7; fi32++ {
            delete(v3, 42)
            _ = fi32
        }
    }
    if v5 {
        for fi33 := 0; fi33 < 10; fi33++ {
            v15 -= 0.0
            if v5 {
                switch "a" {
                case "foo bar":
                    if ((v31 >> ((18446744073709551615) & 63)) > (v31 >> ((65535) & 63))) {
                        v30 -= (v30 | (v30 - v30))
                    } else {
                        if v5 {
                            fmt.Println(fi33, v30)
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        for fi34 := 0; fi34 < 6; fi34++ {
                            if true {
                                v1 += (^(v1))
                            } else {
                                var v35 int16 = 32767
                                if ((v30 ^ (v30 >> ((42) & 63))) >= (^(v30))) {
                                    v15 -= -2.25
                                } else {
                                    var v36 int64 = -1
                                    if len(v2) > 0 {
                                        v2[0] = (fi34 >> ((1) & 63))
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v36
                                }
                                _ = v35
                            }
                            v4 = v4
                            v5 = true
                            v37 := GenS1{F0: 2, F1: (v22 ^ (-(v22)))}
                            _ = fi34
                            _ = v37
                        }
                        _ = len(v2)
                    }
                }
                if (42 != 65535) {
                    for fi38 := 0; fi38 < 15; fi38++ {
                        var v39 uint32 = 2
                        _ = fi38
                        _ = v39
                    }
                    v40 := GenS2{F0: (v31 / (((v31 % ((4294967295) | 1))) | 1))}
                    v30 = v30
                    _ = v40
                }
            }
            delete(v3, 32768)
            _ = fi33
        }
        var v41 int = int(v1)
        var v42 float32 = -2.25
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        _ = len(v3)
        if len(v2) > 0 {
            v2[0] = (fi43 << ((65535) & 63))
        }
        _ = fi43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
