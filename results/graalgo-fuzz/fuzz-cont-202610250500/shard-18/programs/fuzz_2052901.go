package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int64 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint64; F2 string }
func genFn0() uint {
    var v6 uint16 = 2
    var v7 int8 = -98
    switch 2147483647 {
    case -77:
        fmt.Println(v7, v7)
    case 8:
        v8 := []int{440}
        _ = v8
    default:
        v9 := GenS1{F0: 2, F1: (0 > 0), F2: 0}
        _ = v9
    }
    v6 += 905
    var v10 uint8 = 0
    _ = v10
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 int8 = 32
    v2 := make(map[int]int)
    var v3 uint32 = 42
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 815
    v11 := GenS2{F0: v1, F1: -32768}
    var v12 int = 42
    if false {
        for fi13 := 0; fi13 < 9; fi13++ {
            v12 -= (-(v12))
            _ = fi13
        }
        v14 := GenS1{F0: 866, F1: (!false), F2: 42}
        v3 -= 65535
        _ = v14
    } else {
        var v15 float64 = float64(float32(3.14159))
        switch "" {
        case "a":
            v5 = (v5 & (v5 << ((2) & 63)))
        case "foo bar":
            v16 := GenS3{F0: int16(v1), F1: uint8(v3)}
            _ = v16
        }
        _ = v15
    }
    if false {
        var v17 uint64 = 2
        v12 = int(v5)
        _ = v17
    }
    switch v12 {
    case 4294967295:
        v1 += v1
    case 1:
        if true {
            v18 := GenS4{F0: 42, F1: 364, F2: "a"}
            if false {
                v19 := []int{int(v3), v12}
                fmt.Println(v12)
                _ = v19
            }
            switch 1 {
            case -2147483648:
                v5 += v5
            }
            _ = v18
        }
    case -2147483648:
        var v20 float64 = 0.0
        _ = v20
    default:
        v21 := GenS4{F0: -800, F1: 350, F2: "a"}
        _ = v21
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int8 = int8(v12)
        if (v5 > v5) {
            if false {
                fmt.Println(v3, v1, fi22)
                v23 -= v1
            }
            var v24 string = (((("hello" + (("a" + ("hello" + (("foo bar" + ("\n\t" + "a")) + ""))) + "")) + ("hello" + ("foo bar" + "a"))) + ("" + "foo bar")) + "foo bar")
            v23 += v23
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    delete(v2, v12)
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := []int{}
        for fi27 := 0; fi27 < 20; fi27++ {
            v5 -= v5
            if (!((uint8(v3) != 1) && (346 <= 18446744073709551615))) {
                var v28 uint64 = 2
                _ = v28
            } else {
                if false {
                    if false {
                        v3 = (^(v3))
                        v1 -= 127
                        var v29 int16 = -32768
                        _ = v29
                    }
                } else {
                    fmt.Println(v1, fi27, v3)
                    if (1 < 1) {
                        v5 += v5
                        v1 = (v1 >> ((0) & 63))
                    }
                    v30 := GenS2{F0: 127, F1: -32768}
                    var v31 uint32 = (v3 << ((2) & 63))
                    _ = v30
                    _ = v31
                }
            }
            if false {
                v12 = fi25
            } else {
                v32 := GenS3{F0: 0, F1: 111}
                if (1.5 <= 1.5) {
                    if false {
                        v1 += v1
                    } else {
                        var v33 uint64 = 930
                        v5 += v5
                        _ = v33
                    }
                    if false {
                        if (321 == 0) {
                            v34 := v4[v3]
                            var v35 int8 = v1
                            _ = v34
                            _ = v35
                        } else {
                            v12 -= (v12 + (fi27 | 4294967295))
                        }
                        var v36 uint = 42
                        v36 -= 0
                        _ = v36
                    } else {
                        var v37 uint16 = 850
                        var v38 int8 = v1
                        _ = v37
                        _ = v38
                    }
                    fmt.Println(fi25)
                } else {
                    var v39 int32 = 841
                    _ = v39
                }
                _ = v32
            }
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    var v40 float32 = (float32((((1.5 * 0.1) / -2.25) + ((3.14159 + (1.5 + -2.25)) * (3.14159 + (0.0 / 1.0))))) - (-2.25 + float32((float64(1.5) + 0.0))))
    var v41 float64 = float64((v40 * v40))
    v3 += 42
    var v42 float64 = (0.1 / v41)
    for fi43 := 0; fi43 < 18; fi43++ {
        switch (-(fi43)) {
        case 42:
            v12 = v12
        }
        v44 := v4[v3]
        _ = fi43
        _ = v44
    }
    var v45 uint32 = v3
    if (871 < 2) {
        var v46 int16 = -32768
        v41 = v41
        _ = v46
    }
    fmt.Println(v41, v42)
    v47 := GenS1{F0: 255, F1: (!false), F2: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v45
    _ = v47
    _ = v5
}
