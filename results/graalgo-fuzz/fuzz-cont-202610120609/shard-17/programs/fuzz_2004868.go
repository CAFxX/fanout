package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int32; F2 int32 }
func genFn0(p0 uint64, p1 int16) uint16 {
    var v5 bool = true
    if true {
        if (-649 != -2147483648) {
            fmt.Println(p0, p1, p0)
            v6 := []int{904, 128, 32768}
            _ = v6
        }
    } else {
        var v7 int16 = p1
        var v8 uint32 = 2
        _ = v7
        _ = v8
    }
    v9 := []int{-677, -877, 0}
    v10 := GenS1{F0: -2147483648, F1: ("a" + "")}
    p1 = 32767
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    _ = v9
    return 675
}
func genFn1(p0 string, p1 int64, p2 string) uint32 {
    switch 1 {
    case 32768:
        p2 = p2
    case 42:
        var v11 int8 = -128
        _ = v11
    case 682:
        switch "\n\t" {
        case "a":
            if (0 != -1) {
                v12 := GenS1{F0: 0, F1: "a"}
                fmt.Println(p0, p2)
                v13 := []int{}
                _ = v12
                _ = v13
            }
        case "\n\t":
            switch p2 {
            case "hello":
                if true {
                    var v14 int8 = 0
                    _ = v14
                }
            case "foo bar":
                switch 667 {
                case -1:
                    v15 := []int{}
                    _ = v15
                case 127:
                    for fi16 := 0; fi16 < 2; fi16++ {
                        var v17 int16 = -179
                        _ = fi16
                        _ = v17
                    }
                default:
                    v18 := GenS1{F0: 0, F1: p0}
                    _ = v18
                }
            case "a":
                p0 = "\n\t"
            }
        case "foo bar":
            p2 = p0
        }
    }
    v19 := GenS1{F0: -1, F1: ""}
    v20 := GenS2{F0: 1}
    var v21 int16 = -2
    var v22 string = (p0 + p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v21
    _ = v22
    return 4294967295
}

func main() {
    var v1 string = "hello"
    v2 := []int{127, 0}
    var v3 uint8 = 1
    var v4 uint = 4294967295
    var v23 float64 = 0.0
    if (true || (2147483647 < 772)) {
        fmt.Println(v4, v23)
    } else {
        switch "" {
        case "":
            for fi24 := 0; fi24 < 12; fi24++ {
                v3 -= (v3 << ((v4) & 63))
                _ = fi24
            }
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi25
        }
        var v26 string = "hello"
        _ = v26
    }
    var v27 int32 = -2147483648
    var v28 uint8 = v3
    v1 = "\n\t"
    if len(v2) > 0 {
        v2[0] = -984
    }
    var v29 int16 = -32768
    var v30 uint32 = 763
    var v31 int8 = 100
    v32 := GenS3{F0: 5, F1: -2147483648, F2: v27}
    if (543 > 2) {
        switch 1 {
        case 128:
            for fi33 := 0; fi33 < 17; fi33++ {
                for fi34 := 0; fi34 < 13; fi34++ {
                    v30 = (v30 - genFn1(v1, 828, (v1 + "foo bar")))
                    var v35 uint32 = v30
                    v30 = (v35 + 65535)
                    v23 += v23
                    _ = fi34
                    _ = v35
                }
                _ = fi33
            }
        default:
            var v36 bool = true
            _ = v36
        }
        v31 += v31
    }
    v37 := GenS2{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v4
}
