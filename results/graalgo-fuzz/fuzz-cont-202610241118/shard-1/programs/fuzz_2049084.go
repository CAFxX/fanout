package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int; F1 bool; F2 string }
func genFn0(p0 uint8) uint8 {
    var v4 int = -653
    v4 += (v4 + (v4 - 900))
    for fi5 := 0; fi5 < 7; fi5++ {
        if true {
            v6 := []int{1, (v4 % (((fi5 & int(p0))) | 1))}
            p0 = 2
            _ = v6
        }
        fmt.Println(p0, v4)
        _ = fi5
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v4
    return p0
}
func genFn1(p0 int32, p1 uint, p2 int8) uint8 {
    for fi7 := 0; fi7 < 8; fi7++ {
        if true {
            v8 := []int{fi7}
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = v8
        } else {
            v9 := []int{}
            v10 := GenS1{F0: 879, F1: 404, F2: 42}
            switch 9223372036854775807 {
            case 32768:
                p1 = p1
            default:
                for fi11 := 0; fi11 < 17; fi11++ {
                    if (true && false) {
                        var v12 uint16 = 168
                        v13 := GenS1{F0: -1, F1: 158, F2: 32767}
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS2{F0: (^(p1)), F1: 42, F2: int8(p0)}
                    var v15 uint64 = 1
                    v16 := GenS1{F0: 601, F1: 42, F2: int16(fi7)}
                    p0 = (p0 - p0)
                    _ = fi11
                    _ = v14
                    _ = v15
                    _ = v16
                }
            }
            _ = v10
            _ = v9
        }
        v17 := GenS3{F0: 42, F1: (p0 * (p0 - (p0 + p0)))}
        p0 = (p0 + p0)
        _ = fi7
        _ = v17
    }
    switch -2147483648 {
    case -350:
        v18 := GenS4{F0: 32768, F1: (!false), F2: "hello"}
        _ = v18
    default:
        for fi19 := 0; fi19 < 5; fi19++ {
            if true {
                v20 := []int{fi19, (fi19 * (fi19 & (^(fi19)))), (-(fi19)), int(p1)}
                if (!(true || false)) {
                    v21 := GenS1{F0: -423, F1: 4294967295, F2: 0}
                    var v22 float32 = ((0.1 / 0.1) - (3.14159 + 1.5))
                    _ = v21
                    _ = v22
                } else {
                    switch fi19 {
                    case -964:
                        for fi23 := 0; fi23 < 14; fi23++ {
                            p2 -= (p2 + p2)
                            v24 := GenS4{F0: fi19, F1: (42 != (p2 % (((^(p2))) | 1))), F2: (("foo bar" + "\n\t") + ("hello" + ("\n\t" + "\n\t")))}
                            _ = fi23
                            _ = v24
                        }
                    case -2147483648:
                        if len(v20) > 0 {
                            v20[0] = (fi19 / ((0) | 1))
                        }
                    default:
                        _ = len(v20)
                    }
                    p0 = p0
                }
                _ = v20
            }
            _ = fi19
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 int32 = -634
    v2 := []string{"foo bar"}
    var v3 int8 = 127
    v2 = append(v2, "a")
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 float32 = float32(1.5)
        switch ("a" + "a") {
        case "foo bar":
            v3 += -96
        }
        for fi27 := 0; fi27 < 4; fi27++ {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            var v28 uint32 = 2
            v3 = (v3 << ((42) & 63))
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v29 := GenS4{F0: 2147483647, F1: (1 <= 32767), F2: (("a" + ("\n\t" + "")) + ("\n\t" + "hello"))}
    if false {
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 int8 = v3
            _ = fi30
            _ = v31
        }
        v1 = (v1 * v1)
    }
    switch 4294967295 {
    case 32768:
        var v32 float32 = 0.1
        _ = v32
    case 585:
        var v33 float32 = ((float32(float64(float32(0.0))) - (1.5 * 0.1)) / (3.14159 - (1.5 / (0.1 - 1e10))))
        _ = v33
    case -867:
        var v34 int16 = 42
        _ = v34
    default:
        for fi35 := 0; fi35 < 1; fi35++ {
            if ((3.14159 - -2.25) >= 1e10) {
                v3 -= 1
                for fi36 := 0; fi36 < 2; fi36++ {
                    v3 += (v3 ^ v3)
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    var v37 float64 = float64(3.14159)
                    _ = fi36
                    _ = v37
                }
            }
            _ = fi35
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        var v38 uint8 = 246
        v39 := GenS4{F0: 128, F1: true, F2: ("a" + "")}
        var v40 uint16 = 1
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
}
