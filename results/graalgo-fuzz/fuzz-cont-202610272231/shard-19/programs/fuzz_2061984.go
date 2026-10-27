package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint; F2 int16 }
type GenS3 struct { F0 int16; F1 bool; F2 int64 }
type GenS4 struct { F0 uint; F1 uint; F2 uint8 }
type GenS5 struct { F0 int; F1 uint64; F2 int8 }
func genFn0(p0 int) uint64 {
    v7 := GenS1{F0: (p0 / ((p0) | 1)), F1: 879, F2: 18446744073709551615}
    fmt.Println(p0)
    v8 := GenS2{F0: 195, F1: 18446744073709551615, F2: 948}
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 float64 = 0.1
        _ = fi9
        _ = v10
    }
    switch "a" {
    case "hello":
        fmt.Println(p0)
    case "":
        switch 127 {
        case 4294967295:
            p0 = 335
        default:
            var v11 int64 = 1
            _ = v11
        }
    case "foo bar":
        v12 := []int{p0, (p0 ^ (p0 * -501))}
        _ = v12
    }
    _ = p0
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 768
    v2 := []string{"hello", "hello", "foo bar"}
    var v3 int32 = 0
    var v4 uint32 = v1
    v5 := []int{}
    var v6 uint = 18446744073709551615
    v13 := GenS2{F0: -291, F1: (v6 & v6), F2: 1000}
    var v14 bool = (!(!(true && true)))
    v3 += v3
    switch ("hello" + ("a" + (((((("\n\t" + ((("a" + "a") + ((("\n\t" + ("" + (("" + "") + "a"))) + ("foo bar" + "hello")) + ("foo bar" + "foo bar"))) + ("a" + "\n\t"))) + ("a" + (("\n\t" + ("" + "foo bar")) + ("a" + ("a" + "\n\t"))))) + "") + "\n\t") + "\n\t") + ""))) {
    case "\n\t":
        v15 := GenS2{F0: 1000, F1: 532, F2: -32768}
        _ = v15
    case "":
        v4 -= v4
    case "foo bar":
        var v16 uint = (v6 ^ (v6 % ((v6) | 1)))
        _ = v16
    default:
        if v14 {
            v14 = v14
            var v17 uint16 = 42
            var v18 int8 = 74
            _ = v17
            _ = v18
        }
    }
    var v19 uint64 = 1
    if (false || (!v14)) {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        fmt.Println(v3)
    } else {
        v4 = 313
    }
    v20 := GenS1{F0: 748, F1: 2, F2: 2}
    for fi21 := 0; fi21 < 18; fi21++ {
        v22 := GenS2{F0: 32, F1: (v6 - v6), F2: 1000}
        switch "hello" {
        case "a":
            switch fi21 {
            case 42:
                switch -710 {
                case -1:
                    var v23 uint16 = 1
                    _ = v23
                default:
                    var v24 int = -2147483648
                    _ = v24
                }
            case -1:
                for fi25 := 0; fi25 < 11; fi25++ {
                    var v26 uint8 = 42
                    _ = fi25
                    _ = v26
                }
            case 128:
                v2 = append(v2, ("hello" + "foo bar"))
            default:
                var v27 uint16 = 1
                _ = v27
            }
        case "foo bar":
            var v28 int32 = (v3 * v3)
            _ = v28
        default:
            var v29 string = "a"
            _ = v29
        }
        _ = fi21
        _ = v22
    }
    v30 := GenS3{F0: -585, F1: (!v14), F2: 0}
    v31 := GenS3{F0: 95, F1: (!(!(1 <= -1))), F2: -507}
    v32 := GenS1{F0: 42, F1: 2, F2: (^(v19))}
    v14 = v14
    for fi33 := 0; fi33 < 1; fi33++ {
        if (1e10 == float64(float32(3.14159))) {
            for fi34 := 0; fi34 < 7; fi34++ {
                if (v14 || (((v19 | (^(v19))) >= uint64(v6)) || false)) {
                    switch fi34 {
                    case 2147483647:
                        v4 = v4
                    default:
                        switch "hello" {
                        case "\n\t":
                            v35 := GenS2{F0: 1000, F1: (v6 ^ v6), F2: 88}
                            _ = v35
                        default:
                            v36 := GenS4{F0: v6, F1: uint(v19), F2: 249}
                            _ = v36
                        }
                    }
                    if v14 {
                        v4 += v1
                    } else {
                        v37 := GenS5{F0: 32768, F1: v19, F2: 127}
                        if len(v2) > 0 {
                            v2[0] = (("hello" + "\n\t") + ("\n\t" + ("\n\t" + "foo bar")))
                        }
                        v1 -= uint32(v19)
                        _ = v37
                    }
                    v5 = append(v5, (fi34 % ((-2147483648) | 1)))
                } else {
                    if (!((fi34 != (^(fi33))) && v14)) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
                _ = fi34
            }
            var v38 int = (^(fi33))
            v38 = fi33
            _ = v38
        }
        _ = fi33
    }
    v39 := GenS3{F0: 0, F1: ((v1 >> ((v6) & 63)) < 65535), F2: 0}
    for fi40 := 0; fi40 < 6; fi40++ {
        var v41 int = (-(fi40))
        _ = len(v2)
        if true {
            v1 = (v4 % (((v4 << ((uint(v1)) & 63))) | 1))
        }
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
