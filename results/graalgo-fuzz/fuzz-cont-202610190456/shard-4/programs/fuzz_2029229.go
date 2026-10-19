package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 int32; F1 uint8 }

func main() {
    var v1 uint16 = 32767
    var v2 uint32 = 0
    var v3 uint16 = 42
    if (!true) {
        var v4 int8 = -11
        var v5 int = 9223372036854775807
        var v6 uint32 = (v2 >> ((1) & 63))
        _ = v4
        _ = v5
        _ = v6
    }
    v2 += v2
    v1 = v3
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(fi7, v1, v1)
        _ = fi7
    }
    fmt.Println(v1, v2)
    v2 -= 0
    var v8 string = ("\n\t" + (("a" + ("" + "a")) + "foo bar"))
    var v9 uint16 = uint16(v2)
    var v10 int = 9223372036854775807
    v1 = (v9 ^ 0)
    var v11 uint16 = (^(v3))
    switch v8 {
    case "hello":
        switch v8 {
        case "\n\t":
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 float32 = 3.14159
                if (!false) {
                    v1 = (v11 % ((v1) | 1))
                    for fi14 := 0; fi14 < 9; fi14++ {
                        v11 += (v9 | (^(v1)))
                        v9 = (^(v1))
                        _ = fi14
                    }
                } else {
                    fmt.Println(v2, fi12, v13)
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v16 := GenS1{F0: v2}
                        v17 := []int{(fi12 / (((fi12 ^ v10)) | 1)), fi15, fi15, fi15}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    var v18 float32 = ((v13 - v13) + v13)
                    _ = v18
                }
                if true {
                    for fi19 := 0; fi19 < 3; fi19++ {
                        switch (fi12 ^ (-(fi19))) {
                        case 873:
                            v8 = "\n\t"
                        default:
                            v20 := []int{(fi19 | (fi19 + fi19))}
                            _ = v20
                        }
                        _ = fi19
                    }
                }
                _ = fi12
                _ = v13
            }
        }
    case "a":
        if false {
            if false {
                fmt.Println(v2)
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 uint32 = (^(v2))
                    var v23 float32 = float32(float64((3.14159 / 3.14159)))
                    v24 := []int{9223372036854775807, -610, (^(v10)), 42}
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            } else {
                for fi25 := 0; fi25 < 13; fi25++ {
                    v10 += int(v11)
                    _ = fi25
                }
                if ("\n\t" <= v8) {
                    v1 = v1
                    v9 += (v1 << ((2) & 63))
                    v3 += v11
                } else {
                    var v26 uint8 = 164
                    _ = v26
                }
            }
            for fi27 := 0; fi27 < 16; fi27++ {
                v8 = v8
                if (1 >= 0) {
                    v28 := GenS2{F0: -1, F1: 127}
                    fmt.Println(v8, v8)
                    _ = v28
                }
                v2 = 768
                _ = fi27
            }
        }
    case "foo bar":
        if true {
            switch v10 {
            case -9223372036854775808:
                var v29 int64 = 638
                _ = v29
            default:
                if true {
                    for fi30 := 0; fi30 < 2; fi30++ {
                        switch -1 {
                        case 42:
                            for fi31 := 0; fi31 < 17; fi31++ {
                                v2 = v2
                                _ = fi31
                            }
                        }
                        _ = fi30
                    }
                    fmt.Println(v11)
                    var v32 uint = 0
                    _ = v32
                } else {
                    switch -249 {
                    case 4294967295:
                        if (1e10 != float64(-2.25)) {
                            var v33 bool = (1000 == 0)
                            fmt.Println(v2, v1, v3)
                            _ = v33
                        } else {
                            v11 = (v3 & 2)
                        }
                    case -383:
                        if ((3.14159 - -2.25) > (-2.25 * ((1.5 * 0.0) - 1.5))) {
                            v1 = 1
                            var v34 int32 = -994
                            _ = v34
                        }
                    }
                    v35 := GenS3{F0: 42, F1: 0}
                    _ = v35
                }
            }
            v11 = (^(v9))
        } else {
            var v36 int = v10
            v8 = "hello"
            for fi37 := 0; fi37 < 6; fi37++ {
                var v38 uint8 = 128
                _ = fi37
                _ = v38
            }
            _ = v36
        }
    default:
        for fi39 := 0; fi39 < 7; fi39++ {
            for fi40 := 0; fi40 < 20; fi40++ {
                v41 := []int{(fi40 ^ fi40), (fi40 << ((0) & 63)), (fi40 * v10)}
                _ = fi40
                _ = v41
            }
            v42 := GenS4{F0: 1, F1: 127}
            if false {
                var v43 int8 = 42
                _ = v43
            }
            _ = fi39
            _ = v42
        }
    }
    for fi44 := 0; fi44 < 11; fi44++ {
        var v45 string = "foo bar"
        _ = fi44
        _ = v45
    }
    switch "\n\t" {
    case "\n\t":
        v46 := GenS3{F0: 0, F1: 695}
        _ = v46
    case "":
        for fi47 := 0; fi47 < 18; fi47++ {
            switch (v8 + v8) {
            case "a":
                switch v8 {
                case "":
                    v11 -= 65535
                case "\n\t":
                    v11 = (v3 | v3)
                }
            }
            var v48 int16 = -1
            _ = fi47
            _ = v48
        }
    case "a":
        v49 := []int{(^(v10))}
        _ = v49
    default:
        var v50 bool = (1000 != -608)
        _ = v50
    }
    v51 := GenS1{F0: 2147483648}
    var v52 int = v10
    var v53 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v53
    _ = v8
    _ = v9
}
