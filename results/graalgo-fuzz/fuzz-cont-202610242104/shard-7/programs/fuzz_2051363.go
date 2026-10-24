package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int; F2 int32 }
func genFn0(p0 int8, p1 uint32, p2 int64) uint64 {
    switch ((("\n\t" + ("foo bar" + "foo bar")) + "foo bar") + "a") {
    case "":
        var v6 uint64 = 1
        _ = v6
    default:
        v7 := GenS1{F0: 142}
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        p0 = (p0 / ((p0) | 1))
        _ = fi8
    }
    p0 = int8(p2)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float32 = 0.1
    v3 := make(map[int]int)
    var v4 int64 = v1
    var v5 int64 = -221
    for fi9 := 0; fi9 < 20; fi9++ {
        switch "\n\t" {
        case "foo bar":
            if false {
                fmt.Println(v5, v4)
                var v10 uint64 = genFn0(-1, 2, (v4 >> ((1) & 63)))
                _ = v10
            }
        case "\n\t":
            v3[(fi9 | int(v5))] = fi9
        default:
            v5 = v4
        }
        var v11 int = fi9
        for fi12 := 0; fi12 < 12; fi12++ {
            switch fi12 {
            case 4294967295:
                var v13 float32 = v2
                v14 := []int{(fi12 & fi12)}
                _ = v13
                _ = v14
            default:
                var v15 float32 = 1.5
                _ = v15
            }
            var v16 string = ""
            v17 := v3[v11]
            _ = fi12
            _ = v16
            _ = v17
        }
        _ = fi9
        _ = v11
    }
    switch "a" {
    case "\n\t":
        switch "a" {
        case "a":
            v1 += -1
        default:
            fmt.Println(v5, v5, v4)
        }
    case "a":
        if (((v2 - 1.5) * v2) != float32(3.14159)) {
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 uint64 = 594
                v5 -= v5
                for fi20 := 0; fi20 < 15; fi20++ {
                    if (true || true) {
                        switch fi20 {
                        case 1:
                            v1 += v5
                        default:
                            v21 := GenS1{F0: 269}
                            _ = v21
                        }
                    }
                    _ = fi20
                }
                _ = fi18
                _ = v19
            }
            fmt.Println(v1)
            fmt.Println(v1)
        }
    case "":
        v5 = -984
    }
    v22 := GenS2{F0: 2147483648, F1: 0, F2: 380}
    fmt.Println(v4)
    v4 = v4
    v23 := []int{}
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 int16 = -32768
        var v26 int16 = v25
        var v27 uint = 469
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println(v4, v4)
    var v28 bool = (790 != 2)
    switch 2147483647 {
    case -865:
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
    case -111:
        switch (("\n\t" + "hello") + ("\n\t" + "")) {
        case "":
            fmt.Println(v4)
        }
    }
    switch "a" {
    case "":
        var v29 string = ("foo bar" + "foo bar")
        _ = v29
    case "foo bar":
        var v30 bool = true
        _ = v30
    default:
        delete(v3, int(v1))
    }
    if (-1 != -960) {
        for fi31 := 0; fi31 < 19; fi31++ {
            for fi32 := 0; fi32 < 11; fi32++ {
                switch fi31 {
                case 301:
                    var v33 int32 = 0
                    _ = v33
                case -543:
                    fmt.Println(v5, v4)
                case 4294967295:
                    switch "foo bar" {
                    case "a":
                        var v34 uint8 = 128
                        _ = v34
                    }
                }
                _ = fi32
            }
            _ = fi31
        }
    }
    v23 = append(v23, 32768)
    v4 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
