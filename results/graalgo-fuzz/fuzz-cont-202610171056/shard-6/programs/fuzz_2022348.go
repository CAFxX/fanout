package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 bool; F2 uint64 }
func genFn0(p0 uint64, p1 uint64, p2 int) uint32 {
    switch "a" {
    case "a":
        v6 := []int{}
        _ = v6
    case "hello":
        p2 += p2
    default:
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 uint8 = 2
            if (850 == 2147483648) {
                if (!true) {
                    v9 := []int{p2, p2, (fi7 % (((^(fi7))) | 1)), (p2 & (p2 >> ((989) & 63)))}
                    if (true && false) {
                        p0 = p0
                        for fi10 := 0; fi10 < 17; fi10++ {
                            v11 := GenS1{F0: p1}
                            var v12 uint32 = 654
                            var v13 uint64 = 780
                            _ = fi10
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                    }
                    var v14 int = p2
                    _ = v14
                    _ = v9
                } else {
                    var v15 int64 = 158
                    _ = v15
                }
                v16 := GenS2{F0: p0}
                _ = v16
            } else {
                v17 := []int{}
                v18 := GenS1{F0: (^(p0))}
                _ = v17
                _ = v18
            }
            fmt.Println(p2, p0)
            _ = fi7
            _ = v8
        }
    }
    var v19 uint32 = 2147483648
    for fi20 := 0; fi20 < 6; fi20++ {
        if false {
            var v21 int16 = 42
            p0 = 4294967295
            _ = v21
        } else {
            var v22 int = p2
            p1 = 1
            if false {
                if false {
                    var v23 uint8 = 128
                    var v24 uint32 = 2
                    _ = v23
                    _ = v24
                }
            }
            _ = v22
        }
        p2 = 32767
        _ = fi20
    }
    v25 := []int{p2}
    fmt.Println(p2, p0, v19)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v25
    return v19
}

func main() {
    var v1 int32 = -1
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]uint32)
    v4 := []uint32{}
    var v5 uint16 = 65535
    if false {
        for fi26 := 0; fi26 < 15; fi26++ {
            if true {
                _ = len(v3)
                for fi27 := 0; fi27 < 15; fi27++ {
                    if len(v4) > 0 {
                        v4[0] = 2147483648
                    }
                    _ = len(v3)
                    _ = fi27
                }
            }
            var v28 uint64 = 0
            _ = fi26
            _ = v28
        }
        var v29 string = "hello"
        _ = v29
    } else {
        v30 := GenS3{F0: 127, F1: (9223372036854775807 != 127), F2: 0}
        for fi31 := 0; fi31 < 12; fi31++ {
            var v32 uint8 = 1
            v1 -= 42
            var v33 uint8 = (v32 * 2)
            _ = fi31
            _ = v32
            _ = v33
        }
        v5 = v5
        _ = v30
    }
    if true {
        var v34 uint8 = 255
        var v35 float64 = -2.25
        var v36 uint64 = 552
        _ = v34
        _ = v35
        _ = v36
    }
    v2 = (v2 << ((uint(v5)) & 63))
    switch 2147483647 {
    case 32768:
        var v37 uint64 = 2
        _ = v37
    default:
        var v38 int8 = -99
        _ = v38
    }
    v4 = append(v4, 923)
    v5 -= v5
    var v39 uint = 65535
    var v40 uint8 = 1
    for fi41 := 0; fi41 < 11; fi41++ {
        v5 -= v5
        _ = fi41
    }
    v39 = (v39 / ((v39) | 1))
    var v42 string = ""
    _ = len(v4)
    if (!(!true)) {
        var v43 bool = false
        _ = v43
    }
    switch (v42 + (v42 + "foo bar")) {
    case "":
        switch (v42 + v42) {
        case "foo bar":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "a":
            var v44 uint64 = 4294967295
            _ = v44
        case "hello":
            v3[(((("foo bar" + v42) + "hello") + v42) + ("a" + "a"))] = 296
        }
    default:
        delete(v3, v42)
    }
    switch -9223372036854775808 {
    case -1:
        var v45 uint8 = 0
        _ = v45
    case 1:
        if false {
            v40 = 21
        }
    }
    v4 = append(v4, genFn0(42, 958, 4294967295))
    var v46 uint = uint(v40)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v46
    _ = v5
}
