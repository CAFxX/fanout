package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int) uint {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: false, F1: 205}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: ((float32(float64(0.0)) - (3.14159 * 1.5)) != 3.14159), F1: 0}
    fmt.Println(p0, p1)
    v9 := []int{p1, 722, p1, p1}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 bool = true
    v2 := make(map[int]int)
    var v3 uint64 = 0
    var v4 int16 = -102
    var v5 int8 = 1
    v1 = false
    v10 := GenS2{F0: ""}
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            v3 += v3
            var v13 bool = v1
            _ = fi12
            _ = v13
        }
        v5 = (v5 + v5)
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 int64 = -1
            if (!(v1 || (-2.25 <= (0.0 - -2.25)))) {
                v16 := []int{fi14, fi14, 9223372036854775807}
                v1 = (v3 < 414)
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                if ((4294967295 != (fi14 + fi14)) && true) {
                    v16 = append(v16, fi11)
                } else {
                    v17 := GenS1{F0: v1, F1: 127}
                    fmt.Println(fi14, fi11)
                    v3 = (^(v3))
                    _ = v17
                }
                v15 += v15
                v2[42] = (fi14 & (fi11 - (fi14 * 128)))
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        _ = fi11
    }
    v3 = 1
    if (-1 < 1) {
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := []int{(-(fi18)), fi18, 2147483647}
            var v20 string = "foo bar"
            _ = fi18
            _ = v19
            _ = v20
        }
    } else {
        if (true || ("\n\t" < "")) {
            for fi21 := 0; fi21 < 5; fi21++ {
                v22 := []int{fi21}
                if (1 != 134) {
                    fmt.Println(fi21, v3)
                    v3 -= (^(v3))
                    var v23 uint32 = 1
                    _ = v23
                } else {
                    v3 = v3
                    if (0 >= 0) {
                        v4 = (v4 & 0)
                        v24 := GenS3{F0: v3, F1: 0, F2: (^(v4))}
                        if (v1 || true) {
                            switch "foo bar" {
                            case "foo bar":
                                v25 := GenS2{F0: "hello"}
                                _ = v25
                            case "\n\t":
                                v4 += -1
                            }
                            var v26 uint = 542
                            fmt.Println(fi21)
                            _ = v26
                        } else {
                            fmt.Println(v3, v4, fi21)
                        }
                        _ = v24
                    }
                }
                fmt.Println(v5)
                _ = fi21
                _ = v22
            }
            if (float32(1e10) < 3.14159) {
                for fi27 := 0; fi27 < 9; fi27++ {
                    var v28 float32 = 1.5
                    v29 := GenS3{F0: v3, F1: (v3 + 1), F2: int16(fi27)}
                    if ((32767 == v4) || ((^(v3)) >= 4294967295)) {
                        fmt.Println(v1)
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
        } else {
            v30 := []int{-1}
            _ = v30
        }
    }
    fmt.Println(v3, v1, v4)
    for fi31 := 0; fi31 < 17; fi31++ {
        v1 = (uint8(v3) == 127)
        _ = fi31
    }
    v32 := GenS4{F0: 9223372036854775807, F1: 0}
    if false {
        fmt.Println(v3)
        fmt.Println(v1, v3)
        v33 := []int{32768, 948, -9223372036854775808, 42}
        _ = v33
    }
    v34 := GenS4{F0: 42, F1: 111}
    for fi35 := 0; fi35 < 14; fi35++ {
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 uint32 = 42
            _ = fi36
            _ = v37
        }
        v5 = (v5 >> ((273) & 63))
        v38 := []int{fi35}
        _ = fi35
        _ = v38
    }
    v5 -= (v5 & (v5 - (v5 | (v5 / ((42) | 1)))))
    delete(v2, -9223372036854775808)
    v39 := GenS5{F0: 42}
    for fi40 := 0; fi40 < 11; fi40++ {
        for fi41 := 0; fi41 < 2; fi41++ {
            var v42 float32 = (0.1 * 1e10)
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    var v43 bool = (v1 && true)
    v43 = (false || v43)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v39
    _ = v4
    _ = v43
    _ = v5
}
