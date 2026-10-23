package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []int{}
    v2 := make(map[string]string)
    var v3 int32 = 42
    v4 := []int{}
    v5 := GenS1{F0: -9223372036854775808}
    if true {
        var v6 int32 = -2147483648
        fmt.Println(v6)
        v3 = 42
        _ = v6
    } else {
        v7 := GenS1{F0: 0}
        fmt.Println(v3)
        _ = v7
    }
    v8 := GenS2{F0: -9223372036854775808}
    if (!true) {
        switch 9223372036854775807 {
        case 4294967295:
            var v9 int16 = 32767
            _ = v9
        case 32768:
            v3 = v3
        default:
            _ = len(v1)
        }
        var v10 int = 32768
        switch "hello" {
        case "\n\t":
            var v11 uint16 = 2
            _ = v11
        case "":
            fmt.Println(v10)
        case "hello":
            if false {
                v10 += 2147483647
                fmt.Println(v10, v3)
            } else {
                delete(v2, ("" + ("foo bar" + ("\n\t" + "\n\t"))))
            }
        default:
            var v12 bool = false
            _ = v12
        }
        _ = v10
    } else {
        for fi13 := 0; fi13 < 15; fi13++ {
            if true {
                if (!(390 <= 2)) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            } else {
                v3 -= -1
                var v14 uint32 = 2147483648
                _ = v14
            }
            _ = fi13
        }
        var v15 float64 = (0.1 * (1.5 + 3.14159))
        var v16 float64 = v15
        _ = v15
        _ = v16
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 int = fi17
            switch "a" {
            case "a":
                if true {
                    v19 := GenS1{F0: -9223372036854775808}
                    _ = v19
                }
            case "hello":
                var v20 int8 = 42
                _ = v20
            case "foo bar":
                fmt.Println(v3, v3)
            default:
                v18 = v18
            }
            v18 = -9223372036854775808
            _ = fi17
            _ = v18
        }
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 uint32 = 42
        _ = fi21
        _ = v22
    }
    v2[""] = "foo bar"
    if false {
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 int16 = 1000
            v24 = -939
            for fi25 := 0; fi25 < 19; fi25++ {
                var v26 string = "\n\t"
                v26 = v26
                v24 -= v24
                _ = fi25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        if false {
            v2["\n\t"] = "hello"
            switch 32767 {
            case 670:
                for fi27 := 0; fi27 < 8; fi27++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi27
                }
            }
        }
        v3 -= 935
    } else {
        for fi28 := 0; fi28 < 6; fi28++ {
            v29 := GenS1{F0: 42}
            _ = fi28
            _ = v29
        }
        var v30 float64 = ((((0.1 + -2.25) / ((1e10 / float64(1.5)) - 1.5)) / 3.14159) * -2.25)
        var v31 int32 = (v3 & v3)
        _ = v30
        _ = v31
    }
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
