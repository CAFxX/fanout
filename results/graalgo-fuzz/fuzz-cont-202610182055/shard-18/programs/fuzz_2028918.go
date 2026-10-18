package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 string }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS5 struct { F0 string; F1 uint32; F2 uint32 }
type GenS6 struct { F0 uint; F1 bool }
func genFn0(p0 bool) int8 {
    switch "foo bar" {
    case "a":
        v5 := GenS1{F0: 533, F1: (((1 < -1) || p0) || false), F2: ""}
        _ = v5
    case "foo bar":
        p0 = (!(p0 && p0))
    case "\n\t":
        p0 = (42 != 42)
    default:
        fmt.Println(p0)
    }
    p0 = p0
    _ = p0
    return 1
}
func genFn1(p0 uint32, p1 int32, p2 uint8) string {
    v6 := GenS2{F0: (true && ((p0 >> ((uint(p0)) & 63)) < (^(p0)))), F1: false}
    var v7 uint64 = 18446744073709551615
    switch -119 {
    case 1:
        v8 := []int{-16, 1}
        _ = v8
    case 2147483647:
        switch 4294967295 {
        case 485:
            var v9 float32 = 1e10
            _ = v9
        default:
            v10 := GenS3{F0: 255, F1: 65535, F2: (p0 - 2147483648)}
            _ = v10
        }
    default:
        v11 := []int{-1, -1, -9223372036854775808, 32768}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    var v1 int16 = 32767
    var v2 float64 = 1e10
    var v3 float64 = 1e10
    var v4 float64 = 0.0
    switch "a" {
    case "hello":
        v12 := []int{-117, 9223372036854775807, -2147483648}
        _ = v12
    default:
        if (!(538 > 18446744073709551615)) {
            if true {
                var v13 bool = true
                v14 := []int{-907}
                fmt.Println(v13, v4, v4)
                _ = v13
                _ = v14
            }
            if false {
                var v15 int16 = v1
                _ = v15
            }
        }
    }
    v16 := GenS1{F0: 128, F1: true, F2: ("\n\t" + ("" + ("\n\t" + "foo bar")))}
    v17 := []int{1, 127, 453, -9223372036854775808}
    switch 4294967295 {
    case -843:
        if (genFn0((100 != 42)) >= 19) {
            v17 = append(v17, 0)
        }
    case 32767:
        v2 = v4
    }
    v4 = float64(0.1)
    var v18 uint = 42
    var v19 int32 = 830
    var v20 int32 = (-(v19))
    if len(v17) > 0 {
        v17[0] = 0
    }
    if (((9223372036854775807 > 9223372036854775807) && (true || true)) || (!(v20 <= v19))) {
        v2 = v4
        if true {
            v1 = 32767
            v21 := GenS2{F0: true, F1: true}
            _ = v21
        }
        if true {
            v4 += v3
        } else {
            var v22 uint32 = 330
            if false {
                v23 := GenS3{F0: 32767, F1: v18, F2: v22}
                _ = v23
            } else {
                v24 := GenS4{F0: 42, F1: 255, F2: (v20 << (((v18 % (((v18 - (^(v18)))) | 1))) & 63))}
                var v25 uint8 = 111
                v4 += 0.1
                _ = v24
                _ = v25
            }
            _ = v22
        }
    } else {
        v2 = v3
        v26 := GenS3{F0: 45, F1: v18, F2: 514}
        v27 := GenS5{F0: "a", F1: 2147483647, F2: 2147483647}
        _ = v26
        _ = v27
    }
    v17 = append(v17, -690)
    v3 = 0.0
    fmt.Println(v1)
    for fi28 := 0; fi28 < 20; fi28++ {
        for fi29 := 0; fi29 < 3; fi29++ {
            var v30 int32 = v20
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    if (v2 == v2) {
        v4 += float64((0.1 - ((-2.25 * (0.0 + 0.1)) - (((((0.0 * 0.0) * float32(v2)) * 1.5) * -2.25) - 0.0))))
    }
    switch ((("foo bar" + "hello") + "foo bar") + ("" + "hello")) {
    case "":
        for fi31 := 0; fi31 < 7; fi31++ {
            if true {
                var v32 int64 = -751
                _ = v32
            } else {
                var v33 int = (fi31 % ((fi31) | 1))
                _ = v33
            }
            for fi34 := 0; fi34 < 12; fi34++ {
                var v35 float32 = ((1e10 / 0.1) + float32(0.1))
                switch "\n\t" {
                case "hello":
                    switch (fi31 + -2147483648) {
                    case 836:
                        fmt.Println(v3)
                    case 4294967295:
                        if (!false) {
                            v1 += v1
                            v36 := GenS1{F0: (fi31 % ((fi34) | 1)), F1: true, F2: "\n\t"}
                            var v37 string = ""
                            v38 := GenS5{F0: (v37 + v37), F1: 2, F2: 934}
                            _ = v36
                            _ = v37
                            _ = v38
                        } else {
                            var v39 float64 = (v2 * v4)
                            v40 := GenS6{F0: 18446744073709551615, F1: ((127 >= genFn0(((v18 | 255) <= v18))) || (!(-128 == 69)))}
                            _ = v39
                            _ = v40
                        }
                    }
                case "":
                    if (true || false) {
                        var v41 uint32 = 4294967295
                        _ = v41
                    }
                case "a":
                    fmt.Println(v2, v1)
                }
                _ = fi34
                _ = v35
            }
            _ = fi31
        }
    default:
        for fi42 := 0; fi42 < 15; fi42++ {
            switch 9223372036854775807 {
            case 84:
                v43 := GenS2{F0: ((v18 % (((v18 - v18)) | 1)) < 0), F1: true}
                _ = v43
            case -391:
                for fi44 := 0; fi44 < 17; fi44++ {
                    v45 := GenS1{F0: (fi44 * (fi42 - (fi42 - (fi42 / ((0) | 1))))), F1: true, F2: "\n\t"}
                    _ = fi44
                    _ = v45
                }
            case 42:
                for fi46 := 0; fi46 < 18; fi46++ {
                    if (42 == 4294967295) {
                        fmt.Println(fi46, v4, v2)
                    } else {
                        v47 := GenS3{F0: 876, F1: v18, F2: 935}
                        _ = v47
                    }
                    v17 = append(v17, -927)
                    var v48 int32 = 42
                    var v49 uint8 = 1
                    _ = fi46
                    _ = v48
                    _ = v49
                }
            }
            v3 = 1e10
            _ = fi42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
