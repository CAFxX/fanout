package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }

func main() {
    var v1 int8 = 1
    var v2 float32 = 0.1
    v3 := make(map[string]uint32)
    var v4 int32 = 190
    var v5 string = "\n\t"
    var v6 float64 = 1e10
    fmt.Println(v6)
    var v7 float64 = v6
    switch 1 {
    case 9223372036854775807:
        var v8 int8 = v1
        _ = v8
    case -9223372036854775808:
        v9 := GenS1{F0: 9223372036854775807, F1: 126, F2: 1}
        _ = v9
    }
    v10 := GenS1{F0: -9223372036854775808, F1: 682, F2: 0}
    if false {
        var v11 int64 = 716
        if true {
            switch -62 {
            case 30:
                var v12 int64 = v11
                _ = v12
            case -865:
                v13 := GenS1{F0: (^(v11)), F1: 18446744073709551615, F2: 255}
                _ = v13
            default:
                v14 := []int{-1, -512, 9223372036854775807, 128}
                _ = v14
            }
        } else {
            v15 := GenS1{F0: v11, F1: 2, F2: 128}
            _ = v15
        }
        _ = v11
    }
    var v16 uint16 = 0
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 uint = 760
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := GenS1{F0: -116, F1: 0, F2: 2}
            _ = fi20
            _ = v21
        }
        if false {
            v16 = 1
            fmt.Println(v16, fi17)
        }
        _ = fi17
    }
    fmt.Println(v1, v2, v1)
    var v22 float32 = (v2 - 1e10)
    var v23 int64 = 42
    fmt.Println(v5)
    v4 += v4
    v5 = (((v5 + ((v5 + "\n\t") + (v5 + v5))) + v5) + (("foo bar" + "hello") + v5))
    var v24 float64 = 0.0
    v25 := GenS1{F0: (v23 << ((2) & 63)), F1: 2, F2: 42}
    for fi26 := 0; fi26 < 7; fi26++ {
        if false {
            if false {
                var v27 int64 = v23
                _ = v27
            } else {
                switch (v5 + "hello") {
                case "foo bar":
                    var v28 uint8 = 255
                    _ = v28
                case "a":
                    var v29 uint64 = 0
                    _ = v29
                case "\n\t":
                    switch fi26 {
                    case 0:
                        v30 := v3[(v5 + v5)]
                        _ = v30
                    case 32767:
                        var v31 int = fi26
                        _ = v31
                    }
                default:
                    if false {
                        v32 := []int{}
                        v33 := GenS1{F0: (-(v23)), F1: 42, F2: 128}
                        _ = v32
                        _ = v33
                    }
                }
                if true {
                    for fi34 := 0; fi34 < 17; fi34++ {
                        var v35 string = "hello"
                        fmt.Println(v6)
                        _ = fi34
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 14; fi36++ {
                        if (!true) {
                            switch v5 {
                            case "hello":
                                v4 = v4
                            default:
                                var v37 int16 = 980
                                _ = v37
                            }
                            v38 := []int{fi36, (^(fi36)), fi36, (fi26 >> ((4294967295) & 63))}
                            var v39 string = v5
                            _ = v38
                            _ = v39
                        } else {
                            fmt.Println(v5)
                        }
                        v16 -= 754
                        var v40 uint16 = v16
                        _ = fi36
                        _ = v40
                    }
                } else {
                    v1 = v1
                }
                if (v2 >= v2) {
                    _ = len(v3)
                }
            }
            var v41 uint = 1
            v4 = v4
            _ = v41
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
