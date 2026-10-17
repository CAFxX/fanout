package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint }
func genFn0() bool {
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 string = "hello"
        _ = fi6
        _ = v7
    }
    var v8 int16 = 855
    _ = v8
    return false
}

func main() {
    var v1 float64 = (3.14159 - ((((0.1 / (3.14159 - 1.5)) - 3.14159) + 0.1) - -2.25))
    v2 := []int{42, -1}
    v3 := []uint32{2147483647, 1, 2147483647, 0}
    v4 := []uint32{0, 890}
    var v5 int64 = 236
    var v9 uint8 = 94
    v1 = (1.5 / v1)
    for fi10 := 0; fi10 < 5; fi10++ {
        switch ("hello" + ("" + ("\n\t" + ("foo bar" + "hello")))) {
        case "foo bar":
            for fi11 := 0; fi11 < 13; fi11++ {
                v1 = ((((v1 - v1) + (3.14159 - v1)) / (v1 + v1)) / -2.25)
                fmt.Println(fi11, v1, fi10)
                v12 := GenS1{F0: 0, F1: 0}
                _ = fi11
                _ = v12
            }
        default:
            v13 := GenS2{F0: 532, F1: 2147483647, F2: 540}
            _ = v13
        }
        _ = fi10
    }
    v9 = 2
    if (2147483647 != -624) {
        var v14 int64 = v5
        var v15 uint8 = (v9 / ((v9) | 1))
        _ = v14
        _ = v15
    } else {
        var v16 int = 32767
        if (!((-9223372036854775808 > (-(v5))) && true)) {
            _ = len(v3)
            var v17 bool = true
            if ((((v1 + float64(3.14159)) * (1.5 / (((v1 * (1e10 + 0.0)) * ((0.0 * ((v1 - (0.1 / float64(0.1))) - float64(0.1))) / (0.0 + (v1 / v1)))) * 0.1))) == (v1 / 1e10)) && (1 > v5)) {
                for fi18 := 0; fi18 < 7; fi18++ {
                    var v19 string = "hello"
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi18
                    _ = v19
                }
                v9 += (v9 >> ((766) & 63))
                var v20 string = ("\n\t" + "hello")
                _ = v20
            }
            _ = v17
        }
        _ = v16
    }
    v21 := GenS3{F0: 1}
    switch "\n\t" {
    case "":
        if (231 < 307) {
            switch -796 {
            case 1:
                var v22 float32 = 3.14159
                _ = v22
            default:
                v4 = append(v4, 2147483648)
            }
            var v23 string = ("hello" + "a")
            _ = v23
        }
    case "hello":
        v1 += -2.25
    }
    v1 = v1
    v24 := GenS1{F0: -1, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
