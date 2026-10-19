package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 string; F2 int }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := []uint32{2147483648, 82, 90}
    var v4 uint8 = 128
    var v5 uint = 902
    v6 := make(map[uint32]bool)
    v5 = (v5 % (((v5 * 65535)) | 1))
    if true {
        var v7 float32 = (1e10 - -2.25)
        _ = v7
    }
    delete(v6, 0)
    var v8 uint = v5
    var v9 int = 42
    var v10 int8 = 42
    if (1.5 != 1e10) {
        var v11 uint16 = 0
        var v12 float32 = -2.25
        _ = v11
        _ = v12
    } else {
        fmt.Println(v9, v9)
        v13 := GenS1{F0: (v4 * (^(v4)))}
        switch "hello" {
        case "":
            if true {
                if len(v3) > 0 {
                    v3[0] = 65535
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = 199
                }
                for fi14 := 0; fi14 < 7; fi14++ {
                    fmt.Println(v5, v4)
                    v15 := GenS1{F0: v4}
                    if true {
                        var v16 uint8 = (v4 - 248)
                        var v17 int = (v9 & (v9 ^ v9))
                        _ = v16
                        _ = v17
                    } else {
                        var v18 uint64 = 0
                        _ = v18
                    }
                    _ = fi14
                    _ = v15
                }
                if true {
                    var v19 uint = 255
                    v20 := GenS1{F0: (v4 * (v4 >> ((646) & 63)))}
                    v8 -= v8
                    _ = v19
                    _ = v20
                }
            }
        }
        _ = v13
    }
    switch "foo bar" {
    case "\n\t":
        if (!true) {
            if (true && false) {
                if true {
                    for fi21 := 0; fi21 < 18; fi21++ {
                        delete(v1, 2147483647)
                        var v22 float32 = ((0.1 + float32(1.5)) - 3.14159)
                        _ = fi21
                        _ = v22
                    }
                    var v23 uint16 = 0
                    _ = v23
                } else {
                    switch "" {
                    case "a":
                        v24 := GenS1{F0: v4}
                        _ = v24
                    case "":
                        if ((v5 <= (v8 >> ((379) & 63))) && (!(!true))) {
                            if (-2.25 == 1.5) {
                                v10 = 1
                            }
                        }
                    }
                    v4 = v4
                    v10 = (v10 + 42)
                }
            }
            v3 = append(v3, 1)
        } else {
            if true {
                v25 := GenS2{F0: 151}
                var v26 float32 = 0.0
                var v27 uint16 = 1
                _ = v25
                _ = v26
                _ = v27
            } else {
                _ = len(v3)
                if false {
                    v28 := v2[(v9 % ((-2147483648) | 1))]
                    _ = v28
                }
                var v29 uint8 = 1
                _ = v29
            }
            if ((!((v9 >= v9) && false)) && false) {
                v30 := v2[(v9 & (v9 & (v9 % ((v9) | 1))))]
                v30 = (v30 + v9)
                _ = v30
            } else {
                v31 := GenS3{F0: 909, F1: "", F2: 469}
                _ = v31
            }
            switch "\n\t" {
            case "hello":
                fmt.Println(v5, v10, v5)
            case "":
                v9 -= v9
            }
        }
    case "a":
        switch 32768 {
        case -317:
            for fi32 := 0; fi32 < 2; fi32++ {
                if false {
                    for fi33 := 0; fi33 < 14; fi33++ {
                        _ = len(v3)
                        var v34 bool = true
                        v10 -= (v10 + (v10 | (^(v10))))
                        _ = fi33
                        _ = v34
                    }
                }
                if ((541 != 868) && true) {
                    fmt.Println(v8)
                    for fi35 := 0; fi35 < 14; fi35++ {
                        _ = len(v3)
                        switch -80 {
                        case 2147483647:
                            var v36 float64 = (3.14159 / 1e10)
                            _ = v36
                        }
                        fmt.Println(v10)
                        _ = fi35
                    }
                    fmt.Println(v9, v5, v10)
                }
                for fi37 := 0; fi37 < 3; fi37++ {
                    if false {
                        if false {
                            v8 = 0
                            var v38 uint16 = 255
                            var v39 float32 = 0.1
                            _ = v38
                            _ = v39
                        }
                    }
                    var v40 uint16 = 2
                    _ = len(v3)
                    var v41 bool = false
                    _ = fi37
                    _ = v40
                    _ = v41
                }
                _ = fi32
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
