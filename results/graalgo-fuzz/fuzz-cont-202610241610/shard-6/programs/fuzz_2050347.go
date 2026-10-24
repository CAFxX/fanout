package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 bool; F2 uint32 }
type GenS3 struct { F0 int32; F1 bool; F2 int32 }
func genFn0(p0 uint, p1 uint16) int {
    v5 := GenS1{F0: 0}
    p1 = 1
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(fi6)
            var v8 uint16 = 65535
            _ = fi7
            _ = v8
        }
        p0 += 4294967295
        var v9 int16 = 0
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 int = fi10
        p1 = p1
        fmt.Println(v11)
        _ = fi10
        _ = v11
    }
    switch 0 {
    case 0:
        v12 := GenS1{F0: 100}
        _ = v12
    case 9223372036854775807:
        p1 -= p1
    case 32767:
        v13 := []int{-2147483648, -9223372036854775808, 42, 752}
        _ = v13
    default:
        if false {
            v14 := GenS2{F0: 1, F1: false, F2: 753}
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    var v1 string = "a"
    var v2 int8 = 100
    v3 := []int{32767, 2147483647, 1}
    var v4 uint64 = 1
    if (!true) {
        v1 = "hello"
        if true {
            if len(v3) > 0 {
                v3[0] = genFn0(255, 65535)
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                fmt.Println(fi15, v4, v1)
                v1 = v1
                var v16 uint16 = 841
                _ = fi15
                _ = v16
            }
        } else {
            fmt.Println(v4, v1, v2)
            switch ((v1 + "") + "hello") {
            case "\n\t":
                fmt.Println(v4, v2)
            case "":
                v17 := GenS1{F0: 100}
                _ = v17
            }
            _ = len(v3)
        }
    } else {
        v18 := GenS2{F0: 2147483647, F1: ((1.5 != float64(-2.25)) && false), F2: 2}
        if true {
            for fi19 := 0; fi19 < 2; fi19++ {
                v4 = v4
                if (!false) {
                    if false {
                        var v20 int8 = v2
                        v3 = append(v3, fi19)
                        if len(v3) > 0 {
                            v3[0] = (fi19 >> ((204) & 63))
                        }
                        _ = v20
                    } else {
                        fmt.Println(v4, v2, v4)
                        var v21 uint = 65535
                        _ = v21
                    }
                } else {
                    v22 := GenS2{F0: fi19, F1: true, F2: 4294967295}
                    if (!(!(622 != 65535))) {
                        switch -885 {
                        case 42:
                            v23 := GenS2{F0: fi19, F1: true, F2: 42}
                            _ = v23
                        }
                    }
                    for fi24 := 0; fi24 < 8; fi24++ {
                        var v25 uint = 18446744073709551615
                        _ = fi24
                        _ = v25
                    }
                    _ = v22
                }
                v26 := GenS2{F0: fi19, F1: true, F2: 0}
                _ = fi19
                _ = v26
            }
        } else {
            v1 = v1
        }
        _ = v18
    }
    fmt.Println(v1)
    if len(v3) > 0 {
        v3[0] = -354
    }
    v27 := GenS2{F0: 0, F1: false, F2: 2}
    var v28 uint = 48
    fmt.Println(v2, v2, v4)
    v29 := GenS2{F0: 1, F1: ((v28 <= v28) && true), F2: 2147483647}
    switch 4294967295 {
    case 4294967295:
        v30 := GenS3{F0: 1, F1: true, F2: -265}
        _ = v30
    case 875:
        for fi31 := 0; fi31 < 7; fi31++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v32 uint64 = 1
            v2 += 42
            _ = fi31
            _ = v32
        }
    case 2147483647:
        fmt.Println(v1)
    default:
        v33 := GenS1{F0: v2}
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
