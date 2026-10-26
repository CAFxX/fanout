package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint8 }

func main() {
    var v1 int = 128
    v2 := make(map[string]int)
    v3 := []uint32{}
    v2[((((("hello" + "\n\t") + "foo bar") + "hello") + "foo bar") + ("" + "\n\t"))] = (-(v1))
    for fi4 := 0; fi4 < 1; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            switch (fi4 >> ((4294967295) & 63)) {
            case -323:
                var v6 uint = 255
                _ = v6
            case -9223372036854775808:
                if (2147483647 > 0) {
                    v1 = 787
                }
            case 32767:
                v1 = fi5
            }
            var v7 float64 = 1.5
            var v8 int32 = 1
            if false {
                if false {
                    v7 = (v7 * ((1e10 / 0.1) / (v7 * (0.1 / (0.1 - (v7 * v7))))))
                    v7 -= 1e10
                    v9 := v2[("" + "\n\t")]
                    _ = v9
                }
            }
            v8 += v8
            _ = fi5
            _ = v7
            _ = v8
        }
        _ = fi4
    }
    v10 := GenS1{F0: 296}
    if (0 < 0) {
        v1 += v1
        if true {
            var v11 uint64 = 4294967295
            var v12 uint = 0
            if ((false && (255 > 1)) && true) {
                fmt.Println(v11)
                fmt.Println(v12, v12)
            }
            _ = v11
            _ = v12
        }
    } else {
        var v13 uint = 1
        _ = len(v3)
        _ = v13
    }
    switch "hello" {
    case "":
        v14 := GenS1{F0: 65535}
        _ = v14
    }
    var v15 int16 = 1
    var v16 int8 = 59
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint = 18446744073709551615
        _ = fi17
        _ = v18
    }
    v16 = (v16 >> ((1) & 63))
    if (!true) {
        fmt.Println(v16)
        if false {
            fmt.Println(v16, v1)
            var v19 uint = 42
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            _ = v19
        } else {
            if false {
                var v20 int = (v1 * 127)
                _ = v20
            } else {
                v3 = append(v3, 4294967295)
                v21 := GenS2{F0: 217}
                _ = v21
            }
        }
    }
    v16 = (-(v16))
    if true {
        fmt.Println(v15)
    }
    v22 := GenS3{F0: 615, F1: 42, F2: 234}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
}
