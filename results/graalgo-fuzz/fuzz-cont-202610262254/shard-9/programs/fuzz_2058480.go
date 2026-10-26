package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16 }
func genFn0() int64 {
    v4 := GenS1{F0: 2, F1: 100}
    var v5 uint32 = 65535
    _ = v4
    _ = v5
    return -180
}

func main() {
    v1 := []int{-9223372036854775808}
    var v2 bool = (false && false)
    var v3 int16 = 1000
    var v6 uint8 = 109
    for fi7 := 0; fi7 < 6; fi7++ {
        switch 976 {
        case 63:
            v2 = ((!false) && false)
        case 9223372036854775807:
            if true {
                if v2 {
                    switch -458 {
                    case 32768:
                        fmt.Println(fi7, v3)
                    case 0:
                        if v2 {
                            v6 -= (v6 | v6)
                            for fi8 := 0; fi8 < 3; fi8++ {
                                fmt.Println(v6, fi8)
                                v9 := GenS2{F0: -2147483648}
                                _ = fi8
                                _ = v9
                            }
                            if v2 {
                                v10 := GenS3{F0: 32767, F1: 4294967295, F2: -9223372036854775808}
                                for fi11 := 0; fi11 < 9; fi11++ {
                                    if v2 {
                                        var v12 bool = v2
                                        v13 := GenS4{F0: v6}
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v12
                                        _ = v13
                                    } else {
                                        if v2 {
                                            var v14 int16 = v3
                                            v1 = append(v1, (fi7 ^ 32767))
                                            var v15 uint64 = 0
                                            if true {
                                                v16 := GenS1{F0: (v15 * 205), F1: (v6 ^ (v6 >> ((854) & 63)))}
                                                fmt.Println(v15, v3)
                                                _ = v16
                                            }
                                            var v17 bool = false
                                            _ = v14
                                            _ = v15
                                            _ = v17
                                        } else {
                                            switch int(v3) {
                                            case -2147483648:
                                                if len(v1) > 0 {
                                                    v1[0] = (fi7 + -307)
                                                }
                                            case 127:
                                                v2 = v2
                                            case 42:
                                                if false {
                                                    fmt.Println(v6, fi11)
                                                }
                                            }
                                            v3 = (v3 - (v3 << ((677) & 63)))
                                        }
                                        var v18 string = ("foo bar" + ("" + "\n\t"))
                                        _ = v18
                                    }
                                    _ = fi11
                                }
                                _ = len(v1)
                                _ = v10
                            }
                        } else {
                            v19 := GenS3{F0: 2, F1: 500, F2: -515}
                            var v20 bool = v2
                            _ = v19
                            _ = v20
                        }
                    default:
                        v6 = v6
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = fi7
                    }
                    v21 := GenS5{F0: (v3 + (v3 & v3))}
                    _ = v21
                }
                v22 := GenS5{F0: (v3 >> ((65535) & 63))}
                _ = v22
            } else {
                v1 = append(v1, 32767)
                v23 := GenS1{F0: 4294967295, F1: v6}
                if v2 {
                    var v24 int = (fi7 >> ((42) & 63))
                    _ = v24
                } else {
                    v3 -= (v3 ^ (^(v3)))
                }
                _ = v23
            }
        }
        v3 = v3
        _ = fi7
    }
    var v25 int = 128
    var v26 bool = v2
    v25 = (v25 >> ((1) & 63))
    for fi27 := 0; fi27 < 4; fi27++ {
        v25 = 4294967295
        _ = fi27
    }
    var v28 bool = (65535 == 0)
    fmt.Println(v28, v26)
    var v29 uint64 = 42
    var v30 uint = 65535
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    fmt.Println(v25, v3, v28)
    var v31 uint8 = 2
    var v32 uint = uint(v6)
    switch "a" {
    case "\n\t":
        v33 := GenS5{F0: (v3 - v3)}
        _ = v33
    case "":
        fmt.Println(v26)
    default:
        v34 := GenS5{F0: v3}
        _ = v34
    }
    v28 = v28
    switch v25 {
    case 9223372036854775807:
        fmt.Println(v26)
    case 127:
        fmt.Println(v25, v3)
    default:
        v1 = append(v1, v25)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v6
}
