package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int32 = 1
    v2 := []uint32{4294967295, 232, 0}
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]int)
    v5 := []string{}
    if false {
        if (!(true || true)) {
            v6 := GenS1{F0: 255, F1: 0, F2: (v1 - (v1 ^ v1))}
            var v7 uint32 = 42
            _ = v6
            _ = v7
        }
        switch 534 {
        case -1:
            v3 = v3
        case 1:
            v8 := GenS1{F0: 128, F1: 103, F2: 2147483647}
            _ = v8
        }
        v9 := GenS1{F0: 0, F1: 1, F2: v1}
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if false {
            v3 = 42
        } else {
            fmt.Println(fi10, v3, v1)
            if len(v5) > 0 {
                v5[0] = ("foo bar" + ("a" + ""))
            }
            if len(v5) > 0 {
                v5[0] = ""
            }
        }
        var v11 int = fi10
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v3 = v3
        v13 := GenS2{F0: false, F1: 65535}
        var v14 int = 32768
        _ = fi12
        _ = v13
        _ = v14
    }
    switch ("a" + "") {
    case "":
        var v15 string = "hello"
        _ = v15
    case "foo bar":
        v16 := GenS2{F0: true, F1: 65535}
        _ = v16
    case "a":
        if (!true) {
            for fi17 := 0; fi17 < 13; fi17++ {
                _ = len(v2)
                var v18 uint16 = 32767
                v19 := GenS3{F0: 1000}
                _ = fi17
                _ = v18
                _ = v19
            }
            if false {
                for fi20 := 0; fi20 < 12; fi20++ {
                    switch (("foo bar" + "foo bar") + "\n\t") {
                    case "a":
                        if true {
                            v3 += v3
                        } else {
                            if (v3 >= (v3 >> ((65535) & 63))) {
                                v1 = 375
                                v21 := GenS3{F0: int16(v3)}
                                _ = v21
                            }
                            if (132 != 255) {
                                var v22 bool = (!true)
                                delete(v4, 763)
                                _ = v22
                            }
                        }
                    case "\n\t":
                        fmt.Println(v1)
                    case "hello":
                        var v23 uint8 = 132
                        _ = v23
                    default:
                        if (false && (538 < 412)) {
                            var v24 int = 32767
                            _ = v24
                        } else {
                            var v25 int16 = 927
                            _ = v25
                        }
                    }
                    if false {
                        delete(v4, 718)
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = len(v5)
                    }
                    _ = fi20
                }
            }
            var v26 float64 = 1.5
            _ = v26
        } else {
            v27 := GenS2{F0: true, F1: 2147483648}
            _ = v27
        }
    default:
        if (32767 == 646) {
            v3 += (-(v3))
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 int = fi28
                _ = fi28
                _ = v29
            }
        }
    }
    v1 = 0
    if (v3 < 1) {
        var v30 int = 2147483647
        fmt.Println(v1)
        _ = v30
    } else {
        switch -1 {
        case 127:
            v3 -= (v3 << ((34) & 63))
        case 1:
            var v31 uint32 = 324
            _ = v31
        case 42:
            switch -9223372036854775808 {
            case -897:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 898:
                var v32 float64 = float64((1e10 * (1.5 + (1e10 * float32(3.14159)))))
                _ = v32
            case 26:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        if true {
            var v33 int32 = 1
            _ = v33
        } else {
            var v34 int64 = -1
            fmt.Println(v3, v34, v3)
            fmt.Println(v1)
            _ = v34
        }
    }
    v3 -= (-(v3))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
