package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int64; F2 int64 }

func main() {
    var v1 uint64 = 827
    v2 := make(map[int]uint32)
    var v3 int8 = 127
    v4 := []uint32{532, 4294967295, 1, 2147483647}
    switch "" {
    case "a":
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 int32 = 1
            _ = fi5
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v1 -= (v1 | 42)
        if true {
            for fi8 := 0; fi8 < 17; fi8++ {
                _ = len(v2)
                _ = fi8
            }
            v9 := GenS1{F0: "\n\t"}
            for fi10 := 0; fi10 < 3; fi10++ {
                v1 = (v1 / (((v1 & v1)) | 1))
                _ = fi10
            }
            _ = v9
        } else {
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 int64 = 9223372036854775807
                v13 := v2[fi11]
                var v14 float64 = 1.5
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            fmt.Println(v3)
        }
        _ = fi7
    }
    switch ("hello" + "") {
    case "foo bar":
        v1 = v1
    default:
        if true {
            var v15 bool = false
            switch -2147483648 {
            case 32767:
                var v16 int64 = -9223372036854775808
                _ = v16
            }
            _ = v15
        }
    }
    if false {
        if (4294967295 == -1) {
            var v17 uint8 = 0
            _ = v17
        } else {
            v2[-538] = 408
            v18 := GenS1{F0: ("\n\t" + ((("" + ("foo bar" + (("\n\t" + ("" + "foo bar")) + ("hello" + (("hello" + "a") + "hello"))))) + ("\n\t" + "hello")) + (("\n\t" + "a") + "foo bar")))}
            v1 = v1
            _ = v18
        }
        var v19 uint64 = (v1 - v1)
        _ = len(v2)
        _ = v19
    }
    var v20 int64 = 117
    v3 += v3
    fmt.Println(v20, v20)
    switch "a" {
    case "":
        fmt.Println(v3)
    default:
        v1 -= (v1 << ((255) & 63))
    }
    v21 := GenS1{F0: "\n\t"}
    for fi22 := 0; fi22 < 6; fi22++ {
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := GenS1{F0: "\n\t"}
            _ = fi23
            _ = v24
        }
        var v25 uint64 = 177
        _ = fi22
        _ = v25
    }
    v26 := GenS2{F0: v3, F1: 9223372036854775807, F2: v20}
    if ((v3 * int8(v20)) != (v3 >> ((175) & 63))) {
        v27 := GenS2{F0: int8(v1), F1: v20, F2: (v20 << ((2) & 63))}
        var v28 string = ("foo bar" + ("hello" + "a"))
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 uint8 = 128
            _ = fi29
            _ = v30
        }
        _ = v27
        _ = v28
    } else {
        v31 := GenS2{F0: v3, F1: v20, F2: (v20 << ((18446744073709551615) & 63))}
        if (732 != int64(v3)) {
            var v32 uint8 = 85
            _ = v32
        } else {
            var v33 uint16 = 65535
            _ = v33
        }
        var v34 uint64 = (v1 | 42)
        _ = v31
        _ = v34
    }
    v35 := v2[-93]
    var v36 float32 = 0.0
    var v37 uint = 254
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
