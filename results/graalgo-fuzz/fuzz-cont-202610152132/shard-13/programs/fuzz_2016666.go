package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int16; F1 int16; F2 int32 }
type GenS5 struct { F0 uint8; F1 int32; F2 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 1
    var v3 uint8 = 128
    var v4 int = 1
    var v5 uint16 = 977
    for fi6 := 0; fi6 < 11; fi6++ {
        v4 = fi6
        var v7 uint16 = v5
        v8 := GenS1{F0: 950, F1: 65535, F2: 4294967295}
        _ = fi6
        _ = v7
        _ = v8
    }
    switch "\n\t" {
    case "\n\t":
        switch v4 {
        case 42:
            v2 = -9223372036854775808
        case -2147483648:
            if (v2 > (^(v2))) {
                if true {
                    switch "\n\t" {
                    case "hello":
                        v9 := GenS1{F0: 42, F1: 1, F2: 79}
                        _ = v9
                    }
                    var v10 int = v4
                    v2 = v2
                    _ = v10
                }
            } else {
                if (("\n\t" > ("" + "foo bar")) && (false && true)) {
                    v11 := []int{v4}
                    if (false && (((float32(1.5) >= (3.14159 / (0.1 / (0.0 + 3.14159)))) || true) && (!false))) {
                        if (42 != 17) {
                            _ = len(v1)
                            if len(v11) > 0 {
                                v11[0] = v4
                            }
                        }
                        v12 := GenS2{F0: -1, F1: 42}
                        _ = v12
                    }
                    _ = v11
                }
                v13 := []int{(v4 ^ (v4 | v4))}
                _ = v13
            }
        }
    case "hello":
        var v14 string = "hello"
        _ = v14
    case "foo bar":
        v15 := []int{32767}
        _ = v15
    default:
        v16 := GenS2{F0: 42, F1: 1}
        _ = v16
    }
    var v17 int64 = 42
    for fi18 := 0; fi18 < 2; fi18++ {
        v1[637] = 0
        if (!((("hello" + "foo bar") + "foo bar") != "")) {
            v17 = -374
            v5 -= v5
        }
        if true {
            for fi19 := 0; fi19 < 15; fi19++ {
                v3 = 128
                fmt.Println(fi19, fi18, v3)
                _ = fi19
            }
            switch (v4 << ((1) & 63)) {
            case 478:
                fmt.Println(v2)
            case 42:
                var v20 uint = 2
                _ = v20
            }
        }
        _ = fi18
    }
    v21 := GenS2{F0: 169, F1: 429}
    switch "foo bar" {
    case "hello":
        fmt.Println(v2, v4)
    case "":
        for fi22 := 0; fi22 < 19; fi22++ {
            v4 += (-(fi22))
            v23 := GenS1{F0: -32768, F1: 2147483648, F2: 42}
            _ = fi22
            _ = v23
        }
    }
    v24 := GenS2{F0: -196, F1: 0}
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := GenS3{F0: 1, F1: 250}
        var v27 float64 = ((-2.25 - 3.14159) - 1e10)
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := []int{v4, v4, (v4 & (v4 + v4))}
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    v5 -= (^(v5))
    v3 = 219
    if true {
        delete(v1, 42)
    }
    switch v4 {
    case 9223372036854775807:
        v28 = append(v28, v4)
    case 309:
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 int8 = -119
            for fi31 := 0; fi31 < 6; fi31++ {
                v32 := GenS4{F0: 1000, F1: int16(v2), F2: 0}
                _ = fi31
                _ = v32
            }
            for fi33 := 0; fi33 < 6; fi33++ {
                if true {
                    var v34 int64 = v17
                    fmt.Println(fi33, v17, v3)
                    _ = v34
                }
                v17 = -1
                v35 := GenS2{F0: -329, F1: 0}
                _ = fi33
                _ = v35
            }
            _ = fi29
            _ = v30
        }
    default:
        for fi36 := 0; fi36 < 18; fi36++ {
            var v37 bool = false
            v38 := GenS3{F0: 0, F1: v17}
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    v28 = append(v28, v4)
    if (true || ("a" == ("hello" + "\n\t"))) {
        var v39 string = (("a" + "a") + "a")
        v2 = (^(v2))
        _ = v39
    } else {
        v40 := GenS1{F0: -1, F1: 0, F2: 42}
        _ = v40
    }
    v41 := GenS5{F0: v3, F1: 42, F2: (true || ((v5 != (^(v5))) && false))}
    fmt.Println(v3, v2, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
