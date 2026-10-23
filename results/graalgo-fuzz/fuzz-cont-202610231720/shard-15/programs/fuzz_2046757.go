package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
func genFn0(p0 uint16, p1 bool) uint64 {
    switch -1 {
    case 32767:
        v7 := []int{}
        _ = v7
    case 238:
        if (42 == 539) {
            fmt.Println(p1)
            var v8 int = -163
            if false {
                var v9 int32 = -492
                switch "foo bar" {
                case "foo bar":
                    v8 += 2147483647
                case "\n\t":
                    if ((1e10 == 0.0) && (p1 || (!p1))) {
                        p1 = ((!(134 == v8)) || p1)
                        switch v8 {
                        case 1:
                            v10 := []int{(v8 + v8), v8, v8}
                            _ = v10
                        case 32768:
                            fmt.Println(v9)
                        case 42:
                            var v11 int32 = (v9 ^ v9)
                            _ = v11
                        default:
                            fmt.Println(v9, v9)
                        }
                        var v12 uint64 = 4294967295
                        _ = v12
                    }
                }
                var v13 int64 = 0
                _ = v13
                _ = v9
            } else {
                fmt.Println(v8, p1, v8)
                if p1 {
                    fmt.Println(v8, v8)
                    p1 = (3.14159 == (float32(((0.1 / ((1.5 * ((((((0.0 + (0.1 - (0.1 - (1.5 - 3.14159)))) / (float64((0.0 + (0.1 / ((3.14159 / 1.0) * 3.14159)))) / 3.14159)) * -2.25) * -2.25) / -2.25) + 0.1)) * (0.0 - float64((0.1 / 1.5))))) + 3.14159)) / 1.5))
                    switch "foo bar" {
                    case "foo bar":
                        var v14 string = ("hello" + "")
                        _ = v14
                    case "\n\t":
                        for fi15 := 0; fi15 < 5; fi15++ {
                            if false {
                                v16 := []int{}
                                if (1000 > 0) {
                                    var v17 bool = (p1 && p1)
                                    _ = v17
                                }
                                _ = v16
                            }
                            var v18 int8 = -1
                            for fi19 := 0; fi19 < 11; fi19++ {
                                v20 := []int{fi19, (v8 | (^(fi19))), (v8 >> ((65535) & 63)), 4294967295}
                                _ = fi19
                                _ = v20
                            }
                            _ = fi15
                            _ = v18
                        }
                    }
                }
                v21 := []int{32768, (v8 ^ (v8 * -11))}
                _ = v21
            }
            _ = v8
        }
    case -9223372036854775808:
        fmt.Println(p1)
    default:
        p1 = (!false)
    }
    var v22 uint = 42
    switch "foo bar" {
    case "a":
        v22 += (v22 + (v22 ^ 255))
    }
    v23 := GenS1{F0: -556, F1: 56}
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 42
}
func genFn1(p0 uint8) bool {
    var v24 uint16 = 65535
    if false {
        if ((int(p0) == 0) && ((1e10 * (-2.25 + 3.14159)) == 0.0)) {
            var v25 string = "a"
            v26 := []int{}
            v25 = ("hello" + v25)
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 int8 = 100
            p0 += p0
            for fi29 := 0; fi29 < 4; fi29++ {
                p0 = p0
                v30 := []int{(^(fi29))}
                _ = fi29
                _ = v30
            }
            _ = fi27
            _ = v28
        }
    }
    _ = p0
    _ = v24
    return false
}

func main() {
    var v1 int64 = -298
    var v2 string = ""
    v3 := []string{}
    var v4 string = v2
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]string)
    var v31 int16 = -32768
    delete(v5, 4294967295)
    var v32 uint64 = 4294967295
    var v33 uint64 = v32
    v4 = v2
    for fi34 := 0; fi34 < 2; fi34++ {
        v2 = ("foo bar" + (v2 + v4))
        if (fi34 > 42) {
            var v35 float64 = float64((((1e10 + ((((3.14159 - float32((-2.25 - 1.5))) - (1.5 + float32(1.5))) * (1.5 * 1e10)) * (float32((float64(1e10) + 3.14159)) - ((1e10 / 0.1) + 0.1)))) * float32(float64((float32((float64(1.5) + -2.25)) * 1.5)))) * (0.0 - 3.14159)))
            _ = v35
        } else {
            fmt.Println(fi34, v32, v1)
            switch ("a" + v2) {
            case "a":
                _ = len(v3)
            default:
                v36 := GenS1{F0: fi34, F1: 2}
                _ = v36
            }
            v31 -= (v31 << ((4294967295) & 63))
        }
        _ = fi34
    }
    v37 := GenS1{F0: 51, F1: 204}
    var v38 int8 = 100
    var v39 uint8 = 224
    fmt.Println(v31, v39, v38)
    v40 := v5[42]
    v41 := GenS1{F0: 792, F1: (v39 >> ((uint(v40)) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
