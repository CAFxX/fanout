package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint; F1 int16 }

func main() {
    var v1 int = -9223372036854775808
    var v2 int8 = 0
    var v3 uint16 = 65535
    var v4 bool = (2 > 194)
    var v5 string = ("foo bar" + ("" + "foo bar"))
    var v6 float64 = 0.0
    if v4 {
        v7 := GenS1{F0: v2}
        if true {
            v1 = -2147483648
            switch ("foo bar" + "hello") {
            case "foo bar":
                v8 := GenS2{F0: 1000, F1: 42, F2: 42}
                _ = v8
            }
        }
        _ = v7
    }
    v5 = v5
    var v9 int64 = 102
    if v4 {
        v4 = (255 <= 185)
        if v4 {
            if (!(-32768 < 200)) {
                v6 += float64((0.0 + 1e10))
                v1 += v1
                fmt.Println(v6, v5, v6)
            } else {
                var v10 int = (v1 << ((18446744073709551615) & 63))
                _ = v10
            }
            fmt.Println(v2)
            switch 1 {
            case 42:
                var v11 float64 = ((((v6 + 1.5) + v6) - -2.25) - v6)
                _ = v11
            case 9223372036854775807:
                var v12 string = (v5 + v5)
                _ = v12
            }
        }
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        if (255 == 128) {
            var v14 int64 = v9
            _ = v14
        } else {
            v15 := []int{(fi13 ^ v1), (v1 % ((127) | 1))}
            _ = v15
        }
        v16 := GenS2{F0: 32767, F1: 2, F2: (v2 / (((v2 * v2)) | 1))}
        v17 := GenS3{F0: 0, F1: 32767}
        _ = fi13
        _ = v16
        _ = v17
    }
    v9 += v9
    if false {
        v1 = (^(v1))
        v9 += v9
        for fi18 := 0; fi18 < 11; fi18++ {
            v5 = (v5 + "a")
            v19 := []int{127, (-(fi18)), 0, v1}
            _ = fi18
            _ = v19
        }
    }
    v6 += (v6 - (1.5 / 1.0))
    var v20 uint = 4294967295
    v4 = (false || false)
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 float64 = ((v6 - v6) + v6)
        var v23 uint64 = 0
        v24 := []int{}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    var v25 int32 = 166
    var v26 uint8 = 0
    v6 += 0.0
    v6 += v6
    v20 = v20
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
