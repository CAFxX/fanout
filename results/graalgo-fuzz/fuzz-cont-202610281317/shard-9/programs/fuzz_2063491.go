package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 string; F2 int64 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 int64 = 1
    var v2 uint = 18446744073709551615
    v3 := make(map[int]uint32)
    v2 -= (v2 * v2)
    fmt.Println(v1)
    v2 = (v2 - v2)
    for fi4 := 0; fi4 < 20; fi4++ {
        v2 = (v2 >> ((255) & 63))
        _ = fi4
    }
    for fi5 := 0; fi5 < 10; fi5++ {
        v3[(fi5 - (fi5 / (((fi5 | fi5)) | 1)))] = 4294967295
        fmt.Println(fi5)
        v1 = v1
        _ = fi5
    }
    if false {
        v1 -= -1
        var v6 int32 = 474
        if ((!true) && true) {
            v6 = -1
            var v7 int32 = (v6 - v6)
            _ = v7
        }
        _ = v6
    }
    if false {
        v1 -= v1
    } else {
        if (98 > uint8(v2)) {
            var v8 bool = false
            _ = v8
        }
    }
    var v9 uint32 = 42
    if false {
        if (!true) {
            for fi10 := 0; fi10 < 5; fi10++ {
                _ = len(v3)
                var v11 int64 = (^(v1))
                _ = fi10
                _ = v11
            }
            if true {
                v12 := []int{}
                var v13 int = -1
                _ = len(v3)
                _ = v12
                _ = v13
            }
            switch 9223372036854775807 {
            case 42:
                v14 := GenS1{F0: 131, F1: 29, F2: (v2 - 428)}
                _ = v14
            }
        } else {
            var v15 uint16 = 631
            v1 -= 42
            _ = v15
        }
        v16 := GenS2{F0: (("foo bar" + ("\n\t" + "hello")) + "hello")}
        for fi17 := 0; fi17 < 12; fi17++ {
            v9 = 65535
            switch fi17 {
            case 994:
                var v18 float32 = 3.14159
                _ = v18
            case 4294967295:
                switch "hello" {
                case "":
                    fmt.Println(v2, fi17)
                default:
                    var v19 string = ""
                    _ = v19
                }
            default:
                switch (fi17 - (^(fi17))) {
                case -9223372036854775808:
                    v20 := GenS3{F0: v1}
                    _ = v20
                }
            }
            _ = fi17
        }
        _ = v16
    }
    var v21 uint16 = 32767
    v3[42] = uint32(v2)
    v1 = (v1 >> (((v2 ^ v2)) & 63))
    for fi22 := 0; fi22 < 12; fi22++ {
        if false {
            var v23 bool = false
            _ = v23
        } else {
            var v24 bool = (false || (!false))
            fmt.Println(v1)
            _ = v24
        }
        _ = fi22
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        if ((v9 * (^(v9))) != (v9 >> ((v2) & 63))) {
            v21 = (v21 % (((v21 ^ (v21 | (v21 & 2)))) | 1))
            v21 = v21
            switch "" {
            case "":
                if true {
                    v26 := GenS3{F0: 192}
                    if false {
                        var v27 string = "foo bar"
                        fmt.Println(v9, v1, fi25)
                        for fi28 := 0; fi28 < 1; fi28++ {
                            if true {
                                v29 := GenS4{F0: 42, F1: v27, F2: 9223372036854775807}
                                _ = v29
                            }
                            _ = fi28
                        }
                        _ = v27
                    } else {
                        v30 := []int{(fi25 + (fi25 + 128)), (fi25 ^ 2147483647), (^(fi25))}
                        _ = v30
                    }
                    switch (("foo bar" + "a") + "hello") {
                    case "foo bar":
                        v9 = 2
                    }
                    _ = v26
                } else {
                    v31 := GenS3{F0: v1}
                    _ = v31
                }
            default:
                if true {
                    switch "\n\t" {
                    case "a":
                        for fi32 := 0; fi32 < 9; fi32++ {
                            var v33 int = fi25
                            var v34 uint64 = 42
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    case "":
                        var v35 uint64 = 18446744073709551615
                        _ = v35
                    }
                }
            }
        } else {
            for fi36 := 0; fi36 < 4; fi36++ {
                if true {
                    var v37 int32 = 1
                    v37 += (v37 * (v37 * (v37 >> (((^(v2))) & 63))))
                    if false {
                        fmt.Println(v37, fi36)
                        fmt.Println(v9, v9, v37)
                        v1 -= (v1 & v1)
                        var v38 uint32 = (v9 % ((v9) | 1))
                        v39 := GenS2{F0: ""}
                        _ = v38
                        _ = v39
                    } else {
                        fmt.Println(v21, v1, v1)
                        switch "a" {
                        case "":
                            fmt.Println(v9, v2)
                            if (uint8(v21) == 127) {
                                v40 := GenS5{F0: (v37 % ((v37) | 1))}
                                v41 := []int{(fi36 ^ -833)}
                                v2 -= 2
                                _ = v40
                                _ = v41
                            }
                        }
                    }
                    _ = v37
                }
                var v42 float64 = 0.0
                _ = fi36
                _ = v42
            }
        }
        _ = fi25
    }
    v1 += -9223372036854775808
    v43 := GenS2{F0: "foo bar"}
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v43
    _ = v9
}
