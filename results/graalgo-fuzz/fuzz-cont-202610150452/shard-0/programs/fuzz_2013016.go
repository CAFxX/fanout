package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint8; F1 string; F2 string }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint8; F1 int8; F2 int32 }
type GenS5 struct { F0 int; F1 uint64; F2 uint16 }
func genFn0(p0 uint16, p1 string, p2 int16) int8 {
    if false {
        p2 = 1000
        p2 = (p2 >> ((18446744073709551615) & 63))
        v4 := []int{128}
        _ = v4
    }
    v5 := []int{0, 4294967295, -1, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}
func genFn1(p0 string) int64 {
    v6 := []int{32767, 32767}
    p0 = p0
    p0 = (("hello" + ("hello" + p0)) + "\n\t")
    _ = p0
    _ = v6
    return -34
}

func main() {
    var v1 bool = (-519 == -1)
    var v2 string = "hello"
    var v3 int = 392
    if false {
        if (v2 > (v2 + (v2 + v2))) {
            if ((1e10 - 1.5) == 1e10) {
                if (true && (4294967295 > 1)) {
                    v7 := GenS1{F0: 4294967295, F1: 1, F2: 1}
                    _ = v7
                }
                switch ((v2 + v2) + "foo bar") {
                case "foo bar":
                    if v1 {
                        v8 := GenS1{F0: 590, F1: 1, F2: 1}
                        v3 += v3
                        v9 := []int{v3, (-(v3))}
                        _ = v8
                        _ = v9
                    } else {
                        v2 = ("\n\t" + v2)
                    }
                case "":
                    var v10 int16 = 1000
                    _ = v10
                case "hello":
                    var v11 float64 = -2.25
                    _ = v11
                default:
                    v12 := []int{(v3 << ((111) & 63)), (v3 >> ((186) & 63)), 32767}
                    _ = v12
                }
            } else {
                for fi13 := 0; fi13 < 17; fi13++ {
                    v2 = "hello"
                    for fi14 := 0; fi14 < 9; fi14++ {
                        var v15 uint64 = 193
                        if ((90 != 1) && v1) {
                            var v16 int32 = 207
                            var v17 uint64 = (v15 >> ((988) & 63))
                            if true {
                                v18 := []int{fi13}
                                _ = v18
                            } else {
                                switch "foo bar" {
                                case "":
                                    var v19 int32 = (-(v16))
                                    _ = v19
                                case "hello":
                                    var v20 int64 = -9223372036854775808
                                    _ = v20
                                default:
                                    v21 := GenS1{F0: uint32(v17), F1: 0, F2: 1}
                                    _ = v21
                                }
                                v2 = v2
                            }
                            if false {
                                v3 = fi13
                                v22 := GenS1{F0: 4294967295, F1: 128, F2: 124}
                                _ = v22
                            }
                            _ = v16
                            _ = v17
                        } else {
                            v23 := []int{fi13}
                            _ = v23
                        }
                        v1 = false
                        if (!false) {
                            v2 = (v2 + ("a" + v2))
                            var v24 int16 = -1
                            var v25 int = -2147483648
                            _ = v24
                            _ = v25
                        } else {
                            v1 = ((v1 || (v2 > ((v2 + "\n\t") + ("\n\t" + v2)))) || (4294967295 >= 65535))
                            v3 = 42
                            v2 = (((("\n\t" + (v2 + (("\n\t" + v2) + v2))) + v2) + (v2 + ("" + ("a" + "")))) + (("hello" + "a") + v2))
                        }
                        v26 := GenS2{F0: 127, F1: "hello", F2: v2}
                        _ = fi14
                        _ = v15
                        _ = v26
                    }
                    var v27 uint = 18446744073709551615
                    _ = fi13
                    _ = v27
                }
                v28 := GenS3{F0: v3, F1: 0}
                _ = v28
            }
        } else {
            for fi29 := 0; fi29 < 7; fi29++ {
                v30 := []int{v3, (v3 ^ 32767), fi29, 223}
                v2 = v2
                v31 := GenS2{F0: 2, F1: v2, F2: v2}
                _ = fi29
                _ = v30
                _ = v31
            }
            var v32 int64 = 1
            v32 = v32
            _ = v32
        }
        for fi33 := 0; fi33 < 5; fi33++ {
            var v34 float64 = 0.0
            _ = fi33
            _ = v34
        }
    } else {
        var v35 uint32 = 2147483648
        if v1 {
            for fi36 := 0; fi36 < 18; fi36++ {
                switch v2 {
                case "":
                    for fi37 := 0; fi37 < 16; fi37++ {
                        v2 = "a"
                        v3 = fi36
                        _ = fi37
                    }
                }
                v38 := []int{-400, fi36}
                _ = fi36
                _ = v38
            }
            v39 := []int{(v3 * (v3 & (v3 << ((42) & 63))))}
            _ = v39
        } else {
            var v40 uint8 = 127
            _ = v40
        }
        _ = v35
    }
    v3 -= (v3 >> ((0) & 63))
    v1 = (9223372036854775807 <= 0)
    if (-2.25 == float32(0.0)) {
        if v1 {
            v41 := GenS4{F0: 2, F1: -118, F2: 2147483647}
            v42 := []int{}
            var v43 bool = v1
            _ = v41
            _ = v42
            _ = v43
        } else {
            v1 = v1
        }
        v2 = v2
        fmt.Println(v1, v1)
    }
    v44 := []int{(v3 + (v3 ^ -2147483648)), 42, -142}
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    for fi45 := 0; fi45 < 9; fi45++ {
        if ((fi45 + -9223372036854775808) < v3) {
            v1 = (2 <= 0)
            v1 = false
        }
        for fi46 := 0; fi46 < 16; fi46++ {
            v3 += v3
            if ((0.1 - (3.14159 / (3.14159 * (1.5 / (3.14159 / ((0.0 * -2.25) + 0.1)))))) == 0.1) {
                var v47 uint16 = 42
                fmt.Println(fi45, fi45)
                _ = v47
            }
            _ = fi46
        }
        _ = fi45
    }
    _ = len(v44)
    v48 := GenS5{F0: (v3 * v3), F1: 664, F2: 42}
    v3 += 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v48
}
