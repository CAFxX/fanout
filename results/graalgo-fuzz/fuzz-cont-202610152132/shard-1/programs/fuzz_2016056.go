package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 bool }
func genFn0() int {
    if true {
        if (!(0.1 != 0.0)) {
            var v5 string = "a"
            for fi6 := 0; fi6 < 16; fi6++ {
                switch fi6 {
                case 2147483647:
                    v7 := []int{(fi6 * (fi6 & (fi6 + -701))), -761, fi6, fi6}
                    _ = v7
                case 9223372036854775807:
                    v8 := GenS1{F0: 2, F1: 2}
                    _ = v8
                default:
                    v5 = ("a" + ((v5 + (v5 + ("a" + v5))) + ("foo bar" + ((v5 + v5) + v5))))
                }
                if true {
                    for fi9 := 0; fi9 < 10; fi9++ {
                        fmt.Println(fi9)
                        _ = fi9
                    }
                    fmt.Println(v5, fi6)
                } else {
                    v5 = v5
                    v10 := GenS1{F0: 2, F1: 2}
                    _ = v10
                }
                for fi11 := 0; fi11 < 1; fi11++ {
                    v12 := GenS1{F0: 255, F1: 41}
                    v5 = (v5 + "\n\t")
                    var v13 float64 = float64(-2.25)
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi6
            }
            switch -227 {
            case 32767:
                switch "" {
                case "":
                    var v14 string = (v5 + v5)
                    _ = v14
                case "a":
                    v15 := []int{32767, 9223372036854775807, -1, 1}
                    _ = v15
                case "hello":
                    v16 := []int{957, 32768, 2147483647, 9223372036854775807}
                    _ = v16
                default:
                    if false {
                        switch 32768 {
                        case -9223372036854775808:
                            var v17 string = ("a" + v5)
                            _ = v17
                        }
                    } else {
                        v5 = v5
                        var v18 int64 = 9223372036854775807
                        v19 := []int{0, -774}
                        _ = v18
                        _ = v19
                    }
                }
            }
            _ = v5
        }
    }
    var v20 string = "hello"
    v21 := []int{}
    _ = v20
    _ = v21
    return 128
}
func genFn1(p0 int64, p1 bool, p2 int16) uint16 {
    if p1 {
        p0 = (p0 % ((-572) | 1))
    }
    p0 += p0
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 int64 = int64(fi22)
        fmt.Println(p0, p2)
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 uint8 = 0
    var v2 uint8 = (v1 + (v1 & v1))
    var v3 int8 = 100
    var v4 int16 = 42
    var v24 uint = 18446744073709551615
    v25 := GenS1{F0: 128, F1: 255}
    for fi26 := 0; fi26 < 13; fi26++ {
        v27 := GenS2{F0: 81, F1: fi26, F2: (v4 <= (v4 + (v4 | (v4 % ((v4) | 1)))))}
        for fi28 := 0; fi28 < 13; fi28++ {
            v2 += (v1 - 255)
            v29 := []int{(fi28 % ((fi28) | 1)), (fi26 ^ (fi28 * genFn0())), fi26, (fi26 / (((-(fi26))) | 1))}
            fmt.Println(v1, v1)
            v4 = v4
            _ = fi28
            _ = v29
        }
        fmt.Println(v4, v24)
        _ = fi26
        _ = v27
    }
    var v30 uint8 = v2
    v31 := GenS2{F0: 1, F1: 127, F2: (v24 > v24)}
    var v32 float64 = 1e10
    fmt.Println(v30)
    var v33 int8 = (-(v3))
    fmt.Println(v32)
    v34 := GenS1{F0: 137, F1: v2}
    if false {
        switch -1 {
        case 216:
            fmt.Println(v1)
        case 452:
            switch genFn0() {
            case 1:
                switch "a" {
                case "foo bar":
                    v35 := []int{1, 15, 42}
                    _ = v35
                default:
                    var v36 uint = (v24 | v24)
                    _ = v36
                }
            default:
                var v37 uint8 = 209
                _ = v37
            }
        case 9223372036854775807:
            fmt.Println(v4, v24, v2)
        }
    } else {
        if true {
            for fi38 := 0; fi38 < 11; fi38++ {
                for fi39 := 0; fi39 < 2; fi39++ {
                    v24 = v24
                    v33 = (v3 % ((127) | 1))
                    _ = fi39
                }
                var v40 float64 = (v32 - 3.14159)
                var v41 int8 = 1
                _ = fi38
                _ = v40
                _ = v41
            }
            switch "foo bar" {
            case "a":
                v32 = 1.5
            case "hello":
                for fi42 := 0; fi42 < 13; fi42++ {
                    v43 := []int{fi42, 128}
                    var v44 float32 = 0.1
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            }
        }
    }
    v32 = float64(float32(1e10))
    v24 -= 18446744073709551615
    var v45 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v45
}
