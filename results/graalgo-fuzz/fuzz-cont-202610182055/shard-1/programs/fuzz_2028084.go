package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int64; F1 int8 }
func genFn0(p0 uint16, p1 string, p2 int16) uint8 {
    v4 := GenS1{F0: 10, F1: -9223372036854775808}
    v5 := []int{}
    var v6 uint16 = (^(p0))
    fmt.Println(p2, v6, v6)
    var v7 float32 = ((1.5 / -2.25) + -2.25)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint32, p1 bool) int16 {
    var v8 string = ("a" + (("hello" + ("a" + "a")) + "hello"))
    var v9 float64 = 1.5
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -413
}

func main() {
    v1 := []int{}
    var v2 float32 = 3.14159
    v3 := make(map[string]uint32)
    var v10 int16 = -265
    fmt.Println(v2)
    v10 = genFn1(307, ((!false) || true))
    var v11 uint32 = 2147483647
    var v12 string = ("" + ("hello" + ""))
    v13 := GenS1{F0: 1, F1: -570}
    var v14 uint8 = 127
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            fmt.Println(v11, v10, v10)
            var v17 float32 = 1.5
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: 42, F1: 42}
        _ = fi15
        _ = v18
    }
    if (("foo bar" + "a") >= ("\n\t" + v12)) {
        var v19 int64 = 1
        for fi20 := 0; fi20 < 16; fi20++ {
            for fi21 := 0; fi21 < 18; fi21++ {
                if (!(!false)) {
                    _ = len(v3)
                    if ((v19 ^ v19) < (-(v19))) {
                        v22 := GenS1{F0: -112, F1: v19}
                        if true {
                            v23 := GenS2{F0: v19, F1: -128}
                            _ = v23
                        } else {
                            v10 = (v10 << ((0) & 63))
                        }
                        _ = v22
                    }
                    v2 -= (float32(0.1) * (((v2 - 0.1) / 1.0) / v2))
                    v11 = 519
                }
                switch fi21 {
                case 342:
                    fmt.Println(v11)
                default:
                    var v24 string = v12
                    _ = v24
                }
                fmt.Println(v10, v11, v2)
                _ = fi21
            }
            _ = fi20
        }
        v12 = v12
        _ = v19
    }
    switch 0 {
    case 128:
        fmt.Println(v2)
    case -489:
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 string = v12
            fmt.Println(v11)
            _ = fi25
            _ = v26
        }
    case -9223372036854775808:
        for fi27 := 0; fi27 < 4; fi27++ {
            v2 = float32(float64(float32(1.5)))
            _ = fi27
        }
    default:
        for fi28 := 0; fi28 < 3; fi28++ {
            if (uint16(v14) < 65535) {
                v2 += v2
                v29 := GenS2{F0: 432, F1: -5}
                _ = v29
            } else {
                v12 = "\n\t"
            }
            var v30 int = fi28
            v30 += fi28
            _ = fi28
            _ = v30
        }
    }
    switch ((v12 + v12) + v12) {
    case "\n\t":
        switch v12 {
        case "foo bar":
            var v31 uint64 = 2
            _ = v31
        case "\n\t":
            var v32 bool = true
            _ = v32
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            v11 += v11
        }
    default:
        var v33 int64 = -714
        _ = v33
    }
    if false {
        var v34 int64 = 0
        for fi35 := 0; fi35 < 4; fi35++ {
            if (550 <= 4294967295) {
                var v36 uint32 = 1
                if (!(42 > 32767)) {
                    v36 = (^(v36))
                    for fi37 := 0; fi37 < 12; fi37++ {
                        var v38 int8 = 100
                        _ = fi37
                        _ = v38
                    }
                }
                _ = v36
            } else {
                var v39 float64 = -2.25
                if true {
                    v10 += (v10 | (v10 | -1))
                }
                _ = v39
            }
            v40 := GenS1{F0: 127, F1: v34}
            var v41 int8 = 127
            _ = fi35
            _ = v40
            _ = v41
        }
        fmt.Println(v14, v11, v10)
        _ = v34
    } else {
        if len(v1) > 0 {
            v1[0] = -135
        }
        v12 = v12
    }
    var v42 bool = true
    switch 0 {
    case 4294967295:
        for fi43 := 0; fi43 < 7; fi43++ {
            if v42 {
                v12 = "foo bar"
                v12 = v12
            }
            v1 = append(v1, 42)
            v11 -= 1
            _ = fi43
        }
    }
    v44 := GenS2{F0: -976, F1: 59}
    var v45 bool = v42
    v45 = (1 <= -664)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v42
    _ = v44
    _ = v45
}
