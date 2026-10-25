package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0(p0 int64) int16 {
    p0 = p0
    var v7 int32 = -184
    v8 := []int{172}
    _ = p0
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 int32, p1 uint32, p2 int8) int32 {
    v9 := []int{32768, 1, 9223372036854775807}
    v10 := GenS1{F0: "foo bar", F1: "hello"}
    var v11 int32 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return p0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{4294967295}
    var v3 float32 = 3.14159
    var v4 uint64 = 712
    v5 := make(map[int]uint32)
    var v6 int16 = -32768
    switch "a" {
    case "foo bar":
        if false {
            v2 = append(v2, 65535)
            _ = len(v1)
            var v12 int16 = (v6 ^ 32767)
            _ = v12
        }
    case "\n\t":
        v13 := GenS1{F0: ("foo bar" + "\n\t"), F1: "a"}
        _ = v13
    default:
        v14 := v5[127]
        _ = v14
    }
    v15 := GenS1{F0: "hello", F1: "\n\t"}
    switch -127 {
    case -9223372036854775808:
        v16 := v1[4294967295]
        _ = v16
    case -27:
        switch 2147483647 {
        case -1:
            switch 9223372036854775807 {
            case 1:
                fmt.Println(v6, v4)
            case -68:
                for fi17 := 0; fi17 < 1; fi17++ {
                    v18 := GenS1{F0: "foo bar", F1: "hello"}
                    _ = fi17
                    _ = v18
                }
            case 332:
                v19 := GenS1{F0: "hello", F1: "hello"}
                _ = v19
            default:
                var v20 float32 = 3.14159
                _ = v20
            }
        default:
            fmt.Println(v3)
        }
    default:
        var v21 float32 = v3
        _ = v21
    }
    v22 := GenS2{F0: 1, F1: -9223372036854775808}
    for fi23 := 0; fi23 < 17; fi23++ {
        v24 := v5[fi23]
        _ = fi23
        _ = v24
    }
    v3 -= v3
    for fi25 := 0; fi25 < 18; fi25++ {
        for fi26 := 0; fi26 < 8; fi26++ {
            fmt.Println(fi26)
            var v27 string = "\n\t"
            var v28 int = 127
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    var v29 uint16 = 255
    v4 = 203
    var v30 int32 = -1
    var v31 int32 = (v30 % (((v30 * (v30 - int32(v4)))) | 1))
    for fi32 := 0; fi32 < 10; fi32++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v33 uint16 = v29
        var v34 int64 = -9223372036854775808
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 bool = (4294967295 == 20)
    for fi36 := 0; fi36 < 11; fi36++ {
        v4 = v4
        _ = fi36
    }
    switch 1 {
    case 42:
        _ = len(v2)
    case 1:
        if (464 != 580) {
            v31 = v30
            if v35 {
                var v37 bool = v35
                _ = v37
            } else {
                if (255 >= 255) {
                    v35 = ((v6 << ((65535) & 63)) > (-(v6)))
                    v31 = (v31 % ((v30) | 1))
                }
            }
        } else {
            var v38 int8 = -100
            _ = v38
        }
    default:
        switch -9223372036854775808 {
        case -2147483648:
            fmt.Println(v4)
        default:
            fmt.Println(v31, v4)
        }
    }
    var v39 int16 = v6
    for fi40 := 0; fi40 < 17; fi40++ {
        var v41 int32 = (v30 & 905)
        _ = fi40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
