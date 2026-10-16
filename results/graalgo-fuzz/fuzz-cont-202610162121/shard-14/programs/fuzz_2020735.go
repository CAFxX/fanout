package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint8) uint64 {
    p0 += 127
    v5 := []int{515, -710}
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 bool = false
        v7 = v7
        _ = fi6
        _ = v7
    }
    if (42 != -2147483648) {
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(fi8)
            _ = fi8
        }
        var v9 int16 = 42
        if false {
            if len(v5) > 0 {
                v5[0] = 9223372036854775807
            }
            var v10 string = "a"
            _ = v10
        } else {
            fmt.Println(v9)
            v11 := GenS1{F0: -554, F1: 42}
            _ = v11
        }
        _ = v9
    } else {
        for fi12 := 0; fi12 < 15; fi12++ {
            p0 += (^(p0))
            _ = fi12
        }
    }
    _ = p0
    _ = v5
    return 919
}
func genFn1(p0 int64) uint8 {
    if false {
        p0 += (^(p0))
        var v13 int = 517
        _ = v13
    } else {
        p0 += p0
        v14 := []int{42, 9223372036854775807}
        _ = v14
    }
    p0 = (p0 >> ((4294967295) & 63))
    v15 := []int{}
    switch "\n\t" {
    case "foo bar":
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int16 = 32767
            p0 -= p0
            _ = fi16
            _ = v17
        }
    case "a":
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    }
    _ = p0
    _ = v15
    return 2
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 uint8 = 42
    var v4 float32 = (float32(0.1) / 1.5)
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(v3)
        _ = fi18
    }
    var v19 uint32 = 0
    var v20 bool = (422 >= 127)
    v21 := GenS2{F0: -2147483648, F1: (^(v3))}
    v22 := GenS1{F0: 42, F1: 255}
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 uint64 = 42
        v25 := GenS2{F0: -2147483648, F1: 127}
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 uint16 = 370
    for fi27 := 0; fi27 < 17; fi27++ {
        switch "a" {
        case "foo bar":
            v20 = v20
        default:
            fmt.Println(v4, v20)
        }
        _ = fi27
    }
    if (("a" + (("foo bar" + ("" + ("hello" + ""))) + "foo bar")) != "foo bar") {
        v28 := []int{755}
        if (!v20) {
            v29 := GenS3{F0: 32767}
            var v30 bool = (-9223372036854775808 >= 127)
            _ = v29
            _ = v30
        } else {
            for fi31 := 0; fi31 < 16; fi31++ {
                for fi32 := 0; fi32 < 16; fi32++ {
                    _ = len(v1)
                    switch (fi32 << ((109) & 63)) {
                    case 32768:
                        if (-771 == -853) {
                            var v33 float32 = 3.14159
                            v19 = v19
                            _ = v33
                        } else {
                            switch "a" {
                            case "":
                                var v34 uint32 = 2147483648
                                _ = v34
                            case "hello":
                                fmt.Println(v19)
                            default:
                                var v35 int8 = 25
                                _ = v35
                            }
                            v36 := GenS3{F0: (v26 + (v26 & v26))}
                            v37 := v1["\n\t"]
                            _ = v36
                            _ = v37
                        }
                    default:
                        var v38 int8 = -128
                        _ = v38
                    }
                    _ = fi32
                }
                var v39 int64 = 198
                v3 -= genFn1(1)
                _ = fi31
                _ = v39
            }
            _ = len(v28)
            switch 814 {
            case 288:
                v19 = v19
            case 0:
                var v40 int64 = -795
                _ = v40
            }
        }
        _ = v28
    } else {
        if v20 {
            fmt.Println(v4, v26, v3)
        }
    }
    v2[(("hello" + ("a" + "a")) + "hello")] = (-1 > -786)
    v41 := GenS3{F0: v26}
    var v42 bool = v20
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
