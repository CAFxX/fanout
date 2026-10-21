package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint32; F2 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 uint64; F2 bool }
func genFn0(p0 uint32, p1 string) bool {
    var v7 float64 = -2.25
    v8 := GenS1{F0: p1}
    v9 := []int{127, 4294967295}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (!false)
}
func genFn1(p0 string, p1 int32) uint16 {
    switch -197 {
    case 0:
        switch p0 {
        case "\n\t":
            if true {
                v10 := GenS2{F0: -896, F1: 0, F2: "\n\t"}
                fmt.Println(p1, p0)
                _ = v10
            } else {
                for fi11 := 0; fi11 < 17; fi11++ {
                    p0 = ((p0 + p0) + (p0 + "hello"))
                    fmt.Println(p0)
                    v12 := GenS1{F0: (p0 + "")}
                    _ = fi11
                    _ = v12
                }
            }
        default:
            var v13 uint64 = 18446744073709551615
            _ = v13
        }
    }
    var v14 uint = 42
    _ = p0
    _ = p1
    _ = v14
    return 1
}

func main() {
    var v1 int8 = -1
    var v2 uint = 42
    v3 := []int{128, -9223372036854775808, -2147483648, 579}
    var v4 float64 = 0.0
    var v5 uint16 = 2
    var v6 uint64 = 0
    if (uint32(v2) != 2147483648) {
        var v15 uint16 = v5
        v15 += (v15 % ((v15) | 1))
        v1 += 100
        _ = v15
    }
    switch 573 {
    case 32767:
        switch -557 {
        case 128:
            for fi16 := 0; fi16 < 11; fi16++ {
                if false {
                    if len(v3) > 0 {
                        v3[0] = (fi16 / (((fi16 % ((32767) | 1))) | 1))
                    }
                }
                _ = fi16
            }
        case 2147483647:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v5, v1)
    case "\n\t":
        for fi17 := 0; fi17 < 2; fi17++ {
            v4 -= (v4 / 1e10)
            v1 = v1
            _ = fi17
        }
    default:
        for fi18 := 0; fi18 < 17; fi18++ {
            for fi19 := 0; fi19 < 5; fi19++ {
                switch "a" {
                case "a":
                    fmt.Println(fi18, v5)
                default:
                    if false {
                        v4 = float64(((3.14159 - 0.1) / -2.25))
                        var v20 float32 = -2.25
                        v21 := GenS3{F0: 0}
                        _ = v20
                        _ = v21
                    } else {
                        switch (fi19 >> ((234) & 63)) {
                        case 32767:
                            v4 = 1e10
                        case -9223372036854775808:
                            var v22 uint8 = 190
                            _ = v22
                        }
                        var v23 int32 = 42
                        _ = v23
                    }
                }
                var v24 int32 = -2147483648
                if true {
                    v25 := GenS4{F0: 193, F1: uint64(v2), F2: false}
                    _ = v25
                }
                _ = fi19
                _ = v24
            }
            var v26 bool = false
            _ = fi18
            _ = v26
        }
    }
    var v27 bool = (!true)
    v5 = (v5 / ((v5) | 1))
    var v28 uint32 = 65535
    v27 = true
    var v29 int64 = 0
    switch 32767 {
    case 32767:
        var v30 bool = true
        _ = v30
    case 128:
        v2 = v2
    }
    v4 = (3.14159 + ((v4 + (-2.25 + v4)) * 0.0))
    var v31 float32 = (((1e10 - 0.1) * 1e10) + (0.1 - ((-2.25 + ((1.5 - 3.14159) + (((3.14159 - 1.5) * (0.0 - 1e10)) - float32((v4 - 1.5))))) * -2.25)))
    v1 = 127
    var v32 uint16 = uint16(v28)
    for fi33 := 0; fi33 < 9; fi33++ {
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 float32 = v31
            _ = fi34
            _ = v35
        }
        for fi36 := 0; fi36 < 5; fi36++ {
            v29 += -9223372036854775808
            v29 = (v29 >> (((v2 | (v2 * v2))) & 63))
            if ((("\n\t" != "hello") && (v6 != v6)) || (genFn1("hello", 0) > v32)) {
                if ((!v27) || v27) {
                    switch (fi33 >> ((2) & 63)) {
                    case 24:
                        _ = len(v3)
                    case 4294967295:
                        fmt.Println(v31, v28)
                    case 632:
                        fmt.Println(v6)
                        fmt.Println(fi33, v1)
                    default:
                        v27 = (v5 != v32)
                    }
                }
                var v37 int8 = (v1 & (v1 / ((42) | 1)))
                _ = v37
            }
            _ = fi36
        }
        if v27 {
            v5 = v32
            var v38 int16 = 0
            _ = v38
        } else {
            for fi39 := 0; fi39 < 7; fi39++ {
                var v40 uint64 = v6
                var v41 uint16 = v5
                v3 = append(v3, -1)
                _ = fi39
                _ = v40
                _ = v41
            }
            v2 -= v2
            var v42 bool = ((v6 - 220) == (v6 << (((v2 | v2)) & 63)))
            _ = v42
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
