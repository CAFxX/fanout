package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 2
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]int)
    var v5 float32 = 3.14159
    v6 := []int{9223372036854775807}
    delete(v4, 32767)
    v3 = v3
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int16 = 42
        for fi9 := 0; fi9 < 13; fi9++ {
            v8 = (v8 & (v8 | -32768))
            v10 := GenS1{F0: 4294967295, F1: -42}
            v11 := v1[337]
            _ = fi9
            _ = v10
            _ = v11
        }
        v1[0] = -2147483648
        _ = fi7
        _ = v8
    }
    var v12 uint64 = 142
    for fi13 := 0; fi13 < 9; fi13++ {
        switch ("hello" + ("" + "hello")) {
        case "":
            v5 = ((0.0 + 3.14159) - v5)
        default:
            var v14 uint32 = 4294967295
            _ = v14
        }
        v15 := GenS1{F0: 1, F1: v3}
        _ = fi13
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := GenS1{F0: 344, F1: v3}
        for fi18 := 0; fi18 < 13; fi18++ {
            fmt.Println(fi18, fi18)
            v19 := GenS1{F0: 486, F1: v3}
            fmt.Println(fi16)
            if (v5 >= v5) {
                var v20 float32 = (v5 - 3.14159)
                if true {
                    var v21 string = ""
                    if (!((4294967295 == 2147483647) || true)) {
                        v22 := GenS2{F0: -1}
                        var v23 float32 = v20
                        _ = len(v6)
                        _ = v22
                        _ = v23
                    }
                    v6 = append(v6, 0)
                    v5 += v20
                    v3 = v3
                    _ = v21
                }
                v24 := v4[fi18]
                v25 := GenS3{F0: fi18, F1: (v12 | v12)}
                var v26 uint64 = v12
                delete(v1, 2)
                _ = v20
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v27 uint32 = 2147483648
    if (-1 == int32(v27)) {
        if true {
            if (-33 < -444) {
                var v28 uint32 = 0
                for fi29 := 0; fi29 < 3; fi29++ {
                    v30 := GenS2{F0: -394}
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        v5 += v5
                    }
                    switch ((("hello" + "foo bar") + "foo bar") + "hello") {
                    case "hello":
                        v12 = (^(v12))
                    case "a":
                        v31 := GenS4{F0: v2, F1: 0}
                        _ = v31
                    }
                    _ = fi29
                    _ = v30
                }
                _ = v28
            }
            var v32 int32 = -536
            _ = v32
        }
    }
    fmt.Println(v2)
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    if ((v12 > (^(v12))) || false) {
        v33 := GenS5{F0: int16(v12)}
        _ = v33
    }
    v2 -= v2
    v34 := GenS6{F0: v3, F1: 615, F2: 1}
    var v35 int64 = (v3 << ((549) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
