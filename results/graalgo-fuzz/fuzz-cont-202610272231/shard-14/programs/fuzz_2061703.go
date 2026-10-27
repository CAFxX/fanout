package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint64; F2 string }
type GenS4 struct { F0 int; F1 uint }
func genFn0(p0 int32) int8 {
    v6 := GenS1{F0: -379}
    v7 := GenS1{F0: -525}
    var v8 string = ("hello" + ("" + ("" + "a")))
    fmt.Println(v8)
    var v9 bool = (128 > 2)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 74
}
func genFn1(p0 uint8) int32 {
    p0 = (p0 - p0)
    v10 := GenS1{F0: -803}
    if (65535 != 4294967295) {
        if false {
            switch 4294967295 {
            case 32768:
                for fi11 := 0; fi11 < 15; fi11++ {
                    for fi12 := 0; fi12 < 12; fi12++ {
                        fmt.Println(fi11, p0)
                        _ = fi12
                    }
                    var v13 int16 = -32768
                    v13 += v13
                    _ = fi11
                    _ = v13
                }
            case -857:
                fmt.Println(p0)
            case -1:
                for fi14 := 0; fi14 < 11; fi14++ {
                    fmt.Println(p0, fi14)
                    if false {
                        p0 += p0
                    }
                    _ = fi14
                }
            }
            var v15 float64 = 1e10
            if false {
                v16 := GenS2{F0: 1}
                if true {
                    v17 := []int{-1, 32768, -2147483648}
                    _ = v17
                } else {
                    var v18 uint8 = 0
                    _ = v18
                }
                _ = v16
            } else {
                var v19 uint8 = p0
                _ = v19
            }
            _ = v15
        }
    } else {
        fmt.Println(p0)
        fmt.Println(p0)
    }
    p0 -= (p0 * 42)
    v20 := []int{}
    _ = p0
    _ = v10
    _ = v20
    return 0
}

func main() {
    var v1 int32 = -604
    var v2 uint16 = 0
    var v3 int8 = -111
    var v4 int8 = (v3 * (v3 + (v3 | (v3 - (v3 / (((v3 + v3)) | 1))))))
    var v5 uint64 = 4294967295
    var v21 int32 = 1
    fmt.Println(v21)
    fmt.Println(v21)
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := []int{(fi22 >> ((723) & 63)), fi22}
        _ = fi22
        _ = v23
    }
    var v24 float32 = 3.14159
    v24 += v24
    for fi25 := 0; fi25 < 9; fi25++ {
        if ((9223372036854775807 <= 32768) || true) {
            v5 = 42
            v2 = (v2 + v2)
        }
        var v26 string = "a"
        _ = fi25
        _ = v26
    }
    v27 := GenS3{F0: "\n\t", F1: 671, F2: (("foo bar" + ((("" + "foo bar") + "a") + "foo bar")) + ("hello" + ("foo bar" + ("a" + ((("foo bar" + "a") + (("hello" + ("hello" + "")) + "hello")) + (("a" + ("" + "")) + "\n\t"))))))}
    for fi28 := 0; fi28 < 10; fi28++ {
        v29 := GenS4{F0: (fi28 % ((-2147483648) | 1)), F1: 319}
        var v30 int64 = 9223372036854775807
        for fi31 := 0; fi31 < 6; fi31++ {
            fmt.Println(fi28, v2)
            v32 := GenS1{F0: -1}
            v30 -= v30
            v30 -= (v30 % (((v30 | (v30 % (((v30 >> ((2) & 63))) | 1)))) | 1))
            fmt.Println(fi31)
            _ = fi31
            _ = v32
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        if true {
            var v34 int32 = (v1 - genFn1(4))
            _ = v34
        }
        _ = fi33
    }
    var v35 string = "foo bar"
    if (-9223372036854775808 == 667) {
        if (false || (956 < 2)) {
            var v36 float64 = (3.14159 + 1.5)
            fmt.Println(v36, v21, v1)
            _ = v36
        }
    }
    v4 = (v3 & v4)
    v37 := GenS2{F0: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v27
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
