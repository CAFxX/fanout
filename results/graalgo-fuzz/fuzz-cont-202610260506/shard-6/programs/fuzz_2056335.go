package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int) uint64 {
    p0 += 127
    if (!true) {
        var v4 uint = 396
        _ = v4
    } else {
        v5 := []int{9223372036854775807, p1, -1}
        _ = v5
    }
    switch (("\n\t" + "a") + "a") {
    case "a":
        var v6 uint = 975
        _ = v6
    case "hello":
        v7 := []int{p1, -857, p1, (p1 & (p1 / (((p1 & p1)) | 1)))}
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        p0 += (p0 - p0)
        if (!(!true)) {
            fmt.Println(fi8)
            fmt.Println(p0, p0)
        } else {
            p1 = p1
            switch (p1 | (fi8 % ((p1) | 1))) {
            case -2147483648:
                var v9 bool = (!false)
                _ = v9
            case -1:
                fmt.Println(p0, p0, p0)
            }
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 float64 = 3.14159
            fmt.Println(fi10)
            var v12 uint = 219
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    v13 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v13
    return 1
}
func genFn1(p0 int8) int {
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := []int{fi14, 305}
        _ = fi14
        _ = v15
    }
    p0 += p0
    _ = p0
    return 32768
}

func main() {
    var v1 int8 = 0
    var v2 int64 = -331
    var v3 uint64 = 2
    switch "" {
    case "hello":
        v1 = v1
    case "foo bar":
        var v16 int64 = 9223372036854775807
        _ = v16
    case "":
        if false {
            v17 := GenS2{F0: 1, F1: v1}
            if (!false) {
                v18 := GenS1{F0: -1}
                v19 := GenS2{F0: 2147483647, F1: v1}
                _ = v18
                _ = v19
            }
            var v20 uint32 = 65535
            _ = v17
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        for fi22 := 0; fi22 < 11; fi22++ {
            v3 = v3
            var v23 float32 = (1.5 / float32(-2.25))
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    fmt.Println(v3)
    if false {
        for fi24 := 0; fi24 < 8; fi24++ {
            v3 += v3
            _ = fi24
        }
    } else {
        fmt.Println(v3, v3, v2)
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := GenS3{F0: false}
        v3 -= 830
        if true {
            v2 = (v2 >> ((4294967295) & 63))
            if (!(275 > 619)) {
                var v27 int64 = v2
                switch (fi25 | fi25) {
                case 4294967295:
                    v27 += v27
                default:
                    for fi28 := 0; fi28 < 8; fi28++ {
                        v29 := GenS1{F0: 1000}
                        _ = fi28
                        _ = v29
                    }
                }
                _ = v27
            } else {
                if true {
                    v3 = (v3 + v3)
                    v1 = (v1 + (v1 >> ((65535) & 63)))
                    switch "a" {
                    case "":
                        v1 = -22
                    }
                } else {
                    v30 := []int{genFn1(0), (fi25 | (^(fi25)))}
                    var v31 uint8 = 4
                    _ = v30
                    _ = v31
                }
                v32 := GenS4{F0: 9223372036854775807, F1: 268}
                _ = v32
            }
        }
        _ = fi25
        _ = v26
    }
    v33 := []int{-624, 32768}
    _ = len(v33)
    for fi34 := 0; fi34 < 16; fi34++ {
        v35 := GenS5{F0: 2147483647}
        switch fi34 {
        case -961:
            if ((^(v2)) >= (v2 << ((614) & 63))) {
                if false {
                    fmt.Println(v2)
                }
                for fi36 := 0; fi36 < 17; fi36++ {
                    if (!(-2.25 <= (float64(1.5) / 3.14159))) {
                        var v37 int = 32768
                        fmt.Println(v2)
                        v37 = fi36
                        _ = v37
                    } else {
                        switch -2147483648 {
                        case -218:
                            _ = len(v33)
                        default:
                            fmt.Println(v1)
                        }
                    }
                    fmt.Println(v3)
                    v38 := GenS1{F0: 1000}
                    v1 = 42
                    fmt.Println(v2, fi36)
                    _ = fi36
                    _ = v38
                }
            } else {
                v39 := GenS1{F0: -875}
                _ = v39
            }
        }
        _ = fi34
        _ = v35
    }
    var v40 int8 = v1
    var v41 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v40
    _ = v41
}
