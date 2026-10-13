package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int64; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint; F1 int32 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]uint32)
    var v3 int8 = -102
    var v4 bool = false
    v5 := make(map[int]bool)
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(v3)
        if v4 {
            switch 429 {
            case 2147483647:
                v1 = v1
            case 32767:
                fmt.Println(v3, v3, v6)
            default:
                v8 := GenS1{F0: (v1 ^ (v1 ^ (v1 | (^(v1)))))}
                _ = v8
            }
            if (true || false) {
                v3 += -128
                var v9 uint32 = 42
                var v10 string = "hello"
                _ = v10
                _ = v9
            }
        }
        if v4 {
            fmt.Println(v4, v6, v1)
            if true {
                var v11 int8 = 1
                _ = v11
            } else {
                _ = len(v5)
                var v12 int8 = (v3 * 127)
                fmt.Println(v4)
                _ = v12
            }
            if v4 {
                if true {
                    v13 := GenS2{F0: "hello"}
                    v6 = (v6 >> ((v1) & 63))
                    v6 += v6
                    _ = v13
                }
                v14 := []int{fi7, (fi7 / ((fi7) | 1)), fi7, (fi7 << ((v1) & 63))}
                v15 := GenS1{F0: (v1 << (((v1 & v1)) & 63))}
                _ = v14
                _ = v15
            }
        } else {
            fmt.Println(fi7, v3, v3)
            fmt.Println(fi7, v4)
        }
        _ = fi7
    }
    fmt.Println(v4, v6, v6)
    var v16 uint64 = 518
    v1 -= 4294967295
    if v4 {
        var v17 float64 = (1.5 / 1.0)
        v18 := GenS2{F0: "hello"}
        var v19 int64 = 42
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 int32 = 0
    var v21 float32 = 0.1
    var v22 int64 = 422
    var v23 uint32 = 2147483648
    for fi24 := 0; fi24 < 12; fi24++ {
        switch -312 {
        case 32768:
            fmt.Println(v23, v23, v4)
        case 4294967295:
            var v25 int8 = (-(v3))
            _ = v25
        default:
            v26 := GenS3{F0: 42, F1: 0}
            _ = v26
        }
        v27 := GenS4{F0: (v22 ^ (v22 - v22)), F1: -32768, F2: v23}
        fmt.Println(v21, v20, v22)
        _ = fi24
        _ = v27
    }
    if ((v22 - (-(v22))) == (v22 ^ (v22 >> ((v1) & 63)))) {
        var v28 uint64 = v16
        v2["foo bar"] = (v23 * (v23 + (v23 - (v23 << ((65535) & 63)))))
        _ = v28
    } else {
        if ((42 < 65535) || true) {
            if v4 {
                for fi29 := 0; fi29 < 2; fi29++ {
                    var v30 float32 = float32(0.1)
                    var v31 bool = (!(v3 != v3))
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            } else {
                if v4 {
                    if (!true) {
                        _ = len(v5)
                    } else {
                        _ = len(v2)
                    }
                }
                fmt.Println(v4)
            }
            v1 += (^(v1))
            switch -1 {
            case 0:
                v23 -= 2147483647
            default:
                switch ("hello" + (("" + ("foo bar" + "\n\t")) + "")) {
                case "":
                    fmt.Println(v22, v6)
                default:
                    v3 += -128
                }
            }
        } else {
            for fi32 := 0; fi32 < 13; fi32++ {
                v1 -= v1
                _ = fi32
            }
            v4 = ((!(!(int32(v22) >= (v20 + (v20 | (v20 * v20)))))) || v4)
            _ = len(v2)
        }
    }
    v21 += 1.5
    for fi33 := 0; fi33 < 5; fi33++ {
        v34 := GenS3{F0: (v20 & (v20 | v20)), F1: 32767}
        v4 = false
        v35 := []int{(fi33 ^ fi33)}
        _ = fi33
        _ = v34
        _ = v35
    }
    switch "foo bar" {
    case "\n\t":
        v20 -= v20
    case "hello":
        switch "a" {
        case "foo bar":
            v36 := []int{1}
            _ = v36
        case "\n\t":
            if (!(997 <= 2)) {
                v37 := GenS5{F0: uint(v23), F1: v20}
                fmt.Println(v3)
                _ = v37
            }
        }
    case "foo bar":
        v22 = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
