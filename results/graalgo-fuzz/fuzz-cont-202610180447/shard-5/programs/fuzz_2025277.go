package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 int16; F2 bool }
func genFn0(p0 uint16, p1 int, p2 string) uint {
    if (!true) {
        switch p2 {
        case "foo bar":
            var v4 uint16 = p0
            _ = v4
        case "hello":
            for fi5 := 0; fi5 < 15; fi5++ {
                if (-1 > 598) {
                    if (p2 == (p2 + "\n\t")) {
                        switch "hello" {
                        case "foo bar":
                            if false {
                                fmt.Println(p2, p1, p2)
                                switch (-(fi5)) {
                                case 730:
                                    p0 = (p0 / (((p0 - (p0 + 0))) | 1))
                                case 532:
                                    var v6 float64 = 0.0
                                    _ = v6
                                }
                                p2 = (p2 + "foo bar")
                            }
                        }
                        var v7 uint = 521
                        v7 = v7
                        _ = v7
                    } else {
                        p1 -= (-(p1))
                        var v8 uint = 492
                        _ = v8
                    }
                    var v9 uint8 = 42
                    for fi10 := 0; fi10 < 15; fi10++ {
                        var v11 uint = 255
                        v11 -= v11
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                }
                var v12 int64 = 6
                _ = fi5
                _ = v12
            }
        }
    } else {
        var v13 float32 = float32(3.14159)
        var v14 int = (p1 + -9223372036854775808)
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 65535
}
func genFn1() int8 {
    switch -9223372036854775808 {
    case 1:
        var v16 float64 = 3.14159
        _ = v16
    }
    if (379 < 0) {
        var v17 int = 128
        _ = v17
    }
    var v18 int64 = -1
    var v19 float32 = (0.0 + 3.14159)
    _ = v18
    _ = v19
    return -1
}

func main() {
    var v1 float64 = (1.5 / (((3.14159 * 1e10) + 0.0) * (1e10 * (float64(((1e10 + (1e10 - 1.5)) + 1e10)) + 1.5))))
    var v2 string = "foo bar"
    var v3 string = ""
    v1 = 1e10
    var v20 int64 = 341
    fmt.Println(v2)
    v3 = ("hello" + v2)
    for fi21 := 0; fi21 < 17; fi21++ {
        v22 := GenS2{F0: -1, F1: 0, F2: 2}
        v23 := GenS3{F0: 202}
        v24 := []int{598, 60, fi21}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    fmt.Println(v1)
    if (!(v20 > (v20 + (^(v20))))) {
        if true {
            for fi25 := 0; fi25 < 14; fi25++ {
                for fi26 := 0; fi26 < 1; fi26++ {
                    if (-910 <= -2147483648) {
                        v27 := GenS3{F0: 127}
                        _ = v27
                    } else {
                        v28 := GenS4{F0: 22, F1: 1, F2: 2147483648}
                        fmt.Println(v2, fi25)
                        var v29 int8 = 0
                        _ = v28
                        _ = v29
                    }
                    _ = fi26
                }
                _ = fi25
            }
            v2 = (((v3 + "\n\t") + v3) + "a")
        } else {
            var v30 string = v2
            if (1.5 > 0.1) {
                v20 = 611
                fmt.Println(v20, v20, v1)
                v31 := GenS5{F0: v20}
                _ = v31
            }
            _ = v30
        }
        for fi32 := 0; fi32 < 18; fi32++ {
            switch "hello" {
            case "a":
                fmt.Println(v20, fi32)
            case "\n\t":
                var v33 uint = 1
                _ = v33
            case "":
                fmt.Println(v2)
            }
            v2 = v3
            var v34 uint64 = 692
            _ = fi32
            _ = v34
        }
    } else {
        var v35 int32 = 1
        v36 := GenS6{F0: v35, F1: -724, F2: true}
        _ = v35
        _ = v36
    }
    var v37 uint16 = 32767
    var v38 int = 32768
    var v39 uint16 = v37
    for fi40 := 0; fi40 < 4; fi40++ {
        switch fi40 {
        case 128:
            v41 := GenS1{F0: (v37 ^ (v37 | 2))}
            _ = v41
        default:
            for fi42 := 0; fi42 < 18; fi42++ {
                v43 := GenS6{F0: 367, F1: int16(v39), F2: false}
                switch (v38 | (fi42 << ((255) & 63))) {
                case 0:
                    v3 = (v2 + (v2 + v2))
                case 4294967295:
                    fmt.Println(v20, v20)
                default:
                    var v44 int8 = genFn1()
                    _ = v44
                }
                _ = fi42
                _ = v43
            }
        }
        var v45 float64 = 0.0
        _ = fi40
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v38
    _ = v39
}
