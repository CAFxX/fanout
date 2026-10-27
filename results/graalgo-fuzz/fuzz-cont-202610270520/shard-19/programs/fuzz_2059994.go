package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }
func genFn0() int16 {
    if (!true) {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: 0, F1: 2}
            _ = fi7
            _ = v8
        }
    } else {
        if (false && false) {
            var v9 uint64 = 111
            _ = v9
        } else {
            if ((3.14159 < 1e10) || (1 != 0)) {
                var v10 uint = 65535
                _ = v10
            }
            var v11 int8 = -1
            _ = v11
        }
        var v12 float32 = (0.1 / 1.5)
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int32 = 648
        switch (fi13 + fi13) {
        case -1:
            v15 := []int{fi13, 42, (fi13 - 22)}
            _ = v15
        case -85:
            fmt.Println(v14, v14)
        }
        switch "a" {
        case "foo bar":
            for fi16 := 0; fi16 < 7; fi16++ {
                v17 := []int{-882}
                var v18 float64 = float64(0.1)
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        _ = fi13
        _ = v14
    }
    v19 := []int{}
    _ = v19
    return 1000
}
func genFn1(p0 uint) int8 {
    var v20 uint16 = 255
    p0 = (p0 + p0)
    _ = p0
    _ = v20
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 1000
    v3 := []uint32{284, 354, 538}
    var v4 int = 590
    v5 := []string{"a", ("\n\t" + ("\n\t" + ("a" + "hello"))), ("a" + "a"), "a"}
    var v6 int = v4
    var v21 uint8 = 38
    for fi22 := 0; fi22 < 13; fi22++ {
        v6 -= (v4 >> ((18446744073709551615) & 63))
        if false {
            var v23 int32 = -1
            if false {
                var v24 uint32 = 42
                v6 = fi22
                switch "\n\t" {
                case "hello":
                    v6 += (v6 + 2147483647)
                case "foo bar":
                    switch 9223372036854775807 {
                    case 128:
                        _ = len(v1)
                    case 925:
                        v25 := GenS1{F0: 522, F1: (v21 & v21)}
                        _ = v25
                    case 0:
                        v23 = (v23 >> ((145) & 63))
                    default:
                        fmt.Println(v21)
                    }
                case "\n\t":
                    v21 += v21
                }
                _ = v24
            }
            _ = v23
        } else {
            if (4294967295 == uint64(v21)) {
                v6 = (v6 & v6)
                v21 -= 45
                v4 = v6
            }
            if false {
                for fi26 := 0; fi26 < 5; fi26++ {
                    v21 += (v21 << ((255) & 63))
                    v27 := GenS1{F0: 187, F1: (v21 ^ 2)}
                    var v28 string = "\n\t"
                    v2 = v2
                    v29 := GenS2{F0: 18446744073709551615}
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                switch "foo bar" {
                case "foo bar":
                    for fi30 := 0; fi30 < 14; fi30++ {
                        v31 := GenS1{F0: 4294967295, F1: (v21 / (((v21 - 2)) | 1))}
                        v4 = -371
                        _ = fi30
                        _ = v31
                    }
                }
            }
            v32 := GenS3{F0: v2, F1: 0, F2: 2}
            _ = v32
        }
        v33 := GenS2{F0: 4294967295}
        _ = fi22
        _ = v33
    }
    v21 += v21
    switch ("" + (((("a" + (((((("hello" + "") + "foo bar") + ("" + ("\n\t" + "foo bar"))) + "") + "a") + "")) + ("a" + ("foo bar" + ("\n\t" + "hello")))) + ("a" + ("a" + "a"))) + "foo bar")) {
    case "a":
        v21 = v21
    case "":
        v34 := GenS1{F0: 137, F1: (v21 & 0)}
        _ = v34
    default:
        if false {
            v21 -= (v21 * v21)
            var v35 string = ""
            v35 = "\n\t"
            _ = v35
        }
    }
    v36 := GenS4{F0: v21}
    if true {
        fmt.Println(v6, v21)
        v4 = v4
        for fi37 := 0; fi37 < 16; fi37++ {
            fmt.Println(v4)
            _ = fi37
        }
    }
    v21 += (v21 << ((65535) & 63))
    v38 := GenS5{F0: (false && true)}
    for fi39 := 0; fi39 < 6; fi39++ {
        switch (fi39 + v4) {
        case 32767:
            var v40 int32 = -1
            _ = v40
        case -9223372036854775808:
            var v41 uint = 42
            _ = v41
        case 924:
            var v42 bool = (false && ((v2 & v2) <= v2))
            _ = v42
        }
        v5 = append(v5, "hello")
        var v43 float32 = 1.5
        _ = fi39
        _ = v43
    }
    v2 = (v2 - (^(v2)))
    v2 -= (v2 << ((4294967295) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
