package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
func genFn0(p0 int16) uint32 {
    p0 = p0
    if (726 != (p0 << ((42) & 63))) {
        v7 := GenS1{F0: 42, F1: -1}
        _ = v7
    } else {
        p0 = -1
    }
    _ = p0
    return 0
}
func genFn1(p0 int16, p1 uint8) int64 {
    v8 := GenS1{F0: 4294967295, F1: 42}
    p1 += (p1 % (((p1 >> ((1) & 63))) | 1))
    _ = p0
    _ = p1
    _ = v8
    return -1
}

func main() {
    var v1 float64 = 3.14159
    v2 := []int{-545, -9223372036854775808, 127}
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]string)
    var v6 int8 = 33
    v9 := GenS2{F0: 0}
    v1 = v1
    switch ("hello" + ("\n\t" + "\n\t")) {
    case "\n\t":
        if false {
            var v10 float64 = v1
            _ = v10
        } else {
            for fi11 := 0; fi11 < 7; fi11++ {
                switch fi11 {
                case 42:
                    v12 := GenS1{F0: 4294967295, F1: 0}
                    _ = v12
                case 128:
                    for fi13 := 0; fi13 < 15; fi13++ {
                        var v14 int16 = -739
                        v5[2147483648] = (("" + "foo bar") + "")
                        _ = fi13
                        _ = v14
                    }
                case -2147483648:
                    switch (fi11 << ((255) & 63)) {
                    case -9223372036854775808:
                        var v15 uint = 943
                        _ = v15
                    case 4294967295:
                        var v16 string = "\n\t"
                        _ = v16
                    case 32768:
                        var v17 uint16 = 2
                        _ = v17
                    }
                default:
                    v1 = 1e10
                }
                _ = fi11
            }
            switch "\n\t" {
            case "hello":
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            case "\n\t":
                var v18 uint8 = 1
                _ = v18
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case "a":
        var v19 uint64 = 18446744073709551615
        _ = v19
    }
    if (!false) {
        switch "foo bar" {
        case "foo bar":
            v6 += (v6 - (v6 & (v6 >> ((42) & 63))))
        default:
            v6 -= (v6 >> ((227) & 63))
        }
        var v20 int16 = -32768
        _ = v20
    }
    if false {
        v1 += v1
        switch 754 {
        case 32767:
            var v21 int8 = 0
            _ = v21
        case 2147483647:
            if (false && (-9223372036854775808 > -9223372036854775808)) {
                var v22 int = -1
                var v23 string = ("\n\t" + ("foo bar" + ("a" + ("hello" + "hello"))))
                _ = v22
                _ = v23
            } else {
                v24 := GenS1{F0: 549, F1: v6}
                v25 := GenS1{F0: 65535, F1: v6}
                fmt.Println(v6, v6)
                _ = v24
                _ = v25
            }
        default:
            if (true || true) {
                var v26 uint = 122
                _ = v26
            }
        }
    }
    v6 = v6
    if (-9223372036854775808 == -1) {
        switch "foo bar" {
        case "\n\t":
            var v27 int64 = -1
            _ = v27
        }
        v1 -= (v1 + ((1e10 * 3.14159) / 1.5))
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 string = "a"
            v29 = v29
            _ = fi28
            _ = v29
        }
    } else {
        if (4294967295 >= 2) {
            var v30 uint64 = 18446744073709551615
            _ = v30
        } else {
            var v31 int8 = v6
            switch ("hello" + ("a" + (("foo bar" + "hello") + "foo bar"))) {
            case "foo bar":
                switch ("" + "foo bar") {
                case "a":
                    v1 = ((v1 * v1) + float64(3.14159))
                case "":
                    var v32 bool = false
                    _ = v32
                default:
                    var v33 int16 = -989
                    _ = v33
                }
            case "\n\t":
                _ = len(v3)
            }
            if len(v2) > 0 {
                v2[0] = 9223372036854775807
            }
            _ = v31
        }
        if ("\n\t" >= ("foo bar" + "hello")) {
            v34 := GenS3{F0: 42, F1: 248, F2: 255}
            v35 := GenS1{F0: 255, F1: 42}
            _ = v34
            _ = v35
        }
    }
    v1 = (3.14159 / v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
