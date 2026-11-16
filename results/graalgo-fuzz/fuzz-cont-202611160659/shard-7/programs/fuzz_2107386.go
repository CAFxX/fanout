package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 bool }
func genFn0(p0 int64) uint64 {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint8 = 48
        _ = fi7
        _ = v8
    }
    p0 -= p0
    p0 -= (p0 / ((p0) | 1))
    var v9 float64 = -2.25
    _ = p0
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[uint32]int)
    var v3 uint32 = 470
    var v4 string = "foo bar"
    var v5 string = (v4 + (v4 + ""))
    var v6 int64 = 0
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(v5, v3)
        fmt.Println(v3, v1, fi10)
        _ = fi10
    }
    fmt.Println(v1, v1, v5)
    v5 = "foo bar"
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: 56, F1: 0, F2: true}
        if true {
            v1 += ((0.1 + (float32(0.0) * v1)) * float32((float64(1e10) - ((float64(v1) / float64(1e10)) * (-2.25 * (3.14159 * 1e10))))))
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            v5 = "hello"
            v14 := []int{(fi13 >> ((561) & 63)), (-(fi13))}
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    v15 := []int{}
    v4 = v4
    var v16 int64 = (v6 % ((v6) | 1))
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := GenS1{F0: 42, F1: uint(v3), F2: (true || false)}
        var v19 float64 = (1e10 - float64(0.0))
        var v20 uint16 = 319
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    fmt.Println(v6, v6, v1)
    v16 = (v16 * v6)
    v2[(v3 | (v3 << ((0) & 63)))] = 91
    switch "\n\t" {
    case "foo bar":
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 uint16 = 2
            switch v5 {
            case "a":
                if (912 < 2) {
                    if false {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    }
                } else {
                    if true {
                        fmt.Println(v6, v1, v3)
                    }
                    if len(v15) > 0 {
                        v15[0] = fi21
                    }
                    delete(v2, v3)
                }
            case "foo bar":
                v23 := GenS1{F0: 18446744073709551615, F1: 42, F2: false}
                _ = v23
            case "":
                var v24 int32 = 2147483647
                _ = v24
            }
            if true {
                var v25 int = fi21
                var v26 float64 = -2.25
                _ = v25
                _ = v26
            }
            _ = fi21
            _ = v22
        }
    case "hello":
        switch "" {
        case "hello":
            var v27 uint64 = 819
            _ = v27
        case "a":
            if ((v6 >> ((2) & 63)) == 0) {
                v4 = "a"
                for fi28 := 0; fi28 < 18; fi28++ {
                    var v29 int32 = 548
                    var v30 int32 = v29
                    for fi31 := 0; fi31 < 18; fi31++ {
                        switch 2147483647 {
                        case -934:
                            v30 = v30
                        case 42:
                            v5 = v5
                        case -465:
                            v16 = v6
                        }
                        var v32 uint64 = 4294967295
                        _ = fi31
                        _ = v32
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                var v33 bool = (!false)
                _ = v33
            }
        case "":
            var v34 uint32 = v3
            _ = v34
        default:
            if (-535 < v16) {
                v1 = 3.14159
                for fi35 := 0; fi35 < 14; fi35++ {
                    _ = len(v15)
                    _ = fi35
                }
                v15 = append(v15, 9223372036854775807)
            }
        }
    default:
        for fi36 := 0; fi36 < 20; fi36++ {
            for fi37 := 0; fi37 < 9; fi37++ {
                v38 := v2[905]
                v39 := GenS1{F0: 330, F1: 4294967295, F2: (65535 != 2)}
                v40 := GenS1{F0: uint(v3), F1: 4294967295, F2: (255 >= 124)}
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
            var v41 float32 = 0.1
            _ = fi36
            _ = v41
        }
    }
    var v42 uint16 = 2
    fmt.Println(v6, v5, v1)
    if true {
        fmt.Println(v4, v42)
    } else {
        v16 = v6
    }
    v42 -= v42
    delete(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
