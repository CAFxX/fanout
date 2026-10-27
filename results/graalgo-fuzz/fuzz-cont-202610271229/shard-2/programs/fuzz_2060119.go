package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint = 255
    var v2 string = "hello"
    var v3 string = ""
    var v4 int8 = 0
    if true {
        for fi5 := 0; fi5 < 20; fi5++ {
            v6 := []int{0, (^(fi5))}
            v6 = append(v6, (fi5 << (((v1 >> ((4294967295) & 63))) & 63)))
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: v2, F1: -2147483648, F2: 65535}
        v3 = "hello"
        _ = v7
    } else {
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 uint32 = 2147483647
                v4 -= v4
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        fmt.Println(v3, v2, v1)
        fmt.Println(v4, v2, v4)
    }
    var v11 int = 4294967295
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint = v1
        fmt.Println(fi12, v13)
        _ = fi12
        _ = v13
    }
    v3 = "a"
    if (!(!(868 == 4294967295))) {
        var v14 uint32 = 42
        fmt.Println(v11, v1)
        for fi15 := 0; fi15 < 20; fi15++ {
            switch 0 {
            case 4294967295:
                if false {
                    var v16 float32 = 0.0
                    _ = v16
                } else {
                    v2 = ""
                }
            default:
                v17 := []int{int(v14)}
                _ = v17
            }
            v3 = v2
            for fi18 := 0; fi18 < 10; fi18++ {
                switch (v2 + "\n\t") {
                case "foo bar":
                    if (false || (((1 != (v1 | v1)) && (203 <= 234)) && ((("\n\t" + "hello") + v3) < v2))) {
                        v19 := GenS1{F0: v2, F1: 2147483647, F2: (^(v14))}
                        v20 := GenS2{F0: -32768}
                        var v21 string = v2
                        fmt.Println(fi15)
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                }
                fmt.Println(v1)
                _ = fi18
            }
            _ = fi15
        }
        _ = v14
    }
    if (423 <= 30) {
        var v22 int16 = -32768
        var v23 int64 = 160
        _ = v22
        _ = v23
    }
    v24 := []int{}
    v1 += (v1 + 42)
    var v25 int = 32767
    var v26 uint16 = 570
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
