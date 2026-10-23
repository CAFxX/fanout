package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint16, p1 uint16) bool {
    if false {
        v6 := []int{0, 0}
        var v7 uint32 = 1
        v6 = append(v6, -1)
        _ = v6
        _ = v7
    } else {
        v8 := []int{}
        if len(v8) > 0 {
            v8[0] = -1
        }
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        if false {
            var v10 float64 = 0.1
            v10 += (v10 + v10)
            _ = v10
        } else {
            v11 := []int{(-(fi9))}
            v11 = append(v11, fi9)
            v11 = append(v11, -425)
            _ = v11
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            p1 = p1
            p0 = 65535
            _ = fi12
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    return ("hello" == "hello")
}
func genFn1() bool {
    var v13 int16 = 42
    if true {
        v14 := []int{834}
        var v15 int = 996
        _ = v14
        _ = v15
    }
    _ = v13
    return (1 <= 115)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -1
    var v3 string = ((("a" + (("hello" + ("hello" + "\n\t")) + (("\n\t" + "foo bar") + ""))) + ("" + "a")) + "\n\t")
    var v4 uint64 = 18446744073709551615
    var v5 uint = 65535
    fmt.Println(v5, v3, v2)
    if false {
        for fi16 := 0; fi16 < 17; fi16++ {
            if false {
                switch v3 {
                case "":
                    var v17 int16 = 32767
                    _ = v17
                case "\n\t":
                    v4 -= v4
                case "hello":
                    v4 -= (^(v4))
                default:
                    var v18 uint8 = 105
                    _ = v18
                }
            } else {
                for fi19 := 0; fi19 < 6; fi19++ {
                    fmt.Println(v2, v5)
                    switch (fi16 | 1) {
                    case -524:
                        fmt.Println(v2, fi16)
                    case 1:
                        var v20 uint16 = 391
                        _ = v20
                    default:
                        fmt.Println(v5, fi19)
                    }
                    _ = fi19
                }
            }
            v21 := GenS1{F0: 42, F1: 127}
            _ = fi16
            _ = v21
        }
    } else {
        var v22 float64 = (float64((1.5 + (1e10 * (1.5 - -2.25)))) * 1e10)
        v23 := GenS1{F0: 42, F1: 1}
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: int8(v5), F1: 0}
    switch v3 {
    case "":
        var v25 bool = (int(v5) == -1)
        _ = v25
    default:
        v26 := GenS1{F0: 0, F1: 120}
        _ = v26
    }
    switch 4294967295 {
    case 1:
        if (995 < 1) {
            if false {
                fmt.Println(v2, v5, v4)
            } else {
                var v27 uint64 = 42
                _ = v27
            }
            v2 = (^(v2))
        }
    case 9223372036854775807:
        v28 := []int{}
        _ = v28
    default:
        v2 = (v2 - 0)
    }
    v1[65535] = false
    switch 595 {
    case 32767:
        v3 = v3
    case -1:
        var v29 bool = (("a" + (v3 + (((v3 + "a") + "foo bar") + v3))) >= v3)
        _ = v29
    }
    switch -1 {
    case 32768:
        v30 := []int{32768, 4294967295}
        _ = v30
    }
    if false {
        v4 -= (v4 - (v4 >> ((v5) & 63)))
    } else {
        v3 = ((v3 + v3) + "")
        for fi31 := 0; fi31 < 2; fi31++ {
            v4 = v4
            _ = fi31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
