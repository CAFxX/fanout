package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int, p1 int64, p2 uint64) int8 {
    v6 := []int{-885}
    v6 = append(v6, p0)
    if len(v6) > 0 {
        v6[0] = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 108
}
func genFn1(p0 bool, p1 uint64, p2 uint16) int16 {
    p0 = p0
    var v7 int64 = 675
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -32768
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    v5 := []int{}
    var v8 bool = true
    v8 = (104 != 32768)
    v9 := GenS1{F0: -252, F1: 2147483648}
    if (!true) {
        if v8 {
            v8 = v8
        } else {
            if true {
                fmt.Println(v8)
                var v10 string = "a"
                _ = v10
            }
        }
        switch ("\n\t" + "a") {
        case "\n\t":
            switch -2147483648 {
            case -9223372036854775808:
                v8 = (-2147483648 < -2147483648)
            }
        case "":
            v11 := GenS2{F0: 65535, F1: 739}
            _ = v11
        case "foo bar":
            fmt.Println(v8)
        default:
            var v12 int64 = 916
            _ = v12
        }
        var v13 int = 32767
        _ = v13
    } else {
        switch "foo bar" {
        case "a":
            v8 = v8
        }
        v14 := v4[2]
        delete(v3, "a")
        _ = v14
    }
    v8 = (v8 || false)
    v8 = (!true)
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            if (!true) {
                v5 = append(v5, fi15)
                var v17 float64 = 1e10
                var v18 int = 2147483647
                _ = v17
                _ = v18
            } else {
                if true {
                    v8 = v8
                    v8 = ("" < ("\n\t" + "a"))
                    v8 = (v8 && true)
                }
            }
            _ = fi16
        }
        _ = fi15
    }
    v8 = (30 == 18446744073709551615)
    for fi19 := 0; fi19 < 15; fi19++ {
        v8 = false
        v5 = append(v5, -1)
        v8 = (-32768 >= 32767)
        _ = fi19
    }
    if (42 != 2) {
        var v20 int8 = -128
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := GenS1{F0: -9223372036854775808, F1: 2147483647}
        switch fi21 {
        case 32768:
            var v23 uint8 = 1
            _ = v23
        }
        fmt.Println(v8)
        _ = fi21
        _ = v22
    }
    var v24 uint32 = 1
    var v25 uint32 = v24
    var v26 uint8 = 150
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
