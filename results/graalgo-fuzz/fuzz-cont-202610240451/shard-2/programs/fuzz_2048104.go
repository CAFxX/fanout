package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 int16 }
func genFn0(p0 int16) uint64 {
    v4 := GenS1{F0: 540}
    v5 := []int{-2147483648}
    _ = p0
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    var v1 int = 466
    var v2 int = (v1 - (v1 & v1))
    v3 := make(map[string]uint32)
    var v6 uint64 = 18446744073709551615
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS2{F0: (!(false || true)), F1: 950, F2: 2147483647}
        if (!(-1 <= 9223372036854775807)) {
            v2 = (v1 * 9223372036854775807)
            switch 1 {
            case -644:
                v9 := GenS1{F0: 65535}
                _ = v9
            case -9223372036854775808:
                v2 += (fi7 / ((-1) | 1))
            case 4294967295:
                if true {
                    if (2 > 1) {
                        v1 -= (v1 & fi7)
                        v6 = 1
                        for fi10 := 0; fi10 < 15; fi10++ {
                            v11 := GenS3{F0: 164}
                            _ = fi10
                            _ = v11
                        }
                    }
                }
            }
        }
        _ = fi7
        _ = v8
    }
    v1 -= v1
    var v12 uint = 0
    v12 += (v12 >> ((v12) & 63))
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint8 = 128
            if (genFn0(-1) < (v6 & (v6 & uint64(v15)))) {
                switch "foo bar" {
                case "":
                    var v16 float32 = (3.14159 + 3.14159)
                    _ = v16
                case "a":
                    var v17 int64 = -1
                    _ = v17
                case "\n\t":
                    switch (v2 / ((127) | 1)) {
                    case 32768:
                        var v18 uint32 = 0
                        _ = v18
                    case -1:
                        switch "\n\t" {
                        case "\n\t":
                            var v19 int64 = 1
                            _ = v19
                        default:
                            v2 += -452
                        }
                    case 42:
                        v20 := GenS3{F0: 130}
                        _ = v20
                    default:
                        v6 = 552
                    }
                }
                if ((v15 - 128) == (v15 + v15)) {
                    if true {
                        if (!true) {
                            v21 := GenS3{F0: (v15 % (((^(v15))) | 1))}
                            v3[(("\n\t" + "foo bar") + "hello")] = 228
                            v15 = v15
                            _ = v21
                        } else {
                            v12 += 65535
                            v3["hello"] = 65535
                            fmt.Println(v12, v2)
                        }
                        delete(v3, "\n\t")
                    } else {
                        v12 = 42
                        v22 := []int{32767, -2147483648}
                        _ = v22
                    }
                }
                var v23 string = "foo bar"
                _ = v23
            }
            var v24 float64 = (1e10 / 1.0)
            _ = fi14
            _ = v15
            _ = v24
        }
        var v25 uint8 = 255
        _ = fi13
        _ = v25
    }
    var v26 uint8 = 127
    switch "a" {
    case "hello":
        v12 = (v12 ^ uint(v26))
    case "foo bar":
        v27 := []int{v2, v2, 559}
        _ = v27
    case "a":
        for fi28 := 0; fi28 < 20; fi28++ {
            var v29 uint8 = v26
            var v30 int8 = -105
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 uint64 = genFn0(-32768)
        var v33 int = v2
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 int16 = 32767
    v35 := GenS4{F0: (^(v12))}
    switch 2147483647 {
    case 9223372036854775807:
        v1 += v1
    case -2147483648:
        var v36 int16 = v34
        _ = v36
    case -782:
        v26 = 42
    default:
        var v37 string = ""
        _ = v37
    }
    v6 = v6
    for fi38 := 0; fi38 < 1; fi38++ {
        v39 := GenS1{F0: 435}
        if false {
            if false {
                v2 += fi38
                v34 -= 42
                var v40 int8 = 0
                _ = v40
            } else {
                v2 = 0
                var v41 uint8 = (^(v26))
                _ = v41
            }
            var v42 int64 = 9223372036854775807
            for fi43 := 0; fi43 < 2; fi43++ {
                var v44 uint = uint(v26)
                v12 = (v44 + 891)
                v45 := []int{(fi38 >> (((v44 % ((v12) | 1))) & 63)), -764}
                _ = fi43
                _ = v44
                _ = v45
            }
            _ = v42
        } else {
            fmt.Println(v34)
            v34 = v34
        }
        fmt.Println(v1, v1)
        _ = fi38
        _ = v39
    }
    for fi46 := 0; fi46 < 19; fi46++ {
        v47 := GenS5{F0: 1}
        v3[""] = 65535
        fmt.Println(v34, v1)
        _ = fi46
        _ = v47
    }
    if (!(true && true)) {
        v26 = v26
        v12 -= v12
        v34 = (-(v34))
    } else {
        var v48 int16 = (v34 ^ v34)
        _ = v48
    }
    v49 := GenS6{F0: (!(!((((0.1 - -2.25) + float64((-2.25 / 1.5))) + (1e10 / 1.5)) > (1e10 + 3.14159)))), F1: int16(v26)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v49
    _ = v6
}
