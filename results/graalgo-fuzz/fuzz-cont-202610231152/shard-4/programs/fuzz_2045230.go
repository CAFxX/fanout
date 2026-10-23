package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 string }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint64; F1 string; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32 }
func genFn0(p0 string, p1 int16, p2 int16) string {
    switch 9223372036854775807 {
    case 1:
        var v4 int32 = -255
        _ = v4
    case 128:
        v5 := []int{-293, 128, 32767}
        _ = v5
    }
    v6 := GenS1{F0: p0, F1: 0, F2: "hello"}
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(p2)
        var v8 float64 = 1.5
        _ = fi7
        _ = v8
    }
    var v9 int64 = 42
    v9 -= 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return "hello"
}
func genFn1(p0 int8, p1 uint32, p2 uint) uint32 {
    var v10 float32 = -2.25
    switch 9223372036854775807 {
    case -776:
        p2 = (p2 - 1)
    }
    var v11 string = (("a" + ("foo bar" + ((("hello" + ("foo bar" + "")) + (("foo bar" + "foo bar") + "a")) + ("a" + "\n\t")))) + ("\n\t" + "hello"))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return p1
}

func main() {
    var v1 int = -9223372036854775808
    v2 := []int{v1, (v1 | (^(v1))), 0, v1}
    var v3 float64 = 0.1
    for fi12 := 0; fi12 < 16; fi12++ {
        v2 = append(v2, (fi12 << ((255) & 63)))
        if (2 < 1) {
            for fi13 := 0; fi13 < 7; fi13++ {
                v1 += (v1 & (v1 * 4294967295))
                var v14 string = "hello"
                var v15 float32 = 3.14159
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 int32 = 889
            var v17 int64 = 0
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 18; fi18++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v19 uint64 = 931
                _ = fi18
                _ = v19
            }
        }
        v20 := GenS2{F0: 1, F1: "a"}
        _ = fi12
        _ = v20
    }
    v21 := GenS2{F0: 346, F1: ("a" + "a")}
    var v22 bool = false
    v1 = (v1 << ((2) & 63))
    var v23 uint64 = 442
    fmt.Println(v1)
    fmt.Println(v22)
    switch v1 {
    case -2147483648:
        if v22 {
            var v24 bool = false
            _ = v24
        }
    case 938:
        var v25 int16 = 32767
        _ = v25
    case 127:
        var v26 int64 = -1
        _ = v26
    default:
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 uint32 = 1
            for fi29 := 0; fi29 < 1; fi29++ {
                switch "hello" {
                case "hello":
                    var v30 uint64 = 18446744073709551615
                    _ = v30
                case "foo bar":
                    if (false || v22) {
                        var v31 uint16 = 394
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if v22 {
                            var v32 int = 4294967295
                            v3 -= -2.25
                            if true {
                                v32 += -514
                            }
                            _ = v32
                        } else {
                            v1 += -484
                            fmt.Println(v23)
                        }
                        _ = v31
                    }
                case "\n\t":
                    if (v1 > (^(fi29))) {
                        v1 = (fi29 | 128)
                        var v33 int64 = int64(v23)
                        v34 := GenS3{F0: v23, F1: "a", F2: (!(v33 == (v33 - (-(v33)))))}
                        _ = v33
                        _ = v34
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                v22 = true
                v35 := GenS1{F0: "hello", F1: 42, F2: "a"}
                _ = fi29
                _ = v35
            }
            fmt.Println(v22, v3, v3)
            _ = fi27
            _ = v28
        }
    }
    var v36 int32 = -875
    if len(v2) > 0 {
        v2[0] = -1
    }
    _ = len(v2)
    if v22 {
        var v37 int64 = 0
        switch "a" {
        case "\n\t":
            if v22 {
                if (42 <= 100) {
                    v38 := GenS4{F0: ("a" + ((("a" + "") + "foo bar") + ((("a" + ("hello" + "hello")) + "a") + "foo bar")))}
                    _ = v38
                }
                v2 = append(v2, (v1 ^ (v1 / ((v1) | 1))))
            }
        default:
            v36 = 2147483647
        }
        switch v1 {
        case 110:
            _ = len(v2)
        case 128:
            var v39 uint = 2
            _ = v39
        }
        _ = v37
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v40 int32 = v36
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        var v42 uint8 = 128
        for fi43 := 0; fi43 < 14; fi43++ {
            v23 = 1
            v23 += v23
            var v44 uint = 1
            _ = fi43
            _ = v44
        }
        _ = fi41
        _ = v42
    }
    v45 := GenS5{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v36
    _ = v40
    _ = v45
}
