package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 bool; F1 int16 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 string = "hello"
    for fi5 := 0; fi5 < 5; fi5++ {
        switch (fi5 << ((42) & 63)) {
        case 32767:
            if true {
                v6 := []int{fi5, (-(fi5)), (fi5 | (fi5 >> ((v1) & 63)))}
                var v7 int32 = 808
                var v8 uint8 = 107
                _ = v6
                _ = v7
                _ = v8
            }
        case 128:
            v9 := GenS1{F0: 562, F1: 0}
            _ = v9
        case 9223372036854775807:
            fmt.Println(fi5)
        }
        switch 42 {
        case 4294967295:
            switch v4 {
            case "a":
                v10 := GenS2{F0: 0, F1: 0}
                _ = v10
            default:
                v11 := []int{(^(fi5)), fi5, fi5, fi5}
                _ = v11
            }
        case 2147483647:
            v1 += 4294967295
        case 843:
            var v12 bool = (fi5 > fi5)
            _ = v12
        }
        _ = fi5
    }
    var v13 int32 = -2147483648
    v1 = 0
    fmt.Println(v4, v4, v13)
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS2{F0: (fi14 & (fi14 + -767)), F1: 399}
        _ = fi14
        _ = v15
    }
    var v16 bool = (v1 > (v1 | 18446744073709551615))
    var v17 int32 = v13
    if ((!(69 < 42)) || ((!((!(2 != 2)) && true)) && ((!true) || (v13 <= v13)))) {
        v18 := GenS2{F0: -2147483648, F1: 2147483648}
        var v19 uint32 = 1
        _ = v18
        _ = v19
    } else {
        fmt.Println(v13, v17)
        v4 = (v4 + (v4 + "\n\t"))
    }
    v20 := GenS2{F0: -161, F1: 2}
    switch 1 {
    case 32767:
        v21 := GenS1{F0: 0, F1: 2}
        _ = v21
    case 128:
        var v22 float64 = (3.14159 / 1.5)
        _ = v22
    case -9223372036854775808:
        if v16 {
            var v23 uint8 = 177
            var v24 uint16 = 255
            fmt.Println(v17)
            _ = v23
            _ = v24
        }
    }
    fmt.Println(v17, v16)
    for fi25 := 0; fi25 < 1; fi25++ {
        if (false || v16) {
            v26 := GenS3{F0: ((fi25 >> ((v1) & 63)) == fi25), F1: -32768}
            if v16 {
                fmt.Println(v16)
                if true {
                    var v27 uint64 = 42
                    v28 := []int{fi25}
                    _ = v27
                    _ = v28
                } else {
                    v29 := GenS2{F0: fi25, F1: 2147483647}
                    v4 = "\n\t"
                    _ = v29
                }
                for fi30 := 0; fi30 < 8; fi30++ {
                    v16 = (((1e10 - (float64(1e10) - 3.14159)) == 1.5) && true)
                    _ = fi30
                }
            }
            var v31 uint32 = 2147483647
            _ = v26
            _ = v31
        } else {
            var v32 uint64 = 42
            if (42 != -1) {
                for fi33 := 0; fi33 < 2; fi33++ {
                    var v34 uint8 = 205
                    var v35 uint = v1
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                var v36 int32 = (v13 % ((-2147483648) | 1))
                _ = v36
            } else {
                var v37 int32 = v13
                _ = v37
            }
            var v38 int = 483
            _ = v32
            _ = v38
        }
        v4 = "hello"
        for fi39 := 0; fi39 < 9; fi39++ {
            var v40 int16 = 1
            _ = fi39
            _ = v40
        }
        _ = fi25
    }
    fmt.Println(v1)
    fmt.Println(v1, v4, v16)
    if true {
        v4 = v4
    } else {
        v41 := GenS2{F0: 9223372036854775807, F1: 65535}
        _ = v41
    }
    v42 := GenS1{F0: 255, F1: 127}
    var v43 int32 = (v13 / (((^(v13))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}
