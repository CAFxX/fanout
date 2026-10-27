package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 float32 = 1.5
    v4 := make(map[string]uint32)
    var v5 int32 = 1000
    v6 := []string{"foo bar", "a", ""}
    for fi7 := 0; fi7 < 12; fi7++ {
        if false {
            v5 -= v5
        }
        v5 -= (v5 & v5)
        var v8 bool = false
        _ = fi7
        _ = v8
    }
    fmt.Println(v3, v3)
    var v9 int32 = -31
    switch "a" {
    case "hello":
        for fi10 := 0; fi10 < 5; fi10++ {
            if false {
                v11 := GenS1{F0: 285}
                v5 = (^(v9))
                _ = v11
            } else {
                switch ("foo bar" + "a") {
                case "\n\t":
                    switch -914 {
                    case -1:
                        fmt.Println(fi10)
                    case 4294967295:
                        for fi12 := 0; fi12 < 6; fi12++ {
                            v3 += float32(3.14159)
                            var v13 float32 = v3
                            _ = fi12
                            _ = v13
                        }
                    default:
                        fmt.Println(v3, v9, v9)
                    }
                default:
                    if (-97 > 48) {
                        v14 := GenS1{F0: 273}
                        switch "hello" {
                        case "foo bar":
                            for fi15 := 0; fi15 < 16; fi15++ {
                                v16 := GenS2{F0: 1, F1: 927}
                                fmt.Println(fi15)
                                _ = len(v4)
                                v17 := GenS2{F0: 0, F1: 65535}
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        }
                        _ = v14
                    }
                }
                for fi18 := 0; fi18 < 7; fi18++ {
                    v9 += v9
                    if len(v6) > 0 {
                        v6[0] = ("a" + "\n\t")
                    }
                    _ = fi18
                }
                if len(v6) > 0 {
                    v6[0] = ""
                }
            }
            delete(v1, 65535)
            var v19 int = -379
            _ = fi10
            _ = v19
        }
    case "a":
        v20 := v4[""]
        _ = v20
    default:
        v5 = -1
    }
    fmt.Println(v9, v5)
    if false {
        var v21 uint8 = 128
        var v22 int16 = 0
        _ = v21
        _ = v22
    } else {
        v3 -= v3
        for fi23 := 0; fi23 < 18; fi23++ {
            v9 += 997
            _ = fi23
        }
        v24 := GenS1{F0: 657}
        _ = v24
    }
    v3 = (1e10 + -2.25)
    var v25 string = ("a" + "\n\t")
    var v26 int = 398
    var v27 float64 = 1.5
    for fi28 := 0; fi28 < 14; fi28++ {
        v3 -= (((0.1 * 1.5) - float32(((((((v27 / (v27 * float64(((v3 - float32((v27 - v27))) * v3)))) + v27) * v27) - v27) + 1.5) * (v27 + v27)))) - (v3 + v3))
        _ = fi28
    }
    var v29 int32 = (v9 & v5)
    v30 := GenS3{F0: 108}
    fmt.Println(v29, v9)
    switch 4294967295 {
    case -220:
        for fi31 := 0; fi31 < 12; fi31++ {
            switch "\n\t" {
            case "hello":
                switch v26 {
                case 128:
                    var v32 float64 = v27
                    _ = v32
                }
            case "a":
                var v33 uint16 = 255
                _ = v33
            default:
                switch (v26 * (fi31 | fi31)) {
                case 4294967295:
                    _ = len(v1)
                case 9223372036854775807:
                    _ = len(v6)
                }
            }
            v25 = ((v25 + v25) + ("hello" + "\n\t"))
            for fi34 := 0; fi34 < 17; fi34++ {
                var v35 uint16 = 212
                _ = fi34
                _ = v35
            }
            _ = fi31
        }
    case 0:
        v9 += (v29 & v9)
    }
    if ((v3 + v3) > v3) {
        switch ((("a" + ("hello" + v25)) + "hello") + (v25 + (("foo bar" + "hello") + (v25 + v25)))) {
        case "a":
            var v36 int = (v26 / (((v26 ^ -604)) | 1))
            _ = v36
        case "":
            v37 := GenS2{F0: 1, F1: 984}
            _ = v37
        case "foo bar":
            v25 = v25
        default:
            v25 = (v25 + (v25 + v25))
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v38 uint64 = 334
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
