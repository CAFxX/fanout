package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int; F2 int8 }
type GenS3 struct { F0 int; F1 uint32; F2 int8 }
func genFn0(p0 uint64, p1 int, p2 int16) int32 {
    if false {
        v6 := GenS1{F0: 0}
        _ = v6
    }
    var v7 float32 = -2.25
    fmt.Println(p2, p2, p1)
    p2 = 1000
    if (255 == 42) {
        v8 := GenS1{F0: -9223372036854775808}
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 bool = (false || (true || (true || true)))
            switch "hello" {
            case "hello":
                if v10 {
                    var v11 int16 = -128
                    for fi12 := 0; fi12 < 9; fi12++ {
                        fmt.Println(p1, p1)
                        v13 := []int{p1, (p1 ^ fi12), (-(fi12)), -1}
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                }
            case "a":
                p2 = (p2 - -1)
            case "":
                fmt.Println(p2, p2)
            default:
                switch "foo bar" {
                case "\n\t":
                    p2 += p2
                case "":
                    var v14 int8 = 113
                    _ = v14
                default:
                    v15 := []int{}
                    _ = v15
                }
            }
            v16 := []int{p1, p1}
            _ = fi9
            _ = v10
            _ = v16
        }
        switch p1 {
        case -95:
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 uint8 = 127
                v19 := []int{32768, (p1 * (p1 / ((263) | 1))), (fi17 / (((fi17 + (p1 >> ((570) & 63)))) | 1)), fi17}
                _ = fi17
                _ = v18
                _ = v19
            }
        case 128:
            p1 += p1
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 449
}

func main() {
    var v1 uint16 = 42
    var v2 uint32 = 65535
    var v3 int = -78
    v4 := []int{-415}
    var v5 uint32 = v2
    v4 = append(v4, (v3 / (((v3 & v3)) | 1)))
    switch int(v2) {
    case -1:
        v1 = (v1 * v1)
    default:
        v1 -= (v1 / ((v1) | 1))
    }
    v20 := GenS2{F0: false, F1: 9223372036854775807, F2: 0}
    v2 = (v2 % (((^(v2))) | 1))
    if (-181 != 2147483647) {
        fmt.Println(v5, v5)
    }
    switch v3 {
    case 42:
        v5 = 622
    case 734:
        v1 += 627
    case 2147483647:
        if (-1 < 101) {
            if len(v4) > 0 {
                v4[0] = v3
            }
        } else {
            if (uint8(v5) != 255) {
                v4 = append(v4, 108)
                v2 = 583
                var v21 uint = 255
                _ = v21
            }
        }
    }
    var v22 int8 = 127
    _ = len(v4)
    var v23 int32 = -55
    v3 = v3
    if (!(((-604 != (^(v23))) || (uint64(v5) <= 42)) || false)) {
        for fi24 := 0; fi24 < 2; fi24++ {
            switch "hello" {
            case "\n\t":
                v4 = append(v4, -628)
            case "a":
                switch (fi24 >> ((718) & 63)) {
                case 0:
                    v25 := GenS1{F0: 9223372036854775807}
                    _ = v25
                case 1:
                    fmt.Println(v22)
                }
            }
            v2 = v2
            for fi26 := 0; fi26 < 17; fi26++ {
                var v27 uint = 1
                var v28 uint16 = (^(v1))
                var v29 float32 = 1e10
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            _ = fi24
        }
        v4 = append(v4, v3)
    } else {
        fmt.Println(v3)
        switch 1 {
        case 2147483647:
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 float32 = float32((0.1 - 0.0))
                v22 -= -128
                var v32 float32 = v31
                _ = fi30
                _ = v31
                _ = v32
            }
        case -864:
            _ = len(v4)
        }
        if (42 == 128) {
            if len(v4) > 0 {
                v4[0] = -1
            }
            var v33 uint = 891
            var v34 int32 = v23
            _ = v33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi36 := 0; fi36 < 1; fi36++ {
            fmt.Println(v5)
            v37 := GenS1{F0: 1}
            _ = fi36
            _ = v37
        }
        v38 := GenS3{F0: (fi35 / (((v3 - 9223372036854775807)) | 1)), F1: (v5 >> ((255) & 63)), F2: (v22 >> ((1) & 63))}
        _ = fi35
        _ = v38
    }
    switch "foo bar" {
    case "foo bar":
        for fi39 := 0; fi39 < 7; fi39++ {
            for fi40 := 0; fi40 < 17; fi40++ {
                var v41 uint = 0
                _ = fi40
                _ = v41
            }
            _ = fi39
        }
    }
    v23 = 816
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
