package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 bool; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 int64 }

func main() {
    v1 := []int{931, 13}
    var v2 int = 128
    var v3 uint64 = 42
    v4 := GenS1{F0: 2147483648, F1: 24, F2: 2147483647}
    switch "foo bar" {
    case "\n\t":
        v3 = 211
    case "hello":
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 int8 = -106
            fmt.Println(v2, v6)
            _ = fi5
            _ = v6
        }
    case "a":
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 int16 = 32767
            var v9 int16 = v8
            switch int(v9) {
            case 32768:
                var v10 uint = 65535
                _ = v10
            case -561:
                fmt.Println(v8, v8, fi7)
            case -853:
                var v11 uint64 = v3
                _ = v11
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v2 = 2147483647
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(v3)
        var v13 bool = (true || false)
        var v14 int8 = 0
        _ = fi12
        _ = v13
        _ = v14
    }
    v2 = 128
    var v15 uint = 480
    switch v2 {
    case -2147483648:
        v16 := GenS1{F0: 0, F1: uint8(v15), F2: -2147483648}
        _ = v16
    case 2147483647:
        if (!false) {
            if false {
                v17 := GenS2{F0: v2}
                for fi18 := 0; fi18 < 4; fi18++ {
                    v19 := GenS3{F0: 255, F1: true, F2: -993}
                    v20 := GenS1{F0: 2, F1: 255, F2: -1}
                    if false {
                        var v21 int16 = 32767
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = v17
            } else {
                v22 := GenS4{F0: 2}
                _ = v22
            }
            v23 := GenS5{F0: (0 < 807), F1: int64(v2)}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v23
        } else {
            var v24 float64 = 1e10
            for fi25 := 0; fi25 < 9; fi25++ {
                if false {
                    switch "\n\t" {
                    case "a":
                        if (1e10 > v24) {
                            _ = len(v1)
                        }
                    case "foo bar":
                        for fi26 := 0; fi26 < 8; fi26++ {
                            if ((127 != 255) && true) {
                                var v27 uint16 = 153
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v28 := GenS3{F0: (v27 ^ v27), F1: false, F2: -32768}
                                var v29 uint = 2
                                v30 := GenS1{F0: 2, F1: 1, F2: -873}
                                _ = v27
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            v2 = (^(fi25))
                            v31 := GenS4{F0: 131}
                            _ = fi26
                            _ = v31
                        }
                    }
                    var v32 int64 = 42
                    _ = v32
                }
                _ = fi25
            }
            for fi33 := 0; fi33 < 4; fi33++ {
                var v34 float32 = 0.1
                var v35 string = "a"
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = v24
        }
    default:
        for fi36 := 0; fi36 < 10; fi36++ {
            v15 = uint(v3)
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 string = ("" + "foo bar")
                v15 = 4294967295
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
    }
    switch (-(v2)) {
    case 42:
        var v39 int64 = -9223372036854775808
        _ = v39
    default:
        fmt.Println(v3)
    }
    if (42 < 1) {
        if true {
            fmt.Println(v2, v2)
            v2 += -2147483648
        }
    }
    var v40 int16 = 0
    if len(v1) > 0 {
        v1[0] = v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}
