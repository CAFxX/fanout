package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int64 }
type GenS2 struct { F0 int8; F1 string; F2 int }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint8) uint8 {
    v6 := GenS1{F0: 42, F1: 18446744073709551615, F2: 9223372036854775807}
    if (!false) {
        if false {
            fmt.Println(p0)
            for fi7 := 0; fi7 < 17; fi7++ {
                fmt.Println(p0, fi7)
                _ = fi7
            }
            for fi8 := 0; fi8 < 10; fi8++ {
                if false {
                    switch "hello" {
                    case "a":
                        for fi9 := 0; fi9 < 20; fi9++ {
                            p0 = (p0 - 1)
                            _ = fi9
                        }
                    case "\n\t":
                        v10 := GenS2{F0: 47, F1: ("" + "a"), F2: fi8}
                        _ = v10
                    case "hello":
                        v11 := GenS2{F0: 100, F1: ("hello" + "\n\t"), F2: -2}
                        _ = v11
                    }
                    var v12 int32 = -1
                    _ = v12
                }
                for fi13 := 0; fi13 < 11; fi13++ {
                    v14 := []int{128, (fi13 * fi13), fi8, (^(fi8))}
                    p0 = p0
                    _ = fi13
                    _ = v14
                }
                _ = fi8
            }
        } else {
            var v15 uint16 = 182
            v15 = v15
            if (!false) {
                for fi16 := 0; fi16 < 8; fi16++ {
                    fmt.Println(v15, p0)
                    var v17 float64 = (0.0 - 1.5)
                    _ = fi16
                    _ = v17
                }
                fmt.Println(p0)
            }
            _ = v15
        }
        p0 += p0
        var v18 int8 = -1
        _ = v18
    }
    if true {
        var v19 uint8 = (^(p0))
        if (!(!true)) {
            v19 = v19
        } else {
            if (1 != -1) {
                p0 = (v19 << ((0) & 63))
                p0 = (^(p0))
            }
            v19 = (p0 % (((p0 ^ p0)) | 1))
        }
        _ = v19
    } else {
        switch 1 {
        case 0:
            switch 1 {
            case -518:
                var v20 uint64 = 0
                _ = v20
            case 4294967295:
                p0 -= 127
            case 9223372036854775807:
                if (!true) {
                    v21 := GenS1{F0: -217, F1: 18446744073709551615, F2: 706}
                    _ = v21
                }
            }
        case -2147483648:
            var v22 int8 = -49
            _ = v22
        }
        if true {
            v23 := GenS1{F0: 1, F1: 65535, F2: -9223372036854775808}
            v24 := GenS2{F0: 0, F1: ("\n\t" + ("a" + ("\n\t" + "\n\t"))), F2: 2147483647}
            _ = v23
            _ = v24
        } else {
            p0 = 1
            v25 := GenS1{F0: 0, F1: 42, F2: -9223372036854775808}
            var v26 int32 = 987
            _ = v25
            _ = v26
        }
    }
    _ = p0
    _ = v6
    return (^(p0))
}
func genFn1(p0 uint8, p1 bool) string {
    for fi27 := 0; fi27 < 2; fi27++ {
        v28 := GenS2{F0: -33, F1: "a", F2: (fi27 / ((32768) | 1))}
        _ = fi27
        _ = v28
    }
    p1 = p1
    _ = p0
    _ = p1
    return "a"
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 626
    var v3 uint = 2
    var v4 bool = false
    var v5 uint32 = 2147483647
    v29 := GenS2{F0: 0, F1: "foo bar", F2: 2147483647}
    if v4 {
        fmt.Println(v4, v2)
        v30 := []int{}
        _ = v30
    }
    v5 += uint32(v3)
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 uint8 = 127
        _ = len(v1)
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        var v34 string = "hello"
        fmt.Println(v2)
        _ = fi33
        _ = v34
    }
    var v35 uint32 = v5
    _ = len(v1)
    if v4 {
        if v4 {
            _ = len(v1)
            v36 := GenS3{F0: 42, F1: 611}
            _ = v36
        }
        fmt.Println(v35, v5, v5)
        switch -1 {
        case 32768:
            var v37 int = -707
            _ = v37
        default:
            v38 := GenS2{F0: 0, F1: "foo bar", F2: -2147483648}
            _ = v38
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
