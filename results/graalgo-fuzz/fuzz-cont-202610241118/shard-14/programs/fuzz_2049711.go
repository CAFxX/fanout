package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 string, p1 bool, p2 int16) int64 {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 bool = (p0 > (p0 + p0))
        _ = fi7
        _ = v8
    }
    if (456 > 249) {
        if (4294967295 == 705) {
            switch "hello" {
            case "a":
                p2 -= p2
            case "\n\t":
                p0 = p0
            case "foo bar":
                switch p0 {
                case "hello":
                    for fi9 := 0; fi9 < 2; fi9++ {
                        fmt.Println(p0)
                        v10 := []int{(fi9 & (^(fi9)))}
                        switch "foo bar" {
                        case "":
                            if true {
                                var v11 uint8 = 127
                                p1 = true
                                _ = v11
                            }
                        default:
                            fmt.Println(p1, fi9, fi9)
                        }
                        _ = fi9
                        _ = v10
                    }
                case "a":
                    fmt.Println(p1, p1)
                }
            default:
                v12 := []int{}
                _ = v12
            }
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := GenS1{F0: 65535}
                fmt.Println(p2, p2)
                _ = fi13
                _ = v14
            }
        } else {
            var v15 bool = p1
            for fi16 := 0; fi16 < 4; fi16++ {
                if (!v15) {
                    var v17 uint16 = 255
                    _ = v17
                } else {
                    fmt.Println(fi16, fi16, fi16)
                    var v18 uint32 = 2147483648
                    var v19 int32 = -2147483648
                    _ = v18
                    _ = v19
                }
                var v20 uint64 = 4294967295
                _ = fi16
                _ = v20
            }
            v21 := []int{42, -9223372036854775808}
            _ = v15
            _ = v21
        }
    } else {
        p0 = ("foo bar" + p0)
    }
    if (0 < 2) {
        v22 := []int{128, -184, 32767}
        p2 += 1
        _ = v22
    } else {
        p0 = (p0 + (p0 + (("hello" + p0) + (p0 + "foo bar"))))
    }
    p2 += (p2 & 1)
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{65535, 42, 2147483647}
    v3 := make(map[string]bool)
    var v4 uint64 = 18446744073709551615
    v5 := make(map[uint32]int)
    var v6 int8 = 127
    var v23 uint8 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v24 string = ("hello" + "foo bar")
    v25 := GenS1{F0: 1}
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 float32 = ((1e10 + float32((-2.25 / 3.14159))) / 1.5)
        _ = fi26
        _ = v27
    }
    var v28 int = 355
    if true {
        var v29 float32 = -2.25
        v4 += (v4 - (^(v4)))
        _ = v29
    }
    if false {
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 uint32 = 2147483647
            v24 = v24
            _ = fi30
            _ = v31
        }
    } else {
        switch v28 {
        case 1:
            for fi32 := 0; fi32 < 2; fi32++ {
                v33 := GenS1{F0: 32767}
                _ = fi32
                _ = v33
            }
        case 9223372036854775807:
            v28 += v28
        default:
            v34 := GenS1{F0: 1}
            _ = v34
        }
        v24 = (v24 + v24)
    }
    v23 += v23
    v23 += 127
    if ((v4 & v4) < v4) {
        v35 := GenS2{F0: -735}
        _ = v35
    } else {
        var v36 uint64 = (^(v4))
        var v37 uint16 = 2
        _ = len(v2)
        _ = v36
        _ = v37
    }
    if (!true) {
        fmt.Println(v24, v23)
        var v38 uint16 = 2
        _ = v38
    } else {
        v4 -= (v4 << ((826) & 63))
        switch 9223372036854775807 {
        case 32767:
            if (v4 <= 0) {
                _ = len(v2)
                var v39 float32 = (0.1 * (3.14159 + (0.0 / 3.14159)))
                _ = v39
            } else {
                var v40 int32 = -1
                switch 9223372036854775807 {
                case 42:
                    fmt.Println(v4, v28)
                case 630:
                    var v41 uint = 4294967295
                    _ = v41
                }
                _ = v40
            }
        case -9223372036854775808:
            switch v24 {
            case "foo bar":
                var v42 uint = 255
                _ = v42
            case "a":
                v6 = (-(v6))
            default:
                v28 = (-(v28))
            }
        default:
            var v43 float32 = 0.1
            _ = v43
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
