package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int16 }

func main() {
    var v1 int = -2147483648
    v2 := make(map[string]bool)
    var v3 int32 = -1
    var v4 int16 = -32768
    var v5 float32 = 1e10
    fmt.Println(v1)
    v3 = (^(v3))
    v3 = (v3 - v3)
    var v6 int64 = -9223372036854775808
    if (v5 <= (v5 * v5)) {
        var v7 uint64 = 2
        _ = v7
    } else {
        var v8 bool = true
        switch v1 {
        case 32767:
            if (!("" > ("" + "hello"))) {
                delete(v2, "")
                for fi9 := 0; fi9 < 12; fi9++ {
                    v10 := GenS1{F0: fi9}
                    _ = fi9
                    _ = v10
                }
                var v11 int8 = 42
                _ = v11
            }
        case -2147483648:
            switch (v1 | v1) {
            case 4294967295:
                var v12 bool = (v5 == v5)
                _ = v12
            case -9223372036854775808:
                if v8 {
                    var v13 uint32 = 892
                    var v14 string = "\n\t"
                    _ = v13
                    _ = v14
                }
            default:
                v3 -= v3
            }
        case 760:
            fmt.Println(v6, v6)
        default:
            for fi15 := 0; fi15 < 11; fi15++ {
                v16 := GenS1{F0: v1}
                var v17 uint32 = 2
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = v8
    }
    v18 := GenS1{F0: (v1 ^ 0)}
    v19 := GenS2{F0: 271, F1: -80}
    if (4294967295 < 2) {
        if (true && true) {
            v6 += 42
        } else {
            if (65535 < 350) {
                var v20 uint = 0
                var v21 int8 = 1
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 9; fi22++ {
                v1 = fi22
                _ = fi22
            }
        }
    } else {
        if false {
            v23 := []int{-2147483648, v1, v1, v1}
            for fi24 := 0; fi24 < 13; fi24++ {
                switch "" {
                case "a":
                    var v25 int16 = (v4 >> ((42) & 63))
                    _ = v25
                case "":
                    for fi26 := 0; fi26 < 5; fi26++ {
                        var v27 int8 = -1
                        _ = fi26
                        _ = v27
                    }
                case "hello":
                    var v28 uint32 = 65535
                    _ = v28
                }
                v1 = (^(v1))
                _ = fi24
            }
            for fi29 := 0; fi29 < 13; fi29++ {
                v6 += (v6 >> ((807) & 63))
                v5 = v5
                _ = fi29
            }
            _ = v23
        }
        switch (v1 * (v1 + v1)) {
        case 4294967295:
            for fi30 := 0; fi30 < 14; fi30++ {
                var v31 uint8 = 2
                _ = fi30
                _ = v31
            }
        }
    }
    v4 = -910
    switch (("foo bar" + "") + ("foo bar" + "")) {
    case "":
        for fi32 := 0; fi32 < 10; fi32++ {
            if true {
                v3 = (v3 ^ -2147483648)
            }
            for fi33 := 0; fi33 < 20; fi33++ {
                switch fi32 {
                case -9223372036854775808:
                    v5 -= -2.25
                case 2147483647:
                    var v34 uint32 = 65535
                    _ = v34
                case -25:
                    v4 -= (v4 + (v4 >> ((4294967295) & 63)))
                default:
                    v5 += v5
                }
                if ((0.0 <= 0.0) || ((v3 + (v3 & v3)) > (v3 | 42))) {
                    v1 -= 42
                }
                _ = fi33
            }
            _ = fi32
        }
    }
    v35 := GenS3{F0: (^(v6))}
    for fi36 := 0; fi36 < 15; fi36++ {
        v37 := GenS1{F0: v1}
        for fi38 := 0; fi38 < 13; fi38++ {
            if ("foo bar" == "a") {
                v3 += (v3 | (v3 + 0))
                if true {
                    v4 += (v4 << ((255) & 63))
                    v39 := GenS3{F0: v6}
                    switch (fi36 * (fi38 / ((fi36) | 1))) {
                    case -2147483648:
                        v4 = v4
                    case 4294967295:
                        v40 := GenS2{F0: 4294967295, F1: -128}
                        _ = v40
                    case 128:
                        var v41 uint = 983
                        _ = v41
                    default:
                        v42 := GenS4{F0: (^(v6)), F1: (v4 & v4)}
                        _ = v42
                    }
                    fmt.Println(fi38, v1, fi36)
                    _ = v39
                } else {
                    var v43 uint64 = 4294967295
                    var v44 string = "foo bar"
                    _ = v43
                    _ = v44
                }
                v4 += v4
            }
            _ = fi38
        }
        if ("\n\t" < "a") {
            fmt.Println(v4, v6)
        }
        _ = fi36
        _ = v37
    }
    var v45 int16 = v4
    switch (^(v1)) {
    case 9223372036854775807:
        for fi46 := 0; fi46 < 18; fi46++ {
            switch "foo bar" {
            case "a":
                v47 := v2["foo bar"]
                _ = v47
            case "\n\t":
                v45 -= v45
            default:
                switch (("\n\t" + (("" + ((("hello" + (((("\n\t" + "") + "a") + "foo bar") + "\n\t")) + "\n\t") + "")) + ("foo bar" + ("" + ("a" + "foo bar"))))) + (("\n\t" + "foo bar") + (("a" + "") + "\n\t"))) {
                case "a":
                    if true {
                        v48 := GenS4{F0: -1, F1: v45}
                        var v49 bool = true
                        _ = v48
                        _ = v49
                    } else {
                        var v50 float32 = v5
                        var v51 uint8 = 1
                        v51 += v51
                        _ = v50
                        _ = v51
                    }
                }
            }
            var v52 int64 = (v6 * (^(v6)))
            _ = fi46
            _ = v52
        }
    case 0:
        var v53 uint = 2
        _ = v53
    case 1:
        for fi54 := 0; fi54 < 13; fi54++ {
            var v55 int8 = 1
            _ = fi54
            _ = v55
        }
    }
    var v56 uint32 = 4294967295
    switch "foo bar" {
    case "hello":
        v57 := GenS1{F0: (v1 | v1)}
        _ = v57
    default:
        v4 -= v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v45
    _ = v5
    _ = v56
    _ = v6
}
