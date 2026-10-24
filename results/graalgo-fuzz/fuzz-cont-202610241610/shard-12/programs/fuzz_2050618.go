package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int; F1 uint; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint8; F2 uint16 }
func genFn0(p0 uint8) uint64 {
    v5 := []int{9223372036854775807, 786, -79}
    var v6 float32 = float32(0.0)
    var v7 uint = 65535
    p0 += p0
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 uint16, p1 uint64) int32 {
    var v8 float32 = 0.1
    var v9 uint = 628
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = (-2.25 == 0.0)
    var v3 float32 = float32(3.14159)
    v4 := make(map[string]string)
    var v10 int = 32768
    v2 = (0.1 <= 0.1)
    v10 += v10
    v11 := []int{42, v10, v10, (v10 | 0)}
    v4[""] = "a"
    v12 := GenS1{F0: 860, F1: 42}
    if (!(v2 && v2)) {
        if v2 {
            _ = len(v11)
        } else {
            switch (v10 >> ((79) & 63)) {
            case -666:
                fmt.Println(v3, v10, v10)
            case -279:
                switch 2147483647 {
                case -1:
                    if len(v11) > 0 {
                        v11[0] = 891
                    }
                case 127:
                    for fi13 := 0; fi13 < 16; fi13++ {
                        v14 := GenS1{F0: -194, F1: 495}
                        _ = fi13
                        _ = v14
                    }
                default:
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v16 := GenS1{F0: -1, F1: 32767}
                        _ = fi15
                        _ = v16
                    }
                }
            case -175:
                switch (v10 & (v10 * (^(v10)))) {
                case 128:
                    switch -964 {
                    case 128:
                        v2 = (!false)
                    case 127:
                        var v17 uint8 = 42
                        _ = v17
                    }
                default:
                    v18 := GenS2{F0: v10, F1: 42, F2: 1}
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := GenS1{F0: -32768, F1: 65535}
                fmt.Println(v2, v10)
                _ = fi19
                _ = v20
            }
        }
        v21 := GenS3{F0: 0, F1: 255}
        _ = v21
    } else {
        var v22 int = v10
        var v23 int64 = 703
        _ = len(v11)
        _ = v22
        _ = v23
    }
    var v24 uint = 65535
    if false {
        switch "foo bar" {
        case "":
            v2 = (1000 > 32767)
        default:
            v25 := GenS4{F0: 562}
            _ = v25
        }
    } else {
        switch (("hello" + "a") + "foo bar") {
        case "":
            var v26 bool = (599 >= -47)
            _ = v26
        case "foo bar":
            for fi27 := 0; fi27 < 10; fi27++ {
                for fi28 := 0; fi28 < 3; fi28++ {
                    v10 = (v10 / (((-(fi27))) | 1))
                    v2 = ("" <= "hello")
                    v4[("hello" + ("" + ""))] = ("foo bar" + ("a" + "hello"))
                    _ = fi28
                }
                var v29 int32 = -2147483648
                v30 := GenS5{F0: 42, F1: 2, F2: 61}
                _ = fi27
                _ = v29
                _ = v30
            }
        case "a":
            for fi31 := 0; fi31 < 20; fi31++ {
                v24 = v24
                _ = fi31
            }
        }
        var v32 uint8 = 42
        _ = v32
    }
    v10 = (v10 << ((18446744073709551615) & 63))
    var v33 uint64 = 2
    var v34 uint64 = (v33 % ((genFn0(0)) | 1))
    v35 := GenS2{F0: (v10 & (v10 >> (((v24 - 1)) & 63))), F1: 317, F2: 357}
    var v36 uint64 = 1
    switch "" {
    case "foo bar":
        if v2 {
            _ = len(v11)
            _ = len(v11)
            v37 := GenS4{F0: (-(v10))}
            _ = v37
        }
    case "hello":
        if ("" >= ("foo bar" + "hello")) {
            var v38 uint16 = 323
            v39 := GenS3{F0: v38, F1: v24}
            _ = v38
            _ = v39
        } else {
            v34 += v33
            var v40 float32 = v3
            v34 -= 753
            _ = v40
        }
    case "\n\t":
        v24 -= (v24 ^ v24)
    default:
        v24 = v24
    }
    for fi41 := 0; fi41 < 2; fi41++ {
        v34 = (v33 << ((v24) & 63))
        if ((v36 | v36) <= (v34 + (v36 | (v36 * genFn0(211))))) {
            v42 := GenS4{F0: (v10 - -778)}
            _ = v42
        } else {
            var v43 float32 = v3
            for fi44 := 0; fi44 < 15; fi44++ {
                var v45 int64 = 9223372036854775807
                _ = fi44
                _ = v45
            }
            _ = v43
        }
        _ = fi41
    }
    fmt.Println(v33, v2)
    v3 += ((v3 - v3) + v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
