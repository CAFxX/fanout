package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int16, p1 uint8) uint64 {
    p1 = p1
    p1 = (p1 << ((255) & 63))
    p1 = p1
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := []int{800, 4294967295, fi6}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        p1 -= p1
        _ = fi8
    }
    _ = p0
    _ = p1
    return 814
}
func genFn1(p0 uint, p1 uint16, p2 uint) uint64 {
    if true {
        v9 := GenS1{F0: -108}
        _ = v9
    } else {
        switch 1 {
        case 128:
            if false {
                v10 := GenS1{F0: -1}
                var v11 int = 1
                v12 := []int{(-(v11))}
                _ = v10
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 12; fi13++ {
                    p2 = (p0 - (p2 >> ((p0) & 63)))
                    p1 -= (p1 / (((p1 * 0)) | 1))
                    for fi14 := 0; fi14 < 11; fi14++ {
                        var v15 int32 = 0
                        v16 := []int{(fi14 << ((p2) & 63)), fi13, int(p1)}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    _ = fi13
                }
                if false {
                    p2 = p0
                    fmt.Println(p1, p1, p1)
                    var v17 uint = (p2 + 4294967295)
                    _ = v17
                } else {
                    var v18 uint32 = 42
                    var v19 bool = true
                    _ = v18
                    _ = v19
                }
            }
        case 88:
            for fi20 := 0; fi20 < 5; fi20++ {
                fmt.Println(p2, p1, p0)
                var v21 uint = (^(p0))
                for fi22 := 0; fi22 < 15; fi22++ {
                    var v23 int = (fi20 << (((^(p0))) & 63))
                    var v24 string = ""
                    var v25 int32 = -1
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = fi20
                _ = v21
            }
        case 4294967295:
            for fi26 := 0; fi26 < 10; fi26++ {
                if false {
                    p0 += (p0 ^ 714)
                    v27 := []int{951, fi26, fi26}
                    _ = v27
                }
                _ = fi26
            }
        default:
            switch 128 {
            case 2147483647:
                if (-109 >= 56) {
                    p0 = (p0 | (^(p0)))
                    fmt.Println(p2, p0)
                    v28 := []int{-679, 772}
                    _ = v28
                }
            case 9223372036854775807:
                v29 := GenS1{F0: 0}
                _ = v29
            default:
                var v30 int8 = 1
                _ = v30
            }
        }
        p2 -= 42
    }
    var v31 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    return 18446744073709551615
}

func main() {
    var v1 int8 = -1
    var v2 int32 = 451
    v3 := make(map[int]uint32)
    var v4 uint8 = 255
    v5 := make(map[string]string)
    v32 := v3[int(v2)]
    switch -391 {
    case -530:
        v1 += (v1 + 1)
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 int32 = (v2 << ((uint(v32)) & 63))
            var v36 uint = 942
            _ = fi34
            _ = v35
            _ = v36
        }
        v4 = v4
        v37 := []int{fi33, int(v4), fi33, (fi33 << ((1) & 63))}
        _ = fi33
        _ = v37
    }
    if false {
        var v38 bool = false
        for fi39 := 0; fi39 < 10; fi39++ {
            for fi40 := 0; fi40 < 20; fi40++ {
                if (v38 || v38) {
                    var v41 int = fi39
                    v1 += int8(v4)
                    v42 := GenS1{F0: v1}
                    _ = v41
                    _ = v42
                }
                var v43 uint32 = (^(v32))
                _ = fi40
                _ = v43
            }
            v44 := []int{(fi39 & (^(fi39))), (fi39 | fi39)}
            for fi45 := 0; fi45 < 1; fi45++ {
                if true {
                    v46 := GenS1{F0: v1}
                    _ = v46
                }
                v47 := GenS1{F0: 14}
                _ = fi45
                _ = v47
            }
            _ = fi39
            _ = v44
        }
        switch "" {
        case "\n\t":
            v1 = v1
        }
        _ = v38
    } else {
        var v48 int = 482
        v32 -= (v32 / (((^(v32))) | 1))
        _ = v48
    }
    v49 := []int{2147483647, 9223372036854775807, 127, 239}
    fmt.Println(v1)
    var v50 uint64 = 42
    if ((1.5 / (3.14159 / 1.0)) > 1e10) {
        v50 += genFn1(1, 446, 18446744073709551615)
        v4 = v4
        var v51 int64 = int64(v2)
        _ = v51
    }
    v2 += -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v49
    _ = v5
    _ = v50
}
