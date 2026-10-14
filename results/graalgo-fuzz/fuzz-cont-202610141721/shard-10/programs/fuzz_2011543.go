package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0() uint {
    if true {
        if true {
            v5 := GenS1{F0: 749, F1: 42}
            var v6 string = ""
            _ = v5
            _ = v6
        }
        v7 := GenS1{F0: 4294967295, F1: 127}
        _ = v7
    } else {
        if true {
            var v8 uint16 = 65535
            v8 += (^(v8))
            if false {
                switch -9223372036854775808 {
                case -9223372036854775808:
                    switch 4294967295 {
                    case -1:
                        v9 := []int{-1, 543}
                        _ = v9
                    case 2147483647:
                        for fi10 := 0; fi10 < 15; fi10++ {
                            v11 := []int{fi10, (fi10 & (fi10 % (((fi10 % ((fi10) | 1))) | 1)))}
                            _ = fi10
                            _ = v11
                        }
                    default:
                        var v12 int32 = -1
                        _ = v12
                    }
                case 9223372036854775807:
                    v8 = (^(v8))
                case 2147483647:
                    var v13 uint64 = 42
                    _ = v13
                }
                fmt.Println(v8)
                for fi14 := 0; fi14 < 18; fi14++ {
                    v15 := []int{fi14, fi14, (fi14 / ((42) | 1)), fi14}
                    var v16 int32 = 2147483647
                    v16 = (-(v16))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            } else {
                if false {
                    v8 = v8
                    var v17 uint16 = (^(v8))
                    _ = v17
                }
                fmt.Println(v8)
                var v18 bool = true
                _ = v18
            }
            _ = v8
        }
        var v19 uint = 82
        v20 := GenS1{F0: 898, F1: 200}
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: 145, F1: 32}
    switch 1 {
    case 128:
        var v22 int8 = -1
        _ = v22
    default:
        for fi23 := 0; fi23 < 11; fi23++ {
            switch (fi23 << ((4294967295) & 63)) {
            case 1:
                for fi24 := 0; fi24 < 1; fi24++ {
                    var v25 bool = (!(-229 == 750))
                    _ = fi24
                    _ = v25
                }
            case 127:
                for fi26 := 0; fi26 < 17; fi26++ {
                    if true {
                        var v27 int64 = 0
                        _ = v27
                    }
                    _ = fi26
                }
            }
            var v28 string = "a"
            _ = fi23
            _ = v28
        }
    }
    _ = v21
    return 0
}
func genFn1() uint32 {
    switch 32767 {
    case 276:
        switch -1 {
        case 9223372036854775807:
            var v29 bool = (100 > -83)
            _ = v29
        case -1:
            if (1 > 52) {
                var v30 bool = (!(166 != 4294967295))
                _ = v30
            }
        }
    }
    var v31 uint8 = 2
    fmt.Println(v31)
    _ = v31
    return 4294967295
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]bool)
    var v3 uint8 = 5
    v4 := make(map[uint32]bool)
    if false {
        for fi32 := 0; fi32 < 8; fi32++ {
            fmt.Println(v1, v3)
            if false {
                var v33 float32 = (float32((1.5 - 0.0)) - (((1.5 + 0.1) - 0.1) * (-2.25 * (0.0 / float32((3.14159 * -2.25))))))
                _ = v33
            }
            v3 = (v3 - v3)
            _ = fi32
        }
        v34 := []int{4294967295, 4294967295, 9223372036854775807, -9223372036854775808}
        _ = v34
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        v1 = v1
        v2[4294967295] = true
        _ = fi35
    }
    v36 := []int{}
    fmt.Println(v1, v3)
    for fi37 := 0; fi37 < 1; fi37++ {
        if false {
            switch "foo bar" {
            case "foo bar":
                _ = len(v36)
            case "hello":
                v1 = -834
            case "":
                delete(v4, 1)
            }
        }
        _ = fi37
    }
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        if true {
            if true {
                v3 -= v3
                v1 -= (-(v1))
            } else {
                v3 -= (^(v3))
                if true {
                    switch fi38 {
                    case 178:
                        var v39 string = ""
                        _ = v39
                    case 1:
                        v40 := GenS2{F0: 0}
                        _ = v40
                    case 2147483647:
                        _ = len(v2)
                    default:
                        if len(v36) > 0 {
                            v36[0] = (fi38 >> ((2) & 63))
                        }
                    }
                    v1 = (v1 % ((v1) | 1))
                } else {
                    var v41 int = (^(fi38))
                    _ = v41
                }
                v3 = v3
            }
        }
        _ = fi38
    }
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
