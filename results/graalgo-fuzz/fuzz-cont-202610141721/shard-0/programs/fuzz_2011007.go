package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint32 = 193
    var v3 int64 = (v1 & -9223372036854775808)
    var v4 int16 = -258
    var v5 int32 = 0
    var v6 int8 = -1
    v1 -= v1
    for fi7 := 0; fi7 < 3; fi7++ {
        v5 = int32(v2)
        v8 := GenS1{F0: 2}
        var v9 int8 = v6
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int32 = (v5 + v5)
        _ = fi10
        _ = v11
    }
    var v12 float64 = (3.14159 * -2.25)
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 float32 = (-2.25 / 1e10)
        _ = fi13
        _ = v14
    }
    var v15 uint64 = 1
    switch 42 {
    case 32768:
        fmt.Println(v15, v4, v12)
    case -748:
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 uint8 = 42
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS1{F0: 2}
                v20 := []int{4294967295}
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi16
            _ = v17
        }
    case -885:
        if (0 > v2) {
            var v21 uint32 = (v2 << ((18446744073709551615) & 63))
            v22 := GenS1{F0: 32767}
            for fi23 := 0; fi23 < 5; fi23++ {
                v24 := []int{int(v3), fi23}
                var v25 bool = false
                for fi26 := 0; fi26 < 19; fi26++ {
                    if v25 {
                        var v27 int16 = int16(v5)
                        v24 = append(v24, fi26)
                        var v28 int16 = 0
                        _ = v27
                        _ = v28
                    }
                    _ = fi26
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = v21
            _ = v22
        } else {
            var v29 uint16 = 255
            v15 += (^(v15))
            _ = v29
        }
    }
    var v30 uint8 = 128
    v4 -= -371
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
