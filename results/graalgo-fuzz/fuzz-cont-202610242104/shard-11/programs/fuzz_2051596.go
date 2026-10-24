package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint64 }
func genFn0(p0 uint32, p1 uint8) int8 {
    v5 := GenS1{F0: "a"}
    switch 128 {
    case 42:
        var v6 bool = false
        _ = v6
    }
    switch 128 {
    case 127:
        var v7 int64 = 448
        _ = v7
    }
    v8 := GenS2{F0: false, F1: (p1 >> ((18446744073709551615) & 63))}
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return -59
}

func main() {
    v1 := []int{42}
    v2 := make(map[string]int)
    var v3 int16 = 1
    var v4 uint16 = 65535
    v2["a"] = 128
    var v9 uint8 = 128
    if false {
        if true {
            v10 := v2[(("foo bar" + "") + "a")]
            fmt.Println(v9)
            if (true || (true || false)) {
                v4 -= v4
                var v11 int = (-(v10))
                _ = v11
            }
            _ = v10
        } else {
            fmt.Println(v9, v4, v3)
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS3{F0: 47}
                var v14 uint32 = 285
                v1 = append(v1, (^(fi12)))
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        v15 := GenS2{F0: false, F1: 42}
        var v16 int = 127
        _ = v15
        _ = v16
    } else {
        v3 = v3
    }
    v3 = v3
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 uint8 = v9
        for fi19 := 0; fi19 < 19; fi19++ {
            v18 += v9
            var v20 int = (fi19 ^ fi19)
            _ = fi19
            _ = v20
        }
        v21 := GenS1{F0: "foo bar"}
        _ = fi17
        _ = v18
        _ = v21
    }
    var v22 float64 = 1e10
    v4 = (v4 % ((465) | 1))
    var v23 bool = (!false)
    if true {
        if true {
            for fi24 := 0; fi24 < 20; fi24++ {
                for fi25 := 0; fi25 < 17; fi25++ {
                    v23 = ((0 != (-(v3))) || (-9223372036854775808 == 1))
                    var v26 uint32 = 56
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
            v9 -= v9
            for fi27 := 0; fi27 < 18; fi27++ {
                v28 := GenS3{F0: 4294967295}
                if v23 {
                    for fi29 := 0; fi29 < 13; fi29++ {
                        if (2147483647 == 311) {
                            v30 := GenS4{F0: uint64(v9), F1: 679, F2: 791}
                            _ = v30
                        }
                        if len(v1) > 0 {
                            v1[0] = fi29
                        }
                        _ = fi29
                    }
                    v9 = 255
                }
                _ = fi27
                _ = v28
            }
        } else {
            v31 := GenS1{F0: "foo bar"}
            v3 += (^(v3))
            var v32 string = "a"
            _ = v31
            _ = v32
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
