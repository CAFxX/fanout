package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int32 }
func genFn0() int {
    var v6 uint16 = 1
    var v7 string = ("\n\t" + "foo bar")
    var v8 int = 128
    _ = v6
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 int32 = 42
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 4294967295
    var v4 int = -242
    var v5 int16 = -32768
    var v9 int64 = v2
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := []int{(v4 ^ v4)}
        _ = fi10
        _ = v11
    }
    switch v4 {
    case 2147483647:
        v12 := GenS1{F0: 0, F1: -128}
        _ = v12
    case -1:
        v9 -= (v9 / (((v2 ^ 797)) | 1))
    case 4294967295:
        switch -344 {
        case 32767:
            v13 := []int{-791, v4}
            _ = v13
        }
    default:
        v14 := GenS1{F0: 127, F1: 8}
        _ = v14
    }
    v2 = 0
    v9 -= v9
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            if (1e10 < 1.5) {
                var v17 bool = (v5 < 156)
                fmt.Println(v3)
                switch "foo bar" {
                case "":
                    v18 := GenS1{F0: fi15, F1: -1}
                    _ = v18
                default:
                    v3 = (^(v3))
                }
                _ = v17
            }
            _ = fi16
        }
        _ = fi15
    }
    switch 442 {
    case -1:
        for fi19 := 0; fi19 < 7; fi19++ {
            if false {
                if false {
                    for fi20 := 0; fi20 < 9; fi20++ {
                        v5 = int16(v1)
                        v5 = (v5 << ((42) & 63))
                        _ = fi20
                    }
                    var v21 int32 = (v1 + 429)
                    _ = v21
                }
            }
            var v22 uint64 = 42
            _ = fi19
            _ = v22
        }
    case 32768:
        if true {
            if false {
                if true {
                    v23 := GenS2{F0: "a"}
                    _ = v23
                }
                var v24 float32 = 1.5
                for fi25 := 0; fi25 < 4; fi25++ {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        switch fi25 {
                        case 0:
                            v27 := []int{v4, (fi26 + (fi26 & (fi26 - v4))), -9223372036854775808, 189}
                            _ = v27
                        }
                        fmt.Println(v1)
                        _ = fi26
                    }
                    if (v5 <= v5) {
                        var v28 float64 = -2.25
                        _ = v28
                    }
                    _ = fi25
                }
                _ = v24
            } else {
                switch genFn0() {
                case 9223372036854775807:
                    var v29 uint16 = 1
                    _ = v29
                case -148:
                    switch 128 {
                    case -275:
                        v30 := GenS3{F0: "foo bar", F1: v1}
                        _ = v30
                    default:
                        if ((v1 % ((315) | 1)) == (v1 ^ 1)) {
                            if (0 == 2) {
                                var v31 uint32 = 747
                                _ = v31
                            }
                            switch v4 {
                            case 128:
                                var v32 uint32 = 1
                                _ = v32
                            case 0:
                                var v33 uint8 = 2
                                _ = v33
                            }
                        } else {
                            var v34 uint16 = 0
                            if true {
                                if ((153 <= 7) || (!true)) {
                                    v35 := GenS1{F0: int(v3), F1: -128}
                                    var v36 string = "hello"
                                    _ = v35
                                    _ = v36
                                }
                                v1 += 1
                            } else {
                                var v37 int32 = int32(v4)
                                _ = v37
                            }
                            v38 := GenS3{F0: "a", F1: v1}
                            _ = v34
                            _ = v38
                        }
                    }
                default:
                    switch (^(v4)) {
                    case 278:
                        v4 = v4
                    }
                }
                v39 := []int{v4, (v4 >> ((53) & 63)), genFn0(), -443}
                _ = v39
            }
            var v40 int = genFn0()
            v9 += int64(v3)
            _ = v40
        }
    case 2147483647:
        v41 := []int{(v4 + v4), -2147483648}
        _ = v41
    default:
        fmt.Println(v4, v1, v3)
    }
    var v42 uint32 = 162
    for fi43 := 0; fi43 < 13; fi43++ {
        fmt.Println(v4)
        v3 = (v3 | (v3 + v3))
        if (!false) {
            for fi44 := 0; fi44 < 10; fi44++ {
                var v45 float32 = 0.0
                _ = fi44
                _ = v45
            }
            fmt.Println(v5, v4, v4)
            if true {
                for fi46 := 0; fi46 < 8; fi46++ {
                    v47 := []int{fi46, (fi43 >> ((255) & 63))}
                    var v48 uint = 18446744073709551615
                    v49 := GenS2{F0: "\n\t"}
                    _ = len(v47)
                    v5 += v5
                    var v50 string = ("\n\t" + "")
                    _ = fi46
                    _ = v47
                    _ = v48
                    _ = v49
                    _ = v50
                }
                v51 := []int{genFn0()}
                var v52 int32 = 0
                _ = v51
                _ = v52
            }
        }
        _ = fi43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v9
}
