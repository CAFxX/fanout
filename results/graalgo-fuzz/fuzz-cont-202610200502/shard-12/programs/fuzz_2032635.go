package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int64; F1 string }
func genFn0() uint64 {
    if false {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 int16 = 425
            v9 := []int{fi7}
            for fi10 := 0; fi10 < 18; fi10++ {
                v8 = v8
                _ = fi10
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        var v11 float32 = (0.1 - 3.14159)
        _ = v11
    } else {
        var v12 string = "foo bar"
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        if false {
            switch "" {
            case "hello":
                v14 := GenS1{F0: false, F1: -1}
                _ = v14
            default:
                var v15 bool = true
                _ = v15
            }
        } else {
            v16 := []int{fi13}
            if false {
                var v17 uint16 = 292
                for fi18 := 0; fi18 < 4; fi18++ {
                    var v19 uint = 2
                    v19 = (^(v19))
                    _ = fi18
                    _ = v19
                }
                v20 := GenS2{F0: 2147483647, F1: 255}
                _ = v17
                _ = v20
            }
            var v21 uint32 = 2147483648
            _ = v16
            _ = v21
        }
        _ = fi13
    }
    return 42
}

func main() {
    var v1 int32 = -1
    var v2 uint = 560
    var v3 float64 = 1e10
    var v4 string = ("\n\t" + "")
    var v5 uint32 = 959
    var v6 int8 = -1
    v22 := GenS3{F0: 59, F1: (!true)}
    v6 -= v6
    var v23 int8 = v6
    if (false || true) {
        v23 -= (v6 * 21)
    }
    v24 := []int{-2147483648, -1}
    for fi25 := 0; fi25 < 9; fi25++ {
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        _ = fi25
    }
    if (!(!false)) {
        v26 := GenS4{F0: 436, F1: v4}
        _ = v26
    } else {
        var v27 int16 = -1
        _ = v27
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        _ = len(v24)
        v29 := GenS3{F0: (v5 + v5), F1: (557 >= 302)}
        _ = fi28
        _ = v29
    }
    var v30 uint16 = 65535
    v31 := GenS2{F0: 313, F1: 42}
    var v32 float64 = ((1.5 / 1.5) + (1.5 + (v3 + v3)))
    v33 := GenS2{F0: (^(v5)), F1: 255}
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    switch 42 {
    case -637:
        _ = len(v24)
    case 12:
        switch 4294967295 {
        case 356:
            var v34 uint16 = v30
            _ = v34
        default:
            v35 := GenS4{F0: -1, F1: v4}
            _ = v35
        }
    case 32768:
        fmt.Println(v32)
    }
    if ((true || ((!true) || ((0.0 == 3.14159) && ((!(float64((((float32(((1.5 / v32) + v32)) / ((1e10 * 1.5) - (float32(v32) * 0.1))) * -2.25) / 1.5)) < 1.5)) && (true || ((v4 + "\n\t") == "")))))) && true) {
        v6 -= (v23 % ((-128) | 1))
        var v36 uint8 = 1
        _ = v36
    }
    v1 = (v1 / ((-2147483648) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
