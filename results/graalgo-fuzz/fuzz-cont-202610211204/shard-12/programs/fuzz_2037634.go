package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 int32 = 0
    var v2 int8 = 0
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 2; fi5++ {
            v6 := []int{(fi5 & 32768), (fi5 / (((fi5 | -435)) | 1))}
            switch 42 {
            case 4294967295:
                switch fi4 {
                case 4294967295:
                    v7 := GenS1{F0: 4294967295}
                    _ = v7
                default:
                    v2 += (v2 / ((v2) | 1))
                }
            case 1:
                if false {
                    var v8 uint64 = 2
                    v1 = int32(v8)
                    _ = v8
                }
            }
            _ = fi5
            _ = v6
        }
        if false {
            var v9 int64 = 42
            _ = v9
        }
        _ = fi4
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            if true {
                var v12 int64 = -9223372036854775808
                switch "foo bar" {
                case "hello":
                    v1 += (-(v1))
                case "":
                    v13 := GenS1{F0: 72}
                    _ = v13
                }
                fmt.Println(fi10, v12, v12)
                _ = v12
            }
            _ = fi11
        }
        _ = fi10
    }
    fmt.Println(v2, v2)
    switch "a" {
    case "\n\t":
        v14 := GenS1{F0: 733}
        _ = v14
    case "hello":
        v2 += (v2 >> ((2) & 63))
    case "a":
        switch "foo bar" {
        case "a":
            v15 := GenS2{F0: 42}
            _ = v15
        default:
            fmt.Println(v2, v1)
        }
    default:
        if true {
            fmt.Println(v1)
            v1 -= (v1 / (((v1 - (v1 * (-(v1))))) | 1))
        }
    }
    switch "\n\t" {
    case "foo bar":
        var v16 uint32 = 443
        _ = v16
    case "a":
        v17 := GenS3{F0: -1, F1: 818}
        _ = v17
    case "":
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 uint64 = 18446744073709551615
            v20 := GenS2{F0: fi18}
            var v21 bool = (v2 != v2)
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
    default:
        if (float32(-2.25) <= 1e10) {
            for fi22 := 0; fi22 < 18; fi22++ {
                v1 -= 2147483647
                if true {
                    switch 128 {
                    case -671:
                        var v23 uint = 232
                        _ = v23
                    }
                }
                for fi24 := 0; fi24 < 13; fi24++ {
                    v25 := []int{2147483647}
                    _ = fi24
                    _ = v25
                }
                _ = fi22
            }
            fmt.Println(v1, v2)
        } else {
            var v26 uint8 = 128
            for fi27 := 0; fi27 < 15; fi27++ {
                delete(v3, 42)
                _ = fi27
            }
            var v28 uint8 = v26
            _ = v26
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        _ = len(v3)
        var v30 uint = 4294967295
        _ = fi29
        _ = v30
    }
    v31 := []int{128, 127}
    fmt.Println(v1)
    switch -9223372036854775808 {
    case 127:
        switch 127 {
        case 128:
            var v32 int = 32768
            _ = v32
        case 0:
            for fi33 := 0; fi33 < 18; fi33++ {
                delete(v3, 4294967295)
                v1 = (v1 * -431)
                _ = fi33
            }
        case 42:
            var v34 float32 = -2.25
            _ = v34
        }
    case 1:
        v35 := GenS1{F0: 2}
        _ = v35
    case 570:
        var v36 int32 = v1
        _ = v36
    }
    if true {
        var v37 int8 = -39
        v38 := GenS4{F0: 153}
        _ = v37
        _ = v38
    }
    v2 = (v2 - v2)
    switch 127 {
    case -440:
        v1 += -1
    }
    var v39 uint8 = 2
    var v40 uint = 42
    v39 = (v39 & v39)
    _ = len(v31)
    fmt.Println(v40, v2, v1)
    for fi41 := 0; fi41 < 13; fi41++ {
        if len(v31) > 0 {
            v31[0] = (fi41 >> ((4294967295) & 63))
        }
        for fi42 := 0; fi42 < 2; fi42++ {
            var v43 uint = (v40 & (v40 % (((v40 * (v40 << ((v40) & 63)))) | 1)))
            _ = fi42
            _ = v43
        }
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v39
    _ = v40
}
