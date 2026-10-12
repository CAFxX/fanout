package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int8; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int8; F2 int }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint32 }

func main() {
    var v1 int64 = -1
    var v2 int64 = -799
    var v3 int64 = -1
    var v4 int16 = 32767
    fmt.Println(v1, v3)
    if ((!true) && false) {
        var v5 string = "hello"
        _ = v5
    }
    fmt.Println(v3, v4, v4)
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: (v2 * v1)}
        if (2 < 42) {
            v8 := []int{}
            _ = v8
        }
        var v9 int = -9223372036854775808
        _ = fi6
        _ = v7
        _ = v9
    }
    switch "a" {
    case "\n\t":
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := []int{fi10}
            _ = fi10
            _ = v11
        }
    case "a":
        v4 -= v4
    case "foo bar":
        var v12 bool = false
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := []int{fi13, fi13, (fi13 >> ((409) & 63))}
        switch ("hello" + ((("\n\t" + ((("hello" + "foo bar") + "hello") + "a")) + "a") + "hello")) {
        case "\n\t":
            if ("a" <= "") {
                var v15 int32 = -370
                v16 := GenS2{F0: 32767, F1: 1, F2: 127}
                _ = v15
                _ = v16
            } else {
                var v17 float32 = float32(3.14159)
                switch 2147483647 {
                case 127:
                    _ = len(v14)
                }
                _ = v17
            }
        case "a":
            v18 := GenS1{F0: (v3 << ((255) & 63))}
            _ = v18
        default:
            if false {
                for fi19 := 0; fi19 < 10; fi19++ {
                    _ = len(v14)
                    switch (^(fi19)) {
                    case 9223372036854775807:
                        var v20 int64 = v2
                        _ = v20
                    case 0:
                        if (42 >= 318) {
                            v3 = v2
                            if true {
                                if len(v14) > 0 {
                                    v14[0] = fi13
                                }
                            } else {
                                var v21 uint32 = 623
                                var v22 int32 = 1
                                _ = v21
                                _ = v22
                            }
                        }
                    default:
                        v2 = v2
                    }
                    switch "hello" {
                    case "hello":
                        var v23 int64 = -473
                        _ = v23
                    default:
                        v4 = 79
                    }
                    _ = fi19
                }
                fmt.Println(v4)
                v1 = v3
            }
        }
        if false {
            _ = len(v14)
            for fi24 := 0; fi24 < 16; fi24++ {
                if true {
                    switch (-(fi24)) {
                    case -318:
                        v4 += v4
                    case -1:
                        var v25 int32 = 0
                        _ = v25
                    case 493:
                        v26 := GenS1{F0: (v1 - (v3 + (v3 | -60)))}
                        v4 = (v4 * v4)
                        _ = v26
                    default:
                        v4 = (v4 + 32767)
                    }
                    fmt.Println(v3, v4, v2)
                }
                _ = fi24
            }
            switch (fi13 - fi13) {
            case 0:
                if len(v14) > 0 {
                    v14[0] = 32768
                }
            default:
                var v27 uint8 = 29
                _ = v27
            }
        } else {
            var v28 uint16 = 693
            _ = v28
        }
        _ = fi13
        _ = v14
    }
    var v29 bool = (((true || (false && true)) || false) || false)
    v3 += (v3 | 271)
    switch 2147483647 {
    case -819:
        if ((60 != 0) || false) {
            if (!v29) {
                if (!(42 != 2)) {
                    if (!(42 <= 92)) {
                        var v30 uint = 1
                        if v29 {
                            fmt.Println(v1, v3)
                            v3 = -1
                        }
                        _ = v30
                    }
                } else {
                    switch 9223372036854775807 {
                    case 4294967295:
                        switch ("hello" + "hello") {
                        case "foo bar":
                            var v31 int8 = 42
                            _ = v31
                        default:
                            v32 := []int{-415}
                            _ = v32
                        }
                    }
                }
            }
        }
    case 32768:
        switch 368 {
        case -832:
            v33 := GenS3{F0: -588}
            _ = v33
        }
    case 2147483647:
        v34 := GenS4{F0: 100, F1: 10, F2: 407}
        _ = v34
    }
    if (-26 == -128) {
        v35 := []int{4294967295, 1}
        v1 += (v2 ^ v2)
        v3 = v3
        _ = v35
    } else {
        v36 := []int{42, 32768, -967, -773}
        for fi37 := 0; fi37 < 10; fi37++ {
            v38 := GenS3{F0: -2147483648}
            for fi39 := 0; fi39 < 8; fi39++ {
                fmt.Println(fi37, v29)
                v40 := GenS3{F0: int32(fi39)}
                v41 := GenS5{F0: 255, F1: 4294967295}
                _ = fi39
                _ = v40
                _ = v41
            }
            v42 := GenS5{F0: 65535, F1: 586}
            _ = fi37
            _ = v38
            _ = v42
        }
        _ = v36
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        v4 += 281
        _ = fi43
    }
    var v44 string = "foo bar"
    if v29 {
        var v45 uint8 = 42
        switch "a" {
        case "\n\t":
            for fi46 := 0; fi46 < 2; fi46++ {
                var v47 uint32 = 2147483647
                _ = fi46
                _ = v47
            }
        case "a":
            var v48 uint16 = 32767
            _ = v48
        }
        fmt.Println(v1, v29, v29)
        _ = v45
    } else {
        var v49 float32 = 1e10
        for fi50 := 0; fi50 < 11; fi50++ {
            v51 := []int{1, fi50, (fi50 - (fi50 >> ((65535) & 63))), fi50}
            v52 := GenS4{F0: -6, F1: 100, F2: fi50}
            _ = fi50
            _ = v51
            _ = v52
        }
        switch 2147483647 {
        case 4294967295:
            v4 = v4
        case 420:
            if (1 == 0) {
                var v53 int8 = 0
                _ = v53
            }
        case 128:
            v54 := []int{}
            _ = v54
        }
        _ = v49
    }
    var v55 uint16 = 0
    switch v44 {
    case "a":
        var v56 float32 = 1.5
        _ = v56
    case "hello":
        switch v44 {
        case "a":
            if true {
                switch v44 {
                case "":
                    v1 = (v1 >> ((42) & 63))
                case "\n\t":
                    switch "" {
                    case "hello":
                        var v57 uint16 = v55
                        _ = v57
                    case "foo bar":
                        var v58 float64 = 0.1
                        _ = v58
                    case "a":
                        for fi59 := 0; fi59 < 9; fi59++ {
                            v60 := GenS6{F0: 1}
                            _ = fi59
                            _ = v60
                        }
                    default:
                        var v61 int = 128
                        _ = v61
                    }
                case "a":
                    v62 := []int{32768, 9223372036854775807, 128, 1}
                    _ = v62
                }
                var v63 bool = v29
                var v64 float64 = (float64(3.14159) - float64((((((1.5 / 1e10) / ((1.5 - (-2.25 * (float32(3.14159) / 1.5))) + -2.25)) + 1e10) + float32(0.1)) * float32((1.5 + -2.25)))))
                _ = v63
                _ = v64
            }
        case "\n\t":
            v65 := []int{-863}
            _ = v65
        default:
            var v66 float64 = 3.14159
            _ = v66
        }
    case "\n\t":
        v4 = 42
    default:
        if ((v4 ^ v4) != v4) {
            var v67 uint8 = 1
            _ = v67
        }
    }
    var v68 int32 = -111
    fmt.Println(v1, v29, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v55:", v55)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v44
    _ = v55
    _ = v68
}
