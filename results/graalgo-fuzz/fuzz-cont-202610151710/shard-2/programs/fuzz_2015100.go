package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int16; F1 int8; F2 int }

func main() {
    var v1 uint8 = 249
    var v2 bool = (true || (42 != 456))
    v3 := make(map[string]int)
    var v4 int8 = 100
    var v5 float64 = ((((float64(0.0) / (-2.25 - 1.5)) / 3.14159) + 3.14159) - -2.25)
    var v6 uint32 = 65535
    var v7 string = "foo bar"
    v1 = (^(v1))
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: 2147483647, F1: v2}
        _ = fi8
        _ = v9
    }
    v5 += v5
    for fi10 := 0; fi10 < 12; fi10++ {
        v5 = (float64(3.14159) * v5)
        var v11 int8 = v4
        _ = fi10
        _ = v11
    }
    fmt.Println(v4)
    var v12 uint = 65535
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS1{F0: 179, F1: true}
        if ((float64(float32(1.5)) / 3.14159) >= 0.1) {
            delete(v3, (v7 + (v7 + v7)))
            if ((0.1 > ((1.5 * 0.0) / (1.5 + (1.5 - (3.14159 + 1.5))))) && ((("foo bar" + "\n\t") + "\n\t") >= ("\n\t" + v7))) {
                if v2 {
                    if true {
                        v15 := GenS2{F0: -230, F1: v4, F2: 127}
                        switch (fi13 >> (((^(v12))) & 63)) {
                        case 2147483647:
                            v6 += v6
                        }
                        var v16 uint16 = 0
                        _ = v15
                        _ = v16
                    }
                }
                if (791 < 782) {
                    var v17 uint64 = 326
                    _ = v17
                } else {
                    v5 = v5
                    var v18 int16 = -32768
                    _ = v18
                }
                var v19 int32 = 1
                _ = v19
            } else {
                var v20 int32 = 669
                fmt.Println(v20, v5)
                _ = v20
            }
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v4, v6, v7)
    v1 += v1
    if (true || v2) {
        var v21 uint8 = (v1 >> ((14) & 63))
        var v22 float32 = (0.1 + float32(((v5 + ((((0.1 / 1e10) - v5) * 3.14159) / v5)) - v5)))
        _ = v21
        _ = v22
    }
    var v23 int16 = 42
    if ((float64(3.14159) > (0.1 - (((v5 / (v5 * -2.25)) * v5) * v5))) && (4294967295 <= 18446744073709551615)) {
        v23 = v23
        if v2 {
            switch v7 {
            case "foo bar":
                fmt.Println(v4)
            case "a":
                v24 := GenS1{F0: -2147483648, F1: (((-2.25 + v5) / 1e10) == v5)}
                _ = v24
            case "hello":
                for fi25 := 0; fi25 < 17; fi25++ {
                    var v26 int8 = v4
                    if (v6 < v6) {
                        v27 := GenS2{F0: int16(v6), F1: v26, F2: (fi25 + 9223372036854775807)}
                        _ = v27
                    } else {
                        fmt.Println(v6, v4, v26)
                        fmt.Println(v2, v5)
                    }
                    for fi28 := 0; fi28 < 5; fi28++ {
                        if false {
                            v29 := GenS1{F0: -1, F1: false}
                            var v30 float64 = -2.25
                            _ = v29
                            _ = v30
                        } else {
                            v4 += v4
                            var v31 float32 = 0.1
                            v26 += 1
                            fmt.Println(v1)
                            var v32 bool = (((0.0 * v5) + float64(v31)) == v5)
                            _ = v31
                            _ = v32
                        }
                        _ = fi28
                    }
                    _ = fi25
                    _ = v26
                }
            }
            switch -1 {
            case 769:
                var v33 uint8 = 128
                _ = v33
            case 4294967295:
                var v34 string = (("foo bar" + v7) + (((v7 + ("hello" + v7)) + (v7 + "foo bar")) + v7))
                _ = v34
            case -1:
                var v35 uint32 = (^(v6))
                _ = v35
            }
            v4 -= 1
        }
        switch v7 {
        case "\n\t":
            v36 := GenS2{F0: (v23 | int16(v4)), F1: (v4 | v4), F2: 4294967295}
            _ = v36
        default:
            var v37 int64 = 917
            _ = v37
        }
    }
    var v38 uint32 = (v6 | (v6 % ((420) | 1)))
    v39 := []int{32768, 543, -838}
    fmt.Println(v5)
    v7 = "hello"
    var v40 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
