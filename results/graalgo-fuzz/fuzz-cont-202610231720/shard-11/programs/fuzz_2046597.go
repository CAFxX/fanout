package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]bool)
    v3 := []int{}
    var v4 float32 = (-2.25 / 1e10)
    var v5 float32 = ((v4 * v4) * v4)
    if true {
        var v6 int = 850
        switch (v6 * 1) {
        case 4294967295:
            var v7 uint8 = 128
            _ = v7
        case -58:
            var v8 int16 = 1
            _ = v8
        case 2147483647:
            for fi9 := 0; fi9 < 7; fi9++ {
                switch fi9 {
                case 0:
                    v4 += (float32((1.5 + (-2.25 + 0.0))) + v5)
                case -9223372036854775808:
                    v4 += float32((1.5 + float64(v4)))
                case 128:
                    var v10 int8 = -91
                    _ = v10
                default:
                    v3 = append(v3, v6)
                }
                fmt.Println(v1, v1, v5)
                switch "foo bar" {
                case "\n\t":
                    v4 = v4
                case "hello":
                    var v11 int16 = 1
                    _ = v11
                case "a":
                    v1 -= (^(v1))
                }
                _ = fi9
            }
        default:
            for fi12 := 0; fi12 < 2; fi12++ {
                switch "hello" {
                case "foo bar":
                    _ = len(v3)
                case "":
                    v13 := GenS1{F0: (v1 + (v1 % (((v1 << ((65535) & 63))) | 1))), F1: 237, F2: -6}
                    _ = v13
                case "hello":
                    var v14 uint16 = 65535
                    _ = v14
                default:
                    var v15 uint8 = 1
                    _ = v15
                }
                _ = fi12
            }
        }
        var v16 int = v6
        _ = v16
        _ = v6
    } else {
        for fi17 := 0; fi17 < 10; fi17++ {
            fmt.Println(v1, v4, v1)
            v5 = v5
            v1 = (v1 - (v1 >> ((1) & 63)))
            _ = fi17
        }
        for fi18 := 0; fi18 < 20; fi18++ {
            if false {
                var v19 float32 = (-2.25 + (float32(0.0) * v5))
                v4 -= float32((-2.25 / -2.25))
                _ = v19
            } else {
                v1 = (v1 - 830)
                switch "" {
                case "a":
                    v20 := GenS1{F0: 539, F1: 336, F2: fi18}
                    _ = v20
                case "":
                    v4 -= v4
                default:
                    fmt.Println(v5, v4, fi18)
                }
            }
            switch ((("\n\t" + "") + "a") + "\n\t") {
            case "hello":
                fmt.Println(fi18, v5, fi18)
            case "":
                var v21 string = ("\n\t" + (("a" + "hello") + ""))
                _ = v21
            case "foo bar":
                var v22 float64 = 0.1
                _ = v22
            default:
                if len(v3) > 0 {
                    v3[0] = fi18
                }
            }
            _ = fi18
        }
    }
    v4 = v4
    v23 := GenS1{F0: (v1 >> ((42) & 63)), F1: 65535, F2: -2147483648}
    v24 := GenS2{F0: 127, F1: 170}
    switch 128 {
    case 4294967295:
        v5 -= v5
    case 127:
        if (!((false || false) || false)) {
            v1 -= (v1 << ((510) & 63))
            var v25 uint64 = 0
            var v26 uint = 2
            _ = v25
            _ = v26
        }
    }
    delete(v2, 42)
    switch ("a" + "hello") {
    case "a":
        v1 = v1
    default:
        fmt.Println(v4)
    }
    var v27 uint32 = 0
    for fi28 := 0; fi28 < 16; fi28++ {
        _ = len(v2)
        if false {
            var v29 int64 = -1
            for fi30 := 0; fi30 < 20; fi30++ {
                fmt.Println(v5)
                fmt.Println(v4, v29, fi30)
                switch (fi30 << ((0) & 63)) {
                case 46:
                    var v31 int16 = 0
                    if len(v3) > 0 {
                        v3[0] = fi28
                    }
                    _ = v31
                default:
                    var v32 int64 = v29
                    _ = v32
                }
                _ = fi30
            }
            _ = v29
        } else {
            var v33 int16 = 239
            _ = v33
        }
        _ = fi28
    }
    var v34 bool = false
    v34 = (1000 <= 1000)
    var v35 string = "foo bar"
    v34 = (v27 < 433)
    fmt.Println(v35)
    for fi36 := 0; fi36 < 20; fi36++ {
        switch "hello" {
        case "\n\t":
            switch -869 {
            case 0:
                v37 := GenS3{F0: v27}
                _ = v37
            case 32767:
                for fi38 := 0; fi38 < 3; fi38++ {
                    var v39 string = ((((v35 + ("\n\t" + v35)) + ("a" + v35)) + "a") + v35)
                    if v34 {
                        _ = len(v3)
                        fmt.Println(v34)
                        var v40 int = 1
                        var v41 int8 = -1
                        v27 -= (^(v27))
                        _ = v40
                        _ = v41
                    } else {
                        if false {
                            var v42 int8 = 0
                            fmt.Println(v42, v4)
                            var v43 int64 = 0
                            _ = v42
                            _ = v43
                        }
                    }
                    v39 = v35
                    _ = fi38
                    _ = v39
                }
            default:
                var v44 uint64 = (v1 + (^(v1)))
                _ = v44
            }
        case "a":
            v35 = ""
        }
        if len(v3) > 0 {
            v3[0] = (fi36 * (fi36 >> ((2) & 63)))
        }
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
