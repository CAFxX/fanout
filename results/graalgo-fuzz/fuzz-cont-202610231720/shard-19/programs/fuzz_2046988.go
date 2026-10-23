package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 string }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint32; F1 bool }
func genFn0() string {
    var v4 int32 = 1
    var v5 uint8 = 42
    v6 := []int{-1, 4294967295}
    var v7 int16 = -1
    fmt.Println(v5, v4, v4)
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return ("foo bar" + "\n\t")
}

func main() {
    var v1 float64 = -2.25
    var v2 int64 = 0
    v3 := make(map[int]uint32)
    v1 = v1
    v8 := GenS1{F0: -32768, F1: "hello", F2: "\n\t"}
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(v1)
        _ = fi9
    }
    var v10 float64 = v1
    if (232 != 4294967295) {
        v10 = v10
        v2 -= (v2 & (v2 % (((v2 & (v2 | -9223372036854775808))) | 1)))
        var v11 int = 9223372036854775807
        _ = v11
    } else {
        if (255 == 127) {
            fmt.Println(v2, v2, v10)
            var v12 uint32 = 2147483648
            _ = v12
        } else {
            for fi13 := 0; fi13 < 2; fi13++ {
                switch ((("foo bar" + ("foo bar" + "a")) + "\n\t") + ("\n\t" + "hello")) {
                case "a":
                    var v14 int32 = 0
                    _ = v14
                case "\n\t":
                    v15 := GenS2{F0: (128 > 180), F1: 127}
                    _ = v15
                case "hello":
                    v1 -= ((v10 * (v1 - (((v1 * v10) * v1) + v10))) / (v1 / 1.0))
                }
                var v16 uint8 = 43
                _ = fi13
                _ = v16
            }
            v1 = (1.5 * (v1 * v1))
        }
    }
    var v17 uint = 255
    for fi18 := 0; fi18 < 20; fi18++ {
        switch 4294967295 {
        case -9223372036854775808:
            fmt.Println(v2, v10)
        case 9223372036854775807:
            v19 := GenS2{F0: (fi18 == fi18), F1: 255}
            _ = v19
        case 127:
            switch ((("a" + "foo bar") + (("foo bar" + "foo bar") + ("foo bar" + "\n\t"))) + "\n\t") {
            case "":
                for fi20 := 0; fi20 < 4; fi20++ {
                    v10 = v10
                    _ = len(v3)
                    v2 = (v2 >> ((42) & 63))
                    _ = fi20
                }
            }
        }
        _ = fi18
    }
    var v21 int = 9223372036854775807
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 int16 = 1
            _ = fi23
            _ = v24
        }
        if false {
            v25 := []int{fi22, fi22}
            if (65535 <= 42) {
                fmt.Println(v21)
                fmt.Println(v10, fi22)
            } else {
                v26 := GenS1{F0: -430, F1: "\n\t", F2: "a"}
                var v27 int64 = v2
                _ = v26
                _ = v27
            }
            _ = v25
        } else {
            v3[v21] = 568
            var v28 uint64 = 42
            _ = v28
        }
        v10 += (v1 + (v1 + ((float64(1e10) - ((((float64(0.1) * (1.5 * v1)) / (v1 - float64(0.0))) + (-2.25 + v1)) / v1)) - 3.14159)))
        _ = fi22
    }
    v17 += v17
    for fi29 := 0; fi29 < 17; fi29++ {
        for fi30 := 0; fi30 < 5; fi30++ {
            fmt.Println(fi29, fi29)
            if (true && true) {
                v17 -= v17
                v10 = ((1.5 / v10) / v10)
                _ = len(v3)
            } else {
                fmt.Println(v10, fi30)
            }
            if (42 != 0) {
                delete(v3, (v21 << (((^(v17))) & 63)))
                v2 -= -1
            } else {
                var v31 bool = false
                _ = v31
            }
            _ = fi30
        }
        for fi32 := 0; fi32 < 9; fi32++ {
            if ((v17 + (v17 % ((v17) | 1))) != 18446744073709551615) {
                v33 := []int{-1, (fi32 / (((v21 >> ((v17) & 63))) | 1)), int(v2)}
                v34 := GenS3{F0: 42, F1: 255, F2: 42}
                _ = v33
                _ = v34
            } else {
                if (-37 >= 0) {
                    if true {
                        v35 := GenS2{F0: (!false), F1: 127}
                        _ = v35
                    } else {
                        v36 := GenS3{F0: v2, F1: 0, F2: uint64(v17)}
                        v10 = 1e10
                        _ = v36
                    }
                    v37 := []int{-574, (-(v21)), 128, (-(fi32))}
                    v38 := GenS2{F0: true, F1: 42}
                    _ = v37
                    _ = v38
                }
            }
            v39 := []int{fi29, (v21 & (-(v21))), fi32}
            _ = fi32
            _ = v39
        }
        _ = fi29
    }
    v40 := GenS4{F0: 42, F1: true}
    v41 := GenS1{F0: 790, F1: "", F2: (((((("hello" + "\n\t") + ("\n\t" + ("" + "a"))) + ("a" + ("\n\t" + ("\n\t" + ("hello" + "a"))))) + ("foo bar" + "hello")) + "") + "a")}
    if (v2 > (v2 * v2)) {
        fmt.Println(v2)
        var v42 int = (-(v21))
        _ = v42
    } else {
        var v43 string = "hello"
        if (!(557 > (-(v2)))) {
            switch (v21 + 1) {
            case 127:
                if (1.5 != 0.0) {
                    for fi44 := 0; fi44 < 3; fi44++ {
                        v45 := []int{fi44}
                        _ = fi44
                        _ = v45
                    }
                }
            }
            v43 = ""
            var v46 int8 = -128
            _ = v46
        }
        v47 := GenS3{F0: v2, F1: 65535, F2: 4294967295}
        _ = v43
        _ = v47
    }
    var v48 int16 = 1000
    var v49 int = v21
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v40
    _ = v41
    _ = v48
    _ = v49
    _ = v8
}
