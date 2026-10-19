package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint16 }
func genFn0(p0 string) int {
    switch 662 {
    case -2147483648:
        var v7 int8 = 105
        _ = v7
    case 128:
        var v8 uint64 = 18446744073709551615
        _ = v8
    case 189:
        if true {
            if false {
                var v9 uint32 = 4294967295
                if (true || (-9223372036854775808 <= 32768)) {
                    var v10 int8 = -128
                    v10 -= -128
                    _ = v10
                }
                _ = v9
            }
            v11 := []int{4294967295, 42, 32767}
            _ = v11
        } else {
            fmt.Println(p0)
            p0 = p0
            v12 := []int{148, 128}
            _ = v12
        }
    }
    var v13 uint64 = 0
    switch 127 {
    case 127:
        var v14 int64 = -284
        _ = v14
    case -2147483648:
        for fi15 := 0; fi15 < 16; fi15++ {
            for fi16 := 0; fi16 < 5; fi16++ {
                switch "foo bar" {
                case "a":
                    p0 = p0
                case "":
                    fmt.Println(fi16, p0, fi16)
                default:
                    v17 := GenS1{F0: (!true), F1: 1}
                    _ = v17
                }
                if true {
                    var v18 int64 = 42
                    _ = v18
                } else {
                    fmt.Println(p0, fi15, fi15)
                }
                _ = fi16
            }
            var v19 bool = true
            switch "a" {
            case "foo bar":
                for fi20 := 0; fi20 < 10; fi20++ {
                    v21 := GenS2{F0: v19}
                    switch fi15 {
                    case -875:
                        p0 = p0
                    default:
                        var v22 uint8 = 42
                        _ = v22
                    }
                    v19 = v19
                    _ = fi20
                    _ = v21
                }
            case "\n\t":
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 int64 = 1
                    _ = fi23
                    _ = v24
                }
            default:
                v25 := GenS3{F0: (fi15 & (fi15 | (fi15 * (fi15 % (((-(fi15))) | 1))))), F1: uint16(v13)}
                _ = v25
            }
            _ = fi15
            _ = v19
        }
    }
    _ = p0
    _ = v13
    return 9223372036854775807
}
func genFn1(p0 int8) int8 {
    var v26 int = genFn0("")
    v27 := GenS1{F0: true, F1: 1}
    fmt.Println(v26, v26)
    _ = p0
    _ = v26
    _ = v27
    return p0
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint64 = 42
    var v3 uint16 = 32767
    v4 := make(map[int]uint32)
    var v5 int64 = 42
    var v6 int64 = v5
    v3 = v3
    for fi28 := 0; fi28 < 1; fi28++ {
        v2 += (v2 << ((65535) & 63))
        var v29 uint32 = 2147483647
        v30 := GenS3{F0: fi28, F1: v3}
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        v1 += 0.0
        for fi32 := 0; fi32 < 18; fi32++ {
            v2 += v2
            _ = fi32
        }
        _ = fi31
    }
    switch 127 {
    case 127:
        delete(v4, 127)
    case 2147483647:
        var v33 string = "\n\t"
        _ = v33
    case 1:
        if (4294967295 <= v2) {
            var v34 int16 = 32767
            _ = v34
        }
    default:
        v1 += v1
    }
    var v35 int16 = 1000
    var v36 int8 = 26
    var v37 float64 = (1.5 + (0.1 + 3.14159))
    v1 -= 1.5
    var v38 int8 = (v36 ^ (v36 + genFn1((v36 / ((100) | 1)))))
    v38 -= (v36 + v38)
    for fi39 := 0; fi39 < 14; fi39++ {
        for fi40 := 0; fi40 < 12; fi40++ {
            var v41 float64 = 0.0
            _ = fi40
            _ = v41
        }
        v35 += -1
        _ = fi39
    }
    var v42 float32 = -2.25
    fmt.Println(v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
