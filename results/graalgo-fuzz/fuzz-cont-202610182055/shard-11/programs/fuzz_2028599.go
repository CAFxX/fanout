package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []uint32{}
    v2 := make(map[int]uint32)
    var v3 int32 = 1
    var v4 uint = 4294967295
    v4 += v4
    for fi5 := 0; fi5 < 2; fi5++ {
        v3 = (v3 % (((^(v3))) | 1))
        var v6 uint64 = 0
        for fi7 := 0; fi7 < 5; fi7++ {
            if true {
                v3 += (v3 >> (((v4 ^ v4)) & 63))
            }
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    if (false && ((0.0 - 1.5) <= 3.14159)) {
        var v8 uint32 = 114
        var v9 int32 = (v3 / ((0) | 1))
        _ = v8
        _ = v9
    } else {
        var v10 bool = (255 <= 65535)
        for fi11 := 0; fi11 < 10; fi11++ {
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 float32 = 0.1
                _ = fi12
                _ = v13
            }
            if len(v1) > 0 {
                v1[0] = 793
            }
            _ = fi11
        }
        v14 := GenS1{F0: 255, F1: 9223372036854775807}
        _ = v10
        _ = v14
    }
    if false {
        var v15 string = "a"
        _ = v15
    }
    var v16 uint = (v4 >> ((v4) & 63))
    if false {
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 int64 = 9223372036854775807
            fmt.Println(v3)
            _ = fi17
            _ = v18
        }
        if ((("\n\t" + "hello") + ((("a" + "hello") + ("" + "a")) + "foo bar")) != "a") {
            switch 0 {
            case 0:
                v4 = v16
            }
            fmt.Println(v4)
            fmt.Println(v16, v3)
        } else {
            fmt.Println(v3, v3, v4)
        }
        var v19 float32 = 1.5
        _ = v19
    } else {
        _ = len(v2)
        for fi20 := 0; fi20 < 1; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                v22 := GenS2{F0: 101}
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            fmt.Println(fi23, v3, v16)
            var v24 int16 = -624
            _ = fi23
            _ = v24
        }
    }
    v1 = append(v1, 38)
    var v25 int64 = 9223372036854775807
    var v26 uint32 = 593
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
