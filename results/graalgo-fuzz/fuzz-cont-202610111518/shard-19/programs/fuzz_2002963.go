package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint8) int {
    v5 := GenS1{F0: -1}
    p0 -= 42
    _ = p0
    _ = p1
    _ = v5
    return -1
}
func genFn1() string {
    var v6 uint16 = 963
    v6 -= 65535
    fmt.Println(v6)
    _ = v6
    return ("hello" + "foo bar")
}

func main() {
    var v1 int = -370
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    var v4 uint = 65535
    fmt.Println(v4, v1)
    delete(v2, (v1 & 128))
    v4 = v4
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 int16 = 42
            _ = fi8
            _ = v9
        }
        v4 = v4
        if ((!false) && (51 <= 42)) {
            v10 := v3[""]
            _ = v10
        } else {
            v11 := GenS1{F0: -2147483648}
            _ = v11
        }
        _ = fi7
    }
    v12 := GenS1{F0: 955}
    delete(v3, "")
    if false {
        if (237 > 1) {
            var v13 uint = 562
            _ = v13
        }
        v4 += v4
    } else {
        v2[2147483647] = 65535
        var v14 int = (v1 + (v1 / ((v1) | 1)))
        _ = v14
    }
    var v15 uint = (v4 | v4)
    for fi16 := 0; fi16 < 16; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            v1 = (v1 * 68)
            _ = fi17
        }
        _ = fi16
    }
    if false {
        v15 = (v4 ^ v15)
        v1 = (v1 % (((v1 * int(v4))) | 1))
        var v18 uint = 788
        _ = v18
    } else {
        v1 += (-(v1))
        v4 -= v4
        for fi19 := 0; fi19 < 3; fi19++ {
            switch "foo bar" {
            case "foo bar":
                var v20 uint16 = 42
                _ = v20
            case "a":
                var v21 uint32 = 2
                _ = v21
            case "":
                v4 = (^(v15))
            default:
                var v22 float64 = -2.25
                _ = v22
            }
            for fi23 := 0; fi23 < 12; fi23++ {
                v3["hello"] = 821
                var v24 uint64 = 1
                switch v1 {
                case 2147483647:
                    var v25 uint32 = 561
                    _ = v25
                }
                _ = fi23
                _ = v24
            }
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 int16 = int16(v4)
                var v28 uint = v4
                fmt.Println(v28, fi26, v4)
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi19
        }
    }
    v29 := GenS2{F0: -32768}
    delete(v3, "hello")
    v1 = (v1 | (v1 * v1))
    if true {
        var v30 int16 = -1
        _ = v30
    }
    if true {
        v1 += 9223372036854775807
        var v31 int = genFn0(65535, 0)
        v31 += (v1 / ((v1) | 1))
        _ = v31
    } else {
        fmt.Println(v15)
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        for fi33 := 0; fi33 < 9; fi33++ {
            if (!(false || (18446744073709551615 < 601))) {
                v34 := []int{32768, fi33}
                v1 += (fi32 ^ 2147483647)
                var v35 uint64 = 1
                _ = v34
                _ = v35
            } else {
                v4 = (v4 << (((^(v4))) & 63))
            }
            _ = fi33
        }
        v15 = v4
        _ = fi32
    }
    v36 := GenS1{F0: -1}
    delete(v3, "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v36
    _ = v4
}
