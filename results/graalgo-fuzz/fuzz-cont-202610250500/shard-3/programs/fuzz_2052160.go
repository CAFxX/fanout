package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint32; F2 int32 }
type GenS4 struct { F0 int8; F1 int32; F2 string }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t", "\n\t"}
    var v3 uint8 = 128
    v4 := []int{}
    var v5 string = "a"
    var v6 float32 = 3.14159
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v7 := GenS1{F0: 255, F1: -9223372036854775808, F2: 1000}
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int8 = 127
        v10 := GenS1{F0: 32767, F1: 691, F2: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint32 = 4294967295
        if (-128 == 80) {
            var v13 uint64 = 1
            fmt.Println(fi11, v5)
            _ = v13
        } else {
            v14 := GenS2{F0: -128}
            var v15 float64 = (3.14159 - float64((((0.1 + 0.1) + (((1e10 - -2.25) * float32((-2.25 - 1e10))) - 0.0)) + v6)))
            v16 := GenS3{F0: 1000, F1: (v12 | (v12 & uint32(v3))), F2: -574}
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi11
        _ = v12
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 int16 = 1
        if (!(1 != -824)) {
            v5 = v5
            for fi19 := 0; fi19 < 19; fi19++ {
                v5 = ("a" + v5)
                _ = fi19
            }
            _ = len(v4)
        }
        if false {
            var v20 uint16 = 65535
            _ = v20
        } else {
            for fi21 := 0; fi21 < 11; fi21++ {
                if (!true) {
                    var v22 string = v5
                    _ = v22
                }
                switch v5 {
                case "":
                    v4 = append(v4, 314)
                }
                _ = fi21
            }
            v1[0] = false
        }
        _ = fi17
        _ = v18
    }
    var v23 uint = 18
    var v24 int16 = -32768
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 uint32 = 0
        v24 = (v24 ^ 33)
        for fi27 := 0; fi27 < 8; fi27++ {
            v26 = (v26 - (^(v26)))
            var v28 int16 = v24
            fmt.Println(fi27)
            v24 = v24
            v24 = (v24 - (v28 * v28))
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    var v29 uint8 = (v3 << ((31) & 63))
    v30 := GenS4{F0: -1, F1: -2147483648, F2: ("" + v5)}
    switch "hello" {
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "":
        var v31 int8 = 100
        _ = v31
    case "foo bar":
        for fi32 := 0; fi32 < 7; fi32++ {
            v5 = ((v5 + "\n\t") + v5)
            v4 = append(v4, -526)
            fmt.Println(v3, v29)
            _ = fi32
        }
    default:
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 uint = (v23 | (v23 ^ (v23 << ((18446744073709551615) & 63))))
            var v35 uint32 = 98
            if false {
                v23 = (v23 ^ (v23 ^ v34))
            }
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    v36 := GenS2{F0: -1}
    v5 = (v5 + "a")
    v37 := GenS3{F0: v24, F1: 0, F2: 42}
    var v38 uint64 = uint64(v29)
    if false {
        var v39 float32 = float32(0.0)
        var v40 uint32 = 999
        _ = v39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
