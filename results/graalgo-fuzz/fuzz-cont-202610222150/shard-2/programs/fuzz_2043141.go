package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 uint64; F1 string; F2 uint32 }
func genFn0() uint32 {
    if false {
        var v7 int = -1
        _ = v7
    }
    var v8 uint32 = 2
    v8 += (v8 ^ (^(v8)))
    fmt.Println(v8)
    fmt.Println(v8)
    _ = v8
    return (v8 + 985)
}

func main() {
    var v1 float64 = (-2.25 - 1.5)
    var v2 uint32 = 2147483648
    var v3 int = 128
    var v4 bool = (-1 < 42)
    var v5 string = ("foo bar" + ("foo bar" + ""))
    var v6 float64 = (v1 - 1.5)
    var v9 uint8 = 127
    v10 := []int{v3, v3, (v3 * v3), (-(v3))}
    for fi11 := 0; fi11 < 1; fi11++ {
        if false {
            var v12 uint64 = 4294967295
            if (-2147483648 <= 0) {
                fmt.Println(v12)
                switch (v3 & (fi11 | (v3 - fi11))) {
                case 844:
                    if ((v9 & (^(v9))) <= 126) {
                        v4 = v4
                        var v13 float64 = 3.14159
                        _ = v13
                    } else {
                        var v14 int8 = 65
                        if false {
                            v12 -= (v12 << ((2) & 63))
                            fmt.Println(v5, v1)
                            v1 -= float64(0.0)
                        }
                        _ = v14
                    }
                }
            } else {
                var v15 uint32 = (v2 % ((1) | 1))
                _ = v15
            }
            _ = v12
        } else {
            v16 := GenS1{F0: 0, F1: uint8(v2), F2: 4294967295}
            if (!(v4 && false)) {
                switch ((v5 + (v5 + v5)) + "foo bar") {
                case "":
                    v2 = v2
                }
                var v17 int8 = -117
                var v18 uint64 = 0
                _ = v17
                _ = v18
            } else {
                var v19 uint64 = 4294967295
                var v20 int8 = -1
                _ = v19
                _ = v20
            }
            _ = v16
        }
        var v21 string = ("a" + v5)
        if false {
            for fi22 := 0; fi22 < 18; fi22++ {
                if (1e10 >= 3.14159) {
                    if v4 {
                        if v4 {
                            v9 -= v9
                            var v23 uint32 = 1
                            _ = v23
                        }
                        v24 := GenS1{F0: 9223372036854775807, F1: (v9 | 255), F2: 42}
                        _ = v24
                    } else {
                        if (v2 >= v2) {
                            var v25 uint16 = 42
                            var v26 bool = ((v5 + v21) == v5)
                            var v27 int8 = 1
                            _ = v25
                            _ = v26
                            _ = v27
                        } else {
                            v28 := GenS1{F0: 42, F1: v9, F2: 42}
                            v29 := GenS1{F0: 9223372036854775807, F1: v9, F2: 4294967295}
                            _ = v28
                            _ = v29
                        }
                        if (!false) {
                            v30 := GenS2{F0: 47, F1: 18446744073709551615, F2: v4}
                            v5 = "foo bar"
                            _ = v30
                        } else {
                            if v4 {
                                v9 = v9
                            }
                            v31 := GenS3{F0: (fi22 - (fi22 - 127))}
                            v6 = float64(3.14159)
                            _ = v31
                        }
                        v3 += fi22
                    }
                } else {
                    _ = len(v10)
                }
                v32 := GenS2{F0: (v2 | genFn0()), F1: 0, F2: v4}
                _ = fi22
                _ = v32
            }
            for fi33 := 0; fi33 < 14; fi33++ {
                if v4 {
                    if false {
                        var v34 float32 = (0.1 + 0.0)
                        v34 = ((v34 / 3.14159) + ((v34 * 0.1) * 1.5))
                        if v4 {
                            v35 := GenS2{F0: v2, F1: 18446744073709551615, F2: ((0 == 42) && (v4 && true))}
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            var v36 int16 = -32768
                            _ = v35
                            _ = v36
                        }
                        _ = v34
                    }
                    v37 := GenS4{F0: 0}
                    _ = v37
                } else {
                    fmt.Println(v6, v1, v6)
                    v38 := GenS2{F0: genFn0(), F1: 36, F2: (0 < 9223372036854775807)}
                    var v39 uint32 = (v2 & 183)
                    var v40 uint8 = v9
                    _ = v38
                    _ = v39
                    _ = v40
                }
                v41 := GenS1{F0: 42, F1: (v9 | 0), F2: 4294967295}
                var v42 float64 = 1.5
                _ = fi33
                _ = v41
                _ = v42
            }
            v21 = (v5 + v21)
        } else {
            fmt.Println(v6, v3, v6)
            var v43 int16 = -99
            v4 = v4
            _ = v43
        }
        _ = fi11
        _ = v21
    }
    if len(v10) > 0 {
        v10[0] = (v3 / ((4294967295) | 1))
    }
    v44 := GenS5{F0: (v9 >> ((uint(v9)) & 63)), F1: (^(v9)), F2: 1}
    v4 = false
    var v45 bool = (((-2.25 - 1e10) - v1) >= ((float64(3.14159) - 0.0) + v6))
    if (!v4) {
        v3 = v3
        v9 += 0
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        v47 := GenS2{F0: (^(v2)), F1: 42, F2: v45}
        for fi48 := 0; fi48 < 16; fi48++ {
            if ((v2 ^ 2) <= v2) {
                v49 := GenS2{F0: v2, F1: uint(v9), F2: true}
                v9 = v9
                switch -806 {
                case 32768:
                    var v50 int32 = int32(v9)
                    _ = v50
                }
                var v51 uint64 = 4294967295
                fmt.Println(v4)
                fmt.Println(fi48, v1, v1)
                _ = v49
                _ = v51
            } else {
                if v45 {
                    var v52 int16 = 42
                    v53 := GenS5{F0: v9, F1: 122, F2: 65535}
                    v54 := GenS6{F0: 0, F1: 1}
                    _ = v52
                    _ = v53
                    _ = v54
                } else {
                    v5 = ""
                    v55 := GenS1{F0: -9223372036854775808, F1: v9, F2: 2}
                    v56 := GenS7{F0: 2, F1: "hello", F2: (v2 >> ((66) & 63))}
                    _ = v55
                    _ = v56
                }
                var v57 uint64 = 1
                var v58 uint16 = 42
                _ = v57
                _ = v58
            }
            var v59 string = v5
            var v60 int16 = 32767
            fmt.Println(v6, v1)
            _ = fi48
            _ = v59
            _ = v60
        }
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v9
}
