package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int16 }
type GenS2 struct { F0 uint64; F1 string; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 int64 }
func genFn0(p0 int32, p1 int64) int32 {
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := []int{42, fi5, fi5}
        _ = fi5
        _ = v6
    }
    p1 += p1
    var v7 string = "\n\t"
    v8 := GenS1{F0: -2147483648, F1: -1, F2: 42}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (^(p0))
}
func genFn1(p0 string) int {
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS2{F0: 18446744073709551615, F1: p0, F2: 4294967295}
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 924
    v11 -= (v11 % ((698) | 1))
    v12 := []int{42, 32768}
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int = fi13
        var v15 string = (p0 + p0)
        p0 = ""
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v11
    _ = v12
    return 4294967295
}

func main() {
    v1 := []string{("" + "foo bar")}
    var v2 int = -1
    var v3 uint64 = 18446744073709551615
    var v4 int16 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 += genFn1(("\n\t" + "foo bar"))
    if false {
        v16 := GenS3{F0: 9223372036854775807}
        v17 := GenS4{F0: 42, F1: 655}
        if (!false) {
            v2 = genFn1("\n\t")
            for fi18 := 0; fi18 < 9; fi18++ {
                v3 += 1
                fmt.Println(fi18, v2, fi18)
                _ = fi18
            }
        } else {
            v1 = append(v1, "")
        }
        _ = v16
        _ = v17
    } else {
        v19 := GenS1{F0: genFn0(-772, -1), F1: 42, F2: v4}
        switch ("\n\t" + ("a" + ("a" + ("" + "")))) {
        case "":
            var v20 uint64 = (^(v3))
            _ = v20
        case "a":
            v3 -= 4294967295
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v21 uint = 1
        _ = v19
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := GenS5{F0: "a"}
        var v24 uint8 = 127
        _ = fi22
        _ = v23
        _ = v24
    }
    if len(v1) > 0 {
        v1[0] = ("" + (("\n\t" + "\n\t") + "a"))
    }
    var v25 uint8 = 42
    if ((!(!false)) || false) {
        for fi26 := 0; fi26 < 18; fi26++ {
            v27 := GenS4{F0: 504, F1: 555}
            switch -1 {
            case 4294967295:
                v28 := GenS6{F0: (v4 | v4), F1: -966}
                _ = v28
            case 188:
                if ((fi26 / ((fi26) | 1)) == -9223372036854775808) {
                    _ = len(v1)
                }
            }
            _ = fi26
            _ = v27
        }
    } else {
        v29 := GenS4{F0: 0, F1: 42}
        _ = v29
    }
    v3 = (v3 * v3)
    fmt.Println(v3, v4)
    v2 -= (v2 - (v2 >> ((528) & 63)))
    v3 -= (^(v3))
    var v30 int16 = -1
    v30 -= 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
}
