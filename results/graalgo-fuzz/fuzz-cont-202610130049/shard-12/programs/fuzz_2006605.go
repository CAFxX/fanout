package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
func genFn0() uint {
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(fi5)
        var v6 int32 = -1
        if true {
            var v7 int32 = (v6 * 344)
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint16 = 65535
        if (!false) {
            var v10 float64 = 0.1
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 int32 = 935
                v13 := []int{2147483647, fi8, int(v9)}
                var v14 bool = false
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            _ = v10
        }
        switch fi8 {
        case 2147483647:
            if true {
                var v15 uint32 = 277
                v16 := []int{fi8, fi8, (fi8 / (((fi8 ^ (^(fi8)))) | 1)), (fi8 / ((-423) | 1))}
                var v17 uint64 = 18446744073709551615
                _ = v15
                _ = v16
                _ = v17
            } else {
                v9 -= (v9 | v9)
                v18 := []int{(fi8 >> ((255) & 63)), (-(fi8)), fi8, -427}
                _ = v18
            }
        default:
            v19 := []int{(fi8 * fi8), (-(fi8)), -1}
            _ = v19
        }
        _ = fi8
        _ = v9
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        if (false || (2 > 63)) {
            var v21 int64 = 0
            fmt.Println(v21)
            v21 = 9223372036854775807
            _ = v21
        }
        v22 := []int{(fi20 >> ((255) & 63)), fi20, fi20}
        var v23 int16 = 642
        _ = fi20
        _ = v22
        _ = v23
    }
    if (0.0 >= (0.1 / 3.14159)) {
        if true {
            for fi24 := 0; fi24 < 5; fi24++ {
                var v25 uint16 = 255
                v26 := GenS1{F0: ""}
                _ = fi24
                _ = v25
                _ = v26
            }
            v27 := []int{}
            _ = v27
        }
    }
    var v28 uint8 = 30
    _ = v28
    return 91
}

func main() {
    var v1 int32 = -1
    v2 := []uint32{740, 680}
    v3 := make(map[uint32]int)
    v4 := []int{2147483647, -1}
    v29 := GenS1{F0: "foo bar"}
    v1 = (v1 * (^(v1)))
    var v30 int = -9223372036854775808
    var v31 uint = 18446744073709551615
    v1 = v1
    if false {
        v31 = genFn0()
        for fi32 := 0; fi32 < 4; fi32++ {
            if true {
                var v33 uint64 = 1
                _ = v33
            }
            if true {
                fmt.Println(fi32, v1, v30)
            } else {
                v34 := GenS1{F0: "\n\t"}
                switch "" {
                case "hello":
                    fmt.Println(v1, v31)
                case "\n\t":
                    var v35 uint32 = 2
                    _ = v35
                default:
                    for fi36 := 0; fi36 < 6; fi36++ {
                        v1 -= (v1 % ((v1) | 1))
                        v30 -= 32768
                        _ = len(v4)
                        _ = fi36
                    }
                }
                _ = v34
            }
            _ = fi32
        }
        for fi37 := 0; fi37 < 17; fi37++ {
            v1 += 2147483647
            _ = fi37
        }
    }
    v1 -= 42
    v31 += (v31 | (v31 << ((255) & 63)))
    var v38 uint = v31
    switch v30 {
    case 128:
        var v39 uint8 = 0
        _ = v39
    case 226:
        switch ("hello" + "hello") {
        case "":
            fmt.Println(v30, v31, v38)
        case "foo bar":
            v30 -= 128
        default:
            var v40 uint8 = 255
            _ = v40
        }
    case -786:
        switch "foo bar" {
        case "foo bar":
            v41 := GenS2{F0: -32768}
            _ = v41
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "hello":
            v1 += (^(v1))
        }
    default:
        v42 := GenS3{F0: 42}
        _ = v42
    }
    var v43 int = v30
    var v44 uint = 42
    fmt.Println(v44)
    fmt.Println(v38, v44)
    var v45 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
