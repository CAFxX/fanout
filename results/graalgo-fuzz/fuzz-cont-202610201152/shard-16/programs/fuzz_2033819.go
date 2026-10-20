package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int64; F2 uint8 }
func genFn0(p0 uint8, p1 uint, p2 uint64) string {
    v4 := GenS1{F0: false, F1: p1, F2: 1000}
    switch (("a" + "foo bar") + "hello") {
    case "":
        v5 := GenS2{F0: 1}
        _ = v5
    case "hello":
        v6 := GenS1{F0: (!false), F1: p1, F2: -32768}
        _ = v6
    case "a":
        p2 = 42
    default:
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint16 = 255
            fmt.Println(p2, p1, p0)
            _ = fi7
            _ = v8
        }
    }
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return "foo bar"
}

func main() {
    var v1 int16 = -304
    var v2 string = "foo bar"
    v3 := make(map[string]string)
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(v2, v2, v2)
        v1 += (v1 >> ((4294967295) & 63))
        _ = fi10
    }
    var v11 bool = true
    var v12 uint8 = 202
    v1 -= (v1 << ((780) & 63))
    if (32767 == -190) {
        if (-415 == 9223372036854775807) {
            v11 = true
        }
    } else {
        switch -289 {
        case -9223372036854775808:
            for fi13 := 0; fi13 < 15; fi13++ {
                v11 = true
                v14 := v3[v2]
                _ = fi13
                _ = v14
            }
        case -496:
            for fi15 := 0; fi15 < 1; fi15++ {
                fmt.Println(v12, fi15, v11)
                _ = fi15
            }
        }
    }
    if v11 {
        for fi16 := 0; fi16 < 1; fi16++ {
            if v11 {
                var v17 bool = (4294967295 <= 42)
                var v18 uint8 = (v12 ^ (^(v12)))
                for fi19 := 0; fi19 < 11; fi19++ {
                    v20 := []int{fi16, -781, fi16}
                    switch v2 {
                    case "\n\t":
                        v21 := GenS2{F0: 1}
                        _ = v21
                    }
                    v22 := GenS2{F0: 0}
                    _ = fi19
                    _ = v20
                    _ = v22
                }
                _ = v17
                _ = v18
            } else {
                fmt.Println(v1)
                v23 := GenS2{F0: 0}
                _ = v23
            }
            v24 := GenS1{F0: (v11 || (v11 || (!(!(v1 < v1))))), F1: 65535, F2: v1}
            _ = fi16
            _ = v24
        }
    }
    if (!true) {
        v25 := GenS3{F0: (v1 - 895), F1: 986, F2: (v12 | 128)}
        _ = v25
    } else {
        v26 := v3[((v2 + v2) + ("hello" + v2))]
        for fi27 := 0; fi27 < 19; fi27++ {
            var v28 float64 = float64(-2.25)
            var v29 float32 = 1.5
            _ = fi27
            _ = v28
            _ = v29
        }
        var v30 uint = 18446744073709551615
        _ = v26
        _ = v30
    }
    var v31 uint32 = 612
    for fi32 := 0; fi32 < 6; fi32++ {
        var v33 int32 = 1
        var v34 int16 = int16(fi32)
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v31
}
