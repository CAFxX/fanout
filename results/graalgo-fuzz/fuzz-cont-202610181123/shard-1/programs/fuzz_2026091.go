package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int = 127
    var v2 int32 = -1
    var v3 uint = 744
    var v4 int8 = -102
    v5 := make(map[uint32]int)
    v6 := make(map[int]string)
    v7 := []int{}
    for fi8 := 0; fi8 < 5; fi8++ {
        switch ("hello" + "foo bar") {
        case "foo bar":
            _ = len(v7)
        case "a":
            if true {
                v9 := GenS1{F0: 9223372036854775807}
                if false {
                    var v10 uint64 = 42
                    var v11 uint16 = uint16(v10)
                    _ = v10
                    _ = v11
                } else {
                    v2 = 0
                }
                _ = v9
            }
        }
        if (127 == 127) {
            v12 := v6[v1]
            _ = v12
        }
        _ = fi8
    }
    fmt.Println(v4, v2, v2)
    v2 = v2
    for fi13 := 0; fi13 < 20; fi13++ {
        switch (v1 - fi13) {
        case 1:
            v1 = v1
        case -1:
            switch -466 {
            case 128:
                var v14 float64 = ((((3.14159 * 3.14159) - 1e10) - (0.0 * 3.14159)) / float64(0.1))
                _ = v14
            default:
                v15 := GenS1{F0: 678}
                _ = v15
            }
        case -9223372036854775808:
            for fi16 := 0; fi16 < 15; fi16++ {
                switch v1 {
                case -217:
                    var v17 int8 = v4
                    fmt.Println(v4, fi16)
                    if false {
                        v6[302] = "\n\t"
                        fmt.Println(v1, v17)
                    }
                    _ = v17
                case 9223372036854775807:
                    var v18 uint8 = 202
                    _ = v18
                default:
                    var v19 bool = (!(211 <= 516))
                    _ = v19
                }
                fmt.Println(v4, v3, v2)
                v20 := GenS1{F0: 42}
                _ = fi16
                _ = v20
            }
        default:
            v1 = v1
        }
        _ = fi13
    }
    if true {
        v21 := GenS1{F0: 0}
        for fi22 := 0; fi22 < 7; fi22++ {
            v6[int(v2)] = "foo bar"
            v2 = (v2 % (((v2 * 2147483647)) | 1))
            _ = fi22
        }
        if true {
            var v23 uint16 = 0
            v2 = -332
            v3 = (v3 % (((v3 + 65535)) | 1))
            _ = v23
        }
        _ = v21
    }
    v3 = (^(v3))
    v24 := v5[385]
    if true {
        _ = len(v5)
        var v25 int = v1
        _ = v25
    } else {
        v26 := GenS1{F0: 23}
        for fi27 := 0; fi27 < 6; fi27++ {
            delete(v5, 2)
            _ = fi27
        }
        _ = v26
    }
    var v28 string = ("\n\t" + ("" + "a"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
