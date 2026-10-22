package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int16 = 1000
    var v2 int8 = 17
    var v3 int16 = v1
    var v4 uint = 2
    v5 := make(map[uint32]bool)
    v2 = (v2 << (((^(v4))) & 63))
    var v6 int32 = 0
    v7 := GenS1{F0: 9223372036854775807}
    v8 := GenS1{F0: 32767}
    var v9 uint = (v4 & 1)
    v9 += 2
    v10 := GenS1{F0: -9223372036854775808}
    switch -9223372036854775808 {
    case -1:
        for fi11 := 0; fi11 < 11; fi11++ {
            v4 -= (v9 * 557)
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 int = fi11
                var v14 uint32 = 42
                var v15 uint32 = v14
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            v3 -= 32767
            _ = fi11
        }
    }
    switch "foo bar" {
    case "a":
        var v16 bool = false
        _ = v16
    case "":
        v2 = v2
    case "\n\t":
        v9 = v4
    }
    if false {
        for fi17 := 0; fi17 < 4; fi17++ {
            if false {
                var v18 float32 = ((3.14159 * 3.14159) * float32(1.5))
                v19 := GenS2{F0: 0, F1: v9}
                v18 -= 1e10
                _ = v18
                _ = v19
            } else {
                if true {
                    var v20 string = ""
                    v21 := []int{(fi17 >> (((v9 << ((4294967295) & 63))) & 63))}
                    switch (fi17 << (((v9 & v4)) & 63)) {
                    case 42:
                        fmt.Println(v20)
                    case 32767:
                        fmt.Println(fi17, v2)
                    case 1:
                        if false {
                            if ((-2.25 + (0.0 * 1.5)) == 0.0) {
                                if ((1 == 1) || false) {
                                    v20 = "hello"
                                }
                                if (72 < 128) {
                                    if true {
                                        var v22 uint8 = 137
                                        _ = v22
                                    }
                                    switch (fi17 ^ (^(fi17))) {
                                    case 68:
                                        fmt.Println(v2)
                                    case 192:
                                        v23 := GenS3{F0: v3}
                                        _ = v23
                                    case 42:
                                        if false {
                                            _ = len(v5)
                                        } else {
                                            v24 := GenS4{F0: 574, F1: v2}
                                            for fi25 := 0; fi25 < 1; fi25++ {
                                                if (42 >= 10) {
                                                    v26 := GenS4{F0: 42, F1: 22}
                                                    _ = v26
                                                } else {
                                                    v27 := v5[2147483648]
                                                    _ = v27
                                                }
                                                fmt.Println(v20)
                                                var v28 uint8 = 1
                                                _ = fi25
                                                _ = v28
                                            }
                                            _ = v24
                                        }
                                    default:
                                        for fi29 := 0; fi29 < 20; fi29++ {
                                            var v30 uint = 942
                                            switch "hello" {
                                            case "":
                                                fmt.Println(v30, v6, fi17)
                                            case "hello":
                                                if len(v21) >= 2 {
                                                    v21 = v21[1:len(v21)]
                                                }
                                            }
                                            fmt.Println(fi29)
                                            _ = fi29
                                            _ = v30
                                        }
                                    }
                                    for fi31 := 0; fi31 < 12; fi31++ {
                                        v32 := GenS2{F0: 2, F1: 227}
                                        _ = fi31
                                        _ = v32
                                    }
                                } else {
                                    fmt.Println(v6, v3)
                                }
                            }
                        } else {
                            v9 = v4
                            delete(v5, 2147483647)
                        }
                    }
                    _ = v20
                    _ = v21
                }
                _ = len(v5)
            }
            v5[4294967295] = true
            _ = fi17
        }
    } else {
        v2 = v2
        var v33 string = (("a" + "") + "foo bar")
        _ = v33
    }
    if false {
        var v34 float64 = (1e10 + (((1.5 / 3.14159) / -2.25) - -2.25))
        _ = v34
    } else {
        v2 -= v2
        for fi35 := 0; fi35 < 11; fi35++ {
            var v36 int64 = 1
            switch "\n\t" {
            case "hello":
                v6 = 1
            case "foo bar":
                var v37 int = (fi35 << ((430) & 63))
                _ = v37
            default:
                var v38 int8 = (v2 / (((v2 >> (((^(v4))) & 63))) | 1))
                _ = v38
            }
            var v39 uint8 = 128
            _ = fi35
            _ = v36
            _ = v39
        }
        switch "" {
        case "":
            v6 = v6
        case "foo bar":
            for fi40 := 0; fi40 < 9; fi40++ {
                v4 = 984
                _ = fi40
            }
        case "hello":
            for fi41 := 0; fi41 < 6; fi41++ {
                switch "\n\t" {
                case "\n\t":
                    switch "hello" {
                    case "":
                        var v42 int8 = 100
                        _ = v42
                    default:
                        switch -499 {
                        case 256:
                            var v43 uint32 = 2
                            _ = v43
                        }
                    }
                case "a":
                    v1 = 1
                default:
                    v44 := v5[42]
                    _ = v44
                }
                fmt.Println(v9)
                _ = fi41
            }
        }
    }
    var v45 uint8 = uint8(v4)
    if (v2 == v2) {
        v4 = (v9 - v9)
    }
    if (uint32(v45) == 0) {
        for fi46 := 0; fi46 < 1; fi46++ {
            v47 := []int{-937}
            _ = fi46
            _ = v47
        }
        v3 -= (^(v3))
    } else {
        var v48 int = 1
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v45, v2, v6)
        }
        _ = v48
    }
    var v49 string = ""
    var v50 uint64 = 42
    for fi51 := 0; fi51 < 19; fi51++ {
        var v52 int16 = 32767
        _ = fi51
        _ = v52
    }
    if (!(v50 != (v50 / ((v50) | 1)))) {
        v49 = v49
        var v53 uint16 = 2
        switch "foo bar" {
        case "foo bar":
            switch v49 {
            case "":
                fmt.Println(v53, v4)
            }
        }
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
