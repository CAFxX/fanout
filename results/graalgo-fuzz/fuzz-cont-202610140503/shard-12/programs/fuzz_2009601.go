package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint32; F2 int32 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
func genFn0(p0 bool) int64 {
    var v5 int = 32767
    if (2147483647 >= v5) {
        v5 += v5
        v6 := []int{}
        _ = len(v6)
        _ = v6
    }
    var v7 uint64 = 2
    _ = p0
    _ = v5
    _ = v7
    return 42
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]string)
    v3 := []uint32{669, 2147483647, 275}
    v4 := []string{("hello" + "\n\t")}
    for fi8 := 0; fi8 < 11; fi8++ {
        v1 = v1
        var v9 string = "foo bar"
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: v1}
    switch 834 {
    case -1:
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 string = "hello"
            var v13 uint = 18446744073709551615
            v1 = v1
            _ = fi11
            _ = v12
            _ = v13
        }
    case -2147483648:
        var v14 bool = true
        _ = v14
    case 42:
        if false {
            for fi15 := 0; fi15 < 2; fi15++ {
                fmt.Println(fi15, v1)
                _ = fi15
            }
        } else {
            v16 := GenS2{F0: -2147483648, F1: 65535, F2: -1}
            for fi17 := 0; fi17 < 8; fi17++ {
                if true {
                    var v18 float32 = ((float32(((-2.25 - ((1.5 / 0.1) / (1e10 * 0.1))) * (float64(3.14159) + ((1.5 - (1e10 * -2.25)) + 1.5)))) / (1e10 / (0.0 + 0.1))) + 0.1)
                    if (v1 == (v1 - (v1 & 529))) {
                        fmt.Println(v18, v1)
                        fmt.Println(v1)
                        v19 := GenS2{F0: 567, F1: 1, F2: 716}
                        _ = v19
                    }
                    v20 := GenS1{F0: 4294967295}
                    _ = v18
                    _ = v20
                }
                if true {
                    delete(v2, 1)
                }
                v21 := GenS3{F0: -53, F1: 1}
                _ = fi17
                _ = v21
            }
            switch 512 {
            case -2147483648:
                var v22 int64 = -339
                _ = v22
            }
            _ = v16
        }
    }
    if (1.5 < (((1e10 + 3.14159) * (float64(1e10) / (3.14159 + float64(-2.25)))) + ((3.14159 * float64(1.5)) / 0.1))) {
        var v23 string = "foo bar"
        switch v23 {
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
        }
        _ = v23
    } else {
        switch 1 {
        case 128:
            v1 = (^(v1))
        }
        fmt.Println(v1)
    }
    var v24 uint = 0
    var v25 uint = 0
    switch 42 {
    case -1:
        v26 := GenS4{F0: 9223372036854775807, F1: "foo bar"}
        _ = v26
    case 4294967295:
        switch "\n\t" {
        case "\n\t":
            if (true || ((1 <= 0) && true)) {
                v1 = (v1 << (((v25 | (v25 | 4294967295))) & 63))
                var v27 int8 = -41
                _ = v27
            } else {
                v25 = v24
            }
        }
    default:
        v28 := GenS5{F0: -2147483648}
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 uint16 = 114
        v24 -= (^(v24))
        _ = fi29
        _ = v30
    }
    v1 -= 1
    v24 -= 4294967295
    if false {
        var v31 uint16 = 255
        _ = v31
    } else {
        var v32 int32 = 42
        for fi33 := 0; fi33 < 6; fi33++ {
            fmt.Println(v1, v1)
            _ = fi33
        }
        if false {
            for fi34 := 0; fi34 < 16; fi34++ {
                v32 -= (v32 % ((v32) | 1))
                _ = fi34
            }
            v24 = (v24 % (((^(v25))) | 1))
            var v35 uint64 = v1
            _ = v35
        }
        _ = v32
    }
    var v36 uint32 = 2147483648
    switch 32768 {
    case 9223372036854775807:
        if ((!(true && (45 >= -1))) && true) {
            v1 = v1
            var v37 uint32 = 2
            _ = v37
        } else {
            var v38 uint64 = v1
            _ = v38
        }
    case 804:
        v39 := GenS6{F0: 1}
        _ = v39
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        v1 = v1
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    _ = len(v4)
    v25 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v4
}
