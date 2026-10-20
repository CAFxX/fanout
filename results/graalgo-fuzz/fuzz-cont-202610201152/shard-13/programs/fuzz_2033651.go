package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
func genFn0() int32 {
    var v6 uint64 = 338
    v7 := []int{}
    _ = len(v7)
    if ((-1 >= -1) && true) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        switch -473 {
        case 127:
            v7 = append(v7, -2147483648)
        }
    } else {
        var v8 bool = true
        v9 := GenS1{F0: 508}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            v7 = append(v7, fi11)
            v6 = (v6 ^ (v6 & 42))
            v6 -= v6
            _ = fi11
        }
        v12 := GenS1{F0: 0}
        v7 = append(v7, 1)
        _ = fi10
        _ = v12
    }
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 uint16, p1 int16, p2 uint32) bool {
    var v13 bool = false
    v14 := []int{945, 4294967295, 127}
    fmt.Println(p2, p2)
    fmt.Println(v13)
    var v15 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return v13
}

func main() {
    var v1 int64 = 42
    var v2 int8 = 127
    v3 := []uint32{}
    var v4 bool = true
    var v5 int64 = (v1 >> ((151) & 63))
    v16 := GenS1{F0: 4294967295}
    switch "a" {
    case "foo bar":
        var v17 int16 = -1
        _ = v17
    case "":
        if v4 {
            var v18 int64 = v5
            v5 = v5
            var v19 int = int(v5)
            _ = v18
            _ = v19
        } else {
            var v20 int = -213
            var v21 int = (v20 << ((351) & 63))
            v5 -= (v5 >> ((18446744073709551615) & 63))
            _ = v20
            _ = v21
        }
    }
    if v4 {
        var v22 string = ""
        var v23 int = 1
        _ = v22
        _ = v23
    } else {
        switch 9223372036854775807 {
        case 880:
            for fi24 := 0; fi24 < 19; fi24++ {
                var v25 string = ""
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v26 := GenS1{F0: 42}
                _ = fi24
                _ = v25
                _ = v26
            }
        case -799:
            v27 := GenS2{F0: 1}
            _ = v27
        case 128:
            for fi28 := 0; fi28 < 19; fi28++ {
                if true {
                    v3 = append(v3, 65535)
                    var v29 uint32 = 0
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 16; fi30++ {
                        switch 155 {
                        case 128:
                            v31 := GenS1{F0: 138}
                            _ = v31
                        case 1:
                            var v32 int32 = -304
                            _ = v32
                        case -2147483648:
                            var v33 float64 = float64(-2.25)
                            _ = v33
                        default:
                            var v34 int8 = v2
                            var v35 uint64 = 1
                            _ = v34
                            _ = v35
                        }
                        if ("" > "hello") {
                            fmt.Println(v4)
                            v5 += 9223372036854775807
                            v5 += (v1 ^ v5)
                        } else {
                            v2 = (v2 % (((v2 / ((-128) | 1))) | 1))
                            var v36 uint64 = 2
                            _ = v36
                        }
                        _ = fi30
                    }
                }
                fmt.Println(v1, v5, v1)
                _ = fi28
            }
        }
        fmt.Println(v1)
        v5 = 0
    }
    v4 = v4
    v5 -= (v5 * (v5 - v5))
    v37 := GenS1{F0: 691}
    v38 := GenS1{F0: 913}
    fmt.Println(v1, v1)
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    var v39 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
