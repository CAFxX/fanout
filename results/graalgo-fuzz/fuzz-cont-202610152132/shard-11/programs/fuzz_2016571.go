package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int; F2 uint16 }
type GenS3 struct { F0 int8; F1 int64; F2 uint16 }
func genFn0(p0 bool) bool {
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    if (false || (("a" + "hello") <= "hello")) {
        if p0 {
            var v7 string = "a"
            var v8 int32 = -311
            var v9 string = v7
            _ = v7
            _ = v8
            _ = v9
        }
        if (!false) {
            v10 := GenS1{F0: 0}
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := []int{fi11}
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = fi11
            _ = v12
        }
    }
    _ = p0
    return (!(6 < -32768))
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    var v5 int = 4
    var v13 uint64 = v2
    v5 -= (v5 % ((v5) | 1))
    switch "foo bar" {
    case "foo bar":
        if (false && true) {
            var v14 uint64 = v13
            _ = v14
        }
    default:
        var v15 int8 = -1
        _ = v15
    }
    var v16 uint32 = (v1 + (v1 * v1))
    var v17 uint8 = 128
    v16 = 2147483648
    if false {
        v18 := []int{int(v17), (v5 & v5)}
        if (!true) {
            if false {
                if true {
                    var v19 uint32 = v1
                    v1 -= (v1 & v16)
                    _ = v19
                }
                v2 = (v2 + (v2 & 427))
                v16 = uint32(v17)
            }
            switch 9223372036854775807 {
            case -62:
                v20 := GenS2{F0: uint64(v17), F1: v5, F2: 0}
                _ = v20
            case 2147483647:
                v21 := GenS2{F0: (v13 | v13), F1: 32767, F2: 804}
                _ = v21
            }
        } else {
            for fi22 := 0; fi22 < 8; fi22++ {
                fmt.Println(v1, v16, v1)
                v23 := GenS2{F0: 0, F1: (v5 + (fi22 << ((46) & 63))), F2: 2}
                _ = fi22
                _ = v23
            }
            if true {
                v2 = v2
            } else {
                v24 := GenS1{F0: 1000}
                var v25 uint16 = 1
                _ = v24
                _ = v25
            }
            fmt.Println(v2, v16)
        }
        v18 = append(v18, v5)
        _ = v18
    }
    var v26 string = "foo bar"
    for fi27 := 0; fi27 < 2; fi27++ {
        v28 := []int{fi27}
        v29 := GenS3{F0: 0, F1: -375, F2: 40}
        v28 = append(v28, (fi27 / (((fi27 / ((fi27) | 1))) | 1)))
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println(v16, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
