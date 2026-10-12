package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
func genFn0() uint16 {
    if (!false) {
        var v6 int64 = -785
        v6 = -454
        for fi7 := 0; fi7 < 6; fi7++ {
            v6 = 0
            _ = fi7
        }
        _ = v6
    }
    var v8 uint32 = 1
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: -1}
        _ = fi9
        _ = v10
    }
    if false {
        if true {
            v11 := GenS1{F0: 42}
            v8 -= v8
            fmt.Println(v8)
            _ = v11
        } else {
            fmt.Println(v8)
            var v12 float64 = (1e10 - float64(0.1))
            switch -9223372036854775808 {
            case 32768:
                v8 += (v8 * v8)
            case 4294967295:
                fmt.Println(v12, v12)
            }
            _ = v12
        }
        v13 := []int{-2147483648}
        if false {
            if ((!true) || ("a" >= "")) {
                var v14 bool = false
                fmt.Println(v14, v14)
                _ = v14
            } else {
                switch "hello" {
                case "hello":
                    fmt.Println(v8)
                default:
                    v8 = (v8 * (v8 - v8))
                }
                var v15 uint = 916
                for fi16 := 0; fi16 < 20; fi16++ {
                    fmt.Println(v15, v8)
                    switch "\n\t" {
                    case "hello":
                        v15 -= v15
                    case "":
                        switch "hello" {
                        case "hello":
                            var v17 int64 = 653
                            _ = v17
                        }
                    }
                    for fi18 := 0; fi18 < 4; fi18++ {
                        var v19 int64 = -9223372036854775808
                        v19 -= (v19 << (((^(v15))) & 63))
                        _ = fi18
                        _ = v19
                    }
                    _ = fi16
                }
                _ = v15
            }
            v8 = v8
        }
        _ = v13
    } else {
        v20 := []int{}
        _ = v20
    }
    var v21 int16 = 0
    _ = v21
    _ = v8
    return 255
}

func main() {
    var v1 int16 = -245
    var v2 string = "hello"
    var v3 uint16 = 51
    v4 := make(map[string]string)
    var v5 string = (("a" + v2) + "hello")
    var v22 int32 = 2147483647
    v1 -= v1
    fmt.Println(v3, v2, v5)
    v23 := []int{-858, -9223372036854775808}
    for fi24 := 0; fi24 < 19; fi24++ {
        if false {
            if (!(-87 > 42)) {
                v25 := v4[v2]
                fmt.Println(v3)
                _ = len(v23)
                _ = v25
            }
        } else {
            for fi26 := 0; fi26 < 8; fi26++ {
                v5 = v2
                v2 = (v2 + (("" + v2) + ((v5 + v5) + "hello")))
                v27 := GenS1{F0: v22}
                _ = fi26
                _ = v27
            }
            v2 = "a"
        }
        fmt.Println(v2, v2, v2)
        _ = fi24
    }
    fmt.Println(v1, v22)
    var v28 uint32 = 0
    var v29 int32 = 42
    for fi30 := 0; fi30 < 6; fi30++ {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        _ = fi30
    }
    v5 = (v5 + v5)
    switch 0 {
    case 32767:
        fmt.Println(v3, v28)
    }
    fmt.Println(v22)
    v31 := GenS1{F0: v29}
    v32 := GenS2{F0: ((("foo bar" + v5) + v2) + "foo bar")}
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
