package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 string }
func genFn0() string {
    var v5 int32 = 997
    v6 := []int{-904, -9223372036854775808, -19, -9223372036854775808}
    if true {
        v7 := GenS1{F0: 2, F1: (32767 == 1000), F2: ("hello" + "a")}
        _ = v7
    } else {
        v5 -= v5
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v5 = v5
    _ = v5
    _ = v6
    return "a"
}
func genFn1(p0 int32) bool {
    var v8 int32 = 1
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 string = "a"
        var v11 int = (fi9 & (fi9 + fi9))
        for fi12 := 0; fi12 < 18; fi12++ {
            if (-1 == 0) {
                p0 -= (-(p0))
            } else {
                v13 := []int{}
                _ = v13
            }
            p0 += v8
            _ = fi12
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v8
    return true
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 int32 = 42
    var v4 uint64 = 158
    v14 := v1[517]
    v15 := GenS1{F0: 42, F1: true, F2: "a"}
    if true {
        if (!false) {
            fmt.Println(v14)
            var v16 uint = 4294967295
            _ = v16
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            v4 = 0
            var v18 uint32 = 2147483647
            var v19 float64 = float64(float32((-2.25 + 1e10)))
            _ = fi17
            _ = v18
            _ = v19
        }
        if (-9223372036854775808 <= 127) {
            v4 = v4
            for fi20 := 0; fi20 < 12; fi20++ {
                v4 = v4
                if v14 {
                    var v21 float32 = 0.1
                    for fi22 := 0; fi22 < 16; fi22++ {
                        var v23 uint = 255
                        var v24 int16 = 32767
                        var v25 uint64 = v4
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    var v26 bool = ((v14 || (v14 && v14)) && false)
                    _ = v21
                    _ = v26
                } else {
                    v27 := GenS1{F0: 42, F1: ((-2.25 - 0.0) > ((0.0 - (1e10 * ((0.0 * 0.1) * float64(0.0)))) + (float64((0.1 - 1.5)) * float64(float32(-2.25))))), F2: "hello"}
                    _ = v27
                }
                _ = fi20
            }
        } else {
            v28 := GenS1{F0: 42, F1: v14, F2: "hello"}
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 int = int(v3)
        v4 = (v4 ^ v4)
        v4 -= (v4 | v4)
        _ = fi29
        _ = v30
    }
    v3 += v3
    if v14 {
        if false {
            var v31 int64 = 1
            v1[231] = (42 == -1)
            _ = v31
        } else {
            fmt.Println(v3, v14)
        }
    } else {
        if (520 > 2147483648) {
            var v32 uint64 = 927
            v14 = false
            _ = v32
        }
        var v33 int32 = 2147483647
        _ = v33
    }
    v34 := []int{9223372036854775807}
    v14 = (true || false)
    v4 -= v4
    delete(v2, ("foo bar" + (("\n\t" + "") + ("" + ("" + "\n\t")))))
    var v35 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
