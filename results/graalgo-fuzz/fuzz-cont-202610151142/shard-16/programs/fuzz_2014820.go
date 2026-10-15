package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 bool }
type GenS2 struct { F0 uint16; F1 uint; F2 int64 }
type GenS3 struct { F0 int32; F1 int64; F2 uint8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 string, p1 bool, p2 string) bool {
    if p1 {
        p1 = (18446744073709551615 == 42)
        v5 := []int{-400}
        _ = v5
    } else {
        switch ("a" + p2) {
        case "":
            fmt.Println(p1, p1)
        case "\n\t":
            for fi6 := 0; fi6 < 7; fi6++ {
                v7 := []int{fi6, 9223372036854775807, 0}
                p1 = p1
                v8 := GenS1{F0: 9223372036854775807, F1: 42, F2: p1}
                _ = fi6
                _ = v7
                _ = v8
            }
        case "a":
            var v9 float64 = 1e10
            _ = v9
        }
        p2 = ""
    }
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p1
}

func main() {
    v1 := []int{-1, -2147483648, 2147483647, -434}
    var v2 uint64 = 1
    v3 := make(map[string]bool)
    v4 := []string{("foo bar" + "foo bar")}
    v2 += 0
    switch "a" {
    case "a":
        var v11 int8 = 0
        _ = v11
    case "foo bar":
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 int = (-(fi12))
            switch -21 {
            case -693:
                var v14 string = "hello"
                _ = v14
            }
            _ = fi12
            _ = v13
        }
    default:
        switch "\n\t" {
        case "":
            for fi15 := 0; fi15 < 13; fi15++ {
                _ = len(v4)
                for fi16 := 0; fi16 < 14; fi16++ {
                    if (v2 != (v2 - v2)) {
                        v2 = v2
                        v2 = (v2 * 42)
                        fmt.Println(fi15)
                    }
                    v2 = v2
                    _ = fi16
                }
                v17 := GenS1{F0: (fi15 + (fi15 + 27)), F1: v2, F2: false}
                _ = fi15
                _ = v17
            }
        case "foo bar":
            var v18 float32 = 0.1
            _ = v18
        default:
            switch -1 {
            case -1:
                if ((!true) || (!(42 <= 42))) {
                    fmt.Println(v2)
                    v2 += v2
                    if false {
                        if (1 != -32768) {
                            for fi19 := 0; fi19 < 16; fi19++ {
                                _ = len(v3)
                                _ = fi19
                            }
                        }
                    }
                }
            case 42:
                var v20 uint32 = 2147483648
                _ = v20
            case 0:
                fmt.Println(v2)
            default:
                v21 := GenS1{F0: int(v2), F1: 2, F2: false}
                _ = v21
            }
        }
    }
    v3[("hello" + "\n\t")] = false
    v22 := GenS2{F0: 32767, F1: 508, F2: 584}
    var v23 int16 = 1
    var v24 uint16 = 41
    v25 := GenS2{F0: 42, F1: 0, F2: 553}
    v2 -= (v2 >> ((18446744073709551615) & 63))
    for fi26 := 0; fi26 < 16; fi26++ {
        if false {
            if false {
                v27 := v3["hello"]
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 19; fi28++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if true {
                var v29 int64 = 42
                v24 = v24
                _ = v29
            }
            var v30 uint16 = v24
            _ = fi28
            _ = v30
        }
        _ = fi26
    }
    v31 := GenS3{F0: 1, F1: 1, F2: 234}
    switch 2147483647 {
    case -767:
        v32 := GenS4{F0: 0}
        _ = v32
    case -2147483648:
        v33 := GenS2{F0: v24, F1: 255, F2: 9223372036854775807}
        _ = v33
    }
    if (!true) {
        var v34 uint64 = v2
        _ = v34
    }
    v23 = v23
    v2 -= v2
    var v35 uint = 466
    _ = len(v1)
    var v36 uint16 = v24
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
}
