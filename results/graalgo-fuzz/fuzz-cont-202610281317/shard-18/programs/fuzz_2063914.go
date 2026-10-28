package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }

func main() {
    var v1 uint32 = 0
    var v2 uint16 = 42
    var v3 int64 = 646
    var v4 int32 = 1
    v5 := []int{9223372036854775807}
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint = 959
        for fi8 := 0; fi8 < 10; fi8++ {
            if (v2 >= v2) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = fi8
        }
        v2 += (v2 | (^(v2)))
        _ = fi6
        _ = v7
    }
    if true {
        fmt.Println(v2, v2)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v3 = v3
    }
    var v9 uint32 = 0
    if true {
        v2 = v2
        fmt.Println(v3)
    }
    switch "\n\t" {
    case "":
        v9 = v1
    }
    var v10 float32 = 3.14159
    var v11 int32 = -884
    if false {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v1)
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        _ = len(v5)
        if true {
            v3 = v3
            fmt.Println(v11)
            var v13 bool = (620 < 4294967295)
            _ = v13
        } else {
            switch (fi12 / ((fi12) | 1)) {
            case -2147483648:
                var v14 uint = 2
                _ = v14
            case 32768:
                v2 = v2
            default:
                fmt.Println(fi12)
            }
            v4 -= (v4 - v4)
            v1 = (v1 / ((v1) | 1))
        }
        v15 := GenS1{F0: ("a" + (("" + "") + "hello")), F1: 188}
        _ = fi12
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        switch -918 {
        case -1:
            v4 += (^(v11))
        default:
            if ((v9 & (v9 & v1)) > (v1 % ((783) | 1))) {
                for fi17 := 0; fi17 < 17; fi17++ {
                    if false {
                        var v18 float64 = float64(((v10 / (v10 - float32(0.1))) * v10))
                        var v19 uint = 1
                        _ = v18
                        _ = v19
                    }
                    _ = fi17
                }
                var v20 int32 = (v11 | v11)
                v1 += v1
                _ = v20
            } else {
                var v21 uint64 = 849
                _ = v21
            }
        }
        _ = fi16
    }
    if ((v9 * v1) == (v9 * (v9 - (^(v9))))) {
        var v22 uint32 = v9
        v23 := GenS1{F0: "a", F1: 1}
        _ = v22
        _ = v23
    } else {
        if len(v5) > 0 {
            v5[0] = int(v1)
        }
        v5 = append(v5, 2147483647)
        v2 = 822
    }
    var v24 float32 = (1.5 - v10)
    v11 -= (v11 * (v4 >> ((42) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
