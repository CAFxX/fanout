package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    var v4 bool = true
    var v5 int32 = -1
    v6 := []uint32{4294967295, 1, 2147483648}
    for fi7 := 0; fi7 < 2; fi7++ {
        switch "hello" {
        case "foo bar":
            v8 := v2["foo bar"]
            _ = v8
        case "hello":
            switch 4294967295 {
            case 0:
                var v9 int = (fi7 / ((14) | 1))
                _ = v9
            default:
                var v10 uint8 = 238
                _ = v10
            }
        case "":
            if (!true) {
                var v11 int8 = -1
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 int = fi7
                    var v14 bool = v4
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = v11
            }
        default:
            v15 := GenS1{F0: 867, F1: "a"}
            _ = v15
        }
        _ = fi7
    }
    _ = len(v2)
    var v16 int32 = 2147483647
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 uint8 = 128
        for fi19 := 0; fi19 < 13; fi19++ {
            v16 -= v5
            var v20 int64 = -1
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := GenS1{F0: -292, F1: "foo bar"}
        v23 := GenS2{F0: 1}
        if true {
            if false {
                for fi24 := 0; fi24 < 2; fi24++ {
                    fmt.Println(v4, v5, fi24)
                    fmt.Println(fi24, fi24)
                    _ = fi24
                }
            }
            var v25 uint16 = 540
            var v26 uint8 = 0
            _ = v25
            _ = v26
        } else {
            for fi27 := 0; fi27 < 19; fi27++ {
                v4 = ((3.14159 + (1e10 - 0.1)) <= 1e10)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi27
            }
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    if v4 {
        v6 = append(v6, 1)
        if len(v6) > 0 {
            v6[0] = 0
        }
    } else {
        v5 -= -178
    }
    v16 = (v5 / (((v16 - -535)) | 1))
    var v28 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
