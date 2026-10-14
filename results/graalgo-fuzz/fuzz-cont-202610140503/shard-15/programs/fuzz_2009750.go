package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS3 struct { F0 int; F1 int16 }

func main() {
    var v1 uint16 = 255
    var v2 uint32 = 42
    var v3 float32 = 3.14159
    var v4 int32 = 1
    var v5 int32 = (v4 >> ((255) & 63))
    var v6 string = "foo bar"
    var v7 uint16 = (v1 & 42)
    if true {
        switch -930 {
        case -2147483648:
            v8 := []int{-2147483648, -616}
            _ = v8
        case 4294967295:
            fmt.Println(v1, v4, v6)
        default:
            if (v6 >= (("foo bar" + v6) + (("hello" + v6) + (("a" + v6) + (v6 + "a"))))) {
                fmt.Println(v2, v5, v2)
                switch "a" {
                case "":
                    var v9 int16 = -440
                    _ = v9
                case "a":
                    for fi10 := 0; fi10 < 9; fi10++ {
                        var v11 int64 = -185
                        var v12 int64 = v11
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                default:
                    for fi13 := 0; fi13 < 11; fi13++ {
                        var v14 int8 = -101
                        _ = fi13
                        _ = v14
                    }
                }
                fmt.Println(v4)
            }
        }
        var v15 uint8 = 42
        _ = v15
    }
    v16 := GenS1{F0: 197}
    v2 = v2
    for fi17 := 0; fi17 < 20; fi17++ {
        fmt.Println(v4, v3, v3)
        v2 += (v2 | (v2 * v2))
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 int8 = -83
            if (!true) {
                var v20 uint = 199
                var v21 int8 = (v19 / (((v19 << (((v20 ^ (v20 | (^(v20))))) & 63))) | 1))
                switch fi17 {
                case -879:
                    fmt.Println(v5, v20, v1)
                default:
                    var v22 string = v6
                    _ = v22
                }
                _ = v20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v3 -= (v3 * (v3 - v3))
    v23 := GenS2{F0: v1, F1: 18446744073709551615, F2: v7}
    if true {
        v7 = 1
        v6 = ((v6 + v6) + "a")
        v2 = v2
    } else {
        if (127 < 42) {
            var v24 int = 32767
            _ = v24
        }
        var v25 int32 = v4
        _ = v25
    }
    fmt.Println(v6)
    switch "foo bar" {
    case "":
        var v26 int16 = 939
        _ = v26
    }
    v2 -= v2
    switch "" {
    case "hello":
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 uint16 = (v1 >> ((65535) & 63))
            switch (("foo bar" + v6) + ((v6 + v6) + (v6 + v6))) {
            case "a":
                v29 := GenS1{F0: 134}
                _ = v29
            }
            var v30 int = fi27
            _ = fi27
            _ = v28
            _ = v30
        }
    case "foo bar":
        var v31 int16 = 32767
        _ = v31
    case "":
        switch v6 {
        case "hello":
            var v32 uint64 = 4294967295
            _ = v32
        case "foo bar":
            v5 -= (v4 >> ((65535) & 63))
        }
    default:
        v33 := []int{771, 464, 983, 128}
        _ = v33
    }
    switch "\n\t" {
    case "":
        v34 := []int{860}
        _ = v34
    case "hello":
        var v35 uint = 4294967295
        _ = v35
    case "\n\t":
        switch 32768 {
        case -433:
            v36 := GenS1{F0: 42}
            _ = v36
        case -9223372036854775808:
            for fi37 := 0; fi37 < 18; fi37++ {
                var v38 float64 = 1e10
                var v39 int8 = -128
                _ = fi37
                _ = v38
                _ = v39
            }
        case -1:
            var v40 string = (("" + "\n\t") + v6)
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 19; fi41++ {
        var v42 int = fi41
        v43 := GenS3{F0: (fi41 & (fi41 - (fi41 ^ fi41))), F1: 1}
        switch (v6 + ("" + ("a" + v6))) {
        case "a":
            var v44 float64 = (1.5 * -2.25)
            _ = v44
        case "\n\t":
            v45 := []int{-15, v42, fi41}
            _ = v45
        default:
            var v46 int8 = 87
            _ = v46
        }
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
