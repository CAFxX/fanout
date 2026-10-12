package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint; F1 int16; F2 int }
func genFn0(p0 uint64, p1 bool) uint8 {
    var v5 float64 = (0.1 - 0.1)
    for fi6 := 0; fi6 < 1; fi6++ {
        p1 = (!p1)
        var v7 uint64 = 157
        _ = fi6
        _ = v7
    }
    if (!(779 == 32768)) {
        v8 := GenS1{F0: "hello"}
        var v9 uint16 = 255
        v10 := []int{2147483647}
        _ = v10
        _ = v8
        _ = v9
    } else {
        p1 = p1
    }
    var v11 bool = (372 != p0)
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return 1
}
func genFn1() int8 {
    var v12 float32 = 0.0
    var v13 int64 = 9223372036854775807
    var v14 float64 = (1e10 - ((-2.25 + 3.14159) + 1.5))
    if false {
        switch int(v13) {
        case 127:
            v15 := GenS1{F0: "\n\t"}
            _ = v15
        }
    } else {
        var v16 string = "foo bar"
        v17 := []int{9223372036854775807}
        _ = v16
        _ = v17
    }
    v14 += v14
    _ = v12
    _ = v13
    _ = v14
    return -128
}

func main() {
    var v1 bool = true
    v2 := []string{}
    var v3 int32 = 1
    var v4 int32 = v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v18 uint32 = 0
    var v19 int16 = -343
    v20 := GenS2{F0: -9223372036854775808, F1: -119, F2: (v18 >> ((255) & 63))}
    for fi21 := 0; fi21 < 2; fi21++ {
        if (("a" + "hello") < "a") {
            if v1 {
                for fi22 := 0; fi22 < 18; fi22++ {
                    v1 = false
                    var v23 string = "a"
                    _ = fi22
                    _ = v23
                }
                for fi24 := 0; fi24 < 19; fi24++ {
                    if (((float32(float64((0.0 * float32(0.0)))) - 1.5) * float32(1.5)) <= 1e10) {
                        v25 := GenS3{F0: 1, F1: (v19 | (v19 << ((18446744073709551615) & 63))), F2: (fi24 + -681)}
                        var v26 int8 = int8(fi21)
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS2{F0: 625, F1: (v4 - (v3 - 0)), F2: 65535}
                    var v28 float32 = (float32((1e10 * 0.0)) + 0.0)
                    v18 = v18
                    _ = fi24
                    _ = v27
                    _ = v28
                }
                var v29 uint8 = 1
                _ = v29
            }
        }
        v4 -= -365
        if (v1 && ((((4294967295 < 674) && (!false)) || (genFn1() <= -128)) && false)) {
            v30 := GenS3{F0: 42, F1: 0, F2: fi21}
            _ = v30
        } else {
            v3 -= v3
            switch ("a" + "foo bar") {
            case "hello":
                var v31 int64 = 0
                _ = v31
            case "":
                var v32 uint64 = 902
                _ = v32
            case "a":
                v33 := GenS3{F0: 190, F1: (v19 + (-(v19))), F2: (fi21 >> ((764) & 63))}
                _ = v33
            }
        }
        _ = fi21
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        fmt.Println(v1, v4)
        fmt.Println(v3)
        for fi35 := 0; fi35 < 9; fi35++ {
            v1 = ((v1 || (("hello" + "hello") <= (("foo bar" + "\n\t") + "\n\t"))) || v1)
            if len(v2) > 0 {
                v2[0] = "a"
            }
            v36 := GenS3{F0: 4294967295, F1: v19, F2: fi34}
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    v4 = (v3 & v3)
    var v37 uint8 = 128
    if true {
        v18 = v18
        v19 += v19
        v3 = int32(v37)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
}
