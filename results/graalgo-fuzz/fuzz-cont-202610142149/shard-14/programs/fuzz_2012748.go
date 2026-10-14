package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint64) uint32 {
    var v5 int8 = (-(p0))
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 uint8 = 1
        if ((-(v5)) < 42) {
            v8 := []int{fi6, -1}
            var v9 uint32 = 65535
            if false {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 int64 = 162
                    _ = fi10
                    _ = v11
                }
                v12 := GenS1{F0: 255}
                _ = v12
            } else {
                switch 1 {
                case 544:
                    fmt.Println(v5, v7, v9)
                case 32768:
                    v13 := GenS1{F0: 4294967295}
                    _ = v13
                default:
                    v14 := GenS1{F0: 2}
                    _ = v14
                }
            }
            _ = v8
            _ = v9
        }
        v15 := GenS1{F0: uint(v7)}
        _ = fi6
        _ = v15
        _ = v7
    }
    if false {
        p1 = 18446744073709551615
    }
    _ = p0
    _ = p1
    _ = v5
    return 2147483648
}

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 2
    var v3 uint32 = v2
    var v4 int32 = -439
    var v16 string = ("" + "\n\t")
    fmt.Println(v3, v4)
    fmt.Println(v1, v1, v4)
    var v17 uint = 729
    v18 := []int{498, -999}
    for fi19 := 0; fi19 < 18; fi19++ {
        v17 -= (^(v17))
        switch (v16 + v16) {
        case "\n\t":
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        case "":
            var v20 string = v16
            _ = v20
        }
        _ = fi19
    }
    v1 -= (v1 << (((^(v17))) & 63))
    switch v16 {
    case "\n\t":
        v21 := GenS1{F0: 488}
        _ = v21
    default:
        fmt.Println(v17, v17)
    }
    if true {
        v2 = 642
        if ((v16 + v16) == ("\n\t" + ((v16 + ("a" + ("hello" + v16))) + v16))) {
            fmt.Println(v1)
            var v22 string = "foo bar"
            _ = v22
        } else {
            v23 := GenS2{F0: uint(v2), F1: 127, F2: -430}
            if false {
                var v24 uint = v17
                _ = v24
            }
            _ = v23
        }
        switch v16 {
        case "foo bar":
            if true {
                if false {
                    v4 = v4
                } else {
                    for fi25 := 0; fi25 < 2; fi25++ {
                        var v26 uint16 = 42
                        v17 += v17
                        if (!false) {
                            fmt.Println(v17, v16)
                            var v27 bool = ((-(fi25)) == 32767)
                            _ = v27
                        } else {
                            var v28 uint8 = 78
                            var v29 uint8 = (v28 | (v28 | v28))
                            _ = v28
                            _ = v29
                        }
                        _ = fi25
                        _ = v26
                    }
                    v30 := GenS1{F0: (v17 - (v17 / (((v17 ^ v17)) | 1)))}
                    fmt.Println(v4, v17)
                    _ = v30
                }
                for fi31 := 0; fi31 < 6; fi31++ {
                    if false {
                        v3 = (v2 ^ (v3 + v2))
                        for fi32 := 0; fi32 < 17; fi32++ {
                            if false {
                                if (!true) {
                                    switch "foo bar" {
                                    case "hello":
                                        var v33 uint32 = 4294967295
                                        _ = v33
                                    case "foo bar":
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                    default:
                                        fmt.Println(v16, v17, v3)
                                    }
                                    var v34 bool = true
                                    _ = v34
                                }
                                fmt.Println(fi32, v17, v16)
                            }
                            switch v16 {
                            case "\n\t":
                                switch v16 {
                                case "":
                                    v16 = ("hello" + v16)
                                case "foo bar":
                                    v35 := GenS3{F0: uint64(v17), F1: 1, F2: 32767}
                                    _ = v35
                                default:
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                }
                            case "":
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                            }
                            _ = fi32
                        }
                    } else {
                        v36 := GenS4{F0: 1}
                        _ = v36
                    }
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    _ = fi31
                }
            } else {
                v3 -= (v2 ^ v2)
                var v37 string = "hello"
                _ = v37
            }
        default:
            var v38 uint32 = genFn0(0, 274)
            _ = v38
        }
    } else {
        var v39 uint = v17
        _ = v39
    }
    if false {
        for fi40 := 0; fi40 < 13; fi40++ {
            var v41 uint64 = (v1 % (((v1 * 4294967295)) | 1))
            _ = fi40
            _ = v41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
