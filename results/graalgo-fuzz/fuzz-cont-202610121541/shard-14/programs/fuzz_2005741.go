package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 bool; F2 int16 }
type GenS3 struct { F0 int64; F1 int8; F2 string }
func genFn0(p0 uint, p1 uint64) uint32 {
    var v7 uint16 = 32767
    v7 = v7
    var v8 int8 = 100
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 int32, p1 uint16, p2 uint64) uint32 {
    v9 := []int{2147483647}
    p0 += p0
    if (true || true) {
        p1 = p1
        switch "\n\t" {
        case "\n\t":
            switch 4294967295 {
            case -1:
                p1 -= 255
            default:
                var v10 uint8 = 128
                _ = v10
            }
        default:
            for fi11 := 0; fi11 < 7; fi11++ {
                _ = len(v9)
                _ = fi11
            }
        }
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            switch ("a" + ("a" + "hello")) {
            case "a":
                switch ("" + ("a" + ("" + "foo bar"))) {
                case "a":
                    if false {
                        v14 := GenS1{F0: p1}
                        _ = v14
                    }
                case "":
                    var v15 uint32 = 65535
                    _ = v15
                case "foo bar":
                    v16 := GenS2{F0: 42, F1: (2 > 907), F2: int16(p2)}
                    _ = v16
                }
            }
            if (((!true) || (!true)) && false) {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                fmt.Println(fi13, fi12)
                fmt.Println(fi12, p2, p1)
                p1 -= 2
            } else {
                fmt.Println(fi13, p0, p2)
            }
            fmt.Println(fi13, p1, p0)
            fmt.Println(p1, p1)
            _ = fi13
        }
        switch 127 {
        case 0:
            var v17 int8 = -1
            _ = v17
        case 1:
            p0 = 1
        default:
            for fi18 := 0; fi18 < 13; fi18++ {
                if false {
                    p2 -= 42
                    p0 = -1
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    if true {
                        fmt.Println(p2, fi12, p1)
                    }
                    var v19 float64 = (0.0 - -2.25)
                    _ = v19
                }
                switch "foo bar" {
                case "":
                    v20 := GenS2{F0: 65535, F1: false, F2: 871}
                    _ = v20
                case "a":
                    var v21 uint64 = p2
                    _ = v21
                case "foo bar":
                    var v22 int64 = -9223372036854775808
                    _ = v22
                default:
                    p2 = (p2 + p2)
                }
                _ = fi18
            }
        }
        v9 = append(v9, 877)
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return genFn0(255, 405)
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 0
    v3 := make(map[uint32]int)
    var v4 int32 = 0
    v5 := make(map[uint32]bool)
    v6 := []string{"\n\t", "\n\t"}
    var v23 uint16 = 582
    if (!(156 == 42)) {
        v24 := GenS1{F0: 255}
        _ = v24
    } else {
        for fi25 := 0; fi25 < 11; fi25++ {
            v26 := GenS1{F0: v23}
            _ = fi25
            _ = v26
        }
    }
    v6 = append(v6, "a")
    if false {
        v4 -= (v4 / ((-2147483648) | 1))
        v27 := v3[v2]
        _ = v27
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        switch "hello" {
        case "foo bar":
            _ = len(v6)
        case "hello":
            var v29 int64 = 0
            _ = v29
        case "\n\t":
            v4 = v4
        default:
            v30 := GenS2{F0: 42, F1: true, F2: -300}
            _ = v30
        }
        if (!(-1 != fi28)) {
            v6 = append(v6, "\n\t")
        } else {
            v31 := GenS3{F0: -1, F1: 42, F2: ("a" + "foo bar")}
            for fi32 := 0; fi32 < 18; fi32++ {
                var v33 float32 = -2.25
                _ = fi32
                _ = v33
            }
            switch fi28 {
            case 4294967295:
                var v34 int = (fi28 >> ((42) & 63))
                _ = v34
            case -903:
                var v35 int32 = (v4 * -678)
                _ = v35
            case 289:
                _ = len(v6)
            }
            _ = v31
        }
        _ = fi28
    }
    fmt.Println(v2, v23)
    var v36 float32 = 0.1
    fmt.Println(v4)
    var v37 bool = ((("hello" != "a") || false) || false)
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
