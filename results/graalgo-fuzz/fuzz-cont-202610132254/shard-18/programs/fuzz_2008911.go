package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 bool; F1 int64; F2 int32 }

func main() {
    var v1 float32 = -2.25
    var v2 int64 = -9223372036854775808
    var v3 string = "foo bar"
    var v4 bool = true
    v5 := make(map[string]int)
    if false {
        switch 42 {
        case 705:
            v6 := GenS1{F0: 65535}
            _ = v6
        case 1:
            v7 := GenS1{F0: 255}
            _ = v7
        default:
            _ = len(v5)
        }
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := GenS1{F0: 65535}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS2{F0: ("a" + "a"), F1: v2}
            _ = fi10
            _ = v11
        }
    } else {
        switch "foo bar" {
        case "hello":
            fmt.Println(v2, v3)
        case "":
            v12 := v5[""]
            _ = v12
        }
    }
    var v13 int32 = -997
    if false {
        v14 := GenS1{F0: 0}
        v1 = v1
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS1{F0: 577}
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        v17 := GenS1{F0: 912}
        _ = v17
    }
    switch (v3 + "") {
    case "foo bar":
        v18 := GenS1{F0: 2}
        _ = v18
    case "":
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 int = (fi19 << ((18446744073709551615) & 63))
            _ = fi19
            _ = v20
        }
    default:
        fmt.Println(v1, v3)
    }
    v21 := v5[""]
    switch (-(v21)) {
    case 9223372036854775807:
        for fi22 := 0; fi22 < 2; fi22++ {
            if (0 <= 1) {
                for fi23 := 0; fi23 < 3; fi23++ {
                    var v24 uint8 = 123
                    var v25 uint = 65535
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                if (42 < 0) {
                    if ((v21 << ((310) & 63)) >= fi22) {
                        var v26 int = (-(fi22))
                        v1 = (1e10 - v1)
                        _ = v26
                    }
                }
            }
            v3 = (v3 + "hello")
            switch 348 {
            case 0:
                v27 := []int{fi22, (^(v21))}
                _ = v27
            case 1:
                switch v3 {
                case "a":
                    v13 += -1
                case "foo bar":
                    var v28 int16 = -32768
                    _ = v28
                default:
                    v3 = ("\n\t" + ((v3 + "\n\t") + "\n\t"))
                }
            default:
                v2 = (v2 & v2)
            }
            _ = fi22
        }
    }
    var v29 uint = 0
    v30 := []int{(v21 / ((32768) | 1)), (v21 << (((v29 % (((^(v29))) | 1))) & 63)), 128, v21}
    for fi31 := 0; fi31 < 11; fi31++ {
        if v4 {
            v2 += (v2 + (v2 - (v2 | (v2 << (((v29 % ((v29) | 1))) & 63)))))
        }
        _ = fi31
    }
    v1 -= 1.5
    v29 = v29
    if v4 {
        if (!v4) {
            v32 := GenS3{F0: (902 > 255), F1: 1, F2: -1}
            fmt.Println(v21, v29, v4)
            var v33 uint64 = 1
            _ = v32
            _ = v33
        } else {
            if (!(65535 == 65535)) {
                _ = len(v5)
                if (1 >= 128) {
                    if (!false) {
                        var v34 int64 = v2
                        v1 -= 1e10
                        _ = v34
                    }
                    v2 = (v2 * v2)
                } else {
                    fmt.Println(v4)
                    if (!v4) {
                        var v35 int8 = -1
                        var v36 string = v3
                        _ = v35
                        _ = v36
                    }
                    v4 = v4
                }
                v21 = (^(v21))
            } else {
                fmt.Println(v3, v2, v29)
            }
            var v37 int32 = (v13 & v13)
            v13 -= -1
            _ = v37
        }
        v2 = (v2 >> ((v29) & 63))
        fmt.Println(v4, v21, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
