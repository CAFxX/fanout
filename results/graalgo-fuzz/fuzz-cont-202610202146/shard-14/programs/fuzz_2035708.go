package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int16; F2 uint16 }
func genFn0(p0 uint64) int32 {
    switch (("\n\t" + "hello") + ("foo bar" + "")) {
    case "":
        var v7 int = -3
        _ = v7
    default:
        var v8 uint32 = 90
        _ = v8
    }
    if false {
        var v9 int = 292
        _ = v9
    }
    if (!(((!(!(true || false))) || true) || (1 < 379))) {
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := []int{(fi10 - fi10), -2147483648, (fi10 % ((fi10) | 1)), fi10}
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            p0 = p0
            _ = fi10
            _ = v11
        }
        if false {
            p0 -= 513
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 string = "a"
                for fi14 := 0; fi14 < 6; fi14++ {
                    fmt.Println(p0, p0, p0)
                    if (621 < 42) {
                        if (false && (!true)) {
                            var v15 uint8 = 246
                            var v16 float32 = (1.5 + -2.25)
                            var v17 uint = 629
                            v18 := []int{(fi14 * fi12), fi12, (fi12 + (fi14 - fi12)), -125}
                            _ = v15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        v19 := GenS1{F0: (((v13 + ((v13 + ("a" + v13)) + v13)) + "a") + v13)}
                        _ = v19
                    }
                    _ = fi14
                }
                if ((!(100 > -14)) && (!(((0 < 2) && false) || true))) {
                    var v20 string = v13
                    _ = v20
                } else {
                    switch fi12 {
                    case 32768:
                        switch 0 {
                        case 127:
                            v21 := []int{fi12, fi12, (fi12 & (fi12 | (-(fi12)))), fi12}
                            _ = v21
                        default:
                            switch (fi12 | fi12) {
                            case -896:
                                p0 += 117
                            case 1:
                                var v22 uint16 = 255
                                _ = v22
                            case -1:
                                v13 = (v13 + ((("foo bar" + v13) + "hello") + "a"))
                            default:
                                var v23 uint32 = 1
                                _ = v23
                            }
                        }
                    case 32767:
                        fmt.Println(v13, fi12, fi12)
                    case 127:
                        p0 = p0
                    default:
                        p0 += p0
                    }
                }
                _ = fi12
                _ = v13
            }
        }
    } else {
        fmt.Println(p0)
        p0 += 18446744073709551615
        if true {
            var v24 string = "a"
            _ = v24
        }
    }
    if false {
        v25 := GenS2{F0: 2, F1: 1000, F2: uint16(p0)}
        var v26 int = 128
        _ = v25
        _ = v26
    }
    fmt.Println(p0)
    _ = p0
    return 1
}
func genFn1(p0 uint64, p1 int, p2 uint32) uint {
    fmt.Println(p0, p1)
    fmt.Println(p1)
    var v27 int32 = genFn0(uint64(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 4294967295
}

func main() {
    var v1 uint32 = 478
    v2 := make(map[uint32]string)
    v3 := []uint32{(v1 ^ (v1 % ((1) | 1)))}
    v4 := []uint32{}
    v5 := make(map[string]bool)
    v6 := make(map[int]uint32)
    var v28 int32 = 494
    v29 := v2[(v1 * (v1 | 249))]
    fmt.Println(v1)
    for fi30 := 0; fi30 < 20; fi30++ {
        switch (-(fi30)) {
        case 4294967295:
            v1 = 4294967295
        case 127:
            delete(v2, v1)
        case 238:
            if ((true && false) && true) {
                var v31 string = v29
                delete(v2, (^(v1)))
                _ = v31
            } else {
                v32 := GenS2{F0: 265, F1: 0, F2: uint16(v1)}
                _ = v32
            }
        }
        if false {
            if ((1e10 != float32((1.5 + 0.1))) || true) {
                var v33 uint = 868
                v29 = "a"
                _ = v33
            } else {
                var v34 float32 = -2.25
                for fi35 := 0; fi35 < 14; fi35++ {
                    v28 = genFn0(872)
                    _ = fi35
                }
                _ = len(v6)
                _ = v34
            }
            fmt.Println(v29, v28)
            v2[v1] = "foo bar"
        }
        v29 = "foo bar"
        _ = fi30
    }
    var v36 uint8 = 0
    v3 = append(v3, v1)
    if false {
        for fi37 := 0; fi37 < 14; fi37++ {
            var v38 bool = false
            v1 = 2147483648
            v29 = v29
            _ = fi37
            _ = v38
        }
        v36 = v36
    }
    var v39 uint = 4294967295
    var v40 string = (v29 + (((("a" + "hello") + "hello") + "") + v29))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
