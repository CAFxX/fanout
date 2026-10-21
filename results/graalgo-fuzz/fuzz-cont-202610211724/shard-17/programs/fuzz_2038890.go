package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 string; F2 int8 }
func genFn0(p0 int8, p1 bool) bool {
    var v4 bool = (40 > 816)
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := []int{fi5, 32767}
        for fi7 := 0; fi7 < 20; fi7++ {
            if v4 {
                var v8 int32 = 2147483647
                _ = v8
            }
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    if p1 {
        switch -1 {
        case 4294967295:
            switch "a" {
            case "foo bar":
                v9 := []int{int(p0), 1, 128, 42}
                _ = v9
            case "\n\t":
                for fi10 := 0; fi10 < 4; fi10++ {
                    v11 := GenS1{F0: 32767}
                    p0 += p0
                    _ = fi10
                    _ = v11
                }
            default:
                switch ("a" + "foo bar") {
                case "foo bar":
                    var v12 uint64 = 1
                    _ = v12
                }
            }
        case 127:
            p1 = (-1 < p0)
        }
        fmt.Println(p0, p0)
        fmt.Println(p1, p0)
    } else {
        fmt.Println(v4)
        for fi13 := 0; fi13 < 3; fi13++ {
            p0 = (p0 | p0)
            if (-2.25 <= 0.0) {
                v14 := GenS1{F0: 1000}
                v15 := []int{}
                _ = v14
                _ = v15
            }
            var v16 bool = false
            _ = fi13
            _ = v16
        }
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 uint32 = 65535
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if (2 != 65535) {
            v4 = (2 > 2)
            var v20 uint32 = 645
            v21 := []int{}
            _ = v20
            _ = v21
        }
        _ = fi19
    }
    v22 := GenS1{F0: 1000}
    _ = p0
    _ = p1
    _ = v22
    _ = v4
    return p1
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = -2.25
    var v3 uint32 = 1
    switch "" {
    case "hello":
        var v23 float32 = 0.0
        _ = v23
    case "a":
        for fi24 := 0; fi24 < 9; fi24++ {
            v3 -= (v3 % (((v3 + 4294967295)) | 1))
            _ = fi24
        }
    }
    v3 = (v3 >> ((42) & 63))
    var v25 int = 9223372036854775807
    v3 = v3
    for fi26 := 0; fi26 < 2; fi26++ {
        switch 900 {
        case -2147483648:
            var v27 int16 = -1
            _ = v27
        default:
            var v28 uint64 = 583
            _ = v28
        }
        switch "a" {
        case "hello":
            fmt.Println(v3, v25, v3)
        default:
            var v29 bool = true
            _ = v29
        }
        var v30 uint16 = 932
        _ = fi26
        _ = v30
    }
    fmt.Println(v3)
    v25 = v25
    switch "a" {
    case "a":
        if true {
            v25 -= v25
            var v31 uint16 = 0
            _ = v31
        } else {
            fmt.Println(v3, v3, v25)
        }
    case "foo bar":
        v32 := GenS1{F0: -1}
        _ = v32
    case "":
        if (!((100 < 127) && true)) {
            var v33 int64 = 561
            _ = len(v1)
            _ = v33
        }
    default:
        switch "" {
        case "a":
            for fi34 := 0; fi34 < 3; fi34++ {
                v35 := GenS2{F0: 1, F1: ("" + "hello"), F2: -128}
                _ = fi34
                _ = v35
            }
        case "":
            var v36 int32 = 298
            _ = v36
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
}
