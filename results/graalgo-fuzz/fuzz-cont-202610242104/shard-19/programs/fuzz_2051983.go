package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 bool; F2 int64 }
type GenS4 struct { F0 uint }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint16 = 0
    var v3 int64 = (v1 | v1)
    v4 := make(map[uint32]int)
    v3 = (v1 * (-(v3)))
    for fi5 := 0; fi5 < 10; fi5++ {
        v3 += -583
        switch (("\n\t" + ("a" + "hello")) + (("" + "\n\t") + "a")) {
        case "foo bar":
            switch -1 {
            case 9223372036854775807:
                if (1e10 > 1e10) {
                    v6 := GenS1{F0: (v2 << ((2) & 63)), F1: 4294967295}
                    _ = v6
                } else {
                    for fi7 := 0; fi7 < 9; fi7++ {
                        if (!false) {
                            var v8 uint64 = 0
                            _ = v8
                        }
                        var v9 uint16 = (^(v2))
                        var v10 float32 = 0.1
                        _ = fi7
                        _ = v10
                        _ = v9
                    }
                    v11 := GenS1{F0: v2, F1: 1}
                    v12 := GenS2{F0: true}
                    _ = v11
                    _ = v12
                }
            case -2147483648:
                fmt.Println(v2, fi5, v1)
            }
        default:
            if true {
                _ = len(v4)
                var v13 bool = false
                v14 := GenS3{F0: fi5, F1: true, F2: 167}
                _ = v13
                _ = v14
            }
        }
        _ = fi5
    }
    if (true || (1 != -291)) {
        _ = len(v4)
    } else {
        var v15 int16 = -80
        v16 := GenS4{F0: 576}
        for fi17 := 0; fi17 < 9; fi17++ {
            if true {
                v15 = (v15 << ((65535) & 63))
                var v18 int64 = -1
                v1 = 9223372036854775807
                _ = v18
            }
            v15 += (v15 * (v15 & 10))
            for fi19 := 0; fi19 < 9; fi19++ {
                switch "a" {
                case "hello":
                    var v20 float64 = 3.14159
                    _ = v20
                case "\n\t":
                    v21 := GenS3{F0: fi19, F1: (-796 < 42), F2: v1}
                    _ = v21
                case "a":
                    fmt.Println(v15, v2)
                    var v22 int32 = 637
                    _ = v22
                default:
                    v2 -= v2
                }
                v23 := GenS4{F0: 899}
                var v24 uint = 18446744073709551615
                _ = fi19
                _ = v23
                _ = v24
            }
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    switch -887 {
    case -1:
        v25 := v4[2147483648]
        _ = v25
    case 32768:
        if (4294967295 < 305) {
            for fi26 := 0; fi26 < 18; fi26++ {
                v27 := []int{fi26, fi26, (fi26 - 4294967295), fi26}
                var v28 float64 = (0.0 + 1e10)
                _ = fi26
                _ = v27
                _ = v28
            }
        }
    default:
        fmt.Println(v1, v2, v1)
    }
    v29 := GenS2{F0: false}
    v30 := v4[4294967295]
    for fi31 := 0; fi31 < 13; fi31++ {
        v32 := v4[901]
        _ = fi31
        _ = v32
    }
    fmt.Println(v3)
    for fi33 := 0; fi33 < 20; fi33++ {
        v3 = v3
        _ = fi33
    }
    _ = len(v4)
    v30 = (v30 % ((-1) | 1))
    for fi34 := 0; fi34 < 16; fi34++ {
        v1 = (v3 | (v3 << ((505) & 63)))
        _ = fi34
    }
    switch (("foo bar" + "a") + "\n\t") {
    case "a":
        switch v30 {
        case 42:
            fmt.Println(v30, v3, v3)
        default:
            v35 := GenS4{F0: 1}
            _ = v35
        }
    case "":
        v36 := v4[uint32(v2)]
        _ = v36
    default:
        _ = len(v4)
    }
    if (!true) {
        v2 -= (^(v2))
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        var v38 int = v30
        v39 := GenS2{F0: (!(!false))}
        fmt.Println(v1, fi37, v30)
        _ = fi37
        _ = v38
        _ = v39
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
