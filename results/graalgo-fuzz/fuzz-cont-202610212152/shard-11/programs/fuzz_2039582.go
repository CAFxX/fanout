package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
func genFn0(p0 int, p1 uint32) int8 {
    var v7 int64 = 0
    fmt.Println(v7, p1, p0)
    var v8 int16 = -1
    var v9 int64 = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -128
}
func genFn1() int32 {
    var v10 float64 = -2.25
    if false {
        fmt.Println(v10)
    }
    v11 := []int{-656, 9223372036854775807, 0}
    _ = v10
    _ = v11
    return -1
}

func main() {
    var v1 float32 = ((float32((3.14159 / (3.14159 + 3.14159))) + (0.1 / 1.0)) - float32(0.0))
    var v2 uint32 = 0
    var v3 int = 32767
    var v4 int64 = -9223372036854775808
    var v5 float32 = (v1 * 3.14159)
    var v6 int32 = 1
    var v12 uint16 = 0
    v4 = -802
    v13 := GenS1{F0: 894}
    if true {
        var v14 uint64 = 2
        var v15 int = v3
        for fi16 := 0; fi16 < 12; fi16++ {
            v3 += v15
            v3 = fi16
            switch "foo bar" {
            case "hello":
                if ((^(v12)) < (v12 + v12)) {
                    if true {
                        var v17 string = ""
                        if false {
                            if (9223372036854775807 > (v4 + v4)) {
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    var v19 bool = false
                                    _ = fi18
                                    _ = v19
                                }
                            } else {
                                v20 := GenS1{F0: (v14 >> ((719) & 63))}
                                var v21 uint = 18446744073709551615
                                v15 += (fi16 << ((uint(v12)) & 63))
                                _ = v20
                                _ = v21
                            }
                            if true {
                                var v22 uint64 = (v14 >> ((938) & 63))
                                for fi23 := 0; fi23 < 14; fi23++ {
                                    v12 = (v12 | 2)
                                    if false {
                                        fmt.Println(v5, v15)
                                        v4 -= (^(v4))
                                    }
                                    _ = fi23
                                }
                                v5 -= (v1 * (1e10 - v1))
                                _ = v22
                            }
                        }
                        var v24 bool = false
                        _ = v17
                        _ = v24
                    } else {
                        v14 = v14
                        fmt.Println(v6, v15, v4)
                        switch "" {
                        case "a":
                            for fi25 := 0; fi25 < 8; fi25++ {
                                var v26 float32 = -2.25
                                switch "hello" {
                                case "\n\t":
                                    v27 := []int{-1}
                                    _ = v27
                                case "":
                                    v28 := []int{fi16}
                                    _ = v28
                                case "a":
                                    v29 := GenS2{F0: "\n\t"}
                                    _ = v29
                                }
                                _ = fi25
                                _ = v26
                            }
                        default:
                            fmt.Println(v3, v6)
                        }
                    }
                    var v30 uint32 = v2
                    _ = v30
                }
            case "foo bar":
                if (v14 > v14) {
                    v31 := []int{v15}
                    v12 = (v12 + 0)
                    _ = v31
                }
            case "":
                var v32 uint = uint(v12)
                _ = v32
            }
            _ = fi16
        }
        _ = v14
        _ = v15
    } else {
        v1 = v5
    }
    v33 := GenS2{F0: ("a" + "a")}
    var v34 uint16 = v12
    fmt.Println(v3, v12)
    v35 := []int{-2147483648, (-(v3)), 42}
    switch 32768 {
    case 32768:
        v5 += v5
    }
    v1 = (v5 * (v5 - float32(1e10)))
    var v36 bool = false
    v34 = v12
    v35 = append(v35, v3)
    if (v36 || (493 != 2)) {
        v4 -= v4
    } else {
        fmt.Println(v5, v6, v12)
        v5 = v1
    }
    if len(v35) > 0 {
        v35[0] = 32767
    }
    v3 = (v3 - v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
