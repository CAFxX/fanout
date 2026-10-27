package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
func genFn0() int16 {
    if false {
        v6 := []int{32768, 768, 2147483647}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint = 231
            if false {
                var v9 int8 = -128
                var v10 int8 = 1
                v11 := GenS1{F0: fi7}
                _ = v10
                _ = v11
                _ = v9
            }
            switch fi7 {
            case 389:
                var v12 int8 = 0
                _ = v12
            case 0:
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 bool = (32767 != 1)
                    fmt.Println(v14, v14, v14)
                    var v15 int64 = 42
                    var v16 uint8 = 1
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            default:
                if false {
                    var v17 bool = ((fi7 << (((v8 ^ (v8 + (^(v8))))) & 63)) < (-(fi7)))
                    v18 := []int{}
                    _ = v17
                    _ = v18
                }
            }
            _ = fi7
            _ = v8
        }
    }
    v19 := []int{-9223372036854775808, 32768, 1, 9223372036854775807}
    switch 4294967295 {
    case 32768:
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 int64 = 42
            v22 := GenS1{F0: fi20}
            _ = fi20
            _ = v21
            _ = v22
        }
    case -613:
        switch -583 {
        case 2147483647:
            var v23 int8 = -46
            _ = v23
        }
    }
    _ = v19
    return -280
}
func genFn1(p0 int8, p1 int32) string {
    var v24 uint64 = 2
    fmt.Println(p1)
    var v25 uint8 = 239
    v26 := GenS2{F0: p0}
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    _ = v26
    return ""
}

func main() {
    var v1 int64 = 42
    var v2 uint8 = 1
    var v3 uint32 = 2147483648
    var v4 bool = false
    var v5 float64 = 1.5
    v27 := []int{4294967295, -9223372036854775808}
    v28 := GenS2{F0: -66}
    v3 += 2147483647
    if len(v27) > 0 {
        v27[0] = 42
    }
    v29 := GenS2{F0: -1}
    for fi30 := 0; fi30 < 3; fi30++ {
        if v4 {
            switch (fi30 >> ((255) & 63)) {
            case 42:
                var v31 int = fi30
                _ = v31
            case -1:
                var v32 int8 = 42
                _ = v32
            case 1:
                if len(v27) > 0 {
                    v27[0] = fi30
                }
            }
            switch (^(fi30)) {
            case 42:
                v3 += (^(v3))
            }
            v27 = append(v27, 1)
        } else {
            switch "" {
            case "\n\t":
                for fi33 := 0; fi33 < 10; fi33++ {
                    v34 := GenS1{F0: fi33}
                    var v35 int64 = v1
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            case "":
                _ = len(v27)
            }
        }
        var v36 bool = ((false && true) && (false || v4))
        _ = fi30
        _ = v36
    }
    v5 += float64((1e10 * (0.1 - ((1.5 + (1.5 - 1.5)) - 0.0))))
    v37 := GenS2{F0: -128}
    v3 = (v3 >> ((0) & 63))
    v4 = ((v4 && v4) || v4)
    var v38 int = 42
    for fi39 := 0; fi39 < 2; fi39++ {
        fmt.Println(fi39, v2)
        v40 := GenS3{F0: (42 == 42)}
        _ = fi39
        _ = v40
    }
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    v2 += (v2 + v2)
    v41 := GenS4{F0: 42}
    v42 := GenS4{F0: v2}
    for fi43 := 0; fi43 < 2; fi43++ {
        v2 -= v2
        if len(v27) > 0 {
            v27[0] = (fi43 + (v38 / ((v38) | 1)))
        }
        var v44 uint16 = 2
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
