package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint16 = 42
    var v2 bool = (9223372036854775807 < -704)
    var v3 uint32 = 1
    var v4 string = "hello"
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            if (v1 > v1) {
                var v7 float32 = (1e10 * (0.0 * 0.1))
                v8 := GenS1{F0: "", F1: "a", F2: 127}
                v1 = v1
                _ = v7
                _ = v8
            } else {
                fmt.Println(v2, v4, fi5)
                v9 := GenS1{F0: v4, F1: v4, F2: 255}
                _ = v9
            }
            var v10 int64 = 883
            _ = fi6
            _ = v10
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 string = "foo bar"
            v4 = v4
            var v13 bool = v2
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi5
    }
    v14 := GenS1{F0: (((("hello" + v4) + "hello") + v4) + ("hello" + "a")), F1: (((v4 + v4) + "\n\t") + "a"), F2: 50}
    v1 -= (v1 * (v1 * (v1 - v1)))
    v15 := []int{int(v3), 128}
    var v16 uint8 = 128
    var v17 float64 = (-2.25 - 1.5)
    if (true || (543 >= v1)) {
        if false {
            v18 := GenS2{F0: 310, F1: 128}
            _ = v18
        }
        if (!v2) {
            var v19 uint64 = 282
            _ = v19
        } else {
            v3 = v3
            if true {
                for fi20 := 0; fi20 < 4; fi20++ {
                    var v21 uint32 = (^(v3))
                    _ = fi20
                    _ = v21
                }
                v16 += 31
                if (!true) {
                    v22 := GenS2{F0: -32768, F1: 32767}
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v24 := GenS2{F0: -19, F1: 127}
                        var v25 int32 = 0
                        if len(v15) > 0 {
                            v15[0] = (fi23 << ((1) & 63))
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = len(v15)
                    _ = v22
                } else {
                    var v26 int64 = -542
                    v27 := GenS2{F0: -657, F1: 9223372036854775807}
                    _ = v26
                    _ = v27
                }
            } else {
                v4 = "\n\t"
            }
        }
        v16 -= v16
    } else {
        var v28 int16 = 987
        v29 := GenS3{F0: 914}
        v28 += v28
        _ = v28
        _ = v29
    }
    v30 := GenS4{F0: 65535}
    fmt.Println(v1)
    v16 -= v16
    var v31 uint32 = (v3 % ((v3) | 1))
    switch 1 {
    case 32767:
        var v32 int16 = 32767
        _ = v32
    case 2147483647:
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
