package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 int }
func genFn0(p0 int16) string {
    if (34 <= 255) {
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 bool = (1e10 < (-2.25 * -2.25))
            v8 := GenS1{F0: 726}
            _ = fi6
            _ = v7
            _ = v8
        }
    } else {
        v9 := []int{421, 1}
        var v10 float64 = 0.1
        v11 := GenS2{F0: true, F1: 931}
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := []int{127, -494, 127}
    fmt.Println(p0)
    if (32767 < 255) {
        v13 := GenS1{F0: -902}
        if true {
            p0 = -32768
        }
        p0 = (p0 << ((374) & 63))
        _ = v13
    }
    v12 = append(v12, 1)
    _ = p0
    _ = v12
    return "a"
}
func genFn1() int8 {
    var v14 int16 = 42
    v15 := GenS3{F0: 128}
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 int32 = -1
    var v2 uint = 4294967295
    var v3 int16 = 1
    var v4 int = 1
    v5 := make(map[uint32]string)
    v2 = v2
    var v16 int = (v4 & v4)
    if (((true && true) || false) || (false && (false || (v16 >= 2147483647)))) {
        fmt.Println(v4)
        v5[980] = "hello"
        if (42 < 128) {
            v17 := GenS4{F0: -1, F1: 1}
            fmt.Println(v3, v4, v4)
            _ = v17
        } else {
            for fi18 := 0; fi18 < 3; fi18++ {
                for fi19 := 0; fi19 < 7; fi19++ {
                    v20 := GenS5{F0: (fi18 - (v4 << ((v2) & 63)))}
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
            var v21 uint64 = 1
            v16 = (v16 | (v4 % ((128) | 1)))
            _ = v21
        }
    }
    var v22 int64 = 42
    fmt.Println(v16, v22, v16)
    if false {
        fmt.Println(v16, v16)
    } else {
        switch (v4 >> ((v2) & 63)) {
        case 635:
            for fi23 := 0; fi23 < 20; fi23++ {
                v1 = 2147483647
                switch "a" {
                case "hello":
                    v2 += (^(v2))
                case "a":
                    v24 := []int{int(v22), v16}
                    _ = v24
                default:
                    var v25 uint64 = 731
                    _ = v25
                }
                _ = fi23
            }
        default:
            v26 := GenS6{F0: v3, F1: v22}
            _ = v26
        }
        switch "hello" {
        case "\n\t":
            v1 = 0
        case "foo bar":
            v27 := []int{v16, (v16 - 32768), v4, (v16 & (-(v4)))}
            _ = v27
        }
    }
    v1 = 42
    for fi28 := 0; fi28 < 13; fi28++ {
        for fi29 := 0; fi29 < 12; fi29++ {
            v30 := GenS5{F0: fi29}
            v4 -= v4
            _ = fi29
            _ = v30
        }
        if (128 < int(v3)) {
            fmt.Println(v22, v22)
            v31 := GenS3{F0: (v16 << (((v2 + (v2 ^ (v2 + (v2 ^ 42))))) & 63))}
            _ = v31
        }
        _ = fi28
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        v33 := GenS7{F0: v4}
        if (!(!false)) {
            var v34 float32 = 1.5
            if true {
                fmt.Println(v22, fi32)
            }
            v34 += ((float32(0.1) / v34) / v34)
            _ = v34
        } else {
            if (254 < 128) {
                var v35 bool = false
                for fi36 := 0; fi36 < 16; fi36++ {
                    var v37 uint16 = 2
                    var v38 uint16 = (^(v37))
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                for fi39 := 0; fi39 < 10; fi39++ {
                    v40 := []int{(v4 << ((42) & 63)), 4294967295, v4, fi39}
                    switch (-(fi39)) {
                    case -9223372036854775808:
                        v4 = 292
                    case 662:
                        var v41 int8 = -22
                        _ = v41
                    }
                    _ = fi39
                    _ = v40
                }
                _ = v35
            } else {
                var v42 int8 = genFn1()
                v43 := v5[949]
                v22 += (v22 % ((-6) | 1))
                _ = v42
                _ = v43
            }
            var v44 int = v4
            _ = v44
        }
        _ = fi32
        _ = v33
    }
    var v45 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
