package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 int32; F2 int32 }
func genFn0(p0 uint8, p1 int32) int8 {
    p0 = 127
    fmt.Println(p1, p1)
    p1 = (p1 / (((p1 ^ -1)) | 1))
    _ = p0
    _ = p1
    return 77
}

func main() {
    var v1 uint8 = 61
    v2 := []uint32{4294967295}
    var v3 uint64 = 4294967295
    v4 := make(map[int]int)
    v5 := []string{"", (("\n\t" + "a") + "")}
    var v6 float32 = 1.5
    if false {
        var v7 int16 = 42
        if true {
            switch 32767 {
            case 32768:
                v8 := GenS1{F0: v3}
                _ = v8
            case 127:
                v6 = (0.1 * v6)
            case -47:
                for fi9 := 0; fi9 < 1; fi9++ {
                    v1 = 0
                    var v10 int64 = -9223372036854775808
                    _ = fi9
                    _ = v10
                }
            default:
                v2 = append(v2, 61)
            }
            var v11 uint8 = (v1 + (v1 + (v1 + (^(v1)))))
            switch 127 {
            case 32768:
                for fi12 := 0; fi12 < 9; fi12++ {
                    var v13 uint32 = 73
                    if (3.14159 == v6) {
                        for fi14 := 0; fi14 < 20; fi14++ {
                            fmt.Println(fi14, v13, v3)
                            var v15 float32 = v6
                            v11 += v11
                            v16 := GenS2{F0: true}
                            var v17 uint16 = 1
                            v18 := GenS1{F0: 784}
                            _ = fi14
                            _ = v15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    }
                    v19 := GenS1{F0: v3}
                    _ = fi12
                    _ = v13
                    _ = v19
                }
            case -1:
                v6 -= 1.5
            default:
                fmt.Println(v11, v11, v6)
            }
            _ = v11
        } else {
            v1 += v1
            var v20 float32 = 0.1
            v4[127] = 9223372036854775807
            _ = v20
        }
        _ = v7
    }
    if true {
        var v21 uint = uint(v1)
        switch 127 {
        case -2147483648:
            fmt.Println(v6)
        }
        _ = v21
    } else {
        v3 = (v3 * 0)
        if (true || (!true)) {
            for fi22 := 0; fi22 < 17; fi22++ {
                if false {
                    if false {
                        var v23 int8 = 100
                        fmt.Println(v6, v6)
                        for fi24 := 0; fi24 < 11; fi24++ {
                            if ((true || (!false)) || (1 > 835)) {
                                v25 := v4[(fi24 / (((fi22 % ((fi24) | 1))) | 1))]
                                _ = v25
                            } else {
                                delete(v4, (fi24 ^ fi22))
                                v6 = (((0.0 * float32((3.14159 / float64(((v6 + 1.5) * v6))))) - v6) + v6)
                                fmt.Println(v6)
                            }
                            _ = fi24
                        }
                        _ = v23
                    } else {
                        for fi26 := 0; fi26 < 7; fi26++ {
                            var v27 int8 = 42
                            v4[(fi22 - (fi22 + -141))] = fi22
                            _ = fi26
                            _ = v27
                        }
                        if false {
                            _ = len(v4)
                            var v28 string = "hello"
                            if true {
                                v28 = (v28 + "hello")
                                for fi29 := 0; fi29 < 17; fi29++ {
                                    var v30 uint8 = v1
                                    if (true || true) {
                                        var v31 float32 = 3.14159
                                        var v32 int32 = 0
                                        var v33 uint32 = 437
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    } else {
                                        v1 = v1
                                    }
                                    var v34 string = v28
                                    _ = fi29
                                    _ = v30
                                    _ = v34
                                }
                                fmt.Println(v28)
                            } else {
                                var v35 float32 = (0.0 - 0.1)
                                v3 = 4294967295
                                _ = v35
                            }
                            _ = v28
                        }
                    }
                    var v36 uint16 = 42
                    v36 += 255
                    _ = v36
                } else {
                    v37 := GenS1{F0: (v3 * (v3 / ((v3) | 1)))}
                    for fi38 := 0; fi38 < 7; fi38++ {
                        var v39 uint = 653
                        _ = fi38
                        _ = v39
                    }
                    v40 := GenS3{F0: 32767, F1: 996}
                    _ = v37
                    _ = v40
                }
                var v41 int = -9223372036854775808
                _ = fi22
                _ = v41
            }
            v5 = append(v5, ("foo bar" + "a"))
        } else {
            v6 = 1.5
            fmt.Println(v3, v3)
        }
    }
    var v42 string = "hello"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi43 := 0; fi43 < 15; fi43++ {
        v1 = v1
        _ = fi43
    }
    for fi44 := 0; fi44 < 5; fi44++ {
        v42 = v42
        var v45 int64 = -9223372036854775808
        for fi46 := 0; fi46 < 10; fi46++ {
            _ = len(v5)
            v47 := GenS2{F0: true}
            _ = fi46
            _ = v47
        }
        _ = fi44
        _ = v45
    }
    var v48 int8 = 1
    for fi49 := 0; fi49 < 2; fi49++ {
        switch (fi49 ^ (fi49 % ((int(v1)) | 1))) {
        case 2147483647:
            v1 += 183
        default:
            if (int(v3) <= fi49) {
                if true {
                    if true {
                        v48 = int8(v3)
                    }
                    switch 127 {
                    case -9223372036854775808:
                        v3 = (v3 ^ 2)
                    }
                    for fi50 := 0; fi50 < 19; fi50++ {
                        var v51 uint = 290
                        var v52 uint8 = v1
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                } else {
                    fmt.Println(v48, v42, v1)
                }
                v2 = append(v2, 65535)
            } else {
                for fi53 := 0; fi53 < 6; fi53++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if ((!(-2.25 != 1.5)) || false) {
                        fmt.Println(fi49, fi53, fi49)
                        switch fi49 {
                        case 4294967295:
                            var v54 uint32 = 65535
                            _ = v54
                        case 42:
                            v1 += v1
                        case 821:
                            if (v42 >= v42) {
                                if true {
                                    var v55 float64 = 1e10
                                    _ = v55
                                }
                                v56 := GenS4{F0: v3}
                                if true {
                                    fmt.Println(v3)
                                    v57 := GenS5{F0: -643}
                                    v6 += v6
                                    _ = v57
                                } else {
                                    switch (v42 + v42) {
                                    case "a":
                                        fmt.Println(fi53)
                                    }
                                }
                                _ = v56
                            } else {
                                fmt.Println(fi49, v3, fi53)
                                if (!true) {
                                    switch fi53 {
                                    case 1:
                                        var v58 float32 = float32(1e10)
                                        _ = v58
                                    case -763:
                                        v48 -= genFn0(v1, -1)
                                        var v59 int16 = 32767
                                        var v60 uint8 = (v1 << ((802) & 63))
                                        _ = v59
                                        _ = v60
                                    }
                                    var v61 int16 = 1
                                    v42 = ""
                                    _ = v61
                                } else {
                                    var v62 bool = false
                                    _ = v62
                                }
                                v63 := GenS4{F0: v3}
                                _ = v63
                            }
                        }
                        fmt.Println(v3, v42)
                    }
                    _ = fi53
                }
            }
        }
        var v64 string = v42
        _ = fi49
        _ = v64
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v6)
    case "":
        v65 := GenS6{F0: 42, F1: -739, F2: 42}
        _ = v65
    case "\n\t":
        v66 := GenS6{F0: 255, F1: 911, F2: int32(v48)}
        _ = v66
    default:
        if (2 != v3) {
            fmt.Println(v1, v3)
            v42 = "hello"
            var v67 uint = 42
            _ = v67
        }
    }
    v68 := GenS2{F0: (!(true || ((618 >= 0) || false)))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v48
    _ = v5
    _ = v6
    _ = v68
}
