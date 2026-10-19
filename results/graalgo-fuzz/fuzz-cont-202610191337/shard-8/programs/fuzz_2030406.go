package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 uint8; F1 int; F2 string }

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]uint32)
    var v3 uint64 = 606
    var v4 int32 = -989
    var v5 string = "hello"
    if false {
        var v6 uint64 = 4294967295
        _ = v6
    }
    v5 = v5
    _ = len(v2)
    fmt.Println(v5)
    v2[581] = 732
    fmt.Println(v4, v4)
    v7 := GenS1{F0: 255, F1: 1, F2: 255}
    if (true && true) {
        if true {
            var v8 int64 = 42
            _ = v8
        }
    }
    var v9 int16 = 619
    if (2 >= 2) {
        switch v5 {
        case "":
            if true {
                switch 4294967295 {
                case 1:
                    v10 := GenS2{F0: 102, F1: 127, F2: "foo bar"}
                    _ = v10
                }
                var v11 int8 = 101
                _ = v11
            } else {
                if false {
                    var v12 uint8 = (v1 + (^(v1)))
                    _ = v12
                }
            }
        case "hello":
            var v13 uint64 = (v3 + 2)
            _ = v13
        case "foo bar":
            var v14 int32 = v4
            _ = v14
        }
        fmt.Println(v1)
    } else {
        var v15 uint64 = 42
        v4 = v4
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 string = (("hello" + (v5 + (v5 + "a"))) + v5)
            if (!false) {
                var v18 int16 = 1000
                _ = v18
            } else {
                v5 = "a"
            }
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    if false {
        if (776 == 748) {
            fmt.Println(v5, v9)
        } else {
            if (-381 > 741) {
                var v19 uint32 = 0
                delete(v2, 127)
                _ = v19
            }
            var v20 int64 = 1
            for fi21 := 0; fi21 < 8; fi21++ {
                switch fi21 {
                case -181:
                    v22 := []int{fi21}
                    _ = v22
                default:
                    v9 += (v9 << ((uint(v1)) & 63))
                }
                switch "\n\t" {
                case "":
                    v3 = (v3 << ((4294967295) & 63))
                case "foo bar":
                    v2[fi21] = 2147483648
                case "a":
                    fmt.Println(fi21, v1, v1)
                }
                _ = fi21
            }
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
