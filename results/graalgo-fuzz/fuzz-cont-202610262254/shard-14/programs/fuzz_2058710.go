package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
func genFn0() int {
    v6 := []int{0, -2147483648}
    v7 := GenS1{F0: (!(!(18446744073709551615 <= 4294967295)))}
    var v8 uint32 = 557
    v8 = v8
    var v9 int = -9223372036854775808
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1() uint16 {
    v10 := GenS1{F0: false}
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint8 = 0
        if (255 < (v12 / (((v12 / (((v12 * (v12 + (v12 >> ((255) & 63))))) | 1))) | 1))) {
            switch "hello" {
            case "":
                for fi13 := 0; fi13 < 1; fi13++ {
                    if true {
                        v14 := []int{fi13}
                        v12 -= v12
                        v12 += (v12 >> ((65535) & 63))
                        _ = v14
                    } else {
                        v12 += 2
                        v15 := GenS2{F0: -1, F1: "a"}
                        v16 := GenS2{F0: fi11, F1: "a"}
                        _ = v15
                        _ = v16
                    }
                    if false {
                        v17 := []int{116, 577, (fi13 % (((fi13 + genFn0())) | 1))}
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        v12 -= 2
                        _ = v17
                    }
                    _ = fi13
                }
            default:
                for fi18 := 0; fi18 < 6; fi18++ {
                    var v19 bool = (0 > 42)
                    _ = fi18
                    _ = v19
                }
            }
            v12 = v12
        } else {
            var v20 int64 = 9223372036854775807
            for fi21 := 0; fi21 < 16; fi21++ {
                var v22 int64 = 627
                _ = fi21
                _ = v22
            }
            v12 -= (^(v12))
            _ = v20
        }
        _ = fi11
        _ = v12
    }
    var v23 int16 = 0
    _ = v10
    _ = v23
    return 664
}

func main() {
    var v1 uint16 = 32767
    var v2 float64 = (1e10 * 3.14159)
    v3 := make(map[int]uint32)
    var v4 int32 = 783
    var v5 uint16 = (v1 * (v1 - 32767))
    for fi24 := 0; fi24 < 7; fi24++ {
        for fi25 := 0; fi25 < 9; fi25++ {
            v26 := GenS2{F0: (fi24 - -600), F1: "hello"}
            v1 += (v5 * v1)
            if ((964 > 42) && true) {
                if (!(!true)) {
                    fmt.Println(v2, v1, v1)
                }
                if false {
                    var v27 int = (fi25 | -798)
                    var v28 uint8 = 42
                    _ = v27
                    _ = v28
                } else {
                    var v29 uint64 = 169
                    v30 := GenS2{F0: 4294967295, F1: "a"}
                    _ = v29
                    _ = v30
                }
            }
            _ = fi25
            _ = v26
        }
        var v31 int32 = (-(v4))
        _ = fi24
        _ = v31
    }
    v1 -= (v1 & v1)
    v32 := GenS3{F0: (v5 & v5)}
    v33 := v3[724]
    for fi34 := 0; fi34 < 18; fi34++ {
        for fi35 := 0; fi35 < 20; fi35++ {
            if (1 < 18446744073709551615) {
                v33 = v33
                v36 := []int{}
                _ = v36
            }
            var v37 uint32 = (v33 / (((v33 - v33)) | 1))
            if false {
                fmt.Println(v33, v2, v4)
            } else {
                if (1 == -1) {
                    v38 := []int{}
                    if true {
                        v39 := GenS4{F0: int(v5)}
                        if ((v2 * v2) != 0.1) {
                            v33 += v33
                        } else {
                            var v40 int8 = -4
                            _ = v40
                        }
                        _ = v39
                    } else {
                        v41 := GenS4{F0: fi35}
                        var v42 uint16 = (v1 + v5)
                        fmt.Println(v1, fi34)
                        _ = v41
                        _ = v42
                    }
                    _ = len(v3)
                    _ = v38
                }
            }
            fmt.Println(v33, v4)
            _ = fi35
            _ = v37
        }
        _ = fi34
    }
    var v43 string = "hello"
    v1 += (^(v1))
    v44 := GenS5{F0: -1}
    v45 := GenS3{F0: 1}
    v43 = ("a" + v43)
    var v46 int8 = 42
    var v47 string = v43
    v33 = 4294967295
    var v48 string = ""
    switch 486 {
    case -903:
        _ = len(v3)
    default:
        var v49 int32 = v4
        _ = v49
    }
    v50 := []int{}
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v50
}
