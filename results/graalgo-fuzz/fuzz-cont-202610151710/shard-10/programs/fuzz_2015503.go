package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16; F1 uint32; F2 string }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    var v3 bool = (true || true)
    if v3 {
        var v4 int16 = 685
        _ = v4
    } else {
        fmt.Println(v3)
        v3 = ("\n\t" >= ("a" + "\n\t"))
        v5 := []int{-924}
        _ = v5
    }
    fmt.Println(v3)
    switch "\n\t" {
    case "foo bar":
        v3 = (402 < 1000)
    case "hello":
        v1[32767] = 397
    case "\n\t":
        var v6 int64 = 42
        _ = v6
    default:
        switch 567 {
        case 377:
            v3 = (v3 || (3.14159 == (3.14159 / ((0.0 + (0.1 / 3.14159)) - 1e10))))
        case 128:
            fmt.Println(v3)
        case 4294967295:
            for fi7 := 0; fi7 < 5; fi7++ {
                v8 := []int{fi7, (fi7 >> ((18446744073709551615) & 63)), fi7}
                var v9 int64 = 813
                fmt.Println(v3, fi7)
                _ = fi7
                _ = v8
                _ = v9
            }
        default:
            var v10 int8 = -128
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint32 = 284
        v13 := GenS1{F0: true}
        switch ("hello" + "hello") {
        case "\n\t":
            v14 := []int{}
            _ = v14
        case "foo bar":
            if (724 >= 42) {
                v15 := GenS1{F0: (true || (v3 || (v3 && v3)))}
                v16 := GenS1{F0: (232 < 80)}
                _ = v15
                _ = v16
            }
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    v3 = false
    v3 = (v3 || v3)
    for fi17 := 0; fi17 < 20; fi17++ {
        if (874 >= (fi17 & (fi17 % (((fi17 * fi17)) | 1)))) {
            for fi18 := 0; fi18 < 7; fi18++ {
                v3 = v3
                v19 := GenS1{F0: false}
                v3 = ((v3 && (-1 != fi18)) && v3)
                var v20 uint16 = 2
                v21 := GenS2{F0: 255, F1: 18446744073709551615}
                v3 = v3
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 int16 = 645
            if v3 {
                var v24 uint32 = 0
                v3 = v3
                v24 = (v24 & (v24 & v24))
                v25 := GenS3{F0: v23, F1: 2147483648, F2: "\n\t"}
                v26 := GenS4{F0: -1}
                if true {
                    var v27 uint16 = 0
                    _ = v27
                } else {
                    v3 = v3
                }
                _ = v24
                _ = v25
                _ = v26
            } else {
                var v28 uint64 = 0
                _ = v28
            }
            switch "a" {
            case "\n\t":
                v3 = false
            case "hello":
                var v29 int16 = (v23 & v23)
                _ = v29
            default:
                var v30 bool = v3
                _ = v30
            }
            _ = fi22
            _ = v23
        }
        _ = fi17
    }
    fmt.Println(v3)
    fmt.Println(v3)
    switch "\n\t" {
    case "":
        for fi31 := 0; fi31 < 4; fi31++ {
            if v3 {
                fmt.Println(v3)
                v3 = (65535 < 0)
                if v3 {
                    var v32 uint64 = 4294967295
                    var v33 int32 = 1
                    _ = v32
                    _ = v33
                }
            }
            v34 := GenS2{F0: 18446744073709551615, F1: 480}
            for fi35 := 0; fi35 < 12; fi35++ {
                var v36 int32 = 1
                v37 := []int{-560, (fi35 * (fi35 ^ 127)), fi31, fi35}
                var v38 bool = false
                _ = fi35
                _ = v36
                _ = v37
                _ = v38
            }
            _ = fi31
            _ = v34
        }
    case "\n\t":
        var v39 uint8 = 1
        _ = v39
    case "hello":
        v40 := GenS4{F0: 42}
        _ = v40
    default:
        for fi41 := 0; fi41 < 5; fi41++ {
            _ = len(v1)
            _ = fi41
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
