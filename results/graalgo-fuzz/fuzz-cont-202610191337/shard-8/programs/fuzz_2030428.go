package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint8; F1 int64; F2 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 uint; F2 int }
type GenS8 struct { F0 uint64; F1 int; F2 bool }
func genFn0(p0 int32, p1 uint32, p2 int32) int8 {
    p0 = -2147483648
    v6 := []int{}
    v7 := GenS1{F0: 760, F1: 9223372036854775807}
    if (true || (false || true)) {
        v6 = append(v6, 42)
    } else {
        if true {
            v8 := GenS1{F0: 1, F1: 9223372036854775807}
            if (42 >= -1) {
                if (123 >= -1) {
                    p2 += p0
                }
                if true {
                    var v9 string = ("" + ("hello" + ("foo bar" + "")))
                    if (255 == 1) {
                        if len(v6) > 0 {
                            v6[0] = 1
                        }
                    } else {
                        var v10 float32 = (-2.25 * float32((0.0 - float64((1.5 - 1.5)))))
                        var v11 int = -9223372036854775808
                        fmt.Println(v9)
                        _ = v10
                        _ = v11
                    }
                    _ = v9
                }
            } else {
                var v12 uint16 = 330
                _ = v12
            }
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint = 105
    var v2 bool = (((0 >= 1) || false) || (1 != 1))
    v3 := make(map[uint32]bool)
    var v4 int = -1
    var v5 bool = (false || ((-(v4)) >= 2147483647))
    if (!v5) {
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := []int{(-(fi13))}
            v4 -= 898
            _ = fi13
            _ = v14
        }
        if (65535 >= 42) {
            v5 = v2
            var v15 uint8 = 128
            v16 := GenS2{F0: (!v2), F1: 255}
            _ = v15
            _ = v16
        } else {
            switch v4 {
            case 971:
                if ((671 == 1) || true) {
                    var v17 string = ("\n\t" + ((("hello" + "") + ("hello" + "\n\t")) + ""))
                    v18 := []int{v4, v4, v4}
                    v1 += (v1 << ((42) & 63))
                    _ = v17
                    _ = v18
                }
            case 466:
                v2 = (v1 > (^(v1)))
            case 739:
                var v19 int = -1
                _ = v19
            }
            var v20 bool = (!(v2 && true))
            _ = v20
        }
    } else {
        if true {
            v3[42] = v2
            if ((0.1 * (1e10 - 1e10)) <= ((-2.25 / -2.25) / 3.14159)) {
                v21 := GenS2{F0: v5, F1: 568}
                _ = v21
            }
            var v22 int16 = -32768
            _ = v22
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            if (v2 || (!false)) {
                v5 = true
                v1 = 116
                if v5 {
                    v1 = (v1 / ((v1) | 1))
                    var v24 float64 = 0.0
                    _ = v24
                }
            } else {
                v2 = false
                fmt.Println(v5)
                for fi25 := 0; fi25 < 3; fi25++ {
                    if ((0 >= (fi25 + 9223372036854775807)) && v2) {
                        v1 -= 255
                        v4 -= fi23
                    } else {
                        fmt.Println(v1)
                    }
                    if (884 > 2) {
                        v26 := GenS3{F0: -2147483648, F1: 1}
                        v27 := GenS1{F0: 2, F1: fi23}
                        var v28 string = (("\n\t" + "\n\t") + "")
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    _ = fi25
                }
            }
            for fi29 := 0; fi29 < 11; fi29++ {
                var v30 uint8 = 255
                if true {
                    if true {
                        var v31 uint8 = (v30 / ((v30) | 1))
                        var v32 int32 = -2147483648
                        _ = v31
                        _ = v32
                    } else {
                        var v33 int64 = 0
                        v34 := []int{}
                        v35 := GenS2{F0: (v2 && (v2 && (987 >= 582))), F1: 42}
                        var v36 int32 = 452
                        _ = v33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    if v5 {
                        var v37 float64 = (((float64(1e10) * (3.14159 - (0.1 / 1.0))) + (float64((0.1 + (0.0 * 0.0))) - 1e10)) + 0.0)
                        v38 := GenS3{F0: 1, F1: -1}
                        fmt.Println(fi23)
                        _ = v37
                        _ = v38
                    }
                    v39 := GenS4{F0: v30, F1: 1, F2: genFn0(238, 1, 0)}
                    _ = v39
                }
                var v40 int16 = 0
                if ((v30 >> ((882) & 63)) < (^(v30))) {
                    v41 := GenS3{F0: 642, F1: v40}
                    _ = v41
                } else {
                    v42 := GenS4{F0: uint8(v1), F1: 9223372036854775807, F2: -26}
                    _ = v42
                }
                var v43 uint32 = 65535
                delete(v3, v43)
                _ = fi29
                _ = v30
                _ = v40
                _ = v43
            }
            v44 := GenS5{F0: ""}
            _ = fi23
            _ = v44
        }
    }
    v2 = (!v5)
    fmt.Println(v5, v4, v1)
    v45 := v3[520]
    v46 := GenS6{F0: -962}
    v4 = v4
    v3[0] = (!v45)
    var v47 uint = 380
    v48 := v3[335]
    var v49 int32 = 1
    v48 = (1 <= 1)
    v50 := GenS7{F0: (v1 >> ((v1) & 63)), F1: (v47 << ((0) & 63)), F2: v4}
    v5 = (!(2 == 42))
    for fi51 := 0; fi51 < 16; fi51++ {
        v45 = v48
        var v52 uint64 = 4294967295
        v52 -= v52
        _ = fi51
        _ = v52
    }
    v45 = v2
    for fi53 := 0; fi53 < 11; fi53++ {
        delete(v3, 4294967295)
        if v45 {
            v54 := GenS8{F0: 416, F1: (v4 / ((4294967295) | 1)), F2: false}
            _ = v54
        }
        _ = fi53
    }
    v1 -= v47
    if v45 {
        v45 = false
        v55 := []int{(v4 << ((42) & 63)), v4, (v4 & -258)}
        var v56 uint64 = 261
        _ = v55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
