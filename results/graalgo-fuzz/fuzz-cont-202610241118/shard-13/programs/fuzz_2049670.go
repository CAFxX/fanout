package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int16 }
type GenS2 struct { F0 int8; F1 int32; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 int8; F2 string }

func main() {
    var v1 int8 = -128
    var v2 int = 2147483647
    var v3 int32 = 0
    var v4 uint32 = 4294967295
    v5 := make(map[string]bool)
    v6 := make(map[int]uint32)
    switch "foo bar" {
    case "a":
        if ((!true) && (false && true)) {
            fmt.Println(v2, v1, v3)
            switch -200 {
            case 128:
                v7 := GenS1{F0: 311, F1: false, F2: 237}
                _ = v7
            case 9223372036854775807:
                v8 := []int{}
                _ = v8
            case -458:
                _ = len(v6)
            }
            v3 += v3
        }
    default:
        if false {
            var v9 uint16 = 542
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        switch (fi10 * (fi10 % ((-169) | 1))) {
        case -1:
            var v11 int8 = v1
            _ = v11
        default:
            for fi12 := 0; fi12 < 18; fi12++ {
                switch "foo bar" {
                case "\n\t":
                    v13 := GenS1{F0: -725, F1: ((true || (!false)) || false), F2: 622}
                    _ = v13
                case "hello":
                    fmt.Println(fi10, v1)
                case "a":
                    v14 := []int{int(v1), fi10, (fi12 / ((v2) | 1)), (fi10 >> ((42) & 63))}
                    _ = v14
                }
                _ = fi12
            }
        }
        var v15 bool = true
        _ = fi10
        _ = v15
    }
    switch -330 {
    case 32767:
        var v16 uint8 = 128
        _ = v16
    case 127:
        var v17 uint = 1
        _ = v17
    default:
        if (!true) {
            var v18 int8 = (^(v1))
            _ = v18
        } else {
            var v19 string = "foo bar"
            _ = v19
        }
    }
    v20 := GenS1{F0: -1, F1: (0 != 492), F2: 640}
    var v21 uint16 = 42
    v4 = (^(v4))
    var v22 string = "\n\t"
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            v25 := GenS1{F0: 32767, F1: true, F2: 1000}
            v26 := v5[((v22 + (("" + ("\n\t" + v22)) + "foo bar")) + (((v22 + v22) + ("" + ("hello" + "a"))) + v22))]
            v2 -= (v2 + fi23)
            _ = fi24
            _ = v25
            _ = v26
        }
        v3 = v3
        v3 -= (v3 - v3)
        _ = fi23
    }
    if true {
        v27 := GenS2{F0: v1, F1: (v3 - (v3 * v3)), F2: 101}
        _ = v27
    }
    if true {
        for fi28 := 0; fi28 < 8; fi28++ {
            v1 -= (-(v1))
            _ = fi28
        }
        var v29 float64 = 1e10
        _ = v29
    } else {
        if true {
            for fi30 := 0; fi30 < 7; fi30++ {
                for fi31 := 0; fi31 < 14; fi31++ {
                    v21 += (^(v21))
                    fmt.Println(fi30, v21)
                    v32 := GenS3{F0: fi31}
                    _ = fi31
                    _ = v32
                }
                var v33 uint32 = v4
                v34 := GenS4{F0: (v1 - 127), F1: (v1 % ((0) | 1)), F2: ("" + "hello")}
                _ = fi30
                _ = v33
                _ = v34
            }
            if false {
                switch 2147483647 {
                case -624:
                    var v35 int32 = (v3 - -2147483648)
                    _ = v35
                }
                if (4294967295 < 0) {
                    for fi36 := 0; fi36 < 5; fi36++ {
                        v22 = "foo bar"
                        switch 4294967295 {
                        case 128:
                            fmt.Println(v1, v3, v3)
                        case 1:
                            var v37 uint64 = 42
                            _ = v37
                        case 127:
                            var v38 bool = true
                            _ = v38
                        default:
                            var v39 uint16 = v21
                            _ = v39
                        }
                        _ = fi36
                    }
                    if (false && true) {
                        v40 := []int{v2, (v2 | (v2 + 2147483647)), v2, (v2 >> ((302) & 63))}
                        for fi41 := 0; fi41 < 6; fi41++ {
                            var v42 uint16 = (v21 * v21)
                            _ = fi41
                            _ = v42
                        }
                        _ = v40
                    } else {
                        var v43 int = v2
                        _ = v43
                    }
                    v44 := GenS3{F0: (-(v2))}
                    _ = v44
                }
                var v45 float32 = ((1.5 - (1e10 * ((((float32(1e10) / 1.0) + 1.5) / 3.14159) / 3.14159))) / (0.0 - 1.5))
                _ = v45
            }
        } else {
            var v46 bool = false
            _ = v46
        }
        if (false && false) {
            v21 -= (v21 + (v21 & v21))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
