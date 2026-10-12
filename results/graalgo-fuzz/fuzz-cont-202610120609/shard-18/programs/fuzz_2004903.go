package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int16 }
func genFn0(p0 int8, p1 uint8) uint64 {
    if (42 < 128) {
        p0 -= -41
    }
    if (!(585 >= 243)) {
        v4 := GenS1{F0: 0, F1: 857}
        p1 = 58
        if ((182 > 42) || ("a" != "a")) {
            v5 := GenS1{F0: p0, F1: 42}
            if false {
                switch 127 {
                case 105:
                    switch "" {
                    case "a":
                        p1 = (p1 % ((97) | 1))
                    case "hello":
                        if false {
                            fmt.Println(p0, p0)
                            fmt.Println(p0, p1)
                            var v6 float32 = 1e10
                            _ = v6
                        } else {
                            fmt.Println(p0, p1)
                            v7 := []int{}
                            v7 = append(v7, 1)
                            _ = v7
                        }
                    case "\n\t":
                        var v8 float32 = (-2.25 - (0.0 + float32(3.14159)))
                        _ = v8
                    }
                case 963:
                    p0 = p0
                default:
                    v9 := []int{}
                    _ = v9
                }
            }
            v10 := GenS1{F0: (^(p0)), F1: 32767}
            _ = v10
            _ = v5
        }
        _ = v4
    } else {
        if true {
            v11 := []int{32767}
            switch 42 {
            case 4294967295:
                fmt.Println(p1, p1)
            case -1:
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            default:
                var v12 uint16 = 2
                _ = v12
            }
            fmt.Println(p0, p1)
            _ = v11
        }
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 int32 = -1
    v2 := []uint32{1, 2}
    v3 := make(map[int]uint32)
    switch "foo bar" {
    case "hello":
        v1 -= (v1 + (-(v1)))
    case "":
        if false {
            v13 := GenS1{F0: 1, F1: -292}
            _ = len(v3)
            if (127 > 1) {
                fmt.Println(v1)
                for fi14 := 0; fi14 < 16; fi14++ {
                    for fi15 := 0; fi15 < 14; fi15++ {
                        var v16 bool = false
                        v1 = v1
                        var v17 int64 = -1
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = len(v2)
                    _ = fi14
                }
            } else {
                v1 -= (v1 / ((v1) | 1))
            }
            _ = v13
        } else {
            fmt.Println(v1)
        }
    case "a":
        if ("foo bar" > "") {
            v18 := GenS1{F0: 100, F1: 94}
            for fi19 := 0; fi19 < 12; fi19++ {
                if (true && true) {
                    v2 = append(v2, 42)
                    v20 := GenS1{F0: 100, F1: 102}
                    v1 -= v1
                    _ = v20
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v1 = v1
                }
                v1 = -1
                v21 := GenS2{F0: "hello"}
                _ = fi19
                _ = v21
            }
            fmt.Println(v1)
            _ = v18
        } else {
            v22 := GenS1{F0: 127, F1: 814}
            v23 := GenS2{F0: "foo bar"}
            _ = v22
            _ = v23
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v1 -= v1
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    var v24 uint16 = 2
    switch "hello" {
    case "hello":
        v1 += v1
    default:
        switch 153 {
        case 0:
            v1 = (v1 ^ v1)
        default:
            v24 += v24
        }
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := GenS3{F0: v1, F1: 1000}
        switch "foo bar" {
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "":
            var v27 uint32 = 877
            _ = v27
        default:
            for fi28 := 0; fi28 < 2; fi28++ {
                v29 := GenS1{F0: 42, F1: (-(fi25))}
                if (0.0 == ((1.5 * (-2.25 * float64(1.5))) + 3.14159)) {
                    fmt.Println(fi25, v1, fi25)
                    switch (("\n\t" + "hello") + "\n\t") {
                    case "hello":
                        v24 = (v24 >> ((0) & 63))
                    case "":
                        var v30 string = "foo bar"
                        _ = v30
                    }
                    v1 = (v1 & -2147483648)
                    v24 -= (v24 | 0)
                } else {
                    fmt.Println(v24, v24)
                    var v31 uint8 = 127
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v31
                }
                _ = fi28
                _ = v29
            }
        }
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
