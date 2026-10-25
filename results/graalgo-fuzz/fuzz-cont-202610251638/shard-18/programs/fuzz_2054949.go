package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 int64; F2 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16; F1 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{42, 42}
    var v3 string = ""
    var v4 int32 = 131
    v5 := []int{1, -631, 32768}
    v6 := make(map[int]uint32)
    var v7 uint64 = 4294967295
    _ = len(v2)
    if (133 >= 699) {
        var v8 uint32 = 2147483648
        _ = v8
    }
    var v9 string = "a"
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS1{F0: 32767, F1: 32767, F2: 845}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        v13 := GenS2{F0: (v7 - v7)}
        if false {
            v14 := GenS1{F0: 1000, F1: 42, F2: 778}
            var v15 uint16 = 32767
            v3 = ""
            _ = v14
            _ = v15
        } else {
            v9 = ((v3 + (v3 + (v3 + ""))) + ("foo bar" + ""))
            v16 := GenS2{F0: 0}
            if ((485 >= 546) || (false || false)) {
                switch (fi12 + 32768) {
                case -1:
                    for fi17 := 0; fi17 < 18; fi17++ {
                        var v18 int = fi12
                        v19 := GenS1{F0: 32767, F1: 567, F2: -9223372036854775808}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                case 393:
                    v3 = v3
                case 32767:
                    if (false && true) {
                        var v20 uint64 = 42
                        if false {
                            _ = len(v6)
                        }
                        if ((v4 % (((v4 & v4)) | 1)) >= 2147483647) {
                            var v21 int16 = -604
                            v21 -= 1000
                            _ = v21
                        }
                        _ = v20
                    }
                }
                for fi22 := 0; fi22 < 11; fi22++ {
                    v9 = ""
                    fmt.Println(v9, fi22, v4)
                    _ = fi22
                }
                v3 = v9
            } else {
                for fi23 := 0; fi23 < 5; fi23++ {
                    switch fi23 {
                    case -9223372036854775808:
                        var v24 uint32 = 734
                        _ = v24
                    }
                    var v25 int = -75
                    _ = fi23
                    _ = v25
                }
                v26 := GenS3{F0: (fi12 << ((287) & 63)), F1: 0, F2: -9223372036854775808}
                _ = v26
            }
            _ = v16
        }
        v5 = append(v5, fi12)
        _ = fi12
        _ = v13
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        v28 := GenS4{F0: fi27}
        _ = fi27
        _ = v28
    }
    switch 128 {
    case 1:
        v4 += v4
    default:
        var v29 uint = 42
        _ = v29
    }
    switch "" {
    case "a":
        for fi30 := 0; fi30 < 18; fi30++ {
            v3 = v3
            v31 := GenS5{F0: false}
            fmt.Println(v7, v7)
            _ = fi30
            _ = v31
        }
    case "hello":
        for fi32 := 0; fi32 < 16; fi32++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v33 uint64 = v7
            if (((float64(-2.25) + 1.5) < ((((0.0 / 1.0) * -2.25) * 3.14159) + ((0.0 * 3.14159) / 0.1))) || true) {
                delete(v1, 42)
            } else {
                v4 = (v4 ^ (v4 % ((v4) | 1)))
                for fi34 := 0; fi34 < 7; fi34++ {
                    if ((("foo bar" + ((v3 + v9) + v3)) + v3) >= v3) {
                        v9 = (v9 + "a")
                    } else {
                        var v35 int8 = 100
                        v7 = 0
                        _ = v35
                    }
                    v36 := GenS6{F0: -1, F1: 1, F2: 0}
                    var v37 float64 = 3.14159
                    _ = fi34
                    _ = v36
                    _ = v37
                }
            }
            _ = fi32
            _ = v33
        }
    case "":
        var v38 bool = (424 == 18446744073709551615)
        _ = v38
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    for fi39 := 0; fi39 < 5; fi39++ {
        switch "hello" {
        case "hello":
            switch "foo bar" {
            case "\n\t":
                _ = len(v6)
            case "hello":
                if (false || ((!true) || true)) {
                    v1[2] = 858
                } else {
                    v7 = (v7 / (((v7 << ((65535) & 63))) | 1))
                }
            case "":
                v40 := GenS7{F0: 2147483648}
                _ = v40
            }
        case "\n\t":
            switch (fi39 - fi39) {
            case 32768:
                if (!((591 >= -73) && true)) {
                    var v41 int8 = 1
                    _ = v41
                } else {
                    switch v3 {
                    case "a":
                        v42 := GenS8{F0: -32768, F1: 18446744073709551615}
                        _ = v42
                    }
                }
            default:
                var v43 uint16 = 65535
                _ = v43
            }
        }
        switch "a" {
        case "foo bar":
            fmt.Println(v7, fi39, fi39)
        }
        _ = fi39
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v44 string = v3
    var v45 uint32 = 578
    var v46 uint32 = v45
    for fi47 := 0; fi47 < 15; fi47++ {
        var v48 float32 = -2.25
        _ = fi47
        _ = v48
    }
    for fi49 := 0; fi49 < 12; fi49++ {
        var v50 int32 = -2147483648
        switch (fi49 << ((228) & 63)) {
        case -543:
            switch (fi49 / ((fi49) | 1)) {
            case 42:
                v51 := v1[1]
                _ = v51
            case 1:
                fmt.Println(fi49, v46)
            default:
                v7 = v7
            }
        }
        _ = fi49
        _ = v50
    }
    fmt.Println(v3, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
