package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS2 struct { F0 int; F1 int; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint32; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 uint16 }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint16 = 255
        v7 := GenS1{F0: 2, F1: v6, F2: 0}
        v6 = 42
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int = 127
    _ = v8
    return 127
}
func genFn1(p0 int32, p1 uint) uint32 {
    for fi9 := 0; fi9 < 11; fi9++ {
        if (true || false) {
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 uint8 = 128
                if (true || ((p1 | p1) < p1)) {
                    if (!true) {
                        v12 := GenS1{F0: 18446744073709551615, F1: 540, F2: -380}
                        p0 = (^(p0))
                        if (!(true && true)) {
                            v13 := []int{fi9, 630, 26}
                            p0 = -2147483648
                            _ = v13
                        } else {
                            v14 := GenS1{F0: 42, F1: 1, F2: 667}
                            _ = v14
                        }
                        switch "hello" {
                        case "\n\t":
                            p0 = 2147483647
                        case "hello":
                            if (p0 != p0) {
                                v15 := GenS2{F0: (fi10 + fi10), F1: (^(fi10)), F2: 1}
                                _ = v15
                            }
                        }
                        var v16 uint = 0
                        v17 := []int{}
                        _ = v12
                        _ = v16
                        _ = v17
                    } else {
                        var v18 string = ("foo bar" + "a")
                        fmt.Println(fi10, fi9, fi10)
                        p1 = p1
                        _ = v18
                    }
                    fmt.Println(fi9, p1)
                } else {
                    fmt.Println(fi9, p0, p0)
                    switch (fi9 ^ fi9) {
                    case 4294967295:
                        v11 = (v11 & uint8(p1))
                    case 532:
                        v19 := GenS3{F0: 0, F1: 4294967295, F2: 42}
                        _ = v19
                    case -2147483648:
                        var v20 uint = p1
                        _ = v20
                    default:
                        v21 := GenS1{F0: 2, F1: 0, F2: 502}
                        _ = v21
                    }
                }
                _ = fi10
                _ = v11
            }
        } else {
            v22 := []int{(-(fi9)), fi9, (^(fi9))}
            _ = v22
        }
        _ = fi9
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    var v1 int16 = -32768
    var v2 uint32 = 42
    v3 := make(map[int]bool)
    var v4 uint32 = v2
    if false {
        if (!false) {
            fmt.Println(v1, v2)
            fmt.Println(v1)
            v1 = v1
        }
        var v23 uint32 = (v4 / ((v2) | 1))
        _ = v23
    }
    v24 := GenS4{F0: 0}
    var v25 float32 = ((-2.25 + -2.25) * (1.5 + (0.1 * -2.25)))
    fmt.Println(v25, v4)
    var v26 float64 = (1.5 - 1.5)
    if true {
        switch -2147483648 {
        case 145:
            v27 := []int{-897, 650, -1, 0}
            _ = v27
        default:
            switch 128 {
            case 127:
                var v28 int = -9223372036854775808
                _ = v28
            }
        }
        var v29 int = -824
        if true {
            if (42 <= 42) {
                for fi30 := 0; fi30 < 10; fi30++ {
                    v29 = 127
                    _ = fi30
                }
                if (v25 < float32((3.14159 + float64(v25)))) {
                    var v31 float64 = v26
                    v26 = -2.25
                    v25 -= v25
                    _ = v31
                }
                var v32 uint8 = 0
                _ = v32
            } else {
                for fi33 := 0; fi33 < 2; fi33++ {
                    v34 := GenS4{F0: 32767}
                    switch -355 {
                    case 32767:
                        _ = len(v3)
                    default:
                        fmt.Println(v29)
                    }
                    _ = fi33
                    _ = v34
                }
                var v35 int16 = 42
                _ = v35
            }
        }
        _ = v29
    } else {
        if false {
            if false {
                for fi36 := 0; fi36 < 1; fi36++ {
                    if true {
                        v26 = (v26 / (v26 * (float64(v25) * v26)))
                        var v37 int16 = (v1 >> ((999) & 63))
                        _ = v37
                    }
                    _ = fi36
                }
            }
            var v38 float64 = float64(1e10)
            _ = v38
        }
        _ = len(v3)
    }
    var v39 int32 = 42
    if ((2147483647 > v39) && true) {
        if true {
            switch ("hello" + "") {
            case "foo bar":
                v40 := GenS1{F0: 2, F1: uint16(v4), F2: 19}
                _ = v40
            default:
                switch 411 {
                case 127:
                    var v41 uint64 = 4294967295
                    _ = v41
                case 1:
                    for fi42 := 0; fi42 < 1; fi42++ {
                        v4 += (v4 & 2147483648)
                        v26 -= 1.5
                        _ = fi42
                    }
                }
            }
            for fi43 := 0; fi43 < 2; fi43++ {
                v44 := GenS5{F0: v4, F1: 1}
                _ = fi43
                _ = v44
            }
            delete(v3, 32768)
        }
    }
    v4 += v4
    v45 := []int{32768, 32768, 127}
    if len(v45) >= 2 {
        v45 = v45[1:len(v45)]
    }
    var v46 int = 32767
    switch ("hello" + ("a" + "a")) {
    case "a":
        var v47 uint = 2
        _ = v47
    case "\n\t":
        var v48 float32 = v25
        _ = v48
    case "":
        v49 := GenS2{F0: -2147483648, F1: -470, F2: 458}
        _ = v49
    }
    var v50 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v46:", v46)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v45
    _ = v46
    _ = v50
}
