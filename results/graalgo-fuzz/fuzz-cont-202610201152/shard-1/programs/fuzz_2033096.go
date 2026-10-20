package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string }

func main() {
    var v1 uint64 = 280
    v2 := make(map[uint32]string)
    var v3 int8 = 50
    var v4 int8 = (v3 + 100)
    var v5 bool = true
    var v6 int16 = 526
    v7 := GenS1{F0: 0, F1: (v4 >> ((65535) & 63))}
    v8 := []int{-1}
    v9 := GenS1{F0: (v1 ^ v1), F1: -128}
    v4 -= v3
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            v5 = v5
            v6 = (v6 << ((0) & 63))
            switch (^(fi11)) {
            case -176:
                if (2147483647 <= 1) {
                    if v5 {
                        var v12 int = fi10
                        var v13 uint = 819
                        v12 = (v12 & 2147483647)
                        v14 := v2[65535]
                        v15 := GenS2{F0: v3, F1: 2147483647, F2: v1}
                        v5 = ((1.5 * 1.5) <= 3.14159)
                        fmt.Println(v4)
                        _ = v12
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        var v16 int = (fi11 >> ((4294967295) & 63))
                        _ = v16
                    }
                    fmt.Println(v4)
                } else {
                    _ = len(v8)
                    var v17 int16 = -555
                    var v18 uint = 326
                    _ = v17
                    _ = v18
                }
            default:
                switch 127 {
                case 0:
                    var v19 bool = ((2 >= 0) || (0 > -2147483648))
                    _ = v19
                }
            }
            _ = fi11
        }
        _ = fi10
    }
    if ((v6 * (-(v6))) >= 32767) {
        if false {
            if v5 {
                v2[518] = "foo bar"
            }
        } else {
            v6 -= v6
            fmt.Println(v1, v1)
        }
        if ((v3 | -128) <= int8(v6)) {
            v5 = v5
        }
        if v5 {
            for fi20 := 0; fi20 < 12; fi20++ {
                fmt.Println(v4)
                var v21 uint64 = (v1 + 42)
                _ = fi20
                _ = v21
            }
            if v5 {
                if false {
                    v22 := GenS2{F0: v4, F1: 1, F2: (^(v1))}
                    v4 += v3
                    var v23 uint64 = (v1 << ((4294967295) & 63))
                    _ = v22
                    _ = v23
                }
            }
            for fi24 := 0; fi24 < 19; fi24++ {
                var v25 uint64 = (^(v1))
                var v26 uint = 2
                _ = fi24
                _ = v25
                _ = v26
            }
        } else {
            var v27 uint = 65535
            switch -9223372036854775808 {
            case 42:
                fmt.Println(v6)
            case 4294967295:
                fmt.Println(v4)
            case 1:
                v28 := GenS1{F0: (v1 ^ 1), F1: 100}
                _ = v28
            }
            fmt.Println(v27, v3, v5)
            _ = v27
        }
    }
    v4 += v3
    var v29 bool = (4294967295 >= 4294967295)
    v6 += 911
    var v30 float32 = 0.1
    v6 = int16(v4)
    v8 = append(v8, -9223372036854775808)
    v31 := GenS3{F0: v3, F1: 1, F2: (v1 & 4294967295)}
    v32 := GenS4{F0: 42, F1: false}
    switch 32768 {
    case 42:
        switch ("hello" + "\n\t") {
        case "\n\t":
            switch 32767 {
            case 9223372036854775807:
                for fi33 := 0; fi33 < 16; fi33++ {
                    _ = len(v8)
                    _ = fi33
                }
            case -73:
                if (1.5 != 1e10) {
                    v1 = (v1 % (((v1 / ((0) | 1))) | 1))
                    v1 = (v1 >> ((409) & 63))
                    v34 := GenS5{F0: "hello"}
                    _ = v34
                }
            case 4294967295:
                fmt.Println(v4, v5)
            default:
                if v5 {
                    for fi35 := 0; fi35 < 2; fi35++ {
                        for fi36 := 0; fi36 < 5; fi36++ {
                            v5 = (!(!v29))
                            v29 = false
                            _ = fi36
                        }
                        _ = fi35
                    }
                    var v37 int64 = 1
                    _ = v37
                }
            }
        case "foo bar":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
    case 9223372036854775807:
        v29 = v5
    default:
        v38 := GenS6{F0: "\n\t"}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
