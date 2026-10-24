package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint64; F1 string }
func genFn0(p0 int8) bool {
    v7 := []int{-925, -9223372036854775808, 2147483647, 1}
    fmt.Println(p0)
    var v8 int16 = 32767
    _ = p0
    _ = v7
    _ = v8
    return (1 > 65535)
}
func genFn1(p0 uint64, p1 uint8) uint8 {
    v9 := GenS1{F0: -1, F1: p0, F2: -1}
    v10 := []int{1, 1}
    var v11 int16 = -244
    v11 -= (v11 + (v11 | -1))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return p1
}

func main() {
    v1 := []int{4294967295, -1, -9223372036854775808}
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]uint32)
    v5 := []uint32{2147483647, 4294967295, 65535, 204}
    v6 := []string{}
    var v12 bool = (1 == 661)
    var v13 uint = 395
    v14 := GenS2{F0: 287, F1: "a"}
    v13 += v13
    v12 = (-289 >= -2147483648)
    v15 := GenS1{F0: -2147483648, F1: 4294967295, F2: -1}
    v13 -= v13
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 uint16 = 0
        v12 = true
        _ = fi16
        _ = v17
    }
    switch ("hello" + "") {
    case "\n\t":
        switch ((("\n\t" + "\n\t") + "") + ("" + "foo bar")) {
        case "hello":
            fmt.Println(v13, v12)
        }
    case "foo bar":
        if ((v12 && (-1 < 9223372036854775807)) && v12) {
            switch "a" {
            case "a":
                v18 := GenS1{F0: 635, F1: 924, F2: 0}
                _ = v18
            }
        } else {
            v12 = v12
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 int32 = 1
                _ = fi19
                _ = v20
            }
        }
    case "a":
        switch -1 {
        case 2147483647:
            for fi21 := 0; fi21 < 14; fi21++ {
                v22 := GenS2{F0: 18446744073709551615, F1: "\n\t"}
                if (!(!false)) {
                    var v23 uint16 = 42
                    var v24 bool = (!(!((0 != 1) && (true && (-2 == 42)))))
                    if true {
                        if v24 {
                            if (!(2147483648 < 165)) {
                                v25 := GenS2{F0: 667, F1: ("foo bar" + (("\n\t" + "hello") + "\n\t"))}
                                fmt.Println(v23)
                                _ = v25
                            }
                            if v12 {
                                var v26 int8 = 100
                                fmt.Println(v13, v13, v24)
                                _ = v26
                            }
                            if ((v23 << (((v13 & v13)) & 63)) < v23) {
                                v13 = 4294967295
                            }
                        }
                    } else {
                        v23 = (v23 - v23)
                        var v27 float64 = 0.1
                        _ = v27
                    }
                    _ = v23
                    _ = v24
                }
                v13 = v13
                _ = fi21
                _ = v22
            }
        case 32768:
            fmt.Println(v13, v12)
        case -2147483648:
            v28 := GenS1{F0: 795, F1: 2, F2: 100}
            _ = v28
        default:
            var v29 float32 = 1.5
            _ = v29
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
