package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int, p1 uint16) int8 {
    v5 := []int{(p0 / ((9223372036854775807) | 1)), p0, int(p1)}
    p0 += (p0 | (p0 << ((409) & 63)))
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1(p0 int32, p1 bool, p2 int8) int64 {
    p0 = p0
    v6 := []int{-646, 32767, 9223372036854775807, -711}
    var v7 int64 = -9223372036854775808
    v6 = append(v6, 891)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := []string{"a", (("a" + ("hello" + "\n\t")) + ""), (("foo bar" + "a") + ("\n\t" + "")), "foo bar"}
    var v2 int64 = 866
    var v3 uint32 = 768
    var v4 uint64 = 887
    v2 += (v2 - v2)
    if (true || false) {
        fmt.Println(v3)
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 16; fi9++ {
                if (fi9 != (fi8 + (fi9 - (fi8 & fi8)))) {
                    fmt.Println(fi9, v4)
                    var v10 int8 = 0
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v10
                }
                var v11 uint = 65535
                _ = fi9
                _ = v11
            }
            v2 += 9223372036854775807
            switch fi8 {
            case 857:
                fmt.Println(v3, fi8)
            case -595:
                v1 = append(v1, "")
            case 0:
                v3 -= 387
            default:
                if (!true) {
                    v2 = -869
                } else {
                    var v12 float32 = 1.5
                    v2 = (v2 ^ genFn1(-842, true, -113))
                    _ = v12
                }
            }
            _ = fi8
        }
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        switch "foo bar" {
        case "hello":
            fmt.Println(fi13, v2, v3)
        case "foo bar":
            var v14 uint16 = 65535
            _ = v14
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        if true {
            var v15 uint = 3
            _ = v15
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            fmt.Println(v4)
            if (18446744073709551615 == 4294967295) {
                fmt.Println(v2, v4, v4)
                v18 := GenS1{F0: -18}
                _ = v18
            } else {
                fmt.Println(v3, v4, v3)
            }
            var v19 int8 = 53
            _ = fi17
            _ = v19
        }
        v2 = (v2 / ((v2) | 1))
        for fi20 := 0; fi20 < 10; fi20++ {
            fmt.Println(v3)
            _ = len(v1)
            _ = len(v1)
            _ = fi20
        }
        _ = fi16
    }
    fmt.Println(v2)
    if (255 > 191) {
        if (!false) {
            v4 = v4
        }
    }
    var v21 float32 = (((float32(1.5) - 1e10) / 1e10) / 1e10)
    var v22 int64 = (v2 + int64(v3))
    v1 = append(v1, "hello")
    var v23 uint64 = (^(v4))
    v24 := GenS1{F0: 74}
    for fi25 := 0; fi25 < 20; fi25++ {
        for fi26 := 0; fi26 < 19; fi26++ {
            _ = len(v1)
            fmt.Println(v23, v21)
            _ = fi26
        }
        _ = fi25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
