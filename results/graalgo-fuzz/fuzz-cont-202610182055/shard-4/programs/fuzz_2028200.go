package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 bool; F2 bool }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint16) string {
    v4 := GenS1{F0: 42}
    var v5 int8 = 100
    if (2 != 123) {
        v6 := GenS2{F0: 0}
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                var v9 uint = 0
                v10 := []int{(fi8 ^ fi8), (fi8 - fi7), 632}
                _ = fi8
                _ = v10
                _ = v9
            }
            v5 -= (v5 / ((42) | 1))
            fmt.Println(v5, fi7, v5)
            _ = fi7
        }
        var v11 uint32 = 2147483647
        _ = v11
        _ = v6
    }
    _ = p0
    _ = v4
    _ = v5
    return "foo bar"
}
func genFn1() bool {
    v12 := []int{9223372036854775807, -682, 4294967295, 152}
    _ = len(v12)
    var v13 float32 = 0.0
    v13 -= (-2.25 + v13)
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 bool = true
        _ = fi14
        _ = v15
    }
    _ = v12
    _ = v13
    return true
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 int16 = 1
    v3 = (-(v3))
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int = (fi16 - (fi16 / ((128) | 1)))
        v18 := GenS2{F0: 42}
        v19 := []int{}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    v3 = (^(v3))
    for fi20 := 0; fi20 < 18; fi20++ {
        v3 -= 888
        _ = fi20
    }
    fmt.Println(v3)
    var v21 int = -1
    if ((2 < 968) && true) {
        if (-9223372036854775808 > -9223372036854775808) {
            var v22 float32 = float32((-2.25 + 0.1))
            for fi23 := 0; fi23 < 14; fi23++ {
                v22 -= v22
                _ = fi23
            }
            _ = v22
        } else {
            fmt.Println(v21)
        }
    } else {
        v2[4294967295] = ""
        if (18446744073709551615 <= 42) {
            switch (("\n\t" + ((("hello" + "") + "") + (("" + ("hello" + "hello")) + "hello"))) + "hello") {
            case "a":
                var v24 int16 = (v3 - 361)
                _ = v24
            }
            fmt.Println(v21, v3)
            if false {
                var v25 float64 = -2.25
                _ = v25
            }
        }
        v21 -= v21
    }
    if (255 > 576) {
        var v26 int32 = -19
        switch "a" {
        case "foo bar":
            if ((255 > 18446744073709551615) && false) {
                v27 := GenS3{F0: 0, F1: false, F2: true}
                if false {
                    v21 = (v21 ^ v21)
                    v28 := GenS4{F0: 2147483647, F1: 1, F2: 255}
                    _ = v28
                } else {
                    v26 += (v26 % (((-(v26))) | 1))
                }
                _ = v27
            }
        case "":
            v29 := v2[4294967295]
            _ = v29
        case "\n\t":
            var v30 int8 = -1
            _ = v30
        default:
            v31 := []int{}
            _ = v31
        }
        _ = v26
    } else {
        v21 += (v21 % ((-9223372036854775808) | 1))
        var v32 int32 = 2147483647
        var v33 float64 = -2.25
        _ = v32
        _ = v33
    }
    v3 = (v3 / ((32767) | 1))
    fmt.Println(v3)
    for fi34 := 0; fi34 < 15; fi34++ {
        var v35 int = (v21 % ((698) | 1))
        for fi36 := 0; fi36 < 5; fi36++ {
            if (true || true) {
                fmt.Println(v35, fi34, fi36)
            }
            _ = fi36
        }
        v35 += v21
        _ = fi34
        _ = v35
    }
    if (true && (0.0 > 1e10)) {
        for fi37 := 0; fi37 < 1; fi37++ {
            if (false || true) {
                var v38 bool = true
                var v39 uint = 42
                v21 += fi37
                _ = v38
                _ = v39
            }
            _ = fi37
        }
        if true {
            var v40 uint32 = 1
            var v41 int8 = 1
            v42 := GenS5{F0: 819}
            _ = v40
            _ = v41
            _ = v42
        } else {
            var v43 float64 = 0.0
            _ = v43
        }
        var v44 uint8 = 255
        _ = v44
    } else {
        fmt.Println(v21, v3)
        for fi45 := 0; fi45 < 16; fi45++ {
            for fi46 := 0; fi46 < 7; fi46++ {
                v2[2] = ("hello" + "hello")
                fmt.Println(v21, v3, fi46)
                _ = fi46
            }
            var v47 float32 = 0.1
            _ = fi45
            _ = v47
        }
        fmt.Println(v3, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
}
