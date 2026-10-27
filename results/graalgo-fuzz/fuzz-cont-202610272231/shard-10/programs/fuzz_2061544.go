package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int64 }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 6; fi7++ {
        if (2 != 4294967295) {
            if false {
                v8 := []int{401, (-(fi7))}
                _ = v8
            }
            var v9 int16 = -32768
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 uint16 = 421
                var v12 bool = false
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        }
        if false {
            var v13 uint64 = 2
            _ = v13
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int64 = -466
            var v16 float64 = 0.0
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi7
    }
    v17 := []int{9223372036854775807, 32768, 165}
    _ = v17
    return 210
}
func genFn1(p0 int32, p1 int64) int8 {
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 bool = (!((42 == 127) && false))
        for fi20 := 0; fi20 < 18; fi20++ {
            if (!(((true || v19) || v19) && (255 != 0))) {
                v21 := []int{fi18}
                var v22 uint32 = 0
                _ = v21
                _ = v22
            }
            var v23 uint64 = 18446744073709551615
            fmt.Println(fi20, v19)
            _ = fi20
            _ = v23
        }
        _ = fi18
        _ = v19
    }
    switch 2147483647 {
    case 942:
        fmt.Println(p0, p1)
    default:
        p1 += p1
    }
    var v24 string = "hello"
    v24 = ((v24 + "\n\t") + "a")
    if ("\n\t" <= "\n\t") {
        v24 = "a"
    }
    _ = p0
    _ = p1
    _ = v24
    return 90
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2}
    var v3 int32 = 53
    var v4 float32 = 3.14159
    var v5 float32 = 0.1
    v6 := []uint32{2147483648, 1}
    v4 -= v4
    v2 = append(v2, 65535)
    if ((v3 & 42) != (v3 - (v3 % (((v3 >> ((65535) & 63))) | 1)))) {
        var v25 uint8 = 195
        if (true && true) {
            v5 = v5
            for fi26 := 0; fi26 < 5; fi26++ {
                switch fi26 {
                case -2147483648:
                    if false {
                        var v27 uint64 = 421
                        _ = v27
                    }
                default:
                    v28 := GenS1{F0: false, F1: 62}
                    _ = v28
                }
                _ = fi26
            }
            v1[65535] = 775
        } else {
            v29 := v1[genFn0()]
            _ = v29
        }
        _ = v25
    } else {
        for fi30 := 0; fi30 < 7; fi30++ {
            fmt.Println(v4)
            _ = fi30
        }
        fmt.Println(v3)
        for fi31 := 0; fi31 < 3; fi31++ {
            var v32 uint64 = 1
            for fi33 := 0; fi33 < 16; fi33++ {
                switch fi31 {
                case 4294967295:
                    var v34 float64 = 1e10
                    if false {
                        var v35 int = -365
                        if false {
                            v34 = (v34 / (1.5 / float64(-2.25)))
                            v36 := GenS1{F0: (1000 <= 423), F1: 1}
                            _ = v36
                        }
                        _ = v35
                    } else {
                        var v37 bool = false
                        v5 -= v5
                        v3 -= v3
                        _ = v37
                    }
                    _ = v34
                case 762:
                    v4 = (v4 + 1e10)
                default:
                    v5 -= (-2.25 * v4)
                }
                _ = fi33
            }
            v38 := GenS1{F0: (!(9223372036854775807 < -377)), F1: 32767}
            _ = fi31
            _ = v32
            _ = v38
        }
    }
    v39 := GenS1{F0: ((!false) || (true || (870 > 42))), F1: 1}
    var v40 uint64 = 1
    if true {
        v41 := GenS2{F0: -1}
        v40 -= v40
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v41
    } else {
        v40 -= v40
    }
    _ = len(v6)
    var v42 string = ("" + ("foo bar" + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
    _ = v6
}
