package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int16; F2 int }
type GenS4 struct { F0 string; F1 bool; F2 uint16 }
func genFn0(p0 string, p1 int, p2 uint64) uint {
    if true {
        fmt.Println(p1, p2)
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint16 = 42
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    return 716
}

func main() {
    v1 := make(map[string]int)
    var v2 int = -2147483648
    v3 := make(map[string]string)
    var v4 int16 = -1
    var v5 int64 = 1
    var v6 int = int(v4)
    if (0 <= 42) {
        fmt.Println(v5)
    } else {
        v10 := GenS1{F0: 2147483648, F1: 42}
        if (-30 <= v2) {
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 float64 = 0.0
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 int32 = 1
                    switch (v2 - int(v5)) {
                    case 32768:
                        fmt.Println(v6, v5)
                    case -576:
                        fmt.Println(v2)
                    default:
                        var v15 uint8 = 128
                        _ = v15
                    }
                    _ = fi13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
            v16 := GenS1{F0: 4294967295, F1: -45}
            fmt.Println(v5, v6)
            _ = v16
        } else {
            v17 := GenS1{F0: 2, F1: 127}
            switch (v6 * (v6 - v2)) {
            case -1:
                v6 = -309
            case 0:
                v18 := GenS1{F0: 0, F1: 0}
                _ = v18
            case -2147483648:
                v19 := GenS2{F0: 32768}
                _ = v19
            }
            if true {
                for fi20 := 0; fi20 < 19; fi20++ {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        v5 += -946
                        _ = fi21
                    }
                    _ = fi20
                }
            }
            _ = v17
        }
        _ = v10
    }
    var v22 uint = 2
    v22 = v22
    if true {
        var v23 uint8 = 2
        _ = v23
    } else {
        for fi24 := 0; fi24 < 4; fi24++ {
            _ = len(v1)
            _ = fi24
        }
        v5 = v5
        switch (v2 % (((v6 >> ((v22) & 63))) | 1)) {
        case -361:
            switch "hello" {
            case "foo bar":
                v2 -= (^(v6))
            }
        case 2147483647:
            if false {
                v2 = (v6 % ((787) | 1))
                switch "a" {
                case "a":
                    v3["hello"] = ""
                case "hello":
                    v25 := []int{}
                    _ = v25
                case "\n\t":
                    v5 = (v5 >> (((v22 * v22)) & 63))
                }
            } else {
                if ((v22 | (v22 + 936)) == v22) {
                    fmt.Println(v22, v2)
                }
                if (2 < 255) {
                    v26 := GenS3{F0: 2, F1: -32768, F2: -9223372036854775808}
                    _ = v26
                }
                delete(v1, "a")
            }
        case 42:
            _ = len(v3)
        default:
            var v27 float32 = 1.5
            _ = v27
        }
    }
    v5 = v5
    var v28 int16 = 32767
    var v29 string = "a"
    v30 := v1["a"]
    v31 := GenS1{F0: 2147483648, F1: 42}
    var v32 uint = v22
    fmt.Println(v22, v29, v2)
    v33 := GenS4{F0: "foo bar", F1: (0.0 >= 0.0), F2: 360}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
