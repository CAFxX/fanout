package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 int64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint64) int16 {
    var v7 uint64 = p0
    v8 := GenS1{F0: false, F1: 57, F2: 779}
    switch -779 {
    case -1:
        p0 = (v7 * v7)
    case 1:
        p0 = v7
    case -849:
        var v9 uint16 = 42
        _ = v9
    }
    _ = p0
    _ = v7
    _ = v8
    return 1000
}
func genFn1(p0 bool, p1 uint64, p2 int32) string {
    v10 := GenS2{F0: (false || (-103 == 1)), F1: 633, F2: 9223372036854775807}
    if (!p0) {
        for fi11 := 0; fi11 < 18; fi11++ {
            p0 = p0
            for fi12 := 0; fi12 < 11; fi12++ {
                if p0 {
                    p0 = (false && p0)
                }
                _ = fi12
            }
            _ = fi11
        }
        var v13 int16 = -32768
        _ = v13
    } else {
        var v14 bool = p0
        switch 220 {
        case -9223372036854775808:
            v15 := []int{900, -1}
            _ = v15
        }
        p0 = (v14 && ((p2 ^ p2) == (p2 + p2)))
        _ = v14
    }
    if p0 {
        var v16 int32 = p2
        v17 := []int{921, 4294967295}
        _ = v16
        _ = v17
    } else {
        p2 = (p2 >> ((2) & 63))
        p0 = (-128 < 100)
    }
    v18 := []int{127, -9223372036854775808, -9223372036854775808}
    switch 1 {
    case -885:
        switch "\n\t" {
        case "foo bar":
            var v19 int16 = 133
            _ = v19
        default:
            v20 := GenS1{F0: (p1 >= p1), F1: 127, F2: p1}
            _ = v20
        }
    case -2147483648:
        p0 = (!(42 != 210))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v18
    return "foo bar"
}

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 2
    var v3 bool = true
    var v4 uint32 = 2147483647
    var v5 uint8 = (v2 << ((612) & 63))
    v6 := make(map[uint32]int)
    v21 := GenS3{F0: 0}
    v1 = 4294967295
    var v22 uint8 = v2
    v23 := []int{42, -246}
    v22 = (v2 % (((v2 * (v22 * 249))) | 1))
    if ((v1 >> ((1) & 63)) == v1) {
        for fi24 := 0; fi24 < 8; fi24++ {
            fmt.Println(v22, fi24, v3)
            if false {
                for fi25 := 0; fi25 < 6; fi25++ {
                    v4 = v4
                    _ = fi25
                }
                switch "a" {
                case "a":
                    var v26 int = (fi24 ^ (^(fi24)))
                    _ = v26
                }
                fmt.Println(v22)
            }
            var v27 uint16 = 65535
            _ = fi24
            _ = v27
        }
        fmt.Println(v5)
    }
    v2 += 1
    for fi28 := 0; fi28 < 4; fi28++ {
        switch "foo bar" {
        case "":
            var v29 string = ""
            _ = v29
        }
        v2 = (v22 - v2)
        _ = fi28
    }
    fmt.Println(v4, v22, v22)
    v30 := GenS3{F0: 0}
    var v31 bool = false
    for fi32 := 0; fi32 < 6; fi32++ {
        var v33 string = ("" + ("hello" + "hello"))
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
