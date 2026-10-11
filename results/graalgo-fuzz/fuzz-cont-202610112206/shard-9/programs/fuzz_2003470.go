package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint32 }
func genFn0() string {
    v7 := []int{-2147483648}
    switch "foo bar" {
    case "a":
        var v8 int8 = -1
        _ = v8
    case "\n\t":
        var v9 string = "a"
        _ = v9
    }
    switch -778 {
    case 32768:
        var v10 uint64 = 543
        _ = v10
    case 128:
        var v11 float32 = 1e10
        _ = v11
    case -1:
        var v12 uint8 = 168
        _ = v12
    }
    _ = v7
    return ""
}
func genFn1(p0 uint8) bool {
    var v13 int8 = 42
    v14 := GenS1{F0: 2, F1: 128, F2: 24}
    _ = p0
    _ = v13
    _ = v14
    return (!false)
}

func main() {
    var v1 int64 = 1
    var v2 uint16 = 421
    var v3 uint = 1
    var v4 int64 = v1
    v5 := []string{"\n\t", ((("a" + "") + ("\n\t" + ("hello" + "a"))) + ""), ("\n\t" + "\n\t")}
    var v6 uint8 = 0
    v6 += v6
    var v15 int8 = -125
    var v16 uint16 = (v2 ^ (v2 % ((uint16(v6)) | 1)))
    v6 += 42
    fmt.Println(v16, v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch -218 {
    case -927:
        fmt.Println(v15)
    }
    v4 = (v1 << ((v3) & 63))
    var v17 float32 = -2.25
    for fi18 := 0; fi18 < 4; fi18++ {
        if true {
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 string = (("foo bar" + (("\n\t" + "hello") + "\n\t")) + "a")
                _ = fi19
                _ = v20
            }
            fmt.Println(v3, v16)
            v2 -= (v16 ^ 0)
        }
        v6 = 1
        if (("a" >= "\n\t") || (-541 < 371)) {
            v4 -= v1
            switch fi18 {
            case 1:
                v21 := GenS1{F0: 0, F1: 58, F2: 65535}
                _ = v21
            case 0:
                v17 += (1.5 + v17)
            }
        }
        _ = fi18
    }
    var v22 string = "a"
    v5 = append(v5, v22)
    var v23 int32 = 1
    var v24 uint8 = (v6 | 255)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
