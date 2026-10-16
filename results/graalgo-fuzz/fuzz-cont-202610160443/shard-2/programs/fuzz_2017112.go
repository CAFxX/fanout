package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 int16 }

func main() {
    var v1 bool = (1 == 265)
    var v2 int = 32768
    var v3 float64 = 1e10
    var v4 uint32 = 4294967295
    v5 := []int{230}
    switch "" {
    case "\n\t":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "foo bar":
        fmt.Println(v3, v2)
    default:
        fmt.Println(v3)
    }
    if true {
        if len(v5) > 0 {
            v5[0] = (v2 ^ v2)
        }
        for fi6 := 0; fi6 < 3; fi6++ {
            for fi7 := 0; fi7 < 10; fi7++ {
                if v1 {
                    v1 = v1
                    var v8 uint64 = 974
                    v8 += uint64(v4)
                    _ = v8
                } else {
                    v4 = v4
                }
                var v9 int = fi7
                _ = fi7
                _ = v9
            }
            _ = fi6
        }
        for fi10 := 0; fi10 < 10; fi10++ {
            if (!(("\n\t" + "") >= "\n\t")) {
                v11 := GenS1{F0: 1}
                for fi12 := 0; fi12 < 15; fi12++ {
                    v5 = append(v5, (^(fi10)))
                    _ = fi12
                }
                _ = v11
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v13 := GenS2{F0: 22, F1: 65535, F2: 1}
                _ = v13
            }
            if (!false) {
                v14 := GenS3{F0: 351}
                if len(v5) > 0 {
                    v5[0] = 0
                }
                _ = v14
            }
            _ = len(v5)
            _ = fi10
        }
    }
    _ = len(v5)
    v1 = true
    switch "\n\t" {
    case "foo bar":
        v1 = true
    case "hello":
        v15 := GenS2{F0: 50, F1: 65535, F2: -350}
        _ = v15
    default:
        var v16 uint32 = (^(v4))
        _ = v16
    }
    switch (-(v2)) {
    case 968:
        switch v2 {
        case 42:
            var v17 uint32 = (^(v4))
            _ = v17
        case 2147483647:
            for fi18 := 0; fi18 < 13; fi18++ {
                v2 += fi18
                var v19 int16 = -32768
                _ = fi18
                _ = v19
            }
        }
    case 0:
        v2 += (v2 - (v2 - 288))
    default:
        v1 = v1
    }
    var v20 int64 = -223
    for fi21 := 0; fi21 < 8; fi21++ {
        v20 = v20
        v20 -= int64(v2)
        _ = fi21
    }
    switch "" {
    case "a":
        var v22 int16 = 30
        _ = v22
    case "\n\t":
        v23 := GenS4{F0: (v4 % ((704) | 1)), F1: 42}
        _ = v23
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v24 int16 = 1000
    for fi25 := 0; fi25 < 20; fi25++ {
        v24 = v24
        v24 -= v24
        v26 := GenS1{F0: 1}
        _ = fi25
        _ = v26
    }
    switch v2 {
    case 42:
        v27 := GenS1{F0: -1}
        _ = v27
    case -2147483648:
        for fi28 := 0; fi28 < 16; fi28++ {
            for fi29 := 0; fi29 < 12; fi29++ {
                if (v1 && v1) {
                    v3 = v3
                    v30 := GenS4{F0: v4, F1: (v24 / (((v24 << ((255) & 63))) | 1))}
                    fmt.Println(v3, fi29, fi28)
                    _ = v30
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(fi29, v24)
                _ = fi29
            }
            var v31 uint64 = 671
            var v32 int8 = 1
            _ = fi28
            _ = v31
            _ = v32
        }
    default:
        if v1 {
            fmt.Println(v20, v24)
        }
    }
    switch "foo bar" {
    case "a":
        var v33 int64 = 9223372036854775807
        _ = v33
    case "":
        v20 = 0
    case "hello":
        switch int(v4) {
        case 9223372036854775807:
            _ = len(v5)
        }
    }
    var v34 int = v2
    v20 += (v20 >> ((2) & 63))
    _ = len(v5)
    if (!v1) {
        _ = len(v5)
        fmt.Println(v2, v2)
    } else {
        var v35 int16 = (v24 + -1)
        _ = len(v5)
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
