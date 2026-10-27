package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 42
    var v3 uint8 = 42
    var v4 string = ("\n\t" + "\n\t")
    v5 := []int{-1, 4294967295, 1, -9223372036854775808}
    var v6 uint32 = 2147483648
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(v4)
        fmt.Println(fi7)
        _ = fi7
    }
    fmt.Println(v6, v2)
    if true {
        if false {
            var v8 float32 = -2.25
            _ = v8
        } else {
            var v9 int32 = 0
            fmt.Println(v9, v9, v2)
            _ = v9
        }
        fmt.Println(v6)
    } else {
        var v10 int32 = 0
        _ = v10
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v11 string = ((v4 + v4) + (("foo bar" + v4) + v4))
    v4 = v4
    if (!true) {
        if false {
            var v12 uint8 = 127
            _ = v12
        }
    }
    fmt.Println(v4, v6)
    if true {
        switch 4294967295 {
        case 4294967295:
            v11 = "hello"
        default:
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 bool = (!false)
                fmt.Println(v11, fi13, v2)
                for fi15 := 0; fi15 < 17; fi15++ {
                    switch v11 {
                    case "foo bar":
                        v2 = v2
                        fmt.Println(fi15)
                    }
                    var v16 float32 = (float32(1.5) * 0.1)
                    v17 := GenS1{F0: 1, F1: 0}
                    switch -945 {
                    case -697:
                        var v18 int16 = 640
                        _ = v18
                    case 0:
                        if ("foo bar" != "a") {
                            fmt.Println(v14, v11, v3)
                        }
                    case 32768:
                        v14 = false
                    default:
                        var v19 uint8 = 0
                        _ = v19
                    }
                    v6 += v6
                    v14 = (false || ((v11 + "foo bar") <= ""))
                    switch fi15 {
                    case 42:
                        var v20 int64 = 1
                        _ = v20
                    default:
                        v5 = append(v5, (fi13 >> ((18446744073709551615) & 63)))
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = fi13
                _ = v14
            }
        }
    }
    v21 := GenS1{F0: -1, F1: 2}
    v22 := GenS1{F0: 470, F1: 357}
    v23 := GenS1{F0: 0, F1: 248}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
