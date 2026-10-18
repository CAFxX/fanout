package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int }
func genFn0(p0 string, p1 int) int16 {
    fmt.Println(p1, p0)
    p0 = (p0 + p0)
    if true {
        switch p1 {
        case 118:
            v4 := GenS1{F0: true, F1: ("hello" + "hello")}
            _ = v4
        }
    }
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 int32 = -752
        v7 := []int{fi5, fi5, p1, p1}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: (18446744073709551615 != 2), F1: ""}
    _ = p0
    _ = p1
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{151}
    var v2 uint8 = 42
    var v3 int = -1
    fmt.Println(v2)
    switch v3 {
    case 32768:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 448:
        fmt.Println(v3, v3)
    default:
        v9 := GenS1{F0: (26 < 708), F1: ""}
        _ = v9
    }
    v1 = append(v1, 1)
    _ = len(v1)
    var v10 uint = 500
    var v11 uint = 4294967295
    v12 := GenS2{F0: v2, F1: 42}
    v13 := GenS2{F0: 0, F1: 762}
    v11 -= v11
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint = v10
        var v16 int32 = 2147483647
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v2 = (v2 / (((v2 & (v2 + v2))) | 1))
        v3 += fi17
        _ = fi17
    }
    if false {
        v18 := GenS3{F0: 2147483647}
        switch (("foo bar" + ("foo bar" + "\n\t")) + ("foo bar" + "\n\t")) {
        case "foo bar":
            if (1 == 988) {
                if (!(103 != 1)) {
                    v2 += (v2 % ((0) | 1))
                    var v19 string = ""
                    _ = v19
                }
            }
        }
        v1 = append(v1, 42)
        _ = v18
    } else {
        v3 += v3
        fmt.Println(v11, v10, v11)
        var v20 int32 = 57
        _ = v20
    }
    v3 = v3
    var v21 bool = true
    v22 := GenS2{F0: (v2 << (((v11 | v11)) & 63)), F1: 56}
    if ((!(-84 == -1)) || ((0.1 < 1.5) || (!(((v3 ^ v3) < v3) && (433 <= 760))))) {
        if true {
            var v23 float32 = 0.1
            v24 := GenS4{F0: 100, F1: (v10 >> ((v10) & 63))}
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 20; fi25++ {
                var v26 uint64 = 0
                var v27 float32 = 3.14159
                if (0 == -32768) {
                    for fi28 := 0; fi28 < 4; fi28++ {
                        var v29 int8 = 1
                        v10 += 4294967295
                        if true {
                            v30 := GenS2{F0: (v2 | v2), F1: 1}
                            _ = v30
                        }
                        _ = fi28
                        _ = v29
                    }
                    fmt.Println(v27, v11)
                } else {
                    var v31 uint = 459
                    var v32 int64 = -55
                    _ = v31
                    _ = v32
                }
                _ = fi25
                _ = v26
                _ = v27
            }
            for fi33 := 0; fi33 < 15; fi33++ {
                if v21 {
                    switch (^(v3)) {
                    case 32768:
                        for fi34 := 0; fi34 < 6; fi34++ {
                            fmt.Println(fi33)
                            switch "foo bar" {
                            case "hello":
                                _ = len(v1)
                            default:
                                var v35 int32 = -2147483648
                                _ = v35
                            }
                            _ = fi34
                        }
                    case 726:
                        v11 = v11
                    case 127:
                        v2 = (v2 * v2)
                    }
                    var v36 int = fi33
                    _ = v36
                } else {
                    if (0 == 79) {
                        fmt.Println(v10)
                    } else {
                        v37 := GenS5{F0: fi33}
                        _ = v37
                    }
                    fmt.Println(v3)
                }
                if ("foo bar" != "\n\t") {
                    _ = len(v1)
                } else {
                    _ = len(v1)
                    if v21 {
                        v21 = (!v21)
                        var v38 uint8 = v2
                        var v39 int8 = 16
                        _ = v38
                        _ = v39
                    } else {
                        v40 := GenS1{F0: ((v3 / ((int(v2)) | 1)) >= 128), F1: ""}
                        _ = v40
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = fi33
            }
        }
        var v41 uint8 = v2
        for fi42 := 0; fi42 < 14; fi42++ {
            v3 -= (fi42 | fi42)
            var v43 uint = 1
            _ = fi42
            _ = v43
        }
        _ = v41
    }
    v3 += (v3 >> (((v10 % ((255) | 1))) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
