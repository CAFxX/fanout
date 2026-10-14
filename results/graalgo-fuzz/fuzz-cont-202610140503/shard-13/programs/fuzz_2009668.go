package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int32, p1 int8) int8 {
    v5 := []int{-1, 42}
    switch 9223372036854775807 {
    case 707:
        var v6 uint16 = 2
        _ = v6
    default:
        var v7 uint64 = 42
        _ = v7
    }
    var v8 int64 = -539
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return p1
}

func main() {
    v1 := []string{("\n\t" + "hello"), "", "hello"}
    var v2 float32 = 1e10
    var v3 uint = 0
    v4 := make(map[int]int)
    switch 2147483647 {
    case -9223372036854775808:
        var v9 uint8 = 247
        _ = v9
    default:
        v2 -= v2
    }
    v2 = v2
    v3 -= (^(v3))
    for fi10 := 0; fi10 < 10; fi10++ {
        if (genFn0(int32(v3), 127) < 127) {
            var v11 int64 = 9223372036854775807
            v12 := GenS1{F0: uint16(v3)}
            _ = v11
            _ = v12
        } else {
            var v13 int = (-(fi10))
            fmt.Println(fi10, v13, v3)
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 int8 = 42
                if (!(((v2 + (float32((-2.25 - 0.1)) * 1.5)) + (v2 + 3.14159)) <= v2)) {
                    v16 := GenS1{F0: 0}
                    _ = v16
                }
                v17 := GenS1{F0: 89}
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                var v18 uint64 = 575
                _ = fi14
                _ = v15
                _ = v17
                _ = v18
            }
            _ = v13
        }
        var v19 uint8 = 51
        _ = fi10
        _ = v19
    }
    if (1 <= 140) {
        switch 4294967295 {
        case 127:
            var v20 int16 = 139
            _ = v20
        }
        var v21 int = -128
        _ = v21
    }
    v3 += v3
    if (false && (0 < 127)) {
        var v22 bool = (!true)
        var v23 uint32 = 2147483648
        _ = v22
        _ = v23
    }
    fmt.Println(v3, v3)
    var v24 int8 = 13
    var v25 uint32 = 2147483647
    for fi26 := 0; fi26 < 13; fi26++ {
        v25 = 0
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi27 := 0; fi27 < 19; fi27++ {
            if false {
                if (v3 != (v3 * v3)) {
                    _ = len(v4)
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                } else {
                    var v28 float64 = 0.0
                    if false {
                        var v29 uint64 = 455
                        var v30 string = (("foo bar" + "hello") + "foo bar")
                        delete(v4, (fi27 / ((fi27) | 1)))
                        _ = v29
                        _ = v30
                    }
                    v24 = v24
                    _ = v28
                }
                if false {
                    v25 += v25
                    fmt.Println(v2, v25, v24)
                }
                var v31 uint64 = 511
                _ = v31
            }
            var v32 uint64 = 4294967295
            _ = fi27
            _ = v32
        }
        _ = fi26
    }
    v1 = append(v1, ((("foo bar" + ("a" + "\n\t")) + "") + "foo bar"))
    var v33 int32 = 1
    var v34 bool = false
    var v35 int = 950
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
