package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0(p0 int32, p1 int32, p2 uint8) int16 {
    fmt.Println(p0)
    if true {
        v7 := []int{4294967295, 42, -1, 4294967295}
        _ = v7
    } else {
        v8 := GenS1{F0: 0}
        var v9 int32 = (-(p0))
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := []int{fi10}
        switch "" {
        case "\n\t":
            var v12 int8 = 100
            _ = v12
        }
        var v13 string = "hello"
        _ = fi10
        _ = v11
        _ = v13
    }
    v14 := GenS2{F0: 1, F1: 218}
    fmt.Println(p2, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 790
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 777
    var v3 uint32 = 427
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    var v6 int = -1
    if (1 > -1) {
        v15 := v4["a"]
        v16 := []int{v6, 32768, (-(v6))}
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        if (true && (!false)) {
            for fi18 := 0; fi18 < 14; fi18++ {
                fmt.Println(v6, fi18, v6)
                if false {
                    delete(v5, v3)
                    v1[42] = "a"
                    var v19 int8 = 0
                    if true {
                        var v20 string = "hello"
                        _ = v20
                    } else {
                        v19 = v19
                        fmt.Println(v19)
                        v6 = (v6 << ((65535) & 63))
                        v2 -= v2
                        v2 += v2
                        v6 -= fi18
                    }
                    _ = v19
                }
                v21 := GenS2{F0: 27, F1: 0}
                _ = fi18
                _ = v21
            }
            v3 = (v3 - v3)
            switch (v6 | v6) {
            case -979:
                v2 = (v2 / ((556) | 1))
            default:
                var v22 float32 = 0.0
                _ = v22
            }
        } else {
            var v23 int16 = 32767
            v2 = (^(v2))
            _ = v23
        }
        delete(v1, (^(v3)))
        fmt.Println(v2, v2, v6)
        _ = fi17
    }
    v24 := GenS1{F0: v2}
    v3 = (v3 >> ((0) & 63))
    if false {
        if true {
            v25 := GenS2{F0: 18446744073709551615, F1: -32768}
            _ = v25
        }
        switch "foo bar" {
        case "a":
            var v26 int64 = -280
            _ = v26
        case "\n\t":
            var v27 int = v6
            _ = v27
        case "hello":
            for fi28 := 0; fi28 < 12; fi28++ {
                v5[v3] = true
                v5[794] = true
                switch "a" {
                case "hello":
                    for fi29 := 0; fi29 < 4; fi29++ {
                        _ = len(v4)
                        _ = fi29
                    }
                case "a":
                    v3 = 205
                }
                _ = fi28
            }
        default:
            switch ("" + "foo bar") {
            case "hello":
                var v30 string = "\n\t"
                _ = v30
            case "":
                var v31 int64 = 1
                _ = v31
            case "foo bar":
                var v32 float32 = (float32(float64(1.5)) + (3.14159 * 0.0))
                _ = v32
            default:
                if (4294967295 <= 42) {
                    for fi33 := 0; fi33 < 12; fi33++ {
                        var v34 int = (fi33 << ((306) & 63))
                        v6 -= (v6 | v34)
                        _ = fi33
                        _ = v34
                    }
                    v35 := GenS2{F0: (v2 % ((v2) | 1)), F1: -1}
                    v36 := GenS3{F0: 1, F1: 408}
                    _ = v35
                    _ = v36
                }
            }
        }
    } else {
        switch "foo bar" {
        case "\n\t":
            var v37 int64 = -1
            _ = v37
        case "a":
            var v38 uint = 42
            _ = v38
        default:
            v2 = (^(v2))
        }
    }
    v39 := []int{2147483647, v6, (v6 & (v6 ^ (v6 % ((9223372036854775807) | 1))))}
    v40 := GenS1{F0: (v2 / ((782) | 1))}
    for fi41 := 0; fi41 < 20; fi41++ {
        switch ("" + ((((("hello" + "foo bar") + (("" + "\n\t") + "a")) + ("foo bar" + (("foo bar" + "") + ("" + "hello")))) + "a") + "a")) {
        case "":
            v6 = (fi41 / (((^(fi41))) | 1))
        case "\n\t":
            v6 = -2147483648
        case "a":
            v2 += v2
        }
        v42 := GenS2{F0: v2, F1: -1}
        _ = fi41
        _ = v42
    }
    for fi43 := 0; fi43 < 4; fi43++ {
        switch v6 {
        case 158:
            v2 += (v2 * (v2 << ((1) & 63)))
        case -836:
            if len(v39) > 0 {
                v39[0] = (v6 >> ((7) & 63))
            }
        }
        var v44 uint8 = 2
        fmt.Println(v3, v2, v3)
        _ = fi43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
