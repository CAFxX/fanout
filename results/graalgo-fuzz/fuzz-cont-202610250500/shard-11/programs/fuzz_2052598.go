package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int64 = 42
    var v2 float64 = ((1e10 - (float64(-2.25) * 0.1)) * (-2.25 * 0.0))
    v3 := []uint32{}
    v4 := GenS1{F0: 0}
    if (!false) {
        if true {
            var v5 int = 42
            fmt.Println(v2, v2)
            _ = v5
        }
        switch -917 {
        case 2147483647:
            v3 = append(v3, 0)
        default:
            if false {
                v6 := GenS1{F0: -2147483648}
                if true {
                    if false {
                        v7 := GenS1{F0: -635}
                        _ = v7
                    }
                }
                _ = v6
            } else {
                v8 := GenS2{F0: 93}
                _ = v8
            }
        }
    }
    v1 += (v1 + 1)
    var v9 uint32 = 2
    for fi10 := 0; fi10 < 13; fi10++ {
        v1 -= (v1 ^ (v1 >> ((4294967295) & 63)))
        _ = fi10
    }
    switch 1 {
    case 32767:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 127:
        var v11 uint32 = (v9 / (((v9 & v9)) | 1))
        _ = v11
    case -831:
        var v12 int64 = (v1 >> ((0) & 63))
        _ = v12
    }
    var v13 int = -2147483648
    if (-136 > (v1 >> ((963) & 63))) {
        v9 += (v9 & 633)
        v14 := GenS1{F0: 39}
        _ = v14
    }
    v9 += v9
    var v15 float64 = (float64(-2.25) - 0.0)
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS1{F0: 0}
        _ = fi16
        _ = v17
    }
    if ((v9 + (v9 & v9)) < (v9 << ((364) & 63))) {
        var v18 uint = 0
        for fi19 := 0; fi19 < 15; fi19++ {
            fmt.Println(fi19, v15)
            var v20 uint = (v18 << (((^(v18))) & 63))
            _ = fi19
            _ = v20
        }
        _ = v18
    } else {
        v21 := GenS3{F0: (v9 >> ((838) & 63))}
        _ = v21
    }
    if (!((60 >= 2) || ((v15 >= ((v2 - (v2 + ((v2 - (v2 + v15)) * (1e10 + ((v15 - float64(0.0)) - v15))))) + 0.0)) && (true || (127 >= 100))))) {
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
