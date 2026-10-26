package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 int; F2 string }
type GenS6 struct { F0 uint; F1 int64 }
func genFn0(p0 uint, p1 int, p2 bool) bool {
    var v5 string = ("\n\t" + "foo bar")
    if true {
        var v6 string = v5
        _ = v6
    }
    switch p1 {
    case 32768:
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int8 = 42
            var v9 uint16 = 426
            v8 += v8
            _ = fi7
            _ = v8
            _ = v9
        }
    case 522:
        var v10 int64 = 797
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (((((p1 | p1) <= p1) || (!(42 < 1))) || (0.1 != (3.14159 - 1.5))) && p2)
}

func main() {
    v1 := []int{4294967295, -1, 9223372036854775807}
    v2 := []uint32{65535, 721, 2, 2147483647}
    v3 := make(map[int]bool)
    v4 := []string{"hello", "\n\t", ("a" + ("foo bar" + "foo bar")), "hello"}
    switch "" {
    case "foo bar":
        v11 := GenS1{F0: 2}
        _ = v11
    default:
        if (9223372036854775807 > 42) {
            var v12 string = "\n\t"
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint8 = 42
        _ = fi13
        _ = v14
    }
    if false {
        var v15 int32 = 42
        _ = v15
    } else {
        var v16 int64 = 869
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        fmt.Println(fi17)
        var v18 uint8 = 127
        _ = fi17
        _ = v18
    }
    var v19 float64 = float64(3.14159)
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 3; fi21++ {
            v19 -= v19
            v22 := GenS1{F0: 729}
            v19 += 0.0
            v23 := GenS2{F0: 760, F1: (false || (42 > -414))}
            _ = fi21
            _ = v22
            _ = v23
        }
        switch "\n\t" {
        case "a":
            v19 += ((v19 / ((1e10 + v19) / -2.25)) * v19)
        case "hello":
            v1 = append(v1, (^(fi20)))
        default:
            var v24 float32 = (1e10 - 0.1)
            _ = v24
        }
        var v25 int = fi20
        _ = fi20
        _ = v25
    }
    switch 128 {
    case 9223372036854775807:
        if true {
            v19 += v19
            v19 = v19
        }
    case -848:
        v26 := GenS3{F0: -1, F1: 2}
        _ = v26
    case 32768:
        v27 := v3[-9223372036854775808]
        _ = v27
    default:
        var v28 string = "foo bar"
        _ = v28
    }
    v19 -= 0.1
    for fi29 := 0; fi29 < 9; fi29++ {
        for fi30 := 0; fi30 < 12; fi30++ {
            fmt.Println(v19)
            v31 := GenS4{F0: "\n\t"}
            _ = fi30
            _ = v31
        }
        if true {
            if ((255 != 1) && (true || (true || true))) {
                var v32 float64 = ((v19 * v19) * float64(1.5))
                v33 := GenS5{F0: 356, F1: fi29, F2: "foo bar"}
                if ((-263 > 208) && false) {
                    v19 = (v32 + v32)
                    if true {
                        v19 += 3.14159
                        v32 = -2.25
                    } else {
                        for fi34 := 0; fi34 < 9; fi34++ {
                            v35 := GenS5{F0: 1, F1: (-(fi34)), F2: ("hello" + ("a" + (("foo bar" + "a") + "hello")))}
                            v36 := v3[(fi34 % (((fi34 & (fi34 >> ((4294967295) & 63)))) | 1))]
                            v32 = (((1e10 / 0.1) - (((v19 / v32) - 1e10) - float64(0.1))) / ((v19 * (v32 / (0.0 - 3.14159))) - 1.5))
                            var v37 int16 = -932
                            _ = fi34
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 14; fi38++ {
                            var v39 uint = 42
                            v40 := GenS4{F0: ("foo bar" + ("" + "a"))}
                            var v41 uint16 = 458
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    }
                } else {
                    var v42 int32 = 293
                    _ = v42
                }
                _ = v32
                _ = v33
            } else {
                v19 -= v19
                v43 := GenS6{F0: 986, F1: 42}
                v44 := GenS3{F0: (fi29 * (fi29 * (fi29 >> ((65) & 63)))), F1: 931}
                _ = v43
                _ = v44
            }
        } else {
            var v45 uint8 = 42
            _ = len(v3)
            _ = v45
        }
        _ = fi29
    }
    var v46 int8 = 51
    v19 += 3.14159
    fmt.Println(v46, v19)
    var v47 int16 = 1000
    switch 143 {
    case -936:
        v48 := GenS3{F0: 2147483647, F1: 2}
        _ = v48
    case 42:
        var v49 uint16 = 42
        _ = v49
    }
    for fi50 := 0; fi50 < 7; fi50++ {
        v51 := GenS3{F0: 558, F1: 65535}
        v52 := GenS2{F0: 42, F1: true}
        _ = fi50
        _ = v51
        _ = v52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
}
