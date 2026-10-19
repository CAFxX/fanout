package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int, p1 uint, p2 uint8) uint64 {
    var v6 uint16 = 0
    v7 := GenS1{F0: true}
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            v6 = 0
            _ = fi9
        }
        if false {
            v10 := GenS1{F0: false}
            _ = v10
        }
        _ = fi8
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}
func genFn1() uint64 {
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 float64 = 1e10
        var v13 int8 = 127
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{0}
    var v15 int16 = -524
    _ = v14
    _ = v15
    return 400
}

func main() {
    var v1 bool = (0.1 != (3.14159 - 3.14159))
    v2 := make(map[int]bool)
    var v3 uint16 = 42
    var v4 string = "foo bar"
    v5 := make(map[string]uint32)
    v16 := GenS1{F0: (2147483647 <= 0)}
    v17 := GenS1{F0: ((-2.25 <= -2.25) && v1)}
    fmt.Println(v3)
    var v18 int16 = 105
    fmt.Println(v4)
    switch "" {
    case "":
        v19 := GenS1{F0: (9223372036854775807 < 476)}
        _ = v19
    case "a":
        if (!(!(!false))) {
            v3 = v3
        } else {
            v18 += (^(v18))
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 uint8 = 127
                switch fi20 {
                case 42:
                    for fi22 := 0; fi22 < 1; fi22++ {
                        if (v18 > v18) {
                            v23 := GenS1{F0: false}
                            fmt.Println(v3, v21)
                            _ = v23
                        } else {
                            var v24 float64 = 0.0
                            _ = v24
                        }
                        if v1 {
                            var v25 int64 = 9223372036854775807
                            _ = v25
                        }
                        _ = fi22
                    }
                case 128:
                    var v26 float32 = (((-2.25 + ((1.5 - (0.0 * 1e10)) - 1e10)) + (0.0 - 0.0)) - 1e10)
                    _ = v26
                case -129:
                    for fi27 := 0; fi27 < 11; fi27++ {
                        var v28 uint16 = v3
                        _ = fi27
                        _ = v28
                    }
                default:
                    v29 := GenS1{F0: (!v1)}
                    _ = v29
                }
                _ = fi20
                _ = v21
            }
        }
    }
    v30 := []int{42, 128, 32767, 32768}
    v2[-2147483648] = false
    var v31 uint8 = 122
    for fi32 := 0; fi32 < 19; fi32++ {
        v4 = v4
        _ = fi32
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        fmt.Println(v18, v31)
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
