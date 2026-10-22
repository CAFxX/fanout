package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 int16; F1 bool; F2 uint32 }
func genFn0() uint32 {
    switch 9223372036854775807 {
    case -581:
        var v5 uint = 0
        _ = v5
    case 128:
        var v6 int16 = 870
        _ = v6
    case 9223372036854775807:
        if true {
            var v7 string = ("a" + (("\n\t" + "foo bar") + (("" + "\n\t") + "hello")))
            v8 := []int{32768}
            _ = v7
            _ = v8
        }
    default:
        switch 32767 {
        case 841:
            if false {
                var v9 uint64 = 4294967295
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 uint32 = 0
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                switch ("" + ("" + "")) {
                case "\n\t":
                    var v12 int = -60
                    _ = v12
                case "hello":
                    var v13 string = ("" + ("hello" + "a"))
                    _ = v13
                case "a":
                    var v14 uint16 = 65535
                    _ = v14
                default:
                    switch (("hello" + "hello") + "a") {
                    case "a":
                        var v15 uint16 = 32767
                        _ = v15
                    default:
                        var v16 float32 = 1.5
                        _ = v16
                    }
                }
            }
        case -416:
            v17 := GenS1{F0: 42}
            _ = v17
        default:
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 uint16 = 42
                _ = fi18
                _ = v19
            }
        }
    }
    var v20 int32 = -1
    _ = v20
    return 2
}

func main() {
    var v1 float32 = 0.0
    var v2 float64 = (0.1 + float64(1e10))
    var v3 uint = 2
    var v4 float32 = 1.5
    switch "a" {
    case "foo bar":
        switch 832 {
        case 2147483647:
            var v21 float32 = v1
            _ = v21
        case -9223372036854775808:
            v2 = 0.1
        default:
            if (366 >= -170) {
                fmt.Println(v3)
                v22 := []int{9223372036854775807}
                _ = v22
            }
        }
    case "\n\t":
        v23 := GenS2{F0: v3}
        _ = v23
    default:
        var v24 uint = (v3 << ((0) & 63))
        _ = v24
    }
    var v25 int32 = 0
    fmt.Println(v3)
    var v26 uint64 = 42
    for fi27 := 0; fi27 < 6; fi27++ {
        for fi28 := 0; fi28 < 5; fi28++ {
            if ("" >= "\n\t") {
                v29 := GenS3{F0: 127, F1: 127, F2: -443}
                v30 := []int{fi27, (fi27 % ((fi27) | 1)), (-(fi28)), fi27}
                _ = v29
                _ = v30
            } else {
                v31 := GenS4{F0: (^(v26)), F1: 1, F2: v25}
                var v32 int8 = 100
                v33 := GenS1{F0: -804}
                _ = v31
                _ = v32
                _ = v33
            }
            v25 = (v25 * 2147483647)
            _ = fi28
        }
        _ = fi27
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        v35 := GenS4{F0: v26, F1: 255, F2: (v25 / ((v25) | 1))}
        _ = fi34
        _ = v35
    }
    v36 := GenS5{F0: -23}
    var v37 string = "a"
    var v38 string = v37
    if false {
        if false {
            for fi39 := 0; fi39 < 2; fi39++ {
                for fi40 := 0; fi40 < 17; fi40++ {
                    if (65535 <= 276) {
                        var v41 int16 = -1
                        var v42 uint64 = 42
                        _ = v41
                        _ = v42
                    }
                    _ = fi40
                }
                v43 := []int{}
                _ = fi39
                _ = v43
            }
            for fi44 := 0; fi44 < 3; fi44++ {
                var v45 uint64 = (v26 | (v26 >> (((v3 / (((v3 * (v3 << (((v3 / ((v3) | 1))) & 63)))) | 1))) & 63)))
                switch v37 {
                case "\n\t":
                    for fi46 := 0; fi46 < 17; fi46++ {
                        v45 = 18446744073709551615
                        _ = fi46
                    }
                case "":
                    v2 = (3.14159 * ((v2 - v2) + -2.25))
                default:
                    fmt.Println(v3, v1, v3)
                }
                _ = fi44
                _ = v45
            }
        } else {
            var v47 int64 = -9223372036854775808
            if false {
                for fi48 := 0; fi48 < 19; fi48++ {
                    v49 := []int{fi48, (-(fi48))}
                    _ = fi48
                    _ = v49
                }
            }
            v50 := GenS1{F0: v47}
            _ = v47
            _ = v50
        }
        v51 := GenS6{F0: 593, F1: (v25 >> ((65535) & 63))}
        v37 = ""
        _ = v51
    }
    if true {
        v25 = v25
    }
    v52 := GenS7{F0: 0, F1: false, F2: genFn0()}
    switch 1 {
    case 42:
        v37 = v37
    case 32767:
        var v53 uint8 = 187
        _ = v53
    default:
        var v54 int = 159
        _ = v54
    }
    for fi55 := 0; fi55 < 16; fi55++ {
        for fi56 := 0; fi56 < 18; fi56++ {
            var v57 uint = v3
            if (!true) {
                v3 = (v3 ^ v3)
            }
            _ = fi56
            _ = v57
        }
        var v58 int = fi55
        _ = fi55
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v52
}
