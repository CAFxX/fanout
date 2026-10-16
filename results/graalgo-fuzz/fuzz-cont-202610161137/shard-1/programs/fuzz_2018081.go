package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint8) uint16 {
    p0 = (p0 >> ((255) & 63))
    if false {
        fmt.Println(p0)
        if (((!(true || true)) && true) || false) {
            v6 := []int{32768, -1}
            for fi7 := 0; fi7 < 7; fi7++ {
                for fi8 := 0; fi8 < 3; fi8++ {
                    var v9 bool = false
                    _ = fi8
                    _ = v9
                }
                p0 = (p0 / (((p0 - p0)) | 1))
                if (false || (-649 <= 1)) {
                    p0 = (p0 >> ((497) & 63))
                    for fi10 := 0; fi10 < 14; fi10++ {
                        fmt.Println(p0, fi10)
                        _ = fi10
                    }
                    if len(v6) > 0 {
                        v6[0] = fi7
                    }
                }
                _ = fi7
            }
            switch ("foo bar" + ("" + "hello")) {
            case "":
                var v11 int64 = 491
                _ = v11
            case "\n\t":
                p0 = p0
            case "foo bar":
                var v12 bool = (32767 != 959)
                _ = v12
            }
            _ = v6
        } else {
            switch ((("" + "") + "\n\t") + "\n\t") {
            case "a":
                if (!(0 > -2147483648)) {
                    var v13 float32 = 1.5
                    v14 := []int{}
                    _ = v13
                    _ = v14
                }
            case "":
                p0 -= p0
            case "\n\t":
                switch 9223372036854775807 {
                case 975:
                    var v15 uint8 = (p0 >> ((228) & 63))
                    _ = v15
                default:
                    v16 := GenS1{F0: 195, F1: -9223372036854775808}
                    _ = v16
                }
            default:
                p0 = 42
            }
            for fi17 := 0; fi17 < 2; fi17++ {
                p0 = (p0 * 128)
                p0 += (p0 | 172)
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 uint32 = 42
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                fmt.Println(p0, fi20)
                p0 -= p0
                _ = fi20
            }
        }
    } else {
        v21 := []int{}
        var v22 int16 = -1
        _ = v21
        _ = v22
    }
    var v23 int16 = 1000
    fmt.Println(p0, v23)
    _ = p0
    _ = v23
    return 642
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]uint32)
    var v3 uint64 = 4294967295
    v4 := make(map[int]uint32)
    var v5 uint32 = 1
    v24 := v4[32768]
    v24 = v5
    v25 := GenS1{F0: genFn0(127), F1: 350}
    for fi26 := 0; fi26 < 14; fi26++ {
        v24 -= v24
        v27 := GenS1{F0: 496, F1: 1}
        _ = fi26
        _ = v27
    }
    var v28 uint = 255
    for fi29 := 0; fi29 < 16; fi29++ {
        switch ("foo bar" + ("hello" + (("\n\t" + "hello") + "\n\t"))) {
        case "foo bar":
            switch ((((("\n\t" + "a") + "") + (("foo bar" + "foo bar") + ("foo bar" + "foo bar"))) + "a") + "hello") {
            case "":
                v1 = 32767
            case "hello":
                switch fi29 {
                case 331:
                    var v30 int16 = -1
                    _ = v30
                }
            }
        default:
            v5 = uint32(v28)
        }
        v5 -= 600
        _ = fi29
    }
    if false {
        v5 -= (v5 % ((v24) | 1))
        var v31 int8 = 0
        _ = v31
    } else {
        var v32 bool = true
        if v32 {
            for fi33 := 0; fi33 < 14; fi33++ {
                for fi34 := 0; fi34 < 2; fi34++ {
                    v28 += v28
                    _ = fi34
                }
                var v35 int16 = 1000
                _ = fi33
                _ = v35
            }
        }
        _ = v32
    }
    delete(v2, -652)
    v36 := []int{2147483647, 1}
    var v37 int64 = 42
    fmt.Println(v28)
    v24 = (v24 ^ 1)
    switch "\n\t" {
    case "hello":
        delete(v2, -992)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
