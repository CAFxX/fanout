package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int16 }

func main() {
    v1 := []uint32{65535}
    v2 := []uint32{2, 710}
    var v3 uint32 = 692
    v4 := GenS1{F0: 4294967295, F1: 255, F2: -881}
    var v5 uint32 = v3
    v6 := GenS1{F0: 2, F1: 65535, F2: -1}
    if ((311 != 1) && false) {
        for fi7 := 0; fi7 < 10; fi7++ {
            fmt.Println(v3, v5, fi7)
            if false {
                var v8 int64 = 42
                v3 += v3
                _ = v8
            }
            fmt.Println(v5, v5)
            _ = fi7
        }
        fmt.Println(v3)
    }
    v5 = v5
    if (-128 <= 100) {
        v3 -= v5
        for fi9 := 0; fi9 < 9; fi9++ {
            switch "foo bar" {
            case "hello":
                switch (^(fi9)) {
                case 127:
                    var v10 float32 = 3.14159
                    _ = v10
                case -2147483648:
                    var v11 int = (fi9 + fi9)
                    _ = v11
                case -560:
                    switch (fi9 * fi9) {
                    case -2147483648:
                        for fi12 := 0; fi12 < 18; fi12++ {
                            v13 := GenS1{F0: 4294967295, F1: 409, F2: -707}
                            v5 = (v3 & (^(v3)))
                            v3 = v5
                            _ = fi12
                            _ = v13
                        }
                    case -1:
                        for fi14 := 0; fi14 < 8; fi14++ {
                            var v15 int8 = -1
                            fmt.Println(v15, v15)
                            _ = fi14
                            _ = v15
                        }
                    }
                }
            }
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 uint = 255
                switch (fi9 / ((fi9) | 1)) {
                case 1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi16
                _ = v17
            }
            if false {
                var v18 uint16 = 255
                v19 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: 1}
                _ = v18
                _ = v19
            } else {
                var v20 float64 = (0.1 * 0.0)
                var v21 int16 = 42
                v5 -= (v5 ^ v3)
                _ = v20
                _ = v21
            }
            _ = fi9
        }
        switch 127 {
        case -9223372036854775808:
            fmt.Println(v5, v5)
        }
    }
    if true {
        v3 = (v3 >> ((42) & 63))
        v1 = append(v1, 65535)
        switch -2147483648 {
        case -546:
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 uint16 = 0
                for fi24 := 0; fi24 < 16; fi24++ {
                    v23 += 183
                    _ = fi24
                }
                var v25 int = fi22
                _ = fi22
                _ = v23
                _ = v25
            }
        case 4294967295:
            for fi26 := 0; fi26 < 17; fi26++ {
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 string = "a"
                    switch fi26 {
                    case 1:
                        v5 = v3
                    }
                    if ((-2.25 >= 0.0) || (((v28 + (("\n\t" + v28) + "foo bar")) + v28) != "hello")) {
                        if true {
                            var v29 int32 = 278
                            var v30 int32 = v29
                            var v31 uint64 = 18446744073709551615
                            v29 += v29
                            var v32 int32 = (v29 - (^(v29)))
                            _ = v29
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    } else {
                        if (42 == 0) {
                            switch (fi26 + -779) {
                            case 9223372036854775807:
                                if len(v1) > 0 {
                                    v1[0] = v3
                                }
                            }
                        } else {
                            switch (fi26 - (^(fi26))) {
                            case 127:
                                _ = len(v2)
                            case -830:
                                v5 += v5
                            case -625:
                                v3 -= (v5 / ((v3) | 1))
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v5 = v3
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
        case 128:
            if (("a" == "") && (42 <= -552)) {
                v33 := GenS1{F0: 42, F1: 255, F2: -32768}
                var v34 uint16 = 248
                _ = v33
                _ = v34
            }
        }
    } else {
        _ = len(v1)
    }
    v3 -= v5
    var v35 uint8 = 201
    if (float64(-2.25) <= (0.0 * (1.5 - 3.14159))) {
        if (!true) {
            var v36 uint = uint(v3)
            var v37 int8 = 50
            _ = v36
            _ = v37
        } else {
            v38 := GenS1{F0: uint64(v5), F1: 4294967295, F2: -920}
            _ = v38
        }
        for fi39 := 0; fi39 < 3; fi39++ {
            v35 = v35
            if true {
                switch "foo bar" {
                case "":
                    var v40 uint32 = (^(v5))
                    _ = v40
                }
            }
            v5 = 433
            _ = fi39
        }
    }
    v5 = uint32(v35)
    if len(v1) > 0 {
        v1[0] = v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
