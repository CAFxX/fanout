package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 bool; F2 string }
func genFn0(p0 uint8, p1 int, p2 int) int16 {
    v4 := []int{}
    for fi5 := 0; fi5 < 19; fi5++ {
        switch ("" + ("hello" + "foo bar")) {
        case "a":
            if false {
                var v6 float32 = ((1e10 / 1.0) + 1.5)
                var v7 int16 = -1
                _ = v6
                _ = v7
            }
        default:
            v8 := GenS1{F0: 1, F1: 100, F2: 128}
            _ = v8
        }
        v9 := GenS2{F0: -565}
        _ = fi5
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        switch "foo bar" {
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "foo bar":
            for fi11 := 0; fi11 < 5; fi11++ {
                p1 -= p1
                fmt.Println(fi11, fi11)
                if (("foo bar" >= "\n\t") || (false || false)) {
                    p2 = (fi10 >> ((180) & 63))
                    var v12 uint32 = 283
                    _ = v12
                } else {
                    v13 := GenS2{F0: -131}
                    switch ((("foo bar" + "hello") + ("foo bar" + (("hello" + ("hello" + "hello")) + "hello"))) + "hello") {
                    case "hello":
                        if true {
                            p2 = (-(p2))
                            var v14 int = fi10
                            _ = v14
                        }
                        if true {
                            fmt.Println(p1, fi11, fi11)
                            v15 := GenS2{F0: -9223372036854775808}
                            v16 := GenS3{F0: 0, F1: true, F2: ""}
                            _ = v15
                            _ = v16
                        }
                    case "\n\t":
                        var v17 uint8 = p0
                        _ = v17
                    case "":
                        p1 -= fi11
                    }
                    v18 := GenS2{F0: int64(p0)}
                    _ = v13
                    _ = v18
                }
                _ = fi11
            }
        case "hello":
            var v19 int16 = -1
            _ = v19
        }
        _ = fi10
    }
    switch ("\n\t" + "foo bar") {
    case "hello":
        p0 += p0
    case "\n\t":
        var v20 uint64 = 0
        _ = v20
    default:
        switch "a" {
        case "\n\t":
            fmt.Println(p0, p2)
        case "foo bar":
            p1 += (p2 << ((1) & 63))
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return int16(p1)
}

func main() {
    var v1 uint16 = 219
    var v2 bool = true
    var v3 int = -919
    var v21 int8 = 17
    var v22 uint = 0
    for fi23 := 0; fi23 < 12; fi23++ {
        v1 -= v1
        v24 := []int{v3, v3, (fi23 + 128), (fi23 % ((v3) | 1))}
        _ = fi23
        _ = v24
    }
    v25 := []int{(-(v3)), v3, (v3 << (((v22 * v22)) & 63)), (v3 + (v3 & (v3 - 4294967295)))}
    if false {
        var v26 int = (v3 & (v3 / ((v3) | 1)))
        if len(v25) > 0 {
            v25[0] = (^(v3))
        }
        fmt.Println(v3, v22, v2)
        _ = v26
    }
    v21 = (v21 + v21)
    var v27 uint = (v22 / (((v22 - 788)) | 1))
    _ = len(v25)
    var v28 uint = v22
    var v29 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
