package main

import (
    "fmt"
)


func main() {
    var v1 int32 = -1
    v2 := make(map[string]int)
    var v3 string = ("\n\t" + "hello")
    var v4 int8 = -17
    var v5 uint64 = 922
    v6 := make(map[uint32]int)
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 string = (((v3 + v3) + ((v3 + v3) + (v3 + v3))) + v3)
        for fi9 := 0; fi9 < 13; fi9++ {
            if false {
                v8 = "foo bar"
                var v10 string = v8
                v10 = ((("foo bar" + v8) + "foo bar") + "a")
                _ = v10
            } else {
                fmt.Println(fi7, v4)
                var v11 int64 = -1
                _ = v11
            }
            _ = fi9
        }
        var v12 int64 = 9223372036854775807
        _ = fi7
        _ = v12
        _ = v8
    }
    var v13 int16 = 32767
    if (v3 <= ("foo bar" + v3)) {
        var v14 float32 = 1e10
        var v15 uint = 853
        _ = v14
        _ = v15
    }
    v13 -= (v13 - v13)
    var v16 float32 = float32(0.1)
    var v17 int16 = v13
    v3 = v3
    fmt.Println(v13, v5)
    switch -643 {
    case -758:
        v13 = (-(v17))
    case -2147483648:
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 uint16 = 65535
                var v21 uint16 = (v20 & 2)
                _ = fi19
                _ = v20
                _ = v21
            }
            v4 += (v4 << ((4294967295) & 63))
            v22 := v6[2]
            _ = fi18
            _ = v22
        }
    case 1:
        var v23 float32 = (v16 * -2.25)
        _ = v23
    }
    switch -9223372036854775808 {
    case 946:
        var v24 int64 = -9223372036854775808
        _ = v24
    case -9223372036854775808:
        v1 += v1
    default:
        var v25 uint64 = v5
        _ = v25
    }
    var v26 uint8 = 2
    v27 := []int{545, -1, -797}
    var v28 int64 = 9223372036854775807
    for fi29 := 0; fi29 < 1; fi29++ {
        var v30 float64 = 0.1
        _ = fi29
        _ = v30
    }
    var v31 int = -406
    v4 += v4
    v28 = v28
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
