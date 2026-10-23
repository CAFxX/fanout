package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 2147483647
    var v3 string = "a"
    if true {
        switch 844 {
        case 4294967295:
            switch 893 {
            case -125:
                v4 := GenS1{F0: true, F1: (v2 ^ v2)}
                _ = v4
            case -2147483648:
                var v5 int16 = 117
                _ = v5
            case -9223372036854775808:
                v1[42] = 1
            }
        case 32768:
            switch (v3 + "a") {
            case "a":
                for fi6 := 0; fi6 < 12; fi6++ {
                    switch 872 {
                    case 2147483647:
                        v2 = (v2 >> ((18446744073709551615) & 63))
                    case -2147483648:
                        var v7 int16 = 683
                        _ = v7
                    }
                    delete(v1, 0)
                    if true {
                        v8 := GenS1{F0: ("\n\t" >= (v3 + (v3 + "a"))), F1: v2}
                        v9 := GenS1{F0: false, F1: v2}
                        var v10 string = ""
                        _ = v10
                        _ = v8
                        _ = v9
                    }
                    _ = fi6
                }
            }
        case 127:
            v11 := GenS2{F0: 0, F1: 255, F2: 4294967295}
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 2; fi12++ {
            if (892 < 2) {
                if false {
                    var v13 int64 = 0
                    _ = v13
                }
                for fi14 := 0; fi14 < 1; fi14++ {
                    v2 = v2
                    switch (fi12 >> ((1) & 63)) {
                    case 32767:
                        v3 = ((v3 + "") + "\n\t")
                    default:
                        if false {
                            v2 = 2147483647
                        }
                    }
                    _ = fi14
                }
            }
            _ = fi12
        }
        var v15 string = ((("hello" + "a") + (v3 + "\n\t")) + ("hello" + "foo bar"))
        fmt.Println(v3, v2)
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v1[2] = 2147483648
        _ = fi16
    }
    delete(v1, 38)
    switch 32767 {
    case 42:
        switch "" {
        case "a":
            var v17 uint8 = 0
            _ = v17
        case "":
            switch 1 {
            case 2147483647:
                v2 -= (^(v2))
            case 68:
                for fi18 := 0; fi18 < 20; fi18++ {
                    v19 := GenS1{F0: false, F1: 1}
                    _ = fi18
                    _ = v19
                }
            case 32768:
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 int16 = -32768
                    _ = fi20
                    _ = v21
                }
            default:
                v3 = ("\n\t" + ((v3 + "\n\t") + v3))
            }
        }
    default:
        var v22 float32 = 1e10
        _ = v22
    }
    fmt.Println(v2, v3)
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := []int{9223372036854775807, (fi23 | 32768), fi23}
        var v25 float64 = 1e10
        v3 = ""
        _ = fi23
        _ = v24
        _ = v25
    }
    v3 = v3
    switch 32768 {
    case -94:
        if true {
            v3 = v3
            if true {
                v3 = (v3 + v3)
            }
        }
    case -868:
        var v26 float64 = 0.1
        _ = v26
    case 123:
        fmt.Println(v3)
    default:
        fmt.Println(v2)
    }
    v27 := v1[2]
    var v28 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
