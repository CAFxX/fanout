package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int16 }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
func genFn0(p0 int8, p1 int64) int {
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 uint8 = 42
        var v6 bool = true
        _ = fi4
        _ = v5
        _ = v6
    }
    p0 = (p0 * p0)
    var v7 bool = true
    p0 -= (p0 + (p0 & -128))
    var v8 uint = 927
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 uint32, p1 bool) int32 {
    if (63 != 42) {
        var v9 int64 = 324
        _ = v9
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            if ("foo bar" <= "hello") {
                v11 := []int{fi10, (^(fi10)), (fi10 % ((fi10) | 1))}
                p0 -= 1
                _ = v11
            } else {
                var v12 int8 = 24
                for fi13 := 0; fi13 < 13; fi13++ {
                    v14 := []int{(fi10 - (fi13 >> ((364) & 63))), -2147483648, fi13, (fi10 * 32767)}
                    _ = fi13
                    _ = v14
                }
                for fi15 := 0; fi15 < 14; fi15++ {
                    var v16 uint32 = p0
                    var v17 int64 = 0
                    p1 = false
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = v12
            }
            switch fi10 {
            case 42:
                fmt.Println(p1)
            case 850:
                switch "a" {
                case "\n\t":
                    for fi18 := 0; fi18 < 15; fi18++ {
                        v19 := []int{}
                        _ = fi18
                        _ = v19
                    }
                case "foo bar":
                    var v20 float32 = 0.1
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 2; fi21++ {
                        var v22 uint32 = 0
                        v23 := GenS1{F0: 801, F1: 18446744073709551615, F2: 32767}
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                }
            }
            _ = fi10
        }
    }
    p1 = true
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    var v1 uint = 35
    var v2 int16 = -819
    v3 := make(map[string]uint32)
    var v24 int = -1
    for fi25 := 0; fi25 < 14; fi25++ {
        if (!true) {
            v26 := GenS2{F0: -1, F1: 100, F2: 911}
            v27 := []int{fi25, (fi25 - -9223372036854775808), (-(fi25)), (v24 ^ -404)}
            v24 = v24
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            v29 := []int{(v24 * 0), -2147483648, fi28, (fi25 + int(v2))}
            _ = fi28
            _ = v29
        }
        _ = fi25
    }
    var v30 int8 = -108
    if ((v24 << ((4294967295) & 63)) >= v24) {
        v31 := GenS1{F0: 18446744073709551615, F1: v1, F2: 42}
        var v32 uint8 = 0
        if (!false) {
            var v33 uint8 = 67
            v34 := GenS1{F0: 1, F1: v1, F2: (-(v2))}
            _ = v33
            _ = v34
        }
        _ = v31
        _ = v32
    }
    var v35 uint8 = 0
    var v36 uint16 = 0
    v24 = (v24 ^ (v24 | (v24 + 9223372036854775807)))
    v37 := GenS2{F0: (v30 % (((v30 * (v30 * int8(v2)))) | 1)), F1: (-(v30)), F2: (v36 % (((v36 >> ((v1) & 63))) | 1))}
    switch "a" {
    case "":
        if (0 == 1) {
            for fi38 := 0; fi38 < 7; fi38++ {
                fmt.Println(v24)
                v3["foo bar"] = 65535
                for fi39 := 0; fi39 < 14; fi39++ {
                    v35 = 161
                    fmt.Println(fi38, v35)
                    var v40 uint8 = v35
                    _ = fi39
                    _ = v40
                }
                _ = fi38
            }
        }
    case "\n\t":
        v41 := v3["foo bar"]
        _ = v41
    default:
        v2 = v2
    }
    fmt.Println(v35, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
}
