package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int8; F2 int }

func main() {
    v1 := make(map[int]int)
    v2 := []int{32768}
    v3 := make(map[string]string)
    var v4 int16 = 0
    v2 = append(v2, 9223372036854775807)
    var v5 uint8 = 71
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v6 int64 = 9223372036854775807
        _ = v6
    } else {
        var v7 int64 = 0
        _ = v7
    }
    var v8 float32 = 1e10
    for fi9 := 0; fi9 < 12; fi9++ {
        switch fi9 {
        case -79:
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 int8 = 100
                _ = fi10
                _ = v11
            }
        case 9223372036854775807:
            switch fi9 {
            case 127:
                var v12 uint32 = 65535
                _ = v12
            }
        default:
            var v13 int16 = -1
            _ = v13
        }
        if true {
            v8 = 3.14159
            var v14 uint16 = 42
            _ = v14
        }
        var v15 uint8 = (v5 % ((127) | 1))
        _ = fi9
        _ = v15
    }
    var v16 int64 = -9223372036854775808
    for fi17 := 0; fi17 < 18; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            v19 := GenS1{F0: 0, F1: (fi18 | (fi17 >> ((4294967295) & 63)))}
            var v20 uint32 = 1
            switch "\n\t" {
            case "foo bar":
                var v21 float32 = -2.25
                _ = v21
            case "hello":
                var v22 int64 = -950
                _ = v22
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    switch 127 {
    case 32768:
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 float32 = v8
            delete(v3, ("\n\t" + ("\n\t" + "foo bar")))
            fmt.Println(v24)
            _ = fi23
            _ = v24
        }
    case 1:
        v8 += (0.1 * ((v8 * v8) / 0.1))
    default:
        var v25 bool = (true || true)
        _ = v25
    }
    v5 += (v5 + (v5 | (v5 << ((786) & 63))))
    v26 := GenS2{F0: true}
    switch 1 {
    case -573:
        v27 := v3[(("\n\t" + "\n\t") + (((("a" + "foo bar") + ("" + "hello")) + (("foo bar" + "hello") + ("" + ""))) + ("hello" + (("hello" + "") + "hello"))))]
        _ = v27
    default:
        if false {
            for fi28 := 0; fi28 < 16; fi28++ {
                var v29 string = ("\n\t" + "")
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi30 := 0; fi30 < 5; fi30++ {
                    v31 := GenS2{F0: (v29 != ("\n\t" + "a"))}
                    v32 := GenS2{F0: (!false)}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = fi28
                _ = v29
            }
            v33 := GenS3{F0: 2, F1: 23, F2: -9223372036854775808}
            _ = v33
        }
    }
    v8 += (1.5 + v8)
    v8 += v8
    for fi34 := 0; fi34 < 16; fi34++ {
        for fi35 := 0; fi35 < 1; fi35++ {
            if len(v2) > 0 {
                v2[0] = (fi35 ^ (fi34 / ((42) | 1)))
            }
            var v36 bool = true
            v3["foo bar"] = "hello"
            _ = fi35
            _ = v36
        }
        v8 = v8
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
