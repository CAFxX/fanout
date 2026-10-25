package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int8; F2 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint16, p1 int32, p2 uint32) int16 {
    var v6 float32 = 0.0
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 float32 = v6
        var v9 uint16 = (p0 | p0)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    v3 := []string{"", ("foo bar" + ("a" + "\n\t")), ("hello" + ("" + "a")), "a"}
    var v4 int64 = 0
    v5 := make(map[uint32]int)
    v4 -= -694
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 int = fi10
            if (31 == 0) {
                v12 += fi10
                v13 := GenS1{F0: 1, F1: 18446744073709551615, F2: 100}
                _ = v13
            }
            v14 := v2[(-(fi11))]
            _ = fi11
            _ = v12
            _ = v14
        }
        v15 := GenS2{F0: 65535}
        _ = fi10
        _ = v15
    }
    if true {
        v16 := GenS1{F0: 92, F1: 2, F2: 100}
        _ = v16
    }
    v3 = append(v3, "\n\t")
    v4 += v4
    if false {
        switch 127 {
        case 0:
            var v17 int64 = v4
            _ = v17
        }
        if false {
            v4 = (v4 & 42)
            v4 = (v4 + v4)
            fmt.Println(v4)
        } else {
            if true {
                var v18 uint32 = 65535
                _ = v18
            } else {
                var v19 int32 = -2147483648
                delete(v1, 128)
                v20 := GenS3{F0: v19, F1: 42, F2: v4}
                _ = v19
                _ = v20
            }
        }
        var v21 float64 = 1.5
        _ = v21
    } else {
        var v22 int64 = (v4 >> ((255) & 63))
        v23 := GenS2{F0: 255}
        v24 := v1[127]
        _ = v22
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: 255}
    for fi26 := 0; fi26 < 13; fi26++ {
        for fi27 := 0; fi27 < 16; fi27++ {
            v28 := GenS3{F0: 512, F1: 42, F2: 414}
            switch fi26 {
            case 457:
                if (1 != 255) {
                    if (float64(3.14159) == (3.14159 + 1e10)) {
                        fmt.Println(v4)
                        switch int(v4) {
                        case 32768:
                            switch "" {
                            case "a":
                                if false {
                                    if true {
                                        var v29 bool = (-1 != 0)
                                        _ = v29
                                    } else {
                                        v4 = (v4 / ((v4) | 1))
                                        var v30 uint16 = 0
                                        if len(v3) > 0 {
                                            v3[0] = "a"
                                        }
                                        _ = v30
                                    }
                                    v31 := GenS2{F0: 28}
                                    if false {
                                        var v32 int = fi26
                                        v32 += fi27
                                        _ = len(v3)
                                        _ = v32
                                    }
                                    _ = v31
                                }
                            default:
                                var v33 uint8 = 96
                                _ = v33
                            }
                        case 982:
                            v4 += (v4 % ((v4) | 1))
                        case 2147483647:
                            v34 := GenS1{F0: 2147483648, F1: 767, F2: 123}
                            _ = v34
                        default:
                            var v35 int = fi27
                            _ = v35
                        }
                    }
                } else {
                    v36 := GenS2{F0: 452}
                    v37 := GenS4{F0: (v4 ^ -142)}
                    _ = v36
                    _ = v37
                }
                var v38 string = ""
                _ = v38
            default:
                fmt.Println(fi27, fi27, v4)
            }
            _ = fi27
            _ = v28
        }
        var v39 int8 = 0
        _ = len(v3)
        _ = fi26
        _ = v39
    }
    var v40 float32 = 1.5
    if (4294967295 < 1) {
        v41 := GenS3{F0: 0, F1: 117, F2: v4}
        _ = v41
    } else {
        var v42 bool = false
        if v42 {
            if len(v3) > 0 {
                v3[0] = ((("foo bar" + "a") + ((("foo bar" + "") + ("hello" + "hello")) + ("hello" + ("a" + "foo bar")))) + "foo bar")
            }
            v4 = -1
        } else {
            v4 += (v4 | -705)
        }
        for fi43 := 0; fi43 < 14; fi43++ {
            v3 = append(v3, ("\n\t" + "a"))
            _ = fi43
        }
        _ = v42
    }
    v40 -= (v40 * 1.5)
    switch ("" + (("\n\t" + "a") + ("\n\t" + "\n\t"))) {
    case "a":
        v40 = 0.1
    case "":
        if (-1 > int32(v4)) {
            v44 := GenS5{F0: -1, F1: 255}
            fmt.Println(v4)
            v45 := GenS6{F0: -107}
            _ = v44
            _ = v45
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
