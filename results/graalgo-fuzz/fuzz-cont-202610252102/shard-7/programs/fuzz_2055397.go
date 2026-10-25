package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint; F1 uint8; F2 uint32 }
func genFn0(p0 uint64, p1 int32, p2 string) uint64 {
    switch 4294967295 {
    case 1:
        if ((1 < 32767) && ((100 > 0) && (127 > 4294967295))) {
            var v6 int64 = 1
            v7 := GenS1{F0: 1, F1: false}
            v8 := GenS1{F0: 1, F1: true}
            _ = v6
            _ = v7
            _ = v8
        } else {
            p1 = p1
            p1 += (p1 / (((p1 / (((p1 + -516)) | 1))) | 1))
            switch p2 {
            case "hello":
                for fi9 := 0; fi9 < 4; fi9++ {
                    var v10 int16 = 1000
                    v11 := []int{fi9, (fi9 & fi9), 9223372036854775807}
                    var v12 int32 = (p1 / ((p1) | 1))
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            case "":
                v13 := GenS1{F0: 9223372036854775807, F1: true}
                _ = v13
            }
        }
    case 32768:
        var v14 bool = (-1 > 1000)
        _ = v14
    case 184:
        switch ("" + "foo bar") {
        case "hello":
            var v15 uint8 = 193
            _ = v15
        case "":
            if ((!(!true)) || ((((-2.25 + -2.25) * 1.5) > -2.25) || true)) {
                fmt.Println(p1, p0, p2)
                for fi16 := 0; fi16 < 9; fi16++ {
                    v17 := GenS1{F0: 1, F1: (!true)}
                    p0 -= (p0 - p0)
                    _ = fi16
                    _ = v17
                }
            }
        }
    default:
        var v18 int = 127
        _ = v18
    }
    v19 := []int{27}
    v20 := GenS1{F0: 42, F1: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return p0
}

func main() {
    var v1 uint16 = 255
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]bool)
    var v4 bool = false
    var v5 int64 = -170
    var v21 uint = 18446744073709551615
    v5 = int64(v1)
    for fi22 := 0; fi22 < 14; fi22++ {
        switch "hello" {
        case "foo bar":
            for fi23 := 0; fi23 < 13; fi23++ {
                var v24 uint8 = 255
                v25 := []int{851, -1}
                _ = fi23
                _ = v24
                _ = v25
            }
        case "\n\t":
            var v26 int16 = 414
            _ = v26
        case "hello":
            v5 = v5
        }
        _ = fi22
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        fmt.Println(fi27, v21)
        switch ("" + ("\n\t" + "hello")) {
        case "hello":
            switch "foo bar" {
            case "\n\t":
                var v28 int32 = int32(v2)
                _ = v28
            case "":
                var v29 int64 = (v5 - (^(v5)))
                _ = v29
            default:
                for fi30 := 0; fi30 < 13; fi30++ {
                    v31 := []int{}
                    fmt.Println(v5, v5)
                    var v32 string = "\n\t"
                    fmt.Println(fi27)
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
        case "\n\t":
            for fi33 := 0; fi33 < 16; fi33++ {
                var v34 float64 = 1e10
                if v4 {
                    var v35 bool = (!true)
                    v3[1] = (v2 <= 2)
                    v5 = (v5 ^ 215)
                    _ = v35
                }
                switch "foo bar" {
                case "":
                    var v36 uint = v21
                    _ = v36
                }
                _ = fi33
                _ = v34
            }
        default:
            v37 := v3[2]
            _ = v37
        }
        v38 := v3[0]
        _ = fi27
        _ = v38
    }
    var v39 int = 127
    var v40 int = v39
    if (v4 || false) {
        switch v39 {
        case -169:
            for fi41 := 0; fi41 < 16; fi41++ {
                fmt.Println(v5)
                v4 = v4
                _ = fi41
            }
        case -427:
            fmt.Println(v4)
        }
        switch (v39 | 128) {
        case 4294967295:
            var v42 uint = v21
            _ = v42
        case -318:
            v5 += -744
        }
        switch "foo bar" {
        case "":
            switch "\n\t" {
            case "\n\t":
                var v43 uint64 = (^(v2))
                _ = v43
            case "foo bar":
                switch "a" {
                case "foo bar":
                    switch v40 {
                    case -2147483648:
                        v44 := []int{v39, v40}
                        _ = v44
                    }
                }
            case "hello":
                var v45 uint8 = 1
                _ = v45
            }
        case "hello":
            var v46 uint = (^(v21))
            _ = v46
        default:
            v47 := GenS2{F0: v21, F1: 229, F2: 4294967295}
            _ = v47
        }
    }
    if (!(!v4)) {
        fmt.Println(v21, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
