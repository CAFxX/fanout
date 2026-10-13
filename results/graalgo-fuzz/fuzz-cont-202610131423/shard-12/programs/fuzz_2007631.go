package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint) int16 {
    var v5 uint32 = 0
    p0 = p0
    var v6 uint = (^(p0))
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{fi7}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    return 1000
}
func genFn1(p0 int) int16 {
    switch p0 {
    case 4294967295:
        v9 := []int{p0}
        _ = v9
    case 635:
        v10 := []int{192, -803, p0, (p0 / ((-2147483648) | 1))}
        _ = v10
    }
    if true {
        fmt.Println(p0)
    } else {
        var v11 uint = 18446744073709551615
        _ = v11
    }
    var v12 float32 = (((float32(0.1) + 1.5) + 3.14159) - 1.5)
    _ = p0
    _ = v12
    return genFn0(386)
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    fmt.Println(v1)
    fmt.Println(v1)
    fmt.Println(v1)
    for fi13 := 0; fi13 < 3; fi13++ {
        fmt.Println(v1)
        if (false && false) {
            if (!true) {
                var v14 uint32 = 1
                v15 := []int{fi13, (fi13 * (fi13 ^ (fi13 * fi13))), fi13, fi13}
                _ = v14
                _ = v15
            } else {
                var v16 int8 = -90
                _ = v16
            }
        }
        _ = fi13
    }
    if (42 != -9223372036854775808) {
        var v17 float32 = 0.0
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 string = ""
            _ = fi18
            _ = v19
        }
        _ = v17
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            v1 = 26
            _ = fi20
        }
        if (false || false) {
            for fi21 := 0; fi21 < 15; fi21++ {
                var v22 int64 = 42
                v2[(fi21 / ((128) | 1))] = 32767
                _ = fi21
                _ = v22
            }
            v1 += (v1 | (v1 % ((v1) | 1)))
        }
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        v1 += (^(v1))
        fmt.Println(v1, v1)
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 uint64 = 4294967295
            v25 = v25
            v26 := GenS1{F0: ("" + "foo bar")}
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi23
    }
    var v27 int32 = -1
    v28 := GenS1{F0: "a"}
    if (((-2.25 + (0.0 + (0.1 + (3.14159 - -2.25)))) - (((0.1 - 1e10) / 1.0) * 0.1)) >= (0.1 - (3.14159 + (0.0 + (0.1 / float32(-2.25)))))) {
        var v29 int8 = -128
        if ((3.14159 / 1e10) == 0.1) {
            v1 = (v1 * v1)
            if false {
                delete(v4, "hello")
                for fi30 := 0; fi30 < 8; fi30++ {
                    var v31 uint64 = 42
                    v27 += (v27 >> ((18446744073709551615) & 63))
                    if true {
                        var v32 uint64 = uint64(v1)
                        var v33 float64 = (-2.25 / 3.14159)
                        _ = v32
                        _ = v33
                    } else {
                        for fi34 := 0; fi34 < 2; fi34++ {
                            fmt.Println(v1, fi30)
                            var v35 uint = 18446744073709551615
                            _ = fi34
                            _ = v35
                        }
                        v36 := GenS1{F0: ""}
                        _ = v36
                    }
                    _ = fi30
                    _ = v31
                }
            }
        } else {
            v1 = 127
        }
        _ = v29
    } else {
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 float32 = 1.5
            if false {
                var v39 int8 = 42
                v40 := GenS2{F0: 2}
                _ = v39
                _ = v40
            } else {
                var v41 int32 = (^(v27))
                var v42 float64 = 1e10
                fmt.Println(v1, v41, v27)
                _ = v41
                _ = v42
            }
            var v43 float64 = (0.0 + 0.1)
            _ = fi37
            _ = v38
            _ = v43
        }
    }
    fmt.Println(v1, v1)
    v44 := GenS1{F0: "hello"}
    v1 += (v1 | 64)
    v27 = (v27 + (v27 % ((2147483647) | 1)))
    for fi45 := 0; fi45 < 4; fi45++ {
        for fi46 := 0; fi46 < 10; fi46++ {
            var v47 int64 = 1
            _ = fi46
            _ = v47
        }
        if true {
            if ((528 <= 0) || ((!(-2.25 == 1.5)) && (v27 == v27))) {
                var v48 int8 = 42
                var v49 float32 = float32(0.0)
                v49 -= float32((3.14159 * 0.1))
                _ = v48
                _ = v49
            }
        }
        var v50 string = "\n\t"
        _ = fi45
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v44
}
