package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 string; F1 uint; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint8; F2 int16 }
func genFn0(p0 uint) uint8 {
    fmt.Println(p0)
    switch 128 {
    case 9223372036854775807:
        switch "hello" {
        case "hello":
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 uint64 = 938
                v8 -= (v8 << ((p0) & 63))
                v8 = (^(v8))
                _ = fi7
                _ = v8
            }
        case "":
            if ((^(p0)) == (p0 / (((^(p0))) | 1))) {
                v9 := []int{42, 128, 4294967295}
                _ = v9
            } else {
                var v10 uint64 = 4294967295
                v10 = uint64(p0)
                _ = v10
            }
        }
    }
    v11 := GenS1{F0: 4294967295, F1: 674}
    v12 := GenS1{F0: 100, F1: 198}
    _ = p0
    _ = v11
    _ = v12
    return 209
}
func genFn1(p0 uint16, p1 string, p2 uint16) uint64 {
    var v13 float64 = 3.14159
    if true {
        p2 = p2
    } else {
        v14 := []int{42, 366, 4294967295, -9223372036854775808}
        fmt.Println(v13, p1, p2)
        switch 127 {
        case 2147483647:
            for fi15 := 0; fi15 < 9; fi15++ {
                if len(v14) > 0 {
                    v14[0] = (fi15 << ((4294967295) & 63))
                }
                var v16 float64 = ((0.1 * v13) / (v13 / (v13 - 0.1)))
                _ = fi15
                _ = v16
            }
        case 212:
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        }
        _ = v14
    }
    v17 := GenS2{F0: ("foo bar" + ((p1 + (p1 + "\n\t")) + (p1 + "hello"))), F1: 4294967295, F2: 313}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v17
    return 4294967295
}

func main() {
    var v1 float32 = ((((float32(3.14159) / 0.1) + 1e10) * 0.0) + 0.0)
    var v2 string = "foo bar"
    v3 := make(map[uint32]string)
    v4 := []int{566, 127, -2147483648}
    var v5 int32 = 42
    var v6 uint8 = 239
    if (v6 == 97) {
        v1 = (v1 - float32(3.14159))
        for fi18 := 0; fi18 < 10; fi18++ {
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 int = fi18
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        switch v2 {
        case "":
            v3[4294967295] = "foo bar"
        }
    }
    fmt.Println(v2)
    var v21 uint = 707
    fmt.Println(v6, v21, v1)
    var v22 int = 127
    v23 := GenS3{F0: 9223372036854775807, F1: (v6 * v6), F2: 42}
    if false {
        var v24 uint16 = 2
        _ = v24
    } else {
        if true {
            delete(v3, 2147483647)
            var v25 int16 = 0
            _ = v25
        }
        v4 = append(v4, (v22 / ((4294967295) | 1)))
    }
    v22 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
