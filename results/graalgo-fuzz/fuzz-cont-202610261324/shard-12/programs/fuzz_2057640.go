package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 int }
type GenS3 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 uint; F2 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -128
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    _ = len(v3)
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := []int{}
        fmt.Println(fi5, fi5)
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 16; fi8++ {
        v7 = (^(v7))
        var v9 uint32 = 1
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 int64 = -1
            var v12 uint32 = 437
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    if (-619 < 42) {
        fmt.Println(v7, v2)
        if false {
            switch "" {
            case "foo bar":
                v13 := []int{}
                _ = v13
            }
            if (((2 == 4294967295) && true) || (!(0.0 != 3.14159))) {
                if ((("a" + ("hello" + "\n\t")) + "foo bar") <= "a") {
                    v2 = v2
                    switch (("\n\t" + "foo bar") + "a") {
                    case "\n\t":
                        for fi14 := 0; fi14 < 19; fi14++ {
                            v1[(fi14 % ((fi14) | 1))] = "foo bar"
                            if true {
                                v2 = v2
                                v7 = (^(v7))
                                var v15 int8 = (v2 >> ((42) & 63))
                                _ = v15
                            }
                            _ = fi14
                        }
                    case "hello":
                        v16 := GenS1{F0: v2}
                        _ = v16
                    case "foo bar":
                        v7 = (^(v7))
                    default:
                        v17 := []int{2147483647}
                        _ = v17
                    }
                    v18 := v4[(("foo bar" + "\n\t") + "foo bar")]
                    _ = v18
                }
                switch ("hello" + (("\n\t" + "") + ("hello" + ""))) {
                case "foo bar":
                    if true {
                        var v19 uint8 = 2
                        v20 := []int{42, 128, 128}
                        _ = v19
                        _ = v20
                    }
                case "":
                    v21 := GenS2{F0: 2, F1: 0, F2: 42}
                    _ = v21
                default:
                    v22 := []int{}
                    _ = v22
                }
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                v24 := GenS1{F0: (v2 * -37)}
                _ = fi23
                _ = v24
            }
        }
    }
    v25 := []int{-2147483648, 4294967295, 127, -470}
    if false {
        delete(v1, 127)
    } else {
        fmt.Println(v7)
        var v26 string = "foo bar"
        v27 := GenS2{F0: 1, F1: 42, F2: 9223372036854775807}
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        if true {
            var v29 uint32 = 0
            if (!false) {
                var v30 int64 = 1
                var v31 uint8 = 128
                fmt.Println(fi28)
                _ = v30
                _ = v31
            } else {
                fmt.Println(v29, v7)
            }
            switch "hello" {
            case "a":
                _ = len(v1)
            case "":
                for fi32 := 0; fi32 < 2; fi32++ {
                    fmt.Println(fi32)
                    var v33 string = ""
                    _ = fi32
                    _ = v33
                }
            }
            _ = v29
        } else {
            v25 = append(v25, fi28)
            v7 += 976
            v7 = 2
        }
        v34 := GenS3{F0: 244, F1: 0, F2: v7}
        fmt.Println(v7, fi28, v7)
        _ = fi28
        _ = v34
    }
    switch 127 {
    case -460:
        v2 = v2
    case 0:
        for fi35 := 0; fi35 < 8; fi35++ {
            var v36 uint64 = 1
            _ = fi35
            _ = v36
        }
    case 349:
        for fi37 := 0; fi37 < 17; fi37++ {
            v2 += 127
            _ = fi37
        }
    }
    if (4294967295 < -80) {
        v38 := GenS4{F0: -1}
        fmt.Println(v7, v2)
        delete(v1, 32767)
        _ = v38
    } else {
        var v39 uint32 = 2147483647
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
        if (2147483647 > 631) {
            for fi40 := 0; fi40 < 11; fi40++ {
                fmt.Println(v2, fi40, fi40)
                _ = fi40
            }
            if len(v25) > 0 {
                v25[0] = 127
            }
            v7 -= v7
        } else {
            var v41 int32 = 1
            var v42 bool = (-128 != (^(v2)))
            _ = v41
            _ = v42
        }
        _ = v39
    }
    v7 += (v7 & v7)
    var v43 uint8 = 2
    v44 := GenS5{F0: 827, F1: 666, F2: v2}
    v45 := GenS4{F0: 0}
    for fi46 := 0; fi46 < 16; fi46++ {
        v43 = (v43 + v43)
        _ = fi46
    }
    var v47 int8 = v2
    v47 += v47
    var v48 int = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v47
    _ = v48
    _ = v7
}
