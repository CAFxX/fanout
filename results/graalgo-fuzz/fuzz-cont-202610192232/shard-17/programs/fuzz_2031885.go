package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 uint8 }

func main() {
    var v1 uint16 = 0
    var v2 uint32 = 1
    var v3 string = "hello"
    var v4 int16 = -355
    var v5 int64 = -1
    fmt.Println(v4, v2)
    switch 4294967295 {
    case 355:
        var v6 uint16 = (v1 & v1)
        _ = v6
    case 2147483647:
        v7 := GenS1{F0: 128, F1: (v1 & v1)}
        _ = v7
    case 1:
        v2 += 1
    }
    fmt.Println(v3, v3, v3)
    fmt.Println(v2)
    var v8 int16 = v4
    for fi9 := 0; fi9 < 9; fi9++ {
        v1 = v1
        v10 := []int{(-(fi9))}
        if (true && true) {
            v11 := GenS2{F0: 2}
            switch (fi9 - 666) {
            case 0:
                if (false || false) {
                    v4 = (v4 >> ((0) & 63))
                } else {
                    v12 := GenS3{F0: ((v3 + (v3 + "hello")) + "\n\t"), F1: 255}
                    if len(v10) > 0 {
                        v10[0] = 351
                    }
                    var v13 int8 = 127
                    _ = v12
                    _ = v13
                }
            case 9223372036854775807:
                v10 = append(v10, int(v2))
            }
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    v14 := GenS4{F0: 2, F1: (!true)}
    var v15 int8 = 42
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            if false {
                fmt.Println(v4, v4, v3)
                switch v3 {
                case "hello":
                    v15 -= v15
                case "":
                    var v18 int32 = 1
                    _ = v18
                default:
                    if true {
                        v19 := []int{(^(fi17)), fi16, fi17, (fi17 * (fi17 & 42))}
                        _ = v19
                    }
                }
                var v20 float32 = 3.14159
                _ = v20
            } else {
                var v21 int32 = -925
                if (!false) {
                    if false {
                        if false {
                            v22 := GenS5{F0: v15}
                            v21 = v21
                            var v23 uint16 = 65535
                            _ = v22
                            _ = v23
                        }
                        if true {
                            if (true || (-2147483648 < 1)) {
                                var v24 float32 = 1.5
                                v25 := []int{9223372036854775807, 734, (-(fi17))}
                                var v26 int16 = (v8 ^ -699)
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                v27 := GenS5{F0: (-(v15))}
                                v21 = v21
                                v28 := []int{int(v5), fi16}
                                _ = v27
                                _ = v28
                            }
                            var v29 uint16 = v1
                            _ = v29
                        }
                    }
                    switch v3 {
                    case "hello":
                        var v30 uint32 = v2
                        v8 -= v4
                        switch ((v3 + v3) + ("" + "\n\t")) {
                        case "a":
                            var v31 int8 = (v15 | (v15 ^ v15))
                            v31 = -6
                            switch int(v30) {
                            case -260:
                                var v32 uint = 18446744073709551615
                                _ = v32
                            case 32768:
                                switch (v3 + "\n\t") {
                                case "hello":
                                    var v33 uint16 = 340
                                    v34 := []int{(fi17 >> ((uint(v1)) & 63)), fi17, (fi17 * (-(fi16))), fi16}
                                    var v35 int64 = (v5 + v5)
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                            case 2147483647:
                                var v36 uint = 143
                                _ = v36
                            default:
                                v4 -= -1
                            }
                            if true {
                                v37 := GenS2{F0: 8}
                                var v38 int8 = v31
                                _ = v37
                                _ = v38
                            } else {
                                v39 := []int{fi17, fi17}
                                _ = v39
                            }
                            fmt.Println(v3, v1)
                            _ = v31
                        case "hello":
                            v40 := []int{(fi16 & -17), (fi16 + 126), fi16, -2147483648}
                            _ = v40
                        default:
                            v30 += 4294967295
                        }
                        _ = v30
                    case "":
                        if (!(((1e10 * -2.25) / float32(1e10)) >= -2.25)) {
                            if (!true) {
                                var v41 uint = 18446744073709551615
                                fmt.Println(v8, v4, v21)
                                v42 := []int{}
                                var v43 uint16 = 0
                                v44 := GenS2{F0: (v41 | (^(v41)))}
                                _ = v41
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            if true {
                                v45 := []int{331}
                                v46 := GenS1{F0: 127, F1: (v1 + 42)}
                                var v47 uint32 = v2
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                            fmt.Println(v8)
                        } else {
                            if true {
                                if false {
                                    if false {
                                        var v48 int32 = v21
                                        v5 = v5
                                        _ = v48
                                    }
                                    switch (fi16 << ((42) & 63)) {
                                    case 0:
                                        v1 -= v1
                                        v4 = int16(v21)
                                        var v49 uint32 = v2
                                        if (3.14159 >= (3.14159 + 1e10)) {
                                            v50 := []int{}
                                            _ = v50
                                        }
                                        _ = v49
                                    case -362:
                                        v4 -= (v8 - v8)
                                    }
                                }
                            } else {
                                v21 += 0
                                v2 = (v2 ^ v2)
                            }
                            v1 += v1
                        }
                    default:
                        v51 := GenS5{F0: -45}
                        _ = v51
                    }
                    v8 = v8
                    v52 := []int{(fi17 * int(v1)), (fi16 ^ int(v1)), (fi17 & fi16)}
                    if (v15 <= (-(v15))) {
                        var v53 uint = uint(v1)
                        v8 -= int16(v21)
                        _ = v53
                    } else {
                        fmt.Println(v15, v2, v5)
                    }
                    _ = v52
                } else {
                    var v54 int64 = v5
                    v55 := []int{(fi16 | fi17), 0, 458, fi17}
                    _ = v54
                    _ = v55
                }
                v2 = (v2 ^ (v2 - (v2 << ((259) & 63))))
                _ = v21
            }
            var v56 string = v3
            _ = fi17
            _ = v56
        }
        v57 := []int{127, -459, (-(fi16)), (^(fi16))}
        var v58 bool = true
        _ = fi16
        _ = v57
        _ = v58
    }
    switch "\n\t" {
    case "hello":
        for fi59 := 0; fi59 < 6; fi59++ {
            v60 := GenS2{F0: 642}
            _ = fi59
            _ = v60
        }
    }
    var v61 uint = 0
    for fi62 := 0; fi62 < 5; fi62++ {
        v63 := GenS6{F0: v15}
        _ = fi62
        _ = v63
    }
    v64 := GenS7{F0: 128, F1: 127}
    var v65 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v65:", v65)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v64
    _ = v65
    _ = v8
}
