package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }

func main() {
    var v1 float32 = -2.25
    var v2 uint64 = 2
    var v3 uint32 = 151
    for fi4 := 0; fi4 < 17; fi4++ {
        fmt.Println(fi4)
        v5 := []int{fi4, (fi4 << ((509) & 63))}
        v1 = -2.25
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 12; fi6++ {
        if false {
            v7 := GenS1{F0: "\n\t"}
            _ = v7
        }
        _ = fi6
    }
    v8 := GenS1{F0: ""}
    switch "\n\t" {
    case "\n\t":
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint16 = 660
            _ = fi9
            _ = v10
        }
    case "":
        v11 := GenS1{F0: ""}
        _ = v11
    case "a":
        if (-1 < -1) {
            v12 := GenS1{F0: ("a" + "\n\t")}
            var v13 uint16 = 255
            _ = v12
            _ = v13
        }
    default:
        v14 := []int{-81}
        _ = v14
    }
    var v15 uint16 = 926
    var v16 bool = false
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 uint16 = v15
            _ = fi18
            _ = v19
        }
        v15 = (v15 * v15)
        var v20 uint32 = (v3 % ((2147483648) | 1))
        _ = fi17
        _ = v20
    }
    v3 = v3
    fmt.Println(v2, v15)
    if v16 {
        v1 -= v1
        var v21 uint16 = (v15 >> ((4294967295) & 63))
        _ = v21
    }
    switch "\n\t" {
    case "foo bar":
        v22 := []int{-535, -341, 2147483647}
        _ = v22
    case "a":
        switch 42 {
        case -578:
            v23 := []int{0}
            _ = v23
        case 24:
            switch "\n\t" {
            case "a":
                v16 = v16
            default:
                v3 = (v3 << ((526) & 63))
            }
        default:
            v24 := GenS2{F0: 0}
            _ = v24
        }
    case "\n\t":
        if (726 >= 4294967295) {
            for fi25 := 0; fi25 < 19; fi25++ {
                switch (fi25 - 1) {
                case 1:
                    for fi26 := 0; fi26 < 17; fi26++ {
                        v27 := GenS1{F0: "hello"}
                        _ = fi26
                        _ = v27
                    }
                case -9223372036854775808:
                    if true {
                        if (!false) {
                            fmt.Println(v3, fi25)
                            v1 = float32(3.14159)
                        } else {
                            var v28 string = "a"
                            v3 = (v3 >> ((255) & 63))
                            var v29 int = fi25
                            _ = v28
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 8; fi30++ {
                            var v31 float64 = float64(v1)
                            if true {
                                var v32 uint32 = v3
                                v1 = 0.0
                                v16 = (v2 >= uint64(v15))
                                v1 = float32(((float64((float32(v31) + (v1 * (v1 + 1e10)))) / 3.14159) / float64((v1 * v1))))
                                _ = v32
                            }
                            v16 = v16
                            v33 := []int{(-(fi25)), (-(fi30)), (fi30 / (((fi25 & (fi30 ^ (fi30 / ((fi30) | 1))))) | 1))}
                            v34 := GenS2{F0: -184}
                            if false {
                                v35 := GenS1{F0: ("hello" + "\n\t")}
                                v36 := GenS2{F0: 9223372036854775807}
                                fmt.Println(v3)
                                _ = v35
                                _ = v36
                            } else {
                                if v16 {
                                    switch (fi30 - -403) {
                                    case -779:
                                        fmt.Println(v1, v16, fi30)
                                        v1 = v1
                                        var v37 string = "hello"
                                        var v38 uint8 = 2
                                        _ = v37
                                        _ = v38
                                    }
                                    if false {
                                        v39 := GenS1{F0: "\n\t"}
                                        _ = len(v33)
                                        var v40 uint = 65535
                                        _ = v39
                                        _ = v40
                                    }
                                    fmt.Println(v16, fi30, fi30)
                                } else {
                                    var v41 float64 = float64(((v1 / -2.25) - (-2.25 + (v1 - (v1 + v1)))))
                                    v33 = append(v33, fi30)
                                    _ = v41
                                }
                            }
                            _ = fi30
                            _ = v31
                            _ = v33
                            _ = v34
                        }
                        var v42 string = "foo bar"
                        _ = v42
                    } else {
                        var v43 int64 = 9223372036854775807
                        if v16 {
                            if (-73 >= 0) {
                                v44 := []int{-2147483648, 646}
                                for fi45 := 0; fi45 < 2; fi45++ {
                                    var v46 uint16 = (v15 + (v15 - (v15 + 42)))
                                    _ = len(v44)
                                    _ = fi45
                                    _ = v46
                                }
                                _ = v44
                            }
                        } else {
                            fmt.Println(v15, v16)
                        }
                        _ = v43
                    }
                default:
                    v3 = (v3 + (v3 - 2))
                }
                v47 := GenS2{F0: -1}
                _ = fi25
                _ = v47
            }
        } else {
            v48 := []int{}
            _ = v48
        }
    }
    if (!(!v16)) {
        v15 = (v15 / ((v15) | 1))
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
