package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 int8; F1 uint32; F2 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 bool; F2 int32 }

func main() {
    var v1 int32 = 2147483647
    var v2 bool = true
    v3 := make(map[int]string)
    var v4 uint = 65535
    var v5 uint8 = 86
    var v6 int = 2147483647
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 float32 = (0.0 * 0.1)
        v9 := GenS1{F0: 42, F1: v2}
        v8 = ((v8 + v8) + ((v8 * ((0.0 / (1e10 + (0.0 - (v8 + 1.5)))) * 3.14159)) - v8))
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := GenS2{F0: (v5 ^ v5)}
        v4 = (^(v4))
        _ = fi10
        _ = v11
    }
    var v12 bool = ((1e10 - 1e10) != (float64(1.5) - -2.25))
    v13 := GenS3{F0: 0}
    var v14 uint16 = 32767
    v15 := []int{v6, v6, v6}
    if false {
        v12 = v2
        for fi16 := 0; fi16 < 17; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                if true {
                    v5 = (v5 ^ v5)
                    var v18 uint = v4
                    v6 += (v6 << (((v18 ^ v4)) & 63))
                    _ = v18
                }
                switch ("foo bar" + (("foo bar" + "a") + (("\n\t" + "foo bar") + "a"))) {
                case "foo bar":
                    switch "\n\t" {
                    case "hello":
                        delete(v3, (fi17 | fi17))
                    }
                case "hello":
                    if false {
                        fmt.Println(v5, v1, v14)
                        var v19 uint8 = (v5 % ((v5) | 1))
                        _ = v19
                    }
                case "a":
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                default:
                    var v20 uint32 = 1
                    _ = v20
                }
                _ = fi17
            }
            fmt.Println(v5, fi16)
            fmt.Println(v5, v1)
            _ = fi16
        }
    } else {
        _ = len(v15)
        switch 569 {
        case 433:
            for fi21 := 0; fi21 < 12; fi21++ {
                v12 = (v12 && v2)
                for fi22 := 0; fi22 < 18; fi22++ {
                    var v23 int8 = 100
                    v12 = false
                    v5 = 42
                    _ = fi22
                    _ = v23
                }
                _ = fi21
            }
        case 1:
            switch int(v14) {
            case 32767:
                if len(v15) > 0 {
                    v15[0] = 9223372036854775807
                }
            case 79:
                v6 = v6
            case -1:
                v24 := GenS3{F0: 2}
                _ = v24
            }
        }
        var v25 bool = (!(true && (v6 != v6)))
        _ = v25
    }
    var v26 int16 = 1
    var v27 int64 = int64(v1)
    var v28 int32 = (-(v1))
    v29 := GenS4{F0: ("a" + "foo bar")}
    var v30 int64 = v27
    for fi31 := 0; fi31 < 5; fi31++ {
        v27 -= 0
        fmt.Println(v30)
        v1 = (v1 / ((v28) | 1))
        _ = fi31
    }
    switch "foo bar" {
    case "foo bar":
        for fi32 := 0; fi32 < 17; fi32++ {
            switch "" {
            case "\n\t":
                var v33 string = ("\n\t" + ("\n\t" + "foo bar"))
                _ = v33
            case "a":
                delete(v3, fi32)
            case "foo bar":
                v34 := GenS5{F0: (v5 - 42), F1: -41}
                _ = v34
            default:
                for fi35 := 0; fi35 < 4; fi35++ {
                    fmt.Println(v14, v5, v12)
                    var v36 bool = v2
                    if v36 {
                        var v37 int32 = (v28 & (v1 << (((v4 | (v4 - (v4 / ((1) | 1))))) & 63)))
                        v38 := GenS6{F0: 1, F1: 2, F2: 100}
                        v12 = ((v4 | 42) == v4)
                        _ = v37
                        _ = v38
                    }
                    v26 += int16(v4)
                    _ = fi35
                    _ = v36
                }
            }
            _ = fi32
        }
    case "hello":
        if v12 {
            v39 := GenS7{F0: v4}
            for fi40 := 0; fi40 < 19; fi40++ {
                _ = len(v15)
                v41 := GenS8{F0: -32768, F1: (!true), F2: (v28 - v28)}
                var v42 uint32 = 0
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = v39
        }
    }
    switch ("\n\t" + "a") {
    case "\n\t":
        fmt.Println(v12, v26, v26)
    case "":
        for fi43 := 0; fi43 < 12; fi43++ {
            var v44 uint32 = 1
            v6 += int(v1)
            if ((v30 + v27) <= -109) {
                if v2 {
                    var v45 int = fi43
                    fmt.Println(v27, v26)
                    v4 = 1
                    _ = v45
                } else {
                    v28 = (v28 & v28)
                    v46 := GenS6{F0: -65, F1: v44, F2: int8(v14)}
                    _ = v46
                }
                var v47 uint64 = 692
                _ = v47
            } else {
                var v48 uint = v4
                v44 += (^(v44))
                if v12 {
                    v49 := GenS7{F0: 935}
                    v27 += v27
                    var v50 int = 127
                    _ = v49
                    _ = v50
                }
                _ = v48
            }
            _ = fi43
            _ = v44
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
