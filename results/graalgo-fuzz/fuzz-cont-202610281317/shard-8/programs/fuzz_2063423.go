package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int; F1 int }
func genFn0(p0 uint, p1 uint8) uint32 {
    var v7 uint = p0
    v7 = (p0 << ((v7) & 63))
    switch "foo bar" {
    case "a":
        switch -2147483648 {
        case 2147483647:
            var v8 uint8 = p1
            _ = v8
        case 32768:
            for fi9 := 0; fi9 < 2; fi9++ {
                switch fi9 {
                case -104:
                    for fi10 := 0; fi10 < 10; fi10++ {
                        v11 := GenS1{F0: -1, F1: (true && true)}
                        _ = fi10
                        _ = v11
                    }
                case -707:
                    var v12 float64 = (1e10 / ((((1e10 / 1e10) * 1e10) + 1.5) * (1e10 / 3.14159)))
                    _ = v12
                case -2147483648:
                    v7 += (v7 >> ((18446744073709551615) & 63))
                }
                _ = fi9
            }
        case 42:
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := []int{(fi13 * fi13), (-(fi13))}
                if len(v14) > 0 {
                    v14[0] = (^(fi13))
                }
                _ = fi13
                _ = v14
            }
        }
    case "":
        var v15 uint16 = 255
        _ = v15
    default:
        if false {
            v16 := GenS1{F0: 0, F1: true}
            for fi17 := 0; fi17 < 12; fi17++ {
                if ((("hello" + "") > "a") && (100 > 100)) {
                    var v18 uint16 = 680
                    if true {
                        fmt.Println(p0)
                    }
                    _ = v18
                } else {
                    fmt.Println(v7, v7, p1)
                    v7 -= 657
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v20 := GenS2{F0: 4294967295}
                        if false {
                            v21 := []int{fi19}
                            var v22 int16 = 1
                            _ = v21
                            _ = v22
                        } else {
                            v23 := []int{}
                            v7 = v7
                            fmt.Println(fi19, p0, p0)
                            _ = v23
                        }
                        _ = fi19
                        _ = v20
                    }
                }
                fmt.Println(fi17)
                var v24 int32 = 1
                _ = fi17
                _ = v24
            }
            p1 += (p1 >> ((65535) & 63))
            _ = v16
        }
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return 730
}
func genFn1() int {
    var v25 int8 = -81
    var v26 float32 = 0.0
    var v27 uint16 = 2
    for fi28 := 0; fi28 < 18; fi28++ {
        v29 := GenS1{F0: 0, F1: (fi28 < (fi28 + (fi28 - (fi28 * (fi28 | (fi28 >> ((406) & 63)))))))}
        v30 := []int{(fi28 * (fi28 / ((fi28) | 1))), fi28}
        _ = fi28
        _ = v29
        _ = v30
    }
    v25 -= (-(v25))
    _ = v25
    _ = v26
    _ = v27
    return 1
}

func main() {
    var v1 uint64 = 1
    var v2 uint64 = v1
    var v3 uint64 = v2
    v4 := []uint32{808}
    v5 := make(map[string]string)
    var v6 uint64 = (v2 & (v2 ^ v2))
    v31 := v5["a"]
    switch v31 {
    case "hello":
        v31 = v31
    case "foo bar":
        fmt.Println(v3)
    case "\n\t":
        fmt.Println(v2, v3)
    default:
        v32 := GenS3{F0: -1, F1: -32768}
        _ = v32
    }
    v33 := GenS4{F0: 1, F1: genFn1()}
    v3 -= v1
    v34 := GenS2{F0: genFn0(899, 242)}
    _ = len(v4)
    switch v31 {
    case "\n\t":
        v4 = append(v4, 4294967295)
    case "a":
        if len(v4) > 0 {
            v4[0] = 324
        }
    case "":
        switch v31 {
        case "\n\t":
            v35 := GenS2{F0: genFn0(235, 1)}
            _ = v35
        default:
            if false {
                _ = len(v5)
                v4 = append(v4, 575)
            } else {
                v36 := GenS2{F0: genFn0(483, 1)}
                v37 := GenS4{F0: genFn1(), F1: -789}
                _ = v36
                _ = v37
            }
        }
    }
    if (1.5 != (-2.25 + 0.0)) {
        v31 = v31
    } else {
        var v38 int = -9223372036854775808
        v39 := v5["a"]
        switch v31 {
        case "":
            v6 = v3
        case "a":
            delete(v5, "")
        case "\n\t":
            if true {
                fmt.Println(v38, v39, v1)
                if len(v4) > 0 {
                    v4[0] = 0
                }
                for fi40 := 0; fi40 < 19; fi40++ {
                    v41 := GenS1{F0: -1, F1: false}
                    _ = fi40
                    _ = v41
                }
            }
        }
        _ = v38
        _ = v39
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        v43 := GenS1{F0: 942, F1: (false && (65535 <= 0))}
        _ = fi42
        _ = v43
    }
    var v44 uint16 = 255
    v44 += v44
    v44 = (v44 >> ((42) & 63))
    v2 = v6
    v44 = v44
    v1 -= (v6 << ((888) & 63))
    v45 := GenS2{F0: genFn0(0, 0)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
