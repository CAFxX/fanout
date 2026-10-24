package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    var v1 uint = 629
    v2 := make(map[int]string)
    var v3 int = -9223372036854775808
    var v4 uint16 = 42
    if (42 == 42) {
        for fi5 := 0; fi5 < 19; fi5++ {
            fmt.Println(fi5, v1)
            v1 -= v1
            v3 -= v3
            _ = fi5
        }
        v6 := GenS1{F0: 2147483647, F1: 42, F2: v4}
        _ = v6
    } else {
        fmt.Println(v1)
        var v7 bool = (!false)
        for fi8 := 0; fi8 < 20; fi8++ {
            v1 += v1
            _ = len(v2)
            _ = fi8
        }
        _ = v7
    }
    var v9 uint16 = (v4 * (v4 >> (((v1 % ((v1) | 1))) & 63)))
    if ((-91 <= -58) || (v1 >= v1)) {
        var v10 uint = (v1 / ((v1) | 1))
        var v11 int32 = -2147483648
        _ = v10
        _ = v11
    }
    var v12 uint16 = v9
    if (726 != v12) {
        if false {
            if true {
                var v13 uint8 = 0
                _ = v13
            } else {
                var v14 uint64 = 0
                v4 -= uint16(v14)
                for fi15 := 0; fi15 < 1; fi15++ {
                    fmt.Println(v9, v9, v4)
                    v16 := GenS2{F0: "hello", F1: -1}
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            switch "foo bar" {
            case "a":
                for fi17 := 0; fi17 < 2; fi17++ {
                    v18 := []int{(-(v3)), (fi17 - (fi17 & -1))}
                    var v19 uint32 = 1
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            default:
                v9 += v9
            }
            v12 -= v9
        } else {
            v20 := []int{}
            v1 -= 65535
            _ = len(v20)
            _ = v20
        }
        if (!true) {
            switch ("\n\t" + "hello") {
            case "a":
                v21 := v2[v3]
                _ = v21
            case "\n\t":
                var v22 bool = false
                _ = v22
            case "foo bar":
                fmt.Println(v3, v12)
            }
            v23 := GenS1{F0: 615, F1: -1, F2: v9}
            switch -2147483648 {
            case 826:
                var v24 float32 = 0.0
                _ = v24
            }
            _ = v23
        } else {
            v4 = (v9 % ((2) | 1))
            var v25 int64 = 476
            if true {
                var v26 string = ("" + ("\n\t" + "\n\t"))
                v25 -= (v25 - (-(v25)))
                for fi27 := 0; fi27 < 15; fi27++ {
                    v28 := []int{fi27, (v3 & v3), 253, fi27}
                    var v29 string = v26
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v26
            }
            _ = v25
        }
        v30 := GenS2{F0: "", F1: 1}
        _ = v30
    } else {
        switch "" {
        case "a":
            if (false && false) {
                var v31 uint = 18446744073709551615
                _ = v31
            } else {
                v4 = v4
                v3 -= v3
                for fi32 := 0; fi32 < 3; fi32++ {
                    v33 := GenS1{F0: 1, F1: -1, F2: 414}
                    _ = fi32
                    _ = v33
                }
            }
        case "\n\t":
            v34 := []int{(v3 >> ((65535) & 63)), v3}
            _ = v34
        case "hello":
            v12 -= v9
        default:
            if (!((v4 ^ v4) > (v9 << (((v1 + 2)) & 63)))) {
                var v35 uint32 = 2147483647
                var v36 string = ""
                _ = v35
                _ = v36
            }
        }
    }
    fmt.Println(v12, v1, v3)
    for fi37 := 0; fi37 < 2; fi37++ {
        var v38 int32 = -1
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        v2[v3] = "foo bar"
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
