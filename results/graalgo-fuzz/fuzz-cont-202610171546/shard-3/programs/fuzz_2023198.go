package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }

func main() {
    var v1 float64 = -2.25
    var v2 string = "a"
    var v3 bool = (!false)
    v4 := make(map[int]string)
    var v5 uint16 = 929
    var v6 uint16 = v5
    fmt.Println(v6, v1)
    v7 := v4[80]
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 uint = uint(v6)
        delete(v4, fi8)
        _ = fi8
        _ = v9
    }
    if v3 {
        var v10 uint64 = 537
        _ = v10
    } else {
        switch -1 {
        case 32768:
            v3 = v3
        case 127:
            var v11 string = ""
            _ = v11
        default:
            _ = len(v4)
        }
    }
    v5 = 716
    v12 := GenS1{F0: (255 < 0), F1: 42, F2: 1}
    if v3 {
        var v13 bool = (42 > -440)
        var v14 bool = v3
        switch -2147483648 {
        case 32768:
            var v15 int32 = 42
            _ = v15
        case 795:
            v16 := GenS1{F0: (42 >= 1), F1: 208, F2: 2147483647}
            _ = v16
        default:
            fmt.Println(v2)
        }
        _ = v13
        _ = v14
    }
    if (!false) {
        v1 = v1
        var v17 uint8 = 72
        _ = v17
    } else {
        var v18 int16 = 42
        var v19 int32 = -1
        _ = v18
        _ = v19
    }
    var v20 uint64 = 0
    v21 := GenS1{F0: v3, F1: 791, F2: 328}
    var v22 uint16 = (v5 % (((v5 ^ v5)) | 1))
    var v23 bool = v3
    fmt.Println(v6, v5)
    for fi24 := 0; fi24 < 20; fi24++ {
        if ((1e10 < -2.25) && ((v20 ^ (v20 >> ((0) & 63))) <= v20)) {
            v25 := GenS1{F0: (!(v23 || (v23 || (-9223372036854775808 >= 789)))), F1: 1, F2: 42}
            var v26 float32 = 0.1
            v2 = "\n\t"
            _ = v25
            _ = v26
        } else {
            v27 := GenS1{F0: (-1 < 527), F1: 981, F2: 2147483647}
            switch ("a" + v7) {
            case "":
                v20 = 0
            case "foo bar":
                for fi28 := 0; fi28 < 2; fi28++ {
                    v6 = 65535
                    _ = fi28
                }
            case "a":
                v4[fi24] = (v2 + v7)
            }
            switch fi24 {
            case 128:
                v7 = v7
            case 1:
                for fi29 := 0; fi29 < 8; fi29++ {
                    switch 735 {
                    case -9223372036854775808:
                        fmt.Println(v22, v6)
                    default:
                        v30 := GenS1{F0: (219 <= 9223372036854775807), F1: 18446744073709551615, F2: int32(fi29)}
                        var v31 uint16 = (v5 << ((2) & 63))
                        fmt.Println(v5)
                        _ = v30
                        _ = v31
                    }
                    _ = fi29
                }
            }
            _ = v27
        }
        for fi32 := 0; fi32 < 15; fi32++ {
            var v33 int16 = 0
            _ = fi32
            _ = v33
        }
        _ = fi24
    }
    v34 := GenS1{F0: (-446 < 4294967295), F1: 631, F2: 587}
    v35 := GenS1{F0: true, F1: 626, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
