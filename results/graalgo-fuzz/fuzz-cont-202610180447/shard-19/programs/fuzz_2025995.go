package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint8 }
type GenS2 struct { F0 int; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int64; F1 int64; F2 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 int32; F1 bool; F2 uint }
func genFn0() uint64 {
    if (0 >= 42) {
        var v4 float64 = 0.1
        switch -1 {
        case 42:
            if false {
                v5 := []int{1, 32768, -9223372036854775808, 0}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v5
            }
        case 970:
            v6 := GenS1{F0: 147, F1: 127, F2: 42}
            _ = v6
        case 32767:
            var v7 int16 = 32767
            _ = v7
        }
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            var v8 uint8 = 186
            _ = v8
        case "":
            switch "foo bar" {
            case "":
                var v9 bool = true
                _ = v9
            case "foo bar":
                var v10 int16 = 0
                _ = v10
            case "\n\t":
                var v11 int16 = 32767
                _ = v11
            }
        }
        _ = v4
    } else {
        var v12 int32 = -1
        _ = v12
    }
    v13 := GenS2{F0: -1, F1: 0, F2: 42}
    for fi14 := 0; fi14 < 1; fi14++ {
        if false {
            var v15 uint8 = 128
            var v16 uint16 = 1
            _ = v15
            _ = v16
        } else {
            fmt.Println(fi14)
        }
        _ = fi14
    }
    _ = v13
    return 4294967295
}
func genFn1() uint16 {
    var v17 int32 = 2147483647
    var v18 uint16 = 65535
    v19 := []int{1, -9223372036854775808, -1, 127}
    var v20 uint8 = 15
    _ = v17
    _ = v18
    _ = v19
    _ = v20
    return (v18 << ((1) & 63))
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]bool)
    var v3 uint8 = v1
    v3 = 1
    v1 -= 220
    var v21 int64 = 258
    fmt.Println(v3)
    var v22 int16 = -32768
    v23 := []int{945, 32768}
    v1 = v1
    for fi24 := 0; fi24 < 2; fi24++ {
        for fi25 := 0; fi25 < 12; fi25++ {
            if true {
                v26 := GenS2{F0: 758, F1: (v1 / ((0) | 1)), F2: 127}
                v3 = v1
                _ = v26
            }
            if ((false || false) && false) {
                fmt.Println(v21)
                var v27 uint16 = 42
                var v28 int64 = v21
                _ = len(v2)
                fmt.Println(v28)
                _ = v27
                _ = v28
            }
            fmt.Println(fi24, v3)
            v21 -= (v21 / ((v21) | 1))
            var v29 uint16 = 2
            if len(v23) > 0 {
                v23[0] = fi25
            }
            var v30 int32 = 760
            v31 := GenS3{F0: v21, F1: v21, F2: 0}
            _ = fi25
            _ = v29
            _ = v30
            _ = v31
        }
        switch (fi24 & fi24) {
        case -875:
            var v32 bool = (!("" <= (((("hello" + (("\n\t" + ("a" + "")) + "hello")) + "hello") + "") + "\n\t")))
            _ = v32
        default:
            var v33 uint64 = 4294967295
            _ = v33
        }
        _ = fi24
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        var v35 int16 = 1
        for fi36 := 0; fi36 < 19; fi36++ {
            if (259 >= 4294967295) {
                fmt.Println(v21, v21, v22)
                v37 := GenS4{F0: true}
                switch 0 {
                case 0:
                    v1 = v1
                case 128:
                    var v38 int64 = 0
                    _ = v38
                }
                _ = v37
            } else {
                switch (^(fi36)) {
                case 2147483647:
                    if (2147483647 != -959) {
                        v39 := v2[468]
                        v1 += 168
                        v40 := GenS5{F0: 0}
                        _ = v39
                        _ = v40
                    } else {
                        fmt.Println(v21)
                        switch "" {
                        case "":
                            v22 = 32767
                        case "a":
                            var v41 int64 = (^(v21))
                            _ = v41
                        }
                    }
                }
            }
            if true {
                v42 := GenS5{F0: int32(fi34)}
                _ = v42
            } else {
                delete(v2, 463)
            }
            _ = fi36
        }
        v35 -= -355
        _ = fi34
        _ = v35
    }
    var v43 string = "\n\t"
    v44 := GenS6{F0: 32768, F1: false}
    v45 := GenS7{F0: 2147483647, F1: false, F2: 65535}
    v46 := v2[4294967295]
    var v47 int16 = (v22 ^ (v22 | 919))
    for fi48 := 0; fi48 < 12; fi48++ {
        fmt.Println(v46, v1)
        var v49 int = -147
        _ = fi48
        _ = v49
    }
    v21 += (v21 - -523)
    var v50 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v50
}
