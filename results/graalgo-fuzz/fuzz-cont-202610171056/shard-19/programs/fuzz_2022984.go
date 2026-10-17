package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []string{}
    v2 := []uint32{}
    var v3 int16 = -1
    v4 := []uint32{2}
    var v5 bool = ((v3 ^ (v3 + v3)) < -824)
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint8 = 2
        v7 -= v7
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi8
    }
    var v9 int = -1
    for fi10 := 0; fi10 < 20; fi10++ {
        if v5 {
            switch 905 {
            case 2147483647:
                for fi11 := 0; fi11 < 6; fi11++ {
                    fmt.Println(v3)
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    _ = fi11
                }
            case 9223372036854775807:
                switch (-(fi10)) {
                case -749:
                    v5 = v5
                }
            }
        }
        v5 = v5
        _ = len(v1)
        _ = fi10
    }
    v5 = (-9223372036854775808 >= 1)
    var v12 float32 = 1.5
    var v13 bool = (-9223372036854775808 < 0)
    fmt.Println(v9, v9, v12)
    var v14 float32 = v12
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 string = ""
            _ = fi16
            _ = v17
        }
        if false {
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 uint64 = 592
                switch fi15 {
                case 322:
                    fmt.Println(fi15, fi18)
                case 4294967295:
                    v19 -= v19
                    v20 := GenS1{F0: 329}
                    var v21 int64 = 0
                    _ = v20
                    _ = v21
                default:
                    var v22 uint = 4294967295
                    _ = v22
                }
                var v23 float32 = float32(((3.14159 + 1e10) / ((1.5 * (3.14159 - -2.25)) + (1.5 / 0.1))))
                _ = fi18
                _ = v19
                _ = v23
            }
        } else {
            if len(v4) > 0 {
                v4[0] = 2
            }
        }
        v24 := GenS1{F0: 2147483647}
        _ = fi15
        _ = v24
    }
    var v25 float32 = 0.0
    for fi26 := 0; fi26 < 13; fi26++ {
        v3 += -683
        if true {
            v9 = (v9 & v9)
            fmt.Println(v14, v3)
        }
        _ = fi26
    }
    var v27 int = v9
    var v28 uint = 198
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
