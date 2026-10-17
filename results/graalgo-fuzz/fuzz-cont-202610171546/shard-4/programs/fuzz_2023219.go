package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() bool {
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 105}
        if true {
            var v9 uint = 4294967295
            v9 = v9
            if ((1 >= 42) && false) {
                var v10 uint16 = 493
                _ = v10
            }
            _ = v9
        }
        var v11 uint = 0
        _ = fi7
        _ = v11
        _ = v8
    }
    if true {
        if (2 <= 1) {
            for fi12 := 0; fi12 < 6; fi12++ {
                switch 32768 {
                case 556:
                    v13 := []int{}
                    _ = v13
                default:
                    v14 := GenS1{F0: 42}
                    _ = v14
                }
                if true {
                    fmt.Println(fi12)
                }
                _ = fi12
            }
            var v15 float64 = 3.14159
            _ = v15
        } else {
            var v16 int = 42
            if true {
                v17 := []int{(v16 / ((v16) | 1)), 2147483647, 330, (v16 + 1)}
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 int32 = 42
                    v17 = append(v17, fi18)
                    var v20 float64 = 1.5
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = v17
            } else {
                for fi21 := 0; fi21 < 2; fi21++ {
                    v22 := []int{v16, fi21}
                    if true {
                        v16 += fi21
                        v16 -= (fi21 - v16)
                        v16 += -9223372036854775808
                    }
                    _ = fi21
                    _ = v22
                }
                v23 := []int{4294967295, v16, (v16 | (v16 ^ v16))}
                _ = v23
            }
            if (0 < 65535) {
                if true {
                    v16 = v16
                    v24 := []int{808, 1, -1}
                    _ = v24
                } else {
                    if (!(true && ("a" != (((("a" + "") + "") + (("foo bar" + "foo bar") + "hello")) + "")))) {
                        var v25 int32 = -1
                        v25 += v25
                        _ = v25
                    } else {
                        v16 -= 371
                        fmt.Println(v16)
                    }
                    var v26 int8 = 42
                    _ = v26
                }
            } else {
                v16 += v16
            }
            _ = v16
        }
    }
    return false
}

func main() {
    var v1 bool = (-2.25 >= -2.25)
    var v2 uint16 = 255
    var v3 uint = 255
    var v4 int8 = 100
    var v5 uint = v3
    var v6 string = "hello"
    v27 := []int{32768, 127}
    var v28 int8 = (-(v4))
    switch (v6 + "foo bar") {
    case "a":
        v1 = v1
    case "foo bar":
        if false {
            fmt.Println(v5)
        } else {
            var v29 float64 = ((-2.25 + 0.1) * (0.1 + 0.0))
            fmt.Println(v28, v1, v5)
            v2 -= (^(v2))
            _ = v29
        }
    default:
        var v30 int32 = 211
        _ = v30
    }
    v3 += (v5 - v5)
    v3 += 18446744073709551615
    v31 := GenS1{F0: 752}
    var v32 int16 = -849
    switch -9223372036854775808 {
    case 32768:
        v2 -= (v2 / (((v2 * (v2 << ((1) & 63)))) | 1))
    case -2147483648:
        v5 = (^(v5))
    case 2147483647:
        v6 = v6
    default:
        var v33 int32 = 42
        _ = v33
    }
    v3 = 65535
    var v34 string = ((v6 + v6) + "hello")
    for fi35 := 0; fi35 < 9; fi35++ {
        if false {
            var v36 int64 = -1
            _ = v36
        }
        _ = fi35
    }
    if len(v27) > 0 {
        v27[0] = 9223372036854775807
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
