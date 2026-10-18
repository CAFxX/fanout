package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint16 = 1
    var v2 uint = 894
    var v3 uint16 = v1
    var v4 int64 = 9223372036854775807
    var v5 uint = 285
    var v6 string = "\n\t"
    v6 = v6
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 float64 = 1.5
        _ = fi7
        _ = v8
    }
    if false {
        v1 = 65535
        var v9 uint = 529
        v10 := GenS1{F0: v3}
        _ = v10
        _ = v9
    }
    var v11 int8 = 1
    var v12 uint32 = 1
    var v13 uint64 = 2
    v14 := []int{2147483647, 2147483647, 32768}
    v14 = append(v14, int(v2))
    for fi15 := 0; fi15 < 6; fi15++ {
        switch (fi15 ^ fi15) {
        case 2147483647:
            switch v6 {
            case "\n\t":
                if true {
                    _ = len(v14)
                }
            default:
                v4 = v4
            }
        case 842:
            switch fi15 {
            case 441:
                fmt.Println(v4, v3, v13)
            case 4294967295:
                fmt.Println(v13, v4)
            case 32767:
                if false {
                    if (42 < v13) {
                        v6 = (((v6 + ("" + (("a" + "a") + ("hello" + v6)))) + v6) + (v6 + v6))
                    } else {
                        v16 := GenS1{F0: v3}
                        var v17 float32 = 1e10
                        var v18 float64 = 1e10
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    switch (fi15 + (fi15 * int(v2))) {
                    case 128:
                        if true {
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        }
                    }
                }
            }
        case -865:
            fmt.Println(v6)
        default:
            fmt.Println(fi15)
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 int32 = 0
            var v21 int64 = -9223372036854775808
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
