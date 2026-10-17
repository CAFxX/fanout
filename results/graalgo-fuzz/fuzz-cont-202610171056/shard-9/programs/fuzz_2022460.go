package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int64; F1 int; F2 uint64 }
type GenS6 struct { F0 uint; F1 int8; F2 string }
func genFn0(p0 uint) int16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            p0 = p0
            p0 += 113
            v9 := GenS1{F0: 0, F1: 99}
            fmt.Println(fi7, p0, p0)
            p0 += 4294967295
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS2{F0: p0, F1: -554}
    for fi11 := 0; fi11 < 5; fi11++ {
        switch -126 {
        case 42:
            v12 := []int{fi11, 0}
            _ = v12
        case 447:
            switch "foo bar" {
            case "a":
                var v13 int = 128
                _ = v13
            case "":
                fmt.Println(fi11, fi11)
            }
        case -469:
            p0 = 24
        }
        var v14 bool = (!false)
        _ = fi11
        _ = v14
    }
    switch (("foo bar" + ("\n\t" + "foo bar")) + "foo bar") {
    case "foo bar":
        switch "" {
        case "a":
            v15 := []int{}
            _ = v15
        case "\n\t":
            fmt.Println(p0)
        default:
            v16 := []int{-1, 0, 893}
            _ = v16
        }
    case "\n\t":
        v17 := []int{}
        _ = v17
    default:
        if false {
            var v18 bool = (2 <= 336)
            v18 = (v18 || (!false))
            _ = v18
        }
    }
    _ = p0
    _ = v10
    return 637
}
func genFn1() int8 {
    var v19 int16 = 1000
    v20 := []int{-2147483648, 32767}
    for fi21 := 0; fi21 < 17; fi21++ {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = fi21
    }
    fmt.Println(v19)
    v19 += -1
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 int16 = 42
    var v2 bool = true
    v3 := make(map[uint32]bool)
    var v4 float32 = 1e10
    var v5 int32 = 0
    var v6 uint8 = 102
    switch -9223372036854775808 {
    case -2147483648:
        if (0 >= 42) {
            v2 = v2
        }
    case 1:
        switch ("" + "\n\t") {
        case "foo bar":
            switch 9223372036854775807 {
            case 32768:
                for fi22 := 0; fi22 < 20; fi22++ {
                    v6 = (v6 - 99)
                    for fi23 := 0; fi23 < 7; fi23++ {
                        if (!(((v4 * 0.1) * v4) >= v4)) {
                            v4 -= (((3.14159 / v4) - float32(1.5)) + -2.25)
                        }
                        v24 := GenS3{F0: (!false), F1: 42}
                        v25 := GenS4{F0: genFn0(208), F1: int8(v1)}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = fi22
                }
            }
        case "a":
            if ("a" != "a") {
                v26 := []int{2147483647}
                if v2 {
                    var v27 int32 = (v5 << ((0) & 63))
                    _ = v27
                }
                _ = v26
            }
        }
    case 128:
        var v28 string = (((("foo bar" + (("foo bar" + "hello") + "hello")) + ((("foo bar" + "hello") + ("foo bar" + "foo bar")) + "foo bar")) + ("" + (("a" + ("foo bar" + "hello")) + "foo bar"))) + "hello")
        _ = v28
    }
    v29 := GenS5{F0: 1, F1: -9223372036854775808, F2: 18446744073709551615}
    if (1 > 100) {
        var v30 int = 32767
        if v2 {
            for fi31 := 0; fi31 < 6; fi31++ {
                var v32 bool = ((!(v2 && (!false))) && (v2 || false))
                var v33 float32 = 3.14159
                if v2 {
                    v32 = (v32 && false)
                }
                _ = fi31
                _ = v32
                _ = v33
            }
            fmt.Println(v6, v6, v6)
            var v34 int = -2147483648
            _ = v34
        }
        switch v30 {
        case 911:
            v35 := GenS6{F0: 65535, F1: -128, F2: "hello"}
            _ = v35
        case 32767:
            _ = len(v3)
        default:
            var v36 uint8 = v6
            _ = v36
        }
        _ = v30
    }
    var v37 int8 = 100
    v38 := GenS5{F0: -1, F1: 32768, F2: 0}
    if (!(v5 == 1)) {
        fmt.Println(v6, v4)
        var v39 uint64 = 13
        _ = v39
    } else {
        if (42 > 1) {
            v4 = float32(float64(v4))
            var v40 int16 = v1
            _ = v40
        } else {
            v2 = v2
            for fi41 := 0; fi41 < 5; fi41++ {
                var v42 uint64 = 563
                delete(v3, 129)
                _ = fi41
                _ = v42
            }
        }
        v43 := GenS6{F0: 65535, F1: v37, F2: "foo bar"}
        _ = v43
    }
    v2 = true
    v5 += (-(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
