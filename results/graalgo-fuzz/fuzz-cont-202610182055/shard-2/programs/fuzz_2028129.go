package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint8, p1 uint, p2 uint16) int {
    v4 := []int{32768}
    v4 = append(v4, 9223372036854775807)
    var v5 int16 = -1
    switch "a" {
    case "\n\t":
        if (!false) {
            var v6 bool = true
            _ = v6
        }
    default:
        p0 = (p0 + 1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint64 = 484
    var v2 bool = (false && true)
    v3 := make(map[int]bool)
    fmt.Println(v2)
    fmt.Println(v1, v1)
    if v2 {
        v7 := []int{0, 32767}
        _ = v7
    }
    var v8 int16 = -32768
    if (float32((0.0 / 1.5)) != (float32(((0.1 * -2.25) * (-2.25 * 0.0))) + 1.5)) {
        if (true || (v2 || v2)) {
            var v9 float64 = 0.0
            _ = v9
        }
        v2 = false
    }
    v10 := []int{}
    v1 = (v1 + (^(v1)))
    v1 -= 2
    if true {
        for fi11 := 0; fi11 < 2; fi11++ {
            if false {
                for fi12 := 0; fi12 < 14; fi12++ {
                    var v13 int16 = v8
                    _ = fi12
                    _ = v13
                }
                switch ("foo bar" + "foo bar") {
                case "a":
                    fmt.Println(fi11, v2)
                case "foo bar":
                    if ((fi11 & genFn0(25, 255, 0)) >= (fi11 - genFn0(1, 55, 520))) {
                        var v14 uint32 = 42
                        _ = v14
                    } else {
                        if false {
                            v2 = ((v1 << ((1) & 63)) > 20)
                        } else {
                            var v15 int16 = v8
                            fmt.Println(v15, v8, v15)
                            v1 = 4294967295
                            _ = v15
                        }
                    }
                }
            } else {
                v2 = (2 < 0)
                v2 = v2
                switch "\n\t" {
                case "foo bar":
                    v2 = v2
                default:
                    for fi16 := 0; fi16 < 14; fi16++ {
                        var v17 uint32 = 4294967295
                        delete(v3, fi11)
                        if v2 {
                            v1 -= 749
                            if (!v2) {
                                v2 = ((^(v17)) <= v17)
                            }
                            v2 = (719 != v1)
                            var v18 uint8 = 127
                            _ = v18
                        } else {
                            v19 := GenS1{F0: -9223372036854775808}
                            fmt.Println(v8, fi11, fi16)
                            if (!(!(v2 || true))) {
                                v17 = v17
                                v2 = false
                            } else {
                                var v20 uint8 = 55
                                _ = v20
                            }
                            v21 := GenS1{F0: 1}
                            var v22 int64 = -1
                            _ = v19
                            _ = v21
                            _ = v22
                        }
                        fmt.Println(v2, v17)
                        _ = fi16
                        _ = v17
                    }
                }
            }
            fmt.Println(v1, v2)
            _ = fi11
        }
        v23 := GenS1{F0: -234}
        v24 := GenS1{F0: -9223372036854775808}
        _ = v23
        _ = v24
    }
    delete(v3, -1)
    fmt.Println(v2, v2, v2)
    v25 := GenS1{F0: 9223372036854775807}
    v26 := v3[-2147483648]
    var v27 uint64 = 472
    switch "hello" {
    case "":
        switch "a" {
        case "foo bar":
            v2 = (-1 == -738)
        }
    case "foo bar":
        var v28 uint16 = 104
        _ = v28
    }
    switch -2147483648 {
    case 9223372036854775807:
        if false {
            if (42 <= 0) {
                if (65535 == 42) {
                    v1 += 42
                    var v29 float64 = 1e10
                    _ = v29
                }
            }
            for fi30 := 0; fi30 < 2; fi30++ {
                if true {
                    for fi31 := 0; fi31 < 14; fi31++ {
                        fmt.Println(v26)
                        var v32 int64 = -1
                        _ = fi31
                        _ = v32
                    }
                    fmt.Println(v27, v1, v8)
                }
                var v33 int8 = -51
                _ = fi30
                _ = v33
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v8
}
