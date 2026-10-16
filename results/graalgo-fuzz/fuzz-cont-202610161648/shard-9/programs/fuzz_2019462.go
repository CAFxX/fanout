package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0() int8 {
    switch "\n\t" {
    case "foo bar":
        var v4 bool = (!(!true))
        _ = v4
    }
    switch "\n\t" {
    case "":
        var v5 uint = 255
        _ = v5
    default:
        v6 := GenS1{F0: 65535, F1: 1}
        _ = v6
    }
    var v7 float32 = ((0.0 * 0.1) / 0.1)
    fmt.Println(v7)
    _ = v7
    return -1
}
func genFn1(p0 int16, p1 bool) int64 {
    v8 := []int{4294967295, -9223372036854775808}
    var v9 int8 = 0
    v10 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 2
    v2 := []string{}
    v3 := []string{""}
    v1 = (v1 >> ((4294967295) & 63))
    v1 -= v1
    var v11 int64 = genFn1(130, (((false || ((true || false) && (!(2 == 4294967295)))) || false) && true))
    for fi12 := 0; fi12 < 18; fi12++ {
        if len(v2) > 0 {
            v2[0] = ("" + ("foo bar" + "hello"))
        }
        fmt.Println(v11, fi12)
        v1 += v1
        _ = fi12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        if true {
            var v14 bool = true
            _ = v14
        }
        _ = fi13
    }
    fmt.Println(v1)
    if (32768 <= -9223372036854775808) {
        switch ("foo bar" + (("hello" + (("" + "") + "foo bar")) + "a")) {
        case "\n\t":
            v1 -= v1
        case "a":
            for fi15 := 0; fi15 < 1; fi15++ {
                switch fi15 {
                case 2147483647:
                    v1 -= 448
                case 9223372036854775807:
                    _ = len(v3)
                }
                _ = fi15
            }
        case "hello":
            if true {
                for fi16 := 0; fi16 < 4; fi16++ {
                    v3 = append(v3, "")
                    var v17 string = "a"
                    _ = fi16
                    _ = v17
                }
                v1 += v1
                switch 0 {
                case -373:
                    v11 = 0
                }
            }
        default:
            if (false || true) {
                v3 = append(v3, "")
                var v18 bool = (4294967295 > v1)
                v19 := GenS3{F0: 0, F1: 810}
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 1; fi20++ {
                    var v21 int8 = -128
                    _ = fi20
                    _ = v21
                }
            }
        }
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
