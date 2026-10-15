package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
func genFn0() int16 {
    var v7 uint16 = 255
    if false {
        v7 = v7
        if (((!false) || (true && ((-721 > 1000) && false))) || ((!false) && true)) {
            if false {
                v8 := []int{929, 32768}
                v7 -= v7
                var v9 string = ("a" + ("\n\t" + ""))
                _ = v8
                _ = v9
            } else {
                var v10 int16 = 1000
                for fi11 := 0; fi11 < 2; fi11++ {
                    v12 := []int{42, int(v7), (fi11 & fi11), 128}
                    _ = fi11
                    _ = v12
                }
                v13 := []int{-9223372036854775808, -102, -56, 398}
                _ = v10
                _ = v13
            }
            v14 := GenS1{F0: "hello"}
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 int = (fi15 ^ (fi15 | fi15))
                v17 := GenS2{F0: 4294967295}
                v18 := GenS3{F0: 255}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = v14
        } else {
            var v19 uint16 = (^(v7))
            _ = v19
        }
    }
    _ = v7
    return -1
}

func main() {
    var v1 bool = false
    var v2 int8 = 1
    v3 := make(map[int]int)
    var v4 uint8 = 0
    v5 := make(map[int]uint32)
    var v6 float64 = ((float64(((1e10 * ((1.5 * -2.25) / 3.14159)) / 1.5)) + (((0.1 * 3.14159) * float64((0.1 - float32(0.1)))) * 0.0)) - 0.0)
    fmt.Println(v1, v1, v2)
    for fi20 := 0; fi20 < 18; fi20++ {
        switch 2147483647 {
        case 2147483647:
            switch "a" {
            case "foo bar":
                fmt.Println(v4)
            case "\n\t":
                _ = len(v3)
            default:
                v21 := []int{(fi20 / ((fi20) | 1)), fi20}
                _ = v21
            }
        default:
            if (v1 && ((1 == 252) || false)) {
                v2 -= v2
                v22 := GenS1{F0: ("" + "hello")}
                _ = v22
            }
        }
        v6 += (v6 / 1.0)
        v2 += -1
        _ = fi20
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        fmt.Println(v1, v4)
        _ = fi23
    }
    if false {
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := v5[fi24]
            _ = fi24
            _ = v25
        }
    } else {
        for fi26 := 0; fi26 < 19; fi26++ {
            v6 += (float64(1.5) + (v6 * ((v6 - (-2.25 - (v6 * ((float64((-2.25 - 0.0)) - (0.1 / v6)) - v6)))) / v6)))
            _ = fi26
        }
        delete(v3, 374)
    }
    var v27 float64 = 0.0
    var v28 int16 = 32767
    fmt.Println(v4, v27, v4)
    var v29 int64 = -667
    for fi30 := 0; fi30 < 1; fi30++ {
        v31 := GenS4{F0: 572}
        fmt.Println(v4, v4, v1)
        fmt.Println(v27)
        _ = fi30
        _ = v31
    }
    fmt.Println(v27, v6)
    v32 := GenS2{F0: 4294967295}
    if false {
        var v33 int64 = 1
        _ = v33
    }
    var v34 string = ("\n\t" + "")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
