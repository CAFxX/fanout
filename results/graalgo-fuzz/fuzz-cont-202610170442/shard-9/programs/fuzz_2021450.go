package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int64; F1 string }
func genFn0(p0 uint) uint64 {
    var v4 uint64 = 238
    var v5 int = 867
    switch v5 {
    case 127:
        v5 += v5
    default:
        var v6 uint8 = 1
        _ = v6
    }
    _ = p0
    _ = v4
    _ = v5
    return v4
}
func genFn1() uint16 {
    switch 564 {
    case 947:
        var v7 float64 = (0.0 * 1.5)
        _ = v7
    case 476:
        v8 := []int{-1, 32767}
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        if false {
            if (2 < 32) {
                switch "foo bar" {
                case "\n\t":
                    var v10 int64 = 42
                    _ = v10
                case "":
                    for fi11 := 0; fi11 < 6; fi11++ {
                        var v12 int = fi9
                        _ = fi11
                        _ = v12
                    }
                case "foo bar":
                    var v13 uint16 = 65535
                    _ = v13
                default:
                    var v14 float32 = 1e10
                    _ = v14
                }
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 uint8 = 112
                    _ = fi15
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 uint8 = 1
                _ = fi17
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 3; fi19++ {
            if (1 < fi9) {
                var v20 uint = 920
                _ = v20
            } else {
                v21 := GenS1{F0: 4294967295, F1: -4}
                v22 := []int{(fi9 ^ (fi19 & fi19))}
                var v23 string = ("" + "\n\t")
                var v24 uint = 4294967295
                _ = v21
                _ = v22
                _ = v23
                _ = v24
            }
            switch "" {
            case "a":
                if (430 != 4294967295) {
                    switch fi9 {
                    case 674:
                        var v25 int = (-(fi9))
                        _ = v25
                    case 0:
                        v26 := []int{fi19, (fi9 & 4294967295)}
                        _ = v26
                    case -74:
                        v27 := []int{}
                        _ = v27
                    }
                    switch ("\n\t" + "\n\t") {
                    case "a":
                        fmt.Println(fi9, fi19)
                    }
                    v28 := GenS2{F0: 0, F1: "\n\t"}
                    _ = v28
                } else {
                    var v29 uint32 = 337
                    v30 := GenS3{F0: 127, F1: (-1 != 1)}
                    _ = v29
                    _ = v30
                }
                var v31 int16 = 32767
                _ = v31
            default:
                v32 := []int{-271, -718, -9223372036854775808}
                v33 := GenS4{F0: 1, F1: ""}
                if len(v32) >= 2 {
                    v32 = v32[1:len(v32)]
                }
                _ = v32
                _ = v33
            }
            fmt.Println(fi9)
            switch (fi19 - fi19) {
            case 232:
                switch (fi19 * 32768) {
                case -611:
                    var v34 int32 = 482
                    _ = v34
                default:
                    switch (fi9 * (-(fi9))) {
                    case 32767:
                        var v35 uint8 = 249
                        _ = v35
                    case -1:
                        v36 := []int{fi19, (fi9 >> ((2) & 63))}
                        _ = v36
                    default:
                        var v37 uint16 = 2
                        _ = v37
                    }
                }
            case -438:
                var v38 int = fi19
                _ = v38
            case -2147483648:
                v39 := []int{127}
                _ = v39
            default:
                if (!true) {
                    v40 := []int{fi19}
                    if (!false) {
                        var v41 int8 = -128
                        _ = v41
                    }
                    var v42 float64 = (1e10 - 1.5)
                    _ = v40
                    _ = v42
                }
            }
            _ = fi19
        }
        _ = fi9
    }
    return 421
}

func main() {
    var v1 float64 = (-2.25 / 1.5)
    var v2 int16 = 264
    var v3 string = ("foo bar" + (("foo bar" + "") + "hello"))
    var v43 int64 = 9223372036854775807
    v2 = v2
    switch -2147483648 {
    case 9223372036854775807:
        v44 := []int{1, 42}
        _ = v44
    case 42:
        var v45 float32 = ((float32((float64(((3.14159 / 0.1) + 1e10)) * 1e10)) + (((-2.25 * float32(-2.25)) - 1.5) / float32(v1))) - (3.14159 - (0.1 / 0.1)))
        _ = v45
    }
    v46 := []int{42, 0, 2147483647}
    for fi47 := 0; fi47 < 9; fi47++ {
        for fi48 := 0; fi48 < 2; fi48++ {
            var v49 int = 2147483647
            fmt.Println(v3)
            if len(v46) > 0 {
                v46[0] = (fi47 - fi47)
            }
            _ = fi48
            _ = v49
        }
        _ = fi47
    }
    for fi50 := 0; fi50 < 13; fi50++ {
        v46 = append(v46, fi50)
        fmt.Println(v43, v2, v43)
        _ = fi50
    }
    var v51 uint8 = 127
    v2 -= (v2 / (((-(v2))) | 1))
    fmt.Println(v1)
    v52 := GenS3{F0: 12, F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v46
    _ = v51
    _ = v52
}
