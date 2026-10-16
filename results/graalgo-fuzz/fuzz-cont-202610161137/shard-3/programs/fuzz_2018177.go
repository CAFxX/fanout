package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 bool, p1 uint32, p2 uint8) uint64 {
    var v7 int32 = 42
    v8 := GenS1{F0: 2, F1: "a"}
    var v9 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 uint64, p1 int8) int32 {
    switch -9223372036854775808 {
    case 32767:
        var v10 uint = 42
        _ = v10
    case 2147483647:
        v11 := []int{814, 32767}
        _ = v11
    case 4294967295:
        var v12 float64 = (float64(-2.25) / 1e10)
        _ = v12
    }
    v13 := []int{435, -2147483648, 127}
    v14 := GenS1{F0: 783, F1: ("" + "a")}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return -2147483648
}

func main() {
    var v1 int64 = -1
    var v2 uint16 = 255
    v3 := []string{"hello", (("\n\t" + "\n\t") + ("foo bar" + "a"))}
    var v4 uint16 = 2
    var v5 uint64 = 0
    var v6 uint8 = 255
    v4 -= (v4 + (^(v4)))
    v1 = 42
    var v15 int8 = -120
    var v16 uint32 = 65535
    var v17 int32 = 42
    var v18 uint8 = (^(v6))
    if false {
        for fi19 := 0; fi19 < 1; fi19++ {
            v6 -= v6
            _ = fi19
        }
        fmt.Println(v17)
    } else {
        fmt.Println(v4)
    }
    v20 := GenS1{F0: 42, F1: (("" + "foo bar") + "\n\t")}
    var v21 uint32 = v16
    if true {
        switch "hello" {
        case "a":
            if (-2147483648 >= -1) {
                if (false && (false || (!(-2147483648 == 2147483647)))) {
                    fmt.Println(v5, v5, v21)
                    v4 = (^(v2))
                } else {
                    v22 := GenS2{F0: (v15 - v15)}
                    fmt.Println(v17, v15)
                    var v23 float64 = 0.1
                    _ = v22
                    _ = v23
                }
            }
        case "hello":
            for fi24 := 0; fi24 < 16; fi24++ {
                v3 = append(v3, (("\n\t" + (("\n\t" + (("foo bar" + "foo bar") + ("hello" + "hello"))) + ("\n\t" + ""))) + "\n\t"))
                var v25 int64 = v1
                var v26 uint8 = v6
                _ = fi24
                _ = v25
                _ = v26
            }
        default:
            var v27 float32 = 1e10
            _ = v27
        }
    }
    switch -9223372036854775808 {
    case 32767:
        fmt.Println(v21)
    case 128:
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 string = "foo bar"
            if (4294967295 > uint(v18)) {
                v30 := GenS3{F0: -32768}
                var v31 int64 = v1
                _ = v30
                _ = v31
            } else {
                v32 := GenS1{F0: (v2 * 0), F1: "hello"}
                _ = len(v3)
                v33 := GenS3{F0: 42}
                _ = v32
                _ = v33
            }
            switch fi28 {
            case 431:
                v16 = v16
            case 128:
                v17 = genFn1((v5 | uint64(v18)), int8(v5))
            default:
                var v34 int8 = int8(v1)
                _ = v34
            }
            _ = fi28
            _ = v29
        }
    default:
        switch "" {
        case "":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    v5 += (v5 + (v5 - uint64(v18)))
    v15 = (v15 << ((255) & 63))
    v1 = v1
    for fi35 := 0; fi35 < 9; fi35++ {
        var v36 uint16 = (v4 % (((v2 - v4)) | 1))
        switch (fi35 >> ((18446744073709551615) & 63)) {
        case -2147483648:
            var v37 int = -9223372036854775808
            _ = v37
        default:
            v6 = (^(v6))
        }
        _ = fi35
        _ = v36
    }
    if ((v1 << ((703) & 63)) == int64(v5)) {
        var v38 float32 = 0.1
        var v39 uint32 = (v21 + (v21 - 2147483647))
        v2 = (v2 | (^(v4)))
        _ = v38
        _ = v39
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        v18 = v18
        if (false || (!(!false))) {
            var v41 float64 = (1.5 - 3.14159)
            _ = v41
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
