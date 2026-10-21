package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint; F1 int32; F2 uint16 }
type GenS6 struct { F0 int }
func genFn0() uint {
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 uint16 = 0
        switch (fi6 | fi6) {
        case 2147483647:
            v8 := GenS1{F0: 0, F1: 783}
            _ = v8
        case 4294967295:
            v9 := []int{fi6}
            _ = v9
        default:
            if true {
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 int16 = 884
                    _ = fi10
                    _ = v11
                }
            } else {
                v12 := GenS2{F0: -126}
                _ = v12
            }
        }
        var v13 int32 = 3
        _ = fi6
        _ = v13
        _ = v7
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        switch "hello" {
        case "a":
            v15 := []int{(fi14 << ((73) & 63)), -2147483648, 42, fi14}
            _ = v15
        case "\n\t":
            var v16 uint32 = 77
            _ = v16
        default:
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 float32 = float32(3.14159)
                v19 := []int{(-(fi17)), fi14, fi17}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        v20 := GenS3{F0: 1}
        _ = fi14
        _ = v20
    }
    var v21 int64 = 42
    _ = v21
    return 42
}

func main() {
    var v1 float32 = -2.25
    var v2 int32 = 312
    var v3 int32 = v2
    var v4 uint8 = 1
    var v5 int = 127
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 uint = genFn0()
        var v24 int64 = 42
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS4{F0: 42, F1: 42}
    v26 := []int{32767, 1}
    fmt.Println(v3, v5)
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 uint16 = 255
        for fi29 := 0; fi29 < 9; fi29++ {
            if false {
                var v30 int = -9223372036854775808
                _ = v30
            }
            if false {
                fmt.Println(fi27, fi27, v2)
                v2 = 1
                v2 = (v3 % ((v2) | 1))
            } else {
                v5 = fi27
                var v31 bool = ((((3.14159 <= 0.0) && true) && (4 != 100)) || (!(((!false) && false) || (-2.25 < float32(0.0)))))
                var v32 int8 = -90
                _ = v31
                _ = v32
            }
            if false {
                switch ("" + ("a" + ("\n\t" + "hello"))) {
                case "foo bar":
                    v28 -= 0
                case "\n\t":
                    var v33 int = (v5 + -9223372036854775808)
                    _ = v33
                }
                if (4294967295 <= 4294967295) {
                    var v34 float64 = 1e10
                    if false {
                        v4 += (^(v4))
                    } else {
                        var v35 int64 = 9223372036854775807
                        _ = v35
                    }
                    _ = v34
                }
                v36 := GenS5{F0: genFn0(), F1: (v3 ^ (v2 << ((643) & 63))), F2: (v28 & (v28 >> ((809) & 63)))}
                _ = v36
            } else {
                var v37 int8 = 74
                v3 -= int32(v4)
                _ = v37
            }
            _ = fi29
        }
        v38 := GenS6{F0: (fi27 + (fi27 & (v5 & (v5 + 2147483647))))}
        _ = fi27
        _ = v28
        _ = v38
    }
    v26 = append(v26, (v5 % ((v5) | 1)))
    v5 = (v5 >> ((108) & 63))
    if true {
        if len(v26) > 0 {
            v26[0] = (v5 >> ((genFn0()) & 63))
        }
    }
    var v39 int16 = 42
    var v40 uint32 = 878
    var v41 uint = 1
    var v42 uint16 = uint16(v40)
    var v43 int32 = 1
    if len(v26) > 0 {
        v26[0] = v5
    }
    for fi44 := 0; fi44 < 9; fi44++ {
        v26 = append(v26, 1)
        v45 := GenS2{F0: -1}
        _ = fi44
        _ = v45
    }
    switch v5 {
    case -821:
        if false {
            var v46 float64 = 0.1
            var v47 int16 = v39
            v48 := GenS3{F0: v42}
            _ = v46
            _ = v47
            _ = v48
        } else {
            v49 := GenS5{F0: 2, F1: v3, F2: (v42 % (((^(v42))) | 1))}
            _ = v49
        }
    case 128:
        v43 += (^(v2))
    case 32768:
        var v50 uint32 = uint32(v4)
        _ = v50
    default:
        var v51 uint8 = v4
        _ = v51
    }
    for fi52 := 0; fi52 < 16; fi52++ {
        var v53 uint64 = 1
        _ = fi52
        _ = v53
    }
    v54 := GenS4{F0: 0, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v54
}
