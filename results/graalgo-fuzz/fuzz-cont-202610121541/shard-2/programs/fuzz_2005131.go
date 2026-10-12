package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int; F2 string }
func genFn0() int {
    var v5 bool = (-32768 >= -1)
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint32 = 1
        v8 := []int{127}
        if len(v8) > 0 {
            v8[0] = (fi6 / ((fi6) | 1))
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS1{F0: "hello"}
        v11 := []int{(fi9 % (((fi9 ^ (^(fi9)))) | 1))}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v5
    return 897
}
func genFn1(p0 int64) int64 {
    if ((p0 / ((p0) | 1)) < (p0 << ((255) & 63))) {
        var v12 float32 = 1.5
        for fi13 := 0; fi13 < 12; fi13++ {
            if false {
                for fi14 := 0; fi14 < 12; fi14++ {
                    v15 := []int{-952}
                    _ = fi14
                    _ = v15
                }
                if (false && true) {
                    var v16 int16 = -212
                    _ = v16
                }
            }
            p0 += (p0 ^ 42)
            _ = fi13
        }
        _ = v12
    } else {
        p0 = p0
        var v17 uint32 = 0
        _ = v17
    }
    v18 := []int{genFn0(), -163}
    v19 := GenS1{F0: "a"}
    _ = len(v18)
    var v20 int16 = -927
    _ = p0
    _ = v18
    _ = v19
    _ = v20
    return 860
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 42
    var v3 int8 = 55
    var v4 uint64 = 902
    switch "\n\t" {
    case "\n\t":
        var v21 int64 = 0
        _ = v21
    case "":
        if (2147483648 != 1) {
            if true {
                for fi22 := 0; fi22 < 5; fi22++ {
                    v23 := GenS1{F0: "hello"}
                    _ = fi22
                    _ = v23
                }
            } else {
                var v24 int16 = -1
                _ = v24
            }
            var v25 int8 = -107
            var v26 int64 = 9223372036854775807
            _ = v25
            _ = v26
        }
    case "a":
        v3 += (v3 >> (((v2 + (v2 / ((v2) | 1)))) & 63))
    default:
        switch "foo bar" {
        case "hello":
            for fi27 := 0; fi27 < 19; fi27++ {
                var v28 bool = (!(-32768 == 1000))
                v29 := GenS1{F0: "a"}
                v30 := GenS2{F0: (v3 % ((v3) | 1))}
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        case "a":
            v31 := GenS2{F0: (v3 + v3)}
            _ = v31
        }
    }
    var v32 int8 = v3
    fmt.Println(v32, v3)
    switch 2147483647 {
    case 32768:
        var v33 uint8 = 235
        _ = v33
    case 127:
        v2 = 659
    case 2147483647:
        v34 := GenS2{F0: (v3 | (v3 + 127))}
        _ = v34
    default:
        if (1.5 >= (1e10 - float64(0.0))) {
            switch ("\n\t" + "hello") {
            case "\n\t":
                for fi35 := 0; fi35 < 5; fi35++ {
                    var v36 int32 = 0
                    _ = fi35
                    _ = v36
                }
            }
        } else {
            for fi37 := 0; fi37 < 13; fi37++ {
                _ = len(v1)
                v32 = (v3 >> ((uint(v4)) & 63))
                _ = fi37
            }
            fmt.Println(v2, v3, v4)
        }
    }
    v38 := []int{}
    delete(v1, "")
    delete(v1, "")
    switch ("a" + "foo bar") {
    case "hello":
        v3 += (v3 % (((v32 * -128)) | 1))
    }
    switch 1 {
    case 755:
        if (true && (int32(v3) <= 764)) {
            fmt.Println(v32)
            v2 -= (v2 & (^(v2)))
        }
    case 9223372036854775807:
        v39 := GenS3{F0: (("" + ("\n\t" + ("a" + "a"))) + "\n\t"), F1: 42, F2: "a"}
        _ = v39
    case -2147483648:
        var v40 int16 = 42
        _ = v40
    }
    fmt.Println(v3, v3, v4)
    v41 := GenS3{F0: "a", F1: -2147483648, F2: ("a" + "\n\t")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v41
}
