package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 bool; F1 uint; F2 int64 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int8; F1 uint16; F2 uint32 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]string)
    var v3 int16 = 0
    var v4 int32 = 0
    var v5 uint16 = 709
    v6 := GenS1{F0: (v3 >> ((572) & 63)), F1: 255}
    v7 := GenS2{F0: -313, F1: -561}
    for fi8 := 0; fi8 < 5; fi8++ {
        switch "\n\t" {
        case "":
            if (v1 >= v1) {
                var v9 int = fi8
                v2["a"] = (v1 + (v1 + v1))
                if false {
                    _ = len(v2)
                    var v10 uint16 = 42
                    fmt.Println(v10, fi8, v5)
                    _ = v10
                }
                _ = v9
            }
        case "a":
            v11 := v2[v1]
            _ = v11
        }
        _ = fi8
    }
    v12 := GenS3{F0: false, F1: 65535, F2: 9223372036854775807}
    v1 = ""
    if false {
        delete(v2, v1)
        v13 := GenS1{F0: v3, F1: 143}
        v5 = v5
        _ = v13
    } else {
        switch v1 {
        case "hello":
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := GenS4{F0: v3, F1: uint64(v5)}
                v16 := []int{(fi14 + (fi14 >> ((0) & 63))), -2147483648, -1}
                switch (fi14 % ((int(v5)) | 1)) {
                case -975:
                    var v17 uint = 250
                    _ = v17
                case -597:
                    v3 += v3
                case -9223372036854775808:
                    var v18 int16 = v3
                    _ = v18
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            var v19 int = 1
            _ = v19
        }
        if (((!((!false) && (!false))) || true) || (!false)) {
            v20 := GenS5{F0: 100, F1: (v5 - (v5 >> ((42) & 63))), F2: 788}
            for fi21 := 0; fi21 < 1; fi21++ {
                if false {
                    var v22 int = fi21
                    v23 := GenS2{F0: 2147483647, F1: (v4 - (v4 >> ((18446744073709551615) & 63)))}
                    _ = v22
                    _ = v23
                } else {
                    if (1e10 >= 1.5) {
                        fmt.Println(v1)
                        v24 := GenS2{F0: 4294967295, F1: int32(v3)}
                        for fi25 := 0; fi25 < 4; fi25++ {
                            var v26 string = (((v1 + (v1 + v1)) + v1) + (v1 + "a"))
                            v1 = v1
                            _ = fi25
                            _ = v26
                        }
                        _ = v24
                    } else {
                        fmt.Println(v4)
                    }
                }
                if true {
                    v1 = ((("" + v1) + ("\n\t" + v1)) + v1)
                } else {
                    var v27 int16 = (v3 % ((-32768) | 1))
                    v3 -= v27
                    for fi28 := 0; fi28 < 10; fi28++ {
                        switch "hello" {
                        case "\n\t":
                            switch fi28 {
                            case -1:
                                var v29 uint = 255
                                _ = v29
                            }
                        case "a":
                            var v30 int = fi28
                            _ = v30
                        }
                        _ = fi28
                    }
                    _ = v27
                }
                _ = fi21
            }
            _ = v20
        }
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 uint64 = 0
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
