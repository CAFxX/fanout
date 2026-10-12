package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 string = ("foo bar" + ("foo bar" + "a"))
    v2 := make(map[string]string)
    var v3 int64 = 320
    v4 := make(map[int]uint32)
    var v5 int = 32767
    var v6 string = "a"
    fmt.Println(v1)
    switch v5 {
    case 32767:
        v7 := GenS1{F0: 0}
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 1}
        v6 = (((("a" + v6) + ("" + v1)) + v1) + (v1 + "hello"))
        fmt.Println(v1)
        _ = fi8
        _ = v9
    }
    if true {
        switch ("hello" + v6) {
        case "\n\t":
            fmt.Println(v6, v3)
        case "hello":
            for fi10 := 0; fi10 < 4; fi10++ {
                fmt.Println(v5)
                _ = fi10
            }
        case "":
            var v11 uint64 = 18446744073709551615
            _ = v11
        default:
            v12 := GenS1{F0: 675}
            _ = v12
        }
        var v13 int16 = 32767
        _ = v13
    }
    v5 -= (v5 + (v5 % (((-(v5))) | 1)))
    var v14 int = v5
    v14 -= (-(v14))
    var v15 uint32 = 734
    switch v6 {
    case "foo bar":
        v1 = v1
    }
    if false {
        for fi16 := 0; fi16 < 15; fi16++ {
            for fi17 := 0; fi17 < 19; fi17++ {
                v18 := GenS2{F0: "\n\t"}
                _ = fi17
                _ = v18
            }
            var v19 int = fi16
            _ = fi16
            _ = v19
        }
        var v20 uint16 = 372
        _ = v20
    } else {
        _ = len(v2)
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 int64 = 9223372036854775807
            switch v1 {
            case "a":
                if true {
                    if true {
                        v23 := GenS3{F0: v1, F1: 384}
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v14 -= (fi24 * 138)
                        switch (^(fi21)) {
                        case 127:
                            fmt.Println(v22, v14, fi24)
                        case 128:
                            v6 = v6
                        case 0:
                            v5 -= 4294967295
                        }
                        _ = fi24
                    }
                } else {
                    var v25 float32 = 1.5
                    for fi26 := 0; fi26 < 20; fi26++ {
                        if (false && (!((1.5 / (float32(0.1) / 1.5)) < ((v25 - 1.5) / v25)))) {
                            if (959 >= -514) {
                                var v27 uint32 = v15
                                _ = v27
                            } else {
                                if true {
                                    v22 += v3
                                    v25 = (float32(0.1) - float32(-2.25))
                                    v14 = v5
                                    var v28 uint16 = 0
                                    _ = v28
                                } else {
                                    delete(v4, -1)
                                }
                                v15 = (v15 >> ((1) & 63))
                            }
                        }
                        switch (v1 + "\n\t") {
                        case "":
                            v29 := GenS1{F0: 42}
                            _ = v29
                        default:
                            fmt.Println(v6)
                        }
                        switch (v14 / ((v5) | 1)) {
                        case 0:
                            var v30 float64 = -2.25
                            _ = v30
                        default:
                            var v31 int16 = -433
                            _ = v31
                        }
                        _ = fi26
                    }
                    _ = v25
                }
            case "hello":
                fmt.Println(v22, fi21)
            case "foo bar":
                for fi32 := 0; fi32 < 20; fi32++ {
                    switch (fi32 / ((fi21) | 1)) {
                    case 9223372036854775807:
                        var v33 int32 = -2147483648
                        _ = v33
                    }
                    _ = fi32
                }
            }
            switch ((v6 + "\n\t") + "\n\t") {
            case "foo bar":
                var v34 uint32 = v15
                _ = v34
            case "\n\t":
                v35 := []int{1, fi21}
                _ = v35
            }
            _ = fi21
            _ = v22
        }
        for fi36 := 0; fi36 < 16; fi36++ {
            delete(v2, ("hello" + v1))
            var v37 uint16 = 65535
            var v38 int16 = 32767
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    switch (v5 % ((v5) | 1)) {
    case 42:
        fmt.Println(v6)
    case -2147483648:
        switch (v6 + v1) {
        case "hello":
            var v39 string = v1
            _ = v39
        case "foo bar":
            if (363 >= 2) {
                switch v5 {
                case -9223372036854775808:
                    var v40 uint8 = 127
                    _ = v40
                case 467:
                    v6 = ("foo bar" + v1)
                case 32768:
                    v41 := GenS4{F0: 42}
                    _ = v41
                default:
                    v3 = v3
                }
            }
        case "\n\t":
            v6 = v1
        }
    case -1:
        switch 42 {
        case 4294967295:
            for fi42 := 0; fi42 < 15; fi42++ {
                v3 -= v3
                _ = fi42
            }
        case 761:
            switch (v14 * v5) {
            case 1:
                v43 := v2[("\n\t" + v6)]
                _ = v43
            }
        }
    default:
        var v44 float64 = 3.14159
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
