package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 bool; F1 uint32; F2 int16 }
type GenS6 struct { F0 int64 }

func main() {
    v1 := []int{}
    var v2 uint32 = 654
    var v3 uint16 = 42
    var v4 int32 = 2147483647
    v5 := []int{393, -1, 42, 900}
    if true {
        switch ("\n\t" + ((("" + "\n\t") + ("hello" + (("foo bar" + "\n\t") + ("" + "\n\t")))) + "")) {
        case "":
            var v6 float64 = (1e10 / 1e10)
            _ = v6
        }
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: 128, F1: 4294967295}
            _ = fi7
            _ = v8
        }
        var v9 int16 = 1000
        _ = v9
    } else {
        v3 += 2
        v10 := GenS2{F0: 255}
        v11 := GenS3{F0: 718, F1: 2147483647}
        _ = v10
        _ = v11
    }
    v12 := GenS4{F0: "a", F1: 486}
    _ = len(v1)
    fmt.Println(v3)
    var v13 float32 = 0.1
    var v14 uint8 = 42
    switch "a" {
    case "":
        if (true && true) {
            v13 += (v13 * 3.14159)
            if ((v14 ^ 42) <= v14) {
                fmt.Println(v4, v4)
            }
        } else {
            v14 = 163
            var v15 uint = 255
            for fi16 := 0; fi16 < 9; fi16++ {
                v15 = (v15 - (v15 << ((v15) & 63)))
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 uint = v15
                    v19 := GenS2{F0: (^(v14))}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
            _ = v15
        }
    case "hello":
        switch -1 {
        case -386:
            for fi20 := 0; fi20 < 10; fi20++ {
                for fi21 := 0; fi21 < 18; fi21++ {
                    var v22 int16 = 1
                    _ = fi21
                    _ = v22
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi23 := 0; fi23 < 10; fi23++ {
                    v13 -= (((-2.25 / v13) + float32(1.5)) - ((v13 + 0.1) * v13))
                    _ = fi23
                }
                _ = fi20
            }
        case 0:
            fmt.Println(v2, v3)
        case 32767:
            fmt.Println(v13)
        default:
            var v24 string = ("\n\t" + "hello")
            _ = v24
        }
    }
    v25 := GenS2{F0: v14}
    if len(v5) > 0 {
        v5[0] = 1
    }
    v26 := GenS5{F0: true, F1: (v2 % (((v2 * v2)) | 1)), F2: 32767}
    for fi27 := 0; fi27 < 12; fi27++ {
        fmt.Println(v4, v4, v2)
        _ = fi27
    }
    fmt.Println(v4, v14)
    if false {
        var v28 int16 = 32767
        _ = v28
    }
    if ((^(v4)) < v4) {
        var v29 int64 = -1
        _ = v29
    } else {
        if false {
            v13 -= v13
            switch 9223372036854775807 {
            case 407:
                var v30 float64 = 1e10
                _ = v30
            case 970:
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
            default:
                fmt.Println(v14, v13)
            }
            var v31 int8 = -1
            _ = v31
        }
        v2 = (v2 >> ((1) & 63))
    }
    v4 = v4
    for fi32 := 0; fi32 < 17; fi32++ {
        for fi33 := 0; fi33 < 3; fi33++ {
            fmt.Println(v13, fi33, fi32)
            _ = fi33
        }
        for fi34 := 0; fi34 < 3; fi34++ {
            var v35 int64 = -9223372036854775808
            v36 := GenS3{F0: -1, F1: 557}
            fmt.Println(v13, v4)
            _ = fi34
            _ = v35
            _ = v36
        }
        v37 := GenS3{F0: fi32, F1: 9223372036854775807}
        _ = fi32
        _ = v37
    }
    v38 := GenS6{F0: 9223372036854775807}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
