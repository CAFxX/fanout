package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint }

func main() {
    var v1 int8 = 100
    v2 := []int{9223372036854775807, 9223372036854775807, -2147483648}
    var v3 uint8 = 0
    var v4 bool = true
    var v5 string = "a"
    switch 9223372036854775807 {
    case 9223372036854775807:
        switch 1 {
        case 847:
            v1 = v1
        case -9223372036854775808:
            v6 := GenS1{F0: v1}
            _ = v6
        }
    case -404:
        v4 = v4
    default:
        v3 -= 255
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 127}
        v1 = v1
        v9 := GenS1{F0: (-(v1))}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float32 = 3.14159
    var v11 int16 = 1000
    if (v4 || (false || v4)) {
        v4 = ((v11 >= v11) || (1 >= 2))
        for fi12 := 0; fi12 < 11; fi12++ {
            if len(v2) > 0 {
                v2[0] = fi12
            }
            v1 = v1
            _ = fi12
        }
    }
    v1 += v1
    switch 42 {
    case 128:
        if v4 {
            v3 = v3
            if false {
                if len(v2) > 0 {
                    v2[0] = -144
                }
                v13 := GenS1{F0: (v1 << ((482) & 63))}
                _ = v13
            }
            var v14 uint = 2
            _ = v14
        }
    }
    var v15 float32 = (0.1 - ((v10 - 1.5) + v10))
    _ = len(v2)
    v10 -= float32(((1.5 * -2.25) * 1e10))
    fmt.Println(v15)
    if ((-868 < 2147483647) || (!v4)) {
        if false {
            var v16 bool = v4
            v2 = append(v2, 4294967295)
            _ = v16
        } else {
            if v4 {
                var v17 string = v5
                v18 := GenS1{F0: v1}
                if v4 {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        for fi20 := 0; fi20 < 9; fi20++ {
                            v11 = v11
                            if v4 {
                                var v21 string = v17
                                _ = v21
                            } else {
                                v11 = v11
                                var v22 uint = 2
                                if (!v4) {
                                    switch 42 {
                                    case 9223372036854775807:
                                        if len(v2) > 0 {
                                            v2[0] = (fi19 - fi20)
                                        }
                                    default:
                                        var v23 uint32 = 65535
                                        _ = v23
                                    }
                                    fmt.Println(v3, v4, v10)
                                } else {
                                    v5 = (v5 + "foo bar")
                                }
                                _ = v22
                            }
                            switch 657 {
                            case 2147483647:
                                v1 = 100
                            default:
                                var v24 uint32 = 4294967295
                                switch fi20 {
                                case 1:
                                    v25 := GenS2{F0: 42, F1: 315}
                                    _ = v25
                                case -260:
                                    if (-9223372036854775808 == -906) {
                                        v26 := GenS1{F0: (v1 & -128)}
                                        fmt.Println(v17)
                                        var v27 int = -1
                                        var v28 uint64 = 18446744073709551615
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                    v17 = ""
                                    v29 := GenS2{F0: 2, F1: 255}
                                    _ = v29
                                case 127:
                                    v4 = v4
                                default:
                                    var v30 uint16 = 0
                                    _ = v30
                                }
                                _ = v24
                            }
                            v11 = v11
                            v15 = v10
                            _ = fi20
                        }
                        var v31 int64 = 0
                        v3 -= (v3 >> ((18446744073709551615) & 63))
                        _ = fi19
                        _ = v31
                    }
                } else {
                    v11 += (-(v11))
                    fmt.Println(v4)
                }
                _ = v17
                _ = v18
            } else {
                _ = len(v2)
            }
            fmt.Println(v5, v4)
        }
    } else {
        v32 := GenS1{F0: (v1 >> ((719) & 63))}
        var v33 int64 = 0
        v34 := GenS1{F0: (v1 >> ((4294967295) & 63))}
        _ = v32
        _ = v33
        _ = v34
    }
    fmt.Println(v11, v11)
    v11 = v11
    v11 += (v11 + v11)
    v4 = (v11 != v11)
    switch 4294967295 {
    case 32767:
        v3 = (v3 & (^(v3)))
    default:
        var v35 int32 = -52
        _ = v35
    }
    v36 := GenS2{F0: 121, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
