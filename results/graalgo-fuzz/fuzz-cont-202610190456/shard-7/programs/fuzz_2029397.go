package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 bool }
type GenS3 struct { F0 uint64; F1 int64; F2 int32 }
type GenS4 struct { F0 uint32; F1 string; F2 int16 }
func genFn0(p0 int8, p1 bool, p2 uint16) int {
    v5 := GenS1{F0: 2147483647, F1: -2147483648}
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: -60, F1: 0}
        p2 -= 32767
        v8 := []int{}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float64 = 0.1
    p0 += (p0 << ((42) & 63))
    var v10 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v9
    return 161
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]bool)
    var v3 int16 = -1
    var v4 uint = 65535
    v3 = (-(v3))
    v11 := GenS2{F0: 653, F1: 18446744073709551615, F2: (((908 > 636) && ((("hello" == ("a" + "\n\t")) || (("foo bar" + ("" + "")) > "a")) && (true && ((v4 | (v4 & (v4 & v4))) != 65535)))) || true)}
    var v12 string = (((("a" + "\n\t") + "hello") + "hello") + ("hello" + ("hello" + "foo bar")))
    switch v12 {
    case "":
        switch -436 {
        case 1:
            if ((100 < -128) && (!false)) {
                for fi13 := 0; fi13 < 15; fi13++ {
                    v3 -= v3
                    _ = fi13
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        case 2147483647:
            v3 -= v3
        default:
            var v14 uint8 = 2
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        v4 += 1
        _ = fi15
    }
    if true {
        v16 := GenS3{F0: 2, F1: int64(v4), F2: -1}
        fmt.Println(v4, v12, v4)
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 uint8 = 1
        _ = fi17
        _ = v18
    }
    v19 := GenS3{F0: 261, F1: 0, F2: 913}
    var v20 uint32 = 336
    _ = len(v2)
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := GenS2{F0: 0, F1: (^(v4)), F2: ((v12 + "a") < (v12 + (v12 + (v12 + "foo bar"))))}
        _ = len(v1)
        _ = fi21
        _ = v22
    }
    v12 = v12
    if (v12 < (("hello" + (v12 + "foo bar")) + (v12 + v12))) {
        fmt.Println(v12, v4)
        v20 -= (v20 | 2)
    }
    switch "foo bar" {
    case "\n\t":
        var v23 uint = 18446744073709551615
        _ = v23
    case "":
        v24 := GenS2{F0: 2, F1: uint(v20), F2: ("" < "hello")}
        _ = v24
    }
    fmt.Println(v12, v12)
    switch v12 {
    case "a":
        v3 -= (v3 << ((4294967295) & 63))
    case "foo bar":
        if false {
            if (1e10 == 0.1) {
                fmt.Println(v20)
            } else {
                if (true || (9223372036854775807 != 32768)) {
                    if (0.0 < 3.14159) {
                        v25 := GenS4{F0: v20, F1: "\n\t", F2: v3}
                        v1 = append(v1, (v20 / (((v20 << (((v4 & 42)) & 63))) | 1)))
                        _ = len(v1)
                        _ = v25
                    } else {
                        v12 = (("foo bar" + ((v12 + ("" + "")) + ("\n\t" + v12))) + (v12 + v12))
                    }
                } else {
                    fmt.Println(v12, v12, v12)
                    var v26 int32 = 0
                    _ = v26
                }
                var v27 int32 = 0
                fmt.Println(v3)
                _ = v27
            }
        } else {
            var v28 bool = (0 == 0)
            _ = v28
        }
    case "\n\t":
        var v29 uint32 = 65535
        _ = v29
    }
    v20 = (^(v20))
    if false {
        for fi30 := 0; fi30 < 13; fi30++ {
            v31 := GenS4{F0: 826, F1: (v12 + v12), F2: v3}
            if false {
                delete(v2, fi30)
            }
            v32 := GenS1{F0: 2147483647, F1: -1}
            _ = fi30
            _ = v31
            _ = v32
        }
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 uint16 = 1
            delete(v2, (fi33 - -1))
            v20 -= 746
            _ = fi33
            _ = v34
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
