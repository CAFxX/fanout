package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0() uint8 {
    var v4 int32 = 366
    var v5 int16 = -32768
    if (!true) {
        var v6 int8 = 1
        _ = v6
    } else {
        fmt.Println(v4, v4)
        switch "a" {
        case "foo bar":
            v7 := []int{999, -9223372036854775808}
            _ = v7
        case "\n\t":
            v8 := GenS1{F0: (0 == 42)}
            _ = v8
        default:
            switch 128 {
            case -9223372036854775808:
                var v9 uint16 = 2
                _ = v9
            case -439:
                v10 := GenS1{F0: false}
                _ = v10
            case 32768:
                for fi11 := 0; fi11 < 4; fi11++ {
                    fmt.Println(fi11, v5, v5)
                    _ = fi11
                }
            default:
                if true {
                    v12 := GenS1{F0: (0 <= -128)}
                    var v13 uint64 = 223
                    _ = v12
                    _ = v13
                } else {
                    v14 := []int{4294967295, -9223372036854775808, 657}
                    _ = v14
                }
            }
        }
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int8 = 42
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := GenS1{F0: true}
            v19 := []int{fi17, (fi17 >> ((662) & 63)), (fi15 & fi17)}
            var v20 string = "foo bar"
            v4 += (v4 % ((-1) | 1))
            fmt.Println(v4, v20, v20)
            var v21 uint8 = 128
            var v22 float64 = 1.5
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi15
        _ = v16
    }
    _ = v4
    _ = v5
    return 255
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    v3 := []string{"\n\t", ("" + "hello"), "a", "hello"}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            fmt.Println(fi24, v2, fi23)
            var v25 uint32 = 647
            var v26 int8 = 1
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 int16 = -1
        _ = fi23
        _ = v27
    }
    var v28 float64 = 0.0
    var v29 int8 = -128
    v30 := v1[127]
    var v31 uint8 = 2
    v32 := GenS2{F0: (v2 && v2), F1: 0}
    v33 := GenS1{F0: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
