package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 int64 }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := []int{(fi8 & 2147483647), (fi8 >> ((445) & 63)), (fi7 >> ((1) & 63))}
            var v10 uint64 = 18446744073709551615
            v11 := GenS1{F0: 65535, F1: 2}
            v10 = (v10 << ((65535) & 63))
            v10 = (^(v10))
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        v12 := GenS1{F0: 0, F1: 1}
        var v13 int32 = -524
        _ = fi7
        _ = v12
        _ = v13
    }
    var v14 int32 = 1
    fmt.Println(v14)
    v15 := []int{740, -1, -390}
    _ = v14
    _ = v15
    return 4294967295
}
func genFn1(p0 uint8) uint {
    fmt.Println(p0)
    v16 := []int{457}
    p0 += (p0 % (((p0 + p0)) | 1))
    _ = len(v16)
    v17 := GenS1{F0: 767, F1: 4294967295}
    _ = p0
    _ = v16
    _ = v17
    return 539
}

func main() {
    v1 := []uint32{65535}
    var v2 uint32 = 2147483647
    v3 := []uint32{1, v2, (v2 * (v2 | 42))}
    v4 := make(map[uint32]string)
    v5 := []uint32{1, v2, (v2 | (^(v2))), (v2 / ((v2) | 1))}
    var v6 float32 = 1.5
    for fi18 := 0; fi18 < 14; fi18++ {
        v6 -= (float32(1e10) * float32(((1.5 - 0.0) - 3.14159)))
        var v19 string = ""
        v20 := GenS2{F0: (v2 * 0), F1: v19}
        _ = fi18
        _ = v19
        _ = v20
    }
    v2 -= 0
    if true {
        v21 := GenS3{F0: true, F1: 65535}
        v2 = 2147483648
        _ = v21
    }
    var v22 uint8 = 1
    if true {
        v6 -= ((v6 - float32(0.0)) - (1.5 * v6))
    }
    var v23 string = "a"
    v22 += v22
    v24 := GenS1{F0: 955, F1: 2}
    for fi25 := 0; fi25 < 6; fi25++ {
        switch (fi25 ^ fi25) {
        case 371:
            if true {
                var v26 float32 = v6
                var v27 float32 = 1e10
                _ = v26
                _ = v27
            } else {
                if (fi25 > (fi25 | (fi25 ^ 200))) {
                    for fi28 := 0; fi28 < 5; fi28++ {
                        var v29 uint32 = (v2 + 471)
                        switch fi28 {
                        case 9223372036854775807:
                            var v30 float32 = 3.14159
                            _ = v30
                        case -2147483648:
                            delete(v4, (v29 & v2))
                        case -1:
                            if false {
                                switch 32767 {
                                case -2147483648:
                                    var v31 uint64 = 4294967295
                                    _ = v31
                                case 128:
                                    if false {
                                        v6 -= 1e10
                                        switch "foo bar" {
                                        case "a":
                                            var v32 float32 = 1.5
                                            _ = v32
                                        case "hello":
                                            switch (^(fi25)) {
                                            case 906:
                                                var v33 uint8 = (v22 + (v22 + v22))
                                                _ = v33
                                            case 4294967295:
                                                var v34 float64 = (3.14159 + (3.14159 - float64((v6 + (1.5 / 1.0)))))
                                                _ = v34
                                            case -9223372036854775808:
                                                fmt.Println(v29, v22)
                                            default:
                                                v35 := GenS4{F0: 246}
                                                _ = v35
                                            }
                                        case "foo bar":
                                            v4[genFn0()] = ""
                                        }
                                        var v36 uint16 = 400
                                        _ = v36
                                    } else {
                                        switch fi28 {
                                        case 4294967295:
                                            v23 = v23
                                        case -653:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case -9223372036854775808:
                                            var v37 bool = false
                                            _ = v37
                                        }
                                        if (42 > 318) {
                                            var v38 uint64 = 875
                                            _ = v38
                                        } else {
                                            var v39 uint8 = (^(v22))
                                            if true {
                                                var v40 bool = (!((fi25 | (fi28 >> ((743) & 63))) >= (fi25 % (((^(fi25))) | 1))))
                                                _ = v40
                                            }
                                            var v41 bool = true
                                            _ = v39
                                            _ = v41
                                        }
                                        var v42 int = fi25
                                        v43 := GenS2{F0: (v2 | (^(v2))), F1: "a"}
                                        _ = v42
                                        _ = v43
                                    }
                                }
                                v23 = v23
                            } else {
                                _ = len(v3)
                                v44 := v4[924]
                                _ = v44
                            }
                        }
                        var v45 uint8 = v22
                        _ = fi28
                        _ = v29
                        _ = v45
                    }
                } else {
                    for fi46 := 0; fi46 < 14; fi46++ {
                        v47 := GenS5{F0: (fi46 > (fi25 * (^(fi25)))), F1: 42}
                        _ = fi46
                        _ = v47
                    }
                    v4[genFn0()] = "foo bar"
                }
            }
        }
        _ = fi25
    }
    switch ("\n\t" + "") {
    case "hello":
        var v48 int = -554
        _ = v48
    case "foo bar":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "\n\t":
        fmt.Println(v6, v6)
    default:
        if true {
            for fi49 := 0; fi49 < 1; fi49++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                for fi50 := 0; fi50 < 12; fi50++ {
                    v51 := GenS4{F0: 0}
                    _ = fi50
                    _ = v51
                }
                switch 127 {
                case 32768:
                    var v52 uint8 = (v22 / ((v22) | 1))
                    _ = v52
                case 897:
                    v53 := GenS3{F0: true, F1: 326}
                    _ = v53
                case 2147483647:
                    _ = len(v3)
                }
                _ = fi49
            }
            var v54 uint64 = 18446744073709551615
            var v55 float64 = 1e10
            _ = v54
            _ = v55
        } else {
            v23 = "a"
        }
    }
    var v56 bool = false
    v6 = v6
    v56 = v56
    if (true && (670 < -2147483648)) {
        var v57 uint16 = 754
        if v56 {
            delete(v4, (v2 * (v2 | uint32(v22))))
        }
        if v56 {
            var v58 uint8 = v22
            _ = v58
        }
        _ = v57
    } else {
        for fi59 := 0; fi59 < 6; fi59++ {
            v4[v2] = v23
            _ = fi59
        }
        if true {
            v56 = (!(!(1.5 < ((0.1 - 3.14159) + 0.0))))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch "hello" {
            case "a":
                fmt.Println(v23, v6)
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
