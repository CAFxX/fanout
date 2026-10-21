package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 uint; F2 uint16 }
func genFn0(p0 uint16) int16 {
    v6 := GenS1{F0: 597, F1: (!true), F2: 982}
    v7 := GenS2{F0: 0}
    _ = p0
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint, p1 bool, p2 int64) uint64 {
    if p1 {
        v8 := GenS1{F0: 255, F1: (2 > 59), F2: 1}
        _ = v8
    } else {
        v9 := []int{}
        if len(v9) > 0 {
            v9[0] = 4294967295
        }
        _ = v9
    }
    p0 -= 2
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := []int{32768, 32768}
    var v2 uint64 = 1
    var v3 uint16 = 413
    var v4 uint8 = 255
    var v5 uint16 = (v3 - (^(v3)))
    v10 := GenS3{F0: -1, F1: -30}
    for fi11 := 0; fi11 < 18; fi11++ {
        if (true && true) {
            var v12 uint64 = (v2 & v2)
            v1 = append(v1, fi11)
            for fi13 := 0; fi13 < 9; fi13++ {
                v5 = (v5 + 42)
                fmt.Println(v2, v12)
                var v14 bool = false
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        var v15 int64 = 42
        _ = fi11
        _ = v15
    }
    var v16 uint16 = 42
    var v17 uint64 = 42
    if true {
        v18 := GenS1{F0: 4294967295, F1: (((("foo bar" + "hello") + ("\n\t" + ("foo bar" + ("a" + "hello")))) + "foo bar") > "foo bar"), F2: 481}
        _ = v18
    } else {
        v5 = (v3 / ((255) | 1))
        switch 2147483647 {
        case 0:
            for fi19 := 0; fi19 < 5; fi19++ {
                v17 -= (v17 << ((0) & 63))
                _ = fi19
            }
        case 32767:
            if false {
                v20 := GenS2{F0: 2147483648}
                for fi21 := 0; fi21 < 11; fi21++ {
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v17 = v17
                        var v23 int8 = -1
                        var v24 int32 = 42
                        var v25 uint64 = (^(v17))
                        if true {
                            fmt.Println(v5, v24, v17)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            v16 = v5
                            fmt.Println(v17)
                            switch "hello" {
                            case "foo bar":
                                if len(v1) > 0 {
                                    v1[0] = int(v24)
                                }
                            default:
                                switch "a" {
                                case "hello":
                                    v26 := GenS4{F0: true, F1: 42}
                                    _ = v26
                                }
                            }
                            v17 = v2
                        }
                        var v27 int16 = int16(v23)
                        var v28 float64 = 3.14159
                        if false {
                            var v29 uint8 = (v4 << ((65535) & 63))
                            _ = len(v1)
                            switch "foo bar" {
                            case "hello":
                                v17 = (v17 & v17)
                            }
                            var v30 int = (fi21 & 128)
                            v31 := GenS1{F0: 4294967295, F1: (42 > v4), F2: uint32(v25)}
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                        _ = v27
                        _ = v28
                    }
                    _ = fi21
                }
                _ = v20
            }
        default:
            fmt.Println(v17)
        }
    }
    fmt.Println(v16, v5)
    v2 = (^(v17))
    if false {
        for fi32 := 0; fi32 < 18; fi32++ {
            v1 = append(v1, 447)
            _ = len(v1)
            for fi33 := 0; fi33 < 1; fi33++ {
                fmt.Println(v16)
                v34 := GenS5{F0: fi33}
                switch (fi33 << ((18446744073709551615) & 63)) {
                case 2147483647:
                    v35 := GenS6{F0: "foo bar", F1: 512, F2: v16}
                    _ = v35
                }
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        v1 = append(v1, (fi36 << ((4294967295) & 63)))
        _ = len(v1)
        _ = fi36
    }
    if false {
        if false {
            switch "" {
            case "foo bar":
                fmt.Println(v4, v16)
            }
        }
        v5 = uint16(v2)
        if (!false) {
            var v37 uint16 = 65535
            _ = v37
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
