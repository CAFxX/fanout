package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint32 }

func main() {
    var v1 int16 = 1000
    var v2 uint64 = 42
    v3 := make(map[uint32]string)
    fmt.Println(v2, v2)
    if ("\n\t" >= (("a" + "a") + "hello")) {
        v4 := GenS1{F0: 42, F1: 1000}
        var v5 int8 = 24
        _ = v4
        _ = v5
    } else {
        switch 127 {
        case -698:
            var v6 uint16 = 32767
            _ = v6
        case 0:
            var v7 float64 = (float64((-2.25 / -2.25)) * 0.0)
            _ = v7
        }
        v1 -= 32767
    }
    var v8 uint64 = (^(v2))
    v2 -= 713
    var v9 float64 = 0.0
    for fi10 := 0; fi10 < 19; fi10++ {
        if (0.1 >= v9) {
            delete(v3, 2147483648)
        }
        if (0 > 1) {
            var v11 uint = 42
            v8 -= (v8 << ((v11) & 63))
            _ = v11
        }
        var v12 int = (fi10 ^ fi10)
        _ = fi10
        _ = v12
    }
    var v13 int8 = 100
    v8 += 2
    if false {
        fmt.Println(v9, v2)
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 string = "a"
            v16 := GenS1{F0: -1, F1: 0}
            fmt.Println(v15)
            _ = fi14
            _ = v15
            _ = v16
        }
        if ((float32(1.5) >= 3.14159) && (73 == v2)) {
            v17 := GenS1{F0: -2147483648, F1: 29}
            switch ("a" + ("\n\t" + "a")) {
            case "a":
                if false {
                    if ((false || true) && (1 >= 0)) {
                        if true {
                            v1 -= v1
                            v8 = 4294967295
                            v9 += v9
                        } else {
                            if true {
                                v18 := []int{50, 2147483647, 705, 32768}
                                fmt.Println(v13, v2, v9)
                                fmt.Println(v13, v8, v1)
                                _ = v18
                            } else {
                                for fi19 := 0; fi19 < 13; fi19++ {
                                    v8 = v8
                                    _ = fi19
                                }
                            }
                        }
                        var v20 bool = (v13 != -1)
                        _ = v20
                    }
                }
            case "hello":
                for fi21 := 0; fi21 < 18; fi21++ {
                    v22 := GenS1{F0: -322, F1: 86}
                    fmt.Println(v2, fi21, v13)
                    v2 -= (v8 | v8)
                    _ = fi21
                    _ = v22
                }
            default:
                var v23 bool = true
                _ = v23
            }
            _ = v17
        }
    }
    if true {
        var v24 int32 = 2147483647
        _ = v24
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        fmt.Println(v2, fi25, fi25)
        _ = fi25
    }
    delete(v3, 926)
    _ = len(v3)
    v26 := GenS1{F0: 0, F1: 2147483648}
    v27 := GenS2{F0: 856, F1: (1 != 65535), F2: 150}
    if false {
        fmt.Println(v2, v8, v9)
        v9 += 1.5
        v28 := GenS2{F0: 255, F1: (-2147483648 == 32767), F2: 2147483648}
        _ = v28
    } else {
        if (462 == -1) {
            v29 := GenS1{F0: 1, F1: 0}
            _ = v29
        } else {
            var v30 int64 = 0
            if (v30 > 0) {
                for fi31 := 0; fi31 < 6; fi31++ {
                    if false {
                        v8 += v2
                        var v32 float32 = -2.25
                        v33 := GenS2{F0: 2, F1: true, F2: 2147483647}
                        _ = v32
                        _ = v33
                    }
                    _ = fi31
                }
                var v34 uint8 = 128
                v35 := []int{1, 276}
                _ = v34
                _ = v35
            }
            _ = v30
        }
        if false {
            if (!(!(18446744073709551615 < 65535))) {
                for fi36 := 0; fi36 < 18; fi36++ {
                    v8 -= 633
                    var v37 float32 = 3.14159
                    _ = fi36
                    _ = v37
                }
            }
            var v38 uint32 = 610
            _ = v38
        } else {
            for fi39 := 0; fi39 < 14; fi39++ {
                var v40 int16 = -32768
                for fi41 := 0; fi41 < 5; fi41++ {
                    var v42 uint = 2
                    fmt.Println(v2, v9, v13)
                    _ = fi41
                    _ = v42
                }
                _ = fi39
                _ = v40
            }
            fmt.Println(v9, v1, v13)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v8
    _ = v9
}
