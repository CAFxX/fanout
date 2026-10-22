package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 int }
func genFn0(p0 uint32, p1 uint32) int8 {
    var v7 int = -986
    var v8 uint64 = 18446744073709551615
    v9 := GenS1{F0: -447}
    var v10 float32 = 1.5
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 string = ("hello" + ("" + (("\n\t" + "") + ("\n\t" + ("foo bar" + "")))))
    var v2 uint16 = 2
    v3 := []int{9223372036854775807}
    var v4 float32 = float32(0.0)
    var v5 int8 = 127
    v6 := []int{9223372036854775807, -606}
    v2 += v2
    switch -9223372036854775808 {
    case 32768:
        for fi11 := 0; fi11 < 1; fi11++ {
            if (genFn0(65535, 2147483647) >= (v5 << ((18446744073709551615) & 63))) {
                v2 = 2
                for fi12 := 0; fi12 < 19; fi12++ {
                    if (127 != 42) {
                        switch fi11 {
                        case 555:
                            v13 := GenS1{F0: 41}
                            _ = v13
                        case 1:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    }
                    var v14 float32 = 0.0
                    v15 := GenS1{F0: 42}
                    _ = fi12
                    _ = v14
                    _ = v15
                }
            } else {
                if false {
                    v1 = "a"
                }
                for fi16 := 0; fi16 < 20; fi16++ {
                    fmt.Println(v5, v5, fi11)
                    switch 2147483647 {
                    case 1:
                        v5 += 100
                    case 0:
                        if (fi11 >= fi16) {
                            if len(v3) > 0 {
                                v3[0] = -256
                            }
                        } else {
                            var v17 uint = 567
                            if false {
                                v18 := GenS2{F0: v17, F1: 127}
                                _ = v18
                            } else {
                                if false {
                                    switch "hello" {
                                    case "":
                                        var v19 uint32 = 42
                                        _ = v19
                                    case "a":
                                        v20 := GenS3{F0: true}
                                        _ = v20
                                    }
                                } else {
                                    v5 += 42
                                    v4 = (0.1 * (float32(-2.25) * v4))
                                    fmt.Println(fi16)
                                }
                                v21 := GenS1{F0: 1}
                                _ = v21
                            }
                            _ = v17
                        }
                    default:
                        v4 = -2.25
                    }
                    _ = fi16
                }
                var v22 int32 = 1
                _ = v22
            }
            fmt.Println(v5, v2)
            v2 -= (v2 & v2)
            _ = fi11
        }
    }
    if (!(912 == -1)) {
        v5 -= v5
        v5 = genFn0(0, 911)
        for fi23 := 0; fi23 < 8; fi23++ {
            for fi24 := 0; fi24 < 7; fi24++ {
                v1 = (v1 + ("a" + v1))
                var v25 int = (-(fi24))
                _ = fi24
                _ = v25
            }
            _ = len(v3)
            var v26 int32 = -974
            _ = fi23
            _ = v26
        }
    }
    if true {
        v27 := GenS4{F0: 255}
        fmt.Println(v2)
        v2 = (v2 + (v2 >> ((255) & 63)))
        _ = v27
    }
    if (1 > 2) {
        fmt.Println(v1, v5)
        var v28 uint16 = v2
        _ = v28
    } else {
        switch -302 {
        case 0:
            var v29 float64 = (1e10 + 1e10)
            _ = v29
        case 42:
            v5 += (v5 + (v5 * v5))
        case 32767:
            v6 = append(v6, -9223372036854775808)
        }
        switch v1 {
        case "foo bar":
            var v30 bool = (0 < 1)
            _ = v30
        case "":
            v6 = append(v6, 0)
        }
    }
    fmt.Println(v4)
    var v31 float32 = ((v4 / v4) + 1.5)
    if true {
        var v32 int16 = -32768
        var v33 uint16 = (^(v2))
        v5 -= v5
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 uint8 = 255
        _ = fi34
        _ = v35
    }
    if false {
        if true {
            var v36 uint8 = 255
            _ = v36
        }
    } else {
        v37 := GenS5{F0: (v2 - v2), F1: 2147483647}
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
