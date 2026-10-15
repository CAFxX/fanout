package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int16; F1 string }
func genFn0(p0 int64) uint8 {
    for fi5 := 0; fi5 < 18; fi5++ {
        p0 = p0
        _ = fi5
    }
    v6 := []int{}
    var v7 int = 32768
    _ = p0
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 string = "\n\t"
    var v3 int32 = 926
    v4 := []string{}
    v8 := GenS1{F0: 175, F1: 0}
    if (v3 < (v3 ^ (v3 % ((1) | 1)))) {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int16 = -1
            _ = fi9
            _ = v10
        }
        if (1 >= 0) {
            switch ((v2 + "") + v2) {
            case "a":
                if false {
                    v11 := v1[v2]
                    _ = v11
                }
            case "foo bar":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "\n\t":
                var v12 uint64 = 4294967295
                _ = v12
            }
            v3 -= v3
        }
        v2 = v2
    } else {
        if len(v4) > 0 {
            v4[0] = v2
        }
        switch (v2 + "hello") {
        case "":
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 uint16 = 0
                if false {
                    v2 = ("" + "a")
                }
                _ = fi13
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 string = ""
            v2 = "a"
            var v17 uint64 = 42
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    switch "foo bar" {
    case "a":
        _ = len(v1)
    }
    v2 = v2
    v2 = "hello"
    switch 49 {
    case -2147483648:
        var v18 int16 = -664
        _ = v18
    default:
        var v19 int8 = -63
        _ = v19
    }
    var v20 int32 = (v3 & -2147483648)
    var v21 float64 = 0.0
    for fi22 := 0; fi22 < 12; fi22++ {
        v3 = (v3 + (v3 % ((128) | 1)))
        var v23 uint16 = 480
        if (!true) {
            v24 := GenS1{F0: 42, F1: -398}
            _ = v24
        } else {
            var v25 bool = (255 == 557)
            if true {
                var v26 int = (fi22 * fi22)
                if (1e10 == 1.5) {
                    switch v2 {
                    case "hello":
                        v21 -= (float64((1.5 - (-2.25 + 1e10))) / ((v21 - 3.14159) + v21))
                    case "a":
                        for fi27 := 0; fi27 < 20; fi27++ {
                            v28 := GenS2{F0: 9223372036854775807, F1: false}
                            _ = fi27
                            _ = v28
                        }
                    }
                    v25 = (false || (!(-2.25 > 1.5)))
                    fmt.Println(v2, v23)
                }
                v25 = (!(!true))
                _ = v26
            }
            v29 := GenS1{F0: (v23 - v23), F1: -554}
            _ = v25
            _ = v29
        }
        _ = fi22
        _ = v23
    }
    if (0.1 != (-2.25 + float32((v21 - v21)))) {
        for fi30 := 0; fi30 < 4; fi30++ {
            v20 = v20
            var v31 uint16 = 32767
            _ = fi30
            _ = v31
        }
        _ = len(v1)
    }
    if (100 == -1) {
        switch 32768 {
        case -1:
            v32 := GenS3{F0: -32768, F1: v2}
            _ = v32
        case -9223372036854775808:
            var v33 int16 = -32768
            _ = v33
        case -683:
            for fi34 := 0; fi34 < 19; fi34++ {
                v21 = (v21 - v21)
                for fi35 := 0; fi35 < 7; fi35++ {
                    var v36 int16 = 656
                    _ = fi35
                    _ = v36
                }
                _ = fi34
            }
        }
        var v37 string = v2
        _ = v37
    } else {
        var v38 uint32 = 2147483648
        _ = v38
    }
    if (!(1 <= -1)) {
        switch 4294967295 {
        case 4294967295:
            var v39 float64 = (float64((1.5 - (0.0 * float32(((float64(1e10) + (v21 - (0.0 - (3.14159 / (v21 - v21))))) - 1.5))))) / 1.0)
            _ = v39
        case 127:
            for fi40 := 0; fi40 < 10; fi40++ {
                fmt.Println(v3, v20)
                _ = fi40
            }
        case 0:
            switch "a" {
            case "hello":
                fmt.Println(v21, v20, v21)
            case "a":
                switch 527 {
                case 63:
                    var v41 string = (((v2 + "hello") + v2) + v2)
                    _ = v41
                }
            case "\n\t":
                var v42 uint64 = 42
                _ = v42
            default:
                var v43 float32 = (float32((float64(1.5) - v21)) / float32(float64((-2.25 - 1e10))))
                _ = v43
            }
        default:
            var v44 uint = 301
            _ = v44
        }
    }
    fmt.Println(v2)
    v21 -= v21
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v8
}
